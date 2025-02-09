// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # Webhook Resource
//
// A FusionAuth Webhook is intended to consume JSON events emitted by FusionAuth. Creating a Webhook allows you to tell FusionAuth where you would like to receive these JSON events.
//
// [Webhooks API](https://fusionauth.io/docs/v1/tech/apis/webhooks)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthWebhook(ctx, "example", &fusionauth.FusionAuthWebhookArgs{
//				TenantIds: pulumi.StringArray{
//					pulumi.String("00000000-0000-0000-0000-000000000003"),
//					fusionauth_tenant.Example.Id,
//				},
//				ConnectTimeout: pulumi.Int(1000),
//				Description:    pulumi.String("The standard game Webhook"),
//				EventsEnabled: &fusionauth.FusionAuthWebhookEventsEnabledArgs{
//					UserCreate: pulumi.Bool(true),
//					UserDelete: pulumi.Bool(false),
//				},
//				Global: pulumi.Bool(false),
//				Headers: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//					"bar": pulumi.String("baz"),
//				},
//				HttpAuthenticationPassword: pulumi.String("password"),
//				HttpAuthenticationUsername: pulumi.String("username"),
//				ReadTimeout:                pulumi.Int(2000),
//				SslCertificate:             pulumi.String("  -----BEGIN CERTIFICATE-----\\nMIIDUjCCArugAwIBAgIJANZCTNN98L9ZMA0GCSqGSIb3DQEBBQUAMHoxCzAJBgNV\\nBAYTAlVTMQswCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVyMQ8wDQYDVQQKEwZz\\nZXRoLXMxCjAIBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAwHgYJKoZIhvcNAQkB\\nFhFzamZkZkBsc2tkamZjLmNvbTAeFw0xNDA0MDkyMTA2MDdaFw0xNDA1MDkyMTA2\\nMDdaMHoxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVy\\nMQ8wDQYDVQQKEwZzZXRoLXMxCjAIBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAw\\nHgYJKoZIhvcNAQkBFhFzamZkZkBsc2tkamZjLmNvbTCBnzANBgkqhkiG9w0BAQEF\\nAAOBjQAwgYkCgYEAxnQBqyuYvjUE4aFQ6vVZU5RqHmy3KiTg2NcxELIlZztUTK3a\\nVFbJoBB4ixHXCCYslujthILyBjgT3F+IhSpPAcrlu8O5LVPaPCysh/SNrGNwH4lq\\neiW9Z5WAhRO/nG7NZNa0USPHAei6b9Sv9PxuKCY+GJfAIwlO4/bltIH06/kCAwEA\\nAaOB3zCB3DAdBgNVHQ4EFgQUU4SqJEFm1zW+CcLxmLlARrqtMN0wgawGA1UdIwSB\\npDCBoYAUU4SqJEFm1zW+CcLxmLlARrqtMN2hfqR8MHoxCzAJBgNVBAYTAlVTMQsw\\nCQYDVQQIEwJDTzEPMA0GA1UEBxMGZGVudmVyMQ8wDQYDVQQKEwZzZXRoLXMxCjAI\\nBgNVBAsTAXMxDjAMBgNVBAMTBWludmVyMSAwHgYJKoZIhvcNAQkBFhFzamZkZkBs\\nc2tkamZjLmNvbYIJANZCTNN98L9ZMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEF\\nBQADgYEAY/cJsi3w6R4hF4PzAXLhGOg1tzTDYvol3w024WoehJur+qM0AY6UqtoJ\\nneCq9af32IKbbOKkoaok+t1+/tylQVF/0FXMTKepxaMbG22vr4TmN3idPUYYbPfW\\n5GkF7Hh96BjerrtiUPGuBZL50HoLZ5aR5oZUMAu7TXhOFp+vZp8=\\n-----END CERTIFICATE-----\n"),
//				Url:                        pulumi.String("http://mygameserver.local:7001/fusionauth-webhook"),
//				SignatureConfiguration: &fusionauth.FusionAuthWebhookSignatureConfigurationArgs{
//					Enabled:      pulumi.Bool(true),
//					SigningKeyId: pulumi.Any(fusionauth_key.Webhook_key.Id),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthWebhook struct {
	pulumi.CustomResourceState

	// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ConnectTimeout pulumi.IntOutput `pulumi:"connectTimeout"`
	// An object that can hold any information about the Webhook that should be persisted.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// A description of the Webhook. This is used for display purposes only.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A mapping for the events that are enabled for this Webhook.
	EventsEnabled FusionAuthWebhookEventsEnabledPtrOutput `pulumi:"eventsEnabled"`
	// Whether or not this Webhook is used for all events or just for specific Applications.
	Global pulumi.BoolPtrOutput `pulumi:"global"`
	// An object that contains headers that are sent as part of the HTTP request for the events.
	Headers pulumi.StringMapOutput `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrOutput `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrOutput `pulumi:"httpAuthenticationUsername"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntOutput `pulumi:"readTimeout"`
	// Configuration for webhook signing
	SignatureConfiguration FusionAuthWebhookSignatureConfigurationPtrOutput `pulumi:"signatureConfiguration"`
	// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
	SslCertificate pulumi.StringPtrOutput `pulumi:"sslCertificate"`
	// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
	SslCertificateKeyId pulumi.StringPtrOutput `pulumi:"sslCertificateKeyId"`
	// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
	TenantIds pulumi.StringArrayOutput `pulumi:"tenantIds"`
	// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
	Url pulumi.StringOutput `pulumi:"url"`
	// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
	WebhookId pulumi.StringOutput `pulumi:"webhookId"`
}

// NewFusionAuthWebhook registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthWebhook(ctx *pulumi.Context,
	name string, args *FusionAuthWebhookArgs, opts ...pulumi.ResourceOption) (*FusionAuthWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ConnectTimeout'")
	}
	if args.ReadTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ReadTimeout'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.HttpAuthenticationPassword != nil {
		args.HttpAuthenticationPassword = pulumi.ToSecret(args.HttpAuthenticationPassword).(pulumi.StringPtrInput)
	}
	if args.HttpAuthenticationUsername != nil {
		args.HttpAuthenticationUsername = pulumi.ToSecret(args.HttpAuthenticationUsername).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"httpAuthenticationPassword",
		"httpAuthenticationUsername",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthWebhook
	err := ctx.RegisterResource("fusionauth:index/fusionAuthWebhook:FusionAuthWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthWebhook gets an existing FusionAuthWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthWebhookState, opts ...pulumi.ResourceOption) (*FusionAuthWebhook, error) {
	var resource FusionAuthWebhook
	err := ctx.ReadResource("fusionauth:index/fusionAuthWebhook:FusionAuthWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthWebhook resources.
type fusionAuthWebhookState struct {
	// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// An object that can hold any information about the Webhook that should be persisted.
	Data map[string]string `pulumi:"data"`
	// A description of the Webhook. This is used for display purposes only.
	Description *string `pulumi:"description"`
	// A mapping for the events that are enabled for this Webhook.
	EventsEnabled *FusionAuthWebhookEventsEnabled `pulumi:"eventsEnabled"`
	// Whether or not this Webhook is used for all events or just for specific Applications.
	Global *bool `pulumi:"global"`
	// An object that contains headers that are sent as part of the HTTP request for the events.
	Headers map[string]string `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword *string `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername *string `pulumi:"httpAuthenticationUsername"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout *int `pulumi:"readTimeout"`
	// Configuration for webhook signing
	SignatureConfiguration *FusionAuthWebhookSignatureConfiguration `pulumi:"signatureConfiguration"`
	// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
	SslCertificate *string `pulumi:"sslCertificate"`
	// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
	SslCertificateKeyId *string `pulumi:"sslCertificateKeyId"`
	// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
	TenantIds []string `pulumi:"tenantIds"`
	// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
	Url *string `pulumi:"url"`
	// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
	WebhookId *string `pulumi:"webhookId"`
}

