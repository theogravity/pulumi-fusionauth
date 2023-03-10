// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FusionAuthEntityTypePermission struct {
	pulumi.CustomResourceState

	// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// The description of the Permission.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Id of the Entity Type.
	EntityTypeId pulumi.StringOutput `pulumi:"entityTypeId"`
	// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
	// this type if no permissions are provided in a grant request.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// The name of the Permission.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id to use for the new permission. If not specified a secure random UUID will be generated.
	PermissionId pulumi.StringOutput `pulumi:"permissionId"`
}

// NewFusionAuthEntityTypePermission registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthEntityTypePermission(ctx *pulumi.Context,
	name string, args *FusionAuthEntityTypePermissionArgs, opts ...pulumi.ResourceOption) (*FusionAuthEntityTypePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityTypeId == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthEntityTypePermission
	err := ctx.RegisterResource("fusionauth:index/fusionAuthEntityTypePermission:FusionAuthEntityTypePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthEntityTypePermission gets an existing FusionAuthEntityTypePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthEntityTypePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthEntityTypePermissionState, opts ...pulumi.ResourceOption) (*FusionAuthEntityTypePermission, error) {
	var resource FusionAuthEntityTypePermission
	err := ctx.ReadResource("fusionauth:index/fusionAuthEntityTypePermission:FusionAuthEntityTypePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthEntityTypePermission resources.
type fusionAuthEntityTypePermissionState struct {
	// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
	Data *string `pulumi:"data"`
	// The description of the Permission.
	Description *string `pulumi:"description"`
	// The Id of the Entity Type.
	EntityTypeId *string `pulumi:"entityTypeId"`
	// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
	// this type if no permissions are provided in a grant request.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the Permission.
	Name *string `pulumi:"name"`
	// The Id to use for the new permission. If not specified a secure random UUID will be generated.
	PermissionId *string `pulumi:"permissionId"`
}

type FusionAuthEntityTypePermissionState struct {
	// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
	Data pulumi.StringPtrInput
	// The description of the Permission.
	Description pulumi.StringPtrInput
	// The Id of the Entity Type.
	EntityTypeId pulumi.StringPtrInput
	// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
	// this type if no permissions are provided in a grant request.
	IsDefault pulumi.BoolPtrInput
	// The name of the Permission.
	Name pulumi.StringPtrInput
	// The Id to use for the new permission. If not specified a secure random UUID will be generated.
	PermissionId pulumi.StringPtrInput
}

func (FusionAuthEntityTypePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthEntityTypePermissionState)(nil)).Elem()
}

type fusionAuthEntityTypePermissionArgs struct {
	// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
	Data *string `pulumi:"data"`
	// The description of the Permission.
	Description *string `pulumi:"description"`
	// The Id of the Entity Type.
	EntityTypeId string `pulumi:"entityTypeId"`
	// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
	// this type if no permissions are provided in a grant request.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the Permission.
	Name *string `pulumi:"name"`
	// The Id to use for the new permission. If not specified a secure random UUID will be generated.
	PermissionId *string `pulumi:"permissionId"`
}

// The set of arguments for constructing a FusionAuthEntityTypePermission resource.
type FusionAuthEntityTypePermissionArgs struct {
	// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
	Data pulumi.StringPtrInput
	// The description of the Permission.
	Description pulumi.StringPtrInput
	// The Id of the Entity Type.
	EntityTypeId pulumi.StringInput
	// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
	// this type if no permissions are provided in a grant request.
	IsDefault pulumi.BoolPtrInput
	// The name of the Permission.
	Name pulumi.StringPtrInput
	// The Id to use for the new permission. If not specified a secure random UUID will be generated.
	PermissionId pulumi.StringPtrInput
}

func (FusionAuthEntityTypePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthEntityTypePermissionArgs)(nil)).Elem()
}

type FusionAuthEntityTypePermissionInput interface {
	pulumi.Input

	ToFusionAuthEntityTypePermissionOutput() FusionAuthEntityTypePermissionOutput
	ToFusionAuthEntityTypePermissionOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionOutput
}

func (*FusionAuthEntityTypePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthEntityTypePermission)(nil)).Elem()
}

func (i *FusionAuthEntityTypePermission) ToFusionAuthEntityTypePermissionOutput() FusionAuthEntityTypePermissionOutput {
	return i.ToFusionAuthEntityTypePermissionOutputWithContext(context.Background())
}

func (i *FusionAuthEntityTypePermission) ToFusionAuthEntityTypePermissionOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypePermissionOutput)
}

// FusionAuthEntityTypePermissionArrayInput is an input type that accepts FusionAuthEntityTypePermissionArray and FusionAuthEntityTypePermissionArrayOutput values.
// You can construct a concrete instance of `FusionAuthEntityTypePermissionArrayInput` via:
//
//	FusionAuthEntityTypePermissionArray{ FusionAuthEntityTypePermissionArgs{...} }
type FusionAuthEntityTypePermissionArrayInput interface {
	pulumi.Input

	ToFusionAuthEntityTypePermissionArrayOutput() FusionAuthEntityTypePermissionArrayOutput
	ToFusionAuthEntityTypePermissionArrayOutputWithContext(context.Context) FusionAuthEntityTypePermissionArrayOutput
}

type FusionAuthEntityTypePermissionArray []FusionAuthEntityTypePermissionInput

