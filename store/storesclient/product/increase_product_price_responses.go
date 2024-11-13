// Code generated by go-swagger; DO NOT EDIT.

package product

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

// IncreaseProductPriceReader is a Reader for the IncreaseProductPrice structure.
type IncreaseProductPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IncreaseProductPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIncreaseProductPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIncreaseProductPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIncreaseProductPriceOK creates a IncreaseProductPriceOK with default headers values
func NewIncreaseProductPriceOK() *IncreaseProductPriceOK {
	return &IncreaseProductPriceOK{}
}

/*
IncreaseProductPriceOK describes a response with status code 200, with default header values.

A successful response.
*/
type IncreaseProductPriceOK struct {
	Payload models.V1IncreaseProductPriceResponse
}

// IsSuccess returns true when this increase product price o k response has a 2xx status code
func (o *IncreaseProductPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this increase product price o k response has a 3xx status code
func (o *IncreaseProductPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this increase product price o k response has a 4xx status code
func (o *IncreaseProductPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this increase product price o k response has a 5xx status code
func (o *IncreaseProductPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this increase product price o k response a status code equal to that given
func (o *IncreaseProductPriceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the increase product price o k response
func (o *IncreaseProductPriceOK) Code() int {
	return 200
}

func (o *IncreaseProductPriceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/products/{id}/increasePrice][%d] increaseProductPriceOK %s", 200, payload)
}

func (o *IncreaseProductPriceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/products/{id}/increasePrice][%d] increaseProductPriceOK %s", 200, payload)
}

func (o *IncreaseProductPriceOK) GetPayload() models.V1IncreaseProductPriceResponse {
	return o.Payload
}

func (o *IncreaseProductPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncreaseProductPriceDefault creates a IncreaseProductPriceDefault with default headers values
func NewIncreaseProductPriceDefault(code int) *IncreaseProductPriceDefault {
	return &IncreaseProductPriceDefault{
		_statusCode: code,
	}
}

/*
IncreaseProductPriceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IncreaseProductPriceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this increase product price default response has a 2xx status code
func (o *IncreaseProductPriceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this increase product price default response has a 3xx status code
func (o *IncreaseProductPriceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this increase product price default response has a 4xx status code
func (o *IncreaseProductPriceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this increase product price default response has a 5xx status code
func (o *IncreaseProductPriceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this increase product price default response a status code equal to that given
func (o *IncreaseProductPriceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the increase product price default response
func (o *IncreaseProductPriceDefault) Code() int {
	return o._statusCode
}

func (o *IncreaseProductPriceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/products/{id}/increasePrice][%d] increaseProductPrice default %s", o._statusCode, payload)
}

func (o *IncreaseProductPriceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/stores/products/{id}/increasePrice][%d] increaseProductPrice default %s", o._statusCode, payload)
}

func (o *IncreaseProductPriceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *IncreaseProductPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}