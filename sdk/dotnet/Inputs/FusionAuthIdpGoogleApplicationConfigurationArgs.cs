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

    public sealed class FusionAuthIdpGoogleApplicationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Application to apply this configuration to.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        [Input("buttonText")]
        public Input<string>? ButtonText { get; set; }

        /// <summary>
        /// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines if a UserRegistration is created for the User automatically or not. If a user doesn’t exist in FusionAuth and logs in through an identity provider, this boolean controls whether or not FusionAuth creates a registration for the User in the Application they are logging into.
        /// </summary>
        [Input("createRegistration")]
        public Input<bool>? CreateRegistration { get; set; }

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The top-level scope that you are requesting from Google.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public FusionAuthIdpGoogleApplicationConfigurationArgs()
        {
        }
        public static new FusionAuthIdpGoogleApplicationConfigurationArgs Empty => new FusionAuthIdpGoogleApplicationConfigurationArgs();
    }
}
