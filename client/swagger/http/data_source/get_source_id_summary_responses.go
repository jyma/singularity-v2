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

// GetSourceIDSummaryReader is a Reader for the GetSourceIDSummary structure.
type GetSourceIDSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceIDSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceIDSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSourceIDSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSourceIDSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /source/{id}/summary] GetSourceIDSummary", response, response.Code())
	}
}

// NewGetSourceIDSummaryOK creates a GetSourceIDSummaryOK with default headers values
func NewGetSourceIDSummaryOK() *GetSourceIDSummaryOK {
	return &GetSourceIDSummaryOK{}
}

/*
GetSourceIDSummaryOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceIDSummaryOK struct {
	Payload *models.DatasourceChunksByState
}

// IsSuccess returns true when this get source Id summary o k response has a 2xx status code
func (o *GetSourceIDSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get source Id summary o k response has a 3xx status code
func (o *GetSourceIDSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id summary o k response has a 4xx status code
func (o *GetSourceIDSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id summary o k response has a 5xx status code
func (o *GetSourceIDSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id summary o k response a status code equal to that given
func (o *GetSourceIDSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get source Id summary o k response
func (o *GetSourceIDSummaryOK) Code() int {
	return 200
}

func (o *GetSourceIDSummaryOK) Error() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDSummaryOK) String() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDSummaryOK) GetPayload() *models.DatasourceChunksByState {
	return o.Payload
}

func (o *GetSourceIDSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasourceChunksByState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDSummaryBadRequest creates a GetSourceIDSummaryBadRequest with default headers values
func NewGetSourceIDSummaryBadRequest() *GetSourceIDSummaryBadRequest {
	return &GetSourceIDSummaryBadRequest{}
}

/*
GetSourceIDSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSourceIDSummaryBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id summary bad request response has a 2xx status code
func (o *GetSourceIDSummaryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id summary bad request response has a 3xx status code
func (o *GetSourceIDSummaryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id summary bad request response has a 4xx status code
func (o *GetSourceIDSummaryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get source Id summary bad request response has a 5xx status code
func (o *GetSourceIDSummaryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id summary bad request response a status code equal to that given
func (o *GetSourceIDSummaryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get source Id summary bad request response
func (o *GetSourceIDSummaryBadRequest) Code() int {
	return 400
}

func (o *GetSourceIDSummaryBadRequest) Error() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDSummaryBadRequest) String() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDSummaryBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDSummaryInternalServerError creates a GetSourceIDSummaryInternalServerError with default headers values
func NewGetSourceIDSummaryInternalServerError() *GetSourceIDSummaryInternalServerError {
	return &GetSourceIDSummaryInternalServerError{}
}

/*
GetSourceIDSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSourceIDSummaryInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id summary internal server error response has a 2xx status code
func (o *GetSourceIDSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id summary internal server error response has a 3xx status code
func (o *GetSourceIDSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id summary internal server error response has a 4xx status code
func (o *GetSourceIDSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id summary internal server error response has a 5xx status code
func (o *GetSourceIDSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get source Id summary internal server error response a status code equal to that given
func (o *GetSourceIDSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get source Id summary internal server error response
func (o *GetSourceIDSummaryInternalServerError) Code() int {
	return 500
}

func (o *GetSourceIDSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDSummaryInternalServerError) String() string {
	return fmt.Sprintf("[GET /source/{id}/summary][%d] getSourceIdSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDSummaryInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
