// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeRegionBackendServiceInvalidSessionAffinityRule checks the pattern is valid
type GoogleComputeRegionBackendServiceInvalidSessionAffinityRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule returns new rule with default attributes
func NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule() *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule {
	return &GoogleComputeRegionBackendServiceInvalidSessionAffinityRule{
		resourceType:  "google_compute_region_backend_service",
		attributeName: "session_affinity",
	}
}

// Name returns the rule name
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Name() string {
	return "google_compute_region_backend_service_invalid_session_affinity"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
