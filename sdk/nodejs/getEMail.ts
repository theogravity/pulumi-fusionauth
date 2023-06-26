// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Email Resource
 *
 * This data source is used to fetch information about a specific Email Template.
 *
 * [Emails API](https://fusionauth.io/docs/v1/tech/apis/emails)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const defaultBreachedPassword = fusionauth.getEMail({
 *     name: "[FusionAuth Default] Breached Password Notification",
 * });
 * ```
 */
export function getEMail(args: GetEMailArgs, opts?: pulumi.InvokeOptions): Promise<GetEMailResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fusionauth:index/getEMail:getEMail", {
        "fromEmail": args.fromEmail,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getEMail.
 */
export interface GetEMailArgs {
    /**
     * The email address that this email will be sent from.
     */
    fromEmail?: string;
    /**
     * The name of the Email Template.
     */
    name: string;
}

/**
 * A collection of values returned by getEMail.
 */
export interface GetEMailResult {
    /**
     * The default From Name used when sending emails.
     */
    readonly defaultFromName: string;
    /**
     * The default HTML Email Template.
     */
    readonly defaultHtmlTemplate: string;
    /**
     * The default Subject used when sending emails.
     */
    readonly defaultSubject: string;
    /**
     * The default Text Email Template.
     */
    readonly defaultTextTemplate: string;
    /**
     * The email address that this email will be sent from.
     */
    readonly fromEmail?: string;
    /**
     * The Id of the Email Template.
     */
    readonly id: string;
    /**
     * The From Name used when sending emails to users who speak other languages.
     */
    readonly localizedFromNames: {[key: string]: any};
    /**
     * The HTML Email Template used when sending emails to users who speak other languages.
     */
    readonly localizedHtmlTemplates: {[key: string]: any};
    /**
     * The Subject used when sending emails to users who speak other languages.
     */
    readonly localizedSubjects: {[key: string]: any};
    /**
     * The Text Email Template used when sending emails to users who speak other languages.
     */
    readonly localizedTextTemplates: {[key: string]: any};
    readonly name: string;
}
/**
 * ## # Email Resource
 *
 * This data source is used to fetch information about a specific Email Template.
 *
 * [Emails API](https://fusionauth.io/docs/v1/tech/apis/emails)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "@pulumi/fusionauth";
 *
 * const defaultBreachedPassword = fusionauth.getEMail({
 *     name: "[FusionAuth Default] Breached Password Notification",
 * });
 * ```
 */
export function getEMailOutput(args: GetEMailOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEMailResult> {
    return pulumi.output(args).apply((a: any) => getEMail(a, opts))
}

/**
 * A collection of arguments for invoking getEMail.
 */
export interface GetEMailOutputArgs {
    /**
     * The email address that this email will be sent from.
     */
    fromEmail?: pulumi.Input<string>;
    /**
     * The name of the Email Template.
     */
    name: pulumi.Input<string>;
}