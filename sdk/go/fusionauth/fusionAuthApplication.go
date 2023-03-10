// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Application Resource
//
// [Applications API](https://fusionauth.io/docs/v1/tech/apis/applications)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/v2/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthApplication(ctx, "forum", &fusionauth.FusionAuthApplicationArgs{
//				TenantId:                                pulumi.Any(fusionauth_tenant.Portal.Id),
//				AuthenticationTokenConfigurationEnabled: pulumi.Bool(false),
//				FormConfiguration: &fusionauth.FusionAuthApplicationFormConfigurationArgs{
//					AdminRegistrationFormId: pulumi.Any(fusionauth_form.Admin_registration.Id),
//					SelfServiceFormId:       pulumi.Any(fusionauth_form.Self_service.Id),
//				},
//				JwtConfiguration: &fusionauth.FusionAuthApplicationJwtConfigurationArgs{
//					AccessTokenId:          pulumi.Any(fusionauth_key.Access_token.Id),
//					Enabled:                pulumi.Bool(true),
//					IdTokenKeyId:           pulumi.Any(fusionauth_key.Id_token.Id),
//					RefreshTokenTtlMinutes: pulumi.Int(43200),
//					TtlSeconds:             pulumi.Int(3600),
//				},
//				LambdaConfiguration: &fusionauth.FusionAuthApplicationLambdaConfigurationArgs{
//					AccessTokenPopulateId: pulumi.Any(fusionauth_lambda.Token_populate.Id),
//					IdTokenPopulateId:     pulumi.Any(fusionauth_lambda.Id_token_populate.Id),
//				},
//				LoginConfiguration: &fusionauth.FusionAuthApplicationLoginConfigurationArgs{
//					AllowTokenRefresh:     pulumi.Bool(false),
//					GenerateRefreshTokens: pulumi.Bool(false),
//					RequireAuthentication: pulumi.Bool(true),
//				},
//				OauthConfiguration: &fusionauth.FusionAuthApplicationOauthConfigurationArgs{
//					AuthorizedOriginUrls: pulumi.StringArray{
//						pulumi.String("http://www.example.com/oauth-callback"),
//					},
//					EnabledGrants: pulumi.StringArray{
//						pulumi.String("authorization_code"),
//						pulumi.String("implicit"),
//					},
//					GenerateRefreshTokens:       pulumi.Bool(false),
//					LogoutBehavior:              pulumi.String("AllApplications"),
//					LogoutUrl:                   pulumi.String("http://www.example.com/logout"),
//					RequireClientAuthentication: pulumi.Bool(false),
//				},
//				RegistrationConfiguration: &fusionauth.FusionAuthApplicationRegistrationConfigurationArgs{
//					BirthDate: &fusionauth.FusionAuthApplicationRegistrationConfigurationBirthDateArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					ConfirmPassword: pulumi.Bool(false),
//					Enabled:         pulumi.Bool(false),
//					FirstName: &fusionauth.FusionAuthApplicationRegistrationConfigurationFirstNameArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					FullName: &fusionauth.FusionAuthApplicationRegistrationConfigurationFullNameArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					LastName: &fusionauth.FusionAuthApplicationRegistrationConfigurationLastNameArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					LoginIdType: pulumi.String(""),
//					MiddleName: &fusionauth.FusionAuthApplicationRegistrationConfigurationMiddleNameArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					MobilePhone: &fusionauth.FusionAuthApplicationRegistrationConfigurationMobilePhoneArgs{
//						Enabled:  pulumi.Bool(false),
//						Required: pulumi.Bool(false),
//					},
//					Type: pulumi.String(""),
//				},
//				PasswordlessConfigurationEnabled: pulumi.Bool(false),
//				RegistrationDeletePolicy: &fusionauth.FusionAuthApplicationRegistrationDeletePolicyArgs{
//					UnverifiedEnabled:              pulumi.Bool(true),
//					UnverifiedNumberOfDaysToRetain: pulumi.Int(30),
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
type FusionAuthApplication struct {
	pulumi.CustomResourceState

	AccessControlConfiguration FusionAuthApplicationAccessControlConfigurationOutput `pulumi:"accessControlConfiguration"`
	// The Id of the CleanSpeak application that usernames are sent to for moderation.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
	AuthenticationTokenConfigurationEnabled pulumi.BoolPtrOutput                               `pulumi:"authenticationTokenConfigurationEnabled"`
	CleanSpeakConfiguration                 FusionAuthApplicationCleanSpeakConfigurationOutput `pulumi:"cleanSpeakConfiguration"`
	// An object that can hold any information about the Application that should be persisted.
	Data                     pulumi.MapOutput                                    `pulumi:"data"`
	EmailConfiguration       FusionAuthApplicationEmailConfigurationOutput       `pulumi:"emailConfiguration"`
	FormConfiguration        FusionAuthApplicationFormConfigurationOutput        `pulumi:"formConfiguration"`
	JwtConfiguration         FusionAuthApplicationJwtConfigurationOutput         `pulumi:"jwtConfiguration"`
	LambdaConfiguration      FusionAuthApplicationLambdaConfigurationOutput      `pulumi:"lambdaConfiguration"`
	LoginConfiguration       FusionAuthApplicationLoginConfigurationOutput       `pulumi:"loginConfiguration"`
	MultiFactorConfiguration FusionAuthApplicationMultiFactorConfigurationOutput `pulumi:"multiFactorConfiguration"`
	// The name of the Application.
	Name               pulumi.StringOutput                           `pulumi:"name"`
	OauthConfiguration FusionAuthApplicationOauthConfigurationOutput `pulumi:"oauthConfiguration"`
	// Determines if passwordless login is enabled for this application.
	PasswordlessConfigurationEnabled pulumi.BoolPtrOutput                                 `pulumi:"passwordlessConfigurationEnabled"`
	RegistrationConfiguration        FusionAuthApplicationRegistrationConfigurationOutput `pulumi:"registrationConfiguration"`
	RegistrationDeletePolicy         FusionAuthApplicationRegistrationDeletePolicyOutput  `pulumi:"registrationDeletePolicy"`
	Samlv2Configuration              FusionAuthApplicationSamlv2ConfigurationOutput       `pulumi:"samlv2Configuration"`
	TenantId                         pulumi.StringOutput                                  `pulumi:"tenantId"`
	// The unique Id of the theme to be used to style the login page and other end user templates.
	ThemeId pulumi.StringPtrOutput `pulumi:"themeId"`
	// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
	VerificationEmailTemplateId pulumi.StringPtrOutput `pulumi:"verificationEmailTemplateId"`
	// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
	VerificationStrategy pulumi.StringPtrOutput `pulumi:"verificationStrategy"`
	// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
	VerifyRegistration pulumi.BoolPtrOutput `pulumi:"verifyRegistration"`
}

