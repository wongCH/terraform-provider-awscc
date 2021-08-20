// Code generated by generators/resource/main.go; DO NOT EDIT.

package groundstation

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
	registry.AddResourceTypeFactory("awscc_groundstation_config", configResourceType)
}

// configResourceType returns the Terraform awscc_groundstation_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GroundStation::Config resource type.
func configResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"config_data": {
			// Property: ConfigData
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AntennaDownlinkConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "SpectrumConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Bandwidth": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Units": {
			//                   "enum": [
			//                     "GHz",
			//                     "MHz",
			//                     "kHz"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "type": "number"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "CenterFrequency": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Units": {
			//                   "enum": [
			//                     "GHz",
			//                     "MHz",
			//                     "kHz"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "type": "number"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "Polarization": {
			//               "enum": [
			//                 "LEFT_HAND",
			//                 "RIGHT_HAND",
			//                 "NONE"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "AntennaDownlinkDemodDecodeConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DecodeConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "UnvalidatedJSON": {
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "DemodulationConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "UnvalidatedJSON": {
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SpectrumConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Bandwidth": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Units": {
			//                   "enum": [
			//                     "GHz",
			//                     "MHz",
			//                     "kHz"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "type": "number"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "CenterFrequency": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Units": {
			//                   "enum": [
			//                     "GHz",
			//                     "MHz",
			//                     "kHz"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "type": "number"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "Polarization": {
			//               "enum": [
			//                 "LEFT_HAND",
			//                 "RIGHT_HAND",
			//                 "NONE"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "AntennaUplinkConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "SpectrumConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "CenterFrequency": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Units": {
			//                   "enum": [
			//                     "GHz",
			//                     "MHz",
			//                     "kHz"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "type": "number"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "Polarization": {
			//               "enum": [
			//                 "LEFT_HAND",
			//                 "RIGHT_HAND",
			//                 "NONE"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "TargetEirp": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Units": {
			//               "enum": [
			//                 "dBW"
			//               ],
			//               "type": "string"
			//             },
			//             "Value": {
			//               "type": "number"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "TransmitDisabled": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DataflowEndpointConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DataflowEndpointName": {
			//           "type": "string"
			//         },
			//         "DataflowEndpointRegion": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "S3RecordingConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "BucketArn": {
			//           "type": "string"
			//         },
			//         "Prefix": {
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "RoleArn": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "TrackingConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Autotrack": {
			//           "enum": [
			//             "REQUIRED",
			//             "PREFERRED",
			//             "REMOVED"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UplinkEchoConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AntennaUplinkConfigArn": {
			//           "type": "string"
			//         },
			//         "Enabled": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"antenna_downlink_config": {
						// Property: AntennaDownlinkConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"spectrum_config": {
									// Property: SpectrumConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bandwidth": {
												// Property: Bandwidth
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"units": {
															// Property: Units
															Type:     types.StringType,
															Optional: true,
														},
														"value": {
															// Property: Value
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"center_frequency": {
												// Property: CenterFrequency
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"units": {
															// Property: Units
															Type:     types.StringType,
															Optional: true,
														},
														"value": {
															// Property: Value
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"polarization": {
												// Property: Polarization
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"antenna_downlink_demod_decode_config": {
						// Property: AntennaDownlinkDemodDecodeConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"decode_config": {
									// Property: DecodeConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"unvalidated_json": {
												// Property: UnvalidatedJSON
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"demodulation_config": {
									// Property: DemodulationConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"unvalidated_json": {
												// Property: UnvalidatedJSON
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"spectrum_config": {
									// Property: SpectrumConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bandwidth": {
												// Property: Bandwidth
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"units": {
															// Property: Units
															Type:     types.StringType,
															Optional: true,
														},
														"value": {
															// Property: Value
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"center_frequency": {
												// Property: CenterFrequency
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"units": {
															// Property: Units
															Type:     types.StringType,
															Optional: true,
														},
														"value": {
															// Property: Value
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"polarization": {
												// Property: Polarization
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"antenna_uplink_config": {
						// Property: AntennaUplinkConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"spectrum_config": {
									// Property: SpectrumConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"center_frequency": {
												// Property: CenterFrequency
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"units": {
															// Property: Units
															Type:     types.StringType,
															Optional: true,
														},
														"value": {
															// Property: Value
															Type:     types.NumberType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"polarization": {
												// Property: Polarization
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"target_eirp": {
									// Property: TargetEirp
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"units": {
												// Property: Units
												Type:     types.StringType,
												Optional: true,
											},
											"value": {
												// Property: Value
												Type:     types.NumberType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"transmit_disabled": {
									// Property: TransmitDisabled
									Type:     types.BoolType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"dataflow_endpoint_config": {
						// Property: DataflowEndpointConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dataflow_endpoint_name": {
									// Property: DataflowEndpointName
									Type:     types.StringType,
									Optional: true,
								},
								"dataflow_endpoint_region": {
									// Property: DataflowEndpointRegion
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"s3_recording_config": {
						// Property: S3RecordingConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket_arn": {
									// Property: BucketArn
									Type:     types.StringType,
									Optional: true,
								},
								"prefix": {
									// Property: Prefix
									Type:     types.StringType,
									Optional: true,
								},
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"tracking_config": {
						// Property: TrackingConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"autotrack": {
									// Property: Autotrack
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"uplink_echo_config": {
						// Property: UplinkEchoConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"antenna_uplink_config_arn": {
									// Property: AntennaUplinkConfigArn
									Type:     types.StringType,
									Optional: true,
								},
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
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
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS Ground Station config resource type for CloudFormation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::Config").WithTerraformTypeName("awscc_groundstation_config").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"antenna_downlink_config":              "AntennaDownlinkConfig",
		"antenna_downlink_demod_decode_config": "AntennaDownlinkDemodDecodeConfig",
		"antenna_uplink_config":                "AntennaUplinkConfig",
		"antenna_uplink_config_arn":            "AntennaUplinkConfigArn",
		"arn":                                  "Arn",
		"autotrack":                            "Autotrack",
		"bandwidth":                            "Bandwidth",
		"bucket_arn":                           "BucketArn",
		"center_frequency":                     "CenterFrequency",
		"config_data":                          "ConfigData",
		"dataflow_endpoint_config":             "DataflowEndpointConfig",
		"dataflow_endpoint_name":               "DataflowEndpointName",
		"dataflow_endpoint_region":             "DataflowEndpointRegion",
		"decode_config":                        "DecodeConfig",
		"demodulation_config":                  "DemodulationConfig",
		"enabled":                              "Enabled",
		"id":                                   "Id",
		"key":                                  "Key",
		"name":                                 "Name",
		"polarization":                         "Polarization",
		"prefix":                               "Prefix",
		"role_arn":                             "RoleArn",
		"s3_recording_config":                  "S3RecordingConfig",
		"spectrum_config":                      "SpectrumConfig",
		"tags":                                 "Tags",
		"target_eirp":                          "TargetEirp",
		"tracking_config":                      "TrackingConfig",
		"transmit_disabled":                    "TransmitDisabled",
		"type":                                 "Type",
		"units":                                "Units",
		"unvalidated_json":                     "UnvalidatedJSON",
		"uplink_echo_config":                   "UplinkEchoConfig",
		"value":                                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_groundstation_config", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
