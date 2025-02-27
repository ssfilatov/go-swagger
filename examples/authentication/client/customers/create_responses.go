// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ssfilatov/go-swagger/examples/authentication/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCreated creates a CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {
	return &CreateCreated{}
}

/* CreateCreated describes a response with status code 201, with default header values.

created
*/
type CreateCreated struct {
	Payload *models.Customer
}

func (o *CreateCreated) Error() string {
	return fmt.Sprintf("[POST /customers][%d] createCreated  %+v", 201, o.Payload)
}
func (o *CreateCreated) GetPayload() *models.Customer {
	return o.Payload
}

func (o *CreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefault creates a CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	return &CreateDefault{
		_statusCode: code,
	}
}

/* CreateDefault describes a response with status code -1, with default header values.

error
*/
type CreateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create default response
func (o *CreateDefault) Code() int {
	return o._statusCode
}

func (o *CreateDefault) Error() string {
	return fmt.Sprintf("[POST /customers][%d] create default  %+v", o._statusCode, o.Payload)
}
func (o *CreateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
