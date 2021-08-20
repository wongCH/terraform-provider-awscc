// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("awscc_ec2_prefix_list", prefixListResourceType)
}

// prefixListResourceType returns the Terraform awscc_ec2_prefix_list resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::PrefixList resource type.
func prefixListResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"address_family": {
			// Property: AddressFamily
			// CloudFormation resource type schema:
			// {
			//   "description": "Ip Version of Prefix List.",
			//   "enum": [
			//     "IPv4",
			//     "IPv6"
			//   ],
			//   "type": "string"
			// }
			Description: "Ip Version of Prefix List.",
			Type:        types.StringType,
			Required:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Prefix List.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"entries": {
			// Property: Entries
			// CloudFormation resource type schema:
			// {
			//   "description": "Entries of Prefix List.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Cidr": {
			//         "maxLength": 46,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Description": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Cidr"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Entries of Prefix List.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"cidr": {
						// Property: Cidr
						Type:     types.StringType,
						Required: true,
					},
					"description": {
						// Property: Description
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"max_entries": {
			// Property: MaxEntries
			// CloudFormation resource type schema:
			// {
			//   "description": "Max Entries of Prefix List.",
			//   "type": "integer"
			// }
			Description: "Max Entries of Prefix List.",
			Type:        types.NumberType,
			Required:    true,
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "Owner Id of Prefix List.",
			//   "type": "string"
			// }
			Description: "Owner Id of Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"prefix_list_id": {
			// Property: PrefixListId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of Prefix List.",
			//   "type": "string"
			// }
			Description: "Id of Prefix List.",
			Type:        types.StringType,
			Computed:    true,
		},
		"prefix_list_name": {
			// Property: PrefixListName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Prefix List.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of Prefix List.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags for Prefix List",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Tags for Prefix List",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "Version of Prefix List.",
			//   "type": "integer"
			// }
			Description: "Version of Prefix List.",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema of AWS::EC2::PrefixList Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::PrefixList").WithTerraformTypeName("awscc_ec2_prefix_list").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_family":   "AddressFamily",
		"arn":              "Arn",
		"cidr":             "Cidr",
		"description":      "Description",
		"entries":          "Entries",
		"key":              "Key",
		"max_entries":      "MaxEntries",
		"owner_id":         "OwnerId",
		"prefix_list_id":   "PrefixListId",
		"prefix_list_name": "PrefixListName",
		"tags":             "Tags",
		"value":            "Value",
		"version":          "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_prefix_list", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
