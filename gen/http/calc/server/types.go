// Code generated by goa v3.4.2, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"

	goa "goa.design/goa/v3/pkg"
)

// DivDivByZeroResponseBody is the type of the "calc" service "div" endpoint
// HTTP response body for the "DivByZero" error.
type DivDivByZeroResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewDivDivByZeroResponseBody builds the HTTP response body from the result of
// the "div" endpoint of the "calc" service.
func NewDivDivByZeroResponseBody(res *goa.ServiceError) *DivDivByZeroResponseBody {
	body := &DivDivByZeroResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}

// NewDivPayload builds a calc service div endpoint payload.
func NewDivPayload(a int, b int) *calc.DivPayload {
	v := &calc.DivPayload{}
	v.A = a
	v.B = b

	return v
}