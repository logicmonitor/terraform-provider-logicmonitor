// Code generated by go-swagger; DO NOT EDIT.

package data_resource_aws_external_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-logicmonitor/models"
)

// GetAwsExternalIDReader is a Reader for the GetAwsExternalID structure.
type GetAwsExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAwsExternalIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAwsExternalIDOK creates a GetAwsExternalIDOK with default headers values
func NewGetAwsExternalIDOK() *GetAwsExternalIDOK {
	return &GetAwsExternalIDOK{}
}

/* GetAwsExternalIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAwsExternalIDOK struct {
	Payload *models.AwsExternalID
}

func (o *GetAwsExternalIDOK) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalIdOK  %+v", 200, o.Payload)
}
func (o *GetAwsExternalIDOK) GetPayload() *models.AwsExternalID {
	return o.Payload
}

func (o *GetAwsExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsExternalID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsExternalIDDefault creates a GetAwsExternalIDDefault with default headers values
func NewGetAwsExternalIDDefault(code int) *GetAwsExternalIDDefault {
	return &GetAwsExternalIDDefault{
		_statusCode: code,
	}
}

/* GetAwsExternalIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAwsExternalIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get aws external Id default response
func (o *GetAwsExternalIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAwsExternalIDDefault) Error() string {
	return fmt.Sprintf("[GET /aws/externalId][%d] getAwsExternalId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAwsExternalIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAwsExternalIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}