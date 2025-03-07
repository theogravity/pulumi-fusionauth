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
    public static class GetApplicationOAuthScope
    {
        /// <summary>
        /// ## # Application OAuth Scope Resource
        /// 
        /// The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.
        /// 
        /// [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)
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
        ///     var @this = Fusionauth.GetApplicationOAuthScope.Invoke(new()
        ///     {
        ///         ApplicationId = data.Fusionauth_application.This.Id,
        ///         Name = "data:read",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetApplicationOAuthScopeResult> InvokeAsync(GetApplicationOAuthScopeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationOAuthScopeResult>("fusionauth:index/getApplicationOAuthScope:getApplicationOAuthScope", args ?? new GetApplicationOAuthScopeArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application OAuth Scope Resource
        /// 
        /// The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.
        /// 
        /// [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)
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
        ///     var @this = Fusionauth.GetApplicationOAuthScope.Invoke(new()
        ///     {
        ///         ApplicationId = data.Fusionauth_application.This.Id,
        ///         Name = "data:read",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetApplicationOAuthScopeResult> Invoke(GetApplicationOAuthScopeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationOAuthScopeResult>("fusionauth:index/getApplicationOAuthScope:getApplicationOAuthScope", args ?? new GetApplicationOAuthScopeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application OAuth Scope Resource
        /// 
        /// The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.
        /// 
        /// [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)
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
        ///     var @this = Fusionauth.GetApplicationOAuthScope.Invoke(new()
        ///     {
        ///         ApplicationId = data.Fusionauth_application.This.Id,
        ///         Name = "data:read",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetApplicationOAuthScopeResult> Invoke(GetApplicationOAuthScopeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationOAuthScopeResult>("fusionauth:index/getApplicationOAuthScope:getApplicationOAuthScope", args ?? new GetApplicationOAuthScopeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationOAuthScopeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the application that this role is for.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetApplicationOAuthScopeArgs()
        {
        }
        public static new GetApplicationOAuthScopeArgs Empty => new GetApplicationOAuthScopeArgs();
    }

    public sealed class GetApplicationOAuthScopeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the application that this role is for.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetApplicationOAuthScopeInvokeArgs()
        {
        }
        public static new GetApplicationOAuthScopeInvokeArgs Empty => new GetApplicationOAuthScopeInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationOAuthScopeResult
    {
        public readonly string ApplicationId;
        /// <summary>
        /// (Optional) An object that can hold any information about the OAuth Scope that should be persisted.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Data;
        /// <summary>
        /// (Optional) "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        /// </summary>
        public readonly string DefaultConsentDetail;
        /// <summary>
        /// (Optional) The default message to display on the OAuth consent screen if one cannot be found in the theme.
        /// </summary>
        public readonly string DefaultConsentMessage;
        /// <summary>
        /// (Optional) A description of the OAuth Scope. This is used for display purposes only.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// (Optional) Determines if the OAuth Scope is required when requested in an OAuth workflow.
        /// </summary>
        public readonly bool Required;
        /// <summary>
        /// (Optional) The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        /// </summary>
        public readonly string ScopeId;

        [OutputConstructor]
        private GetApplicationOAuthScopeResult(
            string applicationId,

            ImmutableDictionary<string, string> data,

            string defaultConsentDetail,

            string defaultConsentMessage,

            string description,

            string id,

            string name,

            bool required,

            string scopeId)
        {
            ApplicationId = applicationId;
            Data = data;
            DefaultConsentDetail = defaultConsentDetail;
            DefaultConsentMessage = defaultConsentMessage;
            Description = description;
            Id = id;
            Name = name;
            Required = required;
            ScopeId = scopeId;
        }
    }
}
