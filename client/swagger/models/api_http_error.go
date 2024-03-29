// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIHTTPError api HTTP error
//
// swagger:model api.HTTPError
type APIHTTPError struct {

	// err
	Err string `json:"err,omitempty"`
}

// Validate validates this api HTTP error
func (m *APIHTTPError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api HTTP error based on context it is used
func (m *APIHTTPError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIHTTPError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIHTTPError) UnmarshalBinary(b []byte) error {
	var res APIHTTPError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
