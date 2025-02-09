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
    public sealed class FusionAuthWebhookSignatureConfiguration
    {
        /// <summary>
        /// Wether or not webhook signing is enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The UUID key used for signing the Webhook
        /// </summary>
        public readonly string? SigningKeyId;

        [OutputConstructor]
        private FusionAuthWebhookSignatureConfiguration(
            bool? enabled,

            string? signingKeyId)
        {
            Enabled = enabled;
            SigningKeyId = signingKeyId;
        }
    }
}
