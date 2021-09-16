// Code generated by generators/resource/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_macie_findings_filter", findingsFilterResourceType)
}

// findingsFilterResourceType returns the Terraform awscc_macie_findings_filter resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Macie::FindingsFilter resource type.
func findingsFilterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ARCHIVE",
			//     "NOOP"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ARCHIVE",
					"NOOP",
				}),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ARN.",
			//   "type": "string"
			// }
			Description: "Findings filter ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter description",
			//   "type": "string"
			// }
			Description: "Findings filter description",
			Type:        types.StringType,
			Optional:    true,
		},
		"finding_criteria": {
			// Property: FindingCriteria
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "Criterion": {
			//       "description": "Map of filter criteria.",
			//       "patternProperties": {
			//         "": {
			//           "properties": {
			//             "Eq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "Gt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "Gte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "Lt": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "Lte": {
			//               "format": "int64",
			//               "type": "integer"
			//             },
			//             "Neq": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"criterion": {
						// Property: Criterion
						Description: "Map of filter criteria.",
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"eq": {
									// Property: Eq
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"gt": {
									// Property: Gt
									Type:     types.NumberType,
									Optional: true,
								},
								"gte": {
									// Property: Gte
									Type:     types.NumberType,
									Optional: true,
								},
								"lt": {
									// Property: Lt
									Type:     types.NumberType,
									Optional: true,
								},
								"lte": {
									// Property: Lte
									Type:     types.NumberType,
									Optional: true,
								},
								"neq": {
									// Property: Neq
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
							tfsdk.MapNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"findings_filter_list_items": {
			// Property: FindingsFilterListItems
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filters list.",
			//   "items": {
			//     "description": "Returned by ListHandler representing filter name and ID.",
			//     "properties": {
			//       "Id": {
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Findings filters list.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Type:     types.StringType,
						Optional: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter ID.",
			//   "type": "string"
			// }
			Description: "Findings filter ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter name",
			//   "type": "string"
			// }
			Description: "Findings filter name",
			Type:        types.StringType,
			Required:    true,
		},
		"position": {
			// Property: Position
			// CloudFormation resource type schema:
			// {
			//   "description": "Findings filter position.",
			//   "type": "integer"
			// }
			Description: "Findings filter position.",
			Type:        types.NumberType,
			Optional:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Macie FindingsFilter resource schema.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::FindingsFilter").WithTerraformTypeName("awscc_macie_findings_filter")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                     "Action",
		"arn":                        "Arn",
		"criterion":                  "Criterion",
		"description":                "Description",
		"eq":                         "Eq",
		"finding_criteria":           "FindingCriteria",
		"findings_filter_list_items": "FindingsFilterListItems",
		"gt":                         "Gt",
		"gte":                        "Gte",
		"id":                         "Id",
		"lt":                         "Lt",
		"lte":                        "Lte",
		"name":                       "Name",
		"neq":                        "Neq",
		"position":                   "Position",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
