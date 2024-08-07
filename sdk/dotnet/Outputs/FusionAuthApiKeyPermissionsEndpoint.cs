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
    public sealed class FusionAuthApiKeyPermissionsEndpoint
    {
        /// <summary>
        /// HTTP DELETE Verb
        /// </summary>
        public readonly bool? Delete;
        public readonly string Endpoint;
        /// <summary>
        /// HTTP GET Verb
        /// </summary>
        public readonly bool? Get;
        /// <summary>
        /// HTTP PATCH Verb
        /// </summary>
        public readonly bool? Patch;
        /// <summary>
        /// HTTP POST Verb
        /// </summary>
        public readonly bool? Post;
        /// <summary>
        /// HTTP PUT Verb
        /// </summary>
        public readonly bool? Put;

        [OutputConstructor]
        private FusionAuthApiKeyPermissionsEndpoint(
            bool? delete,

            string endpoint,

            bool? get,

            bool? patch,

            bool? post,

            bool? put)
        {
            Delete = delete;
            Endpoint = endpoint;
            Get = get;
            Patch = patch;
            Post = post;
            Put = put;
        }
    }
}
