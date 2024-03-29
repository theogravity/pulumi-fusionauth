// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Application Role Resource
 *
 * This Resource is used to create a role for an Application.
 *
 * [Application Roles API](https://fusionauth.io/docs/v1/tech/apis/applications)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const myAppAdminRole = new fusionauth.FusionAuthApplicationRole("myAppAdminRole", {
 *     applicationId: fusionauth_application.my_app.id,
 *     description: "",
 *     isDefault: false,
 *     isSuperRole: true,
 * });
 * ```
 */
export class FusionAuthApplicationRole extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthApplicationRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthApplicationRoleState, opts?: pulumi.CustomResourceOptions): FusionAuthApplicationRole {
        return new FusionAuthApplicationRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthApplicationRole:FusionAuthApplicationRole';

    /**
     * Returns true if the given object is an instance of FusionAuthApplicationRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthApplicationRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthApplicationRole.__pulumiType;
    }

    /**
     * ID of the application that this role is for.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * A description for the role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether or not the Role is a default role. A default role is automatically assigned to a user during registration if no roles are provided.
     */
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * Whether or not the Role is a considered to be a super user role. This is a marker to indicate that it supersedes all other roles. FusionAuth will attempt to enforce this contract when using the web UI, it is not enforced programmatically when using the API.
     */
    public readonly isSuperRole!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Role.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FusionAuthApplicationRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthApplicationRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthApplicationRoleArgs | FusionAuthApplicationRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthApplicationRoleState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["isSuperRole"] = state ? state.isSuperRole : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FusionAuthApplicationRoleArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["isSuperRole"] = args ? args.isSuperRole : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthApplicationRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthApplicationRole resources.
 */
export interface FusionAuthApplicationRoleState {
    /**
     * ID of the application that this role is for.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * A description for the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not the Role is a default role. A default role is automatically assigned to a user during registration if no roles are provided.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Whether or not the Role is a considered to be a super user role. This is a marker to indicate that it supersedes all other roles. FusionAuth will attempt to enforce this contract when using the web UI, it is not enforced programmatically when using the API.
     */
    isSuperRole?: pulumi.Input<boolean>;
    /**
     * The name of the Role.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthApplicationRole resource.
 */
export interface FusionAuthApplicationRoleArgs {
    /**
     * ID of the application that this role is for.
     */
    applicationId: pulumi.Input<string>;
    /**
     * A description for the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not the Role is a default role. A default role is automatically assigned to a user during registration if no roles are provided.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * Whether or not the Role is a considered to be a super user role. This is a marker to indicate that it supersedes all other roles. FusionAuth will attempt to enforce this contract when using the web UI, it is not enforced programmatically when using the API.
     */
    isSuperRole?: pulumi.Input<boolean>;
    /**
     * The name of the Role.
     */
    name?: pulumi.Input<string>;
}
