// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Apple Identity Provider Resource
//
// The Apple identity provider type will use the Sign in with Apple APIs and will provide a Sign with Apple button on FusionAuth’s login page that will either redirect to an Apple sign in page or leverage native controls when using Safari on macOS or iOS. Additionally, this identity provider will call Apples’s /auth/token API to load additional details about the user and store them in FusionAuth.
//
// FusionAuth will also store the Apple refreshToken that is returned from the /auth/token endpoint in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
//
// [Apple Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/apple/#create-the-apple-identity-provider)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fusionauth/sdk/go/fusionauth"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fusionauth.NewFusionAuthIdpApple(ctx, "apple", &fusionauth.FusionAuthIdpAppleArgs{
// 			ApplicationConfigurations: FusionAuthIdpAppleApplicationConfigurationArray{
// 				&FusionAuthIdpAppleApplicationConfigurationArgs{
// 					ApplicationId:      pulumi.String("1c212e59-0d0e-6b1a-ad48-f4f92793be32"),
// 					CreateRegistration: pulumi.Bool(true),
// 					Enabled:            pulumi.Bool(true),
// 				},
// 			},
// 			ButtonText: pulumi.String("Sign in with Apple"),
// 			Debug:      pulumi.Bool(false),
// 			Enabled:    pulumi.Bool(true),
// 			KeyId:      pulumi.String("2f81529c-4d39-4ce2-982e-cf5fbb1325f6"),
// 			Scope:      pulumi.String("email name"),
// 			ServicesId: pulumi.String("com.piedpiper.webapp"),
// 			TeamId:     pulumi.String("R4NQ1P4UEB"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FusionAuthIdpApple struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpAppleApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The top-level space separated scope that you are requesting from Apple.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	ServicesId pulumi.StringOutput `pulumi:"servicesId"`
	// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpAppleTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
}

// NewFusionAuthIdpApple registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpApple(ctx *pulumi.Context,
	name string, args *FusionAuthIdpAppleArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpApple, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ButtonText == nil {
		return nil, errors.New("invalid value for required argument 'ButtonText'")
	}
	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.ServicesId == nil {
		return nil, errors.New("invalid value for required argument 'ServicesId'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpApple
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpApple:FusionAuthIdpApple", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpApple gets an existing FusionAuthIdpApple resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpApple(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpAppleState, opts ...pulumi.ResourceOption) (*FusionAuthIdpApple, error) {
	var resource FusionAuthIdpApple
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpApple:FusionAuthIdpApple", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpApple resources.
type fusionAuthIdpAppleState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpAppleApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	KeyId *string `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level space separated scope that you are requesting from Apple.
	Scope *string `pulumi:"scope"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	ServicesId *string `pulumi:"servicesId"`
	// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
	TeamId *string `pulumi:"teamId"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpAppleTenantConfiguration `pulumi:"tenantConfigurations"`
}

type FusionAuthIdpAppleState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpAppleApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	KeyId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level space separated scope that you are requesting from Apple.
	Scope pulumi.StringPtrInput
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	ServicesId pulumi.StringPtrInput
	// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
	TeamId pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpAppleTenantConfigurationArrayInput
}

func (FusionAuthIdpAppleState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpAppleState)(nil)).Elem()
}

type fusionAuthIdpAppleArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpAppleApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	KeyId string `pulumi:"keyId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level space separated scope that you are requesting from Apple.
	Scope *string `pulumi:"scope"`
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	ServicesId string `pulumi:"servicesId"`
	// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
	TeamId string `pulumi:"teamId"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpAppleTenantConfiguration `pulumi:"tenantConfigurations"`
}

// The set of arguments for constructing a FusionAuthIdpApple resource.
type FusionAuthIdpAppleArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpAppleApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	KeyId pulumi.StringInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level space separated scope that you are requesting from Apple.
	Scope pulumi.StringPtrInput
	// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
	ServicesId pulumi.StringInput
	// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
	TeamId pulumi.StringInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpAppleTenantConfigurationArrayInput
}

func (FusionAuthIdpAppleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpAppleArgs)(nil)).Elem()
}

type FusionAuthIdpAppleInput interface {
	pulumi.Input

	ToFusionAuthIdpAppleOutput() FusionAuthIdpAppleOutput
	ToFusionAuthIdpAppleOutputWithContext(ctx context.Context) FusionAuthIdpAppleOutput
}

func (*FusionAuthIdpApple) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpApple)(nil)).Elem()
}

func (i *FusionAuthIdpApple) ToFusionAuthIdpAppleOutput() FusionAuthIdpAppleOutput {
	return i.ToFusionAuthIdpAppleOutputWithContext(context.Background())
}

func (i *FusionAuthIdpApple) ToFusionAuthIdpAppleOutputWithContext(ctx context.Context) FusionAuthIdpAppleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpAppleOutput)
}

// FusionAuthIdpAppleArrayInput is an input type that accepts FusionAuthIdpAppleArray and FusionAuthIdpAppleArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpAppleArrayInput` via:
//
//          FusionAuthIdpAppleArray{ FusionAuthIdpAppleArgs{...} }
type FusionAuthIdpAppleArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpAppleArrayOutput() FusionAuthIdpAppleArrayOutput
	ToFusionAuthIdpAppleArrayOutputWithContext(context.Context) FusionAuthIdpAppleArrayOutput
}

type FusionAuthIdpAppleArray []FusionAuthIdpAppleInput

func (FusionAuthIdpAppleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpApple)(nil)).Elem()
}

