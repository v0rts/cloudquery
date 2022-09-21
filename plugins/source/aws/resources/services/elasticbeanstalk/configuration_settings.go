// Code generated by codegen; DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigurationSettings() *schema.Table {
	return &schema.Table{
		Name:      "aws_elasticbeanstalk_configuration_settings",
		Resolver:  fetchElasticbeanstalkConfigurationSettings,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
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
				Name:     "environment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:     "application_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationName"),
			},
			{
				Name:     "date_created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DateCreated"),
			},
			{
				Name:     "date_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DateUpdated"),
			},
			{
				Name:     "deployment_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentStatus"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "environment_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnvironmentName"),
			},
			{
				Name:     "option_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OptionSettings"),
			},
			{
				Name:     "platform_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformArn"),
			},
			{
				Name:     "solution_stack_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SolutionStackName"),
			},
			{
				Name:     "template_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TemplateName"),
			},
			{
				Name:     "application_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationArn"),
			},
		},
	}
}
