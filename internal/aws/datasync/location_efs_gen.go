// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_datasync_location_efs", locationEFSResourceType)
}

// locationEFSResourceType returns the Terraform awscc_datasync_location_efs resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationEFS resource type.
func locationEFSResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"ec_2_config": {
			// Property: Ec2Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The subnet and security group that DataSync uses to access target EFS file system.",
			//   "properties": {
			//     "SecurityGroupArns": {
			//       "description": "The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.",
			//       "insertionOrder": false,
			//       "items": {
			//         "maxLength": 128,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "SubnetArn": {
			//       "description": "The ARN of the subnet that DataSync uses to access the target EFS file system.",
			//       "maxLength": 128,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SecurityGroupArns",
			//     "SubnetArn"
			//   ],
			//   "type": "object"
			// }
			Description: "The subnet and security group that DataSync uses to access target EFS file system.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"security_group_arns": {
						// Property: SecurityGroupArns
						Description: "The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
					},
					"subnet_arn": {
						// Property: SubnetArn
						Description: "The ARN of the subnet that DataSync uses to access the target EFS file system.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
			// Ec2Config is a force-new attribute.
		},
		"efs_filesystem_arn": {
			// Property: EfsFilesystemArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the Amazon EFS file system.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the Amazon EFS file system.",
			Type:        types.StringType,
			Required:    true,
			// EfsFilesystemArn is a force-new attribute.
			// EfsFilesystemArn is a write-only attribute.
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the EFS location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the EFS location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Subdirectory is a force-new attribute.
			// Subdirectory is a write-only attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::DataSync::LocationEFS.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationEFS").WithTerraformTypeName("awscc_datasync_location_efs").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ec_2_config":         "Ec2Config",
		"efs_filesystem_arn":  "EfsFilesystemArn",
		"key":                 "Key",
		"location_arn":        "LocationArn",
		"location_uri":        "LocationUri",
		"security_group_arns": "SecurityGroupArns",
		"subdirectory":        "Subdirectory",
		"subnet_arn":          "SubnetArn",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EfsFilesystemArn",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_datasync_location_efs", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
