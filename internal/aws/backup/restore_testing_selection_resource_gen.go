// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_backup_restore_testing_selection", restoreTestingSelectionResource)
}

// restoreTestingSelectionResource returns the Terraform awscc_backup_restore_testing_selection resource.
// This Terraform resource corresponds to the CloudFormation AWS::Backup::RestoreTestingSelection resource.
func restoreTestingSelectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: IamRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: ProtectedResourceArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"protected_resource_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Optional:   true,
			Computed:   true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProtectedResourceConditions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StringEquals": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "StringNotEquals": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"protected_resource_conditions": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StringEquals
				"string_equals": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
					CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
					Optional:   true,
					Computed:   true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StringNotEquals
				"string_not_equals": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
					CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
					Optional:   true,
					Computed:   true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProtectedResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"protected_resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestoreMetadataOverrides
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"restore_metadata_overrides": // Pattern: ""
		schema.MapAttribute{          /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestoreTestingPlanName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"restore_testing_plan_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestoreTestingSelectionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"restore_testing_selection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ValidationWindowHours
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"validation_window_hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Backup::RestoreTestingSelection",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::RestoreTestingSelection").WithTerraformTypeName("awscc_backup_restore_testing_selection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"iam_role_arn":                   "IamRoleArn",
		"key":                            "Key",
		"protected_resource_arns":        "ProtectedResourceArns",
		"protected_resource_conditions":  "ProtectedResourceConditions",
		"protected_resource_type":        "ProtectedResourceType",
		"restore_metadata_overrides":     "RestoreMetadataOverrides",
		"restore_testing_plan_name":      "RestoreTestingPlanName",
		"restore_testing_selection_name": "RestoreTestingSelectionName",
		"string_equals":                  "StringEquals",
		"string_not_equals":              "StringNotEquals",
		"validation_window_hours":        "ValidationWindowHours",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(5).WithDeleteTimeoutInMinutes(5)

	opts = opts.WithUpdateTimeoutInMinutes(5)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
