// Code generated by codegen; DO NOT EDIT.

package repositories

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Assets() *schema.Table {
	return &schema.Table{
		Name:      "github_release_assets",
		Resolver:  fetchAssets,
		Multiplex: client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "repository_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveGrandParentColumn("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "content_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentType"),
			},
			{
				Name:     "size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "download_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DownloadCount"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "browser_download_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrowserDownloadURL"),
			},
			{
				Name:     "uploader",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Uploader"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
		},
	}
}
