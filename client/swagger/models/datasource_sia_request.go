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

// DatasourceSiaRequest datasource sia request
//
// swagger:model datasource.SiaRequest
type DatasourceSiaRequest struct {

	// Sia Daemon API Password.
	APIPassword string `json:"apiPassword,omitempty"`

	// Sia daemon API URL, like http://sia.daemon.host:9980.
	APIURL *string `json:"apiUrl,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// Siad User Agent
	UserAgent *string `json:"userAgent,omitempty"`
}

// Validate validates this datasource sia request
func (m *DatasourceSiaRequest) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceSiaRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSiaRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSiaRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceSiaRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource sia request based on the context it is used
func (m *DatasourceSiaRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceSiaRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceSiaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceSiaRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceSiaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
