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
    public sealed class FusionAuthApplicationRegistrationConfigurationFirstName
    {
        public readonly bool? Enabled;
        public readonly bool? Required;

        [OutputConstructor]
        private FusionAuthApplicationRegistrationConfigurationFirstName(
            bool? enabled,

            bool? required)
        {
            Enabled = enabled;
            Required = required;
        }
    }
}
