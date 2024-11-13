// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/storesclient/models"
)

// RebrandStoreReader is a Reader for the RebrandStore structure.
type RebrandStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebrandStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRebrandStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRebrandStoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRebrandStoreOK creates a RebrandStoreOK with default headers values
func NewRebrandStoreOK() *RebrandStoreOK {
	return &RebrandStoreOK{}
}

/*
RebrandStoreOK describes a response with status code 200, with default header values.

A successful response.
*/
type RebrandStoreOK struct {
	Payload models.V1RebrandStoreResponse
}

// IsSuccess returns true when this rebrand store o k response has a 2xx status code
func (o *RebrandStoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rebrand store o k response has a 3xx status code
func (o *RebrandStoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rebrand store o k response has a 4xx status code
func (o *RebrandStoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rebrand store o k response has a 5xx status code
func (o *RebrandStoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rebrand store o k response a status code equal to that given
func (o *RebrandStoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rebrand store o k response
func (o *RebrandStoreOK) Code() int {
	return 200
}

func (o *RebrandStoreOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/{id}/rebrand][%d] rebrandStoreOK %s", 200, payload)
}

func (o *RebrandStoreOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/{id}/rebrand][%d] rebrandStoreOK %s", 200, payload)
}

func (o *RebrandStoreOK) GetPayload() models.V1RebrandStoreResponse {
	return o.Payload
}

func (o *RebrandStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRebrandStoreDefault creates a RebrandStoreDefault with default headers values
func NewRebrandStoreDefault(code int) *RebrandStoreDefault {
	return &RebrandStoreDefault{
		_statusCode: code,
	}
}

/*
RebrandStoreDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RebrandStoreDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this rebrand store default response has a 2xx status code
func (o *RebrandStoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rebrand store default response has a 3xx status code
func (o *RebrandStoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rebrand store default response has a 4xx status code
func (o *RebrandStoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rebrand store default response has a 5xx status code
func (o *RebrandStoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rebrand store default response a status code equal to that given
func (o *RebrandStoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rebrand store default response
func (o *RebrandStoreDefault) Code() int {
	return o._statusCode
}

func (o *RebrandStoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/{id}/rebrand][%d] rebrandStore default %s", o._statusCode, payload)
}

func (o *RebrandStoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/{id}/rebrand][%d] rebrandStore default %s", o._statusCode, payload)
}

func (o *RebrandStoreDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RebrandStoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}