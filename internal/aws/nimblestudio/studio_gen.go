// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

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
	registry.AddResourceTypeFactory("awscc_nimblestudio_studio", studioResourceType)
}

// studioResourceType returns the Terraform awscc_nimblestudio_studio resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NimbleStudio::Studio resource type.
func studioResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"admin_role_arn": {
			// Property: AdminRoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sso_client_id": {
			// Property: SsoClientId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"studio_encryption_configuration": {
			// Property: StudioEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KeyArn": {
			//       "type": "string"
			//     },
			//     "KeyType": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "KeyType"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"key_arn": {
						// Property: KeyArn
						Type:     types.StringType,
						Optional: true,
					},
					"key_type": {
						// Property: KeyType
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"studio_name": {
			// Property: StudioName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// StudioName is a force-new attribute.
		},
		"studio_url": {
			// Property: StudioUrl
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"user_role_arn": {
			// Property: UserRoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::NimbleStudio::Studio.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::Studio").WithTerraformTypeName("awscc_nimblestudio_studio").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_role_arn":                  "AdminRoleArn",
		"display_name":                    "DisplayName",
		"home_region":                     "HomeRegion",
		"key_arn":                         "KeyArn",
		"key_type":                        "KeyType",
		"sso_client_id":                   "SsoClientId",
		"studio_encryption_configuration": "StudioEncryptionConfiguration",
		"studio_id":                       "StudioId",
		"studio_name":                     "StudioName",
		"studio_url":                      "StudioUrl",
		"tags":                            "Tags",
		"user_role_arn":                   "UserRoleArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_nimblestudio_studio", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
