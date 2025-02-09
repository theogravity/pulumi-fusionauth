// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # Entity Type Resource
//
// Entity Types categorize Entities. For example, an Entity Type could be `Device`, `API` or `Company`.
//
// [Entity Type API](https://fusionauth.io/docs/v1/tech/apis/entity-management/entity-types/#create-an-entity-type)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"createdBy": "jared@fusionauth.io",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = fusionauth.NewFusionAuthEntityType(ctx, "company", &fusionauth.FusionAuthEntityTypeArgs{
//				Data: pulumi.String(json0),
//				JwtConfiguration: &fusionauth.FusionAuthEntityTypeJwtConfigurationArgs{
//					AccessTokenKeyId:    pulumi.String("a7516c7c-6234-4021-b0b4-8870c807aeb2"),
//					Enabled:             pulumi.Bool(true),
//					TimeToLiveInSeconds: pulumi.Int(3600),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthEntityType struct {
	pulumi.CustomResourceState

	// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
	EntityTypeId pulumi.StringOutput `pulumi:"entityTypeId"`
	// A block to configure JSON Web Token (JWT) options.
	JwtConfiguration FusionAuthEntityTypeJwtConfigurationOutput `pulumi:"jwtConfiguration"`
	// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFusionAuthEntityType registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthEntityType(ctx *pulumi.Context,
	name string, args *FusionAuthEntityTypeArgs, opts ...pulumi.ResourceOption) (*FusionAuthEntityType, error) {
	if args == nil {
		args = &FusionAuthEntityTypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthEntityType
	err := ctx.RegisterResource("fusionauth:index/fusionAuthEntityType:FusionAuthEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthEntityType gets an existing FusionAuthEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthEntityTypeState, opts ...pulumi.ResourceOption) (*FusionAuthEntityType, error) {
	var resource FusionAuthEntityType
	err := ctx.ReadResource("fusionauth:index/fusionAuthEntityType:FusionAuthEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthEntityType resources.
type fusionAuthEntityTypeState struct {
	// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
	Data *string `pulumi:"data"`
	// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
	EntityTypeId *string `pulumi:"entityTypeId"`
	// A block to configure JSON Web Token (JWT) options.
	JwtConfiguration *FusionAuthEntityTypeJwtConfiguration `pulumi:"jwtConfiguration"`
	// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
	Name *string `pulumi:"name"`
}

type FusionAuthEntityTypeState struct {
	// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
	Data pulumi.StringPtrInput
	// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
	EntityTypeId pulumi.StringPtrInput
	// A block to configure JSON Web Token (JWT) options.
	JwtConfiguration FusionAuthEntityTypeJwtConfigurationPtrInput
	// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
	Name pulumi.StringPtrInput
}

func (FusionAuthEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthEntityTypeState)(nil)).Elem()
}

type fusionAuthEntityTypeArgs struct {
	// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
	Data *string `pulumi:"data"`
	// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
	EntityTypeId *string `pulumi:"entityTypeId"`
	// A block to configure JSON Web Token (JWT) options.
	JwtConfiguration *FusionAuthEntityTypeJwtConfiguration `pulumi:"jwtConfiguration"`
	// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FusionAuthEntityType resource.
type FusionAuthEntityTypeArgs struct {
	// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
	Data pulumi.StringPtrInput
	// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
	EntityTypeId pulumi.StringPtrInput
	// A block to configure JSON Web Token (JWT) options.
	JwtConfiguration FusionAuthEntityTypeJwtConfigurationPtrInput
	// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
	Name pulumi.StringPtrInput
}

func (FusionAuthEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthEntityTypeArgs)(nil)).Elem()
}

type FusionAuthEntityTypeInput interface {
	pulumi.Input

	ToFusionAuthEntityTypeOutput() FusionAuthEntityTypeOutput
	ToFusionAuthEntityTypeOutputWithContext(ctx context.Context) FusionAuthEntityTypeOutput
}

func (*FusionAuthEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthEntityType)(nil)).Elem()
}

func (i *FusionAuthEntityType) ToFusionAuthEntityTypeOutput() FusionAuthEntityTypeOutput {
	return i.ToFusionAuthEntityTypeOutputWithContext(context.Background())
}

func (i *FusionAuthEntityType) ToFusionAuthEntityTypeOutputWithContext(ctx context.Context) FusionAuthEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypeOutput)
}

// FusionAuthEntityTypeArrayInput is an input type that accepts FusionAuthEntityTypeArray and FusionAuthEntityTypeArrayOutput values.
// You can construct a concrete instance of `FusionAuthEntityTypeArrayInput` via:
//
//	FusionAuthEntityTypeArray{ FusionAuthEntityTypeArgs{...} }
type FusionAuthEntityTypeArrayInput interface {
	pulumi.Input

	ToFusionAuthEntityTypeArrayOutput() FusionAuthEntityTypeArrayOutput
	ToFusionAuthEntityTypeArrayOutputWithContext(context.Context) FusionAuthEntityTypeArrayOutput
}

type FusionAuthEntityTypeArray []FusionAuthEntityTypeInput

