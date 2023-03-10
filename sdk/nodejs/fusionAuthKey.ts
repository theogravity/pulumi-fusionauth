// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Key Resource
 *
 * Cryptographic keys are used in signing and verifying JWTs and verifying responses for third party identity providers. It is more likely you will interact with keys using the FusionAuth UI in the Key Master menu.
 *
 * [Keys API](https://fusionauth.io/docs/v1/tech/apis/keys)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const adminId = new fusionauth.FusionAuthKey("adminId", {
 *     algorithm: "RS256",
 *     length: 2048,
 * });
 * ```
 */
export class FusionAuthKey extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthKeyState, opts?: pulumi.CustomResourceOptions): FusionAuthKey {
        return new FusionAuthKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthKey:FusionAuthKey';

    /**
     * Returns true if the given object is an instance of FusionAuthKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthKey.__pulumiType;
    }

    /**
     * The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
     */
    public readonly algorithm!: pulumi.Output<string>;
    /**
     * The Id to use for the new key. If not specified a secure random UUID will be generated.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The id used in the JWT header to identify the key used to generate the signature
     */
    public /*out*/ readonly kid!: pulumi.Output<string>;
    /**
     * The length of the RSA or EC certificate. This field is required when generating RSA key types.
     */
    public readonly length!: pulumi.Output<number | undefined>;
    /**
     * The name of the Key.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FusionAuthKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthKeyArgs | FusionAuthKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthKeyState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["kid"] = state ? state.kid : undefined;
            resourceInputs["length"] = state ? state.length : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FusionAuthKeyArgs | undefined;
            if ((!args || args.algorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'algorithm'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["length"] = args ? args.length : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["kid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthKey resources.
 */
export interface FusionAuthKeyState {
    /**
     * The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The Id to use for the new key. If not specified a secure random UUID will be generated.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The id used in the JWT header to identify the key used to generate the signature
     */
    kid?: pulumi.Input<string>;
    /**
     * The length of the RSA or EC certificate. This field is required when generating RSA key types.
     */
    length?: pulumi.Input<number>;
    /**
     * The name of the Key.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthKey resource.
 */
export interface FusionAuthKeyArgs {
    /**
     * The algorithm used to encrypt the Key. The following values represent algorithms supported by FusionAuth:
     */
    algorithm: pulumi.Input<string>;
    /**
     * The Id to use for the new key. If not specified a secure random UUID will be generated.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The length of the RSA or EC certificate. This field is required when generating RSA key types.
     */
    length?: pulumi.Input<number>;
    /**
     * The name of the Key.
     */
    name?: pulumi.Input<string>;
}
