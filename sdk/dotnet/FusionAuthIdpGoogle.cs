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
    /// ## # Google Identity Provider Resource
    /// 
    /// The Google identity provider type will use the Google OAuth v2.0 login API. it will provide a Login with Google button on FusionAuth’s login page that will leverage the Google login pop-up dialog. Additionally, this identity provider will call Google’s /oauth2/v3/tokeninfo API to load additional details about the user and store them in FusionAuth.
    /// 
    /// The email address returned by the Google Token info API will be used to create or lookup the existing user. Additional claims returned by Google can be used to reconcile the User to FusionAuth by using a Google Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available claims returned by Google.
    /// 
    /// FusionAuth will also store the Google access_token that is returned from the login pop-up dialog in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
    /// 
    /// [Google Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/google#create-the-google-identity-provider)
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
    ///     var google = new Fusionauth.FusionAuthIdpGoogle("google", new()
    ///     {
    ///         ApplicationConfigurations = new[]
    ///         {
    ///             new Fusionauth.Inputs.FusionAuthIdpGoogleApplicationConfigurationArgs
    ///             {
    ///                 ApplicationId = fusionauth_application.Myapp.Id,
    ///                 CreateRegistration = true,
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         ButtonText = "Login with Google",
    ///         Debug = false,
    ///         ClientId = "254311943570-8e2i2hds0qdnee4124socceeh2q2mtjl.apps.googleusercontent.com",
    ///         ClientSecret = "BRr7x7xz_-cXxIFznBDIdxF1",
    ///         Scope = "profile",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle")]
    public partial class FusionAuthIdpGoogle : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration for each Application that the identity provider is enabled for.
        /// </summary>
        [Output("applicationConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthIdpGoogleApplicationConfiguration>> ApplicationConfigurations { get; private set; } = null!;

        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        [Output("buttonText")]
        public Output<string> ButtonText { get; private set; } = null!;

        /// <summary>
        /// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
        /// </summary>
        [Output("debug")]
        public Output<bool?> Debug { get; private set; } = null!;

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
        /// </summary>
        [Output("lambdaReconcileId")]
        public Output<string?> LambdaReconcileId { get; private set; } = null!;

        /// <summary>
        /// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        /// </summary>
        [Output("linkingStrategy")]
        public Output<string> LinkingStrategy { get; private set; } = null!;

        /// <summary>
        /// The login method to use for this Identity Provider.
        /// </summary>
        [Output("loginMethod")]
        public Output<string?> LoginMethod { get; private set; } = null!;

        /// <summary>
        /// An object to hold configuration parameters for the Google Identity Services API.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.FusionAuthIdpGoogleProperties?> Properties { get; private set; } = null!;

        /// <summary>
        /// The top-level scope that you are requesting from Google.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        /// </summary>
        [Output("tenantConfigurations")]
        public Output<ImmutableArray<Outputs.FusionAuthIdpGoogleTenantConfiguration>> TenantConfigurations { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthIdpGoogle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthIdpGoogle(string name, FusionAuthIdpGoogleArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle", name, args ?? new FusionAuthIdpGoogleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthIdpGoogle(string name, Input<string> id, FusionAuthIdpGoogleState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle", name, state, MakeResourceOptions(options, id))
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
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthIdpGoogle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthIdpGoogle Get(string name, Input<string> id, FusionAuthIdpGoogleState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthIdpGoogle(name, id, state, options);
        }
    }

    public sealed class FusionAuthIdpGoogleArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationConfigurations")]
        private InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationArgs>? _applicationConfigurations;

        /// <summary>
        /// The configuration for each Application that the identity provider is enabled for.
        /// </summary>
        public InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationArgs> ApplicationConfigurations
        {
            get => _applicationConfigurations ?? (_applicationConfigurations = new InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationArgs>());
            set => _applicationConfigurations = value;
        }

        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        [Input("buttonText", required: true)]
        public Input<string> ButtonText { get; set; } = null!;

        /// <summary>
        /// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
        /// </summary>
        [Input("debug")]
        public Input<bool>? Debug { get; set; }

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
        /// </summary>
        [Input("lambdaReconcileId")]
        public Input<string>? LambdaReconcileId { get; set; }

        /// <summary>
        /// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        /// </summary>
        [Input("linkingStrategy")]
        public Input<string>? LinkingStrategy { get; set; }

        /// <summary>
        /// The login method to use for this Identity Provider.
        /// </summary>
        [Input("loginMethod")]
        public Input<string>? LoginMethod { get; set; }

        /// <summary>
        /// An object to hold configuration parameters for the Google Identity Services API.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.FusionAuthIdpGooglePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The top-level scope that you are requesting from Google.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tenantConfigurations")]
        private InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationArgs>? _tenantConfigurations;

        /// <summary>
        /// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        /// </summary>
        public InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationArgs> TenantConfigurations
        {
            get => _tenantConfigurations ?? (_tenantConfigurations = new InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationArgs>());
            set => _tenantConfigurations = value;
        }

        public FusionAuthIdpGoogleArgs()
        {
        }
        public static new FusionAuthIdpGoogleArgs Empty => new FusionAuthIdpGoogleArgs();
    }

    public sealed class FusionAuthIdpGoogleState : global::Pulumi.ResourceArgs
    {
        [Input("applicationConfigurations")]
        private InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationGetArgs>? _applicationConfigurations;

        /// <summary>
        /// The configuration for each Application that the identity provider is enabled for.
        /// </summary>
        public InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationGetArgs> ApplicationConfigurations
        {
            get => _applicationConfigurations ?? (_applicationConfigurations = new InputList<Inputs.FusionAuthIdpGoogleApplicationConfigurationGetArgs>());
            set => _applicationConfigurations = value;
        }

        /// <summary>
        /// The top-level button text to use on the FusionAuth login page for this Identity Provider.
        /// </summary>
        [Input("buttonText")]
        public Input<string>? ButtonText { get; set; }

        /// <summary>
        /// The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
        /// </summary>
        [Input("debug")]
        public Input<bool>? Debug { get; set; }

        /// <summary>
        /// Determines if this provider is enabled. If it is false then it will be disabled globally.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
        /// </summary>
        [Input("lambdaReconcileId")]
        public Input<string>? LambdaReconcileId { get; set; }

        /// <summary>
        /// The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        /// </summary>
        [Input("linkingStrategy")]
        public Input<string>? LinkingStrategy { get; set; }

        /// <summary>
        /// The login method to use for this Identity Provider.
        /// </summary>
        [Input("loginMethod")]
        public Input<string>? LoginMethod { get; set; }

        /// <summary>
        /// An object to hold configuration parameters for the Google Identity Services API.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.FusionAuthIdpGooglePropertiesGetArgs>? Properties { get; set; }

        /// <summary>
        /// The top-level scope that you are requesting from Google.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tenantConfigurations")]
        private InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationGetArgs>? _tenantConfigurations;

        /// <summary>
        /// The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        /// </summary>
        public InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationGetArgs> TenantConfigurations
        {
            get => _tenantConfigurations ?? (_tenantConfigurations = new InputList<Inputs.FusionAuthIdpGoogleTenantConfigurationGetArgs>());
            set => _tenantConfigurations = value;
        }

        public FusionAuthIdpGoogleState()
        {
        }
        public static new FusionAuthIdpGoogleState Empty => new FusionAuthIdpGoogleState();
    }
}
