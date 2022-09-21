// Code generated by codegen; DO NOT EDIT.

package inspector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:      "aws_inspector_findings",
		Resolver:  fetchInspectorFindings,
		Multiplex: client.ServiceAccountRegionMultiplexer("inspector"),
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
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("Attributes"),
			},
			{
				Name:     "user_attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("UserAttributes"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "asset_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AssetAttributes"),
			},
			{
				Name:     "asset_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssetType"),
			},
			{
				Name:     "confidence",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Confidence"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "indicator_of_compromise",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IndicatorOfCompromise"),
			},
			{
				Name:     "numeric_severity",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("NumericSeverity"),
			},
			{
				Name:     "recommendation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Recommendation"),
			},
			{
				Name:     "schema_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SchemaVersion"),
			},
			{
				Name:     "service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Service"),
			},
			{
				Name:     "service_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceAttributes"),
			},
			{
				Name:     "severity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Severity"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
		},
	}
}
