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
    public static class GetApplicationRole
    {
        /// <summary>
        /// ## # Application Role Resource
        /// 
        /// This Resource is used to create a role for an Application.
        /// 
        /// [Application Roles API](https://fusionauth.io/docs/v1/tech/apis/applications)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var admin = Output.Create(Fusionauth.GetApplicationRole.InvokeAsync(new Fusionauth.GetApplicationRoleArgs
        ///         {
        ///             ApplicationId = data.Fusionauth_application.FusionAuth.Id,
        ///             Name = "admin",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationRoleResult> InvokeAsync(GetApplicationRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationRoleResult>("fusionauth:index/getApplicationRole:getApplicationRole", args ?? new GetApplicationRoleArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application Role Resource
        /// 
        /// This Resource is used to create a role for an Application.
        /// 
        /// [Application Roles API](https://fusionauth.io/docs/v1/tech/apis/applications)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var admin = Output.Create(Fusionauth.GetApplicationRole.InvokeAsync(new Fusionauth.GetApplicationRoleArgs
        ///         {
        ///             ApplicationId = data.Fusionauth_application.FusionAuth.Id,
        ///             Name = "admin",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationRoleResult> Invoke(GetApplicationRoleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationRoleResult>("fusionauth:index/getApplicationRole:getApplicationRole", args ?? new GetApplicationRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationRoleArgs : Pulumi.InvokeArgs
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

        public GetApplicationRoleArgs()
        {
        }
    }

    public sealed class GetApplicationRoleInvokeArgs : Pulumi.InvokeArgs
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

        public GetApplicationRoleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationRoleResult
    {
        public readonly string ApplicationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetApplicationRoleResult(
            string applicationId,

            string id,

            string name)
        {
            ApplicationId = applicationId;
            Id = id;
            Name = name;
        }
    }
}
