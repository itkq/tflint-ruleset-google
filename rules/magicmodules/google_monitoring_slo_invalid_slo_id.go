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
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleMonitoringSloInvalidSloIdRule checks the pattern is valid
type GoogleMonitoringSloInvalidSloIdRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleMonitoringSloInvalidSloIdRule returns new rule with default attributes
func NewGoogleMonitoringSloInvalidSloIdRule() *GoogleMonitoringSloInvalidSloIdRule {
	return &GoogleMonitoringSloInvalidSloIdRule{
		resourceType:  "google_monitoring_slo",
		attributeName: "slo_id",
	}
}

// Name returns the rule name
func (r *GoogleMonitoringSloInvalidSloIdRule) Name() string {
	return "google_monitoring_slo_invalid_slo_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleMonitoringSloInvalidSloIdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleMonitoringSloInvalidSloIdRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleMonitoringSloInvalidSloIdRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleMonitoringSloInvalidSloIdRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^[a-z0-9\-]+$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
