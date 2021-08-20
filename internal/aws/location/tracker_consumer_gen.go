// Code generated by generators/resource/main.go; DO NOT EDIT.

package location

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_location_tracker_consumer", trackerConsumerResourceType)
}

// trackerConsumerResourceType returns the Terraform awscc_location_tracker_consumer resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Location::TrackerConsumer resource type.
func trackerConsumerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"consumer_arn": {
			// Property: ConsumerArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// ConsumerArn is a force-new attribute.
		},
		"tracker_name": {
			// Property: TrackerName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// TrackerName is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::Location::TrackerConsumer Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::TrackerConsumer").WithTerraformTypeName("awscc_location_tracker_consumer").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"consumer_arn": "ConsumerArn",
		"tracker_name": "TrackerName",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_location_tracker_consumer", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
