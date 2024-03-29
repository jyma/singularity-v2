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

// NewPostSourceOosDatasetDatasetNameParams creates a new PostSourceOosDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceOosDatasetDatasetNameParams() *PostSourceOosDatasetDatasetNameParams {
	return &PostSourceOosDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceOosDatasetDatasetNameParamsWithTimeout creates a new PostSourceOosDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceOosDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceOosDatasetDatasetNameParams {
	return &PostSourceOosDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceOosDatasetDatasetNameParamsWithContext creates a new PostSourceOosDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceOosDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceOosDatasetDatasetNameParams {
	return &PostSourceOosDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceOosDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceOosDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceOosDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceOosDatasetDatasetNameParams {
	return &PostSourceOosDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceOosDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source oos dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceOosDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceOosRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source oos dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceOosDatasetDatasetNameParams) WithDefaults() *PostSourceOosDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source oos dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceOosDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceOosDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceOosDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceOosDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceOosDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) WithRequest(request *models.DatasourceOosRequest) *PostSourceOosDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source oos dataset dataset name params
func (o *PostSourceOosDatasetDatasetNameParams) SetRequest(request *models.DatasourceOosRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceOosDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
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
