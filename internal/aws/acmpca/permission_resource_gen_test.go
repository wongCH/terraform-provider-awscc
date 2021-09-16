// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSACMPCAPermission_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ACMPCA::Permission", "awscc_acmpca_permission", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
