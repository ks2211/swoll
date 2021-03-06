// +build !ignore_autogenerated

/*
Copyright 2020 Critical Stack, LLC

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trace) DeepCopyInto(out *Trace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trace.
func (in *Trace) DeepCopy() *Trace {
	if in == nil {
		return nil
	}
	out := new(Trace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceList) DeepCopyInto(out *TraceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceList.
func (in *TraceList) DeepCopy() *TraceList {
	if in == nil {
		return nil
	}
	out := new(TraceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TraceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceSpec) DeepCopyInto(out *TraceSpec) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.FieldSelector.DeepCopyInto(&out.FieldSelector)
	if in.HostSelector != nil {
		in, out := &in.HostSelector, &out.HostSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Syscalls != nil {
		in, out := &in.Syscalls, &out.Syscalls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceSpec.
func (in *TraceSpec) DeepCopy() *TraceSpec {
	if in == nil {
		return nil
	}
	out := new(TraceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceStatus) DeepCopyInto(out *TraceStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceStatus.
func (in *TraceStatus) DeepCopy() *TraceStatus {
	if in == nil {
		return nil
	}
	out := new(TraceStatus)
	in.DeepCopyInto(out)
	return out
}
