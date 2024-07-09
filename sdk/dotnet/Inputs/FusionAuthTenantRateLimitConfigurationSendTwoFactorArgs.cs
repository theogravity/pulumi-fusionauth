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

    public sealed class FusionAuthTenantRateLimitConfigurationSendTwoFactorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether rate limiting is enabled for send two factor.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The number of times a user can request a two-factor code by email or SMS within the configured `time_period_in_seconds` duration.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// The duration for the number of times a user can request a two-factor code by email or SMS before being rate limited.
        /// </summary>
        [Input("timePeriodInSeconds")]
        public Input<int>? TimePeriodInSeconds { get; set; }

        public FusionAuthTenantRateLimitConfigurationSendTwoFactorArgs()
        {
        }
        public static new FusionAuthTenantRateLimitConfigurationSendTwoFactorArgs Empty => new FusionAuthTenantRateLimitConfigurationSendTwoFactorArgs();
    }
}
