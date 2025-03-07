// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Outputs
{

    [OutputType]
    public sealed class FusionAuthTenantExternalIdentifierConfiguration
    {
        /// <summary>
        /// The time in seconds until a OAuth authorization code in no longer valid to be exchanged for an access token. This is essentially the time allowed between the start of an Authorization request during the Authorization code grant and when you request an access token using this authorization code on the Token endpoint. Defaults to 30.
        /// </summary>
        public readonly int? AuthorizationGrantIdTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationChangePasswordIdGenerator? ChangePasswordIdGenerator;
        /// <summary>
        /// The time in seconds until a change password Id is no longer valid and cannot be used by the Change Password API. Value must be greater than 0. Defaults to 600.
        /// </summary>
        public readonly int? ChangePasswordIdTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a device code Id is no longer valid and cannot be used by the Token API. Value must be greater than 0. Defaults to 300.
        /// </summary>
        public readonly int? DeviceCodeTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGenerator? DeviceUserCodeIdGenerator;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationIdGenerator? EmailVerificationIdGenerator;
        /// <summary>
        /// The time in seconds until a email verification Id is no longer valid and cannot be used by the Verify Email API. Value must be greater than 0.
        /// </summary>
        public readonly int? EmailVerificationIdTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationOneTimeCodeGenerator? EmailVerificationOneTimeCodeGenerator;
        /// <summary>
        /// The time in seconds until an external authentication Id is no longer valid and cannot be used by the Token API. Value must be greater than 0. Defaults to 300.
        /// </summary>
        public readonly int? ExternalAuthenticationIdTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a Login Timeout identifier is no longer valid to complete post-authentication steps in the OAuth workflow. Must be greater than 0. Defaults to 1800.
        /// </summary>
        public readonly int? LoginIntentTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a One Time Password is no longer valid and cannot be used by the Login API. Value must be greater than 0. Defaults to 60.
        /// </summary>
        public readonly int? OneTimePasswordTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationPasswordlessLoginGenerator? PasswordlessLoginGenerator;
        /// <summary>
        /// The time in seconds until a passwordless code is no longer valid and cannot be used by the Passwordless API. Value must be greater than 0. Defaults to 180.
        /// </summary>
        public readonly int? PasswordlessLoginTimeToLiveInSeconds;
        /// <summary>
        /// The number of seconds before the pending account link identifier is no longer valid to complete an account link request. Value must be greater than 0. Defaults to 3600
        /// </summary>
        public readonly int? PendingAccountLinkTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationIdGenerator? RegistrationVerificationIdGenerator;
        /// <summary>
        /// The time in seconds until a registration verification Id is no longer valid and cannot be used by the Verify Registration API. Value must be greater than 0.
        /// </summary>
        public readonly int? RegistrationVerificationIdTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationOneTimeCodeGenerator? RegistrationVerificationOneTimeCodeGenerator;
        /// <summary>
        /// The time in seconds until remembered OAuth scope consent choices are no longer valid, and the User will be prompted to consent to requested OAuth scopes even if they have not changed. Applies only when `application.oauthConfiguration.consentMode` is set to RememberDecision. Value must be greater than 0. Note: An Essentials or Enterprise plan is required to utilize advanced OAuth scopes. Defaults to 2592000.
        /// </summary>
        public readonly int? RememberOauthScopeConsentChoiceTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds that a SAML AuthN request will be eligible for use to authenticate with FusionAuth. Defaults to 300.
        /// </summary>
        public readonly int? SamlV2AuthnRequestIdTtlSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGenerator? SetupPasswordIdGenerator;
        /// <summary>
        /// The time in seconds until a setup password Id is no longer valid and cannot be used by the Change Password API. Value must be greater than 0.
        /// </summary>
        public readonly int? SetupPasswordIdTimeToLiveInSeconds;
        /// <summary>
        /// The number of seconds before the Trust Token is no longer valid to complete a request that requires trust. Value must be greater than 0. Defaults to 180
        /// </summary>
        public readonly int? TrustTokenTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a two factor Id is no longer valid and cannot be used by the Two Factor Login API. Value must be greater than 0. Defaults to 300.
        /// </summary>
        public readonly int? TwoFactorIdTimeToLiveInSeconds;
        public readonly Outputs.FusionAuthTenantExternalIdentifierConfigurationTwoFactorOneTimeCodeIdGenerator? TwoFactorOneTimeCodeIdGenerator;
        /// <summary>
        /// The number of seconds before the Two-Factor One Time Code used to enable or disable a two-factor method is no longer valid. Must be greater than 0. Defaults to 60.
        /// </summary>
        public readonly int? TwoFactorOneTimeCodeIdTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until an issued Two Factor trust Id is no longer valid and the User will be Optional to complete Two Factor authentication during the next authentication attempt. Value must be greater than 0.
        /// </summary>
        public readonly int? TwoFactorTrustIdTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a WebAuthn authentication challenge is no longer valid and the User will be required to restart the WebAuthn authentication ceremony by creating a new challenge. This value also controls the timeout for the client-side WebAuthn navigator.credentials.get API call. Value must be greater than 0. Note: A license is required to utilize WebAuthn. Defaults to 180.
        /// </summary>
        public readonly int? WebauthnAuthenticationChallengeTimeToLiveInSeconds;
        /// <summary>
        /// The time in seconds until a WebAuthn registration challenge is no longer valid and the User will be required to restart the WebAuthn registration ceremony by creating a new challenge. This value also controls the timeout for the client-side WebAuthn navigator.credentials.create API call. Value must be greater than 0. Note: A license is required to utilize WebAuthn. Defaults to 180.
        /// </summary>
        public readonly int? WebauthnRegistrationChallengeTimeToLiveInSeconds;

        [OutputConstructor]
        private FusionAuthTenantExternalIdentifierConfiguration(
            int? authorizationGrantIdTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationChangePasswordIdGenerator? changePasswordIdGenerator,

            int? changePasswordIdTimeToLiveInSeconds,

            int? deviceCodeTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationDeviceUserCodeIdGenerator? deviceUserCodeIdGenerator,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationIdGenerator? emailVerificationIdGenerator,

            int? emailVerificationIdTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationEmailVerificationOneTimeCodeGenerator? emailVerificationOneTimeCodeGenerator,

            int? externalAuthenticationIdTimeToLiveInSeconds,

            int? loginIntentTimeToLiveInSeconds,

            int? oneTimePasswordTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationPasswordlessLoginGenerator? passwordlessLoginGenerator,

            int? passwordlessLoginTimeToLiveInSeconds,

            int? pendingAccountLinkTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationIdGenerator? registrationVerificationIdGenerator,

            int? registrationVerificationIdTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationRegistrationVerificationOneTimeCodeGenerator? registrationVerificationOneTimeCodeGenerator,

            int? rememberOauthScopeConsentChoiceTimeToLiveInSeconds,

            int? samlV2AuthnRequestIdTtlSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGenerator? setupPasswordIdGenerator,

            int? setupPasswordIdTimeToLiveInSeconds,

            int? trustTokenTimeToLiveInSeconds,

            int? twoFactorIdTimeToLiveInSeconds,

            Outputs.FusionAuthTenantExternalIdentifierConfigurationTwoFactorOneTimeCodeIdGenerator? twoFactorOneTimeCodeIdGenerator,

            int? twoFactorOneTimeCodeIdTimeToLiveInSeconds,

            int? twoFactorTrustIdTimeToLiveInSeconds,

            int? webauthnAuthenticationChallengeTimeToLiveInSeconds,

            int? webauthnRegistrationChallengeTimeToLiveInSeconds)
        {
            AuthorizationGrantIdTimeToLiveInSeconds = authorizationGrantIdTimeToLiveInSeconds;
            ChangePasswordIdGenerator = changePasswordIdGenerator;
            ChangePasswordIdTimeToLiveInSeconds = changePasswordIdTimeToLiveInSeconds;
            DeviceCodeTimeToLiveInSeconds = deviceCodeTimeToLiveInSeconds;
            DeviceUserCodeIdGenerator = deviceUserCodeIdGenerator;
            EmailVerificationIdGenerator = emailVerificationIdGenerator;
            EmailVerificationIdTimeToLiveInSeconds = emailVerificationIdTimeToLiveInSeconds;
            EmailVerificationOneTimeCodeGenerator = emailVerificationOneTimeCodeGenerator;
            ExternalAuthenticationIdTimeToLiveInSeconds = externalAuthenticationIdTimeToLiveInSeconds;
            LoginIntentTimeToLiveInSeconds = loginIntentTimeToLiveInSeconds;
            OneTimePasswordTimeToLiveInSeconds = oneTimePasswordTimeToLiveInSeconds;
            PasswordlessLoginGenerator = passwordlessLoginGenerator;
            PasswordlessLoginTimeToLiveInSeconds = passwordlessLoginTimeToLiveInSeconds;
            PendingAccountLinkTimeToLiveInSeconds = pendingAccountLinkTimeToLiveInSeconds;
            RegistrationVerificationIdGenerator = registrationVerificationIdGenerator;
            RegistrationVerificationIdTimeToLiveInSeconds = registrationVerificationIdTimeToLiveInSeconds;
            RegistrationVerificationOneTimeCodeGenerator = registrationVerificationOneTimeCodeGenerator;
            RememberOauthScopeConsentChoiceTimeToLiveInSeconds = rememberOauthScopeConsentChoiceTimeToLiveInSeconds;
            SamlV2AuthnRequestIdTtlSeconds = samlV2AuthnRequestIdTtlSeconds;
            SetupPasswordIdGenerator = setupPasswordIdGenerator;
            SetupPasswordIdTimeToLiveInSeconds = setupPasswordIdTimeToLiveInSeconds;
            TrustTokenTimeToLiveInSeconds = trustTokenTimeToLiveInSeconds;
            TwoFactorIdTimeToLiveInSeconds = twoFactorIdTimeToLiveInSeconds;
            TwoFactorOneTimeCodeIdGenerator = twoFactorOneTimeCodeIdGenerator;
            TwoFactorOneTimeCodeIdTimeToLiveInSeconds = twoFactorOneTimeCodeIdTimeToLiveInSeconds;
            TwoFactorTrustIdTimeToLiveInSeconds = twoFactorTrustIdTimeToLiveInSeconds;
            WebauthnAuthenticationChallengeTimeToLiveInSeconds = webauthnAuthenticationChallengeTimeToLiveInSeconds;
            WebauthnRegistrationChallengeTimeToLiveInSeconds = webauthnRegistrationChallengeTimeToLiveInSeconds;
        }
    }
}
