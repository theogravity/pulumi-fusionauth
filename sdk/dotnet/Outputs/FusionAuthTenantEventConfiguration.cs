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
    public sealed class FusionAuthTenantEventConfiguration
    {
        /// <summary>
        /// Whether or not FusionAuth should send these types of events to any configured Webhooks.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The event type
        /// </summary>
        public readonly string? Event;
        /// <summary>
        /// The transaction type that FusionAuth uses when sending these types of events to any configured Webhooks.
        /// </summary>
        public readonly string? TransactionType;

        [OutputConstructor]
        private FusionAuthTenantEventConfiguration(
            bool? enabled,

            string? @event,

            string? transactionType)
        {
            Enabled = enabled;
            Event = @event;
            TransactionType = transactionType;
        }
    }
}