func (FusionAuthEntityTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthEntityType)(nil)).Elem()
}

func (i FusionAuthEntityTypeArray) ToFusionAuthEntityTypeArrayOutput() FusionAuthEntityTypeArrayOutput {
	return i.ToFusionAuthEntityTypeArrayOutputWithContext(context.Background())
}

func (i FusionAuthEntityTypeArray) ToFusionAuthEntityTypeArrayOutputWithContext(ctx context.Context) FusionAuthEntityTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypeArrayOutput)
}

// FusionAuthEntityTypeMapInput is an input type that accepts FusionAuthEntityTypeMap and FusionAuthEntityTypeMapOutput values.
// You can construct a concrete instance of `FusionAuthEntityTypeMapInput` via:
//
//	FusionAuthEntityTypeMap{ "key": FusionAuthEntityTypeArgs{...} }
type FusionAuthEntityTypeMapInput interface {
	pulumi.Input

	ToFusionAuthEntityTypeMapOutput() FusionAuthEntityTypeMapOutput
	ToFusionAuthEntityTypeMapOutputWithContext(context.Context) FusionAuthEntityTypeMapOutput
}

type FusionAuthEntityTypeMap map[string]FusionAuthEntityTypeInput

func (FusionAuthEntityTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthEntityType)(nil)).Elem()
}

func (i FusionAuthEntityTypeMap) ToFusionAuthEntityTypeMapOutput() FusionAuthEntityTypeMapOutput {
	return i.ToFusionAuthEntityTypeMapOutputWithContext(context.Background())
}

func (i FusionAuthEntityTypeMap) ToFusionAuthEntityTypeMapOutputWithContext(ctx context.Context) FusionAuthEntityTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthEntityTypeMapOutput)
}

type FusionAuthEntityTypeOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthEntityType)(nil)).Elem()
}

func (o FusionAuthEntityTypeOutput) ToFusionAuthEntityTypeOutput() FusionAuthEntityTypeOutput {
	return o
}

func (o FusionAuthEntityTypeOutput) ToFusionAuthEntityTypeOutputWithContext(ctx context.Context) FusionAuthEntityTypeOutput {
	return o
}

// An object that can hold any information about the Entity Type that should be persisted. Must be aJSON string.
func (o FusionAuthEntityTypeOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthEntityType) pulumi.StringPtrOutput { return v.Data }).(pulumi.StringPtrOutput)
}

// The ID to use for the new Entity Type. If not specified a secure random UUID will be generated.
func (o FusionAuthEntityTypeOutput) EntityTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthEntityType) pulumi.StringOutput { return v.EntityTypeId }).(pulumi.StringOutput)
}

// A block to configure JSON Web Token (JWT) options.
func (o FusionAuthEntityTypeOutput) JwtConfiguration() FusionAuthEntityTypeJwtConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthEntityType) FusionAuthEntityTypeJwtConfigurationOutput { return v.JwtConfiguration }).(FusionAuthEntityTypeJwtConfigurationOutput)
}

// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
func (o FusionAuthEntityTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthEntityType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FusionAuthEntityTypeArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthEntityType)(nil)).Elem()
}

func (o FusionAuthEntityTypeArrayOutput) ToFusionAuthEntityTypeArrayOutput() FusionAuthEntityTypeArrayOutput {
	return o
}

func (o FusionAuthEntityTypeArrayOutput) ToFusionAuthEntityTypeArrayOutputWithContext(ctx context.Context) FusionAuthEntityTypeArrayOutput {
	return o
}

func (o FusionAuthEntityTypeArrayOutput) Index(i pulumi.IntInput) FusionAuthEntityTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthEntityType {
		return vs[0].([]*FusionAuthEntityType)[vs[1].(int)]
	}).(FusionAuthEntityTypeOutput)
}

type FusionAuthEntityTypeMapOutput struct{ *pulumi.OutputState }

func (FusionAuthEntityTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthEntityType)(nil)).Elem()
}

func (o FusionAuthEntityTypeMapOutput) ToFusionAuthEntityTypeMapOutput() FusionAuthEntityTypeMapOutput {
	return o
}

func (o FusionAuthEntityTypeMapOutput) ToFusionAuthEntityTypeMapOutputWithContext(ctx context.Context) FusionAuthEntityTypeMapOutput {
	return o
}

func (o FusionAuthEntityTypeMapOutput) MapIndex(k pulumi.StringInput) FusionAuthEntityTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthEntityType {
		return vs[0].(map[string]*FusionAuthEntityType)[vs[1].(string)]
	}).(FusionAuthEntityTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypeInput)(nil)).Elem(), &FusionAuthEntityType{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypeArrayInput)(nil)).Elem(), FusionAuthEntityTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthEntityTypeMapInput)(nil)).Elem(), FusionAuthEntityTypeMap{})
	pulumi.RegisterOutputType(FusionAuthEntityTypeOutput{})
	pulumi.RegisterOutputType(FusionAuthEntityTypeArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthEntityTypeMapOutput{})
}
