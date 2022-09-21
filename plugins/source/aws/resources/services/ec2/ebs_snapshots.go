// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EbsSnapshots() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_ebs_snapshots",
		Resolver:  fetchEc2EbsSnapshots,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
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
				Resolver: resolveEbsSnapshotArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attribute",
				Type:     schema.TypeString,
				Resolver: resolveEbsSnapshotAttribute,
			},
			{
				Name:     "data_encryption_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataEncryptionKeyId"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Encrypted"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "owner_alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAlias"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "progress",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Progress"),
			},
			{
				Name:     "restore_expiry_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RestoreExpiryTime"),
			},
			{
				Name:     "snapshot_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotId"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateMessage"),
			},
			{
				Name:     "storage_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageTier"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "volume_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeId"),
			},
			{
				Name:     "volume_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VolumeSize"),
			},
		},
	}
}
