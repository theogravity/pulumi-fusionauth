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
    public sealed class FusionAuthApplicationLambdaConfiguration
    {
        /// <summary>
        /// The Id of the Lambda that will be invoked when an access token is generated for this application. This will be utilized during OAuth2 and OpenID Connect authentication requests as well as when an access token is generated for the Login API.
        /// </summary>
        public readonly string? AccessTokenPopulateId;
        /// <summary>
        /// The Id of the Lambda that will be invoked when an Id token is generated for this application during an OpenID Connect authentication request.
        /// </summary>
        public readonly string? IdTokenPopulateId;
        /// <summary>
        /// The Id of the Lambda that will be invoked when a a SAML response is generated during a SAML authentication request.
        /// </summary>
        public readonly string? Samlv2PopulateId;
        /// <summary>
        /// The unique Id of the lambda that will be used to perform additional validation on registration form steps.
        /// </summary>
        public readonly string? SelfServiceRegistrationValidationId;
        /// <summary>
        /// The Id of the Lambda that will be invoked when a UserInfo response is generated for this application.
        /// </summary>
        public readonly string? UserinfoPopulateId;

        [OutputConstructor]
        private FusionAuthApplicationLambdaConfiguration(
            string? accessTokenPopulateId,

            string? idTokenPopulateId,

            string? samlv2PopulateId,

            string? selfServiceRegistrationValidationId,

            string? userinfoPopulateId)
        {
            AccessTokenPopulateId = accessTokenPopulateId;
            IdTokenPopulateId = idTokenPopulateId;
            Samlv2PopulateId = samlv2PopulateId;
            SelfServiceRegistrationValidationId = selfServiceRegistrationValidationId;
            UserinfoPopulateId = userinfoPopulateId;
        }
    }
}
