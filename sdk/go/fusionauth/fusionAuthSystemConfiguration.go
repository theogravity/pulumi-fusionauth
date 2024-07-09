// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fusionauth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/theogravity/pulumi-fusionauth/sdk/v4/go/fusionauth/internal"
)

// ## # System Configuration Resource
//
// A registration is the association between a User and an Application that they log into.
//
// [System Configuration API](https://fusionauth.io/docs/v1/tech/apis/system)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/theogravity/pulumi-fusionauth/sdk/v4/go/fusionauth"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fusionauth.NewFusionAuthSystemConfiguration(ctx, "example", &fusionauth.FusionAuthSystemConfigurationArgs{
//				AuditLogConfiguration: &fusionauth.FusionAuthSystemConfigurationAuditLogConfigurationArgs{
//					Delete: &fusionauth.FusionAuthSystemConfigurationAuditLogConfigurationDeleteArgs{
//						Enabled:              pulumi.Bool(true),
//						NumberOfDaysToRetain: pulumi.Int(367),
//					},
//				},
//				CorsConfiguration: &fusionauth.FusionAuthSystemConfigurationCorsConfigurationArgs{
//					AllowedMethods: pulumi.StringArray{
//						pulumi.String("POST"),
//						pulumi.String("PUT"),
//					},
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
type FusionAuthSystemConfiguration struct {
	pulumi.CustomResourceState

	AuditLogConfiguration    FusionAuthSystemConfigurationAuditLogConfigurationOutput    `pulumi:"auditLogConfiguration"`
	CorsConfiguration        FusionAuthSystemConfigurationCorsConfigurationOutput        `pulumi:"corsConfiguration"`
	EventLogConfiguration    FusionAuthSystemConfigurationEventLogConfigurationOutput    `pulumi:"eventLogConfiguration"`
	LoginRecordConfiguration FusionAuthSystemConfigurationLoginRecordConfigurationOutput `pulumi:"loginRecordConfiguration"`
	// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
	ReportTimezone  pulumi.StringPtrOutput                             `pulumi:"reportTimezone"`
	UiConfiguration FusionAuthSystemConfigurationUiConfigurationOutput `pulumi:"uiConfiguration"`
}

// NewFusionAuthSystemConfiguration registers a new resource with the given unique name, arguments, and options.
func NewFusionAuthSystemConfiguration(ctx *pulumi.Context,
	name string, args *FusionAuthSystemConfigurationArgs, opts ...pulumi.ResourceOption) (*FusionAuthSystemConfiguration, error) {
	if args == nil {
		args = &FusionAuthSystemConfigurationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FusionAuthSystemConfiguration
	err := ctx.RegisterResource("fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFusionAuthSystemConfiguration gets an existing FusionAuthSystemConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFusionAuthSystemConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAuthSystemConfigurationState, opts ...pulumi.ResourceOption) (*FusionAuthSystemConfiguration, error) {
	var resource FusionAuthSystemConfiguration
	err := ctx.ReadResource("fusionauth:index/fusionAuthSystemConfiguration:FusionAuthSystemConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FusionAuthSystemConfiguration resources.
type fusionAuthSystemConfigurationState struct {
	AuditLogConfiguration    *FusionAuthSystemConfigurationAuditLogConfiguration    `pulumi:"auditLogConfiguration"`
	CorsConfiguration        *FusionAuthSystemConfigurationCorsConfiguration        `pulumi:"corsConfiguration"`
	EventLogConfiguration    *FusionAuthSystemConfigurationEventLogConfiguration    `pulumi:"eventLogConfiguration"`
	LoginRecordConfiguration *FusionAuthSystemConfigurationLoginRecordConfiguration `pulumi:"loginRecordConfiguration"`
	// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
	ReportTimezone  *string                                       `pulumi:"reportTimezone"`
	UiConfiguration *FusionAuthSystemConfigurationUiConfiguration `pulumi:"uiConfiguration"`
}

type FusionAuthSystemConfigurationState struct {
	AuditLogConfiguration    FusionAuthSystemConfigurationAuditLogConfigurationPtrInput
	CorsConfiguration        FusionAuthSystemConfigurationCorsConfigurationPtrInput
	EventLogConfiguration    FusionAuthSystemConfigurationEventLogConfigurationPtrInput
	LoginRecordConfiguration FusionAuthSystemConfigurationLoginRecordConfigurationPtrInput
	// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
	ReportTimezone  pulumi.StringPtrInput
	UiConfiguration FusionAuthSystemConfigurationUiConfigurationPtrInput
}

func (FusionAuthSystemConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthSystemConfigurationState)(nil)).Elem()
}

type fusionAuthSystemConfigurationArgs struct {
	AuditLogConfiguration    *FusionAuthSystemConfigurationAuditLogConfiguration    `pulumi:"auditLogConfiguration"`
	CorsConfiguration        *FusionAuthSystemConfigurationCorsConfiguration        `pulumi:"corsConfiguration"`
	EventLogConfiguration    *FusionAuthSystemConfigurationEventLogConfiguration    `pulumi:"eventLogConfiguration"`
	LoginRecordConfiguration *FusionAuthSystemConfigurationLoginRecordConfiguration `pulumi:"loginRecordConfiguration"`
	// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
	ReportTimezone  *string                                       `pulumi:"reportTimezone"`
	UiConfiguration *FusionAuthSystemConfigurationUiConfiguration `pulumi:"uiConfiguration"`
}

// The set of arguments for constructing a FusionAuthSystemConfiguration resource.
type FusionAuthSystemConfigurationArgs struct {
	AuditLogConfiguration    FusionAuthSystemConfigurationAuditLogConfigurationPtrInput
	CorsConfiguration        FusionAuthSystemConfigurationCorsConfigurationPtrInput
	EventLogConfiguration    FusionAuthSystemConfigurationEventLogConfigurationPtrInput
	LoginRecordConfiguration FusionAuthSystemConfigurationLoginRecordConfigurationPtrInput
	// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
	ReportTimezone  pulumi.StringPtrInput
	UiConfiguration FusionAuthSystemConfigurationUiConfigurationPtrInput
}

func (FusionAuthSystemConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAuthSystemConfigurationArgs)(nil)).Elem()
}

