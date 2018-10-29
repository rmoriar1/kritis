// +build !ignore_autogenerated

/*
Copyright 2018 Google LLC

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestationAuthority) DeepCopyInto(out *AttestationAuthority) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestationAuthority.
func (in *AttestationAuthority) DeepCopy() *AttestationAuthority {
	if in == nil {
		return nil
	}
	out := new(AttestationAuthority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestationAuthority) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestationAuthorityList) DeepCopyInto(out *AttestationAuthorityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AttestationAuthority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestationAuthorityList.
func (in *AttestationAuthorityList) DeepCopy() *AttestationAuthorityList {
	if in == nil {
		return nil
	}
	out := new(AttestationAuthorityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestationAuthorityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestationAuthoritySpec) DeepCopyInto(out *AttestationAuthoritySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestationAuthoritySpec.
func (in *AttestationAuthoritySpec) DeepCopy() *AttestationAuthoritySpec {
	if in == nil {
		return nil
	}
	out := new(AttestationAuthoritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPolicy) DeepCopyInto(out *BuildPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPolicy.
func (in *BuildPolicy) DeepCopy() *BuildPolicy {
	if in == nil {
		return nil
	}
	out := new(BuildPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPolicyList) DeepCopyInto(out *BuildPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPolicyList.
func (in *BuildPolicyList) DeepCopy() *BuildPolicyList {
	if in == nil {
		return nil
	}
	out := new(BuildPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPolicySpec) DeepCopyInto(out *BuildPolicySpec) {
	*out = *in
	out.BuildRequirements = in.BuildRequirements
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPolicySpec.
func (in *BuildPolicySpec) DeepCopy() *BuildPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BuildPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildRequirements) DeepCopyInto(out *BuildRequirements) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildRequirements.
func (in *BuildRequirements) DeepCopy() *BuildRequirements {
	if in == nil {
		return nil
	}
	out := new(BuildRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSecurityPolicy) DeepCopyInto(out *ImageSecurityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSecurityPolicy.
func (in *ImageSecurityPolicy) DeepCopy() *ImageSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(ImageSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSecurityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSecurityPolicyList) DeepCopyInto(out *ImageSecurityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageSecurityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSecurityPolicyList.
func (in *ImageSecurityPolicyList) DeepCopy() *ImageSecurityPolicyList {
	if in == nil {
		return nil
	}
	out := new(ImageSecurityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSecurityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSecurityPolicySpec) DeepCopyInto(out *ImageSecurityPolicySpec) {
	*out = *in
	if in.ImageWhitelist != nil {
		in, out := &in.ImageWhitelist, &out.ImageWhitelist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PackageVulnerabilityRequirements.DeepCopyInto(&out.PackageVulnerabilityRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSecurityPolicySpec.
func (in *ImageSecurityPolicySpec) DeepCopy() *ImageSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ImageSecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KritisConfig) DeepCopyInto(out *KritisConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KritisConfig.
func (in *KritisConfig) DeepCopy() *KritisConfig {
	if in == nil {
		return nil
	}
	out := new(KritisConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KritisConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KritisConfigList) DeepCopyInto(out *KritisConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KritisConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KritisConfigList.
func (in *KritisConfigList) DeepCopy() *KritisConfigList {
	if in == nil {
		return nil
	}
	out := new(KritisConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KritisConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVulnerabilityRequirements) DeepCopyInto(out *PackageVulnerabilityRequirements) {
	*out = *in
	if in.WhitelistCVEs != nil {
		in, out := &in.WhitelistCVEs, &out.WhitelistCVEs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVulnerabilityRequirements.
func (in *PackageVulnerabilityRequirements) DeepCopy() *PackageVulnerabilityRequirements {
	if in == nil {
		return nil
	}
	out := new(PackageVulnerabilityRequirements)
	in.DeepCopyInto(out)
	return out
}
