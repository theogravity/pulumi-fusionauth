// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Lambda Resource
//
// Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.
//
// [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)
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
//			_, err := fusionauth.NewFusionAuthLambda(ctx, "preferred Username", &fusionauth.FusionAuthLambdaArgs{
//				Body:    pulumi.String("// Using the user and registration parameters add additional values to the jwt object.\nfunction populate(jwt, user, registration) {\n  jwt.preferred_username = registration.username;\n}\n  \n"),
//				Enabled: pulumi.Bool(true),
//				Type:    pulumi.String("JWTPopulate"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthLambda struct {
	pulumi.CustomResourceState

	// The lambda function body, a JavaScript function.
	Body pulumi.StringOutput `pulumi:"body"`
	// Whether or not debug event logging is enabled for this Lambda.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Whether or not this Lambda is enabled.
	//
	// Deprecated: Not currently used and may be removed in a future version.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The JavaScript execution engine for the lambda.
	EngineType pulumi.StringPtrOutput `pulumi:"engineType"`
	// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
	LambdaId pulumi.StringOutput `pulumi:"lambdaId"`
	// The name of the lambda.
	Name pulumi.StringOutput `pulumi:"name"`
	// The lambda type. The possible values are:
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFusionAuthLambda registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthLambda(ctx *pulumi.Context,
	name string, args *FusionAuthLambdaArgs, opts ...pulumi.ResourceOption) (*FusionAuthLambda, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthLambda
	err := ctx.RegisterResource("fusionauth:index/fusionAuthLambda:FusionAuthLambda", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthLambda gets an existing FusionAuthLambda resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthLambda(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthLambdaState, opts ...pulumi.ResourceOption) (*FusionAuthLambda, error) {
	var resource FusionAuthLambda
	err := ctx.ReadResource("fusionauth:index/fusionAuthLambda:FusionAuthLambda", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthLambda resources.
type fusionAuthLambdaState struct {
	// The lambda function body, a JavaScript function.
	Body *string `pulumi:"body"`
	// Whether or not debug event logging is enabled for this Lambda.
	Debug *bool `pulumi:"debug"`
	// Whether or not this Lambda is enabled.
	//
	// Deprecated: Not currently used and may be removed in a future version.
	Enabled *bool `pulumi:"enabled"`
	// The JavaScript execution engine for the lambda.
	EngineType *string `pulumi:"engineType"`
	// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
	LambdaId *string `pulumi:"lambdaId"`
	// The name of the lambda.
	Name *string `pulumi:"name"`
	// The lambda type. The possible values are:
	Type *string `pulumi:"type"`
}

type FusionAuthLambdaState struct {
	// The lambda function body, a JavaScript function.
	Body pulumi.StringPtrInput
	// Whether or not debug event logging is enabled for this Lambda.
	Debug pulumi.BoolPtrInput
	// Whether or not this Lambda is enabled.
	//
	// Deprecated: Not currently used and may be removed in a future version.
	Enabled pulumi.BoolPtrInput
	// The JavaScript execution engine for the lambda.
	EngineType pulumi.StringPtrInput
	// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
	LambdaId pulumi.StringPtrInput
	// The name of the lambda.
	Name pulumi.StringPtrInput
	// The lambda type. The possible values are:
	Type pulumi.StringPtrInput
}

func (FusionAuthLambdaState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthLambdaState)(nil)).Elem()
}

type fusionAuthLambdaArgs struct {
	// The lambda function body, a JavaScript function.
	Body string `pulumi:"body"`
	// Whether or not debug event logging is enabled for this Lambda.
	Debug *bool `pulumi:"debug"`
	// Whether or not this Lambda is enabled.
	//
	// Deprecated: Not currently used and may be removed in a future version.
	Enabled *bool `pulumi:"enabled"`
	// The JavaScript execution engine for the lambda.
	EngineType *string `pulumi:"engineType"`
	// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
	LambdaId *string `pulumi:"lambdaId"`
	// The name of the lambda.
	Name *string `pulumi:"name"`
	// The lambda type. The possible values are:
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FusionAuthLambda resource.
type FusionAuthLambdaArgs struct {
	// The lambda function body, a JavaScript function.
	Body pulumi.StringInput
	// Whether or not debug event logging is enabled for this Lambda.
	Debug pulumi.BoolPtrInput
	// Whether or not this Lambda is enabled.
	//
	// Deprecated: Not currently used and may be removed in a future version.
	Enabled pulumi.BoolPtrInput
	// The JavaScript execution engine for the lambda.
	EngineType pulumi.StringPtrInput
	// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
	LambdaId pulumi.StringPtrInput
	// The name of the lambda.
	Name pulumi.StringPtrInput
	// The lambda type. The possible values are:
	Type pulumi.StringInput
}

func (FusionAuthLambdaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthLambdaArgs)(nil)).Elem()
}

type FusionAuthLambdaInput interface {
	pulumi.Input

	ToFusionAuthLambdaOutput() FusionAuthLambdaOutput
	ToFusionAuthLambdaOutputWithContext(ctx context.Context) FusionAuthLambdaOutput
}

func (*FusionAuthLambda) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthLambda)(nil)).Elem()
}

func (i *FusionAuthLambda) ToFusionAuthLambdaOutput() FusionAuthLambdaOutput {
	return i.ToFusionAuthLambdaOutputWithContext(context.Background())
}

