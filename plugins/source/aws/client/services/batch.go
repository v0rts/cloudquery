// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

//go:generate mockgen -package=mocks -destination=../mocks/batch.go -source=batch.go BatchClient
type BatchClient interface {
	DescribeComputeEnvironments(context.Context, *batch.DescribeComputeEnvironmentsInput, ...func(*batch.Options)) (*batch.DescribeComputeEnvironmentsOutput, error)
	DescribeJobDefinitions(context.Context, *batch.DescribeJobDefinitionsInput, ...func(*batch.Options)) (*batch.DescribeJobDefinitionsOutput, error)
	DescribeJobQueues(context.Context, *batch.DescribeJobQueuesInput, ...func(*batch.Options)) (*batch.DescribeJobQueuesOutput, error)
	DescribeJobs(context.Context, *batch.DescribeJobsInput, ...func(*batch.Options)) (*batch.DescribeJobsOutput, error)
	DescribeSchedulingPolicies(context.Context, *batch.DescribeSchedulingPoliciesInput, ...func(*batch.Options)) (*batch.DescribeSchedulingPoliciesOutput, error)
	ListJobs(context.Context, *batch.ListJobsInput, ...func(*batch.Options)) (*batch.ListJobsOutput, error)
	ListSchedulingPolicies(context.Context, *batch.ListSchedulingPoliciesInput, ...func(*batch.Options)) (*batch.ListSchedulingPoliciesOutput, error)
	ListTagsForResource(context.Context, *batch.ListTagsForResourceInput, ...func(*batch.Options)) (*batch.ListTagsForResourceOutput, error)
}
