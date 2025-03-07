// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # API Key
//
// The FusionAuth APIs are primarily secured using API keys. This API can only be accessed using an API key that is set as a keyManager. In order to retrieve, update or delete an API key, an API key with equal or greater permissions must be used. A "tenant-scoped" API key can retrieve, create, update or delete an API key for the same tenant. This page describes APIs that are used to manage API keys.
//
// [API Key](https://fusionauth.io/docs/v1/tech/apis/api-keys/)
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
//			_, err := fusionauth.NewFusionAuthApiKey(ctx, "example", &fusionauth.FusionAuthApiKeyArgs{
//				Description: pulumi.String("my super secret key"),
//				Key:         pulumi.String("super-secret-key"),
//				PermissionsEndpoints: fusionauth.FusionAuthApiKeyPermissionsEndpointArray{
//					&fusionauth.FusionAuthApiKeyPermissionsEndpointArgs{
//						Delete:   pulumi.Bool(true),
//						Endpoint: pulumi.String("/api/application"),
//						Get:      pulumi.Bool(true),
//						Patch:    pulumi.Bool(true),
//						Post:     pulumi.Bool(true),
//						Put:      pulumi.Bool(true),
//					},
//				},
//				TenantId: pulumi.String("94f751c5-4883-4684-a817-6b106778edec"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FusionAuthApiKey struct {
	pulumi.CustomResourceState

	// Description of the key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
	ExpirationInstant pulumi.IntPtrOutput `pulumi:"expirationInstant"`
	// The Id of the IP Access Control List limiting access to this API key.
	IpAccessControlListId pulumi.StringPtrOutput `pulumi:"ipAccessControlListId"`
	// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
	Key pulumi.StringOutput `pulumi:"key"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
	PermissionsEndpoints FusionAuthApiKeyPermissionsEndpointArrayOutput `pulumi:"permissionsEndpoints"`
	// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
}

// NewFusionAuthApiKey registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthApiKey(ctx *pulumi.Context,
	name string, args *FusionAuthApiKeyArgs, opts ...pulumi.ResourceOption) (*FusionAuthApiKey, error) {
	if args == nil {
		args = &FusionAuthApiKeyArgs{}
	}

	if args.Key != nil {
		args.Key = pulumi.ToSecret(args.Key).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthApiKey
	err := ctx.RegisterResource("fusionauth:index/fusionAuthApiKey:FusionAuthApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthApiKey gets an existing FusionAuthApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthApiKeyState, opts ...pulumi.ResourceOption) (*FusionAuthApiKey, error) {
	var resource FusionAuthApiKey
	err := ctx.ReadResource("fusionauth:index/fusionAuthApiKey:FusionAuthApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthApiKey resources.
type fusionAuthApiKeyState struct {
	// Description of the key.
	Description *string `pulumi:"description"`
	// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
	ExpirationInstant *int `pulumi:"expirationInstant"`
	// The Id of the IP Access Control List limiting access to this API key.
	IpAccessControlListId *string `pulumi:"ipAccessControlListId"`
	// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
	Key *string `pulumi:"key"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	KeyId *string `pulumi:"keyId"`
	// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
	PermissionsEndpoints []FusionAuthApiKeyPermissionsEndpoint `pulumi:"permissionsEndpoints"`
	// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
	TenantId *string `pulumi:"tenantId"`
}

type FusionAuthApiKeyState struct {
	// Description of the key.
	Description pulumi.StringPtrInput
	// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
	ExpirationInstant pulumi.IntPtrInput
	// The Id of the IP Access Control List limiting access to this API key.
	IpAccessControlListId pulumi.StringPtrInput
	// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
	Key pulumi.StringPtrInput
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringPtrInput
	// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
	PermissionsEndpoints FusionAuthApiKeyPermissionsEndpointArrayInput
	// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
	TenantId pulumi.StringPtrInput
}

func (FusionAuthApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthApiKeyState)(nil)).Elem()
}

