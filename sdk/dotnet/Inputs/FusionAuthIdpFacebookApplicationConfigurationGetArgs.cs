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

    public sealed class FusionAuthIdpFacebookApplicationConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is an optional Application specific override for the top level `app_id`.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// ID of the FusionAuth Application to apply this configuration to.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level `button_text`.
        /// </summary>
        [Input("buttonText")]
        public Input<string>? ButtonText { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// This is an optional Application specific override for the top level `client_secret`.
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
        /// Determines if a `UserRegistration` is created for the User automatically or not. If a user doesn’t exist in FusionAuth and logs in through an identity provider, this boolean controls whether or not FusionAuth creates a registration for the User in the Application they are logging into.
        /// </summary>
        [Input("createRegistration")]
        public Input<bool>? CreateRegistration { get; set; }

        /// <summary>
        /// Determines if this identity provider is enabled for the Application specified by the `application_id` property.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level `fields`.
        /// </summary>
        [Input("fields")]
        public Input<string>? Fields { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level `permissions`.
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        public FusionAuthIdpFacebookApplicationConfigurationGetArgs()
        {
        }
        public static new FusionAuthIdpFacebookApplicationConfigurationGetArgs Empty => new FusionAuthIdpFacebookApplicationConfigurationGetArgs();
    }
}
