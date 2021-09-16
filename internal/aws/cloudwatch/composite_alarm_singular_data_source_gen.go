// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudwatch_composite_alarm", compositeAlarmDataSourceType)
}

// compositeAlarmDataSourceType returns the Terraform awscc_cloudwatch_composite_alarm data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudWatch::CompositeAlarm resource type.
func compositeAlarmDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"actions_enabled": {
			// Property: ActionsEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"alarm_actions": {
			// Property: AlarmActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"alarm_description": {
			// Property: AlarmDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the alarm",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the alarm",
			Type:        types.StringType,
			Computed:    true,
		},
		"alarm_name": {
			// Property: AlarmName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Composite Alarm",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Composite Alarm",
			Type:        types.StringType,
			Computed:    true,
		},
		"alarm_rule": {
			// Property: AlarmRule
			// CloudFormation resource type schema:
			// {
			//   "description": "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			//   "maxLength": 10240,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the alarm",
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the alarm",
			Type:        types.StringType,
			Computed:    true,
		},
		"insufficient_data_actions": {
			// Property: InsufficientDataActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"ok_actions": {
			// Property: OKActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudWatch::CompositeAlarm",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::CompositeAlarm").WithTerraformTypeName("awscc_cloudwatch_composite_alarm")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions_enabled":           "ActionsEnabled",
		"alarm_actions":             "AlarmActions",
		"alarm_description":         "AlarmDescription",
		"alarm_name":                "AlarmName",
		"alarm_rule":                "AlarmRule",
		"arn":                       "Arn",
		"insufficient_data_actions": "InsufficientDataActions",
		"ok_actions":                "OKActions",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
