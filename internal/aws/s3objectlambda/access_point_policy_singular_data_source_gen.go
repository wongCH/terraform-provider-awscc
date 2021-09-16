// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3objectlambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_s3objectlambda_access_point_policy", accessPointPolicyDataSourceType)
}

// accessPointPolicyDataSourceType returns the Terraform awscc_s3objectlambda_access_point_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPointPolicy resource type.
func accessPointPolicyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"object_lambda_access_point": {
			// Property: ObjectLambdaAccessPoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
			//   "maxLength": 45,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
			//   "type": "object"
			// }
			Description: "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3ObjectLambda::AccessPointPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPointPolicy").WithTerraformTypeName("awscc_s3objectlambda_access_point_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"object_lambda_access_point": "ObjectLambdaAccessPoint",
		"policy_document":            "PolicyDocument",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
