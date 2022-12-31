// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "google.golang.org/genproto/googleapis/cloud/location"

	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
)

func createMetadataLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeMetadataLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeMetadataLocationsRelationsServer{}
	aiplatformpb.RegisterMetadataServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeMetadataLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (f *fakeMetadataLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestMetadataLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, MetadataLocations(), createMetadataLocations, client.TestOptions{})
}

type fakeMetadataLocationsRelationsServer struct {
	aiplatformpb.UnimplementedMetadataServiceServer
}

func (f *fakeMetadataLocationsRelationsServer) ListArtifacts(context.Context, *aiplatformpb.ListArtifactsRequest) (*aiplatformpb.ListArtifactsResponse, error) {
	resp := aiplatformpb.ListArtifactsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeMetadataLocationsRelationsServer) ListContexts(context.Context, *aiplatformpb.ListContextsRequest) (*aiplatformpb.ListContextsResponse, error) {
	resp := aiplatformpb.ListContextsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeMetadataLocationsRelationsServer) ListExecutions(context.Context, *aiplatformpb.ListExecutionsRequest) (*aiplatformpb.ListExecutionsResponse, error) {
	resp := aiplatformpb.ListExecutionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeMetadataLocationsRelationsServer) ListMetadataSchemas(context.Context, *aiplatformpb.ListMetadataSchemasRequest) (*aiplatformpb.ListMetadataSchemasResponse, error) {
	resp := aiplatformpb.ListMetadataSchemasResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (f *fakeMetadataLocationsRelationsServer) ListMetadataStores(context.Context, *aiplatformpb.ListMetadataStoresRequest) (*aiplatformpb.ListMetadataStoresResponse, error) {
	resp := aiplatformpb.ListMetadataStoresResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
