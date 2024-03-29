// Code generated by go-swagger; DO NOT EDIT.

package deal_schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostScheduleIDPauseReader is a Reader for the PostScheduleIDPause structure.
type PostScheduleIDPauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScheduleIDPauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScheduleIDPauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostScheduleIDPauseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostScheduleIDPauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /schedule/{id}/pause] PostScheduleIDPause", response, response.Code())
	}
}

// NewPostScheduleIDPauseOK creates a PostScheduleIDPauseOK with default headers values
func NewPostScheduleIDPauseOK() *PostScheduleIDPauseOK {
	return &PostScheduleIDPauseOK{}
}

/*
PostScheduleIDPauseOK describes a response with status code 200, with default header values.

OK
*/
type PostScheduleIDPauseOK struct {
	Payload *models.ModelSchedule
}

// IsSuccess returns true when this post schedule Id pause o k response has a 2xx status code
func (o *PostScheduleIDPauseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post schedule Id pause o k response has a 3xx status code
func (o *PostScheduleIDPauseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post schedule Id pause o k response has a 4xx status code
func (o *PostScheduleIDPauseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post schedule Id pause o k response has a 5xx status code
func (o *PostScheduleIDPauseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post schedule Id pause o k response a status code equal to that given
func (o *PostScheduleIDPauseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post schedule Id pause o k response
func (o *PostScheduleIDPauseOK) Code() int {
	return 200
}

func (o *PostScheduleIDPauseOK) Error() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseOK  %+v", 200, o.Payload)
}

func (o *PostScheduleIDPauseOK) String() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseOK  %+v", 200, o.Payload)
}

func (o *PostScheduleIDPauseOK) GetPayload() *models.ModelSchedule {
	return o.Payload
}

func (o *PostScheduleIDPauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScheduleIDPauseBadRequest creates a PostScheduleIDPauseBadRequest with default headers values
func NewPostScheduleIDPauseBadRequest() *PostScheduleIDPauseBadRequest {
	return &PostScheduleIDPauseBadRequest{}
}

/*
PostScheduleIDPauseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostScheduleIDPauseBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post schedule Id pause bad request response has a 2xx status code
func (o *PostScheduleIDPauseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post schedule Id pause bad request response has a 3xx status code
func (o *PostScheduleIDPauseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post schedule Id pause bad request response has a 4xx status code
func (o *PostScheduleIDPauseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post schedule Id pause bad request response has a 5xx status code
func (o *PostScheduleIDPauseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post schedule Id pause bad request response a status code equal to that given
func (o *PostScheduleIDPauseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post schedule Id pause bad request response
func (o *PostScheduleIDPauseBadRequest) Code() int {
	return 400
}

func (o *PostScheduleIDPauseBadRequest) Error() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseBadRequest  %+v", 400, o.Payload)
}

func (o *PostScheduleIDPauseBadRequest) String() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseBadRequest  %+v", 400, o.Payload)
}

func (o *PostScheduleIDPauseBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostScheduleIDPauseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScheduleIDPauseInternalServerError creates a PostScheduleIDPauseInternalServerError with default headers values
func NewPostScheduleIDPauseInternalServerError() *PostScheduleIDPauseInternalServerError {
	return &PostScheduleIDPauseInternalServerError{}
}

/*
PostScheduleIDPauseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostScheduleIDPauseInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post schedule Id pause internal server error response has a 2xx status code
func (o *PostScheduleIDPauseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post schedule Id pause internal server error response has a 3xx status code
func (o *PostScheduleIDPauseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post schedule Id pause internal server error response has a 4xx status code
func (o *PostScheduleIDPauseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post schedule Id pause internal server error response has a 5xx status code
func (o *PostScheduleIDPauseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post schedule Id pause internal server error response a status code equal to that given
func (o *PostScheduleIDPauseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post schedule Id pause internal server error response
func (o *PostScheduleIDPauseInternalServerError) Code() int {
	return 500
}

func (o *PostScheduleIDPauseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScheduleIDPauseInternalServerError) String() string {
	return fmt.Sprintf("[POST /schedule/{id}/pause][%d] postScheduleIdPauseInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScheduleIDPauseInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostScheduleIDPauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