type fusionAuthApiKeyArgs struct {
	// Description of the key.
	Description *string `pulumi:"description"`
	// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
	ExpirationInstant *int `pulumi:"expirationInstant"`
	// The Id of the IP Access Control List limiting access to this API key.
	IpAccessControlListId *string `pulumi:"ipAccessControlListId"`
	// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
	Key *string `pulumi:"key"`
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	KeyId *string `pulumi:"keyId"`
	// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
	PermissionsEndpoints []FusionAuthApiKeyPermissionsEndpoint `pulumi:"permissionsEndpoints"`
	// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a FusionAuthApiKey resource.
type FusionAuthApiKeyArgs struct {
	// Description of the key.
	Description pulumi.StringPtrInput
	// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
	ExpirationInstant pulumi.IntPtrInput
	// The Id of the IP Access Control List limiting access to this API key.
	IpAccessControlListId pulumi.StringPtrInput
	// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
	Key pulumi.StringPtrInput
	// The Id to use for the new Form. If not specified a secure random UUID will be generated.
	KeyId pulumi.StringPtrInput
	// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
	PermissionsEndpoints FusionAuthApiKeyPermissionsEndpointArrayInput
	// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
	TenantId pulumi.StringPtrInput
}

func (FusionAuthApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthApiKeyArgs)(nil)).Elem()
}

type FusionAuthApiKeyInput interface {
	pulumi.Input

	ToFusionAuthApiKeyOutput() FusionAuthApiKeyOutput
	ToFusionAuthApiKeyOutputWithContext(ctx context.Context) FusionAuthApiKeyOutput
}

func (*FusionAuthApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthApiKey)(nil)).Elem()
}

func (i *FusionAuthApiKey) ToFusionAuthApiKeyOutput() FusionAuthApiKeyOutput {
	return i.ToFusionAuthApiKeyOutputWithContext(context.Background())
}

func (i *FusionAuthApiKey) ToFusionAuthApiKeyOutputWithContext(ctx context.Context) FusionAuthApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApiKeyOutput)
}

// FusionAuthApiKeyArrayInput is an input type that accepts FusionAuthApiKeyArray and FusionAuthApiKeyArrayOutput values.
// You can construct a concrete instance of `FusionAuthApiKeyArrayInput` via:
//
//	FusionAuthApiKeyArray{ FusionAuthApiKeyArgs{...} }
type FusionAuthApiKeyArrayInput interface {
	pulumi.Input

	ToFusionAuthApiKeyArrayOutput() FusionAuthApiKeyArrayOutput
	ToFusionAuthApiKeyArrayOutputWithContext(context.Context) FusionAuthApiKeyArrayOutput
}

type FusionAuthApiKeyArray []FusionAuthApiKeyInput

func (FusionAuthApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthApiKey)(nil)).Elem()
}

func (i FusionAuthApiKeyArray) ToFusionAuthApiKeyArrayOutput() FusionAuthApiKeyArrayOutput {
	return i.ToFusionAuthApiKeyArrayOutputWithContext(context.Background())
}

func (i FusionAuthApiKeyArray) ToFusionAuthApiKeyArrayOutputWithContext(ctx context.Context) FusionAuthApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApiKeyArrayOutput)
}

// FusionAuthApiKeyMapInput is an input type that accepts FusionAuthApiKeyMap and FusionAuthApiKeyMapOutput values.
// You can construct a concrete instance of `FusionAuthApiKeyMapInput` via:
//
//	FusionAuthApiKeyMap{ "key": FusionAuthApiKeyArgs{...} }
type FusionAuthApiKeyMapInput interface {
	pulumi.Input

	ToFusionAuthApiKeyMapOutput() FusionAuthApiKeyMapOutput
	ToFusionAuthApiKeyMapOutputWithContext(context.Context) FusionAuthApiKeyMapOutput
}

type FusionAuthApiKeyMap map[string]FusionAuthApiKeyInput

func (FusionAuthApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthApiKey)(nil)).Elem()
}

