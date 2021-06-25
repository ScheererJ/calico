// +build !ignore_autogenerated

// Copyright (c) 2016-2021 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	projectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	numorstring "github.com/projectcalico/api/pkg/lib/numorstring"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationAttribute) DeepCopyInto(out *AllocationAttribute) {
	*out = *in
	if in.AttrPrimary != nil {
		in, out := &in.AttrPrimary, &out.AttrPrimary
		*out = new(string)
		**out = **in
	}
	if in.AttrSecondary != nil {
		in, out := &in.AttrSecondary, &out.AttrSecondary
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationAttribute.
func (in *AllocationAttribute) DeepCopy() *AllocationAttribute {
	if in == nil {
		return nil
	}
	out := new(AllocationAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockAffinity) DeepCopyInto(out *BlockAffinity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockAffinity.
func (in *BlockAffinity) DeepCopy() *BlockAffinity {
	if in == nil {
		return nil
	}
	out := new(BlockAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockAffinity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockAffinityList) DeepCopyInto(out *BlockAffinityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockAffinity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockAffinityList.
func (in *BlockAffinityList) DeepCopy() *BlockAffinityList {
	if in == nil {
		return nil
	}
	out := new(BlockAffinityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockAffinityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockAffinitySpec) DeepCopyInto(out *BlockAffinitySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockAffinitySpec.
func (in *BlockAffinitySpec) DeepCopy() *BlockAffinitySpec {
	if in == nil {
		return nil
	}
	out := new(BlockAffinitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMBlock) DeepCopyInto(out *IPAMBlock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMBlock.
func (in *IPAMBlock) DeepCopy() *IPAMBlock {
	if in == nil {
		return nil
	}
	out := new(IPAMBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMBlock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMBlockList) DeepCopyInto(out *IPAMBlockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAMBlock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMBlockList.
func (in *IPAMBlockList) DeepCopy() *IPAMBlockList {
	if in == nil {
		return nil
	}
	out := new(IPAMBlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMBlockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMBlockSpec) DeepCopyInto(out *IPAMBlockSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(string)
		**out = **in
	}
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make([]*int, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(int)
				**out = **in
			}
		}
	}
	if in.Unallocated != nil {
		in, out := &in.Unallocated, &out.Unallocated
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]AllocationAttribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMBlockSpec.
func (in *IPAMBlockSpec) DeepCopy() *IPAMBlockSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMConfig) DeepCopyInto(out *IPAMConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMConfig.
func (in *IPAMConfig) DeepCopy() *IPAMConfig {
	if in == nil {
		return nil
	}
	out := new(IPAMConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMConfigList) DeepCopyInto(out *IPAMConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAMConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMConfigList.
func (in *IPAMConfigList) DeepCopy() *IPAMConfigList {
	if in == nil {
		return nil
	}
	out := new(IPAMConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMConfigSpec) DeepCopyInto(out *IPAMConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMConfigSpec.
func (in *IPAMConfigSpec) DeepCopy() *IPAMConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMHandle) DeepCopyInto(out *IPAMHandle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMHandle.
func (in *IPAMHandle) DeepCopy() *IPAMHandle {
	if in == nil {
		return nil
	}
	out := new(IPAMHandle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMHandle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMHandleList) DeepCopyInto(out *IPAMHandleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAMHandle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMHandleList.
func (in *IPAMHandleList) DeepCopy() *IPAMHandleList {
	if in == nil {
		return nil
	}
	out := new(IPAMHandleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMHandleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMHandleSpec) DeepCopyInto(out *IPAMHandleSpec) {
	*out = *in
	if in.Block != nil {
		in, out := &in.Block, &out.Block
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMHandleSpec.
func (in *IPAMHandleSpec) DeepCopy() *IPAMHandleSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMHandleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPNAT) DeepCopyInto(out *IPNAT) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPNAT.
func (in *IPNAT) DeepCopy() *IPNAT {
	if in == nil {
		return nil
	}
	out := new(IPNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAddress) DeepCopyInto(out *NodeAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAddress.
func (in *NodeAddress) DeepCopy() *NodeAddress {
	if in == nil {
		return nil
	}
	out := new(NodeAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeBGPSpec) DeepCopyInto(out *NodeBGPSpec) {
	*out = *in
	if in.ASNumber != nil {
		in, out := &in.ASNumber, &out.ASNumber
		*out = new(numorstring.ASNumber)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeBGPSpec.
func (in *NodeBGPSpec) DeepCopy() *NodeBGPSpec {
	if in == nil {
		return nil
	}
	out := new(NodeBGPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.BGP != nil {
		in, out := &in.BGP, &out.BGP
		*out = new(NodeBGPSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OrchRefs != nil {
		in, out := &in.OrchRefs, &out.OrchRefs
		*out = make([]OrchRef, len(*in))
		copy(*out, *in)
	}
	if in.Wireguard != nil {
		in, out := &in.Wireguard, &out.Wireguard
		*out = new(NodeWireguardSpec)
		**out = **in
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]NodeAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	if in.PodCIDRs != nil {
		in, out := &in.PodCIDRs, &out.PodCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeWireguardSpec) DeepCopyInto(out *NodeWireguardSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeWireguardSpec.
func (in *NodeWireguardSpec) DeepCopy() *NodeWireguardSpec {
	if in == nil {
		return nil
	}
	out := new(NodeWireguardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchRef) DeepCopyInto(out *OrchRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchRef.
func (in *OrchRef) DeepCopy() *OrchRef {
	if in == nil {
		return nil
	}
	out := new(OrchRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpoint) DeepCopyInto(out *WorkloadEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpoint.
func (in *WorkloadEndpoint) DeepCopy() *WorkloadEndpoint {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointList) DeepCopyInto(out *WorkloadEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointList.
func (in *WorkloadEndpointList) DeepCopy() *WorkloadEndpointList {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointSpec) DeepCopyInto(out *WorkloadEndpointSpec) {
	*out = *in
	if in.IPNetworks != nil {
		in, out := &in.IPNetworks, &out.IPNetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPNATs != nil {
		in, out := &in.IPNATs, &out.IPNATs
		*out = make([]IPNAT, len(*in))
		copy(*out, *in)
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]projectcalicov3.EndpointPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointSpec.
func (in *WorkloadEndpointSpec) DeepCopy() *WorkloadEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointSpec)
	in.DeepCopyInto(out)
	return out
}