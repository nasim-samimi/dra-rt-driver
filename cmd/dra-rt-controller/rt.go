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

	corev1 "k8s.io/api/core/v1"
	resourcev1 "k8s.io/api/resource/v1alpha2"
	"k8s.io/dynamic-resource-allocation/controller"

	nascrd "github.com/nasim-samimi/dra-rt-driver/api/example.com/resource/rt/nas/v1alpha1"
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
		return fmt.Errorf("invalid number of GPUs requested: %v", claimParams.Count)
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

	return onSuccess, nil
}

func (g *rtdriver) Deallocate(crd *nascrd.NodeAllocationState, claim *resourcev1.ResourceClaim) error {
	g.PendingAllocatedClaims.Remove(string(claim.UID))
	return nil
}

func (rt *rtdriver) UnsuitableNode(crd *nascrd.NodeAllocationState, pod *corev1.Pod, rtcas []*controller.ClaimAllocation, allcas []*controller.ClaimAllocation, potentialNode string) error {
	rt.PendingAllocatedClaims.VisitNode(potentialNode, func(claimUID string, allocation nascrd.AllocatedRtCpu) {
		if _, exists := crd.Spec.AllocatedClaims[claimUID]; exists {
			rt.PendingAllocatedClaims.Remove(claimUID)
		} else {
			crd.Spec.AllocatedClaims[claimUID] = allocation
		}
	})

	allocated := rt.allocate(crd, pod, rtcas, allcas, potentialNode)
	for _, ca := range rtcas {
		claimUID := string(ca.Claim.UID)
		claimParams, _ := ca.ClaimParameters.(*rtcrd.RtClaimParametersSpec)

		if claimParams.Count != len(allocated[claimUID]) {
			for _, ca := range allcas {
				ca.UnsuitableNodes = append(ca.UnsuitableNodes, potentialNode)
			}
			return nil
		}

		var devices []nascrd.AllocatedCpu
		for _, gpu := range allocated[claimUID] {
			device := nascrd.AllocatedCpu{
				ID: gpu,
			}
			devices = append(devices, device)
		}

		allocatedDevices := nascrd.AllocatedRtCpu{
			RtCpu: &nascrd.AllocatedCpuset{
				Cpuset: devices,
			},
		}

		rt.PendingAllocatedClaims.Set(claimUID, potentialNode, allocatedDevices)
	}

	return nil
}

func (g *rtdriver) allocate(crd *nascrd.NodeAllocationState, pod *corev1.Pod, gpucas []*controller.ClaimAllocation, allcas []*controller.ClaimAllocation, node string) map[string][]string {
	available := make(map[string]*nascrd.AllocatableCpu)

	for _, device := range crd.Spec.AllocatableRtCpu {
		switch device.Type() {
		case nascrd.RtCpuType:
			available[string(device.RtCpu.ID)] = device.RtCpu
		default:
			// skip other devices
		}
	}

	for _, allocation := range crd.Spec.AllocatedClaims {
		switch allocation.Type() {
		case nascrd.RtCpuType:
			for _, device := range allocation.RtCpu.Cpuset {
				delete(available, device.ID)
			}
		default:
			// skip other devices
		}
	}

	allocated := make(map[string][]string)
	for _, ca := range gpucas {
		claimUID := string(ca.Claim.UID)
		if _, exists := crd.Spec.AllocatedClaims[claimUID]; exists {
			devices := crd.Spec.AllocatedClaims[claimUID].RtCpu.Cpuset
			for _, device := range devices {
				allocated[claimUID] = append(allocated[claimUID], device.ID)
			}
			continue
		}

		claimParams, _ := ca.ClaimParameters.(*rtcrd.RtClaimParametersSpec)
		var devices []string
		for i := 0; i < claimParams.Count; i++ {
			for _, device := range available {
				devices = append(devices, string(device.ID))
				delete(available, string(device.ID))
				break
			}
		}
		allocated[claimUID] = devices
	}

	return allocated
}