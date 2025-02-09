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
    public sealed class FusionAuthIdpSamlv2LoginHintConfiguration
    {
        /// <summary>
        /// When enabled and HTTP-Redirect bindings are used, FusionAuth will provide the username or email address when available to the IdP as a login hint using the configured parameter name set by the `parameter_name` to initiate the AuthN request.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The name of the parameter used to pass the username or email as login hint to the IDP when enabled, and HTTP redirect bindings are used to initiate the AuthN request. The default value is `login_hint`. Required when `enabled` is true.
        /// </summary>
        public readonly string? ParameterName;

        [OutputConstructor]
        private FusionAuthIdpSamlv2LoginHintConfiguration(
            bool? enabled,

            string? parameterName)
        {
            Enabled = enabled;
            ParameterName = parameterName;
        }
    }
}
