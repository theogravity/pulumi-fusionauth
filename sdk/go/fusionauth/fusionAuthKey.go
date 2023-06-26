// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Key Resource
//
// Cryptographic keys are used in signing and verifying JWTs and verifying responses for third party identity providers. It is more likely you will interact with keys using the FusionAuth UI in the Key Master menu.
//
// [Keys API](https://fusionauth.io/docs/v1/tech/apis/keys)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/v3/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthKey(ctx, "adminId", &fusionauth.FusionAuthKeyArgs{
//				Algorithm: pulumi.String("RS256"),
//				Length:    pulumi.Int(2048),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthKey struct {
	pulumi.CustomResourceState

	// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// The Id to use for the new key. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The id used in the JWT header to identify the key used to generate the signature
	Kid pulumi.StringOutput `pulumi:"kid"`
	// The length of the RSA or EC certificate. This field is required when generating RSA key types.
	Length pulumi.IntPtrOutput `pulumi:"length"`
	// The name of the Key.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFusionAuthKey registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthKey(ctx *pulumi.Context,
	name string, args *FusionAuthKeyArgs, opts ...pulumi.ResourceOption) (*FusionAuthKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthKey
	err := ctx.RegisterResource("fusionauth:index/fusionAuthKey:FusionAuthKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthKey gets an existing FusionAuthKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthKeyState, opts ...pulumi.ResourceOption) (*FusionAuthKey, error) {
	var resource FusionAuthKey
	err := ctx.ReadResource("fusionauth:index/fusionAuthKey:FusionAuthKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthKey resources.
type fusionAuthKeyState struct {
	// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
	Algorithm *string `pulumi:"algorithm"`
	// The Id to use for the new key. If not specified a secure random UUID will be generated.
	KeyId *string `pulumi:"keyId"`
	// The id used in the JWT header to identify the key used to generate the signature
	Kid *string `pulumi:"kid"`
	// The length of the RSA or EC certificate. This field is required when generating RSA key types.
	Length *int `pulumi:"length"`
	// The name of the Key.
	Name *string `pulumi:"name"`
}

type FusionAuthKeyState struct {
	// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
	Algorithm pulumi.StringPtrInput
	// The Id to use for the new key. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringPtrInput
	// The id used in the JWT header to identify the key used to generate the signature
	Kid pulumi.StringPtrInput
	// The length of the RSA or EC certificate. This field is required when generating RSA key types.
	Length pulumi.IntPtrInput
	// The name of the Key.
	Name pulumi.StringPtrInput
}

func (FusionAuthKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthKeyState)(nil)).Elem()
}

type fusionAuthKeyArgs struct {
	// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
	Algorithm string `pulumi:"algorithm"`
	// The Id to use for the new key. If not specified a secure random UUID will be generated.
	KeyId *string `pulumi:"keyId"`
	// The length of the RSA or EC certificate. This field is required when generating RSA key types.
	Length *int `pulumi:"length"`
	// The name of the Key.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FusionAuthKey resource.
type FusionAuthKeyArgs struct {
	// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
	Algorithm pulumi.StringInput
	// The Id to use for the new key. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringPtrInput
	// The length of the RSA or EC certificate. This field is required when generating RSA key types.
	Length pulumi.IntPtrInput
	// The name of the Key.
	Name pulumi.StringPtrInput
}

func (FusionAuthKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthKeyArgs)(nil)).Elem()
}

type FusionAuthKeyInput interface {
	pulumi.Input

	ToFusionAuthKeyOutput() FusionAuthKeyOutput
	ToFusionAuthKeyOutputWithContext(ctx context.Context) FusionAuthKeyOutput
}

func (*FusionAuthKey) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthKey)(nil)).Elem()
}

func (i *FusionAuthKey) ToFusionAuthKeyOutput() FusionAuthKeyOutput {
	return i.ToFusionAuthKeyOutputWithContext(context.Background())
}

func (i *FusionAuthKey) ToFusionAuthKeyOutputWithContext(ctx context.Context) FusionAuthKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthKeyOutput)
}

// FusionAuthKeyArrayInput is an input type that accepts FusionAuthKeyArray and FusionAuthKeyArrayOutput values.
// You can construct a concrete instance of `FusionAuthKeyArrayInput` via:
//
//	FusionAuthKeyArray{ FusionAuthKeyArgs{...} }
type FusionAuthKeyArrayInput interface {
	pulumi.Input

	ToFusionAuthKeyArrayOutput() FusionAuthKeyArrayOutput
	ToFusionAuthKeyArrayOutputWithContext(context.Context) FusionAuthKeyArrayOutput
}

