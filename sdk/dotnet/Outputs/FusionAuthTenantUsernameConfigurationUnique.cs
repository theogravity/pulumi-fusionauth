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
    public sealed class FusionAuthTenantUsernameConfigurationUnique
    {
        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The maximum number of digits to use when building a unique suffix for a username. A number will be randomly selected and will be 1 or more digits up to this configured value in length. For example, if this value is 5, the suffix will be a number between 00001 and 99999, inclusive.
        /// </summary>
        public readonly int? NumberOfDigits;
        /// <summary>
        /// A single character to use as a separator from the requested username and a unique suffix that is added when a duplicate username is detected. This value can be a single non-alphanumeric ASCII character.
        /// </summary>
        public readonly string? Separator;
        /// <summary>
        /// When enabled the user’s password will be validated during login. If the password does not meet the currently configured validation rules the user will be required to change their password.
        /// </summary>
        public readonly string? Strategy;

        [OutputConstructor]
        private FusionAuthTenantUsernameConfigurationUnique(
            bool? enabled,

            int? numberOfDigits,

            string? separator,

            string? strategy)
        {
            Enabled = enabled;
            NumberOfDigits = numberOfDigits;
            Separator = separator;
            Strategy = strategy;
        }
    }
}