func (i FusionAuthApiKeyMap) ToFusionAuthApiKeyMapOutput() FusionAuthApiKeyMapOutput {
	return i.ToFusionAuthApiKeyMapOutputWithContext(context.Background())
}

func (i FusionAuthApiKeyMap) ToFusionAuthApiKeyMapOutputWithContext(ctx context.Context) FusionAuthApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthApiKeyMapOutput)
}

type FusionAuthApiKeyOutput struct{ *pulumi.OutputState }

func (FusionAuthApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthApiKey)(nil)).Elem()
}

func (o FusionAuthApiKeyOutput) ToFusionAuthApiKeyOutput() FusionAuthApiKeyOutput {
	return o
}

func (o FusionAuthApiKeyOutput) ToFusionAuthApiKeyOutputWithContext(ctx context.Context) FusionAuthApiKeyOutput {
	return o
}

// Description of the key.
func (o FusionAuthApiKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The expiration instant of this API key. Using an expired API key for API Authentication will result in a 401 response code.
func (o FusionAuthApiKeyOutput) ExpirationInstant() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.IntPtrOutput { return v.ExpirationInstant }).(pulumi.IntPtrOutput)
}

// The Id of the IP Access Control List limiting access to this API key.
func (o FusionAuthApiKeyOutput) IpAccessControlListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.StringPtrOutput { return v.IpAccessControlListId }).(pulumi.StringPtrOutput)
}

// API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
func (o FusionAuthApiKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The Id to use for the new Form. If not specified a secure random UUID will be generated.
func (o FusionAuthApiKeyOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Endpoint permissions for this key. Each key of the object is an endpoint, with the value being an array of the HTTP methods which can be used against the endpoint. An Empty permissionsEndpoints object mean that this is a super key that authorizes this key for all the endpoints.
func (o FusionAuthApiKeyOutput) PermissionsEndpoints() FusionAuthApiKeyPermissionsEndpointArrayOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) FusionAuthApiKeyPermissionsEndpointArrayOutput {
		return v.PermissionsEndpoints
	}).(FusionAuthApiKeyPermissionsEndpointArrayOutput)
}

// The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
func (o FusionAuthApiKeyOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthApiKey) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

type FusionAuthApiKeyArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthApiKey)(nil)).Elem()
}

func (o FusionAuthApiKeyArrayOutput) ToFusionAuthApiKeyArrayOutput() FusionAuthApiKeyArrayOutput {
	return o
}

func (o FusionAuthApiKeyArrayOutput) ToFusionAuthApiKeyArrayOutputWithContext(ctx context.Context) FusionAuthApiKeyArrayOutput {
	return o
}

func (o FusionAuthApiKeyArrayOutput) Index(i pulumi.IntInput) FusionAuthApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthApiKey {
		return vs[0].([]*FusionAuthApiKey)[vs[1].(int)]
	}).(FusionAuthApiKeyOutput)
}

type FusionAuthApiKeyMapOutput struct{ *pulumi.OutputState }

func (FusionAuthApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthApiKey)(nil)).Elem()
}

func (o FusionAuthApiKeyMapOutput) ToFusionAuthApiKeyMapOutput() FusionAuthApiKeyMapOutput {
	return o
}

func (o FusionAuthApiKeyMapOutput) ToFusionAuthApiKeyMapOutputWithContext(ctx context.Context) FusionAuthApiKeyMapOutput {
	return o
}

func (o FusionAuthApiKeyMapOutput) MapIndex(k pulumi.StringInput) FusionAuthApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthApiKey {
		return vs[0].(map[string]*FusionAuthApiKey)[vs[1].(string)]
	}).(FusionAuthApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApiKeyInput)(nil)).Elem(), &FusionAuthApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApiKeyArrayInput)(nil)).Elem(), FusionAuthApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthApiKeyMapInput)(nil)).Elem(), FusionAuthApiKeyMap{})
	pulumi.RegisterOutputType(FusionAuthApiKeyOutput{})
	pulumi.RegisterOutputType(FusionAuthApiKeyArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthApiKeyMapOutput{})
}
