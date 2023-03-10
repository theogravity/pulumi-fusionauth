// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Webhook Resource
 *
 * A FusionAuth Webhook is intended to consume JSON events emitted by FusionAuth. Creating a Webhook allows you to tell FusionAuth where you would like to receive these JSON events.
 *
 * [Webhooks API](https://fusionauth.io/docs/v1/tech/apis/webhooks)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fusionauth from "pulumi-fusionauth";
 *
 * const example = new fusionauth.FusionAuthWebhook("example", {
 *     tenantIds: [
 *         "00000000-0000-0000-0000-000000000003",
 *         fusionauth_tenant.example.id,
 *     ],
 *     connectTimeout: 1000,
 *     description: "The standard game Webhook",
 *     eventsEnabled: {
 *         userCreate: true,
 *         userDelete: false,
 *     },
 *     global: false,
 *     headers: {
 *         foo: "bar",
 *         bar: "baz",
 *     },
 *     httpAuthenticationPassword: "password",
 *     httpAuthenticationUsername: "username",
 *     readTimeout: 2000,
 *     sslCertificate: "  -----BEGIN CERTIFICATE-----\\nMIIDUjCCArugAwIBAgIJANZCTNN98L9ZMA0GCSqGSIb3DQEBBQUAMHoxCzAJBgNV\\nBAYTAlVTMQswCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVyMQ8wDQYDVQQKEwZz\\nZXRoLXMxCjAIBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAwHgYJKoZIhvcNAQkB\\nFhFzamZkZkBsc2tkamZjLmNvbTAeFw0xNDA0MDkyMTA2MDdaFw0xNDA1MDkyMTA2\\nMDdaMHoxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVy\\nMQ8wDQYDVQQKEwZzZXRoLXMxCjAIBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAw\\nHgYJKoZIhvcNAQkBFhFzamZkZkBsc2tkamZjLmNvbTCBnzANBgkqhkiG9w0BAQEF\\nAAOBjQAwgYkCgYEAxnQBqyuYvjUE4aFQ6vVZU5RqHmy3KiTg2NcxELIlZztUTK3a\\nVFbJoBB4ixHXCCYslujthILyBjgT3F+IhSpPAcrlu8O5LVPaPCysh/SNrGNwH4lq\\neiW9Z5WAhRO/nG7NZNa0USPHAei6b9Sv9PxuKCY+GJfAIwlO4/bltIH06/kCAwEA\\nAaOB3zCB3DAdBgNVHQ4EFgQUU4SqJEFm1zW+CcLxmLlARrqtMN0wgawGA1UdIwSB\\npDCBoYAUU4SqJEFm1zW+CcLxmLlARrqtMN2hfqR8MHoxCzAJBgNVBAYTAlVTMQsw\\nCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVyMQ8wDQYDVQQKEwZzZXRoLXMxCjAI\\nBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAwHgYJKoZIhvcNAQkBFhFzamZkZkBs\\nc2tkamZjLmNvbYIJANZCTNN98L9ZMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEF\\nBQADgYEAY/cJsi3w6R4hF4PzAXLhGOg1tzTDYvol3w024WoehJur+qM0AY6UqtoJ\\nneCq9af32IKbbOKkoaok+t1+/tylQVF/0FXMTKepxaMbG22vr4TmN3idPUYYbPfW\\n5GkF7Hh96BjerrtiUPGuBZL50HoLZ5aR5oZUMAu7TXhOFp+vZp8=\\n-----END CERTIFICATE-----\n",
 *     url: "http://mygameserver.local:7001/fusionauth-webhook",
 * });
 * ```
 */
export class FusionAuthWebhook extends pulumi.CustomResource {
    /**
     * Get an existing FusionAuthWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FusionAuthWebhookState, opts?: pulumi.CustomResourceOptions): FusionAuthWebhook {
        return new FusionAuthWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fusionauth:index/fusionAuthWebhook:FusionAuthWebhook';

    /**
     * Returns true if the given object is an instance of FusionAuthWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FusionAuthWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FusionAuthWebhook.__pulumiType;
    }

    /**
     * The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    public readonly connectTimeout!: pulumi.Output<number>;
    /**
     * A description of the Webhook. This is used for display purposes only.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A mapping for the events that are enabled for this Webhook.
     */
    public readonly eventsEnabled!: pulumi.Output<outputs.FusionAuthWebhookEventsEnabled | undefined>;
    /**
     * Whether or not this Webhook is used for all events or just for specific Applications.
     */
    public readonly global!: pulumi.Output<boolean | undefined>;
    /**
     * An object that contains headers that are sent as part of the HTTP request for the events.
     */
    public readonly headers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The HTTP basic authentication password that is sent as part of the HTTP request for the events.
     */
    public readonly httpAuthenticationPassword!: pulumi.Output<string | undefined>;
    /**
     * The HTTP basic authentication username that is sent as part of the HTTP request for the events.
     */
    public readonly httpAuthenticationUsername!: pulumi.Output<string | undefined>;
    /**
     * The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    public readonly readTimeout!: pulumi.Output<number>;
    /**
     * An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
     */
    public readonly sslCertificate!: pulumi.Output<string | undefined>;
    /**
     * The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
     */
    public readonly tenantIds!: pulumi.Output<string[] | undefined>;
    /**
     * The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a FusionAuthWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FusionAuthWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FusionAuthWebhookArgs | FusionAuthWebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FusionAuthWebhookState | undefined;
            resourceInputs["connectTimeout"] = state ? state.connectTimeout : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["eventsEnabled"] = state ? state.eventsEnabled : undefined;
            resourceInputs["global"] = state ? state.global : undefined;
            resourceInputs["headers"] = state ? state.headers : undefined;
            resourceInputs["httpAuthenticationPassword"] = state ? state.httpAuthenticationPassword : undefined;
            resourceInputs["httpAuthenticationUsername"] = state ? state.httpAuthenticationUsername : undefined;
            resourceInputs["readTimeout"] = state ? state.readTimeout : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["tenantIds"] = state ? state.tenantIds : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as FusionAuthWebhookArgs | undefined;
            if ((!args || args.connectTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectTimeout'");
            }
            if ((!args || args.readTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'readTimeout'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["connectTimeout"] = args ? args.connectTimeout : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eventsEnabled"] = args ? args.eventsEnabled : undefined;
            resourceInputs["global"] = args ? args.global : undefined;
            resourceInputs["headers"] = args ? args.headers : undefined;
            resourceInputs["httpAuthenticationPassword"] = args?.httpAuthenticationPassword ? pulumi.secret(args.httpAuthenticationPassword) : undefined;
            resourceInputs["httpAuthenticationUsername"] = args?.httpAuthenticationUsername ? pulumi.secret(args.httpAuthenticationUsername) : undefined;
            resourceInputs["readTimeout"] = args ? args.readTimeout : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["tenantIds"] = args ? args.tenantIds : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["httpAuthenticationPassword", "httpAuthenticationUsername"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FusionAuthWebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FusionAuthWebhook resources.
 */
export interface FusionAuthWebhookState {
    /**
     * The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    connectTimeout?: pulumi.Input<number>;
    /**
     * A description of the Webhook. This is used for display purposes only.
     */
    description?: pulumi.Input<string>;
    /**
     * A mapping for the events that are enabled for this Webhook.
     */
    eventsEnabled?: pulumi.Input<inputs.FusionAuthWebhookEventsEnabled>;
    /**
     * Whether or not this Webhook is used for all events or just for specific Applications.
     */
    global?: pulumi.Input<boolean>;
    /**
     * An object that contains headers that are sent as part of the HTTP request for the events.
     */
    headers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The HTTP basic authentication password that is sent as part of the HTTP request for the events.
     */
    httpAuthenticationPassword?: pulumi.Input<string>;
    /**
     * The HTTP basic authentication username that is sent as part of the HTTP request for the events.
     */
    httpAuthenticationUsername?: pulumi.Input<string>;
    /**
     * The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    readTimeout?: pulumi.Input<number>;
    /**
     * An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
     */
    tenantIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FusionAuthWebhook resource.
 */
export interface FusionAuthWebhookArgs {
    /**
     * The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    connectTimeout: pulumi.Input<number>;
    /**
     * A description of the Webhook. This is used for display purposes only.
     */
    description?: pulumi.Input<string>;
    /**
     * A mapping for the events that are enabled for this Webhook.
     */
    eventsEnabled?: pulumi.Input<inputs.FusionAuthWebhookEventsEnabled>;
    /**
     * Whether or not this Webhook is used for all events or just for specific Applications.
     */
    global?: pulumi.Input<boolean>;
    /**
     * An object that contains headers that are sent as part of the HTTP request for the events.
     */
    headers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The HTTP basic authentication password that is sent as part of the HTTP request for the events.
     */
    httpAuthenticationPassword?: pulumi.Input<string>;
    /**
     * The HTTP basic authentication username that is sent as part of the HTTP request for the events.
     */
    httpAuthenticationUsername?: pulumi.Input<string>;
    /**
     * The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
     */
    readTimeout: pulumi.Input<number>;
    /**
     * An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
     */
    tenantIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
     */
    url: pulumi.Input<string>;
}
