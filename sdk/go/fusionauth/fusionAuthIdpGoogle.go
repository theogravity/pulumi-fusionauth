// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Google Identity Provider Resource
//
// The Google identity provider type will use the Google OAuth v2.0 login API. it will provide a Login with Google button on FusionAuth’s login page that will leverage the Google login pop-up dialog. Additionally, this identity provider will call Google’s /oauth2/v3/tokeninfo API to load additional details about the user and store them in FusionAuth.
//
// The email address returned by the Google Token info API will be used to create or lookup the existing user. Additional claims returned by Google can be used to reconcile the User to FusionAuth by using a Google Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available claims returned by Google.
//
// FusionAuth will also store the Google accessToken that is returned from the login pop-up dialog in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
//
// [Google Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/google#create-the-google-identity-provider)
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
//			_, err := fusionauth.NewFusionAuthIdpGoogle(ctx, "google", &fusionauth.FusionAuthIdpGoogleArgs{
//				ApplicationConfigurations: fusionauth.FusionAuthIdpGoogleApplicationConfigurationArray{
//					&fusionauth.FusionAuthIdpGoogleApplicationConfigurationArgs{
//						ApplicationId:      pulumi.Any(fusionauth_application.Myapp.Id),
//						CreateRegistration: pulumi.Bool(true),
//						Enabled:            pulumi.Bool(true),
//					},
//				},
//				ButtonText:   pulumi.String("Login with Google"),
//				Debug:        pulumi.Bool(false),
//				ClientId:     pulumi.String("254311943570-8e2i2hds0qdnee4124socceeh2q2mtjl.apps.googleusercontent.com"),
//				ClientSecret: pulumi.String("BRr7x7xz_-cXxIFznBDIdxF1"),
//				Scope:        pulumi.String("profile"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthIdpGoogle struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpGoogleApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringOutput `pulumi:"buttonText"`
	// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The login method to use for this Identity Provider.
	LoginMethod pulumi.StringPtrOutput `pulumi:"loginMethod"`
	// The top-level scope that you are requesting from Google.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpGoogleTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
}

// NewFusionAuthIdpGoogle registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpGoogle(ctx *pulumi.Context,
	name string, args *FusionAuthIdpGoogleArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpGoogle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ButtonText == nil {
		return nil, errors.New("invalid value for required argument 'ButtonText'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpGoogle
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpGoogle gets an existing FusionAuthIdpGoogle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpGoogle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpGoogleState, opts ...pulumi.ResourceOption) (*FusionAuthIdpGoogle, error) {
	var resource FusionAuthIdpGoogle
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpGoogle resources.
type fusionAuthIdpGoogleState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpGoogleApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText *string `pulumi:"buttonText"`
	// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientId *string `pulumi:"clientId"`
	// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientSecret *string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The login method to use for this Identity Provider.
	LoginMethod *string `pulumi:"loginMethod"`
	// The top-level scope that you are requesting from Google.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpGoogleTenantConfiguration `pulumi:"tenantConfigurations"`
}

type FusionAuthIdpGoogleState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpGoogleApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringPtrInput
	// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientId pulumi.StringPtrInput
	// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientSecret pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The login method to use for this Identity Provider.
	LoginMethod pulumi.StringPtrInput
	// The top-level scope that you are requesting from Google.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpGoogleTenantConfigurationArrayInput
}

func (FusionAuthIdpGoogleState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpGoogleState)(nil)).Elem()
}

type fusionAuthIdpGoogleArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpGoogleApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText string `pulumi:"buttonText"`
	// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientId string `pulumi:"clientId"`
	// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientSecret *string `pulumi:"clientSecret"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The login method to use for this Identity Provider.
	LoginMethod *string `pulumi:"loginMethod"`
	// The top-level scope that you are requesting from Google.
	Scope *string `pulumi:"scope"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpGoogleTenantConfiguration `pulumi:"tenantConfigurations"`
}

// The set of arguments for constructing a FusionAuthIdpGoogle resource.
type FusionAuthIdpGoogleArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpGoogleApplicationConfigurationArrayInput
	// The top-level button text to use on the FusionAuth login page for this Identity Provider.
	ButtonText pulumi.StringInput
	// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientId pulumi.StringInput
	// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
	ClientSecret pulumi.StringPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The login method to use for this Identity Provider.
	LoginMethod pulumi.StringPtrInput
	// The top-level scope that you are requesting from Google.
	Scope pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpGoogleTenantConfigurationArrayInput
}

func (FusionAuthIdpGoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpGoogleArgs)(nil)).Elem()
}

type FusionAuthIdpGoogleInput interface {
	pulumi.Input

	ToFusionAuthIdpGoogleOutput() FusionAuthIdpGoogleOutput
	ToFusionAuthIdpGoogleOutputWithContext(ctx context.Context) FusionAuthIdpGoogleOutput
}

func (*FusionAuthIdpGoogle) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpGoogle)(nil)).Elem()
}

func (i *FusionAuthIdpGoogle) ToFusionAuthIdpGoogleOutput() FusionAuthIdpGoogleOutput {
	return i.ToFusionAuthIdpGoogleOutputWithContext(context.Background())
}

func (i *FusionAuthIdpGoogle) ToFusionAuthIdpGoogleOutputWithContext(ctx context.Context) FusionAuthIdpGoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpGoogleOutput)
}

// FusionAuthIdpGoogleArrayInput is an input type that accepts FusionAuthIdpGoogleArray and FusionAuthIdpGoogleArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpGoogleArrayInput` via:
//
//	FusionAuthIdpGoogleArray{ FusionAuthIdpGoogleArgs{...} }
type FusionAuthIdpGoogleArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpGoogleArrayOutput() FusionAuthIdpGoogleArrayOutput
	ToFusionAuthIdpGoogleArrayOutputWithContext(context.Context) FusionAuthIdpGoogleArrayOutput
}

