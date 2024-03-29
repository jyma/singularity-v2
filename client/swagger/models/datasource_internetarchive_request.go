// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourceInternetarchiveRequest datasource internetarchive request
//
// swagger:model datasource.InternetarchiveRequest
type DatasourceInternetarchiveRequest struct {

	// IAS3 Access Key.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Don't ask the server to test against MD5 checksum calculated by rclone.
	DisableChecksum *string `json:"disableChecksum,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// IAS3 Endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// Host of InternetArchive Frontend.
	FrontEndpoint *string `json:"frontEndpoint,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// IAS3 Secret Key (password).
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	WaitArchive *string `json:"waitArchive,omitempty"`
}

// Validate validates this datasource internetarchive request
func (m *DatasourceInternetarchiveRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanningState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceInternetarchiveRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceInternetarchiveRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceInternetarchiveRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceInternetarchiveRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource internetarchive request based on the context it is used
func (m *DatasourceInternetarchiveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceInternetarchiveRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceInternetarchiveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceInternetarchiveRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceInternetarchiveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
