// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmNetworkProfileInvalidTypeRule checks the pattern is valid
type AwsDevicefarmNetworkProfileInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDevicefarmNetworkProfileInvalidTypeRule returns new rule with default attributes
func NewAwsDevicefarmNetworkProfileInvalidTypeRule() *AwsDevicefarmNetworkProfileInvalidTypeRule {
	return &AwsDevicefarmNetworkProfileInvalidTypeRule{
		resourceType:  "aws_devicefarm_network_profile",
		attributeName: "type",
		enum: []string{
			"CURATED",
			"PRIVATE",
		},
	}
}

// Name returns the rule name
func (r *AwsDevicefarmNetworkProfileInvalidTypeRule) Name() string {
	return "aws_devicefarm_network_profile_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmNetworkProfileInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmNetworkProfileInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmNetworkProfileInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmNetworkProfileInvalidTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}