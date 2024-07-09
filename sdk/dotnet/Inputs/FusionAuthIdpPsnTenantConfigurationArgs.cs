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

    public sealed class FusionAuthIdpPsnTenantConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, the number of identity provider links a user may create is enforced by maximumLinks.
        /// </summary>
        [Input("limitUserLinkCountEnabled")]
        public Input<bool>? LimitUserLinkCountEnabled { get; set; }

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Input("limitUserLinkCountMaximumLinks")]
        public Input<int>? LimitUserLinkCountMaximumLinks { get; set; }

        /// <summary>
        /// The unique Id of the tenant that this configuration applies to.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public FusionAuthIdpPsnTenantConfigurationArgs()
        {
        }
        public static new FusionAuthIdpPsnTenantConfigurationArgs Empty => new FusionAuthIdpPsnTenantConfigurationArgs();
    }
}
