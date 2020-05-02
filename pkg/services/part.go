package services

import (
	"fmt"

	"golang.org/x/net/context"
)

// Part service
type Part struct{}

// Create new part
func (ps *Part) Create(context.Context, *CreateRequest) (*Response, error) {
	fmt.Println("create part service")

	return &Response{
		Success: true,
		Message: "Part was created successfully",
	}, nil
}

// Update part by id
func (ps *Part) Update(context.Context, *UpdateRequest) (*Response, error) {
	fmt.Println("update part service")
	return nil, nil

}

// Delete part by id
func (ps *Part) Delete(context.Context, *DeleteRequest) (*Response, error) {
	fmt.Println("delete part service")
	return nil, nil

}

// Get part by id
func (ps *Part) Get(context.Context, *GetRequest) (*PartResponse, error) {
	fmt.Println("get part service")

	return nil, nil

}
