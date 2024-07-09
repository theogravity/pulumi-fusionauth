// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Inputs
{

    public sealed class FusionAuthFormFieldValidatorGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if user input should be validated.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// A regular expression used to validate user input. Must be a valid regular expression pattern.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        public FusionAuthFormFieldValidatorGetArgs()
        {
        }
        public static new FusionAuthFormFieldValidatorGetArgs Empty => new FusionAuthFormFieldValidatorGetArgs();
    }
}
