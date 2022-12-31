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

func createIndexLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeIndexLocationsServer{}
	pb.RegisterLocationsServer(gsrv, fakeServer)
	fakeRelationsServer := &fakeIndexLocationsRelationsServer{}
	aiplatformpb.RegisterIndexServiceServer(gsrv, fakeRelationsServer)

	return nil
}

type fakeIndexLocationsServer struct {
	pb.UnimplementedLocationsServer
}

func (f *fakeIndexLocationsServer) ListLocations(context.Context, *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	resp := pb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestIndexLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, IndexLocations(), createIndexLocations, client.TestOptions{})
}

type fakeIndexLocationsRelationsServer struct {
	aiplatformpb.UnimplementedIndexServiceServer
}

func (f *fakeIndexLocationsRelationsServer) ListIndexes(context.Context, *aiplatformpb.ListIndexesRequest) (*aiplatformpb.ListIndexesResponse, error) {
	resp := aiplatformpb.ListIndexesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}
