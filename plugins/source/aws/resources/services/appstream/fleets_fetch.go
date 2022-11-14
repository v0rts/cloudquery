// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppstreamFleets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeFleetsInput

	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeFleets(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Fleets
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
