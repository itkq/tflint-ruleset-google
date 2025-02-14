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

// GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule checks the pattern is valid
type GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule returns new rule with default attributes
func NewGoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule() *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule {
	return &GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule{
		resourceType:  "google_network_services_edge_cache_origin",
		attributeName: "protocol",
	}
}

// Name returns the rule name
func (r *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule) Name() string {
	return "google_network_services_edge_cache_origin_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"HTTP2", "HTTPS", "HTTP", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
