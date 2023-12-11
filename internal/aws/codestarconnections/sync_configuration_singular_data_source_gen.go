// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codestarconnections

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codestarconnections_sync_configuration", syncConfigurationDataSource)
}

// syncConfigurationDataSource returns the Terraform awscc_codestarconnections_sync_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeStarConnections::SyncConfiguration resource.
func syncConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Branch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the branch of the repository from which resources are to be synchronized,",
		//	  "type": "string"
		//	}
		"branch": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the branch of the repository from which resources are to be synchronized,",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigFile
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source provider repository path of the sync configuration file of the respective SyncType.",
		//	  "type": "string"
		//	}
		"config_file": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source provider repository path of the sync configuration file of the respective SyncType.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "the ID of the entity that owns the repository.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "the ID of the entity that owns the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the external provider where your third-party code repository is configured.",
		//	  "pattern": "^(GitHub|Bitbucket|GitHubEnterprise|GitLab)$",
		//	  "type": "string"
		//	}
		"provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the external provider where your third-party code repository is configured.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryLinkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
		//	  "pattern": "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}",
		//	  "type": "string"
		//	}
		"repository_link_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the repository that is being synced to.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the repository that is being synced to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the resource that is being synchronized to the repository.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"resource_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the resource that is being synchronized to the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SyncType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
		//	  "type": "string"
		//	}
		"sync_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeStarConnections::SyncConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarConnections::SyncConfiguration").WithTerraformTypeName("awscc_codestarconnections_sync_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"branch":             "Branch",
		"config_file":        "ConfigFile",
		"owner_id":           "OwnerId",
		"provider_type":      "ProviderType",
		"repository_link_id": "RepositoryLinkId",
		"repository_name":    "RepositoryName",
		"resource_name":      "ResourceName",
		"role_arn":           "RoleArn",
		"sync_type":          "SyncType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
