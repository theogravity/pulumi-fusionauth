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
    public static class GetIdp
    {
        /// <summary>
        /// ## # Application Resource
        /// 
        /// [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fusionAuth = Fusionauth.GetIdp.Invoke(new()
        ///     {
        ///         Name = "Apple",
        ///         Type = "Apple",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIdpResult> InvokeAsync(GetIdpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdpResult>("fusionauth:index/getIdp:getIdp", args ?? new GetIdpArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application Resource
        /// 
        /// [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fusionAuth = Fusionauth.GetIdp.Invoke(new()
        ///     {
        ///         Name = "Apple",
        ///         Type = "Apple",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIdpResult> Invoke(GetIdpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdpResult>("fusionauth:index/getIdp:getIdp", args ?? new GetIdpInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application Resource
        /// 
        /// [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fusionAuth = Fusionauth.GetIdp.Invoke(new()
        ///     {
        ///         Name = "Apple",
        ///         Type = "Apple",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIdpResult> Invoke(GetIdpInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdpResult>("fusionauth:index/getIdp:getIdp", args ?? new GetIdpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The type of the identity provider.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetIdpArgs()
        {
        }
        public static new GetIdpArgs Empty => new GetIdpArgs();
    }

    public sealed class GetIdpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of the identity provider.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetIdpInvokeArgs()
        {
        }
        public static new GetIdpInvokeArgs Empty => new GetIdpInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdpResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetIdpResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
