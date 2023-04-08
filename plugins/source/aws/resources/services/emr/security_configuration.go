package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SecurityConfigurations() *schema.Table {
	tableName := "aws_emr_security_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html`,
		Resolver:    fetchSecurityConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.SecurityConfigurationSummary{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeJSON,
				Resolver: resolveConfiguration,
			},
		},
	}
}

func fetchSecurityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := emr.ListSecurityConfigurationsInput{}
	c := meta.(*client.Client)
	svc := c.Services().Emr
	for {
		response, err := svc.ListSecurityConfigurations(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.SecurityConfigurations

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)
	item := resource.Item.(types.SecurityConfigurationSummary)
	svc := c.Services().Emr
	response, err := svc.DescribeSecurityConfiguration(ctx, &emr.DescribeSecurityConfigurationInput{Name: item.Name})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, response.SecurityConfiguration)
}
