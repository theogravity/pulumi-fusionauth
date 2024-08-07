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

    public sealed class FusionAuthTenantMultiFactorConfigurationEmailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, users may utilize an email address to complete a multi-factor authentication request.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Id of the email template that is used when notifying a user to complete a multi-factor authentication request.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public FusionAuthTenantMultiFactorConfigurationEmailArgs()
        {
        }
        public static new FusionAuthTenantMultiFactorConfigurationEmailArgs Empty => new FusionAuthTenantMultiFactorConfigurationEmailArgs();
    }
}
