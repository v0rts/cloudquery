// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

//go:generate mockgen -package=mocks -destination=../mocks/appstream.go -source=appstream.go AppstreamClient
type AppstreamClient interface {
	DescribeAppBlocks(context.Context, *appstream.DescribeAppBlocksInput, ...func(*appstream.Options)) (*appstream.DescribeAppBlocksOutput, error)
	DescribeApplicationFleetAssociations(context.Context, *appstream.DescribeApplicationFleetAssociationsInput, ...func(*appstream.Options)) (*appstream.DescribeApplicationFleetAssociationsOutput, error)
	DescribeApplications(context.Context, *appstream.DescribeApplicationsInput, ...func(*appstream.Options)) (*appstream.DescribeApplicationsOutput, error)
	DescribeDirectoryConfigs(context.Context, *appstream.DescribeDirectoryConfigsInput, ...func(*appstream.Options)) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeEntitlements(context.Context, *appstream.DescribeEntitlementsInput, ...func(*appstream.Options)) (*appstream.DescribeEntitlementsOutput, error)
	DescribeFleets(context.Context, *appstream.DescribeFleetsInput, ...func(*appstream.Options)) (*appstream.DescribeFleetsOutput, error)
	DescribeImageBuilders(context.Context, *appstream.DescribeImageBuildersInput, ...func(*appstream.Options)) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImagePermissions(context.Context, *appstream.DescribeImagePermissionsInput, ...func(*appstream.Options)) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImages(context.Context, *appstream.DescribeImagesInput, ...func(*appstream.Options)) (*appstream.DescribeImagesOutput, error)
	DescribeSessions(context.Context, *appstream.DescribeSessionsInput, ...func(*appstream.Options)) (*appstream.DescribeSessionsOutput, error)
	DescribeStacks(context.Context, *appstream.DescribeStacksInput, ...func(*appstream.Options)) (*appstream.DescribeStacksOutput, error)
	DescribeUsageReportSubscriptions(context.Context, *appstream.DescribeUsageReportSubscriptionsInput, ...func(*appstream.Options)) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUserStackAssociations(context.Context, *appstream.DescribeUserStackAssociationsInput, ...func(*appstream.Options)) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUsers(context.Context, *appstream.DescribeUsersInput, ...func(*appstream.Options)) (*appstream.DescribeUsersOutput, error)
	ListAssociatedFleets(context.Context, *appstream.ListAssociatedFleetsInput, ...func(*appstream.Options)) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedStacks(context.Context, *appstream.ListAssociatedStacksInput, ...func(*appstream.Options)) (*appstream.ListAssociatedStacksOutput, error)
	ListEntitledApplications(context.Context, *appstream.ListEntitledApplicationsInput, ...func(*appstream.Options)) (*appstream.ListEntitledApplicationsOutput, error)
	ListTagsForResource(context.Context, *appstream.ListTagsForResourceInput, ...func(*appstream.Options)) (*appstream.ListTagsForResourceOutput, error)
}