// NewFusionAuthApplication registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthApplication(ctx *pulumi.Context,
	name string, args *FusionAuthApplicationArgs, opts ...pulumi.ResourceOption) (*FusionAuthApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FusionAuthApplication
	err := ctx.RegisterResource("fusionauth:index/fusionAuthApplication:FusionAuthApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthApplication gets an existing FusionAuthApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthApplicationState, opts ...pulumi.ResourceOption) (*FusionAuthApplication, error) {
	var resource FusionAuthApplication
	err := ctx.ReadResource("fusionauth:index/fusionAuthApplication:FusionAuthApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthApplication resources.
type fusionAuthApplicationState struct {
	AccessControlConfiguration *FusionAuthApplicationAccessControlConfiguration `pulumi:"accessControlConfiguration"`
	// The Id of the CleanSpeak application that usernames are sent to for moderation.
	ApplicationId *string `pulumi:"applicationId"`
	// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
	AuthenticationTokenConfigurationEnabled *bool                                         `pulumi:"authenticationTokenConfigurationEnabled"`
	CleanSpeakConfiguration                 *FusionAuthApplicationCleanSpeakConfiguration `pulumi:"cleanSpeakConfiguration"`
	// An object that can hold any information about the Application that should be persisted.
	Data                     map[string]interface{}                         `pulumi:"data"`
	EmailConfiguration       *FusionAuthApplicationEmailConfiguration       `pulumi:"emailConfiguration"`
	FormConfiguration        *FusionAuthApplicationFormConfiguration        `pulumi:"formConfiguration"`
	JwtConfiguration         *FusionAuthApplicationJwtConfiguration         `pulumi:"jwtConfiguration"`
	LambdaConfiguration      *FusionAuthApplicationLambdaConfiguration      `pulumi:"lambdaConfiguration"`
	LoginConfiguration       *FusionAuthApplicationLoginConfiguration       `pulumi:"loginConfiguration"`
	MultiFactorConfiguration *FusionAuthApplicationMultiFactorConfiguration `pulumi:"multiFactorConfiguration"`
	// The name of the Application.
	Name               *string                                  `pulumi:"name"`
	OauthConfiguration *FusionAuthApplicationOauthConfiguration `pulumi:"oauthConfiguration"`
	// Determines if passwordless login is enabled for this application.
	PasswordlessConfigurationEnabled *bool                                           `pulumi:"passwordlessConfigurationEnabled"`
	RegistrationConfiguration        *FusionAuthApplicationRegistrationConfiguration `pulumi:"registrationConfiguration"`
	RegistrationDeletePolicy         *FusionAuthApplicationRegistrationDeletePolicy  `pulumi:"registrationDeletePolicy"`
	Samlv2Configuration              *FusionAuthApplicationSamlv2Configuration       `pulumi:"samlv2Configuration"`
	TenantId                         *string                                         `pulumi:"tenantId"`
	// The unique Id of the theme to be used to style the login page and other end user templates.
	ThemeId *string `pulumi:"themeId"`
	// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
	VerificationEmailTemplateId *string `pulumi:"verificationEmailTemplateId"`
	// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
	VerificationStrategy *string `pulumi:"verificationStrategy"`
	// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
	VerifyRegistration *bool `pulumi:"verifyRegistration"`
}

type FusionAuthApplicationState struct {
	AccessControlConfiguration FusionAuthApplicationAccessControlConfigurationPtrInput
	// The Id of the CleanSpeak application that usernames are sent to for moderation.
	ApplicationId pulumi.StringPtrInput
	// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
	AuthenticationTokenConfigurationEnabled pulumi.BoolPtrInput
	CleanSpeakConfiguration                 FusionAuthApplicationCleanSpeakConfigurationPtrInput
	// An object that can hold any information about the Application that should be persisted.
	Data                     pulumi.MapInput
	EmailConfiguration       FusionAuthApplicationEmailConfigurationPtrInput
	FormConfiguration        FusionAuthApplicationFormConfigurationPtrInput
	JwtConfiguration         FusionAuthApplicationJwtConfigurationPtrInput
	LambdaConfiguration      FusionAuthApplicationLambdaConfigurationPtrInput
	LoginConfiguration       FusionAuthApplicationLoginConfigurationPtrInput
	MultiFactorConfiguration FusionAuthApplicationMultiFactorConfigurationPtrInput
	// The name of the Application.
	Name               pulumi.StringPtrInput
	OauthConfiguration FusionAuthApplicationOauthConfigurationPtrInput
	// Determines if passwordless login is enabled for this application.
	PasswordlessConfigurationEnabled pulumi.BoolPtrInput
	RegistrationConfiguration        FusionAuthApplicationRegistrationConfigurationPtrInput
	RegistrationDeletePolicy         FusionAuthApplicationRegistrationDeletePolicyPtrInput
	Samlv2Configuration              FusionAuthApplicationSamlv2ConfigurationPtrInput
	TenantId                         pulumi.StringPtrInput
	// The unique Id of the theme to be used to style the login page and other end user templates.
	ThemeId pulumi.StringPtrInput
	// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
	VerificationEmailTemplateId pulumi.StringPtrInput
	// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
	VerificationStrategy pulumi.StringPtrInput
	// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
	VerifyRegistration pulumi.BoolPtrInput
}

func (FusionAuthApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthApplicationState)(nil)).Elem()
}

type fusionAuthApplicationArgs struct {
	AccessControlConfiguration *FusionAuthApplicationAccessControlConfiguration `pulumi:"accessControlConfiguration"`
	// The Id of the CleanSpeak application that usernames are sent to for moderation.
	ApplicationId *string `pulumi:"applicationId"`
	// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
	AuthenticationTokenConfigurationEnabled *bool                                         `pulumi:"authenticationTokenConfigurationEnabled"`
	CleanSpeakConfiguration                 *FusionAuthApplicationCleanSpeakConfiguration `pulumi:"cleanSpeakConfiguration"`
	// An object that can hold any information about the Application that should be persisted.
	Data                     map[string]interface{}                         `pulumi:"data"`
	EmailConfiguration       *FusionAuthApplicationEmailConfiguration       `pulumi:"emailConfiguration"`
	FormConfiguration        *FusionAuthApplicationFormConfiguration        `pulumi:"formConfiguration"`
	JwtConfiguration         *FusionAuthApplicationJwtConfiguration         `pulumi:"jwtConfiguration"`
	LambdaConfiguration      *FusionAuthApplicationLambdaConfiguration      `pulumi:"lambdaConfiguration"`
	LoginConfiguration       *FusionAuthApplicationLoginConfiguration       `pulumi:"loginConfiguration"`
	MultiFactorConfiguration *FusionAuthApplicationMultiFactorConfiguration `pulumi:"multiFactorConfiguration"`
	// The name of the Application.
	Name               *string                                  `pulumi:"name"`
	OauthConfiguration *FusionAuthApplicationOauthConfiguration `pulumi:"oauthConfiguration"`
	// Determines if passwordless login is enabled for this application.
	PasswordlessConfigurationEnabled *bool                                           `pulumi:"passwordlessConfigurationEnabled"`
	RegistrationConfiguration        *FusionAuthApplicationRegistrationConfiguration `pulumi:"registrationConfiguration"`
	RegistrationDeletePolicy         *FusionAuthApplicationRegistrationDeletePolicy  `pulumi:"registrationDeletePolicy"`
	Samlv2Configuration              *FusionAuthApplicationSamlv2Configuration       `pulumi:"samlv2Configuration"`
	TenantId                         string                                          `pulumi:"tenantId"`
	// The unique Id of the theme to be used to style the login page and other end user templates.
	ThemeId *string `pulumi:"themeId"`
	// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
	VerificationEmailTemplateId *string `pulumi:"verificationEmailTemplateId"`
	// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
	VerificationStrategy *string `pulumi:"verificationStrategy"`
	// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
	VerifyRegistration *bool `pulumi:"verifyRegistration"`
}

// The set of arguments for constructing a FusionAuthApplication resource.
type FusionAuthApplicationArgs struct {
	AccessControlConfiguration FusionAuthApplicationAccessControlConfigurationPtrInput
	// The Id of the CleanSpeak application that usernames are sent to for moderation.
	ApplicationId pulumi.StringPtrInput
	// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
	AuthenticationTokenConfigurationEnabled pulumi.BoolPtrInput
	CleanSpeakConfiguration                 FusionAuthApplicationCleanSpeakConfigurationPtrInput
	// An object that can hold any information about the Application that should be persisted.
	Data                     pulumi.MapInput
	EmailConfiguration       FusionAuthApplicationEmailConfigurationPtrInput
	FormConfiguration        FusionAuthApplicationFormConfigurationPtrInput
	JwtConfiguration         FusionAuthApplicationJwtConfigurationPtrInput
	LambdaConfiguration      FusionAuthApplicationLambdaConfigurationPtrInput
	LoginConfiguration       FusionAuthApplicationLoginConfigurationPtrInput
	MultiFactorConfiguration FusionAuthApplicationMultiFactorConfigurationPtrInput
	// The name of the Application.
	Name               pulumi.StringPtrInput
	OauthConfiguration FusionAuthApplicationOauthConfigurationPtrInput
	// Determines if passwordless login is enabled for this application.
	PasswordlessConfigurationEnabled pulumi.BoolPtrInput
	RegistrationConfiguration        FusionAuthApplicationRegistrationConfigurationPtrInput
	RegistrationDeletePolicy         FusionAuthApplicationRegistrationDeletePolicyPtrInput
	Samlv2Configuration              FusionAuthApplicationSamlv2ConfigurationPtrInput
	TenantId                         pulumi.StringInput
	// The unique Id of the theme to be used to style the login page and other end user templates.
	ThemeId pulumi.StringPtrInput
	// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
	VerificationEmailTemplateId pulumi.StringPtrInput
	// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
	VerificationStrategy pulumi.StringPtrInput
	// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
	VerifyRegistration pulumi.BoolPtrInput
}

func (FusionAuthApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthApplicationArgs)(nil)).Elem()
}

