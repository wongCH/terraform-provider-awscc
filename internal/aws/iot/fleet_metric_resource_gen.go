// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_fleet_metric", fleetMetricResourceType)
}

// fleetMetricResourceType returns the Terraform awscc_iot_fleet_metric resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::FleetMetric resource type.
func fleetMetricResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"aggregation_field": {
			// Property: AggregationField
			// CloudFormation resource type schema:
			// {
			//   "description": "The aggregation field to perform aggregation and metric emission",
			//   "type": "string"
			// }
			Description: "The aggregation field to perform aggregation and metric emission",
			Type:        types.StringType,
			Optional:    true,
		},
		"aggregation_type": {
			// Property: AggregationType
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Aggregation types supported by Fleet Indexing",
			//   "properties": {
			//     "Name": {
			//       "description": "Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality",
			//       "type": "string"
			//     },
			//     "Values": {
			//       "description": "Fleet Indexing aggregation type values",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "Values"
			//   ],
			//   "type": "object"
			// }
			Description: "Aggregation types supported by Fleet Indexing",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality",
						Type:        types.StringType,
						Required:    true,
					},
					"values": {
						// Property: Values
						Description: "Fleet Indexing aggregation type values",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Optional: true,
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The creation date of a fleet metric",
			//   "type": "number"
			// }
			Description: "The creation date of a fleet metric",
			Type:        types.NumberType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of a fleet metric",
			//   "type": "string"
			// }
			Description: "The description of a fleet metric",
			Type:        types.StringType,
			Optional:    true,
		},
		"index_name": {
			// Property: IndexName
			// CloudFormation resource type schema:
			// {
			//   "description": "The index name of a fleet metric",
			//   "type": "string"
			// }
			Description: "The index name of a fleet metric",
			Type:        types.StringType,
			Optional:    true,
		},
		"last_modified_date": {
			// Property: LastModifiedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The last modified date of a fleet metric",
			//   "type": "number"
			// }
			Description: "The last modified date of a fleet metric",
			Type:        types.NumberType,
			Computed:    true,
		},
		"metric_arn": {
			// Property: MetricArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) of a fleet metric metric",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) of a fleet metric metric",
			Type:        types.StringType,
			Computed:    true,
		},
		"metric_name": {
			// Property: MetricName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the fleet metric",
			//   "type": "string"
			// }
			Description: "The name of the fleet metric",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"period": {
			// Property: Period
			// CloudFormation resource type schema:
			// {
			//   "description": "The period of metric emission in seconds",
			//   "type": "integer"
			// }
			Description: "The period of metric emission in seconds",
			Type:        types.NumberType,
			Optional:    true,
		},
		"query_string": {
			// Property: QueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The Fleet Indexing query used by a fleet metric",
			//   "type": "string"
			// }
			Description: "The Fleet Indexing query used by a fleet metric",
			Type:        types.StringType,
			Optional:    true,
		},
		"query_version": {
			// Property: QueryVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of a Fleet Indexing query used by a fleet metric",
			//   "type": "string"
			// }
			Description: "The version of a Fleet Indexing query used by a fleet metric",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The tag's value",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"unit": {
			// Property: Unit
			// CloudFormation resource type schema:
			// {
			//   "description": "The unit of data points emitted by a fleet metric",
			//   "type": "string"
			// }
			Description: "The unit of data points emitted by a fleet metric",
			Type:        types.StringType,
			Optional:    true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of a fleet metric",
			//   "type": "number"
			// }
			Description: "The version of a fleet metric",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An aggregated metric of certain devices in your fleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::FleetMetric").WithTerraformTypeName("awscc_iot_fleet_metric")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregation_field":  "AggregationField",
		"aggregation_type":   "AggregationType",
		"creation_date":      "CreationDate",
		"description":        "Description",
		"index_name":         "IndexName",
		"key":                "Key",
		"last_modified_date": "LastModifiedDate",
		"metric_arn":         "MetricArn",
		"metric_name":        "MetricName",
		"name":               "Name",
		"period":             "Period",
		"query_string":       "QueryString",
		"query_version":      "QueryVersion",
		"tags":               "Tags",
		"unit":               "Unit",
		"value":              "Value",
		"values":             "Values",
		"version":            "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
