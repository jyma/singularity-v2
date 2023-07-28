// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ModelSourceType model source type
//
// swagger:model model.SourceType
type ModelSourceType string

func NewModelSourceType(value ModelSourceType) *ModelSourceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ModelSourceType.
func (m ModelSourceType) Pointer() *ModelSourceType {
	return &m
}

const (

	// ModelSourceTypeLocal captures enum value "local"
	ModelSourceTypeLocal ModelSourceType = "local"

	// ModelSourceTypeUpload captures enum value "upload"
	ModelSourceTypeUpload ModelSourceType = "upload"
)

// for schema
var modelSourceTypeEnum []interface{}

func init() {
	var res []ModelSourceType
	if err := json.Unmarshal([]byte(`["local","upload"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelSourceTypeEnum = append(modelSourceTypeEnum, v)
	}
}

func (m ModelSourceType) validateModelSourceTypeEnum(path, location string, value ModelSourceType) error {
	if err := validate.EnumCase(path, location, value, modelSourceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model source type
func (m ModelSourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelSourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model source type based on context it is used
func (m ModelSourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}