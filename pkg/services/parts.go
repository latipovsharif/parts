package services

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PartService struct {}


func (c *PartService) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	return out, nil
}

func (c *PartService) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	return out, nil
}

func (c *PartService) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	return out, nil
}

func (c *PartService) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*PartResponse, error) {
	out := new(PartResponse)
	return out, nil
}