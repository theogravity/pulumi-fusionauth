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
    public sealed class FusionAuthApplicationLoginConfiguration
    {
        /// <summary>
        /// Indicates if a JWT may be refreshed using a Refresh Token for this application. This configuration is separate from issuing new Refresh Tokens which is controlled by the generateRefreshTokens parameter. This configuration indicates specifically if an existing Refresh Token may be used to request a new JWT using the Refresh API.
        /// </summary>
        public readonly bool? AllowTokenRefresh;
        /// <summary>
        /// Indicates if a Refresh Token should be issued from the Login API
        /// </summary>
        public readonly bool? GenerateRefreshTokens;
        /// <summary>
        /// Indicates if the Login API should require an API key. If you set this value to false and your FusionAuth API is on a public network, anyone may attempt to use the Login API.
        /// </summary>
        public readonly bool? RequireAuthentication;

        [OutputConstructor]
        private FusionAuthApplicationLoginConfiguration(
            bool? allowTokenRefresh,

            bool? generateRefreshTokens,

            bool? requireAuthentication)
        {
            AllowTokenRefresh = allowTokenRefresh;
            GenerateRefreshTokens = generateRefreshTokens;
            RequireAuthentication = requireAuthentication;
        }
    }
}
