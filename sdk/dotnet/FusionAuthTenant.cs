// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth
{
    /// <summary>
    /// ## # Tenant Resource
    /// 
    /// A FusionAuth Tenant is a named object that represents a discrete namespace for Users, Applications and Groups. A user is unique by email address or username within a tenant.
    /// 
    /// Tenants may be useful to support a multi-tenant application where you wish to use a single instance of FusionAuth but require the ability to have duplicate users across the tenants in your own application. In this scenario a user may exist multiple times with the same email address and different passwords across tenants.
    /// 
    /// Tenants may also be useful in a test or staging environment to allow multiple users to call APIs and create and modify users without possibility of collision.
    /// 
    /// [Tenants API](https://fusionauth.io/docs/v1/tech/apis/tenants)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Fusionauth = theogravity.Fusionauth;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Fusionauth.FusionAuthTenant("example", new()
    ///     {
    ///         ConnectorPolicies = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthTenantConnectorPolicyArgs
    ///             {
    ///                 ConnectorId = "b57b3d0f-f7a4-4831-a838-549717362ea8",
    ///                 Domains = new[]
    ///                 {
    ///                     "*",
    ///                 },
    ///                 Migrate = false,
    ///             },
    ///         },
    ///         EmailConfiguration = new Fusionauth.Inputs.FusionAuthTenantEmailConfigurationArgs
    ///         {
    ///             ForgotPasswordEmailTemplateId = fusionauth_email.ForgotPassword_Example.Id,
    ///             Host = "smtp.sendgrid.net",
    ///             Password = "password",
    ///             PasswordlessEmailTemplateId = fusionauth_email.PasswordlessLogin_Example.Id,
    ///             Port = 587,
    ///             Security = "TLS",
    ///             SetPasswordEmailTemplateId = fusionauth_email.SetupPassword_Example.Id,
    ///             Username = "username",
    ///             VerifyEmail = true,
    ///             VerifyEmailWhenChanged = true,
    ///         },
    ///         EventConfigurations = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.delete",
    ///                 Enabled = true,
    ///                 TransactionType = "None",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.create",
    ///                 Enabled = true,
    ///                 TransactionType = "None",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.update",
    ///                 Enabled = true,
    ///                 TransactionType = "None",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.deactivate",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.bulk.create",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.reactivate",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "jwt.refresh-token.revoke",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "jwt.refresh",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "jwt.public-key.update",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.login.success",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.login.failed",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.registration.create",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.registration.update",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.registration.delete",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.registration.verified",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.email.verified",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.identity-provider.link",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.identity-provider.unlink",
    ///                 Enabled = true,
    ///                 TransactionType = "Any",
    ///             },
    ///             new Fusionauth.Inputs.FusionAuthTenantEventConfigurationArgs
    ///             {
    ///                 Event = "user.password.breach",
    ///                 Enabled = false,
    ///                 TransactionType = "None",
    ///             },
    ///         },
    ///         ExternalIdentifierConfiguration = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationArgs
    ///         {
    ///             AuthorizationGrantIdTimeToLiveInSeconds = 30,
    ///             ChangePasswordIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationChangePasswordIdGeneratorArgs
    ///             {
    ///                 Length = 32,
    ///                 Type = "randomBytes",
    ///             },
    ///             ChangePasswordIdTimeToLiveInSeconds = 600,
    ///             DeviceCodeTimeToLiveInSeconds = 1800,
    ///             DeviceUserCodeIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGeneratorArgs
    ///             {
    ///                 Length = 6,
    ///                 Type = "randomAlphaNumeric",
    ///             },
    ///             EmailVerificationIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationIdGeneratorArgs
    ///             {
    ///                 Length = 32,
    ///                 Type = "randomBytes",
    ///             },
    ///             EmailVerificationIdTimeToLiveInSeconds = 86400,
    ///             EmailVerificationOneTimeCodeGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationOneTimeCodeGeneratorArgs
    ///             {
    ///                 Length = 6,
    ///                 Type = "randomAlphaNumeric",
    ///             },
    ///             ExternalAuthenticationIdTimeToLiveInSeconds = 300,
    ///             OneTimePasswordTimeToLiveInSeconds = 60,
    ///             PasswordlessLoginGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationPasswordlessLoginGeneratorArgs
    ///             {
    ///                 Length = 32,
    ///                 Type = "randomBytes",
    ///             },
    ///             PasswordlessLoginTimeToLiveInSeconds = 600,
    ///             RegistrationVerificationIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationIdGeneratorArgs
    ///             {
    ///                 Length = 32,
    ///                 Type = "randomBytes",
    ///             },
    ///             RegistrationVerificationIdTimeToLiveInSeconds = 86400,
    ///             RegistrationVerificationOneTimeCodeGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationOneTimeCodeGeneratorArgs
    ///             {
    ///                 Length = 6,
    ///                 Type = "randomAlphaNumeric",
    ///             },
    ///             SamlV2AuthnRequestIdTtlSeconds = 300,
    ///             SetupPasswordIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs
    ///             {
    ///                 Length = 32,
    ///                 Type = "randomBytes",
    ///             },
    ///             SetupPasswordIdTimeToLiveInSeconds = 86400,
    ///             TwoFactorIdTimeToLiveInSeconds = 300,
    ///             TwoFactorOneTimeCodeIdGenerator = new Fusionauth.Inputs.FusionAuthTenantExternalIdentifierConfigurationTwoFactorOneTimeCodeIdGeneratorArgs
    ///             {
    ///                 Length = 6,
    ///                 Type = "randomDigits",
    ///             },
    ///             TwoFactorTrustIdTimeToLiveInSeconds = 2592000,
    ///         },
    ///         FailedAuthenticationConfiguration = new Fusionauth.Inputs.FusionAuthTenantFailedAuthenticationConfigurationArgs
    ///         {
    ///             ActionDuration = 3,
    ///             ActionDurationUnit = "MINUTES",
    ///             ResetCountInSeconds = 60,
    ///             TooManyAttempts = 5,
    ///         },
    ///         FamilyConfiguration = new Fusionauth.Inputs.FusionAuthTenantFamilyConfigurationArgs
    ///         {
    ///             AllowChildRegistrations = true,
    ///             DeleteOrphanedAccounts = false,
    ///             DeleteOrphanedAccountsDays = 30,
    ///             Enabled = true,
    ///             MaximumChildAge = 12,
    ///             MinimumOwnerAge = 21,
    ///             ParentEmailRequired = false,
    ///         },
    ///         FormConfiguration = new Fusionauth.Inputs.FusionAuthTenantFormConfigurationArgs
    ///         {
    ///             AdminUserFormId = "e92751a5-25f4-4bca-ad91-66cdf67725d2",
    ///         },
    ///         HttpSessionMaxInactiveInterval = 3600,
    ///         Issuer = "https://example.com",
    ///         JwtConfigurations = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthTenantJwtConfigurationArgs
    ///             {
    ///                 AccessTokenKeyId = fusionauth_key.Accesstoken.Id,
    ///                 IdTokenKeyId = fusionauth_key.Idtoken.Id,
    ///                 RefreshTokenTimeToLiveInMinutes = 43200,
    ///                 TimeToLiveInSeconds = 3600,
    ///             },
    ///         },
    ///         LoginConfiguration = new Fusionauth.Inputs.FusionAuthTenantLoginConfigurationArgs
    ///         {
    ///             RequireAuthentication = true,
    ///         },
    ///         MaximumPasswordAge = new Fusionauth.Inputs.FusionAuthTenantMaximumPasswordAgeArgs
    ///         {
    ///             Days = 180,
    ///             Enabled = false,
    ///         },
    ///         MinimumPasswordAge = new Fusionauth.Inputs.FusionAuthTenantMinimumPasswordAgeArgs
    ///         {
    ///             Enabled = false,
    ///             Seconds = 30,
    ///         },
    ///         OauthConfigurations = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthTenantOauthConfigurationArgs
    ///             {
    ///                 ClientCredentialsAccessTokenPopulateLambdaId = fusionauth_lambda.Client_jwt_populate.Id,
    ///             },
    ///         },
    ///         PasswordEncryptionConfigurations = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthTenantPasswordEncryptionConfigurationArgs
    ///             {
    ///                 EncryptionScheme = "salted-pbkdf2-hmac-sha256",
    ///                 EncryptionSchemeFactor = 24000,
    ///                 ModifyEncryptionSchemeOnLogin = false,
    ///             },
    ///         },
    ///         PasswordValidationRules = new Fusionauth.Inputs.FusionAuthTenantPasswordValidationRulesArgs
    ///         {
    ///             MaxLength = 256,
    ///             MinLength = 7,
    ///             RememberPreviousPasswords = new Fusionauth.Inputs.FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsArgs
    ///             {
    ///                 Count = 1,
    ///                 Enabled = false,
    ///             },
    ///             RequiredMixedCase = false,
    ///             RequireNonAlpha = false,
    ///             RequireNumber = false,
    ///             ValidateOnLogin = false,
    ///         },
    ///         ThemeId = fusionauth_theme.Example_theme.Id,
    ///         UserDeletePolicy = new Fusionauth.Inputs.FusionAuthTenantUserDeletePolicyArgs
    ///         {
    ///             UnverifiedEnabled = false,
    ///             UnverifiedNumberOfDaysToRetain = 30,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthTenant:FusionAuthTenant")]
    public partial class FusionAuthTenant : global::Pulumi.CustomResource
    {
        [Output("accessControlConfiguration")]
        public Output<Outputs.FusionAuthTenantAccessControlConfiguration?> AccessControlConfiguration { get; private set; } = null!;

        [Output("captchaConfiguration")]
        public Output<Outputs.FusionAuthTenantCaptchaConfiguration> CaptchaConfiguration { get; private set; } = null!;

        /// <summary>
        /// A list of Connector policies. Users will be authenticated against Connectors in order. Each Connector can be included in this list at most once and must exist.
        /// </summary>
        [Output("connectorPolicies")]
        public Output<ImmutableArray<Outputs.FusionAuthTenantConnectorPolicy>> ConnectorPolicies { get; private set; } = null!;

        /// <summary>
        /// An object that can hold any information about the Tenant that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>?> Data { get; private set; } = null!;

        [Output("emailConfiguration")]
        public Output<Outputs.FusionAuthTenantEmailConfiguration> EmailConfiguration { get; private set; } = null!;

        [Output("eventConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthTenantEventConfiguration>> EventConfigurations { get; private set; } = null!;

        [Output("externalIdentifierConfiguration")]
        public Output<Outputs.FusionAuthTenantExternalIdentifierConfiguration> ExternalIdentifierConfiguration { get; private set; } = null!;

        [Output("failedAuthenticationConfiguration")]
        public Output<Outputs.FusionAuthTenantFailedAuthenticationConfiguration> FailedAuthenticationConfiguration { get; private set; } = null!;

        [Output("familyConfiguration")]
        public Output<Outputs.FusionAuthTenantFamilyConfiguration> FamilyConfiguration { get; private set; } = null!;

        [Output("formConfiguration")]
        public Output<Outputs.FusionAuthTenantFormConfiguration> FormConfiguration { get; private set; } = null!;

        /// <summary>
        /// Time in seconds until an inactive session will be invalidated. Used when creating a new session in the FusionAuth OAuth frontend.
        /// </summary>
        [Output("httpSessionMaxInactiveInterval")]
        public Output<int?> HttpSessionMaxInactiveInterval { get; private set; } = null!;

        /// <summary>
        /// The named issuer used to sign tokens, this is generally your public fully qualified domain.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        [Output("jwtConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthTenantJwtConfiguration>> JwtConfigurations { get; private set; } = null!;

        [Output("loginConfiguration")]
        public Output<Outputs.FusionAuthTenantLoginConfiguration?> LoginConfiguration { get; private set; } = null!;

        /// <summary>
        /// The logout redirect URL when sending the user’s browser to the /oauth2/logout URI of the FusionAuth Front End. This value is only used when a logout URL is not defined in your Application.
        /// </summary>
        [Output("logoutUrl")]
        public Output<string?> LogoutUrl { get; private set; } = null!;

        [Output("maximumPasswordAge")]
        public Output<Outputs.FusionAuthTenantMaximumPasswordAge> MaximumPasswordAge { get; private set; } = null!;

        [Output("minimumPasswordAge")]
        public Output<Outputs.FusionAuthTenantMinimumPasswordAge> MinimumPasswordAge { get; private set; } = null!;

        [Output("multiFactorConfiguration")]
        public Output<Outputs.FusionAuthTenantMultiFactorConfiguration> MultiFactorConfiguration { get; private set; } = null!;

        /// <summary>
        /// The unique name of the Tenant.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("oauthConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthTenantOauthConfiguration>> OauthConfigurations { get; private set; } = null!;

        [Output("passwordEncryptionConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthTenantPasswordEncryptionConfiguration>> PasswordEncryptionConfigurations { get; private set; } = null!;

        [Output("passwordValidationRules")]
        public Output<Outputs.FusionAuthTenantPasswordValidationRules> PasswordValidationRules { get; private set; } = null!;

        /// <summary>
        /// The optional Id of an existing Tenant to make a copy of. If present, the tenant.id and tenant.name values of the request body will be applied to the new Tenant, all other values will be copied from the source Tenant to the new Tenant.
        /// </summary>
        [Output("sourceTenantId")]
        public Output<string?> SourceTenantId { get; private set; } = null!;

        /// <summary>
        /// The Id to use for the new Tenant. If not specified a secure random UUID will be generated.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The unique Id of the theme to be used to style the login page and other end user templates.
        /// </summary>
        [Output("themeId")]
        public Output<string> ThemeId { get; private set; } = null!;

        [Output("userDeletePolicy")]
        public Output<Outputs.FusionAuthTenantUserDeletePolicy> UserDeletePolicy { get; private set; } = null!;

        [Output("usernameConfiguration")]
        public Output<Outputs.FusionAuthTenantUsernameConfiguration> UsernameConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthTenant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthTenant(string name, FusionAuthTenantArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthTenant:FusionAuthTenant", name, args ?? new FusionAuthTenantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthTenant(string name, Input<string> id, FusionAuthTenantState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthTenant:FusionAuthTenant", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/theogravity/pulumi-fusionauth/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthTenant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthTenant Get(string name, Input<string> id, FusionAuthTenantState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthTenant(name, id, state, options);
        }
    }

    public sealed class FusionAuthTenantArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessControlConfiguration")]
        public Input<Inputs.FusionAuthTenantAccessControlConfigurationArgs>? AccessControlConfiguration { get; set; }

        [Input("captchaConfiguration")]
        public Input<Inputs.FusionAuthTenantCaptchaConfigurationArgs>? CaptchaConfiguration { get; set; }

        [Input("connectorPolicies")]
        private InputList<Inputs.FusionAuthTenantConnectorPolicyArgs>? _connectorPolicies;

        /// <summary>
        /// A list of Connector policies. Users will be authenticated against Connectors in order. Each Connector can be included in this list at most once and must exist.
        /// </summary>
        public InputList<Inputs.FusionAuthTenantConnectorPolicyArgs> ConnectorPolicies
        {
            get => _connectorPolicies ?? (_connectorPolicies = new InputList<Inputs.FusionAuthTenantConnectorPolicyArgs>());
            set => _connectorPolicies = value;
        }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the Tenant that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        [Input("emailConfiguration", required: true)]
        public Input<Inputs.FusionAuthTenantEmailConfigurationArgs> EmailConfiguration { get; set; } = null!;

        [Input("eventConfigurations")]
        private InputList<Inputs.FusionAuthTenantEventConfigurationArgs>? _eventConfigurations;
        public InputList<Inputs.FusionAuthTenantEventConfigurationArgs> EventConfigurations
        {
            get => _eventConfigurations ?? (_eventConfigurations = new InputList<Inputs.FusionAuthTenantEventConfigurationArgs>());
            set => _eventConfigurations = value;
        }

        [Input("externalIdentifierConfiguration", required: true)]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationArgs> ExternalIdentifierConfiguration { get; set; } = null!;

        [Input("failedAuthenticationConfiguration")]
        public Input<Inputs.FusionAuthTenantFailedAuthenticationConfigurationArgs>? FailedAuthenticationConfiguration { get; set; }

        [Input("familyConfiguration")]
        public Input<Inputs.FusionAuthTenantFamilyConfigurationArgs>? FamilyConfiguration { get; set; }

        [Input("formConfiguration")]
        public Input<Inputs.FusionAuthTenantFormConfigurationArgs>? FormConfiguration { get; set; }

        /// <summary>
        /// Time in seconds until an inactive session will be invalidated. Used when creating a new session in the FusionAuth OAuth frontend.
        /// </summary>
        [Input("httpSessionMaxInactiveInterval")]
        public Input<int>? HttpSessionMaxInactiveInterval { get; set; }

        /// <summary>
        /// The named issuer used to sign tokens, this is generally your public fully qualified domain.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        [Input("jwtConfigurations", required: true)]
        private InputList<Inputs.FusionAuthTenantJwtConfigurationArgs>? _jwtConfigurations;
        public InputList<Inputs.FusionAuthTenantJwtConfigurationArgs> JwtConfigurations
        {
            get => _jwtConfigurations ?? (_jwtConfigurations = new InputList<Inputs.FusionAuthTenantJwtConfigurationArgs>());
            set => _jwtConfigurations = value;
        }

        [Input("loginConfiguration")]
        public Input<Inputs.FusionAuthTenantLoginConfigurationArgs>? LoginConfiguration { get; set; }

        /// <summary>
        /// The logout redirect URL when sending the user’s browser to the /oauth2/logout URI of the FusionAuth Front End. This value is only used when a logout URL is not defined in your Application.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        [Input("maximumPasswordAge")]
        public Input<Inputs.FusionAuthTenantMaximumPasswordAgeArgs>? MaximumPasswordAge { get; set; }

        [Input("minimumPasswordAge")]
        public Input<Inputs.FusionAuthTenantMinimumPasswordAgeArgs>? MinimumPasswordAge { get; set; }

        [Input("multiFactorConfiguration")]
        public Input<Inputs.FusionAuthTenantMultiFactorConfigurationArgs>? MultiFactorConfiguration { get; set; }

        /// <summary>
        /// The unique name of the Tenant.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("oauthConfigurations")]
        private InputList<Inputs.FusionAuthTenantOauthConfigurationArgs>? _oauthConfigurations;
        public InputList<Inputs.FusionAuthTenantOauthConfigurationArgs> OauthConfigurations
        {
            get => _oauthConfigurations ?? (_oauthConfigurations = new InputList<Inputs.FusionAuthTenantOauthConfigurationArgs>());
            set => _oauthConfigurations = value;
        }

        [Input("passwordEncryptionConfigurations")]
        private InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationArgs>? _passwordEncryptionConfigurations;
        public InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationArgs> PasswordEncryptionConfigurations
        {
            get => _passwordEncryptionConfigurations ?? (_passwordEncryptionConfigurations = new InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationArgs>());
            set => _passwordEncryptionConfigurations = value;
        }

        [Input("passwordValidationRules")]
        public Input<Inputs.FusionAuthTenantPasswordValidationRulesArgs>? PasswordValidationRules { get; set; }

        /// <summary>
        /// The optional Id of an existing Tenant to make a copy of. If present, the tenant.id and tenant.name values of the request body will be applied to the new Tenant, all other values will be copied from the source Tenant to the new Tenant.
        /// </summary>
        [Input("sourceTenantId")]
        public Input<string>? SourceTenantId { get; set; }

        /// <summary>
        /// The Id to use for the new Tenant. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The unique Id of the theme to be used to style the login page and other end user templates.
        /// </summary>
        [Input("themeId", required: true)]
        public Input<string> ThemeId { get; set; } = null!;

        [Input("userDeletePolicy")]
        public Input<Inputs.FusionAuthTenantUserDeletePolicyArgs>? UserDeletePolicy { get; set; }

        [Input("usernameConfiguration")]
        public Input<Inputs.FusionAuthTenantUsernameConfigurationArgs>? UsernameConfiguration { get; set; }

        public FusionAuthTenantArgs()
        {
        }
        public static new FusionAuthTenantArgs Empty => new FusionAuthTenantArgs();
    }

    public sealed class FusionAuthTenantState : global::Pulumi.ResourceArgs
    {
        [Input("accessControlConfiguration")]
        public Input<Inputs.FusionAuthTenantAccessControlConfigurationGetArgs>? AccessControlConfiguration { get; set; }

        [Input("captchaConfiguration")]
        public Input<Inputs.FusionAuthTenantCaptchaConfigurationGetArgs>? CaptchaConfiguration { get; set; }

        [Input("connectorPolicies")]
        private InputList<Inputs.FusionAuthTenantConnectorPolicyGetArgs>? _connectorPolicies;

        /// <summary>
        /// A list of Connector policies. Users will be authenticated against Connectors in order. Each Connector can be included in this list at most once and must exist.
        /// </summary>
        public InputList<Inputs.FusionAuthTenantConnectorPolicyGetArgs> ConnectorPolicies
        {
            get => _connectorPolicies ?? (_connectorPolicies = new InputList<Inputs.FusionAuthTenantConnectorPolicyGetArgs>());
            set => _connectorPolicies = value;
        }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the Tenant that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        [Input("emailConfiguration")]
        public Input<Inputs.FusionAuthTenantEmailConfigurationGetArgs>? EmailConfiguration { get; set; }

        [Input("eventConfigurations")]
        private InputList<Inputs.FusionAuthTenantEventConfigurationGetArgs>? _eventConfigurations;
        public InputList<Inputs.FusionAuthTenantEventConfigurationGetArgs> EventConfigurations
        {
            get => _eventConfigurations ?? (_eventConfigurations = new InputList<Inputs.FusionAuthTenantEventConfigurationGetArgs>());
            set => _eventConfigurations = value;
        }

        [Input("externalIdentifierConfiguration")]
        public Input<Inputs.FusionAuthTenantExternalIdentifierConfigurationGetArgs>? ExternalIdentifierConfiguration { get; set; }

        [Input("failedAuthenticationConfiguration")]
        public Input<Inputs.FusionAuthTenantFailedAuthenticationConfigurationGetArgs>? FailedAuthenticationConfiguration { get; set; }

        [Input("familyConfiguration")]
        public Input<Inputs.FusionAuthTenantFamilyConfigurationGetArgs>? FamilyConfiguration { get; set; }

        [Input("formConfiguration")]
        public Input<Inputs.FusionAuthTenantFormConfigurationGetArgs>? FormConfiguration { get; set; }

        /// <summary>
        /// Time in seconds until an inactive session will be invalidated. Used when creating a new session in the FusionAuth OAuth frontend.
        /// </summary>
        [Input("httpSessionMaxInactiveInterval")]
        public Input<int>? HttpSessionMaxInactiveInterval { get; set; }

        /// <summary>
        /// The named issuer used to sign tokens, this is generally your public fully qualified domain.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("jwtConfigurations")]
        private InputList<Inputs.FusionAuthTenantJwtConfigurationGetArgs>? _jwtConfigurations;
        public InputList<Inputs.FusionAuthTenantJwtConfigurationGetArgs> JwtConfigurations
        {
            get => _jwtConfigurations ?? (_jwtConfigurations = new InputList<Inputs.FusionAuthTenantJwtConfigurationGetArgs>());
            set => _jwtConfigurations = value;
        }

        [Input("loginConfiguration")]
        public Input<Inputs.FusionAuthTenantLoginConfigurationGetArgs>? LoginConfiguration { get; set; }

        /// <summary>
        /// The logout redirect URL when sending the user’s browser to the /oauth2/logout URI of the FusionAuth Front End. This value is only used when a logout URL is not defined in your Application.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        [Input("maximumPasswordAge")]
        public Input<Inputs.FusionAuthTenantMaximumPasswordAgeGetArgs>? MaximumPasswordAge { get; set; }

        [Input("minimumPasswordAge")]
        public Input<Inputs.FusionAuthTenantMinimumPasswordAgeGetArgs>? MinimumPasswordAge { get; set; }

        [Input("multiFactorConfiguration")]
        public Input<Inputs.FusionAuthTenantMultiFactorConfigurationGetArgs>? MultiFactorConfiguration { get; set; }

        /// <summary>
        /// The unique name of the Tenant.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("oauthConfigurations")]
        private InputList<Inputs.FusionAuthTenantOauthConfigurationGetArgs>? _oauthConfigurations;
        public InputList<Inputs.FusionAuthTenantOauthConfigurationGetArgs> OauthConfigurations
        {
            get => _oauthConfigurations ?? (_oauthConfigurations = new InputList<Inputs.FusionAuthTenantOauthConfigurationGetArgs>());
            set => _oauthConfigurations = value;
        }

        [Input("passwordEncryptionConfigurations")]
        private InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationGetArgs>? _passwordEncryptionConfigurations;
        public InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationGetArgs> PasswordEncryptionConfigurations
        {
            get => _passwordEncryptionConfigurations ?? (_passwordEncryptionConfigurations = new InputList<Inputs.FusionAuthTenantPasswordEncryptionConfigurationGetArgs>());
            set => _passwordEncryptionConfigurations = value;
        }

        [Input("passwordValidationRules")]
        public Input<Inputs.FusionAuthTenantPasswordValidationRulesGetArgs>? PasswordValidationRules { get; set; }

        /// <summary>
        /// The optional Id of an existing Tenant to make a copy of. If present, the tenant.id and tenant.name values of the request body will be applied to the new Tenant, all other values will be copied from the source Tenant to the new Tenant.
        /// </summary>
        [Input("sourceTenantId")]
        public Input<string>? SourceTenantId { get; set; }

        /// <summary>
        /// The Id to use for the new Tenant. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The unique Id of the theme to be used to style the login page and other end user templates.
        /// </summary>
        [Input("themeId")]
        public Input<string>? ThemeId { get; set; }

        [Input("userDeletePolicy")]
        public Input<Inputs.FusionAuthTenantUserDeletePolicyGetArgs>? UserDeletePolicy { get; set; }

        [Input("usernameConfiguration")]
        public Input<Inputs.FusionAuthTenantUsernameConfigurationGetArgs>? UsernameConfiguration { get; set; }

        public FusionAuthTenantState()
        {
        }
        public static new FusionAuthTenantState Empty => new FusionAuthTenantState();
    }
}
