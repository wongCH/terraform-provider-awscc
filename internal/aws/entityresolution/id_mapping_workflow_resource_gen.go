// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package entityresolution

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_entityresolution_id_mapping_workflow", idMappingWorkflowResource)
}

// idMappingWorkflowResource returns the Terraform awscc_entityresolution_id_mapping_workflow resource.
// This Terraform resource corresponds to the CloudFormation AWS::EntityResolution::IdMappingWorkflow resource.
func idMappingWorkflowResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this IdMappingWorkflow got created",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this IdMappingWorkflow got created",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the IdMappingWorkflow",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the IdMappingWorkflow",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdMappingTechniques
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "IdMappingType": {
		//	      "enum": [
		//	        "PROVIDER"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ProviderProperties": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "IntermediateSourceConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "IntermediateS3Path": {
		//	              "description": "The s3 path that would be used to stage the intermediate data being generated during workflow execution.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "IntermediateS3Path"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ProviderConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ProviderServiceArn": {
		//	          "description": "Arn of the Provider Service being used.",
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:([A-Za-z0-9]+(-[A-Za-z0-9]+)+)::providerservice/[A-Za-z0-9]+/[A-Za-z0-9]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ProviderServiceArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"id_mapping_techniques": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IdMappingType
				"id_mapping_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"PROVIDER",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ProviderProperties
				"provider_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IntermediateSourceConfiguration
						"intermediate_source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IntermediateS3Path
								"intermediate_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The s3 path that would be used to stage the intermediate data being generated during workflow execution.",
									Required:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ProviderConfiguration
						"provider_configuration": // Pattern: ""
						schema.MapAttribute{      /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ProviderServiceArn
						"provider_service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Arn of the Provider Service being used.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-us-gov|aws-cn):entityresolution:([A-Za-z0-9]+(-[A-Za-z0-9]+)+)::providerservice/[A-Za-z0-9]+/[A-Za-z0-9]+$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: InputSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "InputSourceARN": {
		//	        "description": "An Glue table ARN for the input source table",
		//	        "pattern": "arn:(aws|aws-us-gov|aws-cn):.*:.*:[0-9]+:.*$",
		//	        "type": "string"
		//	      },
		//	      "SchemaArn": {
		//	        "description": "The SchemaMapping arn associated with the Schema",
		//	        "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "InputSourceARN",
		//	      "SchemaArn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"input_source_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InputSourceARN
					"input_source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "An Glue table ARN for the input source table",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):.*:.*:[0-9]+:.*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: SchemaArn
					"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The SchemaMapping arn associated with the Schema",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Required:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 20),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: OutputSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "KMSArn": {
		//	        "pattern": "^arn:(aws|aws-us-gov|aws-cn):kms:.*:[0-9]+:.*$",
		//	        "type": "string"
		//	      },
		//	      "OutputS3Path": {
		//	        "description": "The S3 path to which Entity Resolution will write the output table",
		//	        "pattern": "^s3://([^/]+)/?(.*?([^/]+)/?)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "OutputS3Path"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"output_source_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: KMSArn
					"kms_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-us-gov|aws-cn):kms:.*:[0-9]+:.*$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: OutputS3Path
					"output_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The S3 path to which Entity Resolution will write the output table",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^s3://([^/]+)/?(.*?([^/]+)/?)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
			Required:   true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-us-gov|aws-cn):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this IdMappingWorkflow got last updated at",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this IdMappingWorkflow got last updated at",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkflowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default IdMappingWorkflow arn",
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(idmappingworkflow/.*)$",
		//	  "type": "string"
		//	}
		"workflow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default IdMappingWorkflow arn",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkflowName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the IdMappingWorkflow",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"workflow_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the IdMappingWorkflow",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_0-9-]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "IdMappingWorkflow defined in AWS Entity Resolution service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EntityResolution::IdMappingWorkflow").WithTerraformTypeName("awscc_entityresolution_id_mapping_workflow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":                        "CreatedAt",
		"description":                       "Description",
		"id_mapping_techniques":             "IdMappingTechniques",
		"id_mapping_type":                   "IdMappingType",
		"input_source_arn":                  "InputSourceARN",
		"input_source_config":               "InputSourceConfig",
		"intermediate_s3_path":              "IntermediateS3Path",
		"intermediate_source_configuration": "IntermediateSourceConfiguration",
		"key":                               "Key",
		"kms_arn":                           "KMSArn",
		"output_s3_path":                    "OutputS3Path",
		"output_source_config":              "OutputSourceConfig",
		"provider_configuration":            "ProviderConfiguration",
		"provider_properties":               "ProviderProperties",
		"provider_service_arn":              "ProviderServiceArn",
		"role_arn":                          "RoleArn",
		"schema_arn":                        "SchemaArn",
		"tags":                              "Tags",
		"updated_at":                        "UpdatedAt",
		"value":                             "Value",
		"workflow_arn":                      "WorkflowArn",
		"workflow_name":                     "WorkflowName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
