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
    public sealed class FusionAuthApplicationOauthConfiguration
    {
        /// <summary>
        /// An array of URLs that are the authorized origins for FusionAuth OAuth.
        /// </summary>
        public readonly ImmutableArray<string> AuthorizedOriginUrls;
        /// <summary>
        /// An array of URLs that are the authorized redirect URLs for FusionAuth OAuth.
        /// </summary>
        public readonly ImmutableArray<string> AuthorizedRedirectUrls;
        /// <summary>
        /// Determines whether wildcard expressions will be allowed in the authorized_redirect_urls and authorized_origin_urls.
        /// </summary>
        public readonly string? AuthorizedUrlValidationPolicy;
        /// <summary>
        /// Determines the client authentication requirements for the OAuth 2.0 Token endpoint.
        /// </summary>
        public readonly string? ClientAuthenticationPolicy;
        /// <summary>
        /// The OAuth 2.0 client id. If you leave this blank during a POST, a client id will be generated for you. If you leave this blank during PUT, the previous value will be maintained. For both POST and PUT you can provide a value and it will be stored.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The OAuth 2.0 client secret. If you leave this blank during a POST, a secure secret will be generated for you. If you leave this blank during PUT, the previous value will be maintained. For both POST and PUT you can provide a value and it will be stored.
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// Controls the policy for prompting a user to consent to requested OAuth scopes. This configuration only takes effect when `application.oauthConfiguration.relationship` is `ThirdParty`. The possible values are:
        /// * `AlwaysPrompt` - Always prompt the user for consent.
        /// * `RememberDecision` - Remember previous consents; only prompt if the choice expires or if the requested or required scopes have changed. The duration of this persisted choice is controlled by the Tenant’s `externalIdentifierConfiguration.rememberOAuthScopeConsentChoiceTimeToLiveInSeconds` value.
        /// * `NeverPrompt` - The user will be never be prompted to consent to requested OAuth scopes. Permission will be granted implicitly as if this were a `FirstParty` application. This configuration is meant for testing purposes only and should not be used in production.
        /// </summary>
        public readonly string? ConsentMode;
        /// <summary>
        /// Whether or not FusionAuth will log a debug Event Log. This is particular useful for debugging the authorization code exchange with the Token endpoint during an Authorization Code grant."
        /// </summary>
        public readonly bool? Debug;
        /// <summary>
        /// The device verification URL to be used with the Device Code grant type, this field is required when device_code is enabled.
        /// </summary>
        public readonly string? DeviceVerificationUrl;
        /// <summary>
        /// The enabled grants for this application. In order to utilize a particular grant with the OAuth 2.0 endpoints you must have enabled the grant.
        /// </summary>
        public readonly ImmutableArray<string> EnabledGrants;
        /// <summary>
        /// Determines if the OAuth 2.0 Token endpoint will generate a refresh token when the offline_access scope is requested.
        /// </summary>
        public readonly bool? GenerateRefreshTokens;
        /// <summary>
        /// Behavior when /oauth2/logout is called.
        /// </summary>
        public readonly string? LogoutBehavior;
        /// <summary>
        /// The logout URL for the Application. FusionAuth will redirect to this URL after the user logs out of OAuth.
        /// </summary>
        public readonly string? LogoutUrl;
        /// <summary>
        /// Determines the PKCE requirements when using the authorization code grant.
        /// </summary>
        public readonly string? ProofKeyForCodeExchangePolicy;
        /// <summary>
        /// Configures which of the default scopes are enabled and required.
        /// </summary>
        public readonly ImmutableArray<Outputs.FusionAuthApplicationOauthConfigurationProvidedScopePolicy> ProvidedScopePolicies;
        /// <summary>
        /// The application’s relationship to the OAuth server. The possible values are:
        /// * `FirstParty` - The application has the same owner as the authorization server. Consent to requested OAuth scopes is granted implicitly.
        /// * `ThirdParty` - The application is external to the authorization server. Users will be prompted to consent to requested OAuth scopes based on the application object’s `oauthConfiguration.consentMode` value. Note: An Essentials or Enterprise plan is required to utilize third-party applications.
        /// </summary>
        public readonly string? Relationship;
        /// <summary>
        /// Determines if the OAuth 2.0 Token endpoint requires client authentication. If this is enabled, the client must provide client credentials when using the Token endpoint. The client_id and client_secret may be provided using a Basic Authorization HTTP header, or by sending these parameters in the request body using POST data.
        /// </summary>
        public readonly bool? RequireClientAuthentication;
        /// <summary>
        /// When enabled the user will be required to be registered, or complete registration before redirecting to the configured callback in the authorization code grant or the implicit grant. This configuration does not currently apply to any other grant.
        /// </summary>
        public readonly bool? RequireRegistration;
        /// <summary>
        /// Controls the policy for handling of OAuth scopes when populating JWTs and the UserInfo response. The possible values are:
        /// * `Compatibility` - OAuth workflows will populate JWT and UserInfo claims in a manner compatible with versions of FusionAuth before version 1.50.0.
        /// * `Strict` - OAuth workflows will populate token and UserInfo claims according to the OpenID Connect 1.0 specification based on requested and consented scopes.
        /// </summary>
        public readonly string ScopeHandlingPolicy;
        /// <summary>
        /// Controls the policy for handling unknown scopes on an OAuth request. The possible values are:
        /// * `Allow` - Unknown scopes will be allowed on the request, passed through the OAuth workflow, and written to the resulting tokens without consent.
        /// * `Remove` - Unknown scopes will be removed from the OAuth workflow, but the workflow will proceed without them.
        /// * `Reject` - Unknown scopes will be rejected and cause the OAuth workflow to fail with an error.
        /// </summary>
        public readonly string UnknownScopePolicy;

        [OutputConstructor]
        private FusionAuthApplicationOauthConfiguration(
            ImmutableArray<string> authorizedOriginUrls,

            ImmutableArray<string> authorizedRedirectUrls,

            string? authorizedUrlValidationPolicy,

            string? clientAuthenticationPolicy,

            string? clientId,

            string? clientSecret,

            string? consentMode,

            bool? debug,

            string? deviceVerificationUrl,

            ImmutableArray<string> enabledGrants,

            bool? generateRefreshTokens,

            string? logoutBehavior,

            string? logoutUrl,

            string? proofKeyForCodeExchangePolicy,

            ImmutableArray<Outputs.FusionAuthApplicationOauthConfigurationProvidedScopePolicy> providedScopePolicies,

            string? relationship,

            bool? requireClientAuthentication,

            bool? requireRegistration,

            string scopeHandlingPolicy,

            string unknownScopePolicy)
        {
            AuthorizedOriginUrls = authorizedOriginUrls;
            AuthorizedRedirectUrls = authorizedRedirectUrls;
            AuthorizedUrlValidationPolicy = authorizedUrlValidationPolicy;
            ClientAuthenticationPolicy = clientAuthenticationPolicy;
            ClientId = clientId;
            ClientSecret = clientSecret;
            ConsentMode = consentMode;
            Debug = debug;
            DeviceVerificationUrl = deviceVerificationUrl;
            EnabledGrants = enabledGrants;
            GenerateRefreshTokens = generateRefreshTokens;
            LogoutBehavior = logoutBehavior;
            LogoutUrl = logoutUrl;
            ProofKeyForCodeExchangePolicy = proofKeyForCodeExchangePolicy;
            ProvidedScopePolicies = providedScopePolicies;
            Relationship = relationship;
            RequireClientAuthentication = requireClientAuthentication;
            RequireRegistration = requireRegistration;
            ScopeHandlingPolicy = scopeHandlingPolicy;
            UnknownScopePolicy = unknownScopePolicy;
        }
    }
}
