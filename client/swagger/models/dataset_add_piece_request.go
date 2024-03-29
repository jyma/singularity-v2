// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasetAddPieceRequest dataset add piece request
//
// swagger:model dataset.AddPieceRequest
type DatasetAddPieceRequest struct {

	// Path to the CAR file, used to determine the size of the file and root CID
	FilePath string `json:"filePath,omitempty"`

	// CID of the piece
	PieceCid string `json:"pieceCid,omitempty"`

	// Size of the piece
	PieceSize string `json:"pieceSize,omitempty"`

	// Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal
	RootCid string `json:"rootCid,omitempty"`
}

// Validate validates this dataset add piece request
func (m *DatasetAddPieceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataset add piece request based on context it is used
func (m *DatasetAddPieceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasetAddPieceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetAddPieceRequest) UnmarshalBinary(b []byte) error {
	var res DatasetAddPieceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