type FusionAuthKeyArray []FusionAuthKeyInput

func (FusionAuthKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthKey)(nil)).Elem()
}

func (i FusionAuthKeyArray) ToFusionAuthKeyArrayOutput() FusionAuthKeyArrayOutput {
	return i.ToFusionAuthKeyArrayOutputWithContext(context.Background())
}

func (i FusionAuthKeyArray) ToFusionAuthKeyArrayOutputWithContext(ctx context.Context) FusionAuthKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthKeyArrayOutput)
}

// FusionAuthKeyMapInput is an input type that accepts FusionAuthKeyMap and FusionAuthKeyMapOutput values.
// You can construct a concrete instance of `FusionAuthKeyMapInput` via:
//
//	FusionAuthKeyMap{ "key": FusionAuthKeyArgs{...} }
type FusionAuthKeyMapInput interface {
	pulumi.Input

	ToFusionAuthKeyMapOutput() FusionAuthKeyMapOutput
	ToFusionAuthKeyMapOutputWithContext(context.Context) FusionAuthKeyMapOutput
}

type FusionAuthKeyMap map[string]FusionAuthKeyInput

func (FusionAuthKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthKey)(nil)).Elem()
}

func (i FusionAuthKeyMap) ToFusionAuthKeyMapOutput() FusionAuthKeyMapOutput {
	return i.ToFusionAuthKeyMapOutputWithContext(context.Background())
}

func (i FusionAuthKeyMap) ToFusionAuthKeyMapOutputWithContext(ctx context.Context) FusionAuthKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthKeyMapOutput)
}

type FusionAuthKeyOutput struct{ *pulumi.OutputState }

func (FusionAuthKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthKey)(nil)).Elem()
}

func (o FusionAuthKeyOutput) ToFusionAuthKeyOutput() FusionAuthKeyOutput {
	return o
}

func (o FusionAuthKeyOutput) ToFusionAuthKeyOutputWithContext(ctx context.Context) FusionAuthKeyOutput {
	return o
}

// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
func (o FusionAuthKeyOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthKey) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// The Id to use for the new key. If not specified a secure random UUID will be generated.
func (o FusionAuthKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The id used in the JWT header to identify the key used to generate the signature
func (o FusionAuthKeyOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthKey) pulumi.StringOutput { return v.Kid }).(pulumi.StringOutput)
}

// The length of the RSA or EC certificate. This field is required when generating RSA key types.
func (o FusionAuthKeyOutput) Length() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FusionAuthKey) pulumi.IntPtrOutput { return v.Length }).(pulumi.IntPtrOutput)
}

// The name of the Key.
func (o FusionAuthKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FusionAuthKeyArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthKey)(nil)).Elem()
}

func (o FusionAuthKeyArrayOutput) ToFusionAuthKeyArrayOutput() FusionAuthKeyArrayOutput {
	return o
}

func (o FusionAuthKeyArrayOutput) ToFusionAuthKeyArrayOutputWithContext(ctx context.Context) FusionAuthKeyArrayOutput {
	return o
}

func (o FusionAuthKeyArrayOutput) Index(i pulumi.IntInput) FusionAuthKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthKey {
		return vs[0].([]*FusionAuthKey)[vs[1].(int)]
	}).(FusionAuthKeyOutput)
}

type FusionAuthKeyMapOutput struct{ *pulumi.OutputState }

func (FusionAuthKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthKey)(nil)).Elem()
}

func (o FusionAuthKeyMapOutput) ToFusionAuthKeyMapOutput() FusionAuthKeyMapOutput {
	return o
}

func (o FusionAuthKeyMapOutput) ToFusionAuthKeyMapOutputWithContext(ctx context.Context) FusionAuthKeyMapOutput {
	return o
}

func (o FusionAuthKeyMapOutput) MapIndex(k pulumi.StringInput) FusionAuthKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthKey {
		return vs[0].(map[string]*FusionAuthKey)[vs[1].(string)]
	}).(FusionAuthKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthKeyInput)(nil)).Elem(), &FusionAuthKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthKeyArrayInput)(nil)).Elem(), FusionAuthKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthKeyMapInput)(nil)).Elem(), FusionAuthKeyMap{})
	pulumi.RegisterOutputType(FusionAuthKeyOutput{})
	pulumi.RegisterOutputType(FusionAuthKeyArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthKeyMapOutput{})
}
