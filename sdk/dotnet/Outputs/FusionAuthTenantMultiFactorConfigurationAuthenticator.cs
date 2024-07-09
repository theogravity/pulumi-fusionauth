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
    public sealed class FusionAuthTenantMultiFactorConfigurationAuthenticator
    {
        /// <summary>
        /// When enabled, users may utilize an authenticator application to complete a multi-factor authentication request. This method uses TOTP (Time-Based One-Time Password) as defined in RFC 6238 and often uses an native mobile app such as Google Authenticator.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private FusionAuthTenantMultiFactorConfigurationAuthenticator(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
