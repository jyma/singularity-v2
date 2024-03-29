// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPostSourceIDChunkParams creates a new PostSourceIDChunkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceIDChunkParams() *PostSourceIDChunkParams {
	return &PostSourceIDChunkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceIDChunkParamsWithTimeout creates a new PostSourceIDChunkParams object
// with the ability to set a timeout on a request.
func NewPostSourceIDChunkParamsWithTimeout(timeout time.Duration) *PostSourceIDChunkParams {
	return &PostSourceIDChunkParams{
		timeout: timeout,
	}
}

// NewPostSourceIDChunkParamsWithContext creates a new PostSourceIDChunkParams object
// with the ability to set a context for a request.
func NewPostSourceIDChunkParamsWithContext(ctx context.Context) *PostSourceIDChunkParams {
	return &PostSourceIDChunkParams{
		Context: ctx,
	}
}

// NewPostSourceIDChunkParamsWithHTTPClient creates a new PostSourceIDChunkParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceIDChunkParamsWithHTTPClient(client *http.Client) *PostSourceIDChunkParams {
	return &PostSourceIDChunkParams{
		HTTPClient: client,
	}
}

/*
PostSourceIDChunkParams contains all the parameters to send to the API endpoint

	for the post source ID chunk operation.

	Typically these are written to a http.Request.
*/
type PostSourceIDChunkParams struct {

	/* ID.

	   Source ID
	*/
	ID string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceChunkRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source ID chunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDChunkParams) WithDefaults() *PostSourceIDChunkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source ID chunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDChunkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source ID chunk params
func (o *PostSourceIDChunkParams) WithTimeout(timeout time.Duration) *PostSourceIDChunkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source ID chunk params
func (o *PostSourceIDChunkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source ID chunk params
func (o *PostSourceIDChunkParams) WithContext(ctx context.Context) *PostSourceIDChunkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source ID chunk params
func (o *PostSourceIDChunkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source ID chunk params
func (o *PostSourceIDChunkParams) WithHTTPClient(client *http.Client) *PostSourceIDChunkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source ID chunk params
func (o *PostSourceIDChunkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post source ID chunk params
func (o *PostSourceIDChunkParams) WithID(id string) *PostSourceIDChunkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post source ID chunk params
func (o *PostSourceIDChunkParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the post source ID chunk params
func (o *PostSourceIDChunkParams) WithRequest(request *models.DatasourceChunkRequest) *PostSourceIDChunkParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source ID chunk params
func (o *PostSourceIDChunkParams) SetRequest(request *models.DatasourceChunkRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceIDChunkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
