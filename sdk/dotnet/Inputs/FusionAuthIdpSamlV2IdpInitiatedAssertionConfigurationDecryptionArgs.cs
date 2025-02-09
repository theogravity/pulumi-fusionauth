// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Inputs
{

    public sealed class FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationDecryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if FusionAuth requires encrypted assertions in SAML responses from the identity provider. When true, SAML responses from the identity provider containing unencrypted assertions will be rejected by FusionAuth.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Id of the key stored in Key Master that is used to decrypt the symmetric key on the SAML response sent to FusionAuth from the identity provider. The selected Key must contain an RSA private key. Required when `'enabled` is true.
        /// </summary>
        [Input("keyTransportDecryptionKeyId", required: true)]
        public Input<string> KeyTransportDecryptionKeyId { get; set; } = null!;

        public FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationDecryptionArgs()
        {
        }
        public static new FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationDecryptionArgs Empty => new FusionAuthIdpSamlV2IdpInitiatedAssertionConfigurationDecryptionArgs();
    }
}
