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

// GooglePrivatecaCaPoolInvalidTierRule checks the pattern is valid
type GooglePrivatecaCaPoolInvalidTierRule struct {
	resourceType  string
	attributeName string
}

// NewGooglePrivatecaCaPoolInvalidTierRule returns new rule with default attributes
func NewGooglePrivatecaCaPoolInvalidTierRule() *GooglePrivatecaCaPoolInvalidTierRule {
	return &GooglePrivatecaCaPoolInvalidTierRule{
		resourceType:  "google_privateca_ca_pool",
		attributeName: "tier",
	}
}

// Name returns the rule name
func (r *GooglePrivatecaCaPoolInvalidTierRule) Name() string {
	return "google_privateca_ca_pool_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *GooglePrivatecaCaPoolInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GooglePrivatecaCaPoolInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GooglePrivatecaCaPoolInvalidTierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GooglePrivatecaCaPoolInvalidTierRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"ENTERPRISE", "DEVOPS"}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
