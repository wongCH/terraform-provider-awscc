// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol

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
	registry.AddResourceTypeFactory("awscc_route53recoverycontrol_routing_control", routingControlResourceType)
}

// routingControlResourceType returns the Terraform awscc_route53recoverycontrol_routing_control resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryControl::RoutingControl resource type.
func routingControlResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster_arn": {
			// Property: ClusterArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn associated with Control Panel",
			//   "type": "string"
			// }
			Description: "Arn associated with Control Panel",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ClusterArn is a force-new attribute.
			// ClusterArn is a write-only attribute.
		},
		"control_panel_arn": {
			// Property: ControlPanelArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the control panel.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the control panel.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ControlPanelArn is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the routing control. You can use any non-white space character in the name.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the routing control. You can use any non-white space character in the name.",
			Type:        types.StringType,
			Required:    true,
		},
		"routing_control_arn": {
			// Property: RoutingControlArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the routing control.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the routing control.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			//   "enum": [
			//     "PENDING",
			//     "DEPLOYED",
			//     "PENDING_DELETION"
			//   ],
			//   "type": "string"
			// }
			Description: "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS Route53 Recovery Control Routing Control resource schema .",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::RoutingControl").WithTerraformTypeName("awscc_route53recoverycontrol_routing_control").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster_arn":         "ClusterArn",
		"control_panel_arn":   "ControlPanelArn",
		"name":                "Name",
		"routing_control_arn": "RoutingControlArn",
		"status":              "Status",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ClusterArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoverycontrol_routing_control", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
