// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacie2AccountInvalidStatusRule checks the pattern is valid
type AwsMacie2AccountInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMacie2AccountInvalidStatusRule returns new rule with default attributes
func NewAwsMacie2AccountInvalidStatusRule() *AwsMacie2AccountInvalidStatusRule {
	return &AwsMacie2AccountInvalidStatusRule{
		resourceType:  "aws_macie2_account",
		attributeName: "status",
		enum: []string{
			"PAUSED",
			"ENABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsMacie2AccountInvalidStatusRule) Name() string {
	return "aws_macie2_account_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacie2AccountInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacie2AccountInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacie2AccountInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacie2AccountInvalidStatusRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}