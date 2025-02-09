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

    public sealed class FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The length of the secure generator used for generating the change password Id. Defaults to 32.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// The type of the secure generator used for generating the change password Id. Defaults to randomBytes.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs()
        {
        }
        public static new FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs Empty => new FusionAuthTenantExternalIdentifierConfigurationSetupPasswordIdGeneratorArgs();
    }
}
