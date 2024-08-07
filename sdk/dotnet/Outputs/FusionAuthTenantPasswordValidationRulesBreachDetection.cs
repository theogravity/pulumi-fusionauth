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
    public sealed class FusionAuthTenantPasswordValidationRulesBreachDetection
    {
        /// <summary>
        /// Whether to enable Reactor breach detection. Requires an activated license.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The level of severity where Reactor will consider a breach.
        /// </summary>
        public readonly string? MatchMode;
        /// <summary>
        /// The Id of the email template to use when notifying user of breached password. Required if tenant.passwordValidationRules.breachDetection.onLogin is set to NotifyUser.
        /// </summary>
        public readonly string? NotifyUserEmailTemplateId;
        /// <summary>
        /// The behavior when detecting breaches at time of user login
        /// </summary>
        public readonly string? OnLogin;

        [OutputConstructor]
        private FusionAuthTenantPasswordValidationRulesBreachDetection(
            bool? enabled,

            string? matchMode,

            string? notifyUserEmailTemplateId,

            string? onLogin)
        {
            Enabled = enabled;
            MatchMode = matchMode;
            NotifyUserEmailTemplateId = notifyUserEmailTemplateId;
            OnLogin = onLogin;
        }
    }
}
