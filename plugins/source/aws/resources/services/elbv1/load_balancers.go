// Code generated by codegen; DO NOT EDIT.

package elbv1

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv1_load_balancers",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html`,
		Resolver:    fetchElbv1LoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
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
				Resolver: resolveLoadBalancerARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "backend_server_descriptions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackendServerDescriptions"),
			},
			{
				Name:     "canonical_hosted_zone_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CanonicalHostedZoneName"),
			},
			{
				Name:     "canonical_hosted_zone_name_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CanonicalHostedZoneNameID"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSName"),
			},
			{
				Name:     "health_check",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheck"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:     "listener_descriptions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ListenerDescriptions"),
			},
			{
				Name:     "load_balancer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancerName"),
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
			},
			{
				Name:     "scheme",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scheme"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "source_security_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceSecurityGroup"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VPCId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
			},
		},

		Relations: []*schema.Table{
			LoadBalancerPolicies(),
		},
	}
}
