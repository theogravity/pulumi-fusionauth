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

    public sealed class FusionAuthIdpTwitchApplicationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application to apply this configuration to.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level button text.
        /// </summary>
        [Input("buttonText")]
        public Input<string>? ButtonText { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level client_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level client_secret.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Determines if a UserRegistration is created for the User automatically or not. If a user doesn’t exist in FusionAuth and logs in through an identity provider, this boolean controls whether or not FusionAuth creates a registration for the User in the Application they are logging into.
        /// </summary>
        [Input("createRegistration")]
        public Input<bool>? CreateRegistration { get; set; }

        /// <summary>
        /// Determines if this identity provider is enabled for the Application specified by the applicationId key.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level scope.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public FusionAuthIdpTwitchApplicationConfigurationArgs()
        {
        }
        public static new FusionAuthIdpTwitchApplicationConfigurationArgs Empty => new FusionAuthIdpTwitchApplicationConfigurationArgs();
    }
}