func (FusionAuthEntityTypePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthEntityTypePermission)(nil)).Elem()
}

func (i FusionAuthEntityTypePermissionArray) ToFusionAuthEntityTypePermissionArrayOutput() FusionAuthEntityTypePermissionArrayOutput {
	return i.ToFusionAuthEntityTypePermissionArrayOutputWithContext(context.Background())
}

func (i FusionAuthEntityTypePermissionArray) ToFusionAuthEntityTypePermissionArrayOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypePermissionArrayOutput)
}

// FusionAuthEntityTypePermissionMapInput is an input type that accepts FusionAuthEntityTypePermissionMap and FusionAuthEntityTypePermissionMapOutput values.
// You can construct a concrete instance of `FusionAuthEntityTypePermissionMapInput` via:
//
//	FusionAuthEntityTypePermissionMap{ "key": FusionAuthEntityTypePermissionArgs{...} }
type FusionAuthEntityTypePermissionMapInput interface {
	pulumi.Input

	ToFusionAuthEntityTypePermissionMapOutput() FusionAuthEntityTypePermissionMapOutput
	ToFusionAuthEntityTypePermissionMapOutputWithContext(context.Context) FusionAuthEntityTypePermissionMapOutput
}

type FusionAuthEntityTypePermissionMap map[string]FusionAuthEntityTypePermissionInput

func (FusionAuthEntityTypePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthEntityTypePermission)(nil)).Elem()
}

func (i FusionAuthEntityTypePermissionMap) ToFusionAuthEntityTypePermissionMapOutput() FusionAuthEntityTypePermissionMapOutput {
	return i.ToFusionAuthEntityTypePermissionMapOutputWithContext(context.Background())
}

func (i FusionAuthEntityTypePermissionMap) ToFusionAuthEntityTypePermissionMapOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypePermissionMapOutput)
}

type FusionAuthEntityTypePermissionOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthEntityTypePermission)(nil)).Elem()
}

func (o FusionAuthEntityTypePermissionOutput) ToFusionAuthEntityTypePermissionOutput() FusionAuthEntityTypePermissionOutput {
	return o
}

func (o FusionAuthEntityTypePermissionOutput) ToFusionAuthEntityTypePermissionOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionOutput {
	return o
}

// An object that can hold any information about the Permission that should be persisted. Must be a JSON string.
func (o FusionAuthEntityTypePermissionOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.StringPtrOutput { return v.Data }).(pulumi.StringPtrOutput)
}

// The description of the Permission.
func (o FusionAuthEntityTypePermissionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Id of the Entity Type.
func (o FusionAuthEntityTypePermissionOutput) EntityTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.StringOutput { return v.EntityTypeId }).(pulumi.StringOutput)
}

// Whether or not the Permission is a default permission. A default permission is automatically granted to an entity of
// this type if no permissions are provided in a grant request.
func (o FusionAuthEntityTypePermissionOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The name of the Permission.
func (o FusionAuthEntityTypePermissionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Id to use for the new permission. If not specified a secure random UUID will be generated.
func (o FusionAuthEntityTypePermissionOutput) PermissionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthEntityTypePermission) pulumi.StringOutput { return v.PermissionId }).(pulumi.StringOutput)
}

type FusionAuthEntityTypePermissionArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthEntityTypePermission)(nil)).Elem()
}

func (o FusionAuthEntityTypePermissionArrayOutput) ToFusionAuthEntityTypePermissionArrayOutput() FusionAuthEntityTypePermissionArrayOutput {
	return o
}

func (o FusionAuthEntityTypePermissionArrayOutput) ToFusionAuthEntityTypePermissionArrayOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionArrayOutput {
	return o
}

func (o FusionAuthEntityTypePermissionArrayOutput) Index(i pulumi.IntInput) FusionAuthEntityTypePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthEntityTypePermission {
		return vs[0].([]*FusionAuthEntityTypePermission)[vs[1].(int)]
	}).(FusionAuthEntityTypePermissionOutput)
}

type FusionAuthEntityTypePermissionMapOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthEntityTypePermission)(nil)).Elem()
}

func (o FusionAuthEntityTypePermissionMapOutput) ToFusionAuthEntityTypePermissionMapOutput() FusionAuthEntityTypePermissionMapOutput {
	return o
}

func (o FusionAuthEntityTypePermissionMapOutput) ToFusionAuthEntityTypePermissionMapOutputWithContext(ctx context.Context) FusionAuthEntityTypePermissionMapOutput {
	return o
}

func (o FusionAuthEntityTypePermissionMapOutput) MapIndex(k pulumi.StringInput) FusionAuthEntityTypePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthEntityTypePermission {
		return vs[0].(map[string]*FusionAuthEntityTypePermission)[vs[1].(string)]
	}).(FusionAuthEntityTypePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypePermissionInput)(nil)).Elem(), &FusionAuthEntityTypePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypePermissionArrayInput)(nil)).Elem(), FusionAuthEntityTypePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypePermissionMapInput)(nil)).Elem(), FusionAuthEntityTypePermissionMap{})
	pulumi.RegisterOutputType(FusionAuthEntityTypePermissionOutput{})
	pulumi.RegisterOutputType(FusionAuthEntityTypePermissionArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthEntityTypePermissionMapOutput{})
}
