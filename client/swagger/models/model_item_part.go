// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelItemPart model item part
//
// swagger:model model.ItemPart
type ModelItemPart struct {

	// chunk Id
	ChunkID int64 `json:"chunkId,omitempty"`

	// cid
	Cid ModelCID `json:"cid,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// item
	Item *ModelItem `json:"item,omitempty"`

	// item Id
	ItemID int64 `json:"itemId,omitempty"`

	// length
	Length int64 `json:"length,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`
}

// Validate validates this model item part
func (m *ModelItemPart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelItemPart) validateItem(formats strfmt.Registry) error {
	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model item part based on the context it is used
func (m *ModelItemPart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelItemPart) contextValidateItem(ctx context.Context, formats strfmt.Registry) error {

	if m.Item != nil {

		if swag.IsZero(m.Item) { // not required
			return nil
		}

		if err := m.Item.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelItemPart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelItemPart) UnmarshalBinary(b []byte) error {
	var res ModelItemPart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
