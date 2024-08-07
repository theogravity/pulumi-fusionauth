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

    public sealed class FusionAuthTenantUserDeletePolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates that users without a verified email address will be permanently deleted after tenant.userDeletePolicy.unverified.numberOfDaysToRetain days.
        /// </summary>
        [Input("unverifiedEnabled")]
        public Input<bool>? UnverifiedEnabled { get; set; }

        /// <summary>
        /// The number of days from creation users will be retained before being deleted for not completing email verification. This field is required when tenant.userDeletePolicy.unverified.enabled is set to true. Value must be greater than 0.
        /// </summary>
        [Input("unverifiedNumberOfDaysToRetain")]
        public Input<int>? UnverifiedNumberOfDaysToRetain { get; set; }

        public FusionAuthTenantUserDeletePolicyGetArgs()
        {
        }
        public static new FusionAuthTenantUserDeletePolicyGetArgs Empty => new FusionAuthTenantUserDeletePolicyGetArgs();
    }
}
