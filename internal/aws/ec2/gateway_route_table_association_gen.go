// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("awscc_ec2_gateway_route_table_association", gatewayRouteTableAssociationResourceType)
}

// gatewayRouteTableAssociationResourceType returns the Terraform awscc_ec2_gateway_route_table_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::GatewayRouteTableAssociation resource type.
func gatewayRouteTableAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "description": "The route table association ID.",
			//   "type": "string"
			// }
			Description: "The route table association ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"gateway_id": {
			// Property: GatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the gateway.",
			Type:        types.StringType,
			Required:    true,
			// GatewayId is a force-new attribute.
		},
		"route_table_id": {
			// Property: RouteTableId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the route table.",
			//   "type": "string"
			// }
			Description: "The ID of the route table.",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::GatewayRouteTableAssociation").WithTerraformTypeName("awscc_ec2_gateway_route_table_association").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"gateway_id":     "GatewayId",
		"route_table_id": "RouteTableId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_gateway_route_table_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
