// Code generated by codegen; DO NOT EDIT.

package cloudfront

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CachePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudfront_cache_policies",
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html`,
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCachePolicyARN(),
			},
			{
				Name:     "cache_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CachePolicy"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
