// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # LinkedIn Identity Provider Resource
 *
 * The LinkedIn identity provider type will use OAuth 2.0 to authenticate a user with LinkedIn. It will also provide a
 * `Login with LinkedIn` button on FusionAuth’s login page that will direct a user to the LinkedIn login page.
 * Additionally, after successful user authentication, this identity provider will call LinkedIn’s `/v2/me` and
 * `/v2/emailAddress` APIs to load additional details about the user and store them in FusionAuth.
 *
 * The email address returned by the LinkedIn `/v2/emailAddress` API will be used to create or look up the existing user.
 * Additional claims returned by LinkedIn can be used to reconcile the User to FusionAuth by using a LinkedIn Reconcile
 * lambda. Unless you assign a reconcile lambda to this provider, only the email address will be used from the available
 * claims returned by LinkedIn.
 *
 * FusionAuth will also store the LinkedIn `accessToken` returned from the login endpoint in the `identityProviderLink`
 * object. This object is accessible using the Link API.
 *
 * The `identityProviderLink` object stores the token so that you can use it in your application to call LinkedIn APIs on
 * behalf of the user if desired.
 *
 * [LinkedIn Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/linkedin)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fusionauth from "pulumi-fusionauth";
 *
 * const linkedin = new fusionauth.FusionAuthIdpLinkedIn("linkedin", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.myapp.id,
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with LinkedIn",
 *     debug: false,
 *     enabled: true,
 *     clientId: "9876543210",
 *     clientSecret: "716a572f917640698cdb99e9d7e64115",
 *     scope: "r_emailaddress r_liteprofile",
 * });
 * ```
 */
export class FusionAuthIdpLinkedIn extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpLinkedIn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpLinkedInState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpLinkedIn {
        return new FusionAuthIdpLinkedIn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpLinkedIn:FusionAuthIdpLinkedIn';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpLinkedIn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpLinkedIn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpLinkedIn.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpLinkedInApplicationConfiguration[] | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    public readonly lambdaReconcileId!: pulumi.Output<string>;
    /**
     * The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
     * The valid values are:
     */
    public readonly linkingStrategy!: pulumi.Output<string>;
    /**
     * The top-level scope that you are requesting from LinkedIn.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpLinkedInTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpLinkedIn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpLinkedInArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpLinkedInArgs | FusionAuthIdpLinkedInState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpLinkedInState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpLinkedInArgs | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthIdpLinkedIn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpLinkedIn resources.
 */
export interface FusionAuthIdpLinkedInState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpLinkedInApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
     * The valid values are:
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from LinkedIn.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpLinkedInTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpLinkedIn resource.
 */
export interface FusionAuthIdpLinkedInArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpLinkedInApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * The top-level LinkedIn client id for your Application. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    clientId: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the LinkedIn Identity Provider when retrieving the long-lived token. This value is retrieved from the LinkedIn developer website when you set up your LinkedIn app.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, an Event Log is created each time this provider is invoked to reconcile a login.
     */
    debug?: pulumi.Input<boolean>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    lambdaReconcileId?: pulumi.Input<string>;
    /**
     * The linking strategy to use when creating the link between the Facebook Identity Provider and the user.
     * The valid values are:
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The top-level scope that you are requesting from LinkedIn.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpLinkedInTenantConfiguration>[]>;
}