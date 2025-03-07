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
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const admin = fusionauth.getApplicationRole({
 *     applicationId: data.fusionauth_application.FusionAuth.id,
 *     name: "admin",
 * });
 * ```
 */
export function getApplicationRole(args: GetApplicationRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fusionauth:index/getApplicationRole:getApplicationRole", {
        "applicationId": args.applicationId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationRole.
 */
export interface GetApplicationRoleArgs {
    /**
     * ID of the application that this role is for.
     */
    applicationId: string;
    /**
     * The name of the Role.
     */
    name: string;
}

/**
 * A collection of values returned by getApplicationRole.
 */
export interface GetApplicationRoleResult {
    readonly applicationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
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
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const admin = fusionauth.getApplicationRole({
 *     applicationId: data.fusionauth_application.FusionAuth.id,
 *     name: "admin",
 * });
 * ```
 */
export function getApplicationRoleOutput(args: GetApplicationRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("fusionauth:index/getApplicationRole:getApplicationRole", {
        "applicationId": args.applicationId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationRole.
 */
export interface GetApplicationRoleOutputArgs {
    /**
     * ID of the application that this role is for.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The name of the Role.
     */
    name: pulumi.Input<string>;
}
