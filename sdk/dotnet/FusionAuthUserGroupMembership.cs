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
    /// ## # User Group Membership Resource
    /// 
    /// [User Group Membership API](https://fusionauth.io/docs/apis/groups#request-5)
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
    ///     var @this = new Fusionauth.FusionAuthUserGroupMembership("this", new()
    ///     {
    ///         GroupId = fusionauth_group.This.Id,
    ///         UserId = fusionauth_user.This.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthUserGroupMembership:FusionAuthUserGroupMembership")]
    public partial class FusionAuthUserGroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An object that can hold any information about the User for this membership that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, string>?> Data { get; private set; } = null!;

        /// <summary>
        /// The Id of the Group of this membership.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The Id of the User Group Membership. If not provided, a random UUID will be generated.
        /// </summary>
        [Output("membershipId")]
        public Output<string> MembershipId { get; private set; } = null!;

        /// <summary>
        /// "The Id of the User of this membership.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthUserGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthUserGroupMembership(string name, FusionAuthUserGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthUserGroupMembership:FusionAuthUserGroupMembership", name, args ?? new FusionAuthUserGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthUserGroupMembership(string name, Input<string> id, FusionAuthUserGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthUserGroupMembership:FusionAuthUserGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FusionAuthUserGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthUserGroupMembership Get(string name, Input<string> id, FusionAuthUserGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthUserGroupMembership(name, id, state, options);
        }
    }

    public sealed class FusionAuthUserGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the User for this membership that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// The Id of the Group of this membership.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The Id of the User Group Membership. If not provided, a random UUID will be generated.
        /// </summary>
        [Input("membershipId")]
        public Input<string>? MembershipId { get; set; }

        /// <summary>
        /// "The Id of the User of this membership.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public FusionAuthUserGroupMembershipArgs()
        {
        }
        public static new FusionAuthUserGroupMembershipArgs Empty => new FusionAuthUserGroupMembershipArgs();
    }

    public sealed class FusionAuthUserGroupMembershipState : global::Pulumi.ResourceArgs
    {
        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the User for this membership that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// The Id of the Group of this membership.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The Id of the User Group Membership. If not provided, a random UUID will be generated.
        /// </summary>
        [Input("membershipId")]
        public Input<string>? MembershipId { get; set; }

        /// <summary>
        /// "The Id of the User of this membership.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public FusionAuthUserGroupMembershipState()
        {
        }
        public static new FusionAuthUserGroupMembershipState Empty => new FusionAuthUserGroupMembershipState();
    }
}
