// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_backup_backup_selection", backupSelectionDataSourceType)
}

// backupSelectionDataSourceType returns the Terraform awscc_backup_backup_selection data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Backup::BackupSelection resource type.
func backupSelectionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"backup_plan_id": {
			// Property: BackupPlanId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"backup_selection": {
			// Property: BackupSelection
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "IamRoleArn": {
			//       "type": "string"
			//     },
			//     "ListOfTags": {
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ConditionKey": {
			//             "type": "string"
			//           },
			//           "ConditionType": {
			//             "type": "string"
			//           },
			//           "ConditionValue": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ConditionValue",
			//           "ConditionKey",
			//           "ConditionType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "Resources": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SelectionName": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SelectionName",
			//     "IamRoleArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"iam_role_arn": {
						// Property: IamRoleArn
						Type:     types.StringType,
						Computed: true,
					},
					"list_of_tags": {
						// Property: ListOfTags
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"condition_key": {
									// Property: ConditionKey
									Type:     types.StringType,
									Computed: true,
								},
								"condition_type": {
									// Property: ConditionType
									Type:     types.StringType,
									Computed: true,
								},
								"condition_value": {
									// Property: ConditionValue
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"resources": {
						// Property: Resources
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"selection_name": {
						// Property: SelectionName
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"selection_id": {
			// Property: SelectionId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Backup::BackupSelection",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupSelection").WithTerraformTypeName("awscc_backup_backup_selection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"backup_plan_id":   "BackupPlanId",
		"backup_selection": "BackupSelection",
		"condition_key":    "ConditionKey",
		"condition_type":   "ConditionType",
		"condition_value":  "ConditionValue",
		"iam_role_arn":     "IamRoleArn",
		"id":               "Id",
		"list_of_tags":     "ListOfTags",
		"resources":        "Resources",
		"selection_id":     "SelectionId",
		"selection_name":   "SelectionName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
