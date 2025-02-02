// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	registry.AddResourceFactory("awscc_ssmcontacts_plan", planResource)
}

// planResource returns the Terraform awscc_ssmcontacts_plan resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSMContacts::Plan resource.
func planResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the contact.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContactId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Contact ID for the AWS SSM Incident Manager Contact to associate the plan.",
		//	  "pattern": "arn:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]*:[0-9]+:([\\w+=\\/,.@:-]+)*",
		//	  "type": "string"
		//	}
		"contact_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Contact ID for the AWS SSM Incident Manager Contact to associate the plan.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]+:[-\\w+=\\/,.@]*:[0-9]+:([\\w+=\\/,.@:-]+)*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RotationIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rotation Ids to associate with Oncall Contact for engagement.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"rotation_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "Rotation Ids to associate with Oncall Contact for engagement.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// RotationIds is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Stages
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.",
		//	    "properties": {
		//	      "DurationInMinutes": {
		//	        "description": "The time to wait until beginning the next stage.",
		//	        "type": "integer"
		//	      },
		//	      "Targets": {
		//	        "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
		//	          "oneOf": [
		//	            {
		//	              "required": [
		//	                "ChannelTargetInfo"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "ContactTargetInfo"
		//	              ]
		//	            }
		//	          ],
		//	          "properties": {
		//	            "ChannelTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
		//	              "properties": {
		//	                "ChannelId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact channel.",
		//	                  "type": "string"
		//	                },
		//	                "RetryIntervalInMinutes": {
		//	                  "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "ChannelId",
		//	                "RetryIntervalInMinutes"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "ContactTargetInfo": {
		//	              "additionalProperties": false,
		//	              "description": "The contact that SSM Incident Manager is engaging during an incident.",
		//	              "properties": {
		//	                "ContactId": {
		//	                  "description": "The Amazon Resource Name (ARN) of the contact.",
		//	                  "type": "string"
		//	                },
		//	                "IsEssential": {
		//	                  "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "required": [
		//	                "ContactId",
		//	                "IsEssential"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "required": [
		//	      "DurationInMinutes"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"stages": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DurationInMinutes
					"duration_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The time to wait until beginning the next stage.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Targets
					"targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ChannelTargetInfo
								"channel_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ChannelId
										"channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact channel.",
											Required:    true,
										}, /*END ATTRIBUTE*/
										// Property: RetryIntervalInMinutes
										"retry_interval_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Description: "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
											Required:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ContactTargetInfo
								"contact_target_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ContactId
										"contact_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Amazon Resource Name (ARN) of the contact.",
											Required:    true,
										}, /*END ATTRIBUTE*/
										// Property: IsEssential
										"is_essential": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Description: "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
											Required:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The contact that SSM Incident Manager is engaging during an incident.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
						Description: "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
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
		Description: "Engagement Plan for a SSM Incident Manager Contact.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::Plan").WithTerraformTypeName("awscc_ssmcontacts_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"channel_id":                "ChannelId",
		"channel_target_info":       "ChannelTargetInfo",
		"contact_id":                "ContactId",
		"contact_target_info":       "ContactTargetInfo",
		"duration_in_minutes":       "DurationInMinutes",
		"is_essential":              "IsEssential",
		"retry_interval_in_minutes": "RetryIntervalInMinutes",
		"rotation_ids":              "RotationIds",
		"stages":                    "Stages",
		"targets":                   "Targets",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RotationIds",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
