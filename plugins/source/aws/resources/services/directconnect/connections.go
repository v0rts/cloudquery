// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Connections() *schema.Table {
	return &schema.Table{
		Name:      "aws_directconnect_connections",
		Resolver:  fetchDirectconnectConnections,
		Multiplex: client.ServiceAccountRegionMultiplexer("directconnect"),
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
				Resolver: resolveConnectionARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "aws_device",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsDevice"),
			},
			{
				Name:     "aws_device_v_2",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsDeviceV2"),
			},
			{
				Name:     "aws_logical_device_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsLogicalDeviceId"),
			},
			{
				Name:     "bandwidth",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Bandwidth"),
			},
			{
				Name:     "connection_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionName"),
			},
			{
				Name:     "connection_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionState"),
			},
			{
				Name:     "encryption_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionMode"),
			},
			{
				Name:     "has_logical_redundancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HasLogicalRedundancy"),
			},
			{
				Name:     "jumbo_frame_capable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("JumboFrameCapable"),
			},
			{
				Name:     "lag_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LagId"),
			},
			{
				Name:     "loa_issue_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LoaIssueTime"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "mac_sec_capable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MacSecCapable"),
			},
			{
				Name:     "mac_sec_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MacSecKeys"),
			},
			{
				Name:     "owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAccount"),
			},
			{
				Name:     "partner_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PartnerName"),
			},
			{
				Name:     "port_encryption_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PortEncryptionStatus"),
			},
			{
				Name:     "provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderName"),
			},
			{
				Name:     "vlan",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Vlan"),
			},
		},
	}
}
