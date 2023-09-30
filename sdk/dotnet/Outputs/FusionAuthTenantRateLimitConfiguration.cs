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
    public sealed class FusionAuthTenantRateLimitConfiguration
    {
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationFailedLogin? FailedLogin;
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationForgotPassword? ForgotPassword;
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationSendEmailVerification? SendEmailVerification;
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationSendPasswordless? SendPasswordless;
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationSendRegistrationVerification? SendRegistrationVerification;
        public readonly Outputs.FusionAuthTenantRateLimitConfigurationSendTwoFactor? SendTwoFactor;

        [OutputConstructor]
        private FusionAuthTenantRateLimitConfiguration(
            Outputs.FusionAuthTenantRateLimitConfigurationFailedLogin? failedLogin,

            Outputs.FusionAuthTenantRateLimitConfigurationForgotPassword? forgotPassword,

            Outputs.FusionAuthTenantRateLimitConfigurationSendEmailVerification? sendEmailVerification,

            Outputs.FusionAuthTenantRateLimitConfigurationSendPasswordless? sendPasswordless,

            Outputs.FusionAuthTenantRateLimitConfigurationSendRegistrationVerification? sendRegistrationVerification,

            Outputs.FusionAuthTenantRateLimitConfigurationSendTwoFactor? sendTwoFactor)
        {
            FailedLogin = failedLogin;
            ForgotPassword = forgotPassword;
            SendEmailVerification = sendEmailVerification;
            SendPasswordless = sendPasswordless;
            SendRegistrationVerification = sendRegistrationVerification;
            SendTwoFactor = sendTwoFactor;
        }
    }
}
