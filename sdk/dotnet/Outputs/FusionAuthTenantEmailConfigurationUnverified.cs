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
    public sealed class FusionAuthTenantEmailConfigurationUnverified
    {
        /// <summary>
        /// When this value is set to true, the user is allowed to change their email address when they are gated because they haven’t verified their email address.
        /// </summary>
        public readonly bool? AllowEmailChangeWhenGated;
        /// <summary>
        /// = (Optional) The behavior when detecting breaches at time of user login.
        /// </summary>
        public readonly string? Behavior;

        [OutputConstructor]
        private FusionAuthTenantEmailConfigurationUnverified(
            bool? allowEmailChangeWhenGated,

            string? behavior)
        {
            AllowEmailChangeWhenGated = allowEmailChangeWhenGated;
            Behavior = behavior;
        }
    }
}
