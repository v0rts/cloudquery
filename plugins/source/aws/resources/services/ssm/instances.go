package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_instances",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html`,
		Resolver:    fetchSsmInstances,
		Transform:   transformers.TransformWithStruct(&types.InstanceInformation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveInstanceARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			InstanceComplianceItems(),
			InstancePatches(),
		},
	}
}
