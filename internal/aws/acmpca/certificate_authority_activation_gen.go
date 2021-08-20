// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca

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
	registry.AddResourceTypeFactory("awscc_acmpca_certificate_authority_activation", certificateAuthorityActivationResourceType)
}

// certificateAuthorityActivationResourceType returns the Terraform awscc_acmpca_certificate_authority_activation resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ACMPCA::CertificateAuthorityActivation resource type.
func certificateAuthorityActivationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"certificate": {
			// Property: Certificate
			// CloudFormation resource type schema:
			// {
			//   "description": "Certificate Authority certificate that will be installed in the Certificate Authority.",
			//   "type": "string"
			// }
			Description: "Certificate Authority certificate that will be installed in the Certificate Authority.",
			Type:        types.StringType,
			Required:    true,
			// Certificate is a force-new attribute.
			// Certificate is a write-only attribute.
		},
		"certificate_authority_arn": {
			// Property: CertificateAuthorityArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn of the Certificate Authority.",
			//   "type": "string"
			// }
			Description: "Arn of the Certificate Authority.",
			Type:        types.StringType,
			Required:    true,
			// CertificateAuthorityArn is a force-new attribute.
		},
		"certificate_chain": {
			// Property: CertificateChain
			// CloudFormation resource type schema:
			// {
			//   "description": "Certificate chain for the Certificate Authority certificate.",
			//   "type": "string"
			// }
			Description: "Certificate chain for the Certificate Authority certificate.",
			Type:        types.StringType,
			Optional:    true,
			// CertificateChain is a write-only attribute.
		},
		"complete_certificate_chain": {
			// Property: CompleteCertificateChain
			// CloudFormation resource type schema:
			// {
			//   "description": "The complete certificate chain, including the Certificate Authority certificate.",
			//   "type": "string"
			// }
			Description: "The complete certificate chain, including the Certificate Authority certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Certificate Authority.",
			//   "type": "string"
			// }
			Description: "The status of the Certificate Authority.",
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
		Description: "Used to install the certificate authority certificate and update the certificate authority status.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::CertificateAuthorityActivation").WithTerraformTypeName("awscc_acmpca_certificate_authority_activation").WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate":                "Certificate",
		"certificate_authority_arn":  "CertificateAuthorityArn",
		"certificate_chain":          "CertificateChain",
		"complete_certificate_chain": "CompleteCertificateChain",
		"status":                     "Status",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Certificate",
		"/properties/CertificateChain",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_acmpca_certificate_authority_activation", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
