// Code generated by codegen; DO NOT EDIT.

package images

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_images",
		Resolver: fetchImagesImages,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
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
			{
				Name:     "distribution",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Distribution"),
			},
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Slug"),
			},
			{
				Name:     "public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Public"),
			},
			{
				Name:     "regions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Regions"),
			},
			{
				Name:     "min_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinDiskSize"),
			},
			{
				Name:     "size_gigabytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("SizeGigaBytes"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
		},
	}
}
