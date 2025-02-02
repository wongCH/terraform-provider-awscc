// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	registry.AddResourceFactory("awscc_rds_db_proxy", dBProxyResource)
}

// dBProxyResource returns the Terraform awscc_rds_db_proxy resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::DBProxy resource.
func dBProxyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Auth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorization mechanism that the proxy uses.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AuthScheme": {
		//	        "description": "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
		//	        "enum": [
		//	          "SECRETS"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ClientPasswordAuthType": {
		//	        "description": "The type of authentication the proxy uses for connections from clients.",
		//	        "enum": [
		//	          "MYSQL_NATIVE_PASSWORD",
		//	          "POSTGRES_SCRAM_SHA_256",
		//	          "POSTGRES_MD5",
		//	          "SQL_SERVER_AUTHENTICATION"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
		//	        "type": "string"
		//	      },
		//	      "IAMAuth": {
		//	        "description": "Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.",
		//	        "enum": [
		//	          "DISABLED",
		//	          "REQUIRED",
		//	          "ENABLED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "SecretArn": {
		//	        "description": "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"auth": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AuthScheme
					"auth_scheme": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"SECRETS",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ClientPasswordAuthType
					"client_password_auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of authentication the proxy uses for connections from clients.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"MYSQL_NATIVE_PASSWORD",
								"POSTGRES_SCRAM_SHA_256",
								"POSTGRES_MD5",
								"SQL_SERVER_AUTHENTICATION",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: IAMAuth
					"iam_auth": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"DISABLED",
								"REQUIRED",
								"ENABLED",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SecretArn
					"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The authorization mechanism that the proxy uses.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DBProxyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the proxy.",
		//	  "type": "string"
		//	}
		"db_proxy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the proxy.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBProxyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-z]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DebugLogging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether the proxy includes detailed information about SQL statements in its logs.",
		//	  "type": "boolean"
		//	}
		"debug_logging": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the proxy includes detailed information about SQL statements in its logs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The kinds of databases that the proxy can connect to.",
		//	  "enum": [
		//	    "MYSQL",
		//	    "POSTGRESQL",
		//	    "SQLSERVER"
		//	  ],
		//	  "type": "string"
		//	}
		"engine_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The kinds of databases that the proxy can connect to.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"MYSQL",
					"POSTGRESQL",
					"SQLSERVER",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdleClientTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
		//	  "type": "integer"
		//	}
		"idle_client_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RequireTLS
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
		//	  "type": "boolean"
		//	}
		"require_tls": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(128),
							stringvalidator.RegexMatches(regexp.MustCompile("(\\w|\\d|\\s|\\\\|-|\\.:=+-)*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(128),
							stringvalidator.RegexMatches(regexp.MustCompile("(\\w|\\d|\\s|\\\\|-|\\.:=+-)*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC ID to associate with the new DB proxy.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "VPC ID to associate with the new DB proxy.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC security group IDs to associate with the new proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "VPC security group IDs to associate with the new proxy.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC subnet IDs to associate with the new proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array"
		//	}
		"vpc_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "VPC subnet IDs to associate with the new proxy.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(2),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.RequiresReplace(),
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
		Description: "Resource schema for AWS::RDS::DBProxy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxy").WithTerraformTypeName("awscc_rds_db_proxy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auth":                      "Auth",
		"auth_scheme":               "AuthScheme",
		"client_password_auth_type": "ClientPasswordAuthType",
		"db_proxy_arn":              "DBProxyArn",
		"db_proxy_name":             "DBProxyName",
		"debug_logging":             "DebugLogging",
		"description":               "Description",
		"endpoint":                  "Endpoint",
		"engine_family":             "EngineFamily",
		"iam_auth":                  "IAMAuth",
		"idle_client_timeout":       "IdleClientTimeout",
		"key":                       "Key",
		"require_tls":               "RequireTLS",
		"role_arn":                  "RoleArn",
		"secret_arn":                "SecretArn",
		"tags":                      "Tags",
		"value":                     "Value",
		"vpc_id":                    "VpcId",
		"vpc_security_group_ids":    "VpcSecurityGroupIds",
		"vpc_subnet_ids":            "VpcSubnetIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
