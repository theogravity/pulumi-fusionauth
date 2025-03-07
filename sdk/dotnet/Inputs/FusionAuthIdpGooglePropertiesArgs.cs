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

    public sealed class FusionAuthIdpGooglePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google Identity Services login API configuration in a properties file formatted String. Any attribute from Google's documentation can be added. Properties can be referenced in templates that support Google login to initialize the API via HTML or JavaScript. The properties specified in this field should not include the data- prefix on the property name. If the `login_method` is set to UsePopup and this value contains the conflicting ux_mode=redirect property, that single property will be replaced with ux_mode=popup.
        /// </summary>
        [Input("api")]
        public Input<string>? Api { get; set; }

        /// <summary>
        /// Google Identity Services button configuration in a properties file formatted String. Any attribute from Google's documentation can be added. Properties can be referenced in templates that support Google login to render the login button via HTML or JavaScript. The properties specified in this field should not include the data- prefix on the property name.
        /// </summary>
        [Input("button")]
        public Input<string>? Button { get; set; }

        public FusionAuthIdpGooglePropertiesArgs()
        {
        }
        public static new FusionAuthIdpGooglePropertiesArgs Empty => new FusionAuthIdpGooglePropertiesArgs();
    }
}
