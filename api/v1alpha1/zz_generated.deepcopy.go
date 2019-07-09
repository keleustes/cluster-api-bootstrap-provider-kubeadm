// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmBootstrapConfig) DeepCopyInto(out *KubeadmBootstrapConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmBootstrapConfig.
func (in *KubeadmBootstrapConfig) DeepCopy() *KubeadmBootstrapConfig {
	if in == nil {
		return nil
	}
	out := new(KubeadmBootstrapConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmBootstrapConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmBootstrapConfigList) DeepCopyInto(out *KubeadmBootstrapConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmBootstrapConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmBootstrapConfigList.
func (in *KubeadmBootstrapConfigList) DeepCopy() *KubeadmBootstrapConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeadmBootstrapConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmBootstrapConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmBootstrapConfigSpec) DeepCopyInto(out *KubeadmBootstrapConfigSpec) {
	*out = *in
	in.ClusterConfiguration.DeepCopyInto(&out.ClusterConfiguration)
	in.InitConfiguration.DeepCopyInto(&out.InitConfiguration)
	in.JoinConfiguration.DeepCopyInto(&out.JoinConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmBootstrapConfigSpec.
func (in *KubeadmBootstrapConfigSpec) DeepCopy() *KubeadmBootstrapConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmBootstrapConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmBootstrapConfigStatus) DeepCopyInto(out *KubeadmBootstrapConfigStatus) {
	*out = *in
	if in.BootstrapData != nil {
		in, out := &in.BootstrapData, &out.BootstrapData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmBootstrapConfigStatus.
func (in *KubeadmBootstrapConfigStatus) DeepCopy() *KubeadmBootstrapConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmBootstrapConfigStatus)
	in.DeepCopyInto(out)
	return out
}