type FusionAuthIdpGoogleArray []FusionAuthIdpGoogleInput

func (FusionAuthIdpGoogleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpGoogle)(nil)).Elem()
}

func (i FusionAuthIdpGoogleArray) ToFusionAuthIdpGoogleArrayOutput() FusionAuthIdpGoogleArrayOutput {
	return i.ToFusionAuthIdpGoogleArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpGoogleArray) ToFusionAuthIdpGoogleArrayOutputWithContext(ctx context.Context) FusionAuthIdpGoogleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpGoogleArrayOutput)
}

// FusionAuthIdpGoogleMapInput is an input type that accepts FusionAuthIdpGoogleMap and FusionAuthIdpGoogleMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpGoogleMapInput` via:
//
//	FusionAuthIdpGoogleMap{ "key": FusionAuthIdpGoogleArgs{...} }
type FusionAuthIdpGoogleMapInput interface {
	pulumi.Input

	ToFusionAuthIdpGoogleMapOutput() FusionAuthIdpGoogleMapOutput
	ToFusionAuthIdpGoogleMapOutputWithContext(context.Context) FusionAuthIdpGoogleMapOutput
}

type FusionAuthIdpGoogleMap map[string]FusionAuthIdpGoogleInput

func (FusionAuthIdpGoogleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpGoogle)(nil)).Elem()
}

func (i FusionAuthIdpGoogleMap) ToFusionAuthIdpGoogleMapOutput() FusionAuthIdpGoogleMapOutput {
	return i.ToFusionAuthIdpGoogleMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpGoogleMap) ToFusionAuthIdpGoogleMapOutputWithContext(ctx context.Context) FusionAuthIdpGoogleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpGoogleMapOutput)
}

type FusionAuthIdpGoogleOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpGoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpGoogle)(nil)).Elem()
}

func (o FusionAuthIdpGoogleOutput) ToFusionAuthIdpGoogleOutput() FusionAuthIdpGoogleOutput {
	return o
}

func (o FusionAuthIdpGoogleOutput) ToFusionAuthIdpGoogleOutputWithContext(ctx context.Context) FusionAuthIdpGoogleOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpGoogleOutput) ApplicationConfigurations() FusionAuthIdpGoogleApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) FusionAuthIdpGoogleApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpGoogleApplicationConfigurationArrayOutput)
}

// The top-level button text to use on the FusionAuth login page for this Identity Provider.
func (o FusionAuthIdpGoogleOutput) ButtonText() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringOutput { return v.ButtonText }).(pulumi.StringOutput)
}

// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
func (o FusionAuthIdpGoogleOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
func (o FusionAuthIdpGoogleOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
func (o FusionAuthIdpGoogleOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpGoogleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
func (o FusionAuthIdpGoogleOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpGoogleOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The login method to use for this Identity Provider.
func (o FusionAuthIdpGoogleOutput) LoginMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringPtrOutput { return v.LoginMethod }).(pulumi.StringPtrOutput)
}

// The top-level scope that you are requesting from Google.
func (o FusionAuthIdpGoogleOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpGoogleOutput) TenantConfigurations() FusionAuthIdpGoogleTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpGoogle) FusionAuthIdpGoogleTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpGoogleTenantConfigurationArrayOutput)
}

type FusionAuthIdpGoogleArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpGoogleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpGoogle)(nil)).Elem()
}

func (o FusionAuthIdpGoogleArrayOutput) ToFusionAuthIdpGoogleArrayOutput() FusionAuthIdpGoogleArrayOutput {
	return o
}

func (o FusionAuthIdpGoogleArrayOutput) ToFusionAuthIdpGoogleArrayOutputWithContext(ctx context.Context) FusionAuthIdpGoogleArrayOutput {
	return o
}

func (o FusionAuthIdpGoogleArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpGoogleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpGoogle {
		return vs[0].([]*FusionAuthIdpGoogle)[vs[1].(int)]
	}).(FusionAuthIdpGoogleOutput)
}

type FusionAuthIdpGoogleMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpGoogleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpGoogle)(nil)).Elem()
}

func (o FusionAuthIdpGoogleMapOutput) ToFusionAuthIdpGoogleMapOutput() FusionAuthIdpGoogleMapOutput {
	return o
}

func (o FusionAuthIdpGoogleMapOutput) ToFusionAuthIdpGoogleMapOutputWithContext(ctx context.Context) FusionAuthIdpGoogleMapOutput {
	return o
}

func (o FusionAuthIdpGoogleMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpGoogleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpGoogle {
		return vs[0].(map[string]*FusionAuthIdpGoogle)[vs[1].(string)]
	}).(FusionAuthIdpGoogleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpGoogleInput)(nil)).Elem(), &FusionAuthIdpGoogle{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpGoogleArrayInput)(nil)).Elem(), FusionAuthIdpGoogleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpGoogleMapInput)(nil)).Elem(), FusionAuthIdpGoogleMap{})
	pulumi.RegisterOutputType(FusionAuthIdpGoogleOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpGoogleArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpGoogleMapOutput{})
}
