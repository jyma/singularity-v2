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

// PostSourceSwiftDatasetDatasetNameReader is a Reader for the PostSourceSwiftDatasetDatasetName structure.
type PostSourceSwiftDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceSwiftDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceSwiftDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceSwiftDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceSwiftDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/swift/dataset/{datasetName}] PostSourceSwiftDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceSwiftDatasetDatasetNameOK creates a PostSourceSwiftDatasetDatasetNameOK with default headers values
func NewPostSourceSwiftDatasetDatasetNameOK() *PostSourceSwiftDatasetDatasetNameOK {
	return &PostSourceSwiftDatasetDatasetNameOK{}
}

/*
PostSourceSwiftDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceSwiftDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source swift dataset dataset name o k response has a 2xx status code
func (o *PostSourceSwiftDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source swift dataset dataset name o k response has a 3xx status code
func (o *PostSourceSwiftDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source swift dataset dataset name o k response has a 4xx status code
func (o *PostSourceSwiftDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source swift dataset dataset name o k response has a 5xx status code
func (o *PostSourceSwiftDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source swift dataset dataset name o k response a status code equal to that given
func (o *PostSourceSwiftDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source swift dataset dataset name o k response
func (o *PostSourceSwiftDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceSwiftDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceSwiftDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceSwiftDatasetDatasetNameBadRequest creates a PostSourceSwiftDatasetDatasetNameBadRequest with default headers values
func NewPostSourceSwiftDatasetDatasetNameBadRequest() *PostSourceSwiftDatasetDatasetNameBadRequest {
	return &PostSourceSwiftDatasetDatasetNameBadRequest{}
}

/*
PostSourceSwiftDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceSwiftDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source swift dataset dataset name bad request response has a 2xx status code
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source swift dataset dataset name bad request response has a 3xx status code
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source swift dataset dataset name bad request response has a 4xx status code
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source swift dataset dataset name bad request response has a 5xx status code
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source swift dataset dataset name bad request response a status code equal to that given
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source swift dataset dataset name bad request response
func (o *PostSourceSwiftDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceSwiftDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceSwiftDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceSwiftDatasetDatasetNameInternalServerError creates a PostSourceSwiftDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceSwiftDatasetDatasetNameInternalServerError() *PostSourceSwiftDatasetDatasetNameInternalServerError {
	return &PostSourceSwiftDatasetDatasetNameInternalServerError{}
}

/*
PostSourceSwiftDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceSwiftDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source swift dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source swift dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source swift dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source swift dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source swift dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source swift dataset dataset name internal server error response
func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/swift/dataset/{datasetName}][%d] postSourceSwiftDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceSwiftDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
