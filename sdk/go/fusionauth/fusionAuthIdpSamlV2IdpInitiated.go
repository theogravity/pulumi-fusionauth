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

type FusionAuthIdpSamlV2IdpInitiated struct {
	pulumi.CustomResourceState

	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayOutput `pulumi:"applicationConfigurations"`
	// The assertion configuration for the SAML v2 identity provider.
	AssertionConfiguration FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrOutput `pulumi:"assertionConfiguration"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
	// an Event Log will be created.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
	// identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrOutput `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrOutput `pulumi:"idpId"`
	// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
	// this value was required to be a URL.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
	// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
	// component).
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
	LambdaReconcileId pulumi.StringPtrOutput `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringOutput `pulumi:"linkingStrategy"`
	// The name of this SAML v2 identity provider. This is only used for display purposes.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayOutput `pulumi:"tenantConfigurations"`
	// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
	// the `emailClaim` will be used when linking user.
	UniqueIdClaim pulumi.StringPtrOutput `pulumi:"uniqueIdClaim"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
	// processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrOutput `pulumi:"useNameForEmail"`
	// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
	// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
	// LinkByUsernameForExistingUser
	UsernameClaim pulumi.StringPtrOutput `pulumi:"usernameClaim"`
}

// NewFusionAuthIdpSamlV2IdpInitiated registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthIdpSamlV2IdpInitiated(ctx *pulumi.Context,
	name string, args *FusionAuthIdpSamlV2IdpInitiatedArgs, opts ...pulumi.ResourceOption) (*FusionAuthIdpSamlV2IdpInitiated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthIdpSamlV2IdpInitiated
	err := ctx.RegisterResource("fusionauth:index/fusionAuthIdpSamlV2IdpInitiated:FusionAuthIdpSamlV2IdpInitiated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthIdpSamlV2IdpInitiated gets an existing FusionAuthIdpSamlV2IdpInitiated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthIdpSamlV2IdpInitiated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthIdpSamlV2IdpInitiatedState, opts ...pulumi.ResourceOption) (*FusionAuthIdpSamlV2IdpInitiated, error) {
	var resource FusionAuthIdpSamlV2IdpInitiated
	err := ctx.ReadResource("fusionauth:index/fusionAuthIdpSamlV2IdpInitiated:FusionAuthIdpSamlV2IdpInitiated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthIdpSamlV2IdpInitiated resources.
type fusionAuthIdpSamlV2IdpInitiatedState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSamlV2IdpInitiatedApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The assertion configuration for the SAML v2 identity provider.
	AssertionConfiguration *FusionAuthIdpSamlV2IdpInitiatedAssertionConfiguration `pulumi:"assertionConfiguration"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
	// an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
	// identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim *string `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
	// this value was required to be a URL.
	Issuer *string `pulumi:"issuer"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
	// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
	// component).
	KeyId *string `pulumi:"keyId"`
	// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of this SAML v2 identity provider. This is only used for display purposes.
	Name *string `pulumi:"name"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSamlV2IdpInitiatedTenantConfiguration `pulumi:"tenantConfigurations"`
	// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
	// the `emailClaim` will be used when linking user.
	UniqueIdClaim *string `pulumi:"uniqueIdClaim"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
	// processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail *bool `pulumi:"useNameForEmail"`
	// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
	// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
	// LinkByUsernameForExistingUser
	UsernameClaim *string `pulumi:"usernameClaim"`
}

type FusionAuthIdpSamlV2IdpInitiatedState struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayInput
	// The assertion configuration for the SAML v2 identity provider.
	AssertionConfiguration FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
	// an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
	// identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
	// this value was required to be a URL.
	Issuer pulumi.StringPtrInput
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
	// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
	// component).
	KeyId pulumi.StringPtrInput
	// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of this SAML v2 identity provider. This is only used for display purposes.
	Name pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayInput
	// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
	// the `emailClaim` will be used when linking user.
	UniqueIdClaim pulumi.StringPtrInput
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
	// processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrInput
	// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
	// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
	// LinkByUsernameForExistingUser
	UsernameClaim pulumi.StringPtrInput
}

func (FusionAuthIdpSamlV2IdpInitiatedState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSamlV2IdpInitiatedState)(nil)).Elem()
}

type fusionAuthIdpSamlV2IdpInitiatedArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations []FusionAuthIdpSamlV2IdpInitiatedApplicationConfiguration `pulumi:"applicationConfigurations"`
	// The assertion configuration for the SAML v2 identity provider.
	AssertionConfiguration *FusionAuthIdpSamlV2IdpInitiatedAssertionConfiguration `pulumi:"assertionConfiguration"`
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
	// an Event Log will be created.
	Debug *bool `pulumi:"debug"`
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
	// identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim *string `pulumi:"emailClaim"`
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled *bool `pulumi:"enabled"`
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId *string `pulumi:"idpId"`
	// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
	// this value was required to be a URL.
	Issuer string `pulumi:"issuer"`
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
	// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
	// component).
	KeyId string `pulumi:"keyId"`
	// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
	LambdaReconcileId *string `pulumi:"lambdaReconcileId"`
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy *string `pulumi:"linkingStrategy"`
	// The name of this SAML v2 identity provider. This is only used for display purposes.
	Name *string `pulumi:"name"`
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations []FusionAuthIdpSamlV2IdpInitiatedTenantConfiguration `pulumi:"tenantConfigurations"`
	// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
	// the `emailClaim` will be used when linking user.
	UniqueIdClaim *string `pulumi:"uniqueIdClaim"`
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
	// processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail *bool `pulumi:"useNameForEmail"`
	// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
	// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
	// LinkByUsernameForExistingUser
	UsernameClaim *string `pulumi:"usernameClaim"`
}

// The set of arguments for constructing a FusionAuthIdpSamlV2IdpInitiated resource.
type FusionAuthIdpSamlV2IdpInitiatedArgs struct {
	// The configuration for each Application that the identity provider is enabled for.
	ApplicationConfigurations FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayInput
	// The assertion configuration for the SAML v2 identity provider.
	AssertionConfiguration FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrInput
	// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
	// an Event Log will be created.
	Debug pulumi.BoolPtrInput
	// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
	// identity the user. If this is not set, the `useNameForEmail` flag must be true.
	EmailClaim pulumi.StringPtrInput
	// Determines if this provider is enabled. If it is false then it will be disabled globally.
	Enabled pulumi.BoolPtrInput
	// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
	IdpId pulumi.StringPtrInput
	// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
	// this value was required to be a URL.
	Issuer pulumi.StringInput
	// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
	// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
	// component).
	KeyId pulumi.StringInput
	// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
	LambdaReconcileId pulumi.StringPtrInput
	// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
	LinkingStrategy pulumi.StringPtrInput
	// The name of this SAML v2 identity provider. This is only used for display purposes.
	Name pulumi.StringPtrInput
	// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
	TenantConfigurations FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayInput
	// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
	// the `emailClaim` will be used when linking user.
	UniqueIdClaim pulumi.StringPtrInput
	// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
	// processing. If this is false, then the `emailClaim` property must be set.
	UseNameForEmail pulumi.BoolPtrInput
	// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
	// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
	// LinkByUsernameForExistingUser
	UsernameClaim pulumi.StringPtrInput
}

func (FusionAuthIdpSamlV2IdpInitiatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthIdpSamlV2IdpInitiatedArgs)(nil)).Elem()
}

type FusionAuthIdpSamlV2IdpInitiatedInput interface {
	pulumi.Input

	ToFusionAuthIdpSamlV2IdpInitiatedOutput() FusionAuthIdpSamlV2IdpInitiatedOutput
	ToFusionAuthIdpSamlV2IdpInitiatedOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedOutput
}

func (*FusionAuthIdpSamlV2IdpInitiated) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (i *FusionAuthIdpSamlV2IdpInitiated) ToFusionAuthIdpSamlV2IdpInitiatedOutput() FusionAuthIdpSamlV2IdpInitiatedOutput {
	return i.ToFusionAuthIdpSamlV2IdpInitiatedOutputWithContext(context.Background())
}

func (i *FusionAuthIdpSamlV2IdpInitiated) ToFusionAuthIdpSamlV2IdpInitiatedOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlV2IdpInitiatedOutput)
}

// FusionAuthIdpSamlV2IdpInitiatedArrayInput is an input type that accepts FusionAuthIdpSamlV2IdpInitiatedArray and FusionAuthIdpSamlV2IdpInitiatedArrayOutput values.
// You can construct a concrete instance of `FusionAuthIdpSamlV2IdpInitiatedArrayInput` via:
//
//	FusionAuthIdpSamlV2IdpInitiatedArray{ FusionAuthIdpSamlV2IdpInitiatedArgs{...} }
type FusionAuthIdpSamlV2IdpInitiatedArrayInput interface {
	pulumi.Input

	ToFusionAuthIdpSamlV2IdpInitiatedArrayOutput() FusionAuthIdpSamlV2IdpInitiatedArrayOutput
	ToFusionAuthIdpSamlV2IdpInitiatedArrayOutputWithContext(context.Context) FusionAuthIdpSamlV2IdpInitiatedArrayOutput
}

type FusionAuthIdpSamlV2IdpInitiatedArray []FusionAuthIdpSamlV2IdpInitiatedInput

func (FusionAuthIdpSamlV2IdpInitiatedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (i FusionAuthIdpSamlV2IdpInitiatedArray) ToFusionAuthIdpSamlV2IdpInitiatedArrayOutput() FusionAuthIdpSamlV2IdpInitiatedArrayOutput {
	return i.ToFusionAuthIdpSamlV2IdpInitiatedArrayOutputWithContext(context.Background())
}

func (i FusionAuthIdpSamlV2IdpInitiatedArray) ToFusionAuthIdpSamlV2IdpInitiatedArrayOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlV2IdpInitiatedArrayOutput)
}

// FusionAuthIdpSamlV2IdpInitiatedMapInput is an input type that accepts FusionAuthIdpSamlV2IdpInitiatedMap and FusionAuthIdpSamlV2IdpInitiatedMapOutput values.
// You can construct a concrete instance of `FusionAuthIdpSamlV2IdpInitiatedMapInput` via:
//
//	FusionAuthIdpSamlV2IdpInitiatedMap{ "key": FusionAuthIdpSamlV2IdpInitiatedArgs{...} }
type FusionAuthIdpSamlV2IdpInitiatedMapInput interface {
	pulumi.Input

	ToFusionAuthIdpSamlV2IdpInitiatedMapOutput() FusionAuthIdpSamlV2IdpInitiatedMapOutput
	ToFusionAuthIdpSamlV2IdpInitiatedMapOutputWithContext(context.Context) FusionAuthIdpSamlV2IdpInitiatedMapOutput
}

type FusionAuthIdpSamlV2IdpInitiatedMap map[string]FusionAuthIdpSamlV2IdpInitiatedInput

func (FusionAuthIdpSamlV2IdpInitiatedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (i FusionAuthIdpSamlV2IdpInitiatedMap) ToFusionAuthIdpSamlV2IdpInitiatedMapOutput() FusionAuthIdpSamlV2IdpInitiatedMapOutput {
	return i.ToFusionAuthIdpSamlV2IdpInitiatedMapOutputWithContext(context.Background())
}

func (i FusionAuthIdpSamlV2IdpInitiatedMap) ToFusionAuthIdpSamlV2IdpInitiatedMapOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthIdpSamlV2IdpInitiatedMapOutput)
}

type FusionAuthIdpSamlV2IdpInitiatedOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlV2IdpInitiatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (o FusionAuthIdpSamlV2IdpInitiatedOutput) ToFusionAuthIdpSamlV2IdpInitiatedOutput() FusionAuthIdpSamlV2IdpInitiatedOutput {
	return o
}

func (o FusionAuthIdpSamlV2IdpInitiatedOutput) ToFusionAuthIdpSamlV2IdpInitiatedOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedOutput {
	return o
}

// The configuration for each Application that the identity provider is enabled for.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) ApplicationConfigurations() FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayOutput {
		return v.ApplicationConfigurations
	}).(FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArrayOutput)
}

// The assertion configuration for the SAML v2 identity provider.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) AssertionConfiguration() FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrOutput {
		return v.AssertionConfiguration
	}).(FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationPtrOutput)
}

// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
// an Event Log will be created.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
// identity the user. If this is not set, the `useNameForEmail` flag must be true.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) EmailClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringPtrOutput { return v.EmailClaim }).(pulumi.StringPtrOutput)
}

// Determines if this provider is enabled. If it is false then it will be disabled globally.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringPtrOutput { return v.IdpId }).(pulumi.StringPtrOutput)
}

// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
// this value was required to be a URL.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
// identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
// component).
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) LambdaReconcileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringPtrOutput { return v.LambdaReconcileId }).(pulumi.StringPtrOutput)
}

// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) LinkingStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringOutput { return v.LinkingStrategy }).(pulumi.StringOutput)
}

// The name of this SAML v2 identity provider. This is only used for display purposes.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) TenantConfigurations() FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayOutput {
		return v.TenantConfigurations
	}).(FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArrayOutput)
}

// The name of the unique claim in the SAML response that FusionAuth uses to uniquely link the user. If this is not set,
// the `emailClaim` will be used when linking user.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) UniqueIdClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringPtrOutput { return v.UniqueIdClaim }).(pulumi.StringPtrOutput)
}

// Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
// processing. If this is false, then the `emailClaim` property must be set.
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) UseNameForEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.BoolPtrOutput { return v.UseNameForEmail }).(pulumi.BoolPtrOutput)
}

// The name of the claim in the SAML response that FusionAuth uses to identity the username. If this is not set, the NameID
// value will be used to link a user. This property is required when `linkingStategy` is set to LinkByUsername or
// LinkByUsernameForExistingUser
func (o FusionAuthIdpSamlV2IdpInitiatedOutput) UsernameClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthIdpSamlV2IdpInitiated) pulumi.StringPtrOutput { return v.UsernameClaim }).(pulumi.StringPtrOutput)
}

type FusionAuthIdpSamlV2IdpInitiatedArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlV2IdpInitiatedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (o FusionAuthIdpSamlV2IdpInitiatedArrayOutput) ToFusionAuthIdpSamlV2IdpInitiatedArrayOutput() FusionAuthIdpSamlV2IdpInitiatedArrayOutput {
	return o
}

func (o FusionAuthIdpSamlV2IdpInitiatedArrayOutput) ToFusionAuthIdpSamlV2IdpInitiatedArrayOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedArrayOutput {
	return o
}

func (o FusionAuthIdpSamlV2IdpInitiatedArrayOutput) Index(i pulumi.IntInput) FusionAuthIdpSamlV2IdpInitiatedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthIdpSamlV2IdpInitiated {
		return vs[0].([]*FusionAuthIdpSamlV2IdpInitiated)[vs[1].(int)]
	}).(FusionAuthIdpSamlV2IdpInitiatedOutput)
}

type FusionAuthIdpSamlV2IdpInitiatedMapOutput struct{ *pulumi.OutputState }

func (FusionAuthIdpSamlV2IdpInitiatedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthIdpSamlV2IdpInitiated)(nil)).Elem()
}

func (o FusionAuthIdpSamlV2IdpInitiatedMapOutput) ToFusionAuthIdpSamlV2IdpInitiatedMapOutput() FusionAuthIdpSamlV2IdpInitiatedMapOutput {
	return o
}

func (o FusionAuthIdpSamlV2IdpInitiatedMapOutput) ToFusionAuthIdpSamlV2IdpInitiatedMapOutputWithContext(ctx context.Context) FusionAuthIdpSamlV2IdpInitiatedMapOutput {
	return o
}

func (o FusionAuthIdpSamlV2IdpInitiatedMapOutput) MapIndex(k pulumi.StringInput) FusionAuthIdpSamlV2IdpInitiatedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthIdpSamlV2IdpInitiated {
		return vs[0].(map[string]*FusionAuthIdpSamlV2IdpInitiated)[vs[1].(string)]
	}).(FusionAuthIdpSamlV2IdpInitiatedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlV2IdpInitiatedInput)(nil)).Elem(), &FusionAuthIdpSamlV2IdpInitiated{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlV2IdpInitiatedArrayInput)(nil)).Elem(), FusionAuthIdpSamlV2IdpInitiatedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthIdpSamlV2IdpInitiatedMapInput)(nil)).Elem(), FusionAuthIdpSamlV2IdpInitiatedMap{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlV2IdpInitiatedOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlV2IdpInitiatedArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthIdpSamlV2IdpInitiatedMapOutput{})
}
