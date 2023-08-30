// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_backup_backup_plan", backupPlanResource)
}

// backupPlanResource returns the Terraform awscc_backup_backup_plan resource.
// This Terraform resource corresponds to the CloudFormation AWS::Backup::BackupPlan resource.
func backupPlanResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BackupPlan
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AdvancedBackupSettings": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "BackupOptions": {
		//	            "type": "object"
		//	          },
		//	          "ResourceType": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "BackupOptions",
		//	          "ResourceType"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "BackupPlanName": {
		//	      "type": "string"
		//	    },
		//	    "BackupPlanRule": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "CompletionWindowMinutes": {
		//	            "type": "number"
		//	          },
		//	          "CopyActions": {
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "DestinationBackupVaultArn": {
		//	                  "type": "string"
		//	                },
		//	                "Lifecycle": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "DeleteAfterDays": {
		//	                      "type": "number"
		//	                    },
		//	                    "MoveToColdStorageAfterDays": {
		//	                      "type": "number"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "required": [
		//	                "DestinationBackupVaultArn"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "EnableContinuousBackup": {
		//	            "type": "boolean"
		//	          },
		//	          "Lifecycle": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DeleteAfterDays": {
		//	                "type": "number"
		//	              },
		//	              "MoveToColdStorageAfterDays": {
		//	                "type": "number"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "RecoveryPointTags": {
		//	            "additionalProperties": false,
		//	            "patternProperties": {
		//	              "": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "RuleName": {
		//	            "type": "string"
		//	          },
		//	          "ScheduleExpression": {
		//	            "type": "string"
		//	          },
		//	          "ScheduleExpressionTimezone": {
		//	            "type": "string"
		//	          },
		//	          "StartWindowMinutes": {
		//	            "type": "number"
		//	          },
		//	          "TargetBackupVault": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "TargetBackupVault",
		//	          "RuleName"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "required": [
		//	    "BackupPlanName",
		//	    "BackupPlanRule"
		//	  ],
		//	  "type": "object"
		//	}
		"backup_plan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AdvancedBackupSettings
				"advanced_backup_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: BackupOptions
							"backup_options": schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: ResourceType
							"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: BackupPlanName
				"backup_plan_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: BackupPlanRule
				"backup_plan_rule": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CompletionWindowMinutes
							"completion_window_minutes": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: CopyActions
							"copy_actions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DestinationBackupVaultArn
										"destination_backup_vault_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
											Required: true,
										}, /*END ATTRIBUTE*/
										// Property: Lifecycle
										"lifecycle": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: DeleteAfterDays
												"delete_after_days": schema.Float64Attribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
														float64planmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: MoveToColdStorageAfterDays
												"move_to_cold_storage_after_days": schema.Float64Attribute{ /*START ATTRIBUTE*/
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
														float64planmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: EnableContinuousBackup
							"enable_continuous_backup": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Lifecycle
							"lifecycle": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DeleteAfterDays
									"delete_after_days": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
											float64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: MoveToColdStorageAfterDays
									"move_to_cold_storage_after_days": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
											float64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: RecoveryPointTags
							"recovery_point_tags": // Pattern: ""
							schema.MapAttribute{   /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
									mapplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: RuleName
							"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: ScheduleExpression
							"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ScheduleExpressionTimezone
							"schedule_expression_timezone": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: StartWindowMinutes
							"start_window_minutes": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
									float64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TargetBackupVault
							"target_backup_vault": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: BackupPlanArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"backup_plan_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BackupPlanId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"backup_plan_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BackupPlanTags
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
		"backup_plan_tags":  // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Backup::BackupPlan",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupPlan").WithTerraformTypeName("awscc_backup_backup_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"advanced_backup_settings":        "AdvancedBackupSettings",
		"backup_options":                  "BackupOptions",
		"backup_plan":                     "BackupPlan",
		"backup_plan_arn":                 "BackupPlanArn",
		"backup_plan_id":                  "BackupPlanId",
		"backup_plan_name":                "BackupPlanName",
		"backup_plan_rule":                "BackupPlanRule",
		"backup_plan_tags":                "BackupPlanTags",
		"completion_window_minutes":       "CompletionWindowMinutes",
		"copy_actions":                    "CopyActions",
		"delete_after_days":               "DeleteAfterDays",
		"destination_backup_vault_arn":    "DestinationBackupVaultArn",
		"enable_continuous_backup":        "EnableContinuousBackup",
		"lifecycle":                       "Lifecycle",
		"move_to_cold_storage_after_days": "MoveToColdStorageAfterDays",
		"recovery_point_tags":             "RecoveryPointTags",
		"resource_type":                   "ResourceType",
		"rule_name":                       "RuleName",
		"schedule_expression":             "ScheduleExpression",
		"schedule_expression_timezone":    "ScheduleExpressionTimezone",
		"start_window_minutes":            "StartWindowMinutes",
		"target_backup_vault":             "TargetBackupVault",
		"version_id":                      "VersionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
