// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

//go:generate mockgen -package=mocks -destination=../mocks/mq.go -source=mq.go MqClient
type MqClient interface {
	DescribeBroker(context.Context, *mq.DescribeBrokerInput, ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error)
	DescribeBrokerEngineTypes(context.Context, *mq.DescribeBrokerEngineTypesInput, ...func(*mq.Options)) (*mq.DescribeBrokerEngineTypesOutput, error)
	DescribeBrokerInstanceOptions(context.Context, *mq.DescribeBrokerInstanceOptionsInput, ...func(*mq.Options)) (*mq.DescribeBrokerInstanceOptionsOutput, error)
	DescribeConfiguration(context.Context, *mq.DescribeConfigurationInput, ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationRevision(context.Context, *mq.DescribeConfigurationRevisionInput, ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error)
	DescribeUser(context.Context, *mq.DescribeUserInput, ...func(*mq.Options)) (*mq.DescribeUserOutput, error)
	ListBrokers(context.Context, *mq.ListBrokersInput, ...func(*mq.Options)) (*mq.ListBrokersOutput, error)
	ListConfigurationRevisions(context.Context, *mq.ListConfigurationRevisionsInput, ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error)
	ListConfigurations(context.Context, *mq.ListConfigurationsInput, ...func(*mq.Options)) (*mq.ListConfigurationsOutput, error)
	ListTags(context.Context, *mq.ListTagsInput, ...func(*mq.Options)) (*mq.ListTagsOutput, error)
	ListUsers(context.Context, *mq.ListUsersInput, ...func(*mq.Options)) (*mq.ListUsersOutput, error)
}
