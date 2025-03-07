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
    /// ## # Key Resource
    /// 
    /// Cryptographic keys are used in signing and verifying JWTs and verifying responses for third party identity providers. It is more likely you will interact with keys using the FusionAuth UI in the Key Master menu.
    /// 
    /// [Keys API](https://fusionauth.io/docs/v1/tech/apis/keys)
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
    ///     var adminId = new Fusionauth.FusionAuthKey("adminId", new()
    ///     {
    ///         Algorithm = "RS256",
    ///         Length = 2048,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthKey:FusionAuthKey")]
    public partial class FusionAuthKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
        /// * `ES256` - ECDSA using P-256 curve and SHA-256 hash algorithm
        /// * `ES384` - ECDSA using P-384 curve and SHA-384 hash algorithm
        /// * `ES512` - ECDSA using P-521 curve and SHA-512 hash algorithm
        /// * `RS256` - RSA using SHA-256 hash algorithm
        /// * `RS384` - RSA using SHA-384 hash algorithm
        /// * `RS512` - RSA using SHA-512 hash algorithm
        /// * `HS256` - HMAC using SHA-256 hash algorithm
        /// * `HS384` - HMAC using SHA-384 hash algorithm
        /// * `HS512` - HMAC using SHA-512 hash algorithm
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The issuer of the RSA or EC certificate. If omitted, this value will default to the value of tenant issuer on the default tenant. For HMAC keys, this field does not apply and will be ignored if specified, and no default value will be set.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The Id to use for the new key. If not specified a secure random UUID will be generated.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The id used in the JWT header to identify the key used to generate the signature
        /// </summary>
        [Output("kid")]
        public Output<string> Kid { get; private set; } = null!;

        /// <summary>
        /// The length of the RSA or EC certificate. This field is required when generating RSA key types.
        /// </summary>
        [Output("length")]
        public Output<int?> Length { get; private set; } = null!;

        /// <summary>
        /// The name of the Key.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthKey(string name, FusionAuthKeyArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthKey:FusionAuthKey", name, args ?? new FusionAuthKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthKey(string name, Input<string> id, FusionAuthKeyState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthKey:FusionAuthKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/theogravity/pulumi-fusionauth/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthKey Get(string name, Input<string> id, FusionAuthKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthKey(name, id, state, options);
        }
    }

    public sealed class FusionAuthKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
        /// * `ES256` - ECDSA using P-256 curve and SHA-256 hash algorithm
        /// * `ES384` - ECDSA using P-384 curve and SHA-384 hash algorithm
        /// * `ES512` - ECDSA using P-521 curve and SHA-512 hash algorithm
        /// * `RS256` - RSA using SHA-256 hash algorithm
        /// * `RS384` - RSA using SHA-384 hash algorithm
        /// * `RS512` - RSA using SHA-512 hash algorithm
        /// * `HS256` - HMAC using SHA-256 hash algorithm
        /// * `HS384` - HMAC using SHA-384 hash algorithm
        /// * `HS512` - HMAC using SHA-512 hash algorithm
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// The issuer of the RSA or EC certificate. If omitted, this value will default to the value of tenant issuer on the default tenant. For HMAC keys, this field does not apply and will be ignored if specified, and no default value will be set.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The Id to use for the new key. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The length of the RSA or EC certificate. This field is required when generating RSA key types.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// The name of the Key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FusionAuthKeyArgs()
        {
        }
        public static new FusionAuthKeyArgs Empty => new FusionAuthKeyArgs();
    }

    public sealed class FusionAuthKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
        /// * `ES256` - ECDSA using P-256 curve and SHA-256 hash algorithm
        /// * `ES384` - ECDSA using P-384 curve and SHA-384 hash algorithm
        /// * `ES512` - ECDSA using P-521 curve and SHA-512 hash algorithm
        /// * `RS256` - RSA using SHA-256 hash algorithm
        /// * `RS384` - RSA using SHA-384 hash algorithm
        /// * `RS512` - RSA using SHA-512 hash algorithm
        /// * `HS256` - HMAC using SHA-256 hash algorithm
        /// * `HS384` - HMAC using SHA-384 hash algorithm
        /// * `HS512` - HMAC using SHA-512 hash algorithm
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The issuer of the RSA or EC certificate. If omitted, this value will default to the value of tenant issuer on the default tenant. For HMAC keys, this field does not apply and will be ignored if specified, and no default value will be set.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The Id to use for the new key. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The id used in the JWT header to identify the key used to generate the signature
        /// </summary>
        [Input("kid")]
        public Input<string>? Kid { get; set; }

        /// <summary>
        /// The length of the RSA or EC certificate. This field is required when generating RSA key types.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// The name of the Key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FusionAuthKeyState()
        {
        }
        public static new FusionAuthKeyState Empty => new FusionAuthKeyState();
    }
}
