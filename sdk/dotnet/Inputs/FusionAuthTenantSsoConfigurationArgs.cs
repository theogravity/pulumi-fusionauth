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

    public sealed class FusionAuthTenantSsoConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds before a trusted device is reset. When reset, a user is forced to complete captcha during login and complete two factor authentication if applicable.
        /// </summary>
        [Input("deviceTrustTimeToLiveInSeconds")]
        public Input<int>? DeviceTrustTimeToLiveInSeconds { get; set; }

        public FusionAuthTenantSsoConfigurationArgs()
        {
        }
        public static new FusionAuthTenantSsoConfigurationArgs Empty => new FusionAuthTenantSsoConfigurationArgs();
    }
}
