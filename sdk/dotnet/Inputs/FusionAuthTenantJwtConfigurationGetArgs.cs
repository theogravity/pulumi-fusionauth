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

    public sealed class FusionAuthTenantJwtConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique id of the signing key used to sign the access token. Required prior to `1.30.0`.
        /// </summary>
        [Input("accessTokenKeyId")]
        public Input<string>? AccessTokenKeyId { get; set; }

        /// <summary>
        /// The unique id of the signing key used to sign the Id token. Required prior to `1.30.0`.
        /// </summary>
        [Input("idTokenKeyId")]
        public Input<string>? IdTokenKeyId { get; set; }

        /// <summary>
        /// The refresh token expiration policy.
        /// </summary>
        [Input("refreshTokenExpirationPolicy")]
        public Input<string>? RefreshTokenExpirationPolicy { get; set; }

        /// <summary>
        /// When enabled, the refresh token will be revoked when a user action, such as locking an account based on a number of failed login attempts, prevents user login.
        /// </summary>
        [Input("refreshTokenRevocationPolicyOnLoginPrevented")]
        public Input<bool>? RefreshTokenRevocationPolicyOnLoginPrevented { get; set; }

        /// <summary>
        /// When enabled, the refresh token will be revoked when a user changes their password."
        /// </summary>
        [Input("refreshTokenRevocationPolicyOnPasswordChange")]
        public Input<bool>? RefreshTokenRevocationPolicyOnPasswordChange { get; set; }

        /// <summary>
        /// The maximum lifetime of a refresh token when using a refresh token expiration policy of SlidingWindowWithMaximumLifetime. Value must be greater than 0.
        /// </summary>
        [Input("refreshTokenSlidingWindowMaximumTimeToLiveInMinutes")]
        public Input<int>? RefreshTokenSlidingWindowMaximumTimeToLiveInMinutes { get; set; }

        /// <summary>
        /// The length of time in minutes a Refresh Token is valid from the time it was issued. Value must be greater than 0.
        /// </summary>
        [Input("refreshTokenTimeToLiveInMinutes", required: true)]
        public Input<int> RefreshTokenTimeToLiveInMinutes { get; set; } = null!;

        /// <summary>
        /// The refresh token usage policy.
        /// </summary>
        [Input("refreshTokenUsagePolicy")]
        public Input<string>? RefreshTokenUsagePolicy { get; set; }

        /// <summary>
        /// The length of time in seconds this JWT is valid from the time it was issued. Value must be greater than 0.
        /// </summary>
        [Input("timeToLiveInSeconds", required: true)]
        public Input<int> TimeToLiveInSeconds { get; set; } = null!;

        public FusionAuthTenantJwtConfigurationGetArgs()
        {
        }
        public static new FusionAuthTenantJwtConfigurationGetArgs Empty => new FusionAuthTenantJwtConfigurationGetArgs();
    }
}
