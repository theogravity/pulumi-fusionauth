// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth/internal"
)

// ## # Form Resource
//
// A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.
//
// [Forms API](https://fusionauth.io/docs/v1/tech/apis/forms)
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
//			_, err := fusionauth.GetForm(ctx, &fusionauth.GetFormArgs{
//				Name: pulumi.StringRef("Default User Self Service provided by FusionAuth"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetForm(ctx *pulumi.Context, args *GetFormArgs, opts ...pulumi.InvokeOption) (*GetFormResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFormResult
	err := ctx.Invoke("fusionauth:index/getForm:getForm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForm.
type GetFormArgs struct {
	// An object that can hold any information about the Form that should be persisted.
	Data map[string]string `pulumi:"data"`
	// The unique id of the Form. Either `formId` or `name` must be specified.
	FormId *string `pulumi:"formId"`
	// The name of the Form. Either `formId` or `name` must be specified.
	Name *string `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields.
	Steps []GetFormStep `pulumi:"steps"`
	// The form type. The possible values are:
	// * `adminRegistration` - This form be used to customize the add and edit User Registration form in the FusionAuth UI.
	// * `adminUser` - This form can be used to customize the add and edit User form in the FusionAuth UI.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getForm.
type GetFormResult struct {
	// An object that can hold any information about the Form that should be persisted.
	Data   map[string]string `pulumi:"data"`
	FormId string            `pulumi:"formId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The unique name of the Form.
	Name *string `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields.
	Steps []GetFormStep `pulumi:"steps"`
	// The form type. The possible values are:
	// * `adminRegistration` - This form be used to customize the add and edit User Registration form in the FusionAuth UI.
	// * `adminUser` - This form can be used to customize the add and edit User form in the FusionAuth UI.
	Type *string `pulumi:"type"`
}

func GetFormOutput(ctx *pulumi.Context, args GetFormOutputArgs, opts ...pulumi.InvokeOption) GetFormResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFormResultOutput, error) {
			args := v.(GetFormArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("fusionauth:index/getForm:getForm", args, GetFormResultOutput{}, options).(GetFormResultOutput), nil
		}).(GetFormResultOutput)
}

// A collection of arguments for invoking getForm.
type GetFormOutputArgs struct {
	// An object that can hold any information about the Form that should be persisted.
	Data pulumi.StringMapInput `pulumi:"data"`
	// The unique id of the Form. Either `formId` or `name` must be specified.
	FormId pulumi.StringPtrInput `pulumi:"formId"`
	// The name of the Form. Either `formId` or `name` must be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// An ordered list of objects containing one or more Form Fields.
	Steps GetFormStepArrayInput `pulumi:"steps"`
	// The form type. The possible values are:
	// * `adminRegistration` - This form be used to customize the add and edit User Registration form in the FusionAuth UI.
	// * `adminUser` - This form can be used to customize the add and edit User form in the FusionAuth UI.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetFormOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFormArgs)(nil)).Elem()
}

// A collection of values returned by getForm.
type GetFormResultOutput struct{ *pulumi.OutputState }

func (GetFormResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFormResult)(nil)).Elem()
}

func (o GetFormResultOutput) ToGetFormResultOutput() GetFormResultOutput {
	return o
}

func (o GetFormResultOutput) ToGetFormResultOutputWithContext(ctx context.Context) GetFormResultOutput {
	return o
}

// An object that can hold any information about the Form that should be persisted.
func (o GetFormResultOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetFormResult) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

func (o GetFormResultOutput) FormId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFormResult) string { return v.FormId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFormResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFormResult) string { return v.Id }).(pulumi.StringOutput)
}

// The unique name of the Form.
func (o GetFormResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFormResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An ordered list of objects containing one or more Form Fields.
func (o GetFormResultOutput) Steps() GetFormStepArrayOutput {
	return o.ApplyT(func(v GetFormResult) []GetFormStep { return v.Steps }).(GetFormStepArrayOutput)
}

// The form type. The possible values are:
// * `adminRegistration` - This form be used to customize the add and edit User Registration form in the FusionAuth UI.
// * `adminUser` - This form can be used to customize the add and edit User form in the FusionAuth UI.
func (o GetFormResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFormResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFormResultOutput{})
}
