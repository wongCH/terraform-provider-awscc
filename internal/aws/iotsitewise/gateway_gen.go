// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotsitewise_gateway", gatewayResourceType)
}

// gatewayResourceType returns the Terraform awscc_iotsitewise_gateway resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Gateway resource type.
func gatewayResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"gateway_capability_summaries": {
			// Property: GatewayCapabilitySummaries
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of gateway capability summaries that each contain a namespace and status.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Contains a summary of a gateway capability configuration.",
			//     "properties": {
			//       "CapabilityConfiguration": {
			//         "description": "The JSON document that defines the gateway capability's configuration.",
			//         "type": "string"
			//       },
			//       "CapabilityNamespace": {
			//         "description": "The namespace of the capability configuration.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "CapabilityNamespace"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of gateway capability summaries that each contain a namespace and status.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"capability_configuration": {
						// Property: CapabilityConfiguration
						Description: "The JSON document that defines the gateway capability's configuration.",
						Type:        types.StringType,
						Optional:    true,
					},
					"capability_namespace": {
						// Property: CapabilityNamespace
						Description: "The namespace of the capability configuration.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Validators: []tfsdk.AttributeValidator{validate.UniqueItems()},
			Optional:   true,
		},
		"gateway_id": {
			// Property: GatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the gateway device.",
			//   "type": "string"
			// }
			Description: "The ID of the gateway device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"gateway_name": {
			// Property: GatewayName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique, friendly name for the gateway.",
			//   "type": "string"
			// }
			Description: "A unique, friendly name for the gateway.",
			Type:        types.StringType,
			Required:    true,
		},
		"gateway_platform": {
			// Property: GatewayPlatform
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Contains a gateway's platform information.",
			//   "properties": {
			//     "Greengrass": {
			//       "additionalProperties": false,
			//       "description": "Contains the ARN of AWS IoT Greengrass Group that the gateway runs on.",
			//       "properties": {
			//         "GroupArn": {
			//           "description": "The ARN of the Greengrass group.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "GroupArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Greengrass"
			//   ],
			//   "type": "object"
			// }
			Description: "Contains a gateway's platform information.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"greengrass": {
						// Property: Greengrass
						Description: "Contains the ARN of AWS IoT Greengrass Group that the gateway runs on.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_arn": {
									// Property: GroupArn
									Description: "The ARN of the Greengrass group.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
			// GatewayPlatform is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the gateway.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the gateway.",
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
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Resource schema for AWS::IoTSiteWise::Gateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Gateway").WithTerraformTypeName("awscc_iotsitewise_gateway").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"capability_configuration":     "CapabilityConfiguration",
		"capability_namespace":         "CapabilityNamespace",
		"gateway_capability_summaries": "GatewayCapabilitySummaries",
		"gateway_id":                   "GatewayId",
		"gateway_name":                 "GatewayName",
		"gateway_platform":             "GatewayPlatform",
		"greengrass":                   "Greengrass",
		"group_arn":                    "GroupArn",
		"key":                          "Key",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotsitewise_gateway", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