type FusionAuthApplicationInput interface {
	pulumi.Input

	ToFusionAuthApplicationOutput() FusionAuthApplicationOutput
	ToFusionAuthApplicationOutputWithContext(ctx context.Context) FusionAuthApplicationOutput
}

func (*FusionAuthApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthApplication)(nil)).Elem()
}

func (i *FusionAuthApplication) ToFusionAuthApplicationOutput() FusionAuthApplicationOutput {
	return i.ToFusionAuthApplicationOutputWithContext(context.Background())
}

func (i *FusionAuthApplication) ToFusionAuthApplicationOutputWithContext(ctx context.Context) FusionAuthApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApplicationOutput)
}

// FusionAuthApplicationArrayInput is an input type that accepts FusionAuthApplicationArray and FusionAuthApplicationArrayOutput values.
// You can construct a concrete instance of `FusionAuthApplicationArrayInput` via:
//
//	FusionAuthApplicationArray{ FusionAuthApplicationArgs{...} }
type FusionAuthApplicationArrayInput interface {
	pulumi.Input

	ToFusionAuthApplicationArrayOutput() FusionAuthApplicationArrayOutput
	ToFusionAuthApplicationArrayOutputWithContext(context.Context) FusionAuthApplicationArrayOutput
}

type FusionAuthApplicationArray []FusionAuthApplicationInput

