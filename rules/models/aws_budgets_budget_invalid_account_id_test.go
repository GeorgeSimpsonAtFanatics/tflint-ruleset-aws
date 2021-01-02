// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsBudgetsBudgetInvalidAccountIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_budgets_budget" "foo" {
	account_id = "abcdefghijkl"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsBudgetsBudgetInvalidAccountIDRule(),
					Message: `"abcdefghijkl" does not match valid pattern ^\d{12}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_budgets_budget" "foo" {
	account_id = "123456789012"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsBudgetsBudgetInvalidAccountIDRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}