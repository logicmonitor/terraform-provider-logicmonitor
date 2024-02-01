// Code generated by go-swagger; DO NOT EDIT.

package datasource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-logicmonitor/models"
)

// GetDatasourceByIDReader is a Reader for the GetDatasourceByID structure.
type GetDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatasourceByIDOK creates a GetDatasourceByIDOK with default headers values
func NewGetDatasourceByIDOK() *GetDatasourceByIDOK {
	return &GetDatasourceByIDOK{}
}

/*
GetDatasourceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDatasourceByIDOK struct {
	Payload *models.Datasource
}

func (o *GetDatasourceByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}][%d] getDatasourceByIdOK  %+v", 200, o.Payload)
}
func (o *GetDatasourceByIDOK) GetPayload() *models.Datasource {
	return o.Payload
}

func (o *GetDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datasource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceByIDDefault creates a GetDatasourceByIDDefault with default headers values
func NewGetDatasourceByIDDefault(code int) *GetDatasourceByIDDefault {
	return &GetDatasourceByIDDefault{
		_statusCode: code,
	}
}

/*
GetDatasourceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get datasource by Id default response
func (o *GetDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDatasourceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}][%d] getDatasourceById default  %+v", o._statusCode, o.Payload)
}
func (o *GetDatasourceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
