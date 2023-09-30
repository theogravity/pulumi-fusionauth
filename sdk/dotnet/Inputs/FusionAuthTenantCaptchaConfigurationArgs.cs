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

    public sealed class FusionAuthTenantCaptchaConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of captcha method to use. This field is required when tenant.captchaConfiguration.enabled is set to true.
        /// </summary>
        [Input("captchaMethod")]
        public Input<string>? CaptchaMethod { get; set; }

        /// <summary>
        /// When true, FusionAuth will handle username collisions by generating a random suffix.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The secret key for this captcha method. This field is required when tenant.captchaConfiguration.enabled is set to true.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        /// <summary>
        /// The site key for this captcha method. This field is required when tenant.captchaConfiguration.enabled is set to true.
        /// </summary>
        [Input("siteKey")]
        public Input<string>? SiteKey { get; set; }

        /// <summary>
        /// The numeric threshold which separates a passing score from a failing one. This value only applies if using either the Google v3 or HCaptcha Enterprise method, otherwise this value is ignored.
        /// </summary>
        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        public FusionAuthTenantCaptchaConfigurationArgs()
        {
        }
        public static new FusionAuthTenantCaptchaConfigurationArgs Empty => new FusionAuthTenantCaptchaConfigurationArgs();
    }
}
