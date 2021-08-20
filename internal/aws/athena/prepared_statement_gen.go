// Code generated by generators/resource/main.go; DO NOT EDIT.

package athena

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
	registry.AddResourceTypeFactory("awscc_athena_prepared_statement", preparedStatementResourceType)
}

// preparedStatementResourceType returns the Terraform awscc_athena_prepared_statement resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Athena::PreparedStatement resource type.
func preparedStatementResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the prepared statement.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description of the prepared statement.",
			Type:        types.StringType,
			Optional:    true,
		},
		"query_statement": {
			// Property: QueryStatement
			// CloudFormation resource type schema:
			// {
			//   "description": "The query string for the prepared statement.",
			//   "maxLength": 262144,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query string for the prepared statement.",
			Type:        types.StringType,
			Required:    true,
		},
		"statement_name": {
			// Property: StatementName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the prepared statement.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the prepared statement.",
			Type:        types.StringType,
			Required:    true,
			// StatementName is a force-new attribute.
		},
		"work_group": {
			// Property: WorkGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the workgroup to which the prepared statement belongs.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the workgroup to which the prepared statement belongs.",
			Type:        types.StringType,
			Required:    true,
			// WorkGroup is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Athena::PreparedStatement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::PreparedStatement").WithTerraformTypeName("awscc_athena_prepared_statement").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":     "Description",
		"query_statement": "QueryStatement",
		"statement_name":  "StatementName",
		"work_group":      "WorkGroup",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_athena_prepared_statement", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
