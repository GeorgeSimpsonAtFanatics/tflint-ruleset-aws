// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderComponentInvalidDataRule checks the pattern is valid
type AwsImagebuilderComponentInvalidDataRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderComponentInvalidDataRule returns new rule with default attributes
func NewAwsImagebuilderComponentInvalidDataRule() *AwsImagebuilderComponentInvalidDataRule {
	return &AwsImagebuilderComponentInvalidDataRule{
		resourceType:  "aws_imagebuilder_component",
		attributeName: "data",
		max:           16000,
		min:           1,
		pattern:       regexp.MustCompile(`^[^\x00]+$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderComponentInvalidDataRule) Name() string {
	return "aws_imagebuilder_component_invalid_data"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderComponentInvalidDataRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderComponentInvalidDataRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderComponentInvalidDataRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderComponentInvalidDataRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"data must be 16000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"data must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x00]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}