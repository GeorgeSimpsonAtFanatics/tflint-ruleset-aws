// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule() *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule {
	return &AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "tag_key_scope",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_tag_key_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"tag_key_scope must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"tag_key_scope must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}