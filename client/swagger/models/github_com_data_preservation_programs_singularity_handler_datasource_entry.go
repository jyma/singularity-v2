// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry github com data preservation programs singularity handler datasource entry
//
// swagger:model github_com_data-preservation-programs_singularity_handler_datasource.Entry
type GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry struct {

	// is dir
	IsDir bool `json:"isDir,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this github com data preservation programs singularity handler datasource entry
func (m *GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com data preservation programs singularity handler datasource entry based on context it is used
func (m *GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry) UnmarshalBinary(b []byte) error {
	var res GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
