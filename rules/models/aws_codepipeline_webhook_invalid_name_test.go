// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCodepipelineWebhookInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	name = "webhook-github-bar/testing"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCodepipelineWebhookInvalidNameRule(),
					Message: `"webhook-github-bar/testing" does not match valid pattern ^[A-Za-z0-9.@\-_]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	name = "test-webhook-github-bar"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCodepipelineWebhookInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}