// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
	registry.AddResourceFactory("awscc_rds_db_parameter_group", dBParameterGroupResource)
}

// dBParameterGroupResource returns the Terraform awscc_rds_db_parameter_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::DBParameterGroup resource.
func dBParameterGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DBParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the DB parameter group.\n Constraints:\n  +  Must be 1 to 255 letters, numbers, or hyphens.\n  +  First character must be a letter\n  +  Can't end with a hyphen or contain two consecutive hyphens\n  \n If you don't specify a value for ``DBParameterGroupName`` property, a name is automatically created for the DB parameter group.\n  This value is stored as a lowercase string.",
		//	  "pattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"db_parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the DB parameter group.\n Constraints:\n  +  Must be 1 to 255 letters, numbers, or hyphens.\n  +  First character must be a letter\n  +  Can't end with a hyphen or contain two consecutive hyphens\n  \n If you don't specify a value for ``DBParameterGroupName`` property, a name is automatically created for the DB parameter group.\n  This value is stored as a lowercase string.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]{1}(?:-?[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides the customer-specified description for this DB parameter group.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Provides the customer-specified description for this DB parameter group.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Family
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a DB engine and engine version compatible with that DB parameter group family.\n  The DB parameter group family can't be changed when updating a DB parameter group.\n  To list all of the available parameter group families, use the following command:\n ``aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\"``\n The output contains duplicates.\n For more information, see ``CreateDBParameterGroup``.",
		//	  "type": "string"
		//	}
		"family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a DB engine and engine version compatible with that DB parameter group family.\n  The DB parameter group family can't be changed when updating a DB parameter group.\n  To list all of the available parameter group families, use the following command:\n ``aws rds describe-db-engine-versions --query \"DBEngineVersions[].DBParameterGroupFamily\"``\n The output contains duplicates.\n For more information, see ``CreateDBParameterGroup``.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of parameter names and values for the parameter update. At least one parameter name and value must be supplied. Subsequent arguments are optional.\n RDS for Db2 requires you to bring your own Db2 license. You must enter your IBM customer ID (``rds.ibm_customer_id``) and site number (``rds.ibm_site_id``) before starting a Db2 instance.\n For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.\n For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  AWS CloudFormation doesn't support specifying an apply method for each individual ",
		//	  "type": "object"
		//	}
		"parameters": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "An array of parameter names and values for the parameter update. At least one parameter name and value must be supplied. Subsequent arguments are optional.\n RDS for Db2 requires you to bring your own Db2 license. You must enter your IBM customer ID (``rds.ibm_customer_id``) and site number (``rds.ibm_site_id``) before starting a Db2 instance.\n For more information about DB parameters and DB parameter groups for Amazon RDS DB engines, see [Working with DB Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.\n For more information about DB cluster and DB instance parameters and parameter groups for Amazon Aurora DB engines, see [Working with DB Parameter Groups and DB Cluster Parameter Groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  AWS CloudFormation doesn't support specifying an apply method for each individual ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional array of key-value pairs to apply to this DB parameter group.\n  Currently, this is the only property that supports drift detection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata assigned to an Amazon RDS resource consisting of a key-value pair.\n For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "An optional array of key-value pairs to apply to this DB parameter group.\n  Currently, this is the only property that supports drift detection.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
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
		Description: "The ``AWS::RDS::DBParameterGroup`` resource creates a custom parameter group for an RDS database family.\n This type can be declared in a template and referenced in the ``DBParameterGroupName`` property of an ``AWS::RDS::DBInstance`` resource.\n For information about configuring parameters for Amazon RDS DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.\n For information about configuring parameters for Amazon Aurora DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.\n  Applying a parameter group to a DB instance may require the DB instance to reboot, resulting in a database outage for the duration of the reboot.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBParameterGroup").WithTerraformTypeName("awscc_rds_db_parameter_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"db_parameter_group_name": "DBParameterGroupName",
		"description":             "Description",
		"family":                  "Family",
		"key":                     "Key",
		"parameters":              "Parameters",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
