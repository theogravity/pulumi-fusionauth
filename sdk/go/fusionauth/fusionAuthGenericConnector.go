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

// ## # Generic Connector Resource
//
// A FusionAuth Generic Connector is a named object that provides configuration for allowing authentication against external systems.
//
// [Generic Connector API](https://fusionauth.io/docs/v1/tech/apis/connectors/generic/)
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
//			_, err := fusionauth.NewFusionAuthGenericConnector(ctx, "example", &fusionauth.FusionAuthGenericConnectorArgs{
//				AuthenticationUrl: pulumi.String("http://mygameserver.local:7001/fusionauth-connector"),
//				ConnectTimeout:    pulumi.Int(1000),
//				Data: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Debug: pulumi.Bool(false),
//				Headers: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//					"bar": pulumi.String("baz"),
//				},
//				HttpAuthenticationPassword: pulumi.String("supersecret"),
//				HttpAuthenticationUsername: pulumi.String("me"),
//				ReadTimeout:                pulumi.Int(2000),
//				SslCertificateKeyId:        pulumi.String("00000000-0000-0000-0000-000000000678"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthGenericConnector struct {
	pulumi.CustomResourceState

	// The fully qualified URL used to send an HTTP request to authenticate the user.
	AuthenticationUrl pulumi.StringOutput `pulumi:"authenticationUrl"`
	// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
	ConnectTimeout pulumi.IntOutput `pulumi:"connectTimeout"`
	// An object that can hold any information about the Connector that should be persisted.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
	Debug pulumi.BoolPtrOutput `pulumi:"debug"`
	// An object that can hold HTTPHeader key and value pairs.
	Headers pulumi.StringMapOutput `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrOutput `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrOutput `pulumi:"httpAuthenticationUsername"`
	// The unique Connector name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntOutput `pulumi:"readTimeout"`
	// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
	SslCertificateKeyId pulumi.StringPtrOutput `pulumi:"sslCertificateKeyId"`
}

