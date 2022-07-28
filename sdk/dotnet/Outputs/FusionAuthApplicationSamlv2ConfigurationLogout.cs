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
    public sealed class FusionAuthApplicationSamlv2ConfigurationLogout
    {
        /// <summary>
        /// This configuration is functionally equivalent to the Logout Behavior found in the OAuth2 configuration.
        /// </summary>
        public readonly string? Behavior;
        /// <summary>
        /// The unique Id of the Key used to verify the signature if the public key cannot be determined by the KeyInfo element when using POST bindings, or the key used to verify the signature when using HTTP Redirect bindings.
        /// </summary>
        public readonly string? DefaultVerificationKeyId;
        /// <summary>
        /// The unique Id of the Key used to sign the SAML Single Logout response.
        /// </summary>
        public readonly string? KeyId;
        /// <summary>
        /// Set this parameter equal to true to require the SAML v2 Service Provider to sign the Logout request. When this value is true all Logout requests missing a signature will be rejected.
        /// </summary>
        public readonly bool? RequireSignedRequests;
        public readonly Outputs.FusionAuthApplicationSamlv2ConfigurationLogoutSingleLogout? SingleLogout;
        /// <summary>
        /// The XML signature canonicalization method used when digesting and signing the SAML response. Unfortunately, many service providers do not correctly implement the XML signature specifications and force a specific canonicalization method. This setting allows you to change the canonicalization method to match the service provider. Often, service providers don’t even document their required method. You might need to contact enterprise support at the service provider to figure out what method they use.
        /// </summary>
        public readonly string? XmlSignatureCanonicalizationMethod;

        [OutputConstructor]
        private FusionAuthApplicationSamlv2ConfigurationLogout(
            string? behavior,

            string? defaultVerificationKeyId,

            string? keyId,

            bool? requireSignedRequests,

            Outputs.FusionAuthApplicationSamlv2ConfigurationLogoutSingleLogout? singleLogout,

            string? xmlSignatureCanonicalizationMethod)
        {
            Behavior = behavior;
            DefaultVerificationKeyId = defaultVerificationKeyId;
            KeyId = keyId;
            RequireSignedRequests = requireSignedRequests;
            SingleLogout = singleLogout;
            XmlSignatureCanonicalizationMethod = xmlSignatureCanonicalizationMethod;
        }
    }
}