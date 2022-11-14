// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppstreamUsageReportSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeUsageReportSubscriptionsInput

	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeUsageReportSubscriptions(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.UsageReportSubscriptions
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
