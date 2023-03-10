// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Facebook Identity Provider Resource
 *
 * The Facebook identity provider type will use the Facebook OAuth login API. It will provide a `Login with Facebook` button on FusionAuth’s login page that will leverage the Facebook login pop-up dialog. Additionally, this identity provider will call Facebook’s Graph API to load additional details about the user and store them in FusionAuth.
 *
 * The email address returned by the Facebook Graph API will be used to create or lookup the existing user. Additional claims returned by Facebook can be used to reconcile the User to FusionAuth by using a Facebook Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the `email` address will be used from the available claims returned by Facebook.
 *
 * When the `picture` field is not requested FusionAuth will also call Facebook’s `/me/picture` API to load the user’s profile image and store it as the `imageUrl` in FusionAuth. When the `picture` field is requested, the user’s profile image will be returned by the `/me` API and a second request to the `/me/picture` endpoint will not be required.
 *
 * Finally, FusionAuth will call Facebook’s `/oauth/access_token` API to exchange the login token for a long-lived Facebook token. This token is stored in the `UserRegistration` object inside the `tokens` Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
 *
 * Please note if an `idpHint` is appended to the OAuth Authorize endpoint, then the interaction behavior will be defaulted to `redirect`, even if popup interaction is explicitly configured.
 *
 * [Facebook Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/facebook)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const facebook = new fusionauth.FusionAuthIdpFacebook("facebook", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.myapp.id,
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with Facebook",
 *     debug: false,
 *     enabled: true,
 *     appId: "9876543210",
 *     clientSecret: "716a572f917640698cdb99e9d7e64115",
 *     fields: "email",
 *     permissions: "email",
 * });
 * ```
 */
export class FusionAuthIdpFacebook extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpFacebook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpFacebookState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpFacebook {
        return new FusionAuthIdpFacebook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpFacebook:FusionAuthIdpFacebook';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpFacebook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpFacebook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpFacebook.__pulumiType;
    }

    /**
     * The top-level Facebook `appId` for your Application. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpFacebookApplicationConfiguration[] | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * The top-level client secret, also known as 'App Secret', to use with the Facebook Identity Provider when retrieving the long-lived token. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
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
     * The top-level fields that you are requesting from Facebook.
     * Field values are documented at [Facebook Graph API](https://developers.facebook.com/docs/graph-api/using-graph-api/)
     */
    public readonly fields!: pulumi.Output<string | undefined>;
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
     * The login method to use for this Identity Provider.
     * The valid values are:
     */
    public readonly loginMethod!: pulumi.Output<string | undefined>;
    /**
     * The top-level permissions that your application is asking of the user’s Facebook account.
     * Permission values are documented at [Facebook Login API](https://developers.facebook.com/docs/permissions/reference)
     */
    public readonly permissions!: pulumi.Output<string | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpFacebookTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpFacebook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpFacebookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpFacebookArgs | FusionAuthIdpFacebookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpFacebookState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["fields"] = state ? state.fields : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["loginMethod"] = state ? state.loginMethod : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpFacebookArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["loginMethod"] = args ? args.loginMethod : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FusionAuthIdpFacebook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpFacebook resources.
 */
export interface FusionAuthIdpFacebookState {
    /**
     * The top-level Facebook `appId` for your Application. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
     */
    appId?: pulumi.Input<string>;
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpFacebookApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * The top-level client secret, also known as 'App Secret', to use with the Facebook Identity Provider when retrieving the long-lived token. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
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
     * The top-level fields that you are requesting from Facebook.
     * Field values are documented at [Facebook Graph API](https://developers.facebook.com/docs/graph-api/using-graph-api/)
     */
    fields?: pulumi.Input<string>;
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
     * The login method to use for this Identity Provider.
     * The valid values are:
     */
    loginMethod?: pulumi.Input<string>;
    /**
     * The top-level permissions that your application is asking of the user’s Facebook account.
     * Permission values are documented at [Facebook Login API](https://developers.facebook.com/docs/permissions/reference)
     */
    permissions?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpFacebookTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpFacebook resource.
 */
export interface FusionAuthIdpFacebookArgs {
    /**
     * The top-level Facebook `appId` for your Application. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
     */
    appId: pulumi.Input<string>;
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpFacebookApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * The top-level client secret, also known as 'App Secret', to use with the Facebook Identity Provider when retrieving the long-lived token. This value is retrieved from the Facebook developer website when you setup your Facebook developer account.
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
     * The top-level fields that you are requesting from Facebook.
     * Field values are documented at [Facebook Graph API](https://developers.facebook.com/docs/graph-api/using-graph-api/)
     */
    fields?: pulumi.Input<string>;
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
     * The login method to use for this Identity Provider.
     * The valid values are:
     */
    loginMethod?: pulumi.Input<string>;
    /**
     * The top-level permissions that your application is asking of the user’s Facebook account.
     * Permission values are documented at [Facebook Login API](https://developers.facebook.com/docs/permissions/reference)
     */
    permissions?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpFacebookTenantConfiguration>[]>;
}
