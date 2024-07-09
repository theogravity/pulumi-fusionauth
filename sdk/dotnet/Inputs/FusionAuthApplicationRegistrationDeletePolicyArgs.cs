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

    public sealed class FusionAuthApplicationRegistrationDeletePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates that users without a verified registration for this application will have their registration permanently deleted after application.registrationDeletePolicy.unverified.numberOfDaysToRetain days.
        /// </summary>
        [Input("unverifiedEnabled")]
        public Input<bool>? UnverifiedEnabled { get; set; }

        /// <summary>
        /// The number of days from registration a user’s registration will be retained before being deleted for not completing registration verification. This field is required when application.registrationDeletePolicy.enabled is set to true. Value must be greater than 0.
        /// </summary>
        [Input("unverifiedNumberOfDaysToRetain")]
        public Input<int>? UnverifiedNumberOfDaysToRetain { get; set; }

        public FusionAuthApplicationRegistrationDeletePolicyArgs()
        {
        }
        public static new FusionAuthApplicationRegistrationDeletePolicyArgs Empty => new FusionAuthApplicationRegistrationDeletePolicyArgs();
    }
}
