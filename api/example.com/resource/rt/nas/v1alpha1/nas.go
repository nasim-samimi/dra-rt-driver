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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AllocatableGpu represents an allocatable GPU on a node.
type AllocatableCpu struct {
	ID   int `json:"id"`
	Util int `json:"util"`
	// ProductName string `json:"productName"` // let's assume that the UUID is enough for now
}

// AllocatableDevice represents an allocatable device on a node.
type AllocatableRtCpu struct {
	RtCpu *AllocatableCpu `json:"rtcpu,omitempty"`
}

// Type returns the type of AllocatableDevice this represents.
func (d AllocatableRtCpu) Type() string {
	if d.RtCpu != nil {
		return RtCpuType
	}
	return UnknownDeviceType
}

// AllocatedGpu represents an allocated GPU.
type AllocatedCpu struct {
	ID string `json:"id,omitempty"`
}

// AllocatedCpuset represents a set of allocated CPUs.
type AllocatedCpuset struct {
	Cpuset []AllocatedCpu `json:"cpuset"`
}

// AllocatedRtCpu represents a set of allocated CPUs.
type AllocatedRtCpu struct {
	RtCpu *AllocatedCpuset `json:"rtcpu,omitempty"`
}

// Type returns the type of AllocatedDevices this represents.
func (r AllocatedRtCpu) Type() string {
	if r.RtCpu != nil {
		return RtCpuType
	}
	return UnknownDeviceType
}

// PreparedGpu represents a prepared GPU on a node.
type PreparedCpu struct {
	ID string `json:"id"`
}

// PreparedGpus represents a set of prepared GPUs on a node.
type PreparedCpuset struct {
	Cpuset []PreparedCpu `json:"cpuset"`
}

// PreparedDevices represents a set of prepared devices on a node.
type PreparedRtCpu struct {
	RtCpu *PreparedCpuset `json:"rtcpu,omitempty"`
}

// Type returns the type of PreparedDevices this represents.
func (d PreparedRtCpu) Type() string {
	if d.RtCpu != nil {
		return RtCpuType
	}
	return UnknownDeviceType
}

// NodeAllocationStateSpec is the spec for the NodeAllocationState CRD.
type NodeAllocationStateSpec struct {
	AllocatableRtCpu []AllocatableRtCpu        `json:"allocatableRtCpu,omitempty"`
	AllocatedClaims  map[string]AllocatedRtCpu `json:"allocatedClaims,omitempty"`
	PreparedClaims   map[string]PreparedRtCpu  `json:"preparedClaims,omitempty"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:singular=nas

// NodeAllocationState holds the state required for allocation on a node.
type NodeAllocationState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeAllocationStateSpec `json:"spec,omitempty"`
	Status string                  `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeAllocationStateList represents the "plural" of a NodeAllocationState CRD object.
type NodeAllocationStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []NodeAllocationState `json:"items"`
}