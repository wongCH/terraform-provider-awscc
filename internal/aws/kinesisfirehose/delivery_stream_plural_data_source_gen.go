// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package kinesisfirehose

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kinesisfirehose_delivery_streams", deliveryStreamsDataSourceType)
}

// deliveryStreamsDataSourceType returns the Terraform awscc_kinesisfirehose_delivery_streams data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::KinesisFirehose::DeliveryStream resource type.
func deliveryStreamsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::KinesisFirehose::DeliveryStream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KinesisFirehose::DeliveryStream").WithTerraformTypeName("awscc_kinesisfirehose_delivery_streams")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
