// Code generated by codegen; DO NOT EDIT.

package backup

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VaultRecoveryPoints() *schema.Table {
	return &schema.Table{
		Name:      "aws_backup_vault_recovery_points",
		Resolver:  fetchBackupVaultRecoveryPoints,
		Multiplex: client.ServiceAccountRegionMultiplexer("backup"),
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
				Name:     "vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRecoveryPointTags,
			},
			{
				Name:     "backup_size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BackupSizeInBytes"),
			},
			{
				Name:     "backup_vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultArn"),
			},
			{
				Name:     "backup_vault_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultName"),
			},
			{
				Name:     "calculated_lifecycle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CalculatedLifecycle"),
			},
			{
				Name:     "completion_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletionDate"),
			},
			{
				Name:     "created_by",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "encryption_key_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionKeyArn"),
			},
			{
				Name:     "iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleArn"),
			},
			{
				Name:     "is_encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsEncrypted"),
			},
			{
				Name:     "last_restore_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastRestoreTime"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "recovery_point_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoveryPointArn"),
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceArn"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "source_backup_vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceBackupVaultArn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
		},
	}
}
