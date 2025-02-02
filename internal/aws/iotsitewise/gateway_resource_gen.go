// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_iotsitewise_gateway", gatewayResource)
}

// gatewayResource returns the Terraform awscc_iotsitewise_gateway resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTSiteWise::Gateway resource.
func gatewayResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GatewayCapabilitySummaries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of gateway capability summaries that each contain a namespace and status.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Contains a summary of a gateway capability configuration.",
		//	    "properties": {
		//	      "CapabilityConfiguration": {
		//	        "description": "The JSON document that defines the gateway capability's configuration.",
		//	        "type": "string"
		//	      },
		//	      "CapabilityNamespace": {
		//	        "description": "The namespace of the capability configuration.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "CapabilityNamespace"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"gateway_capability_summaries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CapabilityConfiguration
					"capability_configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The JSON document that defines the gateway capability's configuration.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CapabilityNamespace
					"capability_namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The namespace of the capability configuration.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of gateway capability summaries that each contain a namespace and status.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the gateway device.",
		//	  "type": "string"
		//	}
		"gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the gateway device.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique, friendly name for the gateway.",
		//	  "type": "string"
		//	}
		"gateway_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique, friendly name for the gateway.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: GatewayPlatform
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The gateway's platform. You can only specify one platform in a gateway.",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "Greengrass"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "GreengrassV2"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "Greengrass": {
		//	      "additionalProperties": false,
		//	      "description": "A gateway that runs on AWS IoT Greengrass V1.",
		//	      "properties": {
		//	        "GroupArn": {
		//	          "description": "The ARN of the Greengrass group.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "GroupArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "GreengrassV2": {
		//	      "additionalProperties": false,
		//	      "description": "A gateway that runs on AWS IoT Greengrass V2.",
		//	      "properties": {
		//	        "CoreDeviceThingName": {
		//	          "description": "The name of the CoreDevice in GreenGrass V2.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CoreDeviceThingName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"gateway_platform": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Greengrass
				"greengrass": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: GroupArn
						"group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the Greengrass group.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A gateway that runs on AWS IoT Greengrass V1.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: GreengrassV2
				"greengrass_v2": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CoreDeviceThingName
						"core_device_thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the CoreDevice in GreenGrass V2.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A gateway that runs on AWS IoT Greengrass V2.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The gateway's platform. You can only specify one platform in a gateway.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the gateway.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "A list of key-value pairs that contain metadata for the gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::IoTSiteWise::Gateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Gateway").WithTerraformTypeName("awscc_iotsitewise_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"capability_configuration":     "CapabilityConfiguration",
		"capability_namespace":         "CapabilityNamespace",
		"core_device_thing_name":       "CoreDeviceThingName",
		"gateway_capability_summaries": "GatewayCapabilitySummaries",
		"gateway_id":                   "GatewayId",
		"gateway_name":                 "GatewayName",
		"gateway_platform":             "GatewayPlatform",
		"greengrass":                   "Greengrass",
		"greengrass_v2":                "GreengrassV2",
		"group_arn":                    "GroupArn",
		"key":                          "Key",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