func (FusionAuthApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthApplication)(nil)).Elem()
}

func (i FusionAuthApplicationArray) ToFusionAuthApplicationArrayOutput() FusionAuthApplicationArrayOutput {
	return i.ToFusionAuthApplicationArrayOutputWithContext(context.Background())
}

func (i FusionAuthApplicationArray) ToFusionAuthApplicationArrayOutputWithContext(ctx context.Context) FusionAuthApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApplicationArrayOutput)
}

// FusionAuthApplicationMapInput is an input type that accepts FusionAuthApplicationMap and FusionAuthApplicationMapOutput values.
// You can construct a concrete instance of `FusionAuthApplicationMapInput` via:
//
//	FusionAuthApplicationMap{ "key": FusionAuthApplicationArgs{...} }
type FusionAuthApplicationMapInput interface {
	pulumi.Input

	ToFusionAuthApplicationMapOutput() FusionAuthApplicationMapOutput
	ToFusionAuthApplicationMapOutputWithContext(context.Context) FusionAuthApplicationMapOutput
}

type FusionAuthApplicationMap map[string]FusionAuthApplicationInput

func (FusionAuthApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthApplication)(nil)).Elem()
}

func (i FusionAuthApplicationMap) ToFusionAuthApplicationMapOutput() FusionAuthApplicationMapOutput {
	return i.ToFusionAuthApplicationMapOutputWithContext(context.Background())
}

