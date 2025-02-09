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

    public sealed class FusionAuthSystemConfigurationCorsConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access-Control-Allow-Credentials response header values as described by MDN Access-Control-Allow-Credentials.
        /// </summary>
        [Input("allowCredentials")]
        public Input<bool>? AllowCredentials { get; set; }

        [Input("allowedHeaders")]
        private InputList<string>? _allowedHeaders;

        /// <summary>
        /// The Access-Control-Allow-Headers response header values as described by MDN Access-Control-Allow-Headers.
        /// </summary>
        public InputList<string> AllowedHeaders
        {
            get => _allowedHeaders ?? (_allowedHeaders = new InputList<string>());
            set => _allowedHeaders = value;
        }

        [Input("allowedMethods")]
        private InputList<string>? _allowedMethods;

        /// <summary>
        /// The Access-Control-Allow-Methods response header values as described by MDN Access-Control-Allow-Methods.
        /// </summary>
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        [Input("allowedOrigins")]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// The Access-Control-Allow-Origin response header values as described by MDN Access-Control-Allow-Origin. If the wildcard * is specified, no additional domains may be specified.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        /// <summary>
        /// Whether or not FusionAuth will log debug messages to the event log. This is primarily useful for identifying why the FusionAuth CORS filter is rejecting a request and returning an HTTP response status code of 403.
        /// </summary>
        [Input("debug")]
        public Input<bool>? Debug { get; set; }

        /// <summary>
        /// Whether the FusionAuth CORS filter will process requests made to FusionAuth.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exposedHeaders")]
        private InputList<string>? _exposedHeaders;

        /// <summary>
        /// The Access-Control-Expose-Headers response header values as described by MDN Access-Control-Expose-Headers.
        /// </summary>
        public InputList<string> ExposedHeaders
        {
            get => _exposedHeaders ?? (_exposedHeaders = new InputList<string>());
            set => _exposedHeaders = value;
        }

        /// <summary>
        /// The Access-Control-Max-Age response header values as described by MDN Access-Control-Max-Age.
        /// </summary>
        [Input("preflightMaxAgeInSeconds")]
        public Input<int>? PreflightMaxAgeInSeconds { get; set; }

        public FusionAuthSystemConfigurationCorsConfigurationGetArgs()
        {
        }
        public static new FusionAuthSystemConfigurationCorsConfigurationGetArgs Empty => new FusionAuthSystemConfigurationCorsConfigurationGetArgs();
    }
}