type FusionAuthSystemConfigurationInput interface {
	pulumi.Input

	ToFusionAuthSystemConfigurationOutput() FusionAuthSystemConfigurationOutput
	ToFusionAuthSystemConfigurationOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationOutput
}

func (*FusionAuthSystemConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthSystemConfiguration)(nil)).Elem()
}

func (i *FusionAuthSystemConfiguration) ToFusionAuthSystemConfigurationOutput() FusionAuthSystemConfigurationOutput {
	return i.ToFusionAuthSystemConfigurationOutputWithContext(context.Background())
}

func (i *FusionAuthSystemConfiguration) ToFusionAuthSystemConfigurationOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthSystemConfigurationOutput)
}

// FusionAuthSystemConfigurationArrayInput is an input type that accepts FusionAuthSystemConfigurationArray and FusionAuthSystemConfigurationArrayOutput values.
// You can construct a concrete instance of `FusionAuthSystemConfigurationArrayInput` via:
//
//	FusionAuthSystemConfigurationArray{ FusionAuthSystemConfigurationArgs{...} }
type FusionAuthSystemConfigurationArrayInput interface {
	pulumi.Input

	ToFusionAuthSystemConfigurationArrayOutput() FusionAuthSystemConfigurationArrayOutput
	ToFusionAuthSystemConfigurationArrayOutputWithContext(context.Context) FusionAuthSystemConfigurationArrayOutput
}

type FusionAuthSystemConfigurationArray []FusionAuthSystemConfigurationInput

func (FusionAuthSystemConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthSystemConfiguration)(nil)).Elem()
}

func (i FusionAuthSystemConfigurationArray) ToFusionAuthSystemConfigurationArrayOutput() FusionAuthSystemConfigurationArrayOutput {
	return i.ToFusionAuthSystemConfigurationArrayOutputWithContext(context.Background())
}

func (i FusionAuthSystemConfigurationArray) ToFusionAuthSystemConfigurationArrayOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthSystemConfigurationArrayOutput)
}

// FusionAuthSystemConfigurationMapInput is an input type that accepts FusionAuthSystemConfigurationMap and FusionAuthSystemConfigurationMapOutput values.
// You can construct a concrete instance of `FusionAuthSystemConfigurationMapInput` via:
//
//	FusionAuthSystemConfigurationMap{ "key": FusionAuthSystemConfigurationArgs{...} }
type FusionAuthSystemConfigurationMapInput interface {
	pulumi.Input

	ToFusionAuthSystemConfigurationMapOutput() FusionAuthSystemConfigurationMapOutput
	ToFusionAuthSystemConfigurationMapOutputWithContext(context.Context) FusionAuthSystemConfigurationMapOutput
}

type FusionAuthSystemConfigurationMap map[string]FusionAuthSystemConfigurationInput

func (FusionAuthSystemConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthSystemConfiguration)(nil)).Elem()
}

func (i FusionAuthSystemConfigurationMap) ToFusionAuthSystemConfigurationMapOutput() FusionAuthSystemConfigurationMapOutput {
	return i.ToFusionAuthSystemConfigurationMapOutputWithContext(context.Background())
}

func (i FusionAuthSystemConfigurationMap) ToFusionAuthSystemConfigurationMapOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAuthSystemConfigurationMapOutput)
}

