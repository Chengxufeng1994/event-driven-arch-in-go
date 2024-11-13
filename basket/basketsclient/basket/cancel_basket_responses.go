// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/basketsclient/models"
)

// CancelBasketReader is a Reader for the CancelBasket structure.
type CancelBasketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelBasketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelBasketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelBasketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelBasketOK creates a CancelBasketOK with default headers values
func NewCancelBasketOK() *CancelBasketOK {
	return &CancelBasketOK{}
}

/*
CancelBasketOK describes a response with status code 200, with default header values.

A successful response.
*/
type CancelBasketOK struct {
	Payload models.V1CancelBasketResponse
}

// IsSuccess returns true when this cancel basket o k response has a 2xx status code
func (o *CancelBasketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel basket o k response has a 3xx status code
func (o *CancelBasketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel basket o k response has a 4xx status code
func (o *CancelBasketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel basket o k response has a 5xx status code
func (o *CancelBasketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel basket o k response a status code equal to that given
func (o *CancelBasketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel basket o k response
func (o *CancelBasketOK) Code() int {
	return 200
}

func (o *CancelBasketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/baskets/{id}][%d] cancelBasketOK %s", 200, payload)
}

func (o *CancelBasketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/baskets/{id}][%d] cancelBasketOK %s", 200, payload)
}

func (o *CancelBasketOK) GetPayload() models.V1CancelBasketResponse {
	return o.Payload
}

func (o *CancelBasketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelBasketDefault creates a CancelBasketDefault with default headers values
func NewCancelBasketDefault(code int) *CancelBasketDefault {
	return &CancelBasketDefault{
		_statusCode: code,
	}
}

/*
CancelBasketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CancelBasketDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this cancel basket default response has a 2xx status code
func (o *CancelBasketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel basket default response has a 3xx status code
func (o *CancelBasketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel basket default response has a 4xx status code
func (o *CancelBasketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel basket default response has a 5xx status code
func (o *CancelBasketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel basket default response a status code equal to that given
func (o *CancelBasketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cancel basket default response
func (o *CancelBasketDefault) Code() int {
	return o._statusCode
}

func (o *CancelBasketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/baskets/{id}][%d] cancelBasket default %s", o._statusCode, payload)
}

func (o *CancelBasketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/baskets/{id}][%d] cancelBasket default %s", o._statusCode, payload)
}

func (o *CancelBasketDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CancelBasketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
