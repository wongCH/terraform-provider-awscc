// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

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
	registry.AddResourceTypeFactory("awscc_macie_session", sessionResourceType)
}

// sessionResourceType returns the Terraform awscc_macie_session resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Macie::Session resource type.
func sessionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS account ID of customer",
			//   "type": "string"
			// }
			Description: "AWS account ID of customer",
			Type:        types.StringType,
			Computed:    true,
		},
		"finding_publishing_frequency": {
			// Property: FindingPublishingFrequency
			// CloudFormation resource type schema:
			// {
			//   "description": "A enumeration value that specifies how frequently finding updates are published.",
			//   "enum": [
			//     "FIFTEEN_MINUTES",
			//     "ONE_HOUR",
			//     "SIX_HOURS"
			//   ],
			//   "type": "string"
			// }
			Description: "A enumeration value that specifies how frequently finding updates are published.",
			Type:        types.StringType,
			Optional:    true,
		},
		"service_role": {
			// Property: ServiceRole
			// CloudFormation resource type schema:
			// {
			//   "description": "Service role used by Macie",
			//   "type": "string"
			// }
			Description: "Service role used by Macie",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "A enumeration value that specifies the status of the Macie Session.",
			//   "enum": [
			//     "ENABLED",
			//     "PAUSED"
			//   ],
			//   "type": "string"
			// }
			Description: "A enumeration value that specifies the status of the Macie Session.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::Session").WithTerraformTypeName("awscc_macie_session").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_account_id":               "AwsAccountId",
		"finding_publishing_frequency": "FindingPublishingFrequency",
		"service_role":                 "ServiceRole",
		"status":                       "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_macie_session", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
