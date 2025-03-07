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
    /// ## # System Configuration Resource
    /// 
    /// A registration is the association between a User and an Application that they log into.
    /// 
    /// [System Configuration API](https://fusionauth.io/docs/v1/tech/apis/system)
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
    ///     var example = new Fusionauth.FusionAuthSystemConfiguration("example", new()
    ///     {
    ///         AuditLogConfiguration = new Fusionauth.Inputs.FusionAuthSystemConfigurationAuditLogConfigurationArgs
    ///         {
    ///             Delete = new Fusionauth.Inputs.FusionAuthSystemConfigurationAuditLogConfigurationDeleteArgs
    ///             {
    ///                 Enabled = true,
    ///                 NumberOfDaysToRetain = 367,
    ///             },
    ///         },
    ///         CorsConfiguration = new Fusionauth.Inputs.FusionAuthSystemConfigurationCorsConfigurationArgs
    ///         {
    ///             AllowedMethods = new[]
    ///             {
    ///                 "POST",
    ///                 "PUT",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration")]
    public partial class FusionAuthSystemConfiguration : global::Pulumi.CustomResource
    {
        [Output("auditLogConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationAuditLogConfiguration> AuditLogConfiguration { get; private set; } = null!;

        [Output("corsConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationCorsConfiguration> CorsConfiguration { get; private set; } = null!;

        /// <summary>
        /// An object that can hold any information about the System that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, string>?> Data { get; private set; } = null!;

        [Output("eventLogConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationEventLogConfiguration> EventLogConfiguration { get; private set; } = null!;

        [Output("loginRecordConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationLoginRecordConfiguration> LoginRecordConfiguration { get; private set; } = null!;

        /// <summary>
        /// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
        /// </summary>
        [Output("reportTimezone")]
        public Output<string?> ReportTimezone { get; private set; } = null!;

        /// <summary>
        /// The trusted proxy configuration.
        /// </summary>
        [Output("trustedProxyConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationTrustedProxyConfiguration?> TrustedProxyConfiguration { get; private set; } = null!;

        [Output("uiConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationUiConfiguration> UiConfiguration { get; private set; } = null!;

        /// <summary>
        /// The usage data configuration.
        /// </summary>
        [Output("usageDataConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationUsageDataConfiguration?> UsageDataConfiguration { get; private set; } = null!;

        [Output("webhookEventLogConfiguration")]
        public Output<Outputs.FusionAuthSystemConfigurationWebhookEventLogConfiguration> WebhookEventLogConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthSystemConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthSystemConfiguration(string name, FusionAuthSystemConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration", name, args ?? new FusionAuthSystemConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthSystemConfiguration(string name, Input<string> id, FusionAuthSystemConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FusionAuthSystemConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthSystemConfiguration Get(string name, Input<string> id, FusionAuthSystemConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthSystemConfiguration(name, id, state, options);
        }
    }

    public sealed class FusionAuthSystemConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationAuditLogConfigurationArgs>? AuditLogConfiguration { get; set; }

        [Input("corsConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationCorsConfigurationArgs>? CorsConfiguration { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the System that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        [Input("eventLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationEventLogConfigurationArgs>? EventLogConfiguration { get; set; }

        [Input("loginRecordConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationLoginRecordConfigurationArgs>? LoginRecordConfiguration { get; set; }

        /// <summary>
        /// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
        /// </summary>
        [Input("reportTimezone")]
        public Input<string>? ReportTimezone { get; set; }

        /// <summary>
        /// The trusted proxy configuration.
        /// </summary>
        [Input("trustedProxyConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationTrustedProxyConfigurationArgs>? TrustedProxyConfiguration { get; set; }

        [Input("uiConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationUiConfigurationArgs>? UiConfiguration { get; set; }

        /// <summary>
        /// The usage data configuration.
        /// </summary>
        [Input("usageDataConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationUsageDataConfigurationArgs>? UsageDataConfiguration { get; set; }

        [Input("webhookEventLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationWebhookEventLogConfigurationArgs>? WebhookEventLogConfiguration { get; set; }

        public FusionAuthSystemConfigurationArgs()
        {
        }
        public static new FusionAuthSystemConfigurationArgs Empty => new FusionAuthSystemConfigurationArgs();
    }

    public sealed class FusionAuthSystemConfigurationState : global::Pulumi.ResourceArgs
    {
        [Input("auditLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationAuditLogConfigurationGetArgs>? AuditLogConfiguration { get; set; }

        [Input("corsConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationCorsConfigurationGetArgs>? CorsConfiguration { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the System that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        [Input("eventLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationEventLogConfigurationGetArgs>? EventLogConfiguration { get; set; }

        [Input("loginRecordConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationLoginRecordConfigurationGetArgs>? LoginRecordConfiguration { get; set; }

        /// <summary>
        /// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
        /// </summary>
        [Input("reportTimezone")]
        public Input<string>? ReportTimezone { get; set; }

        /// <summary>
        /// The trusted proxy configuration.
        /// </summary>
        [Input("trustedProxyConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationTrustedProxyConfigurationGetArgs>? TrustedProxyConfiguration { get; set; }

        [Input("uiConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationUiConfigurationGetArgs>? UiConfiguration { get; set; }

        /// <summary>
        /// The usage data configuration.
        /// </summary>
        [Input("usageDataConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationUsageDataConfigurationGetArgs>? UsageDataConfiguration { get; set; }

        [Input("webhookEventLogConfiguration")]
        public Input<Inputs.FusionAuthSystemConfigurationWebhookEventLogConfigurationGetArgs>? WebhookEventLogConfiguration { get; set; }

        public FusionAuthSystemConfigurationState()
        {
        }
        public static new FusionAuthSystemConfigurationState Empty => new FusionAuthSystemConfigurationState();
    }
}
