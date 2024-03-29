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

// PostSourcePremiumizemeDatasetDatasetNameReader is a Reader for the PostSourcePremiumizemeDatasetDatasetName structure.
type PostSourcePremiumizemeDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourcePremiumizemeDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourcePremiumizemeDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourcePremiumizemeDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourcePremiumizemeDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/premiumizeme/dataset/{datasetName}] PostSourcePremiumizemeDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourcePremiumizemeDatasetDatasetNameOK creates a PostSourcePremiumizemeDatasetDatasetNameOK with default headers values
func NewPostSourcePremiumizemeDatasetDatasetNameOK() *PostSourcePremiumizemeDatasetDatasetNameOK {
	return &PostSourcePremiumizemeDatasetDatasetNameOK{}
}

/*
PostSourcePremiumizemeDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourcePremiumizemeDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source premiumizeme dataset dataset name o k response has a 2xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source premiumizeme dataset dataset name o k response has a 3xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source premiumizeme dataset dataset name o k response has a 4xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source premiumizeme dataset dataset name o k response has a 5xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source premiumizeme dataset dataset name o k response a status code equal to that given
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source premiumizeme dataset dataset name o k response
func (o *PostSourcePremiumizemeDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourcePremiumizemeDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourcePremiumizemeDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourcePremiumizemeDatasetDatasetNameBadRequest creates a PostSourcePremiumizemeDatasetDatasetNameBadRequest with default headers values
func NewPostSourcePremiumizemeDatasetDatasetNameBadRequest() *PostSourcePremiumizemeDatasetDatasetNameBadRequest {
	return &PostSourcePremiumizemeDatasetDatasetNameBadRequest{}
}

/*
PostSourcePremiumizemeDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourcePremiumizemeDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source premiumizeme dataset dataset name bad request response has a 2xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source premiumizeme dataset dataset name bad request response has a 3xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source premiumizeme dataset dataset name bad request response has a 4xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source premiumizeme dataset dataset name bad request response has a 5xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source premiumizeme dataset dataset name bad request response a status code equal to that given
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source premiumizeme dataset dataset name bad request response
func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourcePremiumizemeDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourcePremiumizemeDatasetDatasetNameInternalServerError creates a PostSourcePremiumizemeDatasetDatasetNameInternalServerError with default headers values
func NewPostSourcePremiumizemeDatasetDatasetNameInternalServerError() *PostSourcePremiumizemeDatasetDatasetNameInternalServerError {
	return &PostSourcePremiumizemeDatasetDatasetNameInternalServerError{}
}

/*
PostSourcePremiumizemeDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourcePremiumizemeDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source premiumizeme dataset dataset name internal server error response has a 2xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source premiumizeme dataset dataset name internal server error response has a 3xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source premiumizeme dataset dataset name internal server error response has a 4xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source premiumizeme dataset dataset name internal server error response has a 5xx status code
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source premiumizeme dataset dataset name internal server error response a status code equal to that given
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source premiumizeme dataset dataset name internal server error response
func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/premiumizeme/dataset/{datasetName}][%d] postSourcePremiumizemeDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourcePremiumizemeDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
