// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Documents() *schema.Table {
	return &schema.Table{
		Name:      "aws_ssm_documents",
		Resolver:  fetchSsmDocuments,
		Multiplex: client.ServiceAccountRegionMultiplexer("ssm"),
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
				Resolver: resolveDocumentARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: resolveDocumentPermission,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "approved_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApprovedVersion"),
			},
			{
				Name:     "attachments_information",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttachmentsInformation"),
			},
			{
				Name:     "author",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Author"),
			},
			{
				Name:     "category",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Category"),
			},
			{
				Name:     "category_enum",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CategoryEnum"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "default_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultVersion"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "document_format",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentFormat"),
			},
			{
				Name:     "document_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentType"),
			},
			{
				Name:     "document_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentVersion"),
			},
			{
				Name:     "hash",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Hash"),
			},
			{
				Name:     "hash_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HashType"),
			},
			{
				Name:     "latest_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "pending_review_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingReviewVersion"),
			},
			{
				Name:     "platform_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PlatformTypes"),
			},
			{
				Name:     "requires",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Requires"),
			},
			{
				Name:     "review_information",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReviewInformation"),
			},
			{
				Name:     "review_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReviewStatus"),
			},
			{
				Name:     "schema_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SchemaVersion"),
			},
			{
				Name:     "sha_1",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sha1"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_information",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusInformation"),
			},
			{
				Name:     "target_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetType"),
			},
			{
				Name:     "version_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VersionName"),
			},
		},
	}
}
