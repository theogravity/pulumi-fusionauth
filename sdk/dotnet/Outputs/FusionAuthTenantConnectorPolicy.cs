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
    public sealed class FusionAuthTenantConnectorPolicy
    {
        /// <summary>
        /// The identifier of the Connector to which this policy refers.
        /// </summary>
        public readonly string? ConnectorId;
        /// <summary>
        /// A list of email domains to which this connector should apply. A value of ["*"] indicates this connector applies to all users.
        /// </summary>
        public readonly ImmutableArray<string> Domains;
        /// <summary>
        /// If true, the user’s data will be migrated to FusionAuth at first successful authentication; subsequent authentications will occur against the FusionAuth datastore. If false, the Connector’s source will be treated as authoritative.
        /// </summary>
        public readonly bool? Migrate;

        [OutputConstructor]
        private FusionAuthTenantConnectorPolicy(
            string? connectorId,

            ImmutableArray<string> domains,

            bool? migrate)
        {
            ConnectorId = connectorId;
            Domains = domains;
            Migrate = migrate;
        }
    }
}
