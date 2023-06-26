// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Registration Resource
//
// A registration is the association between a User and an Application that they log into.
//
// [Registrations API](https://fusionauth.io/docs/v1/tech/apis/registrations)
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
//			_, err := fusionauth.NewFusionAuthRegistration(ctx, "example", &fusionauth.FusionAuthRegistrationArgs{
//				UserId:        pulumi.Any(fusionauth_user.Example.Id),
//				ApplicationId: pulumi.Any(data.Fusionauth_application.FusionAuth.Id),
//				Roles: pulumi.StringArray{
//					pulumi.String("admin"),
//				},
//				Username: pulumi.String("theadmin"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthRegistration struct {
	pulumi.CustomResourceState

	// The Id of the Application that this registration is for.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
	AuthenticationToken pulumi.StringOutput `pulumi:"authenticationToken"`
	// An object that can hold any information about the User for this registration that should be persisted.
	Data pulumi.MapOutput `pulumi:"data"`
	// Determines if FusionAuth should generate an Authentication Token for this registration.
	GenerateAuthenticationToken pulumi.BoolPtrOutput `pulumi:"generateAuthenticationToken"`
	// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
	PreferredLanguages pulumi.StringArrayOutput `pulumi:"preferredLanguages"`
	// The list of roles that the User has for this registration.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
	SkipRegistrationValidation pulumi.BoolPtrOutput `pulumi:"skipRegistrationValidation"`
	// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// The Id of the User that is registering for the Application.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The username of the User for this registration only.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewFusionAuthRegistration registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthRegistration(ctx *pulumi.Context,
	name string, args *FusionAuthRegistrationArgs, opts ...pulumi.ResourceOption) (*FusionAuthRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.AuthenticationToken != nil {
		args.AuthenticationToken = pulumi.ToSecret(args.AuthenticationToken).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authenticationToken",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthRegistration
	err := ctx.RegisterResource("fusionauth:index/fusionAuthRegistration:FusionAuthRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthRegistration gets an existing FusionAuthRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthRegistrationState, opts ...pulumi.ResourceOption) (*FusionAuthRegistration, error) {
	var resource FusionAuthRegistration
	err := ctx.ReadResource("fusionauth:index/fusionAuthRegistration:FusionAuthRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthRegistration resources.
type fusionAuthRegistrationState struct {
	// The Id of the Application that this registration is for.
	ApplicationId *string `pulumi:"applicationId"`
	// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
	AuthenticationToken *string `pulumi:"authenticationToken"`
	// An object that can hold any information about the User for this registration that should be persisted.
	Data map[string]interface{} `pulumi:"data"`
	// Determines if FusionAuth should generate an Authentication Token for this registration.
	GenerateAuthenticationToken *bool `pulumi:"generateAuthenticationToken"`
	// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
	PreferredLanguages []string `pulumi:"preferredLanguages"`
	// The list of roles that the User has for this registration.
	Roles []string `pulumi:"roles"`
	// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
	SkipRegistrationValidation *bool `pulumi:"skipRegistrationValidation"`
	// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
	Timezone *string `pulumi:"timezone"`
	// The Id of the User that is registering for the Application.
	UserId *string `pulumi:"userId"`
	// The username of the User for this registration only.
	Username *string `pulumi:"username"`
}

type FusionAuthRegistrationState struct {
	// The Id of the Application that this registration is for.
	ApplicationId pulumi.StringPtrInput
	// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
	AuthenticationToken pulumi.StringPtrInput
	// An object that can hold any information about the User for this registration that should be persisted.
	Data pulumi.MapInput
	// Determines if FusionAuth should generate an Authentication Token for this registration.
	GenerateAuthenticationToken pulumi.BoolPtrInput
	// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
	PreferredLanguages pulumi.StringArrayInput
	// The list of roles that the User has for this registration.
	Roles pulumi.StringArrayInput
	// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
	SkipRegistrationValidation pulumi.BoolPtrInput
	// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
	Timezone pulumi.StringPtrInput
	// The Id of the User that is registering for the Application.
	UserId pulumi.StringPtrInput
	// The username of the User for this registration only.
	Username pulumi.StringPtrInput
}

func (FusionAuthRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthRegistrationState)(nil)).Elem()
}

type fusionAuthRegistrationArgs struct {
	// The Id of the Application that this registration is for.
	ApplicationId string `pulumi:"applicationId"`
	// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
	AuthenticationToken *string `pulumi:"authenticationToken"`
	// An object that can hold any information about the User for this registration that should be persisted.
	Data map[string]interface{} `pulumi:"data"`
	// Determines if FusionAuth should generate an Authentication Token for this registration.
	GenerateAuthenticationToken *bool `pulumi:"generateAuthenticationToken"`
	// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
	PreferredLanguages []string `pulumi:"preferredLanguages"`
	// The list of roles that the User has for this registration.
	Roles []string `pulumi:"roles"`
	// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
	SkipRegistrationValidation *bool `pulumi:"skipRegistrationValidation"`
	// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
	Timezone *string `pulumi:"timezone"`
	// The Id of the User that is registering for the Application.
	UserId string `pulumi:"userId"`
	// The username of the User for this registration only.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a FusionAuthRegistration resource.
type FusionAuthRegistrationArgs struct {
	// The Id of the Application that this registration is for.
	ApplicationId pulumi.StringInput
	// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
	AuthenticationToken pulumi.StringPtrInput
	// An object that can hold any information about the User for this registration that should be persisted.
	Data pulumi.MapInput
	// Determines if FusionAuth should generate an Authentication Token for this registration.
	GenerateAuthenticationToken pulumi.BoolPtrInput
	// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
	PreferredLanguages pulumi.StringArrayInput
	// The list of roles that the User has for this registration.
	Roles pulumi.StringArrayInput
	// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
	SkipRegistrationValidation pulumi.BoolPtrInput
	// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
	Timezone pulumi.StringPtrInput
	// The Id of the User that is registering for the Application.
	UserId pulumi.StringInput
	// The username of the User for this registration only.
	Username pulumi.StringPtrInput
}

func (FusionAuthRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthRegistrationArgs)(nil)).Elem()
}

type FusionAuthRegistrationInput interface {
	pulumi.Input

	ToFusionAuthRegistrationOutput() FusionAuthRegistrationOutput
	ToFusionAuthRegistrationOutputWithContext(ctx context.Context) FusionAuthRegistrationOutput
}

func (*FusionAuthRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthRegistration)(nil)).Elem()
}

