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
    public static class GetLambda
    {
        /// <summary>
        /// ## # Lambda Resource
        /// 
        /// Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.
        /// 
        /// [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)
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
        ///     var defaultGoogleReconcile = Fusionauth.GetLambda.Invoke(new()
        ///     {
        ///         Name = "Default Google Reconcile provided by FusionAuth",
        ///         Type = "GoogleReconcile",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLambdaResult> InvokeAsync(GetLambdaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLambdaResult>("fusionauth:index/getLambda:getLambda", args ?? new GetLambdaArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Lambda Resource
        /// 
        /// Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.
        /// 
        /// [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)
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
        ///     var defaultGoogleReconcile = Fusionauth.GetLambda.Invoke(new()
        ///     {
        ///         Name = "Default Google Reconcile provided by FusionAuth",
        ///         Type = "GoogleReconcile",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLambdaResult> Invoke(GetLambdaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLambdaResult>("fusionauth:index/getLambda:getLambda", args ?? new GetLambdaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLambdaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Lambda. At least one of `id` or `name` must be specified.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the Lambda. At least one of `id` or `name` must be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Lambda type. The possible values are:
        /// - `JWTPopulate`
        /// - `OpenIDReconcile`
        /// - `SAMLv2Reconcile`
        /// - `SAMLv2Populate`
        /// - `AppleReconcile`
        /// - `ExternalJWTReconcile`
        /// - `FacebookReconcile`
        /// - `GoogleReconcile`
        /// - `HYPRReconcile`
        /// - `TwitterReconcile`
        /// - `LDAPConnectorReconcile`
        /// - `LinkedInReconcile`
        /// - `EpicGamesReconcile`
        /// - `NintendoReconcile`
        /// - `SonyPSNReconcile`
        /// - `SteamReconcile`
        /// - `TwitchReconcile`
        /// - `XboxReconcile`
        /// - `SelfServiceRegistrationValidation`
        /// - `ClientCredentialsJWTPopulate`
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetLambdaArgs()
        {
        }
        public static new GetLambdaArgs Empty => new GetLambdaArgs();
    }

    public sealed class GetLambdaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Lambda. At least one of `id` or `name` must be specified.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the Lambda. At least one of `id` or `name` must be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Lambda type. The possible values are:
        /// - `JWTPopulate`
        /// - `OpenIDReconcile`
        /// - `SAMLv2Reconcile`
        /// - `SAMLv2Populate`
        /// - `AppleReconcile`
        /// - `ExternalJWTReconcile`
        /// - `FacebookReconcile`
        /// - `GoogleReconcile`
        /// - `HYPRReconcile`
        /// - `TwitterReconcile`
        /// - `LDAPConnectorReconcile`
        /// - `LinkedInReconcile`
        /// - `EpicGamesReconcile`
        /// - `NintendoReconcile`
        /// - `SonyPSNReconcile`
        /// - `SteamReconcile`
        /// - `TwitchReconcile`
        /// - `XboxReconcile`
        /// - `SelfServiceRegistrationValidation`
        /// - `ClientCredentialsJWTPopulate`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetLambdaInvokeArgs()
        {
        }
        public static new GetLambdaInvokeArgs Empty => new GetLambdaInvokeArgs();
    }


    [OutputType]
    public sealed class GetLambdaResult
    {
        /// <summary>
        /// The lambda function body, a JavaScript function.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// Whether or not debug event logging is enabled for this Lambda.
        /// </summary>
        public readonly bool Debug;
        public readonly string Id;
        public readonly string? Name;
        public readonly string Type;

        [OutputConstructor]
        private GetLambdaResult(
            string body,

            bool debug,

            string id,

            string? name,

            string type)
        {
            Body = body;
            Debug = debug;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
