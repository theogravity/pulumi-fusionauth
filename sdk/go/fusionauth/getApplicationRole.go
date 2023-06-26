// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Application Role Resource
//
// This Resource is used to create a role for an Application.
//
// [Application Roles API](https://fusionauth.io/docs/v1/tech/apis/applications)
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
//			_, err := fusionauth.GetApplicationRole(ctx, &fusionauth.GetApplicationRoleArgs{
//				ApplicationId: data.Fusionauth_application.FusionAuth.Id,
//				Name:          "admin",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetApplicationRole(ctx *pulumi.Context, args *GetApplicationRoleArgs, opts ...pulumi.InvokeOption) (*GetApplicationRoleResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetApplicationRoleResult
	err := ctx.Invoke("fusionauth:index/getApplicationRole:getApplicationRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationRole.
type GetApplicationRoleArgs struct {
	// ID of the application that this role is for.
	ApplicationId string `pulumi:"applicationId"`
	// The name of the Role.
	Name string `pulumi:"name"`
}

// A collection of values returned by getApplicationRole.
type GetApplicationRoleResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func GetApplicationRoleOutput(ctx *pulumi.Context, args GetApplicationRoleOutputArgs, opts ...pulumi.InvokeOption) GetApplicationRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationRoleResult, error) {
			args := v.(GetApplicationRoleArgs)
			r, err := GetApplicationRole(ctx, &args, opts...)
			var s GetApplicationRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationRoleResultOutput)
}

// A collection of arguments for invoking getApplicationRole.
type GetApplicationRoleOutputArgs struct {
	// ID of the application that this role is for.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// The name of the Role.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetApplicationRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationRoleArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationRole.
type GetApplicationRoleResultOutput struct{ *pulumi.OutputState }

func (GetApplicationRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationRoleResult)(nil)).Elem()
}

func (o GetApplicationRoleResultOutput) ToGetApplicationRoleResultOutput() GetApplicationRoleResultOutput {
	return o
}

func (o GetApplicationRoleResultOutput) ToGetApplicationRoleResultOutputWithContext(ctx context.Context) GetApplicationRoleResultOutput {
	return o
}

func (o GetApplicationRoleResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationRoleResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApplicationRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationRoleResultOutput{})
}
