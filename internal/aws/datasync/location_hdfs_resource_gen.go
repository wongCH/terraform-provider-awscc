// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
	registry.AddResourceFactory("awscc_datasync_location_hdfs", locationHDFSResource)
}

// locationHDFSResource returns the Terraform awscc_datasync_location_hdfs resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationHDFS resource.
func locationHDFSResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN(s) of the agent(s) to use for an HDFS location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "ARN(s) of the agent(s) to use for an HDFS location.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 4),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(128),
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$"), ""),
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthenticationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authentication mode used to determine identity of user.",
		//	  "enum": [
		//	    "SIMPLE",
		//	    "KERBEROS"
		//	  ],
		//	  "type": "string"
		//	}
		"authentication_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The authentication mode used to determine identity of user.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SIMPLE",
					"KERBEROS",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: BlockSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.",
		//	  "format": "int64",
		//	  "maximum": 1073741824,
		//	  "minimum": 1048576,
		//	  "type": "integer"
		//	}
		"block_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1048576, 1073741824),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KerberosKeytab
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Base64 string representation of the Keytab file.",
		//	  "maxLength": 87384,
		//	  "type": "string"
		//	}
		"kerberos_keytab": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Base64 string representation of the Keytab file.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(87384),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// KerberosKeytab is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: KerberosKrb5Conf
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.",
		//	  "maxLength": 174764,
		//	  "type": "string"
		//	}
		"kerberos_krb_5_conf": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(174764),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// KerberosKrb5Conf is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: KerberosPrincipal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identity, or principal, to which Kerberos can assign tickets.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"kerberos_principal": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identity, or principal, to which Kerberos can assign tickets.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyProviderUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^kms:\\/\\/http[s]?@(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])(;(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9]))*:[0-9]{1,5}\\/kms$",
		//	  "type": "string"
		//	}
		"kms_key_provider_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^kms:\\/\\/http[s]?@(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])(;(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9]))*:[0-9]{1,5}\\/kms$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the HDFS location.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the HDFS location.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the HDFS location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw|hdfs)://[a-zA-Z0-9.:/\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the HDFS location that was described.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NameNodes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of Name Node(s) of the HDFS location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "HDFS Name Node IP and port information.",
		//	    "properties": {
		//	      "Hostname": {
		//	        "description": "The DNS name or IP address of the Name Node in the customer's on premises HDFS cluster.",
		//	        "maxLength": 255,
		//	        "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
		//	        "type": "string"
		//	      },
		//	      "Port": {
		//	        "description": "The port on which the Name Node is listening on for client requests.",
		//	        "maximum": 65536,
		//	        "minimum": 1,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "Hostname",
		//	      "Port"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"name_nodes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Hostname
					"hostname": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The DNS name or IP address of the Name Node in the customer's on premises HDFS cluster.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile("^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Port
					"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The port on which the Name Node is listening on for client requests.",
						Required:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(1, 65536),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "An array of Name Node(s) of the HDFS location.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: QopConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration information for RPC Protection and Data Transfer Protection. These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY.",
		//	  "properties": {
		//	    "DataTransferProtection": {
		//	      "default": "PRIVACY",
		//	      "description": "Configuration for Data Transfer Protection.",
		//	      "enum": [
		//	        "AUTHENTICATION",
		//	        "INTEGRITY",
		//	        "PRIVACY",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "RpcProtection": {
		//	      "default": "PRIVACY",
		//	      "description": "Configuration for RPC Protection.",
		//	      "enum": [
		//	        "AUTHENTICATION",
		//	        "INTEGRITY",
		//	        "PRIVACY",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"qop_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataTransferProtection
				"data_transfer_protection": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Configuration for Data Transfer Protection.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"AUTHENTICATION",
							"INTEGRITY",
							"PRIVACY",
							"DISABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						generic.StringDefaultValue("PRIVACY"),
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RpcProtection
				"rpc_protection": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Configuration for RPC Protection.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"AUTHENTICATION",
							"INTEGRITY",
							"PRIVACY",
							"DISABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						generic.StringDefaultValue("PRIVACY"),
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration information for RPC Protection and Data Transfer Protection. These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationFactor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 3,
		//	  "description": "Number of copies of each block that exists inside the HDFS cluster.",
		//	  "format": "int64",
		//	  "maximum": 512,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"replication_factor": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Number of copies of each block that exists inside the HDFS cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 512),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				generic.Int64DefaultValue(3),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SimpleUser
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user name that has read and write permissions on the specified HDFS cluster.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[_.A-Za-z0-9][-_.A-Za-z0-9]*$",
		//	  "type": "string"
		//	}
		"simple_user": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user name that has read and write permissions on the specified HDFS cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[_.A-Za-z0-9][-_.A-Za-z0-9]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(4096),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Subdirectory is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
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
		//	  "maxItems": 50,
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
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::DataSync::LocationHDFS.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationHDFS").WithTerraformTypeName("awscc_datasync_location_hdfs")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":               "AgentArns",
		"authentication_type":      "AuthenticationType",
		"block_size":               "BlockSize",
		"data_transfer_protection": "DataTransferProtection",
		"hostname":                 "Hostname",
		"kerberos_keytab":          "KerberosKeytab",
		"kerberos_krb_5_conf":      "KerberosKrb5Conf",
		"kerberos_principal":       "KerberosPrincipal",
		"key":                      "Key",
		"kms_key_provider_uri":     "KmsKeyProviderUri",
		"location_arn":             "LocationArn",
		"location_uri":             "LocationUri",
		"name_nodes":               "NameNodes",
		"port":                     "Port",
		"qop_configuration":        "QopConfiguration",
		"replication_factor":       "ReplicationFactor",
		"rpc_protection":           "RpcProtection",
		"simple_user":              "SimpleUser",
		"subdirectory":             "Subdirectory",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Subdirectory",
		"/properties/KerberosKeytab",
		"/properties/KerberosKrb5Conf",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
