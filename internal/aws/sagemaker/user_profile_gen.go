// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sagemaker_user_profile", userProfileResourceType)
}

// userProfileResourceType returns the Terraform awscc_sagemaker_user_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::UserProfile resource type.
func userProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"domain_id": {
			// Property: DomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the associated Domain.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ID of the associated Domain.",
			Type:        types.StringType,
			Required:    true,
			// DomainId is a force-new attribute.
		},
		"single_sign_on_user_identifier": {
			// Property: SingleSignOnUserIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is \"UserName\". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is \"UserName\". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// SingleSignOnUserIdentifier is a force-new attribute.
		},
		"single_sign_on_user_value": {
			// Property: SingleSignOnUserValue
			// CloudFormation resource type schema:
			// {
			//   "description": "The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// SingleSignOnUserValue is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the user profile.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of tags to apply to the user profile.",
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
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
			// Tags is a write-only attribute.
		},
		"user_profile_arn": {
			// Property: UserProfileArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The user profile Amazon Resource Name (ARN).",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The user profile Amazon Resource Name (ARN).",
			Type:        types.StringType,
			Computed:    true,
		},
		"user_profile_name": {
			// Property: UserProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the UserProfile.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A name for the UserProfile.",
			Type:        types.StringType,
			Required:    true,
			// UserProfileName is a force-new attribute.
		},
		"user_settings": {
			// Property: UserSettings
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the CreateUserProfile API is called, and as DefaultUserSettings when the CreateDomain API is called.",
			//   "properties": {
			//     "ExecutionRole": {
			//       "description": "The user profile Amazon Resource Name (ARN).",
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "JupyterServerAppSettings": {
			//       "additionalProperties": false,
			//       "description": "The JupyterServer app settings.",
			//       "properties": {
			//         "DefaultResourceSpec": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "InstanceType": {
			//               "description": "The instance type that the image version runs on.",
			//               "enum": [
			//                 "system",
			//                 "ml.t3.micro",
			//                 "ml.t3.small",
			//                 "ml.t3.medium",
			//                 "ml.t3.large",
			//                 "ml.t3.xlarge",
			//                 "ml.t3.2xlarge",
			//                 "ml.m5.large",
			//                 "ml.m5.xlarge",
			//                 "ml.m5.2xlarge",
			//                 "ml.m5.4xlarge",
			//                 "ml.m5.8xlarge",
			//                 "ml.m5.12xlarge",
			//                 "ml.m5.16xlarge",
			//                 "ml.m5.24xlarge",
			//                 "ml.c5.large",
			//                 "ml.c5.xlarge",
			//                 "ml.c5.2xlarge",
			//                 "ml.c5.4xlarge",
			//                 "ml.c5.9xlarge",
			//                 "ml.c5.12xlarge",
			//                 "ml.c5.18xlarge",
			//                 "ml.c5.24xlarge",
			//                 "ml.p3.2xlarge",
			//                 "ml.p3.8xlarge",
			//                 "ml.p3.16xlarge",
			//                 "ml.g4dn.xlarge",
			//                 "ml.g4dn.2xlarge",
			//                 "ml.g4dn.4xlarge",
			//                 "ml.g4dn.8xlarge",
			//                 "ml.g4dn.12xlarge",
			//                 "ml.g4dn.16xlarge"
			//               ],
			//               "type": "string"
			//             },
			//             "SageMakerImageArn": {
			//               "description": "The ARN of the SageMaker image that the image version belongs to.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "SageMakerImageVersionArn": {
			//               "description": "The ARN of the image version created on the instance.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "KernelGatewayAppSettings": {
			//       "additionalProperties": false,
			//       "description": "The kernel gateway app settings.",
			//       "properties": {
			//         "CustomImages": {
			//           "description": "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "A custom SageMaker image.",
			//             "properties": {
			//               "AppImageConfigName": {
			//                 "description": "The Name of the AppImageConfig.",
			//                 "maxLength": 63,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "ImageName": {
			//                 "description": "The name of the CustomImage. Must be unique to your account.",
			//                 "maxLength": 63,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "ImageVersionNumber": {
			//                 "description": "The version number of the CustomImage.",
			//                 "type": "integer"
			//               }
			//             },
			//             "required": [
			//               "AppImageConfigName",
			//               "ImageName"
			//             ],
			//             "type": "object"
			//           },
			//           "maxItems": 30,
			//           "minItems": 0,
			//           "type": "array",
			//           "uniqueItems": false
			//         },
			//         "DefaultResourceSpec": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "InstanceType": {
			//               "description": "The instance type that the image version runs on.",
			//               "enum": [
			//                 "system",
			//                 "ml.t3.micro",
			//                 "ml.t3.small",
			//                 "ml.t3.medium",
			//                 "ml.t3.large",
			//                 "ml.t3.xlarge",
			//                 "ml.t3.2xlarge",
			//                 "ml.m5.large",
			//                 "ml.m5.xlarge",
			//                 "ml.m5.2xlarge",
			//                 "ml.m5.4xlarge",
			//                 "ml.m5.8xlarge",
			//                 "ml.m5.12xlarge",
			//                 "ml.m5.16xlarge",
			//                 "ml.m5.24xlarge",
			//                 "ml.c5.large",
			//                 "ml.c5.xlarge",
			//                 "ml.c5.2xlarge",
			//                 "ml.c5.4xlarge",
			//                 "ml.c5.9xlarge",
			//                 "ml.c5.12xlarge",
			//                 "ml.c5.18xlarge",
			//                 "ml.c5.24xlarge",
			//                 "ml.p3.2xlarge",
			//                 "ml.p3.8xlarge",
			//                 "ml.p3.16xlarge",
			//                 "ml.g4dn.xlarge",
			//                 "ml.g4dn.2xlarge",
			//                 "ml.g4dn.4xlarge",
			//                 "ml.g4dn.8xlarge",
			//                 "ml.g4dn.12xlarge",
			//                 "ml.g4dn.16xlarge"
			//               ],
			//               "type": "string"
			//             },
			//             "SageMakerImageArn": {
			//               "description": "The ARN of the SageMaker image that the image version belongs to.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "SageMakerImageVersionArn": {
			//               "description": "The ARN of the image version created on the instance.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "SecurityGroups": {
			//       "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
			//       "items": {
			//         "maxLength": 32,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "minItems": 0,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SharingSettings": {
			//       "additionalProperties": false,
			//       "description": "Specifies options when sharing an Amazon SageMaker Studio notebook. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called.",
			//       "properties": {
			//         "NotebookOutputOption": {
			//           "description": "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
			//           "enum": [
			//             "Allowed",
			//             "Disabled"
			//           ],
			//           "type": "string"
			//         },
			//         "S3KmsKeyId": {
			//           "description": "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
			//           "maxLength": 2048,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "S3OutputPath": {
			//           "description": "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
			//           "maxLength": 1024,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the CreateUserProfile API is called, and as DefaultUserSettings when the CreateDomain API is called.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"execution_role": {
						// Property: ExecutionRole
						Description: "The user profile Amazon Resource Name (ARN).",
						Type:        types.StringType,
						Optional:    true,
					},
					"jupyter_server_app_settings": {
						// Property: JupyterServerAppSettings
						Description: "The JupyterServer app settings.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"default_resource_spec": {
									// Property: DefaultResourceSpec
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"instance_type": {
												// Property: InstanceType
												Description: "The instance type that the image version runs on.",
												Type:        types.StringType,
												Optional:    true,
											},
											"sage_maker_image_arn": {
												// Property: SageMakerImageArn
												Description: "The ARN of the SageMaker image that the image version belongs to.",
												Type:        types.StringType,
												Optional:    true,
											},
											"sage_maker_image_version_arn": {
												// Property: SageMakerImageVersionArn
												Description: "The ARN of the image version created on the instance.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"kernel_gateway_app_settings": {
						// Property: KernelGatewayAppSettings
						Description: "The kernel gateway app settings.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"custom_images": {
									// Property: CustomImages
									Description: "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"app_image_config_name": {
												// Property: AppImageConfigName
												Description: "The Name of the AppImageConfig.",
												Type:        types.StringType,
												Required:    true,
											},
											"image_name": {
												// Property: ImageName
												Description: "The name of the CustomImage. Must be unique to your account.",
												Type:        types.StringType,
												Required:    true,
											},
											"image_version_number": {
												// Property: ImageVersionNumber
												Description: "The version number of the CustomImage.",
												Type:        types.NumberType,
												Optional:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{
											MinItems: 0,
											MaxItems: 30,
										},
									),
									Optional: true,
								},
								"default_resource_spec": {
									// Property: DefaultResourceSpec
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"instance_type": {
												// Property: InstanceType
												Description: "The instance type that the image version runs on.",
												Type:        types.StringType,
												Optional:    true,
											},
											"sage_maker_image_arn": {
												// Property: SageMakerImageArn
												Description: "The ARN of the SageMaker image that the image version belongs to.",
												Type:        types.StringType,
												Optional:    true,
											},
											"sage_maker_image_version_arn": {
												// Property: SageMakerImageVersionArn
												Description: "The ARN of the image version created on the instance.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"security_groups": {
						// Property: SecurityGroups
						Description: "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"sharing_settings": {
						// Property: SharingSettings
						Description: "Specifies options when sharing an Amazon SageMaker Studio notebook. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"notebook_output_option": {
									// Property: NotebookOutputOption
									Description: "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
									Type:        types.StringType,
									Optional:    true,
								},
								"s3_kms_key_id": {
									// Property: S3KmsKeyId
									Description: "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
									Type:        types.StringType,
									Optional:    true,
								},
								"s3_output_path": {
									// Property: S3OutputPath
									Description: "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
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
		Description: "Resource Type definition for AWS::SageMaker::UserProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::UserProfile").WithTerraformTypeName("awscc_sagemaker_user_profile").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_name":          "AppImageConfigName",
		"custom_images":                  "CustomImages",
		"default_resource_spec":          "DefaultResourceSpec",
		"domain_id":                      "DomainId",
		"execution_role":                 "ExecutionRole",
		"image_name":                     "ImageName",
		"image_version_number":           "ImageVersionNumber",
		"instance_type":                  "InstanceType",
		"jupyter_server_app_settings":    "JupyterServerAppSettings",
		"kernel_gateway_app_settings":    "KernelGatewayAppSettings",
		"key":                            "Key",
		"notebook_output_option":         "NotebookOutputOption",
		"s3_kms_key_id":                  "S3KmsKeyId",
		"s3_output_path":                 "S3OutputPath",
		"sage_maker_image_arn":           "SageMakerImageArn",
		"sage_maker_image_version_arn":   "SageMakerImageVersionArn",
		"security_groups":                "SecurityGroups",
		"sharing_settings":               "SharingSettings",
		"single_sign_on_user_identifier": "SingleSignOnUserIdentifier",
		"single_sign_on_user_value":      "SingleSignOnUserValue",
		"tags":                           "Tags",
		"user_profile_arn":               "UserProfileArn",
		"user_profile_name":              "UserProfileName",
		"user_settings":                  "UserSettings",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_sagemaker_user_profile", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
