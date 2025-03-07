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

    public sealed class FusionAuthIdpGoogleApplicationConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is an optional Application specific override for the top level properties.api . If this `login_method` is set to UsePopup, or the Application configuration is unset and the top level loginMethod is set to UsePopup, and this value contains the conflicting ux_mode=redirect property, that single property will be replaced with ux_mode=popup.
        /// </summary>
        [Input("api")]
        public Input<string>? Api { get; set; }

        /// <summary>
        /// This is an optional Application specific override for the top level `button`.
        /// </summary>
        [Input("button")]
        public Input<string>? Button { get; set; }

        public FusionAuthIdpGoogleApplicationConfigurationPropertiesArgs()
        {
        }
        public static new FusionAuthIdpGoogleApplicationConfigurationPropertiesArgs Empty => new FusionAuthIdpGoogleApplicationConfigurationPropertiesArgs();
    }
}
