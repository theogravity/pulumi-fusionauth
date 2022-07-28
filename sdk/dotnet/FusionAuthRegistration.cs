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
    /// ## # Registration Resource
    /// 
    /// A registration is the association between a User and an Application that they log into.
    /// 
    /// [Registrations API](https://fusionauth.io/docs/v1/tech/apis/registrations)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Fusionauth = theogravity.Fusionauth;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Fusionauth.FusionAuthRegistration("example", new Fusionauth.FusionAuthRegistrationArgs
    ///         {
    ///             UserId = fusionauth_user.Example.Id,
    ///             ApplicationId = data.Fusionauth_application.FusionAuth.Id,
    ///             Roles = 
    ///             {
    ///                 "admin",
    ///             },
    ///             Username = "theadmin",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthRegistration:FusionAuthRegistration")]
    public partial class FusionAuthRegistration : Pulumi.CustomResource
    {
        /// <summary>
        /// The Id of the Application that this registration is for.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
        /// </summary>
        [Output("authenticationToken")]
        public Output<string> AuthenticationToken { get; private set; } = null!;

        /// <summary>
        /// An object that can hold any information about the User for this registration that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>?> Data { get; private set; } = null!;

        /// <summary>
        /// Determines if FusionAuth should generate an Authentication Token for this registration.
        /// </summary>
        [Output("generateAuthenticationToken")]
        public Output<bool?> GenerateAuthenticationToken { get; private set; } = null!;

        /// <summary>
        /// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
        /// </summary>
        [Output("preferredLanguages")]
        public Output<ImmutableArray<string>> PreferredLanguages { get; private set; } = null!;

        /// <summary>
        /// The list of roles that the User has for this registration.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
        /// </summary>
        [Output("skipRegistrationValidation")]
        public Output<bool?> SkipRegistrationValidation { get; private set; } = null!;

        /// <summary>
        /// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// The Id of the User that is registering for the Application.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// The username of the User for this registration only.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthRegistration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthRegistration(string name, FusionAuthRegistrationArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthRegistration:FusionAuthRegistration", name, args ?? new FusionAuthRegistrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthRegistration(string name, Input<string> id, FusionAuthRegistrationState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthRegistration:FusionAuthRegistration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FusionAuthRegistration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthRegistration Get(string name, Input<string> id, FusionAuthRegistrationState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthRegistration(name, id, state, options);
        }
    }

    public sealed class FusionAuthRegistrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the Application that this registration is for.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
        /// </summary>
        [Input("authenticationToken")]
        public Input<string>? AuthenticationToken { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the User for this registration that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// Determines if FusionAuth should generate an Authentication Token for this registration.
        /// </summary>
        [Input("generateAuthenticationToken")]
        public Input<bool>? GenerateAuthenticationToken { get; set; }

        [Input("preferredLanguages")]
        private InputList<string>? _preferredLanguages;

        /// <summary>
        /// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
        /// </summary>
        public InputList<string> PreferredLanguages
        {
            get => _preferredLanguages ?? (_preferredLanguages = new InputList<string>());
            set => _preferredLanguages = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The list of roles that the User has for this registration.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
        /// </summary>
        [Input("skipRegistrationValidation")]
        public Input<bool>? SkipRegistrationValidation { get; set; }

        /// <summary>
        /// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// The Id of the User that is registering for the Application.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        /// <summary>
        /// The username of the User for this registration only.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public FusionAuthRegistrationArgs()
        {
        }
    }

    public sealed class FusionAuthRegistrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the Application that this registration is for.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The authentication token that may be used in place of the User’s password when authenticating against this application represented by this registration. This parameter is ignored if generateAuthenticationToken is set to true and instead the value will be generated.
        /// </summary>
        [Input("authenticationToken")]
        public Input<string>? AuthenticationToken { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the User for this registration that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// Determines if FusionAuth should generate an Authentication Token for this registration.
        /// </summary>
        [Input("generateAuthenticationToken")]
        public Input<bool>? GenerateAuthenticationToken { get; set; }

        [Input("preferredLanguages")]
        private InputList<string>? _preferredLanguages;

        /// <summary>
        /// An array of locale strings that give, in order, the User’s preferred languages for this registration. These are important for email templates and other localizable text.
        /// </summary>
        public InputList<string> PreferredLanguages
        {
            get => _preferredLanguages ?? (_preferredLanguages = new InputList<string>());
            set => _preferredLanguages = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The list of roles that the User has for this registration.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// Indicates to FusionAuth that it should skip registration verification even if it is enabled for the Application.
        /// </summary>
        [Input("skipRegistrationValidation")]
        public Input<bool>? SkipRegistrationValidation { get; set; }

        /// <summary>
        /// The User’s preferred timezone for this registration. The string will be in an IANA time zone format.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// The Id of the User that is registering for the Application.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// The username of the User for this registration only.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public FusionAuthRegistrationState()
        {
        }
    }
}