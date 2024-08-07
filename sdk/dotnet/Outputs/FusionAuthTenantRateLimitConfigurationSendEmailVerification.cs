// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Outputs
{

    [OutputType]
    public sealed class FusionAuthTenantRateLimitConfigurationSendEmailVerification
    {
        /// <summary>
        /// Whether rate limiting is enabled for send email verification.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The number of times a user can request a verification email within the configured `time_period_in_seconds` duration.
        /// </summary>
        public readonly int? Limit;
        /// <summary>
        /// The duration for the number of times a user can request a verification email before being rate limited.
        /// </summary>
        public readonly int? TimePeriodInSeconds;

        [OutputConstructor]
        private FusionAuthTenantRateLimitConfigurationSendEmailVerification(
            bool? enabled,

            int? limit,

            int? timePeriodInSeconds)
        {
            Enabled = enabled;
            Limit = limit;
            TimePeriodInSeconds = timePeriodInSeconds;
        }
    }
}
