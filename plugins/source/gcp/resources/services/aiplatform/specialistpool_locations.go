// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/location"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"cloud.google.com/go/aiplatform/apiv1"
)

func SpecialistpoolLocations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_specialistpool_locations",
		Description: `https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location`,
		Resolver:    fetchSpecialistpoolLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationId"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
		},

		Relations: []*schema.Table{
			SpecialistPools(),
		},
	}
}

func fetchSpecialistpoolLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListLocationsRequest{
		Name: "projects/" + c.ProjectId,
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewSpecialistPoolClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListLocations(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
