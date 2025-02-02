// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package shield

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_shield_protection_group", protectionGroupResource)
}

// protectionGroupResource returns the Terraform awscc_shield_protection_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::Shield::ProtectionGroup resource.
func protectionGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Aggregation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.\n* Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.\n* Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.\n* Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront and origin resources for CloudFront distributions.",
		//	  "enum": [
		//	    "SUM",
		//	    "MEAN",
		//	    "MAX"
		//	  ],
		//	  "type": "string"
		//	}
		"aggregation": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.\n* Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.\n* Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.\n* Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront and origin resources for CloudFront distributions.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SUM",
					"MEAN",
					"MAX",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Members
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of the resources to include in the protection group. You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 10000,
		//	  "type": "array"
		//	}
		"members": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The Amazon Resource Names (ARNs) of the resources to include in the protection group. You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(10000),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 2048),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Pattern
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type.",
		//	  "enum": [
		//	    "ALL",
		//	    "ARBITRARY",
		//	    "BY_RESOURCE_TYPE"
		//	  ],
		//	  "type": "string"
		//	}
		"pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ALL",
					"ARBITRARY",
					"BY_RESOURCE_TYPE",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ProtectionGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon Resource Name) of the protection group.",
		//	  "type": "string"
		//	}
		"protection_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon Resource Name) of the protection group.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProtectionGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the protection group. You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.",
		//	  "maxLength": 36,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9\\-]*",
		//	  "type": "string"
		//	}
		"protection_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the protection group. You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9\\-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource type to include in the protection group. All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.",
		//	  "enum": [
		//	    "CLOUDFRONT_DISTRIBUTION",
		//	    "ROUTE_53_HOSTED_ZONE",
		//	    "ELASTIC_IP_ALLOCATION",
		//	    "CLASSIC_LOAD_BALANCER",
		//	    "APPLICATION_LOAD_BALANCER",
		//	    "GLOBAL_ACCELERATOR"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource type to include in the protection group. All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CLOUDFRONT_DISTRIBUTION",
					"ROUTE_53_HOSTED_ZONE",
					"ELASTIC_IP_ALLOCATION",
					"CLASSIC_LOAD_BALANCER",
					"APPLICATION_LOAD_BALANCER",
					"GLOBAL_ACCELERATOR",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tag key-value pairs for the Protection object.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as \"environment\", and the tag value represents a specific value within that category, such as \"test,\" \"development,\" or \"production\". Or you might set the tag key to \"customer\" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "One or more tag key-value pairs for the Protection object.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
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
		Description: "A grouping of protected resources so they can be handled as a collective. This resource grouping improves the accuracy of detection and reduces false positives.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Shield::ProtectionGroup").WithTerraformTypeName("awscc_shield_protection_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregation":          "Aggregation",
		"key":                  "Key",
		"members":              "Members",
		"pattern":              "Pattern",
		"protection_group_arn": "ProtectionGroupArn",
		"protection_group_id":  "ProtectionGroupId",
		"resource_type":        "ResourceType",
		"tags":                 "Tags",
		"value":                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
