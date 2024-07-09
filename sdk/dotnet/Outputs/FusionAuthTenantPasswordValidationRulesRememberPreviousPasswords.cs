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
    public sealed class FusionAuthTenantPasswordValidationRulesRememberPreviousPasswords
    {
        /// <summary>
        /// The number of previous passwords to remember. Value must be greater than 0.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// Whether to prevent a user from using any of their previous passwords.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private FusionAuthTenantPasswordValidationRulesRememberPreviousPasswords(
            int? count,

            bool? enabled)
        {
            Count = count;
            Enabled = enabled;
        }
    }
}
