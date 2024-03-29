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

// PostSourceIDPushReader is a Reader for the PostSourceIDPush structure.
type PostSourceIDPushReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceIDPushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSourceIDPushCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceIDPushBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostSourceIDPushConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceIDPushInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/{id}/push] PostSourceIDPush", response, response.Code())
	}
}

// NewPostSourceIDPushCreated creates a PostSourceIDPushCreated with default headers values
func NewPostSourceIDPushCreated() *PostSourceIDPushCreated {
	return &PostSourceIDPushCreated{}
}

/*
PostSourceIDPushCreated describes a response with status code 201, with default header values.

Created
*/
type PostSourceIDPushCreated struct {
	Payload *models.ModelItem
}

// IsSuccess returns true when this post source Id push created response has a 2xx status code
func (o *PostSourceIDPushCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source Id push created response has a 3xx status code
func (o *PostSourceIDPushCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id push created response has a 4xx status code
func (o *PostSourceIDPushCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Id push created response has a 5xx status code
func (o *PostSourceIDPushCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Id push created response a status code equal to that given
func (o *PostSourceIDPushCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post source Id push created response
func (o *PostSourceIDPushCreated) Code() int {
	return 201
}

func (o *PostSourceIDPushCreated) Error() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushCreated  %+v", 201, o.Payload)
}

func (o *PostSourceIDPushCreated) String() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushCreated  %+v", 201, o.Payload)
}

func (o *PostSourceIDPushCreated) GetPayload() *models.ModelItem {
	return o.Payload
}

func (o *PostSourceIDPushCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceIDPushBadRequest creates a PostSourceIDPushBadRequest with default headers values
func NewPostSourceIDPushBadRequest() *PostSourceIDPushBadRequest {
	return &PostSourceIDPushBadRequest{}
}

/*
PostSourceIDPushBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceIDPushBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post source Id push bad request response has a 2xx status code
func (o *PostSourceIDPushBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Id push bad request response has a 3xx status code
func (o *PostSourceIDPushBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id push bad request response has a 4xx status code
func (o *PostSourceIDPushBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source Id push bad request response has a 5xx status code
func (o *PostSourceIDPushBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Id push bad request response a status code equal to that given
func (o *PostSourceIDPushBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source Id push bad request response
func (o *PostSourceIDPushBadRequest) Code() int {
	return 400
}

func (o *PostSourceIDPushBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceIDPushBadRequest) String() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceIDPushBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostSourceIDPushBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceIDPushConflict creates a PostSourceIDPushConflict with default headers values
func NewPostSourceIDPushConflict() *PostSourceIDPushConflict {
	return &PostSourceIDPushConflict{}
}

/*
PostSourceIDPushConflict describes a response with status code 409, with default header values.

Item already exists
*/
type PostSourceIDPushConflict struct {
	Payload string
}

// IsSuccess returns true when this post source Id push conflict response has a 2xx status code
func (o *PostSourceIDPushConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Id push conflict response has a 3xx status code
func (o *PostSourceIDPushConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id push conflict response has a 4xx status code
func (o *PostSourceIDPushConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source Id push conflict response has a 5xx status code
func (o *PostSourceIDPushConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Id push conflict response a status code equal to that given
func (o *PostSourceIDPushConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post source Id push conflict response
func (o *PostSourceIDPushConflict) Code() int {
	return 409
}

func (o *PostSourceIDPushConflict) Error() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushConflict  %+v", 409, o.Payload)
}

func (o *PostSourceIDPushConflict) String() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushConflict  %+v", 409, o.Payload)
}

func (o *PostSourceIDPushConflict) GetPayload() string {
	return o.Payload
}

func (o *PostSourceIDPushConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceIDPushInternalServerError creates a PostSourceIDPushInternalServerError with default headers values
func NewPostSourceIDPushInternalServerError() *PostSourceIDPushInternalServerError {
	return &PostSourceIDPushInternalServerError{}
}

/*
PostSourceIDPushInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceIDPushInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post source Id push internal server error response has a 2xx status code
func (o *PostSourceIDPushInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Id push internal server error response has a 3xx status code
func (o *PostSourceIDPushInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id push internal server error response has a 4xx status code
func (o *PostSourceIDPushInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Id push internal server error response has a 5xx status code
func (o *PostSourceIDPushInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source Id push internal server error response a status code equal to that given
func (o *PostSourceIDPushInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source Id push internal server error response
func (o *PostSourceIDPushInternalServerError) Code() int {
	return 500
}

func (o *PostSourceIDPushInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceIDPushInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/{id}/push][%d] postSourceIdPushInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceIDPushInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostSourceIDPushInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
