// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetSourceIDItemsReader is a Reader for the GetSourceIDItems structure.
type GetSourceIDItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceIDItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceIDItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSourceIDItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSourceIDItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /source/{id}/items] GetSourceIDItems", response, response.Code())
	}
}

// NewGetSourceIDItemsOK creates a GetSourceIDItemsOK with default headers values
func NewGetSourceIDItemsOK() *GetSourceIDItemsOK {
	return &GetSourceIDItemsOK{}
}

/*
GetSourceIDItemsOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceIDItemsOK struct {
	Payload []*models.ModelItem
}

// IsSuccess returns true when this get source Id items o k response has a 2xx status code
func (o *GetSourceIDItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get source Id items o k response has a 3xx status code
func (o *GetSourceIDItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id items o k response has a 4xx status code
func (o *GetSourceIDItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id items o k response has a 5xx status code
func (o *GetSourceIDItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id items o k response a status code equal to that given
func (o *GetSourceIDItemsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get source Id items o k response
func (o *GetSourceIDItemsOK) Code() int {
	return 200
}

func (o *GetSourceIDItemsOK) Error() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDItemsOK) String() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDItemsOK) GetPayload() []*models.ModelItem {
	return o.Payload
}

func (o *GetSourceIDItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDItemsBadRequest creates a GetSourceIDItemsBadRequest with default headers values
func NewGetSourceIDItemsBadRequest() *GetSourceIDItemsBadRequest {
	return &GetSourceIDItemsBadRequest{}
}

/*
GetSourceIDItemsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSourceIDItemsBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id items bad request response has a 2xx status code
func (o *GetSourceIDItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id items bad request response has a 3xx status code
func (o *GetSourceIDItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id items bad request response has a 4xx status code
func (o *GetSourceIDItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get source Id items bad request response has a 5xx status code
func (o *GetSourceIDItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id items bad request response a status code equal to that given
func (o *GetSourceIDItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get source Id items bad request response
func (o *GetSourceIDItemsBadRequest) Code() int {
	return 400
}

func (o *GetSourceIDItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDItemsBadRequest) String() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDItemsBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDItemsInternalServerError creates a GetSourceIDItemsInternalServerError with default headers values
func NewGetSourceIDItemsInternalServerError() *GetSourceIDItemsInternalServerError {
	return &GetSourceIDItemsInternalServerError{}
}

/*
GetSourceIDItemsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSourceIDItemsInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id items internal server error response has a 2xx status code
func (o *GetSourceIDItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id items internal server error response has a 3xx status code
func (o *GetSourceIDItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id items internal server error response has a 4xx status code
func (o *GetSourceIDItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id items internal server error response has a 5xx status code
func (o *GetSourceIDItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get source Id items internal server error response a status code equal to that given
func (o *GetSourceIDItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get source Id items internal server error response
func (o *GetSourceIDItemsInternalServerError) Code() int {
	return 500
}

func (o *GetSourceIDItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDItemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /source/{id}/items][%d] getSourceIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDItemsInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
