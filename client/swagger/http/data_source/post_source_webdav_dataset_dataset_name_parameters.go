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

// NewPostSourceWebdavDatasetDatasetNameParams creates a new PostSourceWebdavDatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceWebdavDatasetDatasetNameParams() *PostSourceWebdavDatasetDatasetNameParams {
	return &PostSourceWebdavDatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceWebdavDatasetDatasetNameParamsWithTimeout creates a new PostSourceWebdavDatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceWebdavDatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceWebdavDatasetDatasetNameParams {
	return &PostSourceWebdavDatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceWebdavDatasetDatasetNameParamsWithContext creates a new PostSourceWebdavDatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceWebdavDatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceWebdavDatasetDatasetNameParams {
	return &PostSourceWebdavDatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceWebdavDatasetDatasetNameParamsWithHTTPClient creates a new PostSourceWebdavDatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceWebdavDatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceWebdavDatasetDatasetNameParams {
	return &PostSourceWebdavDatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceWebdavDatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source webdav dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceWebdavDatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceWebdavRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source webdav dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceWebdavDatasetDatasetNameParams) WithDefaults() *PostSourceWebdavDatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source webdav dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceWebdavDatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceWebdavDatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceWebdavDatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceWebdavDatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceWebdavDatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) WithRequest(request *models.DatasourceWebdavRequest) *PostSourceWebdavDatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source webdav dataset dataset name params
func (o *PostSourceWebdavDatasetDatasetNameParams) SetRequest(request *models.DatasourceWebdavRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceWebdavDatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
