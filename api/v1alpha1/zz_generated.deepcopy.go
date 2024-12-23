//go:build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and nginx-ingress-controller-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressController) DeepCopyInto(out *NginxIngressController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressController.
func (in *NginxIngressController) DeepCopy() *NginxIngressController {
	if in == nil {
		return nil
	}
	out := new(NginxIngressController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NginxIngressController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerList) DeepCopyInto(out *NginxIngressControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NginxIngressController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerList.
func (in *NginxIngressControllerList) DeepCopy() *NginxIngressControllerList {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NginxIngressControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerSpec) DeepCopyInto(out *NginxIngressControllerSpec) {
	*out = *in
	in.JSON.DeepCopyInto(&out.JSON)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerSpec.
func (in *NginxIngressControllerSpec) DeepCopy() *NginxIngressControllerSpec {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerStatus) DeepCopyInto(out *NginxIngressControllerStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerStatus.
func (in *NginxIngressControllerStatus) DeepCopy() *NginxIngressControllerStatus {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerStatus)
	in.DeepCopyInto(out)
	return out
}
