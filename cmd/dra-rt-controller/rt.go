/*
 * Copyright 2023 The Kubernetes Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"sort"
	"strconv"

	nascrd "github.com/nasim-samimi/dra-rt-driver/api/example.com/resource/rt/nas/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	resourcev1 "k8s.io/api/resource/v1alpha2"
	"k8s.io/dynamic-resource-allocation/controller"

	rtcrd "github.com/nasim-samimi/dra-rt-driver/api/example.com/resource/rt/v1alpha1"
)

type rtdriver struct {
	PendingAllocatedClaims *PerNodeAllocatedClaims
}

func NewRtDriver() *rtdriver {
	return &rtdriver{
		PendingAllocatedClaims: NewPerNodeAllocatedClaims(),
	}
}

func (g *rtdriver) ValidateClaimParameters(claimParams *rtcrd.RtClaimParametersSpec) error {
	if claimParams.Count < 1 {
		return fmt.Errorf("invalid number of HCBS requested: %v", claimParams.Count)
	}
	return nil
}

func (g *rtdriver) Allocate(crd *nascrd.NodeAllocationState, claim *resourcev1.ResourceClaim, claimParams *rtcrd.RtClaimParametersSpec, class *resourcev1.ResourceClass, classParams *rtcrd.DeviceClassParametersSpec, selectedNode string) (OnSuccessCallback, error) {
	claimUID := string(claim.UID)

	if !g.PendingAllocatedClaims.Exists(claimUID, selectedNode) {
		return nil, fmt.Errorf("no allocations generated for claim '%v' on node '%v' yet", claim.UID, selectedNode)
	}

	crd.Spec.AllocatedClaims[claimUID] = g.PendingAllocatedClaims.Get(claimUID, selectedNode)
	onSuccess := func() {
		g.PendingAllocatedClaims.Remove(claimUID)
	}
	crd.Spec.AllocatedUtilToCpu = g.PendingAllocatedClaims.GetUtil(selectedNode)

	return onSuccess, nil
}

func (g *rtdriver) Deallocate(crd *nascrd.NodeAllocationState, claim *resourcev1.ResourceClaim) error {
	g.PendingAllocatedClaims.Remove(string(claim.UID))
	g.PendingAllocatedClaims.RemoveUtil(string(claim.UID))
	return nil
}

func (rt *rtdriver) UnsuitableNode(crd *nascrd.NodeAllocationState, pod *corev1.Pod, rtcas []*controller.ClaimAllocation, allcas []*controller.ClaimAllocation, potentialNode string) error {
	rt.PendingAllocatedClaims.VisitNode(potentialNode, func(claimUID string, allocation nascrd.AllocatedCpuset, utilisation nascrd.AllocatedUtilset, cgroups nascrd.PodCgroup) {
		if _, exists := crd.Spec.AllocatedClaims[claimUID]; exists {
			rt.PendingAllocatedClaims.Remove(claimUID)
		} else {
			crd.Spec.AllocatedClaims[claimUID] = allocation
			crd.Spec.AllocatedUtilToCpu = utilisation
			crd.Spec.AllocatedPodCgroups[crd.Spec.AllocatedClaims[claimUID].RtCpu.CgoupUID] = cgroups
		}
	})
	cgroupUID := cgroupUIDGenerator()
	for id, pcg := range crd.Spec.AllocatedPodCgroups {
		fmt.Println("id", id)
		fmt.Println("pcg", pcg)
	}

	allocated, allocatedUtil, podCgroup := rt.allocate(crd, pod, rtcas, allcas, potentialNode)
	fmt.Println("Allocated: ", podCgroup, "potentialNode: ", potentialNode)

	for _, ca := range rtcas {
		claimUID := string(ca.Claim.UID)
		claimParams, _ := ca.ClaimParameters.(*rtcrd.RtClaimParametersSpec)

		if claimParams.Count != len(allocated[claimUID]) {
			for _, ca := range allcas {
				ca.UnsuitableNodes = append(ca.UnsuitableNodes, potentialNode)
			}
			return nil
		} // it puts everything on only one node

		var devices []nascrd.AllocatedCpu
		for _, cpu := range allocated[claimUID] {
			device := cpu
			devices = append(devices, device)
		}

		allocatedDevices := nascrd.AllocatedCpuset{
			RtCpu: &nascrd.AllocatedRtCpu{
				Cpuset:   devices,
				CgoupUID: cgroupUID,
			},
		}

		allocatedUtilisations := nascrd.AllocatedUtilset{
			Cpus: allocatedUtil,
		}

		rt.PendingAllocatedClaims.Set(claimUID, potentialNode, allocatedDevices)
		rt.PendingAllocatedClaims.SetUtil(potentialNode, allocatedUtilisations)
		rt.PendingAllocatedClaims.SetCgroup(cgroupUID, potentialNode, podCgroup)
	}

	return nil
}

func (rt *rtdriver) allocate(crd *nascrd.NodeAllocationState, pod *corev1.Pod, cpucas []*controller.ClaimAllocation, allcas []*controller.ClaimAllocation, node string) (map[string][]nascrd.AllocatedCpu, map[string]nascrd.AllocatedUtil, nascrd.PodCgroup) {
	available := make(map[int]*nascrd.AllocatableCpu)
	util := crd.Spec.AllocatedUtilToCpu.Cpus
	// util := make(map[string]nascrd.AllocatedUtil)
	allocated := make(map[string][]nascrd.AllocatedCpu)
	containerCG := make(map[string]nascrd.ContainerCgroup)

	for _, device := range crd.Spec.AllocatableCpuset {
		switch device.Type() {
		case nascrd.RtCpuType:
			available[device.RtCpu.ID] = device.RtCpu
		default:
			// skip other devices
		}
	}

	for _, ca := range cpucas {
		claimUID := string(ca.Claim.UID)
		if _, exists := crd.Spec.AllocatedClaims[claimUID]; exists {
			devices := crd.Spec.AllocatedClaims[claimUID].RtCpu.Cpuset
			for _, device := range devices {
				allocated[claimUID] = append(allocated[claimUID], device)
			}
			continue
		}

		claimParams, _ := ca.ClaimParameters.(*rtcrd.RtClaimParametersSpec)
		claimUtil := (claimParams.Runtime * 1000 / claimParams.Period)
		var devices []nascrd.AllocatedCpu
		for i := 0; i < claimParams.Count; i++ {
			// for _, device := range available {
			worstFitCpus := cpuPartitioning(util, claimUtil, 1, "worstFit") //must get the policy from the user
			if worstFitCpus == nil {
				return nil, nil, nascrd.PodCgroup{}
			}
			worstFitCpusStr, _ := strconv.Atoi(worstFitCpus[0])
			d := nascrd.AllocatedCpu{
				ID:      worstFitCpusStr,
				Runtime: claimParams.Runtime,
				Period:  claimParams.Period,
			}
			util[strconv.Itoa(d.ID)] = nascrd.AllocatedUtil{
				Util: util[strconv.Itoa(d.ID)].Util + claimUtil,
			}
			if util[strconv.Itoa(d.ID)].Util >= 1000 {
				delete(available, d.ID)
			}
			devices = append(devices, d)

			break
		}
		allocated[claimUID] = devices
		fmt.Println("podUID", pod.UID)

		rt.containerCgroups(containerCG, devices, ca.PodClaimName, pod)
		for c, claims := range containerCG {
			fmt.Println("containername", c)
			for typed, cgroup := range claims {
				fmt.Println("cgroup", cgroup)
				fmt.Println("typed", typed)
			}
		}
		fmt.Println("claimlabels", ca.Claim.Labels)
	}
	podCG := rt.podCgroups(containerCG, crd, pod)
	// crd.Spec.AllocatedUtilToCpu = nascrd.AllocatedUtilset{
	// 	Cpus: util,
	// }

	return allocated, util, podCG
}

func cpuPartitioning(spec map[string]nascrd.AllocatedUtil, reqUtil int, reqCpus int, policy string) []string {
	type scoredCpu struct {
		cpu   string
		score int
	}
	var scoredCpus []scoredCpu
	for id, cpuinfo := range spec {
		score := 1000 - cpuinfo.Util - reqUtil
		if score > 0 {
			scoredCpus = append(scoredCpus, scoredCpu{
				cpu:   id,
				score: score,
			})
		}
	}

	if int(len(scoredCpus)) < reqCpus {
		return nil
	}
	switch policy {
	case "worstFit":
		sort.SliceStable(scoredCpus, func(i, j int) bool {
			if scoredCpus[i].score > scoredCpus[j].score {
				return true
			}
			return false
		})
	case "bestFit":
		sort.SliceStable(scoredCpus, func(i, j int) bool {
			if scoredCpus[i].score < scoredCpus[j].score {
				return true
			}
			return false
		})
	default:
		sort.SliceStable(scoredCpus, func(i, j int) bool {
			if scoredCpus[i].score > scoredCpus[j].score {
				return true
			}
			return false
		}) //default is worstFit
	}

	var fittingCpus []string
	for i := int(0); i < reqCpus; i++ {
		fittingCpus = append(fittingCpus, scoredCpus[i].cpu)
	}

	return fittingCpus
}
