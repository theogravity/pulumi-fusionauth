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
    public sealed class FusionAuthTenantOauthConfiguration
    {
        /// <summary>
        /// The Id of a lambda that will be called to populate the JWT during a client credentials grant. **Note:** A paid edition of FusionAuth is required to utilize client credentials grant.
        /// </summary>
        public readonly string? ClientCredentialsAccessTokenPopulateLambdaId;

        [OutputConstructor]
        private FusionAuthTenantOauthConfiguration(string? clientCredentialsAccessTokenPopulateLambdaId)
        {
            ClientCredentialsAccessTokenPopulateLambdaId = clientCredentialsAccessTokenPopulateLambdaId;
        }
    }
}
