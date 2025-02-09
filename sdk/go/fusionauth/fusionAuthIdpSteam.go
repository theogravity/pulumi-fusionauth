// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # Steam Identity Provider Resource
//
// The Steam identity provider type will use the Steam OAuth login API. It will also provide a Login with Steam button on FusionAuth’s login page that will direct a user to the Steam login page. The Steam login uses the implicit OAuth grant and will return to the callback URL with token and state in the URL Fragment. This is handled by the FusionAuth /oauth2/implicit javascript function to pass those values to the /oauth2/callback endpoint.
//
// This identity provider will call Steam’s API to load the Steam user’s personaname and use that as username to lookup or create a user in FusionAuth depending on the linking strategy configured for this identity provider. However, Steam does not allow access to user emails, so neither email linking strategy can be used and user’s will not be able to login or be created.
//
// FusionAuth will also store the Steam token that is returned from the Steam login in the link between the user and the identity provider. This token can be used by an application to make further requests to Steam APIs on behalf of the user.
//
// [Steam Identity Provider APIs](https://fusionauth.io/docs/v1/tech/apis/identity-providers/steam/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthIdpSteam(ctx, "steam", &fusionauth.FusionAuthIdpSteamArgs{
//				ApplicationConfigurations: fusionauth.FusionAuthIdpSteamApplicationConfigurationArray{
//					&fusionauth.FusionAuthIdpSteamApplicationConfigurationArgs{
//						ApplicationId:      pulumi.Any(fusionauth_application.GPS_Insight.Id),
//						CreateRegistration: pulumi.Bool(true),
//						Enabled:            pulumi.Bool(true),
//					},
//				},
//				ButtonText: pulumi.String("Login with Steam"),
//				ClientId:   pulumi.String("0eb1ce3c-2fb1-4ae9-b361-d49fc6e764cc"),
//				Scope:      pulumi.String("Xboxlive.signin Xboxlive.offline_access"),
//				WebApiKey:  pulumi.String("HG0A97ACKPQ5ZLPU0007BN6674OA25TY"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthIdpSteam struct {
	pulumi.CustomResourceState

	// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
	ApiMode pulumi.StringPtrOutput `pulumi:"apiMode"`
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSteamApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrOutput `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Steam.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSteamTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
	// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	WebApiKey pulumi.StringOutput `pulumi:"webApiKey"`
}

// NewFusionAuthIdpSteam registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpSteam(ctx *pulumi.Context,
	name string, args *FusionAuthIdpSteamArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpSteam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ButtonText == nil {
		return nil, errors.New("invalid value for required argument 'ButtonText'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.WebApiKey == nil {
		return nil, errors.New("invalid value for required argument 'WebApiKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpSteam
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpSteam:FusionAuthIdpSteam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpSteam gets an existing FusionAuthIdpSteam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpSteam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpSteamState, opts ...pulumi.ResourceOption) (*FusionAuthIdpSteam, error) {
	var resource FusionAuthIdpSteam
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpSteam:FusionAuthIdpSteam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpSteam resources.
type fusionAuthIdpSteamState struct {
	// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
	ApiMode *string `pulumi:"apiMode"`
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSteamApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	ClientId *string `pulumi:"clientId"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Steam.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSteamTenantConfiguration `pulumi:"tenantConfigurations"`
	// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	WebApiKey *string `pulumi:"webApiKey"`
}

type FusionAuthIdpSteamState struct {
	// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
	ApiMode pulumi.StringPtrInput
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSteamApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	ClientId pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from Steam.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSteamTenantConfigurationArrayInput
	// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	WebApiKey pulumi.StringPtrInput
}

func (FusionAuthIdpSteamState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSteamState)(nil)).Elem()
}

type fusionAuthIdpSteamArgs struct {
	// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
	ApiMode *string `pulumi:"apiMode"`
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSteamApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	ClientId string `pulumi:"clientId"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The top-level scope that you are requesting from Steam.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSteamTenantConfiguration `pulumi:"tenantConfigurations"`
	// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	WebApiKey string `pulumi:"webApiKey"`
}

// The set of arguments for constructing a FusionAuthIdpSteam resource.
type FusionAuthIdpSteamArgs struct {
	// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
	ApiMode pulumi.StringPtrInput
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSteamApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	ClientId pulumi.StringInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The top-level scope that you are requesting from Steam.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSteamTenantConfigurationArrayInput
	// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
	WebApiKey pulumi.StringInput
}

func (FusionAuthIdpSteamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSteamArgs)(nil)).Elem()
}

type FusionAuthIdpSteamInput interface {
	pulumi.Input

	ToFusionAuthIdpSteamOutput() FusionAuthIdpSteamOutput
	ToFusionAuthIdpSteamOutputWithContext(ctx context.Context) FusionAuthIdpSteamOutput
}

func (*FusionAuthIdpSteam) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSteam)(nil)).Elem()
}

func (i *FusionAuthIdpSteam) ToFusionAuthIdpSteamOutput() FusionAuthIdpSteamOutput {
	return i.ToFusionAuthIdpSteamOutputWithContext(context.Background())
}

func (i *FusionAuthIdpSteam) ToFusionAuthIdpSteamOutputWithContext(ctx context.Context) FusionAuthIdpSteamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSteamOutput)
}

// FusionAuthIdpSteamArrayInput is an input type that accepts FusionAuthIdpSteamArray and FusionAuthIdpSteamArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpSteamArrayInput` via:
//
//	FusionAuthIdpSteamArray{ FusionAuthIdpSteamArgs{...} }
type FusionAuthIdpSteamArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpSteamArrayOutput() FusionAuthIdpSteamArrayOutput
	ToFusionAuthIdpSteamArrayOutputWithContext(context.Context) FusionAuthIdpSteamArrayOutput
}

