// +build !ignore_autogenerated

// Copyright 2019 Machine Zone, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSecret) DeepCopyInto(out *ConfigMapSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSecret.
func (in *ConfigMapSecret) DeepCopy() *ConfigMapSecret {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigMapSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSecretCondition) DeepCopyInto(out *ConfigMapSecretCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSecretCondition.
func (in *ConfigMapSecretCondition) DeepCopy() *ConfigMapSecretCondition {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSecretCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSecretList) DeepCopyInto(out *ConfigMapSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigMapSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSecretList.
func (in *ConfigMapSecretList) DeepCopy() *ConfigMapSecretList {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigMapSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSecretSpec) DeepCopyInto(out *ConfigMapSecretSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]TemplateVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSecretSpec.
func (in *ConfigMapSecretSpec) DeepCopy() *ConfigMapSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapSecretStatus) DeepCopyInto(out *ConfigMapSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConfigMapSecretCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapSecretStatus.
func (in *ConfigMapSecretStatus) DeepCopy() *ConfigMapSecretStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigMapSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapTemplate) DeepCopyInto(out *ConfigMapTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapTemplate.
func (in *ConfigMapTemplate) DeepCopy() *ConfigMapTemplate {
	if in == nil {
		return nil
	}
	out := new(ConfigMapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateVariable) DeepCopyInto(out *TemplateVariable) {
	*out = *in
	if in.SecretValue != nil {
		in, out := &in.SecretValue, &out.SecretValue
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapValue != nil {
		in, out := &in.ConfigMapValue, &out.ConfigMapValue
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateVariable.
func (in *TemplateVariable) DeepCopy() *TemplateVariable {
	if in == nil {
		return nil
	}
	out := new(TemplateVariable)
	in.DeepCopyInto(out)
	return out
}