// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # User Group Membership Resource
 *
 * [User Group Membership API](https://fusionauth.io/docs/apis/groups#request-5)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const _this = new fusionauth.FusionAuthUserGroupMembership("this", {
 *     groupId: fusionauth_group["this"].id,
 *     userId: fusionauth_user["this"].id,
 * });
 * ```
 */
export class FusionAuthUserGroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthUserGroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthUserGroupMembershipState, opts?: pulumi.CustomResourceOptions): FusionAuthUserGroupMembership {
        return new FusionAuthUserGroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthUserGroupMembership:FusionAuthUserGroupMembership';

    /**
     * Returns true if the given object is an instance of FusionAuthUserGroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthUserGroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthUserGroupMembership.__pulumiType;
    }

    /**
     * An object that can hold any information about the User for this membership that should be persisted.
     */
    public readonly data!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Id of the Group of this membership.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The Id of the User Group Membership. If not provided, a random UUID will be generated.
     */
    public readonly membershipId!: pulumi.Output<string>;
    /**
     * "The Id of the User of this membership.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a FusionAuthUserGroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthUserGroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthUserGroupMembershipArgs | FusionAuthUserGroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthUserGroupMembershipState | undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["membershipId"] = state ? state.membershipId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as FusionAuthUserGroupMembershipArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["membershipId"] = args ? args.membershipId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "fusionauth:index/userGroupMembership:UserGroupMembership" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(FusionAuthUserGroupMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthUserGroupMembership resources.
 */
export interface FusionAuthUserGroupMembershipState {
    /**
     * An object that can hold any information about the User for this membership that should be persisted.
     */
    data?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Id of the Group of this membership.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The Id of the User Group Membership. If not provided, a random UUID will be generated.
     */
    membershipId?: pulumi.Input<string>;
    /**
     * "The Id of the User of this membership.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthUserGroupMembership resource.
 */
export interface FusionAuthUserGroupMembershipArgs {
    /**
     * An object that can hold any information about the User for this membership that should be persisted.
     */
    data?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Id of the Group of this membership.
     */
    groupId: pulumi.Input<string>;
    /**
     * The Id of the User Group Membership. If not provided, a random UUID will be generated.
     */
    membershipId?: pulumi.Input<string>;
    /**
     * "The Id of the User of this membership.
     */
    userId: pulumi.Input<string>;
}
