// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jaeger) DeepCopyInto(out *Jaeger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jaeger.
func (in *Jaeger) DeepCopy() *Jaeger {
	if in == nil {
		return nil
	}
	out := new(Jaeger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jaeger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAgentSpec) DeepCopyInto(out *JaegerAgentSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAgentSpec.
func (in *JaegerAgentSpec) DeepCopy() *JaegerAgentSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAllInOneSpec) DeepCopyInto(out *JaegerAllInOneSpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAllInOneSpec.
func (in *JaegerAllInOneSpec) DeepCopy() *JaegerAllInOneSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAllInOneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCollectorSpec) DeepCopyInto(out *JaegerCollectorSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCollectorSpec.
func (in *JaegerCollectorSpec) DeepCopy() *JaegerCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngressSpec) DeepCopyInto(out *JaegerIngressSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngressSpec.
func (in *JaegerIngressSpec) DeepCopy() *JaegerIngressSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerList) DeepCopyInto(out *JaegerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jaeger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerList.
func (in *JaegerList) DeepCopy() *JaegerList {
	if in == nil {
		return nil
	}
	out := new(JaegerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JaegerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerQuerySpec) DeepCopyInto(out *JaegerQuerySpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerQuerySpec.
func (in *JaegerQuerySpec) DeepCopy() *JaegerQuerySpec {
	if in == nil {
		return nil
	}
	out := new(JaegerQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerSpec) DeepCopyInto(out *JaegerSpec) {
	*out = *in
	in.AllInOne.DeepCopyInto(&out.AllInOne)
	in.Query.DeepCopyInto(&out.Query)
	in.Collector.DeepCopyInto(&out.Collector)
	in.Agent.DeepCopyInto(&out.Agent)
	in.Storage.DeepCopyInto(&out.Storage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerSpec.
func (in *JaegerSpec) DeepCopy() *JaegerSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStatus) DeepCopyInto(out *JaegerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStatus.
func (in *JaegerStatus) DeepCopy() *JaegerStatus {
	if in == nil {
		return nil
	}
	out := new(JaegerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStorageSpec) DeepCopyInto(out *JaegerStorageSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStorageSpec.
func (in *JaegerStorageSpec) DeepCopy() *JaegerStorageSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	if in.opts != nil {
		in, out := &in.opts, &out.opts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}
