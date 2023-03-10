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

    public sealed class FusionAuthUserTwoFactorMethodArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used by the TOTP authenticator. With the current implementation, this will always be HmacSHA1.
        /// </summary>
        [Input("authenticatorAlgorithm")]
        public Input<string>? AuthenticatorAlgorithm { get; set; }

        /// <summary>
        /// The length of code generated by the TOTP. With the current implementation, this will always be 6.
        /// </summary>
        [Input("authenticatorCodeLength")]
        public Input<int>? AuthenticatorCodeLength { get; set; }

        /// <summary>
        /// The time-step size in seconds. With the current implementation, this will always be 30.
        /// </summary>
        [Input("authenticatorTimeStep")]
        public Input<int>? AuthenticatorTimeStep { get; set; }

        /// <summary>
        /// The value of the email address for this method.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The type of this method. There will also be an object with the same value containing additional information about this method.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// The value of the mobile phone for this method.
        /// </summary>
        [Input("mobilePhone")]
        public Input<string>? MobilePhone { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// A base64 encoded secret
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("twoFactorMethodId")]
        public Input<string>? TwoFactorMethodId { get; set; }

        public FusionAuthUserTwoFactorMethodArgs()
        {
        }
        public static new FusionAuthUserTwoFactorMethodArgs Empty => new FusionAuthUserTwoFactorMethodArgs();
    }
}
