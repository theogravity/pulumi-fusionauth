// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth
{
    /// <summary>
    /// ## # API Key
    /// 
    /// The FusionAuth APIs are primarily secured using API keys. This API can only be accessed using an API key that is set as a keyManager. In order to retrieve, update or delete an API key, an API key with equal or greater permissions must be used. A "tenant-scoped" API key can retrieve, create, update or delete an API key for the same tenant. This page describes APIs that are used to manage API keys.
    /// 
    /// [API Key](https://fusionauth.io/docs/v1/tech/apis/api-keys/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fusionauth = theogravity.Fusionauth;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Fusionauth.FusionAuthApiKey("example", new()
    ///     {
    ///         Description = "my super secret key",
    ///         Key = "super-secret-key",
    ///         PermissionsEndpoints = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthApiKeyPermissionsEndpointArgs
    ///             {
    ///                 Delete = true,
    ///                 Endpoint = "/api/application",
    ///                 Get = true,
    ///                 Patch = true,
    ///                 Post = true,
    ///                 Put = true,
    ///             },
    ///         },
    ///         TenantId = "94f751c5-4883-4684-a817-6b106778edec",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthApiKey:FusionAuthApiKey")]
    public partial class FusionAuthApiKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the key.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
        /// </summary>
        [Output("expirationInstant")]
        public Output<int?> ExpirationInstant { get; private set; } = null!;

        /// <summary>
        /// The Id of the IP Access Control List limiting access to this API key.
        /// </summary>
        [Output("ipAccessControlListId")]
        public Output<string?> IpAccessControlListId { get; private set; } = null!;

        /// <summary>
        /// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The Id to use for the new Form. If not specified a secure random UUID will be generated.
        /// </summary>
        [Output("keyId")]
        public Output<string?> KeyId { get; private set; } = null!;

        /// <summary>
        /// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissions_endpoints object mean that this is a super key that authorizes this key for all the endpoints.
        /// </summary>
        [Output("permissionsEndpoints")]
        public Output<ImmutableArray<Outputs.FusionAuthApiKeyPermissionsEndpoint>> PermissionsEndpoints { get; private set; } = null!;

        /// <summary>
        /// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthApiKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthApiKey(string name, FusionAuthApiKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthApiKey:FusionAuthApiKey", name, args ?? new FusionAuthApiKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthApiKey(string name, Input<string> id, FusionAuthApiKeyState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthApiKey:FusionAuthApiKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/theogravity/pulumi-fusionauth/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthApiKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthApiKey Get(string name, Input<string> id, FusionAuthApiKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthApiKey(name, id, state, options);
        }
    }

    public sealed class FusionAuthApiKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the key.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
        /// </summary>
        [Input("expirationInstant")]
        public Input<int>? ExpirationInstant { get; set; }

        /// <summary>
        /// The Id of the IP Access Control List limiting access to this API key.
        /// </summary>
        [Input("ipAccessControlListId")]
        public Input<string>? IpAccessControlListId { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Id to use for the new Form. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("permissionsEndpoints")]
        private InputList<Inputs.FusionAuthApiKeyPermissionsEndpointArgs>? _permissionsEndpoints;

        /// <summary>
        /// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissions_endpoints object mean that this is a super key that authorizes this key for all the endpoints.
        /// </summary>
        public InputList<Inputs.FusionAuthApiKeyPermissionsEndpointArgs> PermissionsEndpoints
        {
            get => _permissionsEndpoints ?? (_permissionsEndpoints = new InputList<Inputs.FusionAuthApiKeyPermissionsEndpointArgs>());
            set => _permissionsEndpoints = value;
        }

        /// <summary>
        /// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public FusionAuthApiKeyArgs()
        {
        }
        public static new FusionAuthApiKeyArgs Empty => new FusionAuthApiKeyArgs();
    }

    public sealed class FusionAuthApiKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the key.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
        /// </summary>
        [Input("expirationInstant")]
        public Input<int>? ExpirationInstant { get; set; }

        /// <summary>
        /// The Id of the IP Access Control List limiting access to this API key.
        /// </summary>
        [Input("ipAccessControlListId")]
        public Input<string>? IpAccessControlListId { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Id to use for the new Form. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("permissionsEndpoints")]
        private InputList<Inputs.FusionAuthApiKeyPermissionsEndpointGetArgs>? _permissionsEndpoints;

        /// <summary>
        /// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissions_endpoints object mean that this is a super key that authorizes this key for all the endpoints.
        /// </summary>
        public InputList<Inputs.FusionAuthApiKeyPermissionsEndpointGetArgs> PermissionsEndpoints
        {
            get => _permissionsEndpoints ?? (_permissionsEndpoints = new InputList<Inputs.FusionAuthApiKeyPermissionsEndpointGetArgs>());
            set => _permissionsEndpoints = value;
        }

        /// <summary>
        /// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public FusionAuthApiKeyState()
        {
        }
        public static new FusionAuthApiKeyState Empty => new FusionAuthApiKeyState();
    }
}
