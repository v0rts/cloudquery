// Code generated by codegen; DO NOT EDIT.

package account

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountEvents() *schema.Table {
	return &schema.Table{
		Name:        "fastly_account_events",
		Description: `https://developer.fastly.com/reference/api/account/events/`,
		Resolver:    fetchAccountEvents,
		Columns: []schema.Column{
			{
				Name:     "admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Admin"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "customer_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerID"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "event_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventType"),
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
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IP"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceID"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserID"),
			},
		},
	}
}
