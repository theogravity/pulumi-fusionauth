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

    public sealed class FusionAuthTenantEventConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not FusionAuth should send these types of events to any configured Webhooks.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The event type
        /// </summary>
        [Input("event")]
        public Input<string>? Event { get; set; }

        /// <summary>
        /// The transaction type that FusionAuth uses when sending these types of events to any configured Webhooks.
        /// </summary>
        [Input("transactionType")]
        public Input<string>? TransactionType { get; set; }

        public FusionAuthTenantEventConfigurationGetArgs()
        {
        }
        public static new FusionAuthTenantEventConfigurationGetArgs Empty => new FusionAuthTenantEventConfigurationGetArgs();
    }
}
