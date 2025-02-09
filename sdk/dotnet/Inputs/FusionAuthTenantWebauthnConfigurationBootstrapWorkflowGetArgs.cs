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

    public sealed class FusionAuthTenantWebauthnConfigurationBootstrapWorkflowGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the authenticator attachment requirement for WebAuthn passkey registration when using the bootstrap workflow. The possible values are: Any, CrossPlatform and Platform. Note: A license is required to utilize WebAuthn and an Enterprise plan is required to utilize WebAuthn cross-platform authenticators..
        /// </summary>
        [Input("authenticatorAttachmentPreference")]
        public Input<string>? AuthenticatorAttachmentPreference { get; set; }

        /// <summary>
        /// Whether or not this tenant has the WebAuthn bootstrap workflow enabled. The bootstrap workflow is used when the user must "bootstrap" the authentication process by identifying themselves prior to the WebAuthn ceremony and can be used to authenticate from a new device using WebAuthn. Note: A license is required to utilize WebAuthn..
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Determines the user verification requirement for WebAuthn passkey registration when using the bootstrap workflow. The possible values are: Discouraged, Preferred and Required. Note: A license is required to utilize WebAuthn..
        /// </summary>
        [Input("userVerificationRequirement")]
        public Input<string>? UserVerificationRequirement { get; set; }

        public FusionAuthTenantWebauthnConfigurationBootstrapWorkflowGetArgs()
        {
        }
        public static new FusionAuthTenantWebauthnConfigurationBootstrapWorkflowGetArgs Empty => new FusionAuthTenantWebauthnConfigurationBootstrapWorkflowGetArgs();
    }
}
