// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("awscc_networkmanager_device", deviceResourceType)
}

// deviceResourceType returns the Terraform awscc_networkmanager_device resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::Device resource type.
func deviceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the device.",
			//   "type": "string"
			// }
			Description: "The description of the device.",
			Type:        types.StringType,
			Optional:    true,
		},
		"device_arn": {
			// Property: DeviceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the device.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the device.",
			//   "type": "string"
			// }
			Description: "The ID of the device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the global network.",
			//   "type": "string"
			// }
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			// GlobalNetworkId is a force-new attribute.
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The site location.",
			//   "properties": {
			//     "Address": {
			//       "description": "The physical address.",
			//       "type": "string"
			//     },
			//     "Latitude": {
			//       "description": "The latitude.",
			//       "type": "string"
			//     },
			//     "Longitude": {
			//       "description": "The longitude.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The site location.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Description: "The physical address.",
						Type:        types.StringType,
						Optional:    true,
					},
					"latitude": {
						// Property: Latitude
						Description: "The latitude.",
						Type:        types.StringType,
						Optional:    true,
					},
					"longitude": {
						// Property: Longitude
						Description: "The longitude.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"model": {
			// Property: Model
			// CloudFormation resource type schema:
			// {
			//   "description": "The device model",
			//   "type": "string"
			// }
			Description: "The device model",
			Type:        types.StringType,
			Optional:    true,
		},
		"serial_number": {
			// Property: SerialNumber
			// CloudFormation resource type schema:
			// {
			//   "description": "The device serial number.",
			//   "type": "string"
			// }
			Description: "The device serial number.",
			Type:        types.StringType,
			Optional:    true,
		},
		"site_id": {
			// Property: SiteId
			// CloudFormation resource type schema:
			// {
			//   "description": "The site ID.",
			//   "type": "string"
			// }
			Description: "The site ID.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the device.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a device resource.",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the device.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The device type.",
			//   "type": "string"
			// }
			Description: "The device type.",
			Type:        types.StringType,
			Optional:    true,
		},
		"vendor": {
			// Property: Vendor
			// CloudFormation resource type schema:
			// {
			//   "description": "The device vendor.",
			//   "type": "string"
			// }
			Description: "The device vendor.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::NetworkManager::Device type describes a device.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::Device").WithTerraformTypeName("awscc_networkmanager_device").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":           "Address",
		"description":       "Description",
		"device_arn":        "DeviceArn",
		"device_id":         "DeviceId",
		"global_network_id": "GlobalNetworkId",
		"key":               "Key",
		"latitude":          "Latitude",
		"location":          "Location",
		"longitude":         "Longitude",
		"model":             "Model",
		"serial_number":     "SerialNumber",
		"site_id":           "SiteId",
		"tags":              "Tags",
		"type":              "Type",
		"value":             "Value",
		"vendor":            "Vendor",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkmanager_device", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
