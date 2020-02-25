// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRLink) DeepCopyInto(out *CRLink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRLink.
func (in *CRLink) DeepCopy() *CRLink {
	if in == nil {
		return nil
	}
	out := new(CRLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseModulesList) DeepCopyInto(out *DatabaseModulesList) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseModulesList.
func (in *DatabaseModulesList) DeepCopy() *DatabaseModulesList {
	if in == nil {
		return nil
	}
	out := new(DatabaseModulesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabase) DeepCopyInto(out *PostgresqlDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabase.
func (in *PostgresqlDatabase) DeepCopy() *PostgresqlDatabase {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseList) DeepCopyInto(out *PostgresqlDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseList.
func (in *PostgresqlDatabaseList) DeepCopy() *PostgresqlDatabaseList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseSpec) DeepCopyInto(out *PostgresqlDatabaseSpec) {
	*out = *in
	in.Schemas.DeepCopyInto(&out.Schemas)
	in.Extensions.DeepCopyInto(&out.Extensions)
	if in.EngineConfiguration != nil {
		in, out := &in.EngineConfiguration, &out.EngineConfiguration
		*out = new(CRLink)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseSpec.
func (in *PostgresqlDatabaseSpec) DeepCopy() *PostgresqlDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseStatus) DeepCopyInto(out *PostgresqlDatabaseStatus) {
	*out = *in
	out.Roles = in.Roles
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseStatus.
func (in *PostgresqlDatabaseStatus) DeepCopy() *PostgresqlDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfiguration) DeepCopyInto(out *PostgresqlEngineConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfiguration.
func (in *PostgresqlEngineConfiguration) DeepCopy() *PostgresqlEngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlEngineConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationList) DeepCopyInto(out *PostgresqlEngineConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlEngineConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationList.
func (in *PostgresqlEngineConfigurationList) DeepCopy() *PostgresqlEngineConfigurationList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlEngineConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationSpec) DeepCopyInto(out *PostgresqlEngineConfigurationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationSpec.
func (in *PostgresqlEngineConfigurationSpec) DeepCopy() *PostgresqlEngineConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlEngineConfigurationStatus) DeepCopyInto(out *PostgresqlEngineConfigurationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlEngineConfigurationStatus.
func (in *PostgresqlEngineConfigurationStatus) DeepCopy() *PostgresqlEngineConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlEngineConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUser) DeepCopyInto(out *PostgresqlUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUser.
func (in *PostgresqlUser) DeepCopy() *PostgresqlUser {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserList) DeepCopyInto(out *PostgresqlUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserList.
func (in *PostgresqlUserList) DeepCopy() *PostgresqlUserList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserSpec) DeepCopyInto(out *PostgresqlUserSpec) {
	*out = *in
	out.Database = in.Database
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserSpec.
func (in *PostgresqlUserSpec) DeepCopy() *PostgresqlUserSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserStatus) DeepCopyInto(out *PostgresqlUserStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserStatus.
func (in *PostgresqlUserStatus) DeepCopy() *PostgresqlUserStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusPostgresRoles) DeepCopyInto(out *StatusPostgresRoles) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusPostgresRoles.
func (in *StatusPostgresRoles) DeepCopy() *StatusPostgresRoles {
	if in == nil {
		return nil
	}
	out := new(StatusPostgresRoles)
	in.DeepCopyInto(out)
	return out
}
