// Code generated by pulumi-gen-awsconf DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/t0yv0/pulumi-12709/sdk/go/awsconf/internal"
)

func ConfigureProvider(ctx *pulumi.Context, args *ConfigureProviderArgs, opts ...pulumi.InvokeOption) (*ConfigureProviderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ConfigureProviderResult
	err := ctx.Invoke("awsconf:index:ConfigureProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ConfigureProviderArgs struct {
	Profile *string `pulumi:"profile"`
	Region  *string `pulumi:"region"`
}

type ConfigureProviderResult struct {
	AwsProvider *aws.Provider `pulumi:"awsProvider"`
}

func ConfigureProviderOutput(ctx *pulumi.Context, args ConfigureProviderOutputArgs, opts ...pulumi.InvokeOption) ConfigureProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConfigureProviderResult, error) {
			args := v.(ConfigureProviderArgs)
			r, err := ConfigureProvider(ctx, &args, opts...)
			var s ConfigureProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConfigureProviderResultOutput)
}

type ConfigureProviderOutputArgs struct {
	Profile pulumi.StringPtrInput `pulumi:"profile"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (ConfigureProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigureProviderArgs)(nil)).Elem()
}

type ConfigureProviderResultOutput struct{ *pulumi.OutputState }

func (ConfigureProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigureProviderResult)(nil)).Elem()
}

func (o ConfigureProviderResultOutput) ToConfigureProviderResultOutput() ConfigureProviderResultOutput {
	return o
}

func (o ConfigureProviderResultOutput) ToConfigureProviderResultOutputWithContext(ctx context.Context) ConfigureProviderResultOutput {
	return o
}

func (o ConfigureProviderResultOutput) AwsProvider() aws.ProviderOutput {
	return o.ApplyT(func(v ConfigureProviderResult) *aws.Provider { return v.AwsProvider }).(aws.ProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigureProviderResultOutput{})
}
