// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func AddOnAttachments() *schema.Table {
	return &schema.Table{
		Name:        "heroku_add_on_attachments",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment`,
		Resolver:    fetchAddOnAttachments,
		Columns: []schema.Column{
			{
				Name:     "addon",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Addon"),
			},
			{
				Name:     "app",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("App"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
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
				Name:     "log_input_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogInputURL"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "web_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebURL"),
			},
		},
	}
}

func fetchAddOnAttachments(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AddOnAttachmentList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
