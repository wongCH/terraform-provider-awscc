// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

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
	registry.AddResourceTypeFactory("awscc_glue_schema_version_metadata", schemaVersionMetadataResourceType)
}

// schemaVersionMetadataResourceType returns the Terraform awscc_glue_schema_version_metadata resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Glue::SchemaVersionMetadata resource type.
func schemaVersionMetadataResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"key": {
			// Property: Key
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata key",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Metadata key",
			Type:        types.StringType,
			Required:    true,
			// Key is a force-new attribute.
		},
		"schema_version_id": {
			// Property: SchemaVersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "Represents the version ID associated with the schema version.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Represents the version ID associated with the schema version.",
			Type:        types.StringType,
			Required:    true,
			// SchemaVersionId is a force-new attribute.
		},
		"value": {
			// Property: Value
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata value",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Metadata value",
			Type:        types.StringType,
			Required:    true,
			// Value is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "This resource adds Key-Value metadata to a Schema version of Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::SchemaVersionMetadata").WithTerraformTypeName("awscc_glue_schema_version_metadata").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":               "Key",
		"schema_version_id": "SchemaVersionId",
		"value":             "Value",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_glue_schema_version_metadata", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