func (i FusionAuthIdpAppleArray) ToFusionAuthIdpAppleArrayOutput() FusionAuthIdpAppleArrayOutput {
	return i.ToFusionAuthIdpAppleArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpAppleArray) ToFusionAuthIdpAppleArrayOutputWithContext(ctx context.Context) FusionAuthIdpAppleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpAppleArrayOutput)
}

// FusionAuthIdpAppleMapInput is an input type that accepts FusionAuthIdpAppleMap and FusionAuthIdpAppleMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpAppleMapInput` via:
//
//          FusionAuthIdpAppleMap{ "key": FusionAuthIdpAppleArgs{...} }
type FusionAuthIdpAppleMapInput interface {
	pulumi.Input

	ToFusionAuthIdpAppleMapOutput() FusionAuthIdpAppleMapOutput
	ToFusionAuthIdpAppleMapOutputWithContext(context.Context) FusionAuthIdpAppleMapOutput
}

type FusionAuthIdpAppleMap map[string]FusionAuthIdpAppleInput

func (FusionAuthIdpAppleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpApple)(nil)).Elem()
}

func (i FusionAuthIdpAppleMap) ToFusionAuthIdpAppleMapOutput() FusionAuthIdpAppleMapOutput {
	return i.ToFusionAuthIdpAppleMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpAppleMap) ToFusionAuthIdpAppleMapOutputWithContext(ctx context.Context) FusionAuthIdpAppleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpAppleMapOutput)
}

type FusionAuthIdpAppleOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpAppleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpApple)(nil)).Elem()
}

func (o FusionAuthIdpAppleOutput) ToFusionAuthIdpAppleOutput() FusionAuthIdpAppleOutput {
	return o
}

func (o FusionAuthIdpAppleOutput) ToFusionAuthIdpAppleOutputWithContext(ctx context.Context) FusionAuthIdpAppleOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpAppleOutput) ApplicationConfigurations() FusionAuthIdpAppleApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) FusionAuthIdpAppleApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpAppleApplicationConfigurationArrayOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpAppleOutput) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpAppleOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpAppleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
func (o FusionAuthIdpAppleOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpAppleOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpAppleOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The top-level space separated scope that you are requesting from Apple.
func (o FusionAuthIdpAppleOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
func (o FusionAuthIdpAppleOutput) ServicesId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringOutput { return v.ServicesId }).(pulumi.StringOutput)
}

// The Apple App ID Prefix, or Team ID found in your Apple Developer Account which has been configured for Sign in with Apple.
func (o FusionAuthIdpAppleOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpAppleOutput) TenantConfigurations() FusionAuthIdpAppleTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpApple) FusionAuthIdpAppleTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpAppleTenantConfigurationArrayOutput)
}

type FusionAuthIdpAppleArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpAppleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpApple)(nil)).Elem()
}

func (o FusionAuthIdpAppleArrayOutput) ToFusionAuthIdpAppleArrayOutput() FusionAuthIdpAppleArrayOutput {
	return o
}

func (o FusionAuthIdpAppleArrayOutput) ToFusionAuthIdpAppleArrayOutputWithContext(ctx context.Context) FusionAuthIdpAppleArrayOutput {
	return o
}

func (o FusionAuthIdpAppleArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpAppleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpApple {
		return vs[0].([]*FusionAuthIdpApple)[vs[1].(int)]
	}).(FusionAuthIdpAppleOutput)
}

type FusionAuthIdpAppleMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpAppleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpApple)(nil)).Elem()
}

func (o FusionAuthIdpAppleMapOutput) ToFusionAuthIdpAppleMapOutput() FusionAuthIdpAppleMapOutput {
	return o
}

func (o FusionAuthIdpAppleMapOutput) ToFusionAuthIdpAppleMapOutputWithContext(ctx context.Context) FusionAuthIdpAppleMapOutput {
	return o
}

func (o FusionAuthIdpAppleMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpAppleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpApple {
		return vs[0].(map[string]*FusionAuthIdpApple)[vs[1].(string)]
	}).(FusionAuthIdpAppleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpAppleInput)(nil)).Elem(), &FusionAuthIdpApple{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpAppleArrayInput)(nil)).Elem(), FusionAuthIdpAppleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpAppleMapInput)(nil)).Elem(), FusionAuthIdpAppleMap{})
	pulumi.RegisterOutputType(FusionAuthIdpAppleOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpAppleArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpAppleMapOutput{})
}