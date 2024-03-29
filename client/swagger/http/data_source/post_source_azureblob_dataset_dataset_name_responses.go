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

// PostSourceAzureblobDatasetDatasetNameReader is a Reader for the PostSourceAzureblobDatasetDatasetName structure.
type PostSourceAzureblobDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceAzureblobDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceAzureblobDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceAzureblobDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceAzureblobDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/azureblob/dataset/{datasetName}] PostSourceAzureblobDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceAzureblobDatasetDatasetNameOK creates a PostSourceAzureblobDatasetDatasetNameOK with default headers values
func NewPostSourceAzureblobDatasetDatasetNameOK() *PostSourceAzureblobDatasetDatasetNameOK {
	return &PostSourceAzureblobDatasetDatasetNameOK{}
}

/*
PostSourceAzureblobDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceAzureblobDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source azureblob dataset dataset name o k response has a 2xx status code
func (o *PostSourceAzureblobDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source azureblob dataset dataset name o k response has a 3xx status code
func (o *PostSourceAzureblobDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source azureblob dataset dataset name o k response has a 4xx status code
func (o *PostSourceAzureblobDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source azureblob dataset dataset name o k response has a 5xx status code
func (o *PostSourceAzureblobDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source azureblob dataset dataset name o k response a status code equal to that given
func (o *PostSourceAzureblobDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source azureblob dataset dataset name o k response
func (o *PostSourceAzureblobDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceAzureblobDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceAzureblobDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceAzureblobDatasetDatasetNameBadRequest creates a PostSourceAzureblobDatasetDatasetNameBadRequest with default headers values
func NewPostSourceAzureblobDatasetDatasetNameBadRequest() *PostSourceAzureblobDatasetDatasetNameBadRequest {
	return &PostSourceAzureblobDatasetDatasetNameBadRequest{}
}

/*
PostSourceAzureblobDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceAzureblobDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source azureblob dataset dataset name bad request response has a 2xx status code
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source azureblob dataset dataset name bad request response has a 3xx status code
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source azureblob dataset dataset name bad request response has a 4xx status code
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source azureblob dataset dataset name bad request response has a 5xx status code
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source azureblob dataset dataset name bad request response a status code equal to that given
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source azureblob dataset dataset name bad request response
func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceAzureblobDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceAzureblobDatasetDatasetNameInternalServerError creates a PostSourceAzureblobDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceAzureblobDatasetDatasetNameInternalServerError() *PostSourceAzureblobDatasetDatasetNameInternalServerError {
	return &PostSourceAzureblobDatasetDatasetNameInternalServerError{}
}

/*
PostSourceAzureblobDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceAzureblobDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source azureblob dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source azureblob dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source azureblob dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source azureblob dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source azureblob dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source azureblob dataset dataset name internal server error response
func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/azureblob/dataset/{datasetName}][%d] postSourceAzureblobDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceAzureblobDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
