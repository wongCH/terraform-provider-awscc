// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_amplify_branch", branchDataSourceType)
}

// branchDataSourceType returns the Terraform awscc_amplify_branch data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Amplify::Branch resource type.
func branchDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"basic_auth_config": {
			// Property: BasicAuthConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EnableBasicAuth": {
			//       "type": "boolean"
			//     },
			//     "Password": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Username": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Username",
			//     "Password"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_basic_auth": {
						// Property: EnableBasicAuth
						Type:     types.BoolType,
						Computed: true,
					},
					"password": {
						// Property: Password
						Type:     types.StringType,
						Computed: true,
					},
					"username": {
						// Property: Username
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"branch_name": {
			// Property: BranchName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"build_spec": {
			// Property: BuildSpec
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 25000,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"enable_auto_build": {
			// Property: EnableAutoBuild
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"enable_performance_mode": {
			// Property: EnablePerformanceMode
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"enable_pull_request_preview": {
			// Property: EnablePullRequestPreview
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"environment_variables": {
			// Property: EnvironmentVariables
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 255,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 5500,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"pull_request_environment_name": {
			// Property: PullRequestEnvironmentName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"stage": {
			// Property: Stage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "EXPERIMENTAL",
			//     "BETA",
			//     "PULL_REQUEST",
			//     "PRODUCTION",
			//     "DEVELOPMENT"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "insertionOrder": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
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
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Amplify::Branch",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Branch").WithTerraformTypeName("awscc_amplify_branch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                        "AppId",
		"arn":                           "Arn",
		"basic_auth_config":             "BasicAuthConfig",
		"branch_name":                   "BranchName",
		"build_spec":                    "BuildSpec",
		"description":                   "Description",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"key":                           "Key",
		"name":                          "Name",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"stage":                         "Stage",
		"tags":                          "Tags",
		"username":                      "Username",
		"value":                         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
