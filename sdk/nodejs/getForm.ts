// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Form Resource
 *
 * A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.
 *
 * [Forms API](https://fusionauth.io/docs/v1/tech/apis/forms)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const default = fusionauth.getForm({
 *     name: "Default User Self Service provided by FusionAuth",
 * });
 * ```
 */
export function getForm(args?: GetFormArgs, opts?: pulumi.InvokeOptions): Promise<GetFormResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fusionauth:index/getForm:getForm", {
        "data": args.data,
        "formId": args.formId,
        "name": args.name,
        "steps": args.steps,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getForm.
 */
export interface GetFormArgs {
    /**
     * An object that can hold any information about the Form that should be persisted.
     */
    data?: {[key: string]: any};
    /**
     * The unique id of the Form. Either `formId` or `name` must be specified.
     */
    formId?: string;
    /**
     * The name of the Form. Either `formId` or `name` must be specified.
     */
    name?: string;
    /**
     * An ordered list of objects containing one or more Form Fields.
     */
    steps?: inputs.GetFormStep[];
    /**
     * The form type. The possible values are:
     */
    type?: string;
}

/**
 * A collection of values returned by getForm.
 */
export interface GetFormResult {
    /**
     * An object that can hold any information about the Form that should be persisted.
     */
    readonly data?: {[key: string]: any};
    readonly formId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The unique name of the Form.
     */
    readonly name?: string;
    /**
     * An ordered list of objects containing one or more Form Fields.
     */
    readonly steps?: outputs.GetFormStep[];
    /**
     * The form type. The possible values are:
     */
    readonly type?: string;
}
/**
 * ## # Form Resource
 *
 * A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.
 *
 * [Forms API](https://fusionauth.io/docs/v1/tech/apis/forms)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const default = fusionauth.getForm({
 *     name: "Default User Self Service provided by FusionAuth",
 * });
 * ```
 */
export function getFormOutput(args?: GetFormOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormResult> {
    return pulumi.output(args).apply((a: any) => getForm(a, opts))
}

/**
 * A collection of arguments for invoking getForm.
 */
export interface GetFormOutputArgs {
    /**
     * An object that can hold any information about the Form that should be persisted.
     */
    data?: pulumi.Input<{[key: string]: any}>;
    /**
     * The unique id of the Form. Either `formId` or `name` must be specified.
     */
    formId?: pulumi.Input<string>;
    /**
     * The name of the Form. Either `formId` or `name` must be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * An ordered list of objects containing one or more Form Fields.
     */
    steps?: pulumi.Input<pulumi.Input<inputs.GetFormStepArgs>[]>;
    /**
     * The form type. The possible values are:
     */
    type?: pulumi.Input<string>;
}
