// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_asset", assetDataSource)
}

// assetDataSource returns the Terraform awscc_iotsitewise_asset data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::Asset resource.
func assetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the asset",
		//	  "type": "string"
		//	}
		"asset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the asset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the asset",
		//	  "type": "string"
		//	}
		"asset_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the asset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetHierarchies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A hierarchy specifies allowed parent/child asset relationships.",
		//	    "properties": {
		//	      "ChildAssetId": {
		//	        "description": "The ID of the child asset to be associated.",
		//	        "type": "string"
		//	      },
		//	      "LogicalId": {
		//	        "description": "The LogicalID of a hierarchy in the parent asset's model.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LogicalId",
		//	      "ChildAssetId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"asset_hierarchies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ChildAssetId
					"child_asset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the child asset to be associated.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LogicalId
					"logical_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The LogicalID of a hierarchy in the parent asset's model.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AssetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the asset",
		//	  "type": "string"
		//	}
		"asset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the asset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetModelId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the asset model from which to create the asset.",
		//	  "type": "string"
		//	}
		"asset_model_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the asset model from which to create the asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique, friendly name for the asset.",
		//	  "type": "string"
		//	}
		"asset_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique, friendly name for the asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssetProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The asset property's definition, alias, unit, and notification state.",
		//	    "properties": {
		//	      "Alias": {
		//	        "description": "The property alias that identifies the property.",
		//	        "type": "string"
		//	      },
		//	      "LogicalId": {
		//	        "description": "Customer provided ID for property.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "NotificationState": {
		//	        "description": "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
		//	        "enum": [
		//	          "ENABLED",
		//	          "DISABLED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Unit": {
		//	        "description": "The unit of measure (such as Newtons or RPM) of the asset property. If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LogicalId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"asset_properties": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Alias
					"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The property alias that identifies the property.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LogicalId
					"logical_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Customer provided ID for property.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: NotificationState
					"notification_state": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Unit
					"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The unit of measure (such as Newtons or RPM) of the asset property. If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the asset.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the asset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Asset",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Asset").WithTerraformTypeName("awscc_iotsitewise_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":              "Alias",
		"asset_arn":          "AssetArn",
		"asset_description":  "AssetDescription",
		"asset_hierarchies":  "AssetHierarchies",
		"asset_id":           "AssetId",
		"asset_model_id":     "AssetModelId",
		"asset_name":         "AssetName",
		"asset_properties":   "AssetProperties",
		"child_asset_id":     "ChildAssetId",
		"key":                "Key",
		"logical_id":         "LogicalId",
		"notification_state": "NotificationState",
		"tags":               "Tags",
		"unit":               "Unit",
		"value":              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
