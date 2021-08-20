// Code generated by generators/resource/main.go; DO NOT EDIT.

package cassandra

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cassandra_table", tableResourceType)
}

// tableResourceType returns the Terraform awscc_cassandra_table resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Cassandra::Table resource type.
func tableResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"billing_mode": {
			// Property: BillingMode
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Mode": {
			//       "description": "Capacity mode for the specified table",
			//       "enum": [
			//         "PROVISIONED",
			//         "ON_DEMAND"
			//       ],
			//       "type": "string"
			//     },
			//     "ProvisionedThroughput": {
			//       "additionalProperties": false,
			//       "description": "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
			//       "properties": {
			//         "ReadCapacityUnits": {
			//           "type": "integer"
			//         },
			//         "WriteCapacityUnits": {
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "ReadCapacityUnits",
			//         "WriteCapacityUnits"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Mode"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"mode": {
						// Property: Mode
						Description: "Capacity mode for the specified table",
						Type:        types.StringType,
						Required:    true,
					},
					"provisioned_throughput": {
						// Property: ProvisionedThroughput
						Description: "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"read_capacity_units": {
									// Property: ReadCapacityUnits
									Type:     types.NumberType,
									Required: true,
								},
								"write_capacity_units": {
									// Property: WriteCapacityUnits
									Type:     types.NumberType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"clustering_key_columns": {
			// Property: ClusteringKeyColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Clustering key columns of the table",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Column": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ColumnName": {
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "ColumnType": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ColumnName",
			//           "ColumnType"
			//         ],
			//         "type": "object"
			//       },
			//       "OrderBy": {
			//         "enum": [
			//           "ASC",
			//           "DESC"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Column"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Clustering key columns of the table",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"column": {
						// Property: Column
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"column_name": {
									// Property: ColumnName
									Type:     types.StringType,
									Required: true,
								},
								"column_type": {
									// Property: ColumnType
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"order_by": {
						// Property: OrderBy
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Validators: []tfsdk.AttributeValidator{validate.UniqueItems()},
			Optional:   true,
			Computed:   true,
			// ClusteringKeyColumns is a force-new attribute.
		},
		"encryption_specification": {
			// Property: EncryptionSpecification
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Represents the settings used to enable server-side encryption",
			//   "properties": {
			//     "EncryptionType": {
			//       "description": "Server-side encryption type",
			//       "enum": [
			//         "AWS_OWNED_KMS_KEY",
			//         "CUSTOMER_MANAGED_KMS_KEY"
			//       ],
			//       "type": "string"
			//     },
			//     "KmsKeyIdentifier": {
			//       "description": "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "EncryptionType"
			//   ],
			//   "type": "object"
			// }
			Description: "Represents the settings used to enable server-side encryption",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "Server-side encryption type",
						Type:        types.StringType,
						Required:    true,
					},
					"kms_key_identifier": {
						// Property: KmsKeyIdentifier
						Description: "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"keyspace_name": {
			// Property: KeyspaceName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for Cassandra keyspace",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name for Cassandra keyspace",
			Type:        types.StringType,
			Required:    true,
			// KeyspaceName is a force-new attribute.
		},
		"partition_key_columns": {
			// Property: PartitionKeyColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Partition key columns of the table",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ColumnName": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ColumnType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ColumnName",
			//       "ColumnType"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Partition key columns of the table",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"column_name": {
						// Property: ColumnName
						Type:     types.StringType,
						Required: true,
					},
					"column_type": {
						// Property: ColumnType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Validators: []tfsdk.AttributeValidator{validate.UniqueItems()},
			Required:   true,
			// PartitionKeyColumns is a force-new attribute.
		},
		"point_in_time_recovery_enabled": {
			// Property: PointInTimeRecoveryEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
			//   "type": "boolean"
			// }
			Description: "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
			Type:        types.BoolType,
			Optional:    true,
		},
		"regular_columns": {
			// Property: RegularColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Non-key columns of the table",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ColumnName": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ColumnType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ColumnName",
			//       "ColumnType"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Non-key columns of the table",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"column_name": {
						// Property: ColumnName
						Type:     types.StringType,
						Required: true,
					},
					"column_type": {
						// Property: ColumnType
						Type:     types.StringType,
						Required: true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for Cassandra table",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name for Cassandra table",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// TableName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to apply to the resource",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource",
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
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Validators: []tfsdk.AttributeValidator{validate.UniqueItems()},
			Optional:   true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Cassandra::Table",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Cassandra::Table").WithTerraformTypeName("awscc_cassandra_table").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"billing_mode":                   "BillingMode",
		"clustering_key_columns":         "ClusteringKeyColumns",
		"column":                         "Column",
		"column_name":                    "ColumnName",
		"column_type":                    "ColumnType",
		"encryption_specification":       "EncryptionSpecification",
		"encryption_type":                "EncryptionType",
		"key":                            "Key",
		"keyspace_name":                  "KeyspaceName",
		"kms_key_identifier":             "KmsKeyIdentifier",
		"mode":                           "Mode",
		"order_by":                       "OrderBy",
		"partition_key_columns":          "PartitionKeyColumns",
		"point_in_time_recovery_enabled": "PointInTimeRecoveryEnabled",
		"provisioned_throughput":         "ProvisionedThroughput",
		"read_capacity_units":            "ReadCapacityUnits",
		"regular_columns":                "RegularColumns",
		"table_name":                     "TableName",
		"tags":                           "Tags",
		"value":                          "Value",
		"write_capacity_units":           "WriteCapacityUnits",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cassandra_table", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
