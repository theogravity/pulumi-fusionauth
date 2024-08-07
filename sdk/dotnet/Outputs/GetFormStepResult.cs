// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth.Outputs
{

    [OutputType]
    public sealed class GetFormStepResult
    {
        /// <summary>
        /// An ordered list of Form Field Ids assigned to this step.
        /// </summary>
        public readonly ImmutableArray<string> Fields;

        [OutputConstructor]
        private GetFormStepResult(ImmutableArray<string> fields)
        {
            Fields = fields;
        }
    }
}
