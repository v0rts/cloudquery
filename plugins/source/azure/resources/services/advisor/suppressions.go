// Code generated by codegen2; DO NOT EDIT.
package advisor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Suppressions() *schema.Table {
	return &schema.Table{
		Name:      "azure_advisor_suppressions",
		Resolver:  fetchSuppressions,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_advisor),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSuppressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armadvisor.NewSuppressionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
