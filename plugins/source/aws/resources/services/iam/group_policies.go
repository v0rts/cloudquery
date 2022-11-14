// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupPolicies() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_group_policies",
		Resolver:            fetchIamGroupPolicies,
		PreResourceResolver: getGroupPolicy,
		Multiplex:           client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamGroupPolicyPolicyDocument,
			},
			{
				Name:     "group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupName"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
