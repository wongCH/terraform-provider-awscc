// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_cloudfront_response_headers_policy", responseHeadersPolicyResource)
}

// responseHeadersPolicyResource returns the Terraform awscc_cloudfront_response_headers_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFront::ResponseHeadersPolicy resource.
func responseHeadersPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResponseHeadersPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Comment": {
		//	      "type": "string"
		//	    },
		//	    "CorsConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AccessControlAllowCredentials": {
		//	          "type": "boolean"
		//	        },
		//	        "AccessControlAllowHeaders": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Items": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Items"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "AccessControlAllowMethods": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Items": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Items"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "AccessControlAllowOrigins": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Items": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Items"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "AccessControlExposeHeaders": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Items": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Items"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "AccessControlMaxAgeSec": {
		//	          "type": "integer"
		//	        },
		//	        "OriginOverride": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "AccessControlAllowOrigins",
		//	        "AccessControlAllowHeaders",
		//	        "AccessControlAllowMethods",
		//	        "AccessControlAllowCredentials",
		//	        "OriginOverride"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "CustomHeadersConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Items": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Header": {
		//	                "type": "string"
		//	              },
		//	              "Override": {
		//	                "type": "boolean"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Header",
		//	              "Value",
		//	              "Override"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "Items"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Name": {
		//	      "type": "string"
		//	    },
		//	    "RemoveHeadersConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Items": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Header": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Header"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "Items"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "SecurityHeadersConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ContentSecurityPolicy": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ContentSecurityPolicy": {
		//	              "type": "string"
		//	            },
		//	            "Override": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override",
		//	            "ContentSecurityPolicy"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ContentTypeOptions": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Override": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "FrameOptions": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "FrameOption": {
		//	              "pattern": "^(DENY|SAMEORIGIN)$",
		//	              "type": "string"
		//	            },
		//	            "Override": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override",
		//	            "FrameOption"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ReferrerPolicy": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Override": {
		//	              "type": "boolean"
		//	            },
		//	            "ReferrerPolicy": {
		//	              "pattern": "^(no-referrer|no-referrer-when-downgrade|origin|origin-when-cross-origin|same-origin|strict-origin|strict-origin-when-cross-origin|unsafe-url)$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override",
		//	            "ReferrerPolicy"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "StrictTransportSecurity": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AccessControlMaxAgeSec": {
		//	              "type": "integer"
		//	            },
		//	            "IncludeSubdomains": {
		//	              "type": "boolean"
		//	            },
		//	            "Override": {
		//	              "type": "boolean"
		//	            },
		//	            "Preload": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override",
		//	            "AccessControlMaxAgeSec"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "XSSProtection": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ModeBlock": {
		//	              "type": "boolean"
		//	            },
		//	            "Override": {
		//	              "type": "boolean"
		//	            },
		//	            "Protection": {
		//	              "type": "boolean"
		//	            },
		//	            "ReportUri": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Override",
		//	            "Protection"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ServerTimingHeadersConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        },
		//	        "SamplingRate": {
		//	          "maximum": 100,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enabled"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name"
		//	  ],
		//	  "type": "object"
		//	}
		"response_headers_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CorsConfig
				"cors_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AccessControlAllowCredentials
						"access_control_allow_credentials": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: AccessControlAllowHeaders
						"access_control_allow_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Items
								"items": schema.ListAttribute{ /*START ATTRIBUTE*/
									CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
									Required:   true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: AccessControlAllowMethods
						"access_control_allow_methods": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Items
								"items": schema.ListAttribute{ /*START ATTRIBUTE*/
									CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
									Required:   true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: AccessControlAllowOrigins
						"access_control_allow_origins": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Items
								"items": schema.ListAttribute{ /*START ATTRIBUTE*/
									CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
									Required:   true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: AccessControlExposeHeaders
						"access_control_expose_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Items
								"items": schema.ListAttribute{ /*START ATTRIBUTE*/
									CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
									Required:   true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: AccessControlMaxAgeSec
						"access_control_max_age_sec": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: OriginOverride
						"origin_override": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CustomHeadersConfig
				"custom_headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Items
						"items": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Header
									"header": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
									// Property: Override
									"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
									// Property: Value
									"value": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							CustomType: cctypes.NewMultisetTypeOf[types.Object](ctx),
							Required:   true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: RemoveHeadersConfig
				"remove_headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Items
						"items": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Header
									"header": schema.StringAttribute{ /*START ATTRIBUTE*/
										Required: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SecurityHeadersConfig
				"security_headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ContentSecurityPolicy
						"content_security_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ContentSecurityPolicy
								"content_security_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ContentTypeOptions
						"content_type_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: FrameOptions
						"frame_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: FrameOption
								"frame_option": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.RegexMatches(regexp.MustCompile("^(DENY|SAMEORIGIN)$"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ReferrerPolicy
						"referrer_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: ReferrerPolicy
								"referrer_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.RegexMatches(regexp.MustCompile("^(no-referrer|no-referrer-when-downgrade|origin|origin-when-cross-origin|same-origin|strict-origin|strict-origin-when-cross-origin|unsafe-url)$"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: StrictTransportSecurity
						"strict_transport_security": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccessControlMaxAgeSec
								"access_control_max_age_sec": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: IncludeSubdomains
								"include_subdomains": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: Preload
								"preload": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: XSSProtection
						"xss_protection": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ModeBlock
								"mode_block": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Override
								"override": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: Protection
								"protection": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: ReportUri
								"report_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServerTimingHeadersConfig
				"server_timing_headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: SamplingRate
						"sampling_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 100.000000),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::CloudFront::ResponseHeadersPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::ResponseHeadersPolicy").WithTerraformTypeName("awscc_cloudfront_response_headers_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_control_allow_credentials": "AccessControlAllowCredentials",
		"access_control_allow_headers":     "AccessControlAllowHeaders",
		"access_control_allow_methods":     "AccessControlAllowMethods",
		"access_control_allow_origins":     "AccessControlAllowOrigins",
		"access_control_expose_headers":    "AccessControlExposeHeaders",
		"access_control_max_age_sec":       "AccessControlMaxAgeSec",
		"comment":                          "Comment",
		"content_security_policy":          "ContentSecurityPolicy",
		"content_type_options":             "ContentTypeOptions",
		"cors_config":                      "CorsConfig",
		"custom_headers_config":            "CustomHeadersConfig",
		"enabled":                          "Enabled",
		"frame_option":                     "FrameOption",
		"frame_options":                    "FrameOptions",
		"header":                           "Header",
		"id":                               "Id",
		"include_subdomains":               "IncludeSubdomains",
		"items":                            "Items",
		"last_modified_time":               "LastModifiedTime",
		"mode_block":                       "ModeBlock",
		"name":                             "Name",
		"origin_override":                  "OriginOverride",
		"override":                         "Override",
		"preload":                          "Preload",
		"protection":                       "Protection",
		"referrer_policy":                  "ReferrerPolicy",
		"remove_headers_config":            "RemoveHeadersConfig",
		"report_uri":                       "ReportUri",
		"response_headers_policy_config":   "ResponseHeadersPolicyConfig",
		"sampling_rate":                    "SamplingRate",
		"security_headers_config":          "SecurityHeadersConfig",
		"server_timing_headers_config":     "ServerTimingHeadersConfig",
		"strict_transport_security":        "StrictTransportSecurity",
		"value":                            "Value",
		"xss_protection":                   "XSSProtection",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