func (i *FusionAuthRegistration) ToFusionAuthRegistrationOutput() FusionAuthRegistrationOutput {
	return i.ToFusionAuthRegistrationOutputWithContext(context.Background())
}

func (i *FusionAuthRegistration) ToFusionAuthRegistrationOutputWithContext(ctx context.Context) FusionAuthRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthRegistrationOutput)
}

// FusionAuthRegistrationArrayInput is an input type that accepts FusionAuthRegistrationArray and FusionAuthRegistrationArrayOutput values.
// You can construct a concrete instance of `FusionAuthRegistrationArrayInput` via:
//
//	FusionAuthRegistrationArray{ FusionAuthRegistrationArgs{...} }
type FusionAuthRegistrationArrayInput interface {
	pulumi.Input

	ToFusionAuthRegistrationArrayOutput() FusionAuthRegistrationArrayOutput
	ToFusionAuthRegistrationArrayOutputWithContext(context.Context) FusionAuthRegistrationArrayOutput
}

type FusionAuthRegistrationArray []FusionAuthRegistrationInput

func (FusionAuthRegistrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthRegistration)(nil)).Elem()
}

func (i FusionAuthRegistrationArray) ToFusionAuthRegistrationArrayOutput() FusionAuthRegistrationArrayOutput {
	return i.ToFusionAuthRegistrationArrayOutputWithContext(context.Background())
}

func (i FusionAuthRegistrationArray) ToFusionAuthRegistrationArrayOutputWithContext(ctx context.Context) FusionAuthRegistrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthRegistrationArrayOutput)
}

// FusionAuthRegistrationMapInput is an input type that accepts FusionAuthRegistrationMap and FusionAuthRegistrationMapOutput values.
// You can construct a concrete instance of `FusionAuthRegistrationMapInput` via:
//
//	FusionAuthRegistrationMap{ "key": FusionAuthRegistrationArgs{...} }
type FusionAuthRegistrationMapInput interface {
	pulumi.Input

	ToFusionAuthRegistrationMapOutput() FusionAuthRegistrationMapOutput
	ToFusionAuthRegistrationMapOutputWithContext(context.Context) FusionAuthRegistrationMapOutput
}

