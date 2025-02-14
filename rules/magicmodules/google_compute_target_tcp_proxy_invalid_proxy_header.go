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

// GoogleComputeTargetTcpProxyInvalidProxyHeaderRule checks the pattern is valid
type GoogleComputeTargetTcpProxyInvalidProxyHeaderRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeTargetTcpProxyInvalidProxyHeaderRule returns new rule with default attributes
func NewGoogleComputeTargetTcpProxyInvalidProxyHeaderRule() *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule {
	return &GoogleComputeTargetTcpProxyInvalidProxyHeaderRule{
		resourceType:  "google_compute_target_tcp_proxy",
		attributeName: "proxy_header",
	}
}

// Name returns the rule name
func (r *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule) Name() string {
	return "google_compute_target_tcp_proxy_invalid_proxy_header"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeTargetTcpProxyInvalidProxyHeaderRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