func (i FusionAuthApplicationMap) ToFusionAuthApplicationMapOutputWithContext(ctx context.Context) FusionAuthApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApplicationMapOutput)
}

type FusionAuthApplicationOutput struct{ *pulumi.OutputState }

func (FusionAuthApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthApplication)(nil)).Elem()
}

func (o FusionAuthApplicationOutput) ToFusionAuthApplicationOutput() FusionAuthApplicationOutput {
	return o
}

func (o FusionAuthApplicationOutput) ToFusionAuthApplicationOutputWithContext(ctx context.Context) FusionAuthApplicationOutput {
	return o
}

func (o FusionAuthApplicationOutput) AccessControlConfiguration() FusionAuthApplicationAccessControlConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationAccessControlConfigurationOutput {
		return v.AccessControlConfiguration
	}).(FusionAuthApplicationAccessControlConfigurationOutput)
}

// The Id of the CleanSpeak application that usernames are sent to for moderation.
func (o FusionAuthApplicationOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
func (o FusionAuthApplicationOutput) AuthenticationTokenConfigurationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.BoolPtrOutput { return v.AuthenticationTokenConfigurationEnabled }).(pulumi.BoolPtrOutput)
}

func (o FusionAuthApplicationOutput) CleanSpeakConfiguration() FusionAuthApplicationCleanSpeakConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationCleanSpeakConfigurationOutput {
		return v.CleanSpeakConfiguration
	}).(FusionAuthApplicationCleanSpeakConfigurationOutput)
}

// An object that can hold any information about the Application that should be persisted.
func (o FusionAuthApplicationOutput) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

func (o FusionAuthApplicationOutput) EmailConfiguration() FusionAuthApplicationEmailConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationEmailConfigurationOutput {
		return v.EmailConfiguration
	}).(FusionAuthApplicationEmailConfigurationOutput)
}

func (o FusionAuthApplicationOutput) FormConfiguration() FusionAuthApplicationFormConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationFormConfigurationOutput {
		return v.FormConfiguration
	}).(FusionAuthApplicationFormConfigurationOutput)
}

func (o FusionAuthApplicationOutput) JwtConfiguration() FusionAuthApplicationJwtConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationJwtConfigurationOutput { return v.JwtConfiguration }).(FusionAuthApplicationJwtConfigurationOutput)
}

func (o FusionAuthApplicationOutput) LambdaConfiguration() FusionAuthApplicationLambdaConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationLambdaConfigurationOutput {
		return v.LambdaConfiguration
	}).(FusionAuthApplicationLambdaConfigurationOutput)
}

