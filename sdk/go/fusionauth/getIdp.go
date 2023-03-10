// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Application Resource
//
// [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/v2/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.GetIdp(ctx, &fusionauth.GetIdpArgs{
//				Name: "Apple",
//				Type: "Apple",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIdp(ctx *pulumi.Context, args *GetIdpArgs, opts ...pulumi.InvokeOption) (*GetIdpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIdpResult
	err := ctx.Invoke("fusionauth:index/getIdp:getIdp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdp.
type GetIdpArgs struct {
	// The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
	Name string `pulumi:"name"`
	// The type of the identity provider.
	Type string `pulumi:"type"`
}

// A collection of values returned by getIdp.
type GetIdpResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

func GetIdpOutput(ctx *pulumi.Context, args GetIdpOutputArgs, opts ...pulumi.InvokeOption) GetIdpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIdpResult, error) {
			args := v.(GetIdpArgs)
			r, err := GetIdp(ctx, &args, opts...)
			var s GetIdpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIdpResultOutput)
}

// A collection of arguments for invoking getIdp.
type GetIdpOutputArgs struct {
	// The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
	Name pulumi.StringInput `pulumi:"name"`
	// The type of the identity provider.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetIdpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdpArgs)(nil)).Elem()
}

// A collection of values returned by getIdp.
type GetIdpResultOutput struct{ *pulumi.OutputState }

func (GetIdpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIdpResult)(nil)).Elem()
}

func (o GetIdpResultOutput) ToGetIdpResultOutput() GetIdpResultOutput {
	return o
}

func (o GetIdpResultOutput) ToGetIdpResultOutputWithContext(ctx context.Context) GetIdpResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIdpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIdpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdpResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIdpResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetIdpResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIdpResultOutput{})
}
