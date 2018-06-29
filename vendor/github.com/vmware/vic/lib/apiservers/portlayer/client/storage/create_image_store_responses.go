package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// CreateImageStoreReader is a Reader for the CreateImageStore structure.
type CreateImageStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImageStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateImageStoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateImageStoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateImageStoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateImageStoreCreated creates a CreateImageStoreCreated with default headers values
func NewCreateImageStoreCreated() *CreateImageStoreCreated {
	return &CreateImageStoreCreated{}
}

/*CreateImageStoreCreated handles this case with default header values.

Created
*/
type CreateImageStoreCreated struct {
	Payload *models.StoreURL
}

func (o *CreateImageStoreCreated) Error() string {
	return fmt.Sprintf("[POST /storage][%d] createImageStoreCreated  %+v", 201, o.Payload)
}

func (o *CreateImageStoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoreURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageStoreConflict creates a CreateImageStoreConflict with default headers values
func NewCreateImageStoreConflict() *CreateImageStoreConflict {
	return &CreateImageStoreConflict{}
}

/*CreateImageStoreConflict handles this case with default header values.

An image store with that name already exists.
*/
type CreateImageStoreConflict struct {
	Payload *models.Error
}

func (o *CreateImageStoreConflict) Error() string {
	return fmt.Sprintf("[POST /storage][%d] createImageStoreConflict  %+v", 409, o.Payload)
}

func (o *CreateImageStoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageStoreDefault creates a CreateImageStoreDefault with default headers values
func NewCreateImageStoreDefault(code int) *CreateImageStoreDefault {
	return &CreateImageStoreDefault{
		_statusCode: code,
	}
}

/*CreateImageStoreDefault handles this case with default header values.

error
*/
type CreateImageStoreDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create image store default response
func (o *CreateImageStoreDefault) Code() int {
	return o._statusCode
}

func (o *CreateImageStoreDefault) Error() string {
	return fmt.Sprintf("[POST /storage][%d] CreateImageStore default  %+v", o._statusCode, o.Payload)
}

func (o *CreateImageStoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}