func (o FusionAuthApplicationOutput) LoginConfiguration() FusionAuthApplicationLoginConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationLoginConfigurationOutput {
		return v.LoginConfiguration
	}).(FusionAuthApplicationLoginConfigurationOutput)
}

func (o FusionAuthApplicationOutput) MultiFactorConfiguration() FusionAuthApplicationMultiFactorConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationMultiFactorConfigurationOutput {
		return v.MultiFactorConfiguration
	}).(FusionAuthApplicationMultiFactorConfigurationOutput)
}

// The name of the Application.
func (o FusionAuthApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FusionAuthApplicationOutput) OauthConfiguration() FusionAuthApplicationOauthConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationOauthConfigurationOutput {
		return v.OauthConfiguration
	}).(FusionAuthApplicationOauthConfigurationOutput)
}

// Determines if passwordless login is enabled for this application.
func (o FusionAuthApplicationOutput) PasswordlessConfigurationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.BoolPtrOutput { return v.PasswordlessConfigurationEnabled }).(pulumi.BoolPtrOutput)
}

func (o FusionAuthApplicationOutput) RegistrationConfiguration() FusionAuthApplicationRegistrationConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationRegistrationConfigurationOutput {
		return v.RegistrationConfiguration
	}).(FusionAuthApplicationRegistrationConfigurationOutput)
}

func (o FusionAuthApplicationOutput) RegistrationDeletePolicy() FusionAuthApplicationRegistrationDeletePolicyOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationRegistrationDeletePolicyOutput {
		return v.RegistrationDeletePolicy
	}).(FusionAuthApplicationRegistrationDeletePolicyOutput)
}

func (o FusionAuthApplicationOutput) Samlv2Configuration() FusionAuthApplicationSamlv2ConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthApplication) FusionAuthApplicationSamlv2ConfigurationOutput {
		return v.Samlv2Configuration
	}).(FusionAuthApplicationSamlv2ConfigurationOutput)
}

func (o FusionAuthApplicationOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The unique Id of the theme to be used to style the login page and other end user templates.
func (o FusionAuthApplicationOutput) ThemeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringPtrOutput { return v.ThemeId }).(pulumi.StringPtrOutput)
}

// The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
func (o FusionAuthApplicationOutput) VerificationEmailTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringPtrOutput { return v.VerificationEmailTemplateId }).(pulumi.StringPtrOutput)
}

// The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
func (o FusionAuthApplicationOutput) VerificationStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.StringPtrOutput { return v.VerificationStrategy }).(pulumi.StringPtrOutput)
}

// Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
func (o FusionAuthApplicationOutput) VerifyRegistration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthApplication) pulumi.BoolPtrOutput { return v.VerifyRegistration }).(pulumi.BoolPtrOutput)
}

type FusionAuthApplicationArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthApplication)(nil)).Elem()
}

func (o FusionAuthApplicationArrayOutput) ToFusionAuthApplicationArrayOutput() FusionAuthApplicationArrayOutput {
	return o
}

func (o FusionAuthApplicationArrayOutput) ToFusionAuthApplicationArrayOutputWithContext(ctx context.Context) FusionAuthApplicationArrayOutput {
	return o
}

func (o FusionAuthApplicationArrayOutput) Index(i pulumi.IntInput) FusionAuthApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthApplication {
		return vs[0].([]*FusionAuthApplication)[vs[1].(int)]
	}).(FusionAuthApplicationOutput)
}

type FusionAuthApplicationMapOutput struct{ *pulumi.OutputState }

func (FusionAuthApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthApplication)(nil)).Elem()
}

func (o FusionAuthApplicationMapOutput) ToFusionAuthApplicationMapOutput() FusionAuthApplicationMapOutput {
	return o
}

func (o FusionAuthApplicationMapOutput) ToFusionAuthApplicationMapOutputWithContext(ctx context.Context) FusionAuthApplicationMapOutput {
	return o
}

func (o FusionAuthApplicationMapOutput) MapIndex(k pulumi.StringInput) FusionAuthApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthApplication {
		return vs[0].(map[string]*FusionAuthApplication)[vs[1].(string)]
	}).(FusionAuthApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApplicationInput)(nil)).Elem(), &FusionAuthApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApplicationArrayInput)(nil)).Elem(), FusionAuthApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApplicationMapInput)(nil)).Elem(), FusionAuthApplicationMap{})
	pulumi.RegisterOutputType(FusionAuthApplicationOutput{})
	pulumi.RegisterOutputType(FusionAuthApplicationArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthApplicationMapOutput{})
}
