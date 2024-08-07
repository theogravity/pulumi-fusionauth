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

    public sealed class FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of previous passwords to remember. Value must be greater than 0.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// Whether to prevent a user from using any of their previous passwords.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsGetArgs()
        {
        }
        public static new FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsGetArgs Empty => new FusionAuthTenantPasswordValidationRulesRememberPreviousPasswordsGetArgs();
    }
}
