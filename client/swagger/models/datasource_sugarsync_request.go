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

// DatasourceSugarsyncRequest datasource sugarsync request
//
// swagger:model datasource.SugarsyncRequest
type DatasourceSugarsyncRequest struct {

	// Sugarsync Access Key ID.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Sugarsync App ID.
	AppID string `json:"appId,omitempty"`

	// Sugarsync authorization.
	Authorization string `json:"authorization,omitempty"`

	// Sugarsync authorization expiry.
	AuthorizationExpiry string `json:"authorizationExpiry,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Sugarsync deleted folder id.
	DeletedID string `json:"deletedId,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Permanently delete files if true
	HardDelete *string `json:"hardDelete,omitempty"`

	// Sugarsync Private Access Key.
	PrivateAccessKey string `json:"privateAccessKey,omitempty"`

	// Sugarsync refresh token.
	RefreshToken string `json:"refreshToken,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Sugarsync root id.
	RootID string `json:"rootId,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// Sugarsync user.
	User string `json:"user,omitempty"`
}

// Validate validates this datasource sugarsync request
func (m *DatasourceSugarsyncRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
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

func (m *DatasourceSugarsyncRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSugarsyncRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceSugarsyncRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasource sugarsync request based on context it is used
func (m *DatasourceSugarsyncRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceSugarsyncRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceSugarsyncRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceSugarsyncRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}