// Code generated by codegen; DO NOT EDIT.

package sqs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Queues() *schema.Table {
	return &schema.Table{
		Name:                "aws_sqs_queues",
		Resolver:            fetchSqsQueues,
		PreResourceResolver: getQueue,
		Multiplex:           client.ServiceAccountRegionMultiplexer("sqs"),
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSqsQueueTags,
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: client.MarshaledJsonResolver("Policy"),
			},
			{
				Name:     "redrive_policy",
				Type:     schema.TypeJSON,
				Resolver: client.MarshaledJsonResolver("RedrivePolicy"),
			},
			{
				Name:     "redrive_allow_policy",
				Type:     schema.TypeJSON,
				Resolver: client.MarshaledJsonResolver("RedriveAllowPolicy"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "approximate_number_of_messages",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ApproximateNumberOfMessages"),
			},
			{
				Name:     "approximate_number_of_messages_delayed",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ApproximateNumberOfMessagesDelayed"),
			},
			{
				Name:     "approximate_number_of_messages_not_visible",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ApproximateNumberOfMessagesNotVisible"),
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreatedTimestamp"),
			},
			{
				Name:     "delay_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DelaySeconds"),
			},
			{
				Name:     "last_modified_timestamp",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LastModifiedTimestamp"),
			},
			{
				Name:     "maximum_message_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumMessageSize"),
			},
			{
				Name:     "message_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MessageRetentionPeriod"),
			},
			{
				Name:     "receive_message_wait_time_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReceiveMessageWaitTimeSeconds"),
			},
			{
				Name:     "visibility_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VisibilityTimeout"),
			},
			{
				Name:     "kms_master_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsMasterKeyId"),
			},
			{
				Name:     "kms_data_key_reuse_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("KmsDataKeyReusePeriodSeconds"),
			},
			{
				Name:     "sqs_managed_sse_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SqsManagedSseEnabled"),
			},
			{
				Name:     "fifo_queue",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FifoQueue"),
			},
			{
				Name:     "content_based_deduplication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ContentBasedDeduplication"),
			},
			{
				Name:     "deduplication_scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeduplicationScope"),
			},
			{
				Name:     "fifo_throughput_limit",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FifoThroughputLimit"),
			},
			{
				Name:     "unknown_fields",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UnknownFields"),
			},
		},
	}
}
