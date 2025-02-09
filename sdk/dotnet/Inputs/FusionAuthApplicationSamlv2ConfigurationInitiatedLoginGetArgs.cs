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

    public sealed class FusionAuthApplicationSamlv2ConfigurationInitiatedLoginGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if SAML v2 IdP initiated login is enabled for this application. See application.samlv2Configuration.authorizedRedirectURLs for information on which destination URLs are allowed.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The value sent in the AuthN response to the SAML v2 Service Provider in the NameID assertion.
        /// </summary>
        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        public FusionAuthApplicationSamlv2ConfigurationInitiatedLoginGetArgs()
        {
        }
        public static new FusionAuthApplicationSamlv2ConfigurationInitiatedLoginGetArgs Empty => new FusionAuthApplicationSamlv2ConfigurationInitiatedLoginGetArgs();
    }
}
