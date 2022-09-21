// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserAttachedPolicies() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_user_attached_policies",
		Resolver:  fetchIamUserAttachedPolicies,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:     "policy_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyArn"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
			},
		},
	}
}
