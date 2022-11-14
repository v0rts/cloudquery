// Code generated by codegen; DO NOT EDIT.

package dynamodb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tables() *schema.Table {
	return &schema.Table{
		Name:                "aws_dynamodb_tables",
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html`,
		Resolver:            fetchDynamodbTables,
		PreResourceResolver: getTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer("dynamodb"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDynamodbTableTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "archival_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ArchivalSummary"),
			},
			{
				Name:     "attribute_definitions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttributeDefinitions"),
			},
			{
				Name:     "billing_mode_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BillingModeSummary"),
			},
			{
				Name:     "creation_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDateTime"),
			},
			{
				Name:     "global_secondary_indexes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalSecondaryIndexes"),
			},
			{
				Name:     "global_table_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalTableVersion"),
			},
			{
				Name:     "item_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ItemCount"),
			},
			{
				Name:     "key_schema",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KeySchema"),
			},
			{
				Name:     "latest_stream_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestStreamArn"),
			},
			{
				Name:     "latest_stream_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestStreamLabel"),
			},
			{
				Name:     "local_secondary_indexes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocalSecondaryIndexes"),
			},
			{
				Name:     "provisioned_throughput",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProvisionedThroughput"),
			},
			{
				Name:     "replicas",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Replicas"),
			},
			{
				Name:     "restore_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RestoreSummary"),
			},
			{
				Name:     "sse_description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SSEDescription"),
			},
			{
				Name:     "stream_specification",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StreamSpecification"),
			},
			{
				Name:     "table_class_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TableClassSummary"),
			},
			{
				Name:     "table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableId"),
			},
			{
				Name:     "table_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableName"),
			},
			{
				Name:     "table_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TableSizeBytes"),
			},
			{
				Name:     "table_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableStatus"),
			},
		},

		Relations: []*schema.Table{
			TableReplicaAutoScalings(),
			TableContinuousBackups(),
		},
	}
}
