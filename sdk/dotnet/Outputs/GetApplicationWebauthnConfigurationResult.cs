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
    public sealed class GetApplicationWebauthnConfigurationResult
    {
        public readonly bool BootstrapWorkflowEnabled;
        public readonly bool Enabled;
        public readonly bool ReauthenticationWorkflowEnabled;

        [OutputConstructor]
        private GetApplicationWebauthnConfigurationResult(
            bool bootstrapWorkflowEnabled,

            bool enabled,

            bool reauthenticationWorkflowEnabled)
        {
            BootstrapWorkflowEnabled = bootstrapWorkflowEnabled;
            Enabled = enabled;
            ReauthenticationWorkflowEnabled = reauthenticationWorkflowEnabled;
        }
    }
}