type FusionAuthIdpSteamArray []FusionAuthIdpSteamInput

func (FusionAuthIdpSteamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSteam)(nil)).Elem()
}

func (i FusionAuthIdpSteamArray) ToFusionAuthIdpSteamArrayOutput() FusionAuthIdpSteamArrayOutput {
	return i.ToFusionAuthIdpSteamArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpSteamArray) ToFusionAuthIdpSteamArrayOutputWithContext(ctx context.Context) FusionAuthIdpSteamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSteamArrayOutput)
}

// FusionAuthIdpSteamMapInput is an input type that accepts FusionAuthIdpSteamMap and FusionAuthIdpSteamMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpSteamMapInput` via:
//
//	FusionAuthIdpSteamMap{ "key": FusionAuthIdpSteamArgs{...} }
type FusionAuthIdpSteamMapInput interface {
	pulumi.Input

	ToFusionAuthIdpSteamMapOutput() FusionAuthIdpSteamMapOutput
	ToFusionAuthIdpSteamMapOutputWithContext(context.Context) FusionAuthIdpSteamMapOutput
}

type FusionAuthIdpSteamMap map[string]FusionAuthIdpSteamInput

func (FusionAuthIdpSteamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSteam)(nil)).Elem()
}

func (i FusionAuthIdpSteamMap) ToFusionAuthIdpSteamMapOutput() FusionAuthIdpSteamMapOutput {
	return i.ToFusionAuthIdpSteamMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpSteamMap) ToFusionAuthIdpSteamMapOutputWithContext(ctx context.Context) FusionAuthIdpSteamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSteamMapOutput)
}

type FusionAuthIdpSteamOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSteamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSteam)(nil)).Elem()
}

func (o FusionAuthIdpSteamOutput) ToFusionAuthIdpSteamOutput() FusionAuthIdpSteamOutput {
	return o
}

func (o FusionAuthIdpSteamOutput) ToFusionAuthIdpSteamOutputWithContext(ctx context.Context) FusionAuthIdpSteamOutput {
	return o
}

// Determines which Steam API to utilize. The possible values are: `Partner` and `Public`
func (o FusionAuthIdpSteamOutput) ApiMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringPtrOutput { return v.ApiMode }).(pulumi.StringPtrOutput)
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpSteamOutput) ApplicationConfigurations() FusionAuthIdpSteamApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) FusionAuthIdpSteamApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpSteamApplicationConfigurationArrayOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpSteamOutput) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// The top-level Steam client id for your Application. This value is retrieved from the Steam developer website when you setup your Steam developer account.
func (o FusionAuthIdpSteamOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpSteamOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpSteamOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
func (o FusionAuthIdpSteamOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringPtrOutput { return v.IdpId }).(pulumi.StringPtrOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpSteamOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpSteamOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The top-level scope that you are requesting from Steam.
func (o FusionAuthIdpSteamOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpSteamOutput) TenantConfigurations() FusionAuthIdpSteamTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) FusionAuthIdpSteamTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpSteamTenantConfigurationArrayOutput)
}

// The top-level web API key to use with the Steam Identity Provider when retrieving the player summary info. This value is retrieved from the Steam developer website when you setup your Steam developer account.
func (o FusionAuthIdpSteamOutput) WebApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSteam) pulumi.StringOutput { return v.WebApiKey }).(pulumi.StringOutput)
}

type FusionAuthIdpSteamArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSteamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSteam)(nil)).Elem()
}

func (o FusionAuthIdpSteamArrayOutput) ToFusionAuthIdpSteamArrayOutput() FusionAuthIdpSteamArrayOutput {
	return o
}

func (o FusionAuthIdpSteamArrayOutput) ToFusionAuthIdpSteamArrayOutputWithContext(ctx context.Context) FusionAuthIdpSteamArrayOutput {
	return o
}

func (o FusionAuthIdpSteamArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpSteamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpSteam {
		return vs[0].([]*FusionAuthIdpSteam)[vs[1].(int)]
	}).(FusionAuthIdpSteamOutput)
}

type FusionAuthIdpSteamMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSteamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSteam)(nil)).Elem()
}

func (o FusionAuthIdpSteamMapOutput) ToFusionAuthIdpSteamMapOutput() FusionAuthIdpSteamMapOutput {
	return o
}

func (o FusionAuthIdpSteamMapOutput) ToFusionAuthIdpSteamMapOutputWithContext(ctx context.Context) FusionAuthIdpSteamMapOutput {
	return o
}

func (o FusionAuthIdpSteamMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpSteamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpSteam {
		return vs[0].(map[string]*FusionAuthIdpSteam)[vs[1].(string)]
	}).(FusionAuthIdpSteamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSteamInput)(nil)).Elem(), &FusionAuthIdpSteam{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSteamArrayInput)(nil)).Elem(), FusionAuthIdpSteamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSteamMapInput)(nil)).Elem(), FusionAuthIdpSteamMap{})
	pulumi.RegisterOutputType(FusionAuthIdpSteamOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpSteamArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpSteamMapOutput{})
}
