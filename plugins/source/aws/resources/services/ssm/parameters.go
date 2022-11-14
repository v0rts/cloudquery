// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Parameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_parameters",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html`,
		Resolver:    fetchSsmParameters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: `The parameter name`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "allowed_pattern",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowedPattern"),
			},
			{
				Name:     "data_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyId"),
			},
			{
				Name:     "last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedDate"),
			},
			{
				Name:     "last_modified_user",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastModifiedUser"),
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}