type FusionAuthRegistrationMap map[string]FusionAuthRegistrationInput

func (FusionAuthRegistrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthRegistration)(nil)).Elem()
}

func (i FusionAuthRegistrationMap) ToFusionAuthRegistrationMapOutput() FusionAuthRegistrationMapOutput {
	return i.ToFusionAuthRegistrationMapOutputWithContext(context.Background())
}

func (i FusionAuthRegistrationMap) ToFusionAuthRegistrationMapOutputWithContext(ctx context.Context) FusionAuthRegistrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthRegistrationMapOutput)
}

type FusionAuthRegistrationOutput struct{ *pulumi.OutputState }

func (FusionAuthRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthRegistration)(nil)).Elem()
}

func (o FusionAuthRegistrationOutput) ToFusionAuthRegistrationOutput() FusionAuthRegistrationOutput {
	return o
}

func (o FusionAuthRegistrationOutput) ToFusionAuthRegistrationOutputWithContext(ctx context.Context) FusionAuthRegistrationOutput {
	return o
}

// The Id of the Application that this registration is for.
func (o FusionAuthRegistrationOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
func (o FusionAuthRegistrationOutput) AuthenticationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringOutput { return v.AuthenticationToken }).(pulumi.StringOutput)
}

// An object that can hold any information about the User for this registration that should be persisted.
func (o FusionAuthRegistrationOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

// Determines if FusionAuth should generate an Authentication Token for this registration.
func (o FusionAuthRegistrationOutput) GenerateAuthenticationToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.BoolPtrOutput { return v.GenerateAuthenticationToken }).(pulumi.BoolPtrOutput)
}

// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
func (o FusionAuthRegistrationOutput) PreferredLanguages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringArrayOutput { return v.PreferredLanguages }).(pulumi.StringArrayOutput)
}

// The list of roles that the User has for this registration.
func (o FusionAuthRegistrationOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
func (o FusionAuthRegistrationOutput) SkipRegistrationValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.BoolPtrOutput { return v.SkipRegistrationValidation }).(pulumi.BoolPtrOutput)
}

// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
func (o FusionAuthRegistrationOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// The Id of the User that is registering for the Application.
func (o FusionAuthRegistrationOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// The username of the User for this registration only.
func (o FusionAuthRegistrationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthRegistration) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type FusionAuthRegistrationArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthRegistrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthRegistration)(nil)).Elem()
}

func (o FusionAuthRegistrationArrayOutput) ToFusionAuthRegistrationArrayOutput() FusionAuthRegistrationArrayOutput {
	return o
}

func (o FusionAuthRegistrationArrayOutput) ToFusionAuthRegistrationArrayOutputWithContext(ctx context.Context) FusionAuthRegistrationArrayOutput {
	return o
}

func (o FusionAuthRegistrationArrayOutput) Index(i pulumi.IntInput) FusionAuthRegistrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthRegistration {
		return vs[0].([]*FusionAuthRegistration)[vs[1].(int)]
	}).(FusionAuthRegistrationOutput)
}

type FusionAuthRegistrationMapOutput struct{ *pulumi.OutputState }

func (FusionAuthRegistrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthRegistration)(nil)).Elem()
}

func (o FusionAuthRegistrationMapOutput) ToFusionAuthRegistrationMapOutput() FusionAuthRegistrationMapOutput {
	return o
}

func (o FusionAuthRegistrationMapOutput) ToFusionAuthRegistrationMapOutputWithContext(ctx context.Context) FusionAuthRegistrationMapOutput {
	return o
}

func (o FusionAuthRegistrationMapOutput) MapIndex(k pulumi.StringInput) FusionAuthRegistrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthRegistration {
		return vs[0].(map[string]*FusionAuthRegistration)[vs[1].(string)]
	}).(FusionAuthRegistrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthRegistrationInput)(nil)).Elem(), &FusionAuthRegistration{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthRegistrationArrayInput)(nil)).Elem(), FusionAuthRegistrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthRegistrationMapInput)(nil)).Elem(), FusionAuthRegistrationMap{})
	pulumi.RegisterOutputType(FusionAuthRegistrationOutput{})
	pulumi.RegisterOutputType(FusionAuthRegistrationArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthRegistrationMapOutput{})
}
