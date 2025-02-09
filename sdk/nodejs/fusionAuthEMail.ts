// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Email Resource
 *
 * This resource contains the APIs for managing Email Templates.
 *
 * [Emails API](https://fusionauth.io/docs/v1/tech/apis/emails)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const helloWorld = new fusionauth.FusionAuthEMail("helloWorld", {
 *     defaultFromName: "Welcome Team",
 *     defaultHtmlTemplate: fs.readFileSync(`${path.module}/email_templates/HelloWorld.html.ftl`, "utf8"),
 *     defaultSubject: "Hello",
 *     defaultTextTemplate: fs.readFileSync(`${path.module}/email_templates/HelloWorld.txt.ftl`, "utf8"),
 *     fromEmail: "welcome@example.com.com",
 * });
 * ```
 */
export class FusionAuthEMail extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthEMail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthEMailState, opts?: pulumi.CustomResourceOptions): FusionAuthEMail {
        return new FusionAuthEMail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthEMail:FusionAuthEMail';

    /**
     * Returns true if the given object is an instance of FusionAuthEMail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthEMail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthEMail.__pulumiType;
    }

    /**
     * The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
     */
    public readonly defaultFromName!: pulumi.Output<string | undefined>;
    /**
     * The default HTML Email Template.
     */
    public readonly defaultHtmlTemplate!: pulumi.Output<string>;
    /**
     * The default Subject used when sending emails.
     */
    public readonly defaultSubject!: pulumi.Output<string>;
    /**
     * The default Text Email Template.
     */
    public readonly defaultTextTemplate!: pulumi.Output<string>;
    /**
     * The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
     */
    public readonly emailId!: pulumi.Output<string | undefined>;
    /**
     * The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
     */
    public readonly fromEmail!: pulumi.Output<string | undefined>;
    /**
     * The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
     */
    public readonly localizedFromNames!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
     */
    public readonly localizedHtmlTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
     */
    public readonly localizedSubjects!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
     */
    public readonly localizedTextTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A descriptive name for the email template (i.e. "April 2016 Coupon Email")
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FusionAuthEMail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthEMailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthEMailArgs | FusionAuthEMailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthEMailState | undefined;
            resourceInputs["defaultFromName"] = state ? state.defaultFromName : undefined;
            resourceInputs["defaultHtmlTemplate"] = state ? state.defaultHtmlTemplate : undefined;
            resourceInputs["defaultSubject"] = state ? state.defaultSubject : undefined;
            resourceInputs["defaultTextTemplate"] = state ? state.defaultTextTemplate : undefined;
            resourceInputs["emailId"] = state ? state.emailId : undefined;
            resourceInputs["fromEmail"] = state ? state.fromEmail : undefined;
            resourceInputs["localizedFromNames"] = state ? state.localizedFromNames : undefined;
            resourceInputs["localizedHtmlTemplates"] = state ? state.localizedHtmlTemplates : undefined;
            resourceInputs["localizedSubjects"] = state ? state.localizedSubjects : undefined;
            resourceInputs["localizedTextTemplates"] = state ? state.localizedTextTemplates : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FusionAuthEMailArgs | undefined;
            if ((!args || args.defaultHtmlTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultHtmlTemplate'");
            }
            if ((!args || args.defaultSubject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultSubject'");
            }
            if ((!args || args.defaultTextTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultTextTemplate'");
            }
            resourceInputs["defaultFromName"] = args ? args.defaultFromName : undefined;
            resourceInputs["defaultHtmlTemplate"] = args ? args.defaultHtmlTemplate : undefined;
            resourceInputs["defaultSubject"] = args ? args.defaultSubject : undefined;
            resourceInputs["defaultTextTemplate"] = args ? args.defaultTextTemplate : undefined;
            resourceInputs["emailId"] = args ? args.emailId : undefined;
            resourceInputs["fromEmail"] = args ? args.fromEmail : undefined;
            resourceInputs["localizedFromNames"] = args ? args.localizedFromNames : undefined;
            resourceInputs["localizedHtmlTemplates"] = args ? args.localizedHtmlTemplates : undefined;
            resourceInputs["localizedSubjects"] = args ? args.localizedSubjects : undefined;
            resourceInputs["localizedTextTemplates"] = args ? args.localizedTextTemplates : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FusionAuthEMail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthEMail resources.
 */
export interface FusionAuthEMailState {
    /**
     * The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
     */
    defaultFromName?: pulumi.Input<string>;
    /**
     * The default HTML Email Template.
     */
    defaultHtmlTemplate?: pulumi.Input<string>;
    /**
     * The default Subject used when sending emails.
     */
    defaultSubject?: pulumi.Input<string>;
    /**
     * The default Text Email Template.
     */
    defaultTextTemplate?: pulumi.Input<string>;
    /**
     * The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
     */
    emailId?: pulumi.Input<string>;
    /**
     * The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
     */
    fromEmail?: pulumi.Input<string>;
    /**
     * The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
     */
    localizedFromNames?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
     */
    localizedHtmlTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
     */
    localizedSubjects?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
     */
    localizedTextTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A descriptive name for the email template (i.e. "April 2016 Coupon Email")
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthEMail resource.
 */
export interface FusionAuthEMailArgs {
    /**
     * The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
     */
    defaultFromName?: pulumi.Input<string>;
    /**
     * The default HTML Email Template.
     */
    defaultHtmlTemplate: pulumi.Input<string>;
    /**
     * The default Subject used when sending emails.
     */
    defaultSubject: pulumi.Input<string>;
    /**
     * The default Text Email Template.
     */
    defaultTextTemplate: pulumi.Input<string>;
    /**
     * The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
     */
    emailId?: pulumi.Input<string>;
    /**
     * The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
     */
    fromEmail?: pulumi.Input<string>;
    /**
     * The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
     */
    localizedFromNames?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
     */
    localizedHtmlTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
     */
    localizedSubjects?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
     */
    localizedTextTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A descriptive name for the email template (i.e. "April 2016 Coupon Email")
     */
    name?: pulumi.Input<string>;
}
