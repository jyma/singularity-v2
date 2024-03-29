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

// ModelDealState model deal state
//
// swagger:model model.DealState
type ModelDealState string

func NewModelDealState(value ModelDealState) *ModelDealState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ModelDealState.
func (m ModelDealState) Pointer() *ModelDealState {
	return &m
}

const (

	// ModelDealStateProposed captures enum value "proposed"
	ModelDealStateProposed ModelDealState = "proposed"

	// ModelDealStatePublished captures enum value "published"
	ModelDealStatePublished ModelDealState = "published"

	// ModelDealStateActive captures enum value "active"
	ModelDealStateActive ModelDealState = "active"

	// ModelDealStateExpired captures enum value "expired"
	ModelDealStateExpired ModelDealState = "expired"

	// ModelDealStateProposalExpired captures enum value "proposal_expired"
	ModelDealStateProposalExpired ModelDealState = "proposal_expired"

	// ModelDealStateRejected captures enum value "rejected"
	ModelDealStateRejected ModelDealState = "rejected"

	// ModelDealStateSlashed captures enum value "slashed"
	ModelDealStateSlashed ModelDealState = "slashed"

	// ModelDealStateError captures enum value "error"
	ModelDealStateError ModelDealState = "error"
)

// for schema
var modelDealStateEnum []interface{}

func init() {
	var res []ModelDealState
	if err := json.Unmarshal([]byte(`["proposed","published","active","expired","proposal_expired","rejected","slashed","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelDealStateEnum = append(modelDealStateEnum, v)
	}
}

func (m ModelDealState) validateModelDealStateEnum(path, location string, value ModelDealState) error {
	if err := validate.EnumCase(path, location, value, modelDealStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model deal state
func (m ModelDealState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelDealStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model deal state based on context it is used
func (m ModelDealState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
