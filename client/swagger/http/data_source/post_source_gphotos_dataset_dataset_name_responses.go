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

// PostSourceGphotosDatasetDatasetNameReader is a Reader for the PostSourceGphotosDatasetDatasetName structure.
type PostSourceGphotosDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceGphotosDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceGphotosDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceGphotosDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceGphotosDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/gphotos/dataset/{datasetName}] PostSourceGphotosDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceGphotosDatasetDatasetNameOK creates a PostSourceGphotosDatasetDatasetNameOK with default headers values
func NewPostSourceGphotosDatasetDatasetNameOK() *PostSourceGphotosDatasetDatasetNameOK {
	return &PostSourceGphotosDatasetDatasetNameOK{}
}

/*
PostSourceGphotosDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceGphotosDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source gphotos dataset dataset name o k response has a 2xx status code
func (o *PostSourceGphotosDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source gphotos dataset dataset name o k response has a 3xx status code
func (o *PostSourceGphotosDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gphotos dataset dataset name o k response has a 4xx status code
func (o *PostSourceGphotosDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source gphotos dataset dataset name o k response has a 5xx status code
func (o *PostSourceGphotosDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source gphotos dataset dataset name o k response a status code equal to that given
func (o *PostSourceGphotosDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source gphotos dataset dataset name o k response
func (o *PostSourceGphotosDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceGphotosDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceGphotosDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceGphotosDatasetDatasetNameBadRequest creates a PostSourceGphotosDatasetDatasetNameBadRequest with default headers values
func NewPostSourceGphotosDatasetDatasetNameBadRequest() *PostSourceGphotosDatasetDatasetNameBadRequest {
	return &PostSourceGphotosDatasetDatasetNameBadRequest{}
}

/*
PostSourceGphotosDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceGphotosDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source gphotos dataset dataset name bad request response has a 2xx status code
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source gphotos dataset dataset name bad request response has a 3xx status code
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gphotos dataset dataset name bad request response has a 4xx status code
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source gphotos dataset dataset name bad request response has a 5xx status code
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source gphotos dataset dataset name bad request response a status code equal to that given
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source gphotos dataset dataset name bad request response
func (o *PostSourceGphotosDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceGphotosDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceGphotosDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceGphotosDatasetDatasetNameInternalServerError creates a PostSourceGphotosDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceGphotosDatasetDatasetNameInternalServerError() *PostSourceGphotosDatasetDatasetNameInternalServerError {
	return &PostSourceGphotosDatasetDatasetNameInternalServerError{}
}

/*
PostSourceGphotosDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceGphotosDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source gphotos dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source gphotos dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gphotos dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source gphotos dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source gphotos dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source gphotos dataset dataset name internal server error response
func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/gphotos/dataset/{datasetName}][%d] postSourceGphotosDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceGphotosDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
