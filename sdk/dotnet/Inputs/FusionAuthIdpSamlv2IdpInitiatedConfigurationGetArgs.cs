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

    public sealed class FusionAuthIdpSamlv2IdpInitiatedConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if FusionAuth will accept IdP initiated login requests from this SAMLv2 Identity Provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Required when `enabled` is true.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        public FusionAuthIdpSamlv2IdpInitiatedConfigurationGetArgs()
        {
        }
        public static new FusionAuthIdpSamlv2IdpInitiatedConfigurationGetArgs Empty => new FusionAuthIdpSamlv2IdpInitiatedConfigurationGetArgs();
    }
}
