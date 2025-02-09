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
    /// ## # Form Field Resource
    /// 
    /// A FusionAuth Form Field is an object that can be customized to receive input within a FusionAuth Form.
    /// 
    /// [Form Field API](https://fusionauth.io/docs/v1/tech/apis/form-fields/)
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
    ///     var field = new Fusionauth.FusionAuthFormField("field", new()
    ///     {
    ///         Confirm = true,
    ///         Data = 
    ///         {
    ///             { "leftAddOn", "send" },
    ///         },
    ///         Description = "Information about this custom field",
    ///         Key = "user.firstName",
    ///         Required = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthFormField:FusionAuthFormField")]
    public partial class FusionAuthFormField : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if the user input should be confirmed by requiring the value to be entered twice. If true, a confirmation field is included.
        /// </summary>
        [Output("confirm")]
        public Output<bool?> Confirm { get; private set; } = null!;

        /// <summary>
        /// The Id of an existing Consent. This field will be required when the type is set to consent.
        /// </summary>
        [Output("consentId")]
        public Output<string?> ConsentId { get; private set; } = null!;

        /// <summary>
        /// The Form Field control
        /// </summary>
        [Output("control")]
        public Output<string> Control { get; private set; } = null!;

        /// <summary>
        /// An object that can hold any information about the Form Field that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, string>?> Data { get; private set; } = null!;

        /// <summary>
        /// A description of the Form Field.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Id to use for the new Form Field. If not specified a secure random UUID will be generated.
        /// </summary>
        [Output("formFieldId")]
        public Output<string?> FormFieldId { get; private set; } = null!;

        /// <summary>
        /// The key is the path to the value in the user or registration object.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The unique name of the Form Field.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of options that are applied to checkbox, radio, or select controls.
        /// </summary>
        [Output("options")]
        public Output<ImmutableArray<string>> Options { get; private set; } = null!;

        /// <summary>
        /// Determines if a value is required to complete the form.
        /// </summary>
        [Output("required")]
        public Output<bool?> Required { get; private set; } = null!;

        /// <summary>
        /// The data type used to store the value in FusionAuth.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("validator")]
        public Output<Outputs.FusionAuthFormFieldValidator> Validator { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthFormField resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthFormField(string name, FusionAuthFormFieldArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthFormField:FusionAuthFormField", name, args ?? new FusionAuthFormFieldArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthFormField(string name, Input<string> id, FusionAuthFormFieldState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthFormField:FusionAuthFormField", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FusionAuthFormField resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthFormField Get(string name, Input<string> id, FusionAuthFormFieldState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthFormField(name, id, state, options);
        }
    }

    public sealed class FusionAuthFormFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the user input should be confirmed by requiring the value to be entered twice. If true, a confirmation field is included.
        /// </summary>
        [Input("confirm")]
        public Input<bool>? Confirm { get; set; }

        /// <summary>
        /// The Id of an existing Consent. This field will be required when the type is set to consent.
        /// </summary>
        [Input("consentId")]
        public Input<string>? ConsentId { get; set; }

        /// <summary>
        /// The Form Field control
        /// </summary>
        [Input("control")]
        public Input<string>? Control { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the Form Field that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// A description of the Form Field.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Id to use for the new Form Field. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("formFieldId")]
        public Input<string>? FormFieldId { get; set; }

        /// <summary>
        /// The key is the path to the value in the user or registration object.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The unique name of the Form Field.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputList<string>? _options;

        /// <summary>
        /// A list of options that are applied to checkbox, radio, or select controls.
        /// </summary>
        public InputList<string> Options
        {
            get => _options ?? (_options = new InputList<string>());
            set => _options = value;
        }

        /// <summary>
        /// Determines if a value is required to complete the form.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The data type used to store the value in FusionAuth.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("validator")]
        public Input<Inputs.FusionAuthFormFieldValidatorArgs>? Validator { get; set; }

        public FusionAuthFormFieldArgs()
        {
        }
        public static new FusionAuthFormFieldArgs Empty => new FusionAuthFormFieldArgs();
    }

    public sealed class FusionAuthFormFieldState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the user input should be confirmed by requiring the value to be entered twice. If true, a confirmation field is included.
        /// </summary>
        [Input("confirm")]
        public Input<bool>? Confirm { get; set; }

        /// <summary>
        /// The Id of an existing Consent. This field will be required when the type is set to consent.
        /// </summary>
        [Input("consentId")]
        public Input<string>? ConsentId { get; set; }

        /// <summary>
        /// The Form Field control
        /// </summary>
        [Input("control")]
        public Input<string>? Control { get; set; }

        [Input("data")]
        private InputMap<string>? _data;

        /// <summary>
        /// An object that can hold any information about the Form Field that should be persisted.
        /// </summary>
        public InputMap<string> Data
        {
            get => _data ?? (_data = new InputMap<string>());
            set => _data = value;
        }

        /// <summary>
        /// A description of the Form Field.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Id to use for the new Form Field. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("formFieldId")]
        public Input<string>? FormFieldId { get; set; }

        /// <summary>
        /// The key is the path to the value in the user or registration object.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The unique name of the Form Field.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputList<string>? _options;

        /// <summary>
        /// A list of options that are applied to checkbox, radio, or select controls.
        /// </summary>
        public InputList<string> Options
        {
            get => _options ?? (_options = new InputList<string>());
            set => _options = value;
        }

        /// <summary>
        /// Determines if a value is required to complete the form.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The data type used to store the value in FusionAuth.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("validator")]
        public Input<Inputs.FusionAuthFormFieldValidatorGetArgs>? Validator { get; set; }

        public FusionAuthFormFieldState()
        {
        }
        public static new FusionAuthFormFieldState Empty => new FusionAuthFormFieldState();
    }
}
