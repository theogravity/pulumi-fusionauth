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

    public sealed class FusionAuthApplicationCleanSpeakConfigurationUsernameModerationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the CleanSpeak application that usernames are sent to for moderation.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// True if CleanSpeak username moderation is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public FusionAuthApplicationCleanSpeakConfigurationUsernameModerationArgs()
        {
        }
        public static new FusionAuthApplicationCleanSpeakConfigurationUsernameModerationArgs Empty => new FusionAuthApplicationCleanSpeakConfigurationUsernameModerationArgs();
    }
}
