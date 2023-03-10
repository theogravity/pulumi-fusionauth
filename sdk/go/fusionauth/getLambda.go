// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Lambda Resource
//
// Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.
//
// [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)
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
//			_, err := fusionauth.GetLambda(ctx, &fusionauth.GetLambdaArgs{
//				Name: pulumi.StringRef("Default Google Reconcile provided by FusionAuth"),
//				Type: "GoogleReconcile",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLambda(ctx *pulumi.Context, args *GetLambdaArgs, opts ...pulumi.InvokeOption) (*GetLambdaResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetLambdaResult
	err := ctx.Invoke("fusionauth:index/getLambda:getLambda", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLambda.
type GetLambdaArgs struct {
	// The ID of the Lambda. At least one of `id` or `name` must be specified.
	Id *string `pulumi:"id"`
	// The name of the Lambda. At least one of `id` or `name` must be specified.
	Name *string `pulumi:"name"`
	// The Lambda type. The possible values are:
	Type string `pulumi:"type"`
}

// A collection of values returned by getLambda.
type GetLambdaResult struct {
	// The lambda function body, a JavaScript function.
	Body string `pulumi:"body"`
	// Whether or not debug event logging is enabled for this Lambda.
	Debug bool    `pulumi:"debug"`
	Id    string  `pulumi:"id"`
	Name  *string `pulumi:"name"`
	Type  string  `pulumi:"type"`
}

func GetLambdaOutput(ctx *pulumi.Context, args GetLambdaOutputArgs, opts ...pulumi.InvokeOption) GetLambdaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLambdaResult, error) {
			args := v.(GetLambdaArgs)
			r, err := GetLambda(ctx, &args, opts...)
			var s GetLambdaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLambdaResultOutput)
}

// A collection of arguments for invoking getLambda.
type GetLambdaOutputArgs struct {
	// The ID of the Lambda. At least one of `id` or `name` must be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the Lambda. At least one of `id` or `name` must be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Lambda type. The possible values are:
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetLambdaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLambdaArgs)(nil)).Elem()
}

// A collection of values returned by getLambda.
type GetLambdaResultOutput struct{ *pulumi.OutputState }

func (GetLambdaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLambdaResult)(nil)).Elem()
}

func (o GetLambdaResultOutput) ToGetLambdaResultOutput() GetLambdaResultOutput {
	return o
}

func (o GetLambdaResultOutput) ToGetLambdaResultOutputWithContext(ctx context.Context) GetLambdaResultOutput {
	return o
}

// The lambda function body, a JavaScript function.
func (o GetLambdaResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetLambdaResult) string { return v.Body }).(pulumi.StringOutput)
}

// Whether or not debug event logging is enabled for this Lambda.
func (o GetLambdaResultOutput) Debug() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLambdaResult) bool { return v.Debug }).(pulumi.BoolOutput)
}

func (o GetLambdaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLambdaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLambdaResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLambdaResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetLambdaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetLambdaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLambdaResultOutput{})
}
