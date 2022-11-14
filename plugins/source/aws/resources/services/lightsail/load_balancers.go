// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_load_balancers",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancer.html`,
		Resolver:    fetchLightsailLoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "configuration_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfigurationOptions"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnsName"),
			},
			{
				Name:     "health_check_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckPath"),
			},
			{
				Name:     "https_redirection_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HttpsRedirectionEnabled"),
			},
			{
				Name:     "instance_health_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceHealthSummary"),
			},
			{
				Name:     "instance_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InstancePort"),
			},
			{
				Name:     "ip_address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddressType"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Protocol"),
			},
			{
				Name:     "public_ports",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("PublicPorts"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tls_certificate_summaries",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TlsCertificateSummaries"),
			},
			{
				Name:     "tls_policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TlsPolicyName"),
			},
		},

		Relations: []*schema.Table{
			LoadBalancerTlsCertificates(),
		},
	}
}
