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
	registry.AddResourceTypeFactory("awscc_sagemaker_project", projectResourceType)
}

// projectResourceType returns the Terraform awscc_sagemaker_project resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Project resource type.
func projectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time at which the project was created.",
			//   "type": "string"
			// }
			Description: "The time at which the project was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_arn": {
			// Property: ProjectArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Project.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_description": {
			// Property: ProjectDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the project.",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The description of the project.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ProjectDescription is a force-new attribute.
		},
		"project_id": {
			// Property: ProjectId
			// CloudFormation resource type schema:
			// {
			//   "description": "Project Id.",
			//   "maxLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Project Id.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the project.",
			//   "maxLength": 32,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the project.",
			Type:        types.StringType,
			Required:    true,
			// ProjectName is a force-new attribute.
		},
		"project_status": {
			// Property: ProjectStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of a project.",
			//   "enum": [
			//     "Pending",
			//     "CreateInProgress",
			//     "CreateCompleted",
			//     "CreateFailed",
			//     "DeleteInProgress",
			//     "DeleteFailed",
			//     "DeleteCompleted"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of a project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_catalog_provisioned_product_details": {
			// Property: ServiceCatalogProvisionedProductDetails
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Provisioned ServiceCatalog  Details",
			//   "properties": {
			//     "ProvisionedProductId": {
			//       "description": "The identifier of the provisioning artifact (also known as a version).",
			//       "maxLength": 100,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "ProvisionedProductStatusMessage": {
			//       "description": "Provisioned Product Status Message",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Provisioned ServiceCatalog  Details",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"provisioned_product_id": {
						// Property: ProvisionedProductId
						Description: "The identifier of the provisioning artifact (also known as a version).",
						Type:        types.StringType,
						Optional:    true,
					},
					"provisioned_product_status_message": {
						// Property: ProvisionedProductStatusMessage
						Description: "Provisioned Product Status Message",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Computed: true,
		},
		"service_catalog_provisioning_details": {
			// Property: ServiceCatalogProvisioningDetails
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Input ServiceCatalog Provisioning Details",
			//   "properties": {
			//     "PathId": {
			//       "description": "The path identifier of the product.",
			//       "maxLength": 100,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "ProductId": {
			//       "description": "Service Catalog product identifier.",
			//       "maxLength": 100,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "ProvisioningArtifactId": {
			//       "description": "The identifier of the provisioning artifact (also known as a version).",
			//       "maxLength": 100,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "ProvisioningParameters": {
			//       "description": "Parameters specified by the administrator that are required for provisioning the product.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Information about a parameter used to provision a product.",
			//         "properties": {
			//           "Key": {
			//             "description": "The parameter key.",
			//             "maxLength": 1000,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "description": "The parameter value.",
			//             "maxLength": 4096,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Key",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "ProductId",
			//     "ProvisioningArtifactId"
			//   ],
			//   "type": "object"
			// }
			Description: "Input ServiceCatalog Provisioning Details",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"path_id": {
						// Property: PathId
						Description: "The path identifier of the product.",
						Type:        types.StringType,
						Optional:    true,
					},
					"product_id": {
						// Property: ProductId
						Description: "Service Catalog product identifier.",
						Type:        types.StringType,
						Required:    true,
					},
					"provisioning_artifact_id": {
						// Property: ProvisioningArtifactId
						Description: "The identifier of the provisioning artifact (also known as a version).",
						Type:        types.StringType,
						Required:    true,
					},
					"provisioning_parameters": {
						// Property: ProvisioningParameters
						Description: "Parameters specified by the administrator that are required for provisioning the product.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "The parameter key.",
									Type:        types.StringType,
									Required:    true,
								},
								"value": {
									// Property: Value
									Description: "The parameter value.",
									Type:        types.StringType,
									Required:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Required: true,
			// ServiceCatalogProvisioningDetails is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 40,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 40,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SageMaker::Project",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Project").WithTerraformTypeName("awscc_sagemaker_project").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                      "CreationTime",
		"key":                                "Key",
		"path_id":                            "PathId",
		"product_id":                         "ProductId",
		"project_arn":                        "ProjectArn",
		"project_description":                "ProjectDescription",
		"project_id":                         "ProjectId",
		"project_name":                       "ProjectName",
		"project_status":                     "ProjectStatus",
		"provisioned_product_id":             "ProvisionedProductId",
		"provisioned_product_status_message": "ProvisionedProductStatusMessage",
		"provisioning_artifact_id":           "ProvisioningArtifactId",
		"provisioning_parameters":            "ProvisioningParameters",
		"service_catalog_provisioned_product_details": "ServiceCatalogProvisionedProductDetails",
		"service_catalog_provisioning_details":        "ServiceCatalogProvisioningDetails",
		"tags":                                        "Tags",
		"value":                                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_sagemaker_project", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
