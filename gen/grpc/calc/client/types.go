// Code generated by goa v3.4.2, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen calc/design

package client

import (
	calc "calc/gen/calc"
	calcpb "calc/gen/grpc/calc/pb"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "calc" service.
func NewAddRequest(payload *calc.AddPayload) *calcpb.AddRequest {
	message := &calcpb.AddRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "calc"
// service from the gRPC response type.
func NewAddResult(message *calcpb.AddResponse) int {
	result := int(message.Field)
	return result
}

// NewDivRequest builds the gRPC request type from the payload of the "div"
// endpoint of the "calc" service.
func NewDivRequest(payload *calc.DivPayload) *calcpb.DivRequest {
	message := &calcpb.DivRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewDivResult builds the result type of the "div" endpoint of the "calc"
// service from the gRPC response type.
func NewDivResult(message *calcpb.DivResponse) int {
	result := int(message.Field)
	return result
}