// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # Application OAuth Scope Resource
//
// The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.
//
// [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)
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
//			_, err := fusionauth.LookupApplicationOauthScope(ctx, &fusionauth.LookupApplicationOauthScopeArgs{
//				ApplicationId: data.Fusionauth_application.This.Id,
//				Name:          "data:read",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupApplicationOauthScope(ctx *pulumi.Context, args *LookupApplicationOauthScopeArgs, opts ...pulumi.InvokeOption) (*LookupApplicationOauthScopeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationOauthScopeResult
	err := ctx.Invoke("fusionauth:index/getApplicationOauthScope:getApplicationOauthScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationOauthScope.
type LookupApplicationOauthScopeArgs struct {
	// ID of the application that this role is for.
	ApplicationId string `pulumi:"applicationId"`
	// The name of the Role.
	Name string `pulumi:"name"`
}

// A collection of values returned by getApplicationOauthScope.
type LookupApplicationOauthScopeResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// (Optional) An object that can hold any information about the OAuth Scope that should be persisted.
	Data map[string]string `pulumi:"data"`
	// (Optional) "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
	DefaultConsentDetail string `pulumi:"defaultConsentDetail"`
	// (Optional) The default message to display on the OAuth consent screen if one cannot be found in the theme.
	DefaultConsentMessage string `pulumi:"defaultConsentMessage"`
	// (Optional) A description of the OAuth Scope. This is used for display purposes only.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// (Optional) Determines if the OAuth Scope is required when requested in an OAuth workflow.
	Required bool `pulumi:"required"`
	// (Optional) The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
	ScopeId string `pulumi:"scopeId"`
}

func LookupApplicationOauthScopeOutput(ctx *pulumi.Context, args LookupApplicationOauthScopeOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationOauthScopeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApplicationOauthScopeResultOutput, error) {
			args := v.(LookupApplicationOauthScopeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("fusionauth:index/getApplicationOauthScope:getApplicationOauthScope", args, LookupApplicationOauthScopeResultOutput{}, options).(LookupApplicationOauthScopeResultOutput), nil
		}).(LookupApplicationOauthScopeResultOutput)
}

// A collection of arguments for invoking getApplicationOauthScope.
type LookupApplicationOauthScopeOutputArgs struct {
	// ID of the application that this role is for.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// The name of the Role.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupApplicationOauthScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationOauthScopeArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationOauthScope.
type LookupApplicationOauthScopeResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationOauthScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationOauthScopeResult)(nil)).Elem()
}

func (o LookupApplicationOauthScopeResultOutput) ToLookupApplicationOauthScopeResultOutput() LookupApplicationOauthScopeResultOutput {
	return o
}

func (o LookupApplicationOauthScopeResultOutput) ToLookupApplicationOauthScopeResultOutputWithContext(ctx context.Context) LookupApplicationOauthScopeResultOutput {
	return o
}

func (o LookupApplicationOauthScopeResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// (Optional) An object that can hold any information about the OAuth Scope that should be persisted.
func (o LookupApplicationOauthScopeResultOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

// (Optional) "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
func (o LookupApplicationOauthScopeResultOutput) DefaultConsentDetail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.DefaultConsentDetail }).(pulumi.StringOutput)
}

// (Optional) The default message to display on the OAuth consent screen if one cannot be found in the theme.
func (o LookupApplicationOauthScopeResultOutput) DefaultConsentMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.DefaultConsentMessage }).(pulumi.StringOutput)
}

// (Optional) A description of the OAuth Scope. This is used for display purposes only.
func (o LookupApplicationOauthScopeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApplicationOauthScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationOauthScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// (Optional) Determines if the OAuth Scope is required when requested in an OAuth workflow.
func (o LookupApplicationOauthScopeResultOutput) Required() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) bool { return v.Required }).(pulumi.BoolOutput)
}

// (Optional) The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
func (o LookupApplicationOauthScopeResultOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationOauthScopeResult) string { return v.ScopeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationOauthScopeResultOutput{})
}