func (i *FusionAuthLambda) ToFusionAuthLambdaOutputWithContext(ctx context.Context) FusionAuthLambdaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthLambdaOutput)
}

// FusionAuthLambdaArrayInput is an input type that accepts FusionAuthLambdaArray and FusionAuthLambdaArrayOutput values.
// You can construct a concrete instance of `FusionAuthLambdaArrayInput` via:
//
//	FusionAuthLambdaArray{ FusionAuthLambdaArgs{...} }
type FusionAuthLambdaArrayInput interface {
	pulumi.Input

	ToFusionAuthLambdaArrayOutput() FusionAuthLambdaArrayOutput
	ToFusionAuthLambdaArrayOutputWithContext(context.Context) FusionAuthLambdaArrayOutput
}

type FusionAuthLambdaArray []FusionAuthLambdaInput

func (FusionAuthLambdaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthLambda)(nil)).Elem()
}

func (i FusionAuthLambdaArray) ToFusionAuthLambdaArrayOutput() FusionAuthLambdaArrayOutput {
	return i.ToFusionAuthLambdaArrayOutputWithContext(context.Background())
}

func (i FusionAuthLambdaArray) ToFusionAuthLambdaArrayOutputWithContext(ctx context.Context) FusionAuthLambdaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthLambdaArrayOutput)
}

// FusionAuthLambdaMapInput is an input type that accepts FusionAuthLambdaMap and FusionAuthLambdaMapOutput values.
// You can construct a concrete instance of `FusionAuthLambdaMapInput` via:
//
//	FusionAuthLambdaMap{ "key": FusionAuthLambdaArgs{...} }
type FusionAuthLambdaMapInput interface {
	pulumi.Input

	ToFusionAuthLambdaMapOutput() FusionAuthLambdaMapOutput
	ToFusionAuthLambdaMapOutputWithContext(context.Context) FusionAuthLambdaMapOutput
}

type FusionAuthLambdaMap map[string]FusionAuthLambdaInput

func (FusionAuthLambdaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthLambda)(nil)).Elem()
}

func (i FusionAuthLambdaMap) ToFusionAuthLambdaMapOutput() FusionAuthLambdaMapOutput {
	return i.ToFusionAuthLambdaMapOutputWithContext(context.Background())
}

func (i FusionAuthLambdaMap) ToFusionAuthLambdaMapOutputWithContext(ctx context.Context) FusionAuthLambdaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthLambdaMapOutput)
}

type FusionAuthLambdaOutput struct{ *pulumi.OutputState }

func (FusionAuthLambdaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthLambda)(nil)).Elem()
}

func (o FusionAuthLambdaOutput) ToFusionAuthLambdaOutput() FusionAuthLambdaOutput {
	return o
}

func (o FusionAuthLambdaOutput) ToFusionAuthLambdaOutputWithContext(ctx context.Context) FusionAuthLambdaOutput {
	return o
}

// The lambda function body, a JavaScript function.
func (o FusionAuthLambdaOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

// Whether or not debug event logging is enabled for this Lambda.
func (o FusionAuthLambdaOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Whether or not this Lambda is enabled.
//
// Deprecated: Not currently used and may be removed in a future version.
func (o FusionAuthLambdaOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The JavaScript execution engine for the lambda.
func (o FusionAuthLambdaOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

// The Id to use for the new lambda. If not specified a secure random UUID will be generated.
func (o FusionAuthLambdaOutput) LambdaId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.StringOutput { return v.LambdaId }).(pulumi.StringOutput)
}

// The name of the lambda.
func (o FusionAuthLambdaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The lambda type. The possible values are:
func (o FusionAuthLambdaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthLambda) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FusionAuthLambdaArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthLambdaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthLambda)(nil)).Elem()
}

func (o FusionAuthLambdaArrayOutput) ToFusionAuthLambdaArrayOutput() FusionAuthLambdaArrayOutput {
	return o
}

func (o FusionAuthLambdaArrayOutput) ToFusionAuthLambdaArrayOutputWithContext(ctx context.Context) FusionAuthLambdaArrayOutput {
	return o
}

func (o FusionAuthLambdaArrayOutput) Index(i pulumi.IntInput) FusionAuthLambdaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthLambda {
		return vs[0].([]*FusionAuthLambda)[vs[1].(int)]
	}).(FusionAuthLambdaOutput)
}

type FusionAuthLambdaMapOutput struct{ *pulumi.OutputState }

func (FusionAuthLambdaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthLambda)(nil)).Elem()
}

func (o FusionAuthLambdaMapOutput) ToFusionAuthLambdaMapOutput() FusionAuthLambdaMapOutput {
	return o
}

func (o FusionAuthLambdaMapOutput) ToFusionAuthLambdaMapOutputWithContext(ctx context.Context) FusionAuthLambdaMapOutput {
	return o
}

func (o FusionAuthLambdaMapOutput) MapIndex(k pulumi.StringInput) FusionAuthLambdaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthLambda {
		return vs[0].(map[string]*FusionAuthLambda)[vs[1].(string)]
	}).(FusionAuthLambdaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthLambdaInput)(nil)).Elem(), &FusionAuthLambda{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthLambdaArrayInput)(nil)).Elem(), FusionAuthLambdaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthLambdaMapInput)(nil)).Elem(), FusionAuthLambdaMap{})
	pulumi.RegisterOutputType(FusionAuthLambdaOutput{})
	pulumi.RegisterOutputType(FusionAuthLambdaArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthLambdaMapOutput{})
}