type FusionAuthWebhookState struct {
	// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ConnectTimeout pulumi.IntPtrInput
	// An object that can hold any information about the Webhook that should be persisted.
	Data pulumi.StringMapInput
	// A description of the Webhook. This is used for display purposes only.
	Description pulumi.StringPtrInput
	// A mapping for the events that are enabled for this Webhook.
	EventsEnabled FusionAuthWebhookEventsEnabledPtrInput
	// Whether or not this Webhook is used for all events or just for specific Applications.
	Global pulumi.BoolPtrInput
	// An object that contains headers that are sent as part of the HTTP request for the events.
	Headers pulumi.StringMapInput
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrInput
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrInput
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntPtrInput
	// Configuration for webhook signing
	SignatureConfiguration FusionAuthWebhookSignatureConfigurationPtrInput
	// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
	SslCertificate pulumi.StringPtrInput
	// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
	SslCertificateKeyId pulumi.StringPtrInput
	// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
	TenantIds pulumi.StringArrayInput
	// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
	Url pulumi.StringPtrInput
	// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
	WebhookId pulumi.StringPtrInput
}

func (FusionAuthWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthWebhookState)(nil)).Elem()
}

type fusionAuthWebhookArgs struct {
	// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ConnectTimeout int `pulumi:"connectTimeout"`
	// An object that can hold any information about the Webhook that should be persisted.
	Data map[string]string `pulumi:"data"`
	// A description of the Webhook. This is used for display purposes only.
	Description *string `pulumi:"description"`
	// A mapping for the events that are enabled for this Webhook.
	EventsEnabled *FusionAuthWebhookEventsEnabled `pulumi:"eventsEnabled"`
	// Whether or not this Webhook is used for all events or just for specific Applications.
	Global *bool `pulumi:"global"`
	// An object that contains headers that are sent as part of the HTTP request for the events.
	Headers map[string]string `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword *string `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername *string `pulumi:"httpAuthenticationUsername"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout int `pulumi:"readTimeout"`
	// Configuration for webhook signing
	SignatureConfiguration *FusionAuthWebhookSignatureConfiguration `pulumi:"signatureConfiguration"`
	// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
	SslCertificate *string `pulumi:"sslCertificate"`
	// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
	SslCertificateKeyId *string `pulumi:"sslCertificateKeyId"`
	// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
	TenantIds []string `pulumi:"tenantIds"`
	// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
	Url string `pulumi:"url"`
	// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
	WebhookId *string `pulumi:"webhookId"`
}

// The set of arguments for constructing a FusionAuthWebhook resource.
type FusionAuthWebhookArgs struct {
	// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ConnectTimeout pulumi.IntInput
	// An object that can hold any information about the Webhook that should be persisted.
	Data pulumi.StringMapInput
	// A description of the Webhook. This is used for display purposes only.
	Description pulumi.StringPtrInput
	// A mapping for the events that are enabled for this Webhook.
	EventsEnabled FusionAuthWebhookEventsEnabledPtrInput
	// Whether or not this Webhook is used for all events or just for specific Applications.
	Global pulumi.BoolPtrInput
	// An object that contains headers that are sent as part of the HTTP request for the events.
	Headers pulumi.StringMapInput
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrInput
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrInput
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntInput
	// Configuration for webhook signing
	SignatureConfiguration FusionAuthWebhookSignatureConfigurationPtrInput
	// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
	SslCertificate pulumi.StringPtrInput
	// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
	SslCertificateKeyId pulumi.StringPtrInput
	// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
	TenantIds pulumi.StringArrayInput
	// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
	Url pulumi.StringInput
	// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
	WebhookId pulumi.StringPtrInput
}

func (FusionAuthWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthWebhookArgs)(nil)).Elem()
}

type FusionAuthWebhookInput interface {
	pulumi.Input

	ToFusionAuthWebhookOutput() FusionAuthWebhookOutput
	ToFusionAuthWebhookOutputWithContext(ctx context.Context) FusionAuthWebhookOutput
}

func (*FusionAuthWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthWebhook)(nil)).Elem()
}

func (i *FusionAuthWebhook) ToFusionAuthWebhookOutput() FusionAuthWebhookOutput {
	return i.ToFusionAuthWebhookOutputWithContext(context.Background())
}

func (i *FusionAuthWebhook) ToFusionAuthWebhookOutputWithContext(ctx context.Context) FusionAuthWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthWebhookOutput)
}

// FusionAuthWebhookArrayInput is an input type that accepts FusionAuthWebhookArray and FusionAuthWebhookArrayOutput values.
// You can construct a concrete instance of `FusionAuthWebhookArrayInput` via:
//
//	FusionAuthWebhookArray{ FusionAuthWebhookArgs{...} }
type FusionAuthWebhookArrayInput interface {
	pulumi.Input

	ToFusionAuthWebhookArrayOutput() FusionAuthWebhookArrayOutput
	ToFusionAuthWebhookArrayOutputWithContext(context.Context) FusionAuthWebhookArrayOutput
}

type FusionAuthWebhookArray []FusionAuthWebhookInput

func (FusionAuthWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthWebhook)(nil)).Elem()
}

func (i FusionAuthWebhookArray) ToFusionAuthWebhookArrayOutput() FusionAuthWebhookArrayOutput {
	return i.ToFusionAuthWebhookArrayOutputWithContext(context.Background())
}

func (i FusionAuthWebhookArray) ToFusionAuthWebhookArrayOutputWithContext(ctx context.Context) FusionAuthWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthWebhookArrayOutput)
}

// FusionAuthWebhookMapInput is an input type that accepts FusionAuthWebhookMap and FusionAuthWebhookMapOutput values.
// You can construct a concrete instance of `FusionAuthWebhookMapInput` via:
//
//	FusionAuthWebhookMap{ "key": FusionAuthWebhookArgs{...} }
type FusionAuthWebhookMapInput interface {
	pulumi.Input

	ToFusionAuthWebhookMapOutput() FusionAuthWebhookMapOutput
	ToFusionAuthWebhookMapOutputWithContext(context.Context) FusionAuthWebhookMapOutput
}

type FusionAuthWebhookMap map[string]FusionAuthWebhookInput

func (FusionAuthWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthWebhook)(nil)).Elem()
}

func (i FusionAuthWebhookMap) ToFusionAuthWebhookMapOutput() FusionAuthWebhookMapOutput {
	return i.ToFusionAuthWebhookMapOutputWithContext(context.Background())
}

func (i FusionAuthWebhookMap) ToFusionAuthWebhookMapOutputWithContext(ctx context.Context) FusionAuthWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthWebhookMapOutput)
}

type FusionAuthWebhookOutput struct{ *pulumi.OutputState }

func (FusionAuthWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthWebhook)(nil)).Elem()
}

func (o FusionAuthWebhookOutput) ToFusionAuthWebhookOutput() FusionAuthWebhookOutput {
	return o
}

func (o FusionAuthWebhookOutput) ToFusionAuthWebhookOutputWithContext(ctx context.Context) FusionAuthWebhookOutput {
	return o
}

// The connection timeout in milliseconds used when FusionAuth sends events to the Webhook.
func (o FusionAuthWebhookOutput) ConnectTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.IntOutput { return v.ConnectTimeout }).(pulumi.IntOutput)
}

// An object that can hold any information about the Webhook that should be persisted.
func (o FusionAuthWebhookOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// A description of the Webhook. This is used for display purposes only.
func (o FusionAuthWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A mapping for the events that are enabled for this Webhook.
func (o FusionAuthWebhookOutput) EventsEnabled() FusionAuthWebhookEventsEnabledPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) FusionAuthWebhookEventsEnabledPtrOutput { return v.EventsEnabled }).(FusionAuthWebhookEventsEnabledPtrOutput)
}

// Whether or not this Webhook is used for all events or just for specific Applications.
func (o FusionAuthWebhookOutput) Global() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.BoolPtrOutput { return v.Global }).(pulumi.BoolPtrOutput)
}

// An object that contains headers that are sent as part of the HTTP request for the events.
func (o FusionAuthWebhookOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringMapOutput { return v.Headers }).(pulumi.StringMapOutput)
}

// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
func (o FusionAuthWebhookOutput) HttpAuthenticationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringPtrOutput { return v.HttpAuthenticationPassword }).(pulumi.StringPtrOutput)
}

// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
func (o FusionAuthWebhookOutput) HttpAuthenticationUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringPtrOutput { return v.HttpAuthenticationUsername }).(pulumi.StringPtrOutput)
}

// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
func (o FusionAuthWebhookOutput) ReadTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.IntOutput { return v.ReadTimeout }).(pulumi.IntOutput)
}

// Configuration for webhook signing
func (o FusionAuthWebhookOutput) SignatureConfiguration() FusionAuthWebhookSignatureConfigurationPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) FusionAuthWebhookSignatureConfigurationPtrOutput {
		return v.SignatureConfiguration
	}).(FusionAuthWebhookSignatureConfigurationPtrOutput)
}

// An SSL certificate in PEM format that is used to establish the a SSL (TLS specifically) connection to the Webhook.
func (o FusionAuthWebhookOutput) SslCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringPtrOutput { return v.SslCertificate }).(pulumi.StringPtrOutput)
}

// The Id of an existing Key. The X.509 certificate is used for client certificate authentication in requests to the Webhook.
func (o FusionAuthWebhookOutput) SslCertificateKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringPtrOutput { return v.SslCertificateKeyId }).(pulumi.StringPtrOutput)
}

// The Ids of the tenants that this Webhook should be associated with. If no Ids are specified and the global field is false, this Webhook will not be used.
func (o FusionAuthWebhookOutput) TenantIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringArrayOutput { return v.TenantIds }).(pulumi.StringArrayOutput)
}

// The fully qualified URL of the Webhook’s endpoint that will accept the event requests from FusionAuth.
func (o FusionAuthWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The Id to use for the new Webhook. If not specified a secure random UUID will be generated.
func (o FusionAuthWebhookOutput) WebhookId() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthWebhook) pulumi.StringOutput { return v.WebhookId }).(pulumi.StringOutput)
}

type FusionAuthWebhookArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthWebhook)(nil)).Elem()
}

func (o FusionAuthWebhookArrayOutput) ToFusionAuthWebhookArrayOutput() FusionAuthWebhookArrayOutput {
	return o
}

func (o FusionAuthWebhookArrayOutput) ToFusionAuthWebhookArrayOutputWithContext(ctx context.Context) FusionAuthWebhookArrayOutput {
	return o
}

func (o FusionAuthWebhookArrayOutput) Index(i pulumi.IntInput) FusionAuthWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthWebhook {
		return vs[0].([]*FusionAuthWebhook)[vs[1].(int)]
	}).(FusionAuthWebhookOutput)
}

type FusionAuthWebhookMapOutput struct{ *pulumi.OutputState }

func (FusionAuthWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthWebhook)(nil)).Elem()
}

func (o FusionAuthWebhookMapOutput) ToFusionAuthWebhookMapOutput() FusionAuthWebhookMapOutput {
	return o
}

func (o FusionAuthWebhookMapOutput) ToFusionAuthWebhookMapOutputWithContext(ctx context.Context) FusionAuthWebhookMapOutput {
	return o
}

func (o FusionAuthWebhookMapOutput) MapIndex(k pulumi.StringInput) FusionAuthWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthWebhook {
		return vs[0].(map[string]*FusionAuthWebhook)[vs[1].(string)]
	}).(FusionAuthWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthWebhookInput)(nil)).Elem(), &FusionAuthWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthWebhookArrayInput)(nil)).Elem(), FusionAuthWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthWebhookMapInput)(nil)).Elem(), FusionAuthWebhookMap{})
	pulumi.RegisterOutputType(FusionAuthWebhookOutput{})
	pulumi.RegisterOutputType(FusionAuthWebhookArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthWebhookMapOutput{})
}