// NewFusionAuthGenericConnector registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthGenericConnector(ctx *pulumi.Context,
	name string, args *FusionAuthGenericConnectorArgs, opts ...pulumi.ResourceOption) (*FusionAuthGenericConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationUrl == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationUrl'")
	}
	if args.ConnectTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ConnectTimeout'")
	}
	if args.ReadTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ReadTimeout'")
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
	var resource FusionAuthGenericConnector
	err := ctx.RegisterResource("fusionauth:index/fusionAuthGenericConnector:FusionAuthGenericConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthGenericConnector gets an existing FusionAuthGenericConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthGenericConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthGenericConnectorState, opts ...pulumi.ResourceOption) (*FusionAuthGenericConnector, error) {
	var resource FusionAuthGenericConnector
	err := ctx.ReadResource("fusionauth:index/fusionAuthGenericConnector:FusionAuthGenericConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthGenericConnector resources.
type fusionAuthGenericConnectorState struct {
	// The fully qualified URL used to send an HTTP request to authenticate the user.
	AuthenticationUrl *string `pulumi:"authenticationUrl"`
	// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
	ConnectTimeout *int `pulumi:"connectTimeout"`
	// An object that can hold any information about the Connector that should be persisted.
	Data map[string]string `pulumi:"data"`
	// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
	Debug *bool `pulumi:"debug"`
	// An object that can hold HTTPHeader key and value pairs.
	Headers map[string]string `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword *string `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername *string `pulumi:"httpAuthenticationUsername"`
	// The unique Connector name.
	Name *string `pulumi:"name"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout *int `pulumi:"readTimeout"`
	// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
	SslCertificateKeyId *string `pulumi:"sslCertificateKeyId"`
}

type FusionAuthGenericConnectorState struct {
	// The fully qualified URL used to send an HTTP request to authenticate the user.
	AuthenticationUrl pulumi.StringPtrInput
	// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
	ConnectTimeout pulumi.IntPtrInput
	// An object that can hold any information about the Connector that should be persisted.
	Data pulumi.StringMapInput
	// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
	Debug pulumi.BoolPtrInput
	// An object that can hold HTTPHeader key and value pairs.
	Headers pulumi.StringMapInput
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrInput
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrInput
	// The unique Connector name.
	Name pulumi.StringPtrInput
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntPtrInput
	// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
	SslCertificateKeyId pulumi.StringPtrInput
}

func (FusionAuthGenericConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthGenericConnectorState)(nil)).Elem()
}

type fusionAuthGenericConnectorArgs struct {
	// The fully qualified URL used to send an HTTP request to authenticate the user.
	AuthenticationUrl string `pulumi:"authenticationUrl"`
	// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
	ConnectTimeout int `pulumi:"connectTimeout"`
	// An object that can hold any information about the Connector that should be persisted.
	Data map[string]string `pulumi:"data"`
	// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
	Debug *bool `pulumi:"debug"`
	// An object that can hold HTTPHeader key and value pairs.
	Headers map[string]string `pulumi:"headers"`
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword *string `pulumi:"httpAuthenticationPassword"`
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername *string `pulumi:"httpAuthenticationUsername"`
	// The unique Connector name.
	Name *string `pulumi:"name"`
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout int `pulumi:"readTimeout"`
	// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
	SslCertificateKeyId *string `pulumi:"sslCertificateKeyId"`
}

// The set of arguments for constructing a FusionAuthGenericConnector resource.
type FusionAuthGenericConnectorArgs struct {
	// The fully qualified URL used to send an HTTP request to authenticate the user.
	AuthenticationUrl pulumi.StringInput
	// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
	ConnectTimeout pulumi.IntInput
	// An object that can hold any information about the Connector that should be persisted.
	Data pulumi.StringMapInput
	// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
	Debug pulumi.BoolPtrInput
	// An object that can hold HTTPHeader key and value pairs.
	Headers pulumi.StringMapInput
	// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
	HttpAuthenticationPassword pulumi.StringPtrInput
	// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
	HttpAuthenticationUsername pulumi.StringPtrInput
	// The unique Connector name.
	Name pulumi.StringPtrInput
	// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
	ReadTimeout pulumi.IntInput
	// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
	SslCertificateKeyId pulumi.StringPtrInput
}

func (FusionAuthGenericConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthGenericConnectorArgs)(nil)).Elem()
}

type FusionAuthGenericConnectorInput interface {
	pulumi.Input

	ToFusionAuthGenericConnectorOutput() FusionAuthGenericConnectorOutput
	ToFusionAuthGenericConnectorOutputWithContext(ctx context.Context) FusionAuthGenericConnectorOutput
}

func (*FusionAuthGenericConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthGenericConnector)(nil)).Elem()
}

func (i *FusionAuthGenericConnector) ToFusionAuthGenericConnectorOutput() FusionAuthGenericConnectorOutput {
	return i.ToFusionAuthGenericConnectorOutputWithContext(context.Background())
}

func (i *FusionAuthGenericConnector) ToFusionAuthGenericConnectorOutputWithContext(ctx context.Context) FusionAuthGenericConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthGenericConnectorOutput)
}

// FusionAuthGenericConnectorArrayInput is an input type that accepts FusionAuthGenericConnectorArray and FusionAuthGenericConnectorArrayOutput values.
// You can construct a concrete instance of `FusionAuthGenericConnectorArrayInput` via:
//
//	FusionAuthGenericConnectorArray{ FusionAuthGenericConnectorArgs{...} }
type FusionAuthGenericConnectorArrayInput interface {
	pulumi.Input

	ToFusionAuthGenericConnectorArrayOutput() FusionAuthGenericConnectorArrayOutput
	ToFusionAuthGenericConnectorArrayOutputWithContext(context.Context) FusionAuthGenericConnectorArrayOutput
}

type FusionAuthGenericConnectorArray []FusionAuthGenericConnectorInput

func (FusionAuthGenericConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthGenericConnector)(nil)).Elem()
}

func (i FusionAuthGenericConnectorArray) ToFusionAuthGenericConnectorArrayOutput() FusionAuthGenericConnectorArrayOutput {
	return i.ToFusionAuthGenericConnectorArrayOutputWithContext(context.Background())
}

func (i FusionAuthGenericConnectorArray) ToFusionAuthGenericConnectorArrayOutputWithContext(ctx context.Context) FusionAuthGenericConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthGenericConnectorArrayOutput)
}

// FusionAuthGenericConnectorMapInput is an input type that accepts FusionAuthGenericConnectorMap and FusionAuthGenericConnectorMapOutput values.
// You can construct a concrete instance of `FusionAuthGenericConnectorMapInput` via:
//
//	FusionAuthGenericConnectorMap{ "key": FusionAuthGenericConnectorArgs{...} }
type FusionAuthGenericConnectorMapInput interface {
	pulumi.Input

	ToFusionAuthGenericConnectorMapOutput() FusionAuthGenericConnectorMapOutput
	ToFusionAuthGenericConnectorMapOutputWithContext(context.Context) FusionAuthGenericConnectorMapOutput
}

type FusionAuthGenericConnectorMap map[string]FusionAuthGenericConnectorInput

func (FusionAuthGenericConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthGenericConnector)(nil)).Elem()
}

func (i FusionAuthGenericConnectorMap) ToFusionAuthGenericConnectorMapOutput() FusionAuthGenericConnectorMapOutput {
	return i.ToFusionAuthGenericConnectorMapOutputWithContext(context.Background())
}

func (i FusionAuthGenericConnectorMap) ToFusionAuthGenericConnectorMapOutputWithContext(ctx context.Context) FusionAuthGenericConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthGenericConnectorMapOutput)
}

type FusionAuthGenericConnectorOutput struct{ *pulumi.OutputState }

func (FusionAuthGenericConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthGenericConnector)(nil)).Elem()
}

func (o FusionAuthGenericConnectorOutput) ToFusionAuthGenericConnectorOutput() FusionAuthGenericConnectorOutput {
	return o
}

func (o FusionAuthGenericConnectorOutput) ToFusionAuthGenericConnectorOutputWithContext(ctx context.Context) FusionAuthGenericConnectorOutput {
	return o
}

// The fully qualified URL used to send an HTTP request to authenticate the user.
func (o FusionAuthGenericConnectorOutput) AuthenticationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringOutput { return v.AuthenticationUrl }).(pulumi.StringOutput)
}

// The connect timeout for the HTTP connection, in milliseconds. Value must be greater than 0.
func (o FusionAuthGenericConnectorOutput) ConnectTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.IntOutput { return v.ConnectTimeout }).(pulumi.IntOutput)
}

// An object that can hold any information about the Connector that should be persisted.
func (o FusionAuthGenericConnectorOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// Determines if debug should be enabled to create an event log to assist in debugging integration errors. Defaults to false.
func (o FusionAuthGenericConnectorOutput) Debug() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.BoolPtrOutput { return v.Debug }).(pulumi.BoolPtrOutput)
}

// An object that can hold HTTPHeader key and value pairs.
func (o FusionAuthGenericConnectorOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringMapOutput { return v.Headers }).(pulumi.StringMapOutput)
}

// The HTTP basic authentication password that is sent as part of the HTTP request for the events.
func (o FusionAuthGenericConnectorOutput) HttpAuthenticationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringPtrOutput { return v.HttpAuthenticationPassword }).(pulumi.StringPtrOutput)
}

// The HTTP basic authentication username that is sent as part of the HTTP request for the events.
func (o FusionAuthGenericConnectorOutput) HttpAuthenticationUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringPtrOutput { return v.HttpAuthenticationUsername }).(pulumi.StringPtrOutput)
}

// The unique Connector name.
func (o FusionAuthGenericConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The read timeout in milliseconds used when FusionAuth sends events to the Webhook.
func (o FusionAuthGenericConnectorOutput) ReadTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.IntOutput { return v.ReadTimeout }).(pulumi.IntOutput)
}

// The Id of an existing [Key](https://fusionauth.io/docs/v1/tech/apis/keys/). The X509 certificate is used for client certificate authentication in requests to the Connector.
func (o FusionAuthGenericConnectorOutput) SslCertificateKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthGenericConnector) pulumi.StringPtrOutput { return v.SslCertificateKeyId }).(pulumi.StringPtrOutput)
}

type FusionAuthGenericConnectorArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthGenericConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthGenericConnector)(nil)).Elem()
}

func (o FusionAuthGenericConnectorArrayOutput) ToFusionAuthGenericConnectorArrayOutput() FusionAuthGenericConnectorArrayOutput {
	return o
}

func (o FusionAuthGenericConnectorArrayOutput) ToFusionAuthGenericConnectorArrayOutputWithContext(ctx context.Context) FusionAuthGenericConnectorArrayOutput {
	return o
}

func (o FusionAuthGenericConnectorArrayOutput) Index(i pulumi.IntInput) FusionAuthGenericConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthGenericConnector {
		return vs[0].([]*FusionAuthGenericConnector)[vs[1].(int)]
	}).(FusionAuthGenericConnectorOutput)
}

type FusionAuthGenericConnectorMapOutput struct{ *pulumi.OutputState }

func (FusionAuthGenericConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthGenericConnector)(nil)).Elem()
}

func (o FusionAuthGenericConnectorMapOutput) ToFusionAuthGenericConnectorMapOutput() FusionAuthGenericConnectorMapOutput {
	return o
}

func (o FusionAuthGenericConnectorMapOutput) ToFusionAuthGenericConnectorMapOutputWithContext(ctx context.Context) FusionAuthGenericConnectorMapOutput {
	return o
}

func (o FusionAuthGenericConnectorMapOutput) MapIndex(k pulumi.StringInput) FusionAuthGenericConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthGenericConnector {
		return vs[0].(map[string]*FusionAuthGenericConnector)[vs[1].(string)]
	}).(FusionAuthGenericConnectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthGenericConnectorInput)(nil)).Elem(), &FusionAuthGenericConnector{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthGenericConnectorArrayInput)(nil)).Elem(), FusionAuthGenericConnectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthGenericConnectorMapInput)(nil)).Elem(), FusionAuthGenericConnectorMap{})
	pulumi.RegisterOutputType(FusionAuthGenericConnectorOutput{})
	pulumi.RegisterOutputType(FusionAuthGenericConnectorArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthGenericConnectorMapOutput{})
}
