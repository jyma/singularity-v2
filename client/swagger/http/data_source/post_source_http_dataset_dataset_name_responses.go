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

// PostSourceHTTPDatasetDatasetNameReader is a Reader for the PostSourceHTTPDatasetDatasetName structure.
type PostSourceHTTPDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceHTTPDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceHTTPDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceHTTPDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceHTTPDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/http/dataset/{datasetName}] PostSourceHTTPDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceHTTPDatasetDatasetNameOK creates a PostSourceHTTPDatasetDatasetNameOK with default headers values
func NewPostSourceHTTPDatasetDatasetNameOK() *PostSourceHTTPDatasetDatasetNameOK {
	return &PostSourceHTTPDatasetDatasetNameOK{}
}

/*
PostSourceHTTPDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceHTTPDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source Http dataset dataset name o k response has a 2xx status code
func (o *PostSourceHTTPDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source Http dataset dataset name o k response has a 3xx status code
func (o *PostSourceHTTPDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Http dataset dataset name o k response has a 4xx status code
func (o *PostSourceHTTPDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Http dataset dataset name o k response has a 5xx status code
func (o *PostSourceHTTPDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Http dataset dataset name o k response a status code equal to that given
func (o *PostSourceHTTPDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source Http dataset dataset name o k response
func (o *PostSourceHTTPDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceHTTPDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceHTTPDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceHTTPDatasetDatasetNameBadRequest creates a PostSourceHTTPDatasetDatasetNameBadRequest with default headers values
func NewPostSourceHTTPDatasetDatasetNameBadRequest() *PostSourceHTTPDatasetDatasetNameBadRequest {
	return &PostSourceHTTPDatasetDatasetNameBadRequest{}
}

/*
PostSourceHTTPDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceHTTPDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source Http dataset dataset name bad request response has a 2xx status code
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Http dataset dataset name bad request response has a 3xx status code
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Http dataset dataset name bad request response has a 4xx status code
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source Http dataset dataset name bad request response has a 5xx status code
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Http dataset dataset name bad request response a status code equal to that given
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source Http dataset dataset name bad request response
func (o *PostSourceHTTPDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceHTTPDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceHTTPDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceHTTPDatasetDatasetNameInternalServerError creates a PostSourceHTTPDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceHTTPDatasetDatasetNameInternalServerError() *PostSourceHTTPDatasetDatasetNameInternalServerError {
	return &PostSourceHTTPDatasetDatasetNameInternalServerError{}
}

/*
PostSourceHTTPDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceHTTPDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source Http dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Http dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Http dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Http dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source Http dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source Http dataset dataset name internal server error response
func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/http/dataset/{datasetName}][%d] postSourceHttpDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceHTTPDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
