//go:build !ignore_autogenerated

/*
 * Copyright 2024 The Kubernetes Authors.
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableCpu) DeepCopyInto(out *AllocatableCpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableCpu.
func (in *AllocatableCpu) DeepCopy() *AllocatableCpu {
	if in == nil {
		return nil
	}
	out := new(AllocatableCpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatableCpuset) DeepCopyInto(out *AllocatableCpuset) {
	*out = *in
	if in.RtCpu != nil {
		in, out := &in.RtCpu, &out.RtCpu
		*out = new(AllocatableCpu)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatableCpuset.
func (in *AllocatableCpuset) DeepCopy() *AllocatableCpuset {
	if in == nil {
		return nil
	}
	out := new(AllocatableCpuset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedCpu) DeepCopyInto(out *AllocatedCpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedCpu.
func (in *AllocatedCpu) DeepCopy() *AllocatedCpu {
	if in == nil {
		return nil
	}
	out := new(AllocatedCpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedCpuset) DeepCopyInto(out *AllocatedCpuset) {
	*out = *in
	if in.RtCpu != nil {
		in, out := &in.RtCpu, &out.RtCpu
		*out = new(AllocatedRtCpu)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedCpuset.
func (in *AllocatedCpuset) DeepCopy() *AllocatedCpuset {
	if in == nil {
		return nil
	}
	out := new(AllocatedCpuset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedRtCpu) DeepCopyInto(out *AllocatedRtCpu) {
	*out = *in
	if in.Cpuset != nil {
		in, out := &in.Cpuset, &out.Cpuset
		*out = make([]AllocatedCpu, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedRtCpu.
func (in *AllocatedRtCpu) DeepCopy() *AllocatedRtCpu {
	if in == nil {
		return nil
	}
	out := new(AllocatedRtCpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedUtil) DeepCopyInto(out *AllocatedUtil) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedUtil.
func (in *AllocatedUtil) DeepCopy() *AllocatedUtil {
	if in == nil {
		return nil
	}
	out := new(AllocatedUtil)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedUtilset) DeepCopyInto(out *AllocatedUtilset) {
	*out = *in
	if in.Cpus != nil {
		in, out := &in.Cpus, &out.Cpus
		*out = make(MappedUtil, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedUtilset.
func (in *AllocatedUtilset) DeepCopy() *AllocatedUtilset {
	if in == nil {
		return nil
	}
	out := new(AllocatedUtilset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimCgroup) DeepCopyInto(out *ClaimCgroup) {
	*out = *in
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = make(MappedCgroup, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ContainerPeriod != nil {
		in, out := &in.ContainerPeriod, &out.ContainerPeriod
		*out = make(MappedCgroup, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimCgroup.
func (in *ClaimCgroup) DeepCopy() *ClaimCgroup {
	if in == nil {
		return nil
	}
	out := new(ClaimCgroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ContainerCgroup) DeepCopyInto(out *ContainerCgroup) {
	{
		in := &in
		*out = make(ContainerCgroup, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerCgroup.
func (in ContainerCgroup) DeepCopy() ContainerCgroup {
	if in == nil {
		return nil
	}
	out := new(ContainerCgroup)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MappedCgroup) DeepCopyInto(out *MappedCgroup) {
	{
		in := &in
		*out = make(MappedCgroup, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappedCgroup.
func (in MappedCgroup) DeepCopy() MappedCgroup {
	if in == nil {
		return nil
	}
	out := new(MappedCgroup)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MappedUtil) DeepCopyInto(out *MappedUtil) {
	{
		in := &in
		*out = make(MappedUtil, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappedUtil.
func (in MappedUtil) DeepCopy() MappedUtil {
	if in == nil {
		return nil
	}
	out := new(MappedUtil)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationState) DeepCopyInto(out *NodeAllocationState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationState.
func (in *NodeAllocationState) DeepCopy() *NodeAllocationState {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateConfig) DeepCopyInto(out *NodeAllocationStateConfig) {
	*out = *in
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(v1.OwnerReference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateConfig.
func (in *NodeAllocationStateConfig) DeepCopy() *NodeAllocationStateConfig {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateList) DeepCopyInto(out *NodeAllocationStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeAllocationState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateList.
func (in *NodeAllocationStateList) DeepCopy() *NodeAllocationStateList {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAllocationStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAllocationStateSpec) DeepCopyInto(out *NodeAllocationStateSpec) {
	*out = *in
	if in.AllocatableCpuset != nil {
		in, out := &in.AllocatableCpuset, &out.AllocatableCpuset
		*out = make([]AllocatableCpuset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllocatedClaims != nil {
		in, out := &in.AllocatedClaims, &out.AllocatedClaims
		*out = make(map[string]AllocatedCpuset, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PreparedClaims != nil {
		in, out := &in.PreparedClaims, &out.PreparedClaims
		*out = make(map[string]PreparedCpuset, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.AllocatedUtilToCpu.DeepCopyInto(&out.AllocatedUtilToCpu)
	if in.AllocatedPodCgroups != nil {
		in, out := &in.AllocatedPodCgroups, &out.AllocatedPodCgroups
		*out = make(map[string]PodCgroup, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAllocationStateSpec.
func (in *NodeAllocationStateSpec) DeepCopy() *NodeAllocationStateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeAllocationStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodCgroup) DeepCopyInto(out *PodCgroup) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]ContainerCgroup, len(*in))
		for key, val := range *in {
			var outVal map[string]ClaimCgroup
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(ContainerCgroup, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodCgroup.
func (in *PodCgroup) DeepCopy() *PodCgroup {
	if in == nil {
		return nil
	}
	out := new(PodCgroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreparedCpu) DeepCopyInto(out *PreparedCpu) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreparedCpu.
func (in *PreparedCpu) DeepCopy() *PreparedCpu {
	if in == nil {
		return nil
	}
	out := new(PreparedCpu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreparedCpuset) DeepCopyInto(out *PreparedCpuset) {
	*out = *in
	if in.RtCpu != nil {
		in, out := &in.RtCpu, &out.RtCpu
		*out = new(PreparedRtCpu)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreparedCpuset.
func (in *PreparedCpuset) DeepCopy() *PreparedCpuset {
	if in == nil {
		return nil
	}
	out := new(PreparedCpuset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreparedRtCpu) DeepCopyInto(out *PreparedRtCpu) {
	*out = *in
	if in.Cpuset != nil {
		in, out := &in.Cpuset, &out.Cpuset
		*out = make([]PreparedCpu, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreparedRtCpu.
func (in *PreparedRtCpu) DeepCopy() *PreparedRtCpu {
	if in == nil {
		return nil
	}
	out := new(PreparedRtCpu)
	in.DeepCopyInto(out)
	return out
}
