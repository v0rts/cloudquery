// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReservedCacheNodesOfferings() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_reserved_cache_nodes_offerings",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNodesOffering.html`,
		Resolver:    fetchElasticacheReservedCacheNodesOfferings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCacheNodesOfferingArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "duration",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Duration"),
			},
			{
				Name:     "fixed_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("FixedPrice"),
			},
			{
				Name:     "offering_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OfferingType"),
			},
			{
				Name:     "product_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductDescription"),
			},
			{
				Name:     "recurring_charges",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RecurringCharges"),
			},
			{
				Name:     "reserved_cache_nodes_offering_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedCacheNodesOfferingId"),
			},
			{
				Name:     "usage_price",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("UsagePrice"),
			},
		},
	}
}
