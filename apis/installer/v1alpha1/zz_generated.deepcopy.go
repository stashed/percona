// +build !ignore_autogenerated

/*
Copyright The Stash Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRef) DeepCopyInto(out *ImageRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRef.
func (in *ImageRef) DeepCopy() *ImageRef {
	if in == nil {
		return nil
	}
	out := new(ImageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBBackup) DeepCopyInto(out *PerconaXtraDBBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBBackup.
func (in *PerconaXtraDBBackup) DeepCopy() *PerconaXtraDBBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBRestore) DeepCopyInto(out *PerconaXtraDBRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBRestore.
func (in *PerconaXtraDBRestore) DeepCopy() *PerconaXtraDBRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashPerconaXtraDB) DeepCopyInto(out *StashPerconaXtraDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashPerconaXtraDB.
func (in *StashPerconaXtraDB) DeepCopy() *StashPerconaXtraDB {
	if in == nil {
		return nil
	}
	out := new(StashPerconaXtraDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashPerconaXtraDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashPerconaXtraDBList) DeepCopyInto(out *StashPerconaXtraDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StashPerconaXtraDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashPerconaXtraDBList.
func (in *StashPerconaXtraDBList) DeepCopy() *StashPerconaXtraDBList {
	if in == nil {
		return nil
	}
	out := new(StashPerconaXtraDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashPerconaXtraDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashPerconaXtraDBSpec) DeepCopyInto(out *StashPerconaXtraDBSpec) {
	*out = *in
	out.Image = in.Image
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashPerconaXtraDBSpec.
func (in *StashPerconaXtraDBSpec) DeepCopy() *StashPerconaXtraDBSpec {
	if in == nil {
		return nil
	}
	out := new(StashPerconaXtraDBSpec)
	in.DeepCopyInto(out)
	return out
}