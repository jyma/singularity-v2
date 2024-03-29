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

// PostSourceOpendriveDatasetDatasetNameReader is a Reader for the PostSourceOpendriveDatasetDatasetName structure.
type PostSourceOpendriveDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceOpendriveDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceOpendriveDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceOpendriveDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceOpendriveDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/opendrive/dataset/{datasetName}] PostSourceOpendriveDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceOpendriveDatasetDatasetNameOK creates a PostSourceOpendriveDatasetDatasetNameOK with default headers values
func NewPostSourceOpendriveDatasetDatasetNameOK() *PostSourceOpendriveDatasetDatasetNameOK {
	return &PostSourceOpendriveDatasetDatasetNameOK{}
}

/*
PostSourceOpendriveDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceOpendriveDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source opendrive dataset dataset name o k response has a 2xx status code
func (o *PostSourceOpendriveDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source opendrive dataset dataset name o k response has a 3xx status code
func (o *PostSourceOpendriveDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source opendrive dataset dataset name o k response has a 4xx status code
func (o *PostSourceOpendriveDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source opendrive dataset dataset name o k response has a 5xx status code
func (o *PostSourceOpendriveDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source opendrive dataset dataset name o k response a status code equal to that given
func (o *PostSourceOpendriveDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source opendrive dataset dataset name o k response
func (o *PostSourceOpendriveDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceOpendriveDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceOpendriveDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceOpendriveDatasetDatasetNameBadRequest creates a PostSourceOpendriveDatasetDatasetNameBadRequest with default headers values
func NewPostSourceOpendriveDatasetDatasetNameBadRequest() *PostSourceOpendriveDatasetDatasetNameBadRequest {
	return &PostSourceOpendriveDatasetDatasetNameBadRequest{}
}

/*
PostSourceOpendriveDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceOpendriveDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source opendrive dataset dataset name bad request response has a 2xx status code
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source opendrive dataset dataset name bad request response has a 3xx status code
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source opendrive dataset dataset name bad request response has a 4xx status code
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source opendrive dataset dataset name bad request response has a 5xx status code
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source opendrive dataset dataset name bad request response a status code equal to that given
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source opendrive dataset dataset name bad request response
func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceOpendriveDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceOpendriveDatasetDatasetNameInternalServerError creates a PostSourceOpendriveDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceOpendriveDatasetDatasetNameInternalServerError() *PostSourceOpendriveDatasetDatasetNameInternalServerError {
	return &PostSourceOpendriveDatasetDatasetNameInternalServerError{}
}

/*
PostSourceOpendriveDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceOpendriveDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source opendrive dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source opendrive dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source opendrive dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source opendrive dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source opendrive dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source opendrive dataset dataset name internal server error response
func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/opendrive/dataset/{datasetName}][%d] postSourceOpendriveDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceOpendriveDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
