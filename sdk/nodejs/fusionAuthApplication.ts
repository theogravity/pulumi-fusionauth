// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Application Resource
 *
 * [Applications API](https://fusionauth.io/docs/v1/tech/apis/applications)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const forum = new fusionauth.FusionAuthApplication("forum", {
 *     tenantId: fusionauth_tenant.portal.id,
 *     authenticationTokenConfigurationEnabled: false,
 *     formConfiguration: {
 *         adminRegistrationFormId: fusionauth_form.admin_registration.id,
 *         selfServiceFormId: fusionauth_form.self_service.id,
 *     },
 *     jwtConfiguration: {
 *         accessTokenId: fusionauth_key.access_token.id,
 *         enabled: true,
 *         idTokenKeyId: fusionauth_key.id_token.id,
 *         refreshTokenTtlMinutes: 43200,
 *         ttlSeconds: 3600,
 *     },
 *     lambdaConfiguration: {
 *         accessTokenPopulateId: fusionauth_lambda.token_populate.id,
 *         idTokenPopulateId: fusionauth_lambda.id_token_populate.id,
 *     },
 *     loginConfiguration: {
 *         allowTokenRefresh: false,
 *         generateRefreshTokens: false,
 *         requireAuthentication: true,
 *     },
 *     multiFactorConfiguration: {
 *         emailTemplateId: "859f394b-22a6-4fa6-ba55-de700df9e950",
 *         smsTemplateId: "17760f96-dca7-448b-9a8f-c49016aa7210",
 *         loginPolicy: "Required",
 *         trustPolicy: "Any",
 *     },
 *     oauthConfiguration: {
 *         authorizedOriginUrls: ["http://www.example.com/oauth-callback"],
 *         authorizedUrlValidationPolicy: "ExactMatch",
 *         enabledGrants: [
 *             "authorization_code",
 *             "implicit",
 *         ],
 *         generateRefreshTokens: false,
 *         logoutBehavior: "AllApplications",
 *         logoutUrl: "http://www.example.com/logout",
 *         requireClientAuthentication: false,
 *         providedScopePolicies: [{
 *             address: {
 *                 enabled: false,
 *                 required: false,
 *             },
 *             email: {
 *                 enabled: false,
 *                 required: false,
 *             },
 *             phone: {
 *                 enabled: false,
 *                 required: false,
 *             },
 *             profile: {
 *                 enabled: false,
 *                 required: false,
 *             },
 *         }],
 *     },
 *     registrationConfiguration: {
 *         birthDate: {
 *             enabled: false,
 *             required: false,
 *         },
 *         confirmPassword: false,
 *         enabled: false,
 *         firstName: {
 *             enabled: false,
 *             required: false,
 *         },
 *         fullName: {
 *             enabled: false,
 *             required: false,
 *         },
 *         lastName: {
 *             enabled: false,
 *             required: false,
 *         },
 *         loginIdType: "",
 *         middleName: {
 *             enabled: false,
 *             required: false,
 *         },
 *         mobilePhone: {
 *             enabled: false,
 *             required: false,
 *         },
 *         preferredLanguages: {
 *             enabled: false,
 *             required: false,
 *         },
 *         type: "",
 *     },
 *     passwordlessConfigurationEnabled: false,
 *     registrationDeletePolicy: {
 *         unverifiedEnabled: true,
 *         unverifiedNumberOfDaysToRetain: 30,
 *     },
 * });
 * ```
 */
export class FusionAuthApplication extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthApplicationState, opts?: pulumi.CustomResourceOptions): FusionAuthApplication {
        return new FusionAuthApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthApplication:FusionAuthApplication';

    /**
     * Returns true if the given object is an instance of FusionAuthApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthApplication.__pulumiType;
    }

    public readonly accessControlConfiguration!: pulumi.Output<outputs.FusionAuthApplicationAccessControlConfiguration>;
    /**
     * The Id to use for the new Application. If not specified a secure random UUID will be generated.
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
     */
    public readonly authenticationTokenConfigurationEnabled!: pulumi.Output<boolean | undefined>;
    public readonly cleanSpeakConfiguration!: pulumi.Output<outputs.FusionAuthApplicationCleanSpeakConfiguration>;
    /**
     * An object that can hold any information about the Application that should be persisted.
     */
    public readonly data!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly emailConfiguration!: pulumi.Output<outputs.FusionAuthApplicationEmailConfiguration>;
    public readonly formConfiguration!: pulumi.Output<outputs.FusionAuthApplicationFormConfiguration | undefined>;
    /**
     * The instant that the Application was added to the FusionAuth database.
     */
    public /*out*/ readonly insertInstant!: pulumi.Output<number>;
    public readonly jwtConfiguration!: pulumi.Output<outputs.FusionAuthApplicationJwtConfiguration | undefined>;
    public readonly lambdaConfiguration!: pulumi.Output<outputs.FusionAuthApplicationLambdaConfiguration>;
    /**
     * The instant that the Application was last updated in the FusionAuth database.
     */
    public /*out*/ readonly lastUpdateInstant!: pulumi.Output<number>;
    public readonly loginConfiguration!: pulumi.Output<outputs.FusionAuthApplicationLoginConfiguration>;
    public readonly multiFactorConfiguration!: pulumi.Output<outputs.FusionAuthApplicationMultiFactorConfiguration>;
    /**
     * The name of the Application.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly oauthConfiguration!: pulumi.Output<outputs.FusionAuthApplicationOauthConfiguration>;
    /**
     * Determines if passwordless login is enabled for this application.
     */
    public readonly passwordlessConfigurationEnabled!: pulumi.Output<boolean | undefined>;
    public readonly registrationConfiguration!: pulumi.Output<outputs.FusionAuthApplicationRegistrationConfiguration>;
    public readonly registrationDeletePolicy!: pulumi.Output<outputs.FusionAuthApplicationRegistrationDeletePolicy>;
    public readonly samlv2Configuration!: pulumi.Output<outputs.FusionAuthApplicationSamlv2Configuration>;
    /**
     * The Id of the Tenant that this Application belongs to.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The unique Id of the theme to be used to style the login page and other end user templates.
     */
    public readonly themeId!: pulumi.Output<string | undefined>;
    /**
     * The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
     */
    public readonly verificationEmailTemplateId!: pulumi.Output<string | undefined>;
    /**
     * The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
     */
    public readonly verificationStrategy!: pulumi.Output<string | undefined>;
    /**
     * Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
     */
    public readonly verifyRegistration!: pulumi.Output<boolean | undefined>;
    public readonly webauthnConfiguration!: pulumi.Output<outputs.FusionAuthApplicationWebauthnConfiguration | undefined>;

    /**
     * Create a FusionAuthApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthApplicationArgs | FusionAuthApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthApplicationState | undefined;
            resourceInputs["accessControlConfiguration"] = state ? state.accessControlConfiguration : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["authenticationTokenConfigurationEnabled"] = state ? state.authenticationTokenConfigurationEnabled : undefined;
            resourceInputs["cleanSpeakConfiguration"] = state ? state.cleanSpeakConfiguration : undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["emailConfiguration"] = state ? state.emailConfiguration : undefined;
            resourceInputs["formConfiguration"] = state ? state.formConfiguration : undefined;
            resourceInputs["insertInstant"] = state ? state.insertInstant : undefined;
            resourceInputs["jwtConfiguration"] = state ? state.jwtConfiguration : undefined;
            resourceInputs["lambdaConfiguration"] = state ? state.lambdaConfiguration : undefined;
            resourceInputs["lastUpdateInstant"] = state ? state.lastUpdateInstant : undefined;
            resourceInputs["loginConfiguration"] = state ? state.loginConfiguration : undefined;
            resourceInputs["multiFactorConfiguration"] = state ? state.multiFactorConfiguration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oauthConfiguration"] = state ? state.oauthConfiguration : undefined;
            resourceInputs["passwordlessConfigurationEnabled"] = state ? state.passwordlessConfigurationEnabled : undefined;
            resourceInputs["registrationConfiguration"] = state ? state.registrationConfiguration : undefined;
            resourceInputs["registrationDeletePolicy"] = state ? state.registrationDeletePolicy : undefined;
            resourceInputs["samlv2Configuration"] = state ? state.samlv2Configuration : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["themeId"] = state ? state.themeId : undefined;
            resourceInputs["verificationEmailTemplateId"] = state ? state.verificationEmailTemplateId : undefined;
            resourceInputs["verificationStrategy"] = state ? state.verificationStrategy : undefined;
            resourceInputs["verifyRegistration"] = state ? state.verifyRegistration : undefined;
            resourceInputs["webauthnConfiguration"] = state ? state.webauthnConfiguration : undefined;
        } else {
            const args = argsOrState as FusionAuthApplicationArgs | undefined;
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["accessControlConfiguration"] = args ? args.accessControlConfiguration : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["authenticationTokenConfigurationEnabled"] = args ? args.authenticationTokenConfigurationEnabled : undefined;
            resourceInputs["cleanSpeakConfiguration"] = args ? args.cleanSpeakConfiguration : undefined;
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["emailConfiguration"] = args ? args.emailConfiguration : undefined;
            resourceInputs["formConfiguration"] = args ? args.formConfiguration : undefined;
            resourceInputs["jwtConfiguration"] = args ? args.jwtConfiguration : undefined;
            resourceInputs["lambdaConfiguration"] = args ? args.lambdaConfiguration : undefined;
            resourceInputs["loginConfiguration"] = args ? args.loginConfiguration : undefined;
            resourceInputs["multiFactorConfiguration"] = args ? args.multiFactorConfiguration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauthConfiguration"] = args ? args.oauthConfiguration : undefined;
            resourceInputs["passwordlessConfigurationEnabled"] = args ? args.passwordlessConfigurationEnabled : undefined;
            resourceInputs["registrationConfiguration"] = args ? args.registrationConfiguration : undefined;
            resourceInputs["registrationDeletePolicy"] = args ? args.registrationDeletePolicy : undefined;
            resourceInputs["samlv2Configuration"] = args ? args.samlv2Configuration : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["themeId"] = args ? args.themeId : undefined;
            resourceInputs["verificationEmailTemplateId"] = args ? args.verificationEmailTemplateId : undefined;
            resourceInputs["verificationStrategy"] = args ? args.verificationStrategy : undefined;
            resourceInputs["verifyRegistration"] = args ? args.verifyRegistration : undefined;
            resourceInputs["webauthnConfiguration"] = args ? args.webauthnConfiguration : undefined;
            resourceInputs["insertInstant"] = undefined /*out*/;
            resourceInputs["lastUpdateInstant"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthApplication resources.
 */
export interface FusionAuthApplicationState {
    accessControlConfiguration?: pulumi.Input<inputs.FusionAuthApplicationAccessControlConfiguration>;
    /**
     * The Id to use for the new Application. If not specified a secure random UUID will be generated.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
     */
    authenticationTokenConfigurationEnabled?: pulumi.Input<boolean>;
    cleanSpeakConfiguration?: pulumi.Input<inputs.FusionAuthApplicationCleanSpeakConfiguration>;
    /**
     * An object that can hold any information about the Application that should be persisted.
     */
    data?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    emailConfiguration?: pulumi.Input<inputs.FusionAuthApplicationEmailConfiguration>;
    formConfiguration?: pulumi.Input<inputs.FusionAuthApplicationFormConfiguration>;
    /**
     * The instant that the Application was added to the FusionAuth database.
     */
    insertInstant?: pulumi.Input<number>;
    jwtConfiguration?: pulumi.Input<inputs.FusionAuthApplicationJwtConfiguration>;
    lambdaConfiguration?: pulumi.Input<inputs.FusionAuthApplicationLambdaConfiguration>;
    /**
     * The instant that the Application was last updated in the FusionAuth database.
     */
    lastUpdateInstant?: pulumi.Input<number>;
    loginConfiguration?: pulumi.Input<inputs.FusionAuthApplicationLoginConfiguration>;
    multiFactorConfiguration?: pulumi.Input<inputs.FusionAuthApplicationMultiFactorConfiguration>;
    /**
     * The name of the Application.
     */
    name?: pulumi.Input<string>;
    oauthConfiguration?: pulumi.Input<inputs.FusionAuthApplicationOauthConfiguration>;
    /**
     * Determines if passwordless login is enabled for this application.
     */
    passwordlessConfigurationEnabled?: pulumi.Input<boolean>;
    registrationConfiguration?: pulumi.Input<inputs.FusionAuthApplicationRegistrationConfiguration>;
    registrationDeletePolicy?: pulumi.Input<inputs.FusionAuthApplicationRegistrationDeletePolicy>;
    samlv2Configuration?: pulumi.Input<inputs.FusionAuthApplicationSamlv2Configuration>;
    /**
     * The Id of the Tenant that this Application belongs to.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The unique Id of the theme to be used to style the login page and other end user templates.
     */
    themeId?: pulumi.Input<string>;
    /**
     * The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
     */
    verificationEmailTemplateId?: pulumi.Input<string>;
    /**
     * The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
     */
    verificationStrategy?: pulumi.Input<string>;
    /**
     * Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
     */
    verifyRegistration?: pulumi.Input<boolean>;
    webauthnConfiguration?: pulumi.Input<inputs.FusionAuthApplicationWebauthnConfiguration>;
}

/**
 * The set of arguments for constructing a FusionAuthApplication resource.
 */
export interface FusionAuthApplicationArgs {
    accessControlConfiguration?: pulumi.Input<inputs.FusionAuthApplicationAccessControlConfiguration>;
    /**
     * The Id to use for the new Application. If not specified a secure random UUID will be generated.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * Determines if Users can have Authentication Tokens associated with this Application. This feature may not be enabled for the FusionAuth application.
     */
    authenticationTokenConfigurationEnabled?: pulumi.Input<boolean>;
    cleanSpeakConfiguration?: pulumi.Input<inputs.FusionAuthApplicationCleanSpeakConfiguration>;
    /**
     * An object that can hold any information about the Application that should be persisted.
     */
    data?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    emailConfiguration?: pulumi.Input<inputs.FusionAuthApplicationEmailConfiguration>;
    formConfiguration?: pulumi.Input<inputs.FusionAuthApplicationFormConfiguration>;
    jwtConfiguration?: pulumi.Input<inputs.FusionAuthApplicationJwtConfiguration>;
    lambdaConfiguration?: pulumi.Input<inputs.FusionAuthApplicationLambdaConfiguration>;
    loginConfiguration?: pulumi.Input<inputs.FusionAuthApplicationLoginConfiguration>;
    multiFactorConfiguration?: pulumi.Input<inputs.FusionAuthApplicationMultiFactorConfiguration>;
    /**
     * The name of the Application.
     */
    name?: pulumi.Input<string>;
    oauthConfiguration?: pulumi.Input<inputs.FusionAuthApplicationOauthConfiguration>;
    /**
     * Determines if passwordless login is enabled for this application.
     */
    passwordlessConfigurationEnabled?: pulumi.Input<boolean>;
    registrationConfiguration?: pulumi.Input<inputs.FusionAuthApplicationRegistrationConfiguration>;
    registrationDeletePolicy?: pulumi.Input<inputs.FusionAuthApplicationRegistrationDeletePolicy>;
    samlv2Configuration?: pulumi.Input<inputs.FusionAuthApplicationSamlv2Configuration>;
    /**
     * The Id of the Tenant that this Application belongs to.
     */
    tenantId: pulumi.Input<string>;
    /**
     * The unique Id of the theme to be used to style the login page and other end user templates.
     */
    themeId?: pulumi.Input<string>;
    /**
     * The Id of the Email Template that is used to send the Registration Verification emails to users. If the verifyRegistration field is true this field is required.
     */
    verificationEmailTemplateId?: pulumi.Input<string>;
    /**
     * The process by which the user will verify their email address. Possible values are `ClickableLink` or `FormField`
     */
    verificationStrategy?: pulumi.Input<string>;
    /**
     * Whether or not registrations to this Application may be verified. When this is set to true the verificationEmailTemplateId parameter is also required.
     */
    verifyRegistration?: pulumi.Input<boolean>;
    webauthnConfiguration?: pulumi.Input<inputs.FusionAuthApplicationWebauthnConfiguration>;
}
