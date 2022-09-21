// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_jobs",
		Resolver:  fetchIotJobs,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
				Resolver: ResolveIotJobTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "abort_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AbortConfig"),
			},
			{
				Name:     "comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Comment"),
			},
			{
				Name:     "completed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletedAt"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "document_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DocumentParameters"),
			},
			{
				Name:     "force_canceled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ForceCanceled"),
			},
			{
				Name:     "is_concurrent",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsConcurrent"),
			},
			{
				Name:     "job_executions_retry_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobExecutionsRetryConfig"),
			},
			{
				Name:     "job_executions_rollout_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobExecutionsRolloutConfig"),
			},
			{
				Name:     "job_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobId"),
			},
			{
				Name:     "job_process_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobProcessDetails"),
			},
			{
				Name:     "job_template_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobTemplateArn"),
			},
			{
				Name:     "last_updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedAt"),
			},
			{
				Name:     "namespace_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NamespaceId"),
			},
			{
				Name:     "presigned_url_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PresignedUrlConfig"),
			},
			{
				Name:     "reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReasonCode"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "target_selection",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetSelection"),
			},
			{
				Name:     "targets",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Targets"),
			},
			{
				Name:     "timeout_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TimeoutConfig"),
			},
		},
	}
}
