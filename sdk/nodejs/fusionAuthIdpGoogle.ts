// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Google Identity Provider Resource
 *
 * The Google identity provider type will use the Google OAuth v2.0 login API. it will provide a Login with Google button on FusionAuth’s login page that will leverage the Google login pop-up dialog. Additionally, this identity provider will call Google’s /oauth2/v3/tokeninfo API to load additional details about the user and store them in FusionAuth.
 *
 * The email address returned by the Google Token info API will be used to create or lookup the existing user. Additional claims returned by Google can be used to reconcile the User to FusionAuth by using a Google Reconcile Lambda. Unless you assign a reconcile lambda to this provider, on the email address will be used from the available claims returned by Google.
 *
 * FusionAuth will also store the Google accessToken that is returned from the login pop-up dialog in the UserRegistration object inside the tokens Map. This Map stores the tokens from the various identity providers so that you can use them in your application to call their APIs.
 *
 * [Google Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/google#create-the-google-identity-provider)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const google = new fusionauth.FusionAuthIdpGoogle("google", {
 *     applicationConfigurations: [{
 *         applicationId: fusionauth_application.myapp.id,
 *         createRegistration: true,
 *         enabled: true,
 *     }],
 *     buttonText: "Login with Google",
 *     debug: false,
 *     clientId: "254311943570-8e2i2hds0qdnee4124socceeh2q2mtjl.apps.googleusercontent.com",
 *     clientSecret: "BRr7x7xz_-cXxIFznBDIdxF1",
 *     scope: "profile",
 * });
 * ```
 */
export class FusionAuthIdpGoogle extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthIdpGoogle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthIdpGoogleState, opts?: pulumi.CustomResourceOptions): FusionAuthIdpGoogle {
        return new FusionAuthIdpGoogle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthIdpGoogle:FusionAuthIdpGoogle';

    /**
     * Returns true if the given object is an instance of FusionAuthIdpGoogle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthIdpGoogle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthIdpGoogle.__pulumiType;
    }

    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    public readonly applicationConfigurations!: pulumi.Output<outputs.FusionAuthIdpGoogleApplicationConfiguration[] | undefined>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    public readonly buttonText!: pulumi.Output<string>;
    /**
     * The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
     */
    public readonly debug!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if this provider is enabled. If it is false then it will be disabled globally.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique Id of the lambda to used during the user reconcile process to map custom claims from the external identity provider to the FusionAuth user.
     */
    public readonly lambdaReconcileId!: pulumi.Output<string | undefined>;
    /**
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    public readonly linkingStrategy!: pulumi.Output<string>;
    /**
     * The login method to use for this Identity Provider.
     */
    public readonly loginMethod!: pulumi.Output<string | undefined>;
    /**
     * An object to hold configuration parameters for the Google Identity Services API.
     */
    public readonly properties!: pulumi.Output<outputs.FusionAuthIdpGoogleProperties | undefined>;
    /**
     * The top-level scope that you are requesting from Google.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    public readonly tenantConfigurations!: pulumi.Output<outputs.FusionAuthIdpGoogleTenantConfiguration[] | undefined>;

    /**
     * Create a FusionAuthIdpGoogle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthIdpGoogleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthIdpGoogleArgs | FusionAuthIdpGoogleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthIdpGoogleState | undefined;
            resourceInputs["applicationConfigurations"] = state ? state.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = state ? state.buttonText : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["debug"] = state ? state.debug : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["lambdaReconcileId"] = state ? state.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = state ? state.linkingStrategy : undefined;
            resourceInputs["loginMethod"] = state ? state.loginMethod : undefined;
            resourceInputs["properties"] = state ? state.properties : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["tenantConfigurations"] = state ? state.tenantConfigurations : undefined;
        } else {
            const args = argsOrState as FusionAuthIdpGoogleArgs | undefined;
            if ((!args || args.buttonText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buttonText'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            resourceInputs["applicationConfigurations"] = args ? args.applicationConfigurations : undefined;
            resourceInputs["buttonText"] = args ? args.buttonText : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["debug"] = args ? args.debug : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["lambdaReconcileId"] = args ? args.lambdaReconcileId : undefined;
            resourceInputs["linkingStrategy"] = args ? args.linkingStrategy : undefined;
            resourceInputs["loginMethod"] = args ? args.loginMethod : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tenantConfigurations"] = args ? args.tenantConfigurations : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FusionAuthIdpGoogle.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthIdpGoogle resources.
 */
export interface FusionAuthIdpGoogleState {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpGoogleApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText?: pulumi.Input<string>;
    /**
     * The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
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
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The login method to use for this Identity Provider.
     */
    loginMethod?: pulumi.Input<string>;
    /**
     * An object to hold configuration parameters for the Google Identity Services API.
     */
    properties?: pulumi.Input<inputs.FusionAuthIdpGoogleProperties>;
    /**
     * The top-level scope that you are requesting from Google.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpGoogleTenantConfiguration>[]>;
}

/**
 * The set of arguments for constructing a FusionAuthIdpGoogle resource.
 */
export interface FusionAuthIdpGoogleArgs {
    /**
     * The configuration for each Application that the identity provider is enabled for.
     */
    applicationConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpGoogleApplicationConfiguration>[]>;
    /**
     * The top-level button text to use on the FusionAuth login page for this Identity Provider.
     */
    buttonText: pulumi.Input<string>;
    /**
     * The top-level Google client id for your Application. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    clientId: pulumi.Input<string>;
    /**
     * The top-level client secret to use with the Google Identity Provider when retrieving the long-lived token. This value is retrieved from the Google developer website when you setup your Google developer account.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login an Event Log will be created.
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
     * The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
     */
    linkingStrategy?: pulumi.Input<string>;
    /**
     * The login method to use for this Identity Provider.
     */
    loginMethod?: pulumi.Input<string>;
    /**
     * An object to hold configuration parameters for the Google Identity Services API.
     */
    properties?: pulumi.Input<inputs.FusionAuthIdpGoogleProperties>;
    /**
     * The top-level scope that you are requesting from Google.
     */
    scope?: pulumi.Input<string>;
    /**
     * The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
     */
    tenantConfigurations?: pulumi.Input<pulumi.Input<inputs.FusionAuthIdpGoogleTenantConfiguration>[]>;
}
