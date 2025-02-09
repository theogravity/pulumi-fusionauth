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

    public sealed class FusionAuthTenantScimServerConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Entity Type that will be used to represent SCIM Clients for this tenant. Note: An Enterprise plan is required to utilize SCIM. Required when `scim_server_configuration.enabled` is true.
        /// </summary>
        [Input("clientEntityTypeId", required: true)]
        public Input<string> ClientEntityTypeId { get; set; } = null!;

        /// <summary>
        /// Whether or not this tenant has the SCIM endpoints enabled. Note: An Enterprise plan is required to utilize SCIM.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// SON formatted as a SCIM Schemas endpoint response. Because the SCIM lambdas may modify the JSON response, ensure the Schema's response matches that generated by the response lambdas. More about Schema definitions. When this parameter is not provided, it will default to EnterpriseUser, Group, and User schema definitions as defined by the SCIM core schemas spec. Note: An Enterprise plan is required to utilize SCIM.
        /// </summary>
        [Input("schemas")]
        public Input<string>? Schemas { get; set; }

        /// <summary>
        /// The Entity Type that will be used to represent SCIM Servers for this tenant. Note: An Enterprise plan is required to utilize SCIM. Required when `scim_server_configuration.enabled` is true.
        /// </summary>
        [Input("serverEntityTypeId", required: true)]
        public Input<string> ServerEntityTypeId { get; set; } = null!;

        public FusionAuthTenantScimServerConfigurationGetArgs()
        {
        }
        public static new FusionAuthTenantScimServerConfigurationGetArgs Empty => new FusionAuthTenantScimServerConfigurationGetArgs();
    }
}
