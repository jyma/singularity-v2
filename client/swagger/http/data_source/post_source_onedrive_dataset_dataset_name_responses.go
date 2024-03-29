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

// PostSourceOnedriveDatasetDatasetNameReader is a Reader for the PostSourceOnedriveDatasetDatasetName structure.
type PostSourceOnedriveDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceOnedriveDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceOnedriveDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceOnedriveDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceOnedriveDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/onedrive/dataset/{datasetName}] PostSourceOnedriveDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceOnedriveDatasetDatasetNameOK creates a PostSourceOnedriveDatasetDatasetNameOK with default headers values
func NewPostSourceOnedriveDatasetDatasetNameOK() *PostSourceOnedriveDatasetDatasetNameOK {
	return &PostSourceOnedriveDatasetDatasetNameOK{}
}

/*
PostSourceOnedriveDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceOnedriveDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source onedrive dataset dataset name o k response has a 2xx status code
func (o *PostSourceOnedriveDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source onedrive dataset dataset name o k response has a 3xx status code
func (o *PostSourceOnedriveDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source onedrive dataset dataset name o k response has a 4xx status code
func (o *PostSourceOnedriveDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source onedrive dataset dataset name o k response has a 5xx status code
func (o *PostSourceOnedriveDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source onedrive dataset dataset name o k response a status code equal to that given
func (o *PostSourceOnedriveDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source onedrive dataset dataset name o k response
func (o *PostSourceOnedriveDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceOnedriveDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceOnedriveDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceOnedriveDatasetDatasetNameBadRequest creates a PostSourceOnedriveDatasetDatasetNameBadRequest with default headers values
func NewPostSourceOnedriveDatasetDatasetNameBadRequest() *PostSourceOnedriveDatasetDatasetNameBadRequest {
	return &PostSourceOnedriveDatasetDatasetNameBadRequest{}
}

/*
PostSourceOnedriveDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceOnedriveDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source onedrive dataset dataset name bad request response has a 2xx status code
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source onedrive dataset dataset name bad request response has a 3xx status code
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source onedrive dataset dataset name bad request response has a 4xx status code
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source onedrive dataset dataset name bad request response has a 5xx status code
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source onedrive dataset dataset name bad request response a status code equal to that given
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source onedrive dataset dataset name bad request response
func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceOnedriveDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceOnedriveDatasetDatasetNameInternalServerError creates a PostSourceOnedriveDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceOnedriveDatasetDatasetNameInternalServerError() *PostSourceOnedriveDatasetDatasetNameInternalServerError {
	return &PostSourceOnedriveDatasetDatasetNameInternalServerError{}
}

/*
PostSourceOnedriveDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceOnedriveDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source onedrive dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source onedrive dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source onedrive dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source onedrive dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source onedrive dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source onedrive dataset dataset name internal server error response
func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/onedrive/dataset/{datasetName}][%d] postSourceOnedriveDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceOnedriveDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
