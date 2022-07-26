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
    public sealed class FusionAuthApplicationEmailConfiguration
    {
        /// <summary>
        /// The Id of the Email Template used to send emails to users when their email address is updated. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? EmailUpdateTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users to verify that their email address is valid. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? EmailVerificationTemplateId;
        /// <summary>
        /// The Id of the Email Template used to verify user emails. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? EmailVerifiedTemplateId;
        /// <summary>
        /// The Id of the Email Template that is used when a user is sent a forgot password email. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? ForgotPasswordTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when another user attempts to create an account with their login Id. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? LoginIdInUseOnCreateTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when another user attempts to update an existing account to use their login Id. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? LoginIdInUseOnUpdateTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when they log in on a new device. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? LoginNewDeviceTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when a suspicious login occurs. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? LoginSuspiciousTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when they have completed a 'forgot password' workflow and their password has been reset. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? PasswordResetSuccessTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when their password has been updated. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? PasswordUpdateTemplateId;
        /// <summary>
        /// The Id of the Passwordless Email Template, sent to users when they start a passwordless login. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? PasswordlessEmailTemplateId;
        /// <summary>
        /// The Id of the Email Template that is used when a user had their account created for them and they must set their password manually and they are sent an email to set their password. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? SetPasswordEmailTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when a MFA method has been added to their account. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? TwoFactorMethodAddTemplateId;
        /// <summary>
        /// The Id of the Email Template used to send emails to users when a MFA method has been removed from their account. When configured, this value will take precedence over the same configuration from the Tenant when an application context is known.
        /// </summary>
        public readonly string? TwoFactorMethodRemoveTemplateId;

        [OutputConstructor]
        private FusionAuthApplicationEmailConfiguration(
            string? emailUpdateTemplateId,

            string? emailVerificationTemplateId,

            string? emailVerifiedTemplateId,

            string? forgotPasswordTemplateId,

            string? loginIdInUseOnCreateTemplateId,

            string? loginIdInUseOnUpdateTemplateId,

            string? loginNewDeviceTemplateId,

            string? loginSuspiciousTemplateId,

            string? passwordResetSuccessTemplateId,

            string? passwordUpdateTemplateId,

            string? passwordlessEmailTemplateId,

            string? setPasswordEmailTemplateId,

            string? twoFactorMethodAddTemplateId,

            string? twoFactorMethodRemoveTemplateId)
        {
            EmailUpdateTemplateId = emailUpdateTemplateId;
            EmailVerificationTemplateId = emailVerificationTemplateId;
            EmailVerifiedTemplateId = emailVerifiedTemplateId;
            ForgotPasswordTemplateId = forgotPasswordTemplateId;
            LoginIdInUseOnCreateTemplateId = loginIdInUseOnCreateTemplateId;
            LoginIdInUseOnUpdateTemplateId = loginIdInUseOnUpdateTemplateId;
            LoginNewDeviceTemplateId = loginNewDeviceTemplateId;
            LoginSuspiciousTemplateId = loginSuspiciousTemplateId;
            PasswordResetSuccessTemplateId = passwordResetSuccessTemplateId;
            PasswordUpdateTemplateId = passwordUpdateTemplateId;
            PasswordlessEmailTemplateId = passwordlessEmailTemplateId;
            SetPasswordEmailTemplateId = setPasswordEmailTemplateId;
            TwoFactorMethodAddTemplateId = twoFactorMethodAddTemplateId;
            TwoFactorMethodRemoveTemplateId = twoFactorMethodRemoveTemplateId;
        }
    }
}
