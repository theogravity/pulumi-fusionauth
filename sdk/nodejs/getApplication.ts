// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const fusionAuth = fusionauth.getApplication({
 *     name: "FusionAuth",
 * });
 * ```
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fusionauth:index/getApplication:getApplication", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    /**
     * The name of the Application.
     */
    name: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
/**
 * ## # Application Resource
 *
 * [Applications API](https://fusionauth.io/docs/v1/tech/apis/applications)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const fusionAuth = fusionauth.getApplication({
 *     name: "FusionAuth",
 * });
 * ```
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply((a: any) => getApplication(a, opts))
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationOutputArgs {
    /**
     * The name of the Application.
     */
    name: pulumi.Input<string>;
}
