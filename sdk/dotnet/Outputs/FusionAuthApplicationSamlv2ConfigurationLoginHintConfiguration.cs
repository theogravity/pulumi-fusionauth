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
    public sealed class FusionAuthApplicationSamlv2ConfigurationLoginHintConfiguration
    {
        /// <summary>
        /// When enabled, FusionAuth will accept a username or email address as a login hint on a custom HTTP request parameter.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The name of the parameter that will be used to pass the login hint to the SAML v2 IdP.
        /// </summary>
        public readonly string? ParameterName;

        [OutputConstructor]
        private FusionAuthApplicationSamlv2ConfigurationLoginHintConfiguration(
            bool? enabled,

            string? parameterName)
        {
            Enabled = enabled;
            ParameterName = parameterName;
        }
    }
}
