// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_transit_gateway_multicast_group_member", transitGatewayMulticastGroupMemberResourceType)
}

// transitGatewayMulticastGroupMemberResourceType returns the Terraform awscc_ec2_transit_gateway_multicast_group_member resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastGroupMember resource type.
func transitGatewayMulticastGroupMemberResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"group_ip_address": {
			// Property: GroupIpAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP address assigned to the transit gateway multicast group.",
			//   "type": "string"
			// }
			Description: "The IP address assigned to the transit gateway multicast group.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"group_member": {
			// Property: GroupMember
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates that the resource is a transit gateway multicast group member.",
			//   "type": "boolean"
			// }
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"group_source": {
			// Property: GroupSource
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates that the resource is a transit gateway multicast group member.",
			//   "type": "boolean"
			// }
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"member_type": {
			// Property: MemberType
			// CloudFormation resource type schema:
			// {
			//   "description": "The member type (for example, static).",
			//   "type": "string"
			// }
			Description: "The member type (for example, static).",
			Type:        types.StringType,
			Computed:    true,
		},
		"network_interface_id": {
			// Property: NetworkInterfaceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway attachment.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the resource.",
			//   "type": "string"
			// }
			Description: "The ID of the resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of resource, for example a VPC attachment.",
			//   "type": "string"
			// }
			Description: "The type of resource, for example a VPC attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"source_type": {
			// Property: SourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The source type.",
			//   "type": "string"
			// }
			Description: "The source type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the subnet.",
			//   "type": "string"
			// }
			Description: "The ID of the subnet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_multicast_domain_id": {
			// Property: TransitGatewayMulticastDomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway multicast domain.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::EC2::TransitGatewayMulticastGroupMember registers and deregisters members and sources (network interfaces) with the transit gateway multicast group",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastGroupMember").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_group_member")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_ip_address":                    "GroupIpAddress",
		"group_member":                        "GroupMember",
		"group_source":                        "GroupSource",
		"member_type":                         "MemberType",
		"network_interface_id":                "NetworkInterfaceId",
		"resource_id":                         "ResourceId",
		"resource_type":                       "ResourceType",
		"source_type":                         "SourceType",
		"subnet_id":                           "SubnetId",
		"transit_gateway_attachment_id":       "TransitGatewayAttachmentId",
		"transit_gateway_multicast_domain_id": "TransitGatewayMulticastDomainId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