type FusionAuthSystemConfigurationOutput struct{ *pulumi.OutputState }

func (FusionAuthSystemConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAuthSystemConfiguration)(nil)).Elem()
}

func (o FusionAuthSystemConfigurationOutput) ToFusionAuthSystemConfigurationOutput() FusionAuthSystemConfigurationOutput {
	return o
}

func (o FusionAuthSystemConfigurationOutput) ToFusionAuthSystemConfigurationOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationOutput {
	return o
}

func (o FusionAuthSystemConfigurationOutput) AuditLogConfiguration() FusionAuthSystemConfigurationAuditLogConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) FusionAuthSystemConfigurationAuditLogConfigurationOutput {
		return v.AuditLogConfiguration
	}).(FusionAuthSystemConfigurationAuditLogConfigurationOutput)
}

func (o FusionAuthSystemConfigurationOutput) CorsConfiguration() FusionAuthSystemConfigurationCorsConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) FusionAuthSystemConfigurationCorsConfigurationOutput {
		return v.CorsConfiguration
	}).(FusionAuthSystemConfigurationCorsConfigurationOutput)
}

func (o FusionAuthSystemConfigurationOutput) EventLogConfiguration() FusionAuthSystemConfigurationEventLogConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) FusionAuthSystemConfigurationEventLogConfigurationOutput {
		return v.EventLogConfiguration
	}).(FusionAuthSystemConfigurationEventLogConfigurationOutput)
}

func (o FusionAuthSystemConfigurationOutput) LoginRecordConfiguration() FusionAuthSystemConfigurationLoginRecordConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) FusionAuthSystemConfigurationLoginRecordConfigurationOutput {
		return v.LoginRecordConfiguration
	}).(FusionAuthSystemConfigurationLoginRecordConfigurationOutput)
}

// The time zone used to adjust the stored UTC time when generating reports. Since reports are usually rolled up hourly, this timezone will be used for demarcating the hours.
func (o FusionAuthSystemConfigurationOutput) ReportTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) pulumi.StringPtrOutput { return v.ReportTimezone }).(pulumi.StringPtrOutput)
}

func (o FusionAuthSystemConfigurationOutput) UiConfiguration() FusionAuthSystemConfigurationUiConfigurationOutput {
	return o.ApplyT(func(v *FusionAuthSystemConfiguration) FusionAuthSystemConfigurationUiConfigurationOutput {
		return v.UiConfiguration
	}).(FusionAuthSystemConfigurationUiConfigurationOutput)
}

type FusionAuthSystemConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FusionAuthSystemConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FusionAuthSystemConfiguration)(nil)).Elem()
}

func (o FusionAuthSystemConfigurationArrayOutput) ToFusionAuthSystemConfigurationArrayOutput() FusionAuthSystemConfigurationArrayOutput {
	return o
}

func (o FusionAuthSystemConfigurationArrayOutput) ToFusionAuthSystemConfigurationArrayOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationArrayOutput {
	return o
}

func (o FusionAuthSystemConfigurationArrayOutput) Index(i pulumi.IntInput) FusionAuthSystemConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FusionAuthSystemConfiguration {
		return vs[0].([]*FusionAuthSystemConfiguration)[vs[1].(int)]
	}).(FusionAuthSystemConfigurationOutput)
}

type FusionAuthSystemConfigurationMapOutput struct{ *pulumi.OutputState }

func (FusionAuthSystemConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FusionAuthSystemConfiguration)(nil)).Elem()
}

func (o FusionAuthSystemConfigurationMapOutput) ToFusionAuthSystemConfigurationMapOutput() FusionAuthSystemConfigurationMapOutput {
	return o
}

func (o FusionAuthSystemConfigurationMapOutput) ToFusionAuthSystemConfigurationMapOutputWithContext(ctx context.Context) FusionAuthSystemConfigurationMapOutput {
	return o
}

func (o FusionAuthSystemConfigurationMapOutput) MapIndex(k pulumi.StringInput) FusionAuthSystemConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FusionAuthSystemConfiguration {
		return vs[0].(map[string]*FusionAuthSystemConfiguration)[vs[1].(string)]
	}).(FusionAuthSystemConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthSystemConfigurationInput)(nil)).Elem(), &FusionAuthSystemConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthSystemConfigurationArrayInput)(nil)).Elem(), FusionAuthSystemConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FusionAuthSystemConfigurationMapInput)(nil)).Elem(), FusionAuthSystemConfigurationMap{})
	pulumi.RegisterOutputType(FusionAuthSystemConfigurationOutput{})
	pulumi.RegisterOutputType(FusionAuthSystemConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FusionAuthSystemConfigurationMapOutput{})
}
