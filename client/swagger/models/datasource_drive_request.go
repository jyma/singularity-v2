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

// DatasourceDriveRequest datasource drive request
//
// swagger:model datasource.DriveRequest
type DatasourceDriveRequest struct {

	// Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	AcknowledgeAbuse *string `json:"acknowledgeAbuse,omitempty"`

	// Allow the filetype to change when uploading Google docs.
	AllowImportNameChange *string `json:"allowImportNameChange,omitempty"`

	// Deprecated: No longer needed.
	AlternateExport *string `json:"alternateExport,omitempty"`

	// Only consider files owned by the authenticated user.
	AuthOwnerOnly *string `json:"authOwnerOnly,omitempty"`

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Upload chunk size.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Google Application Client Id
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// Server side copy contents of shortcuts instead of the shortcut.
	CopyShortcutContent *string `json:"copyShortcutContent,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Disable drive using http2.
	DisableHttp2 *string `json:"disableHttp2,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Comma separated list of preferred formats for downloading Google docs.
	ExportFormats *string `json:"exportFormats,omitempty"`

	// Deprecated: See export_formats.
	Formats string `json:"formats,omitempty"`

	// Impersonate this user when using a service account.
	Impersonate string `json:"impersonate,omitempty"`

	// Comma separated list of preferred formats for uploading Google docs.
	ImportFormats string `json:"importFormats,omitempty"`

	// Keep new head revision of each file forever.
	KeepRevisionForever *string `json:"keepRevisionForever,omitempty"`

	// Size of listing chunk 100-1000, 0 to disable.
	ListChunk *string `json:"listChunk,omitempty"`

	// Number of API calls to allow without sleeping.
	PacerBurst *string `json:"pacerBurst,omitempty"`

	// Minimum time to sleep between API calls.
	PacerMinSleep *string `json:"pacerMinSleep,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Resource key for accessing a link-shared file.
	ResourceKey string `json:"resourceKey,omitempty"`

	// ID of the root folder.
	RootFolderID string `json:"rootFolderId,omitempty"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// Scope that rclone should use when requesting access from drive.
	Scope string `json:"scope,omitempty"`

	// Allow server-side operations (e.g. copy) to work across different drive configs.
	ServerSideAcrossConfigs *string `json:"serverSideAcrossConfigs,omitempty"`

	// Service Account Credentials JSON blob.
	ServiceAccountCredentials string `json:"serviceAccountCredentials,omitempty"`

	// Service Account Credentials JSON file path.
	ServiceAccountFile string `json:"serviceAccountFile,omitempty"`

	// Only show files that are shared with me.
	SharedWithMe *string `json:"sharedWithMe,omitempty"`

	// Show sizes as storage quota usage, not actual size.
	SizeAsQuota *string `json:"sizeAsQuota,omitempty"`

	// Skip MD5 checksum on Google photos and videos only.
	SkipChecksumGphotos *string `json:"skipChecksumGphotos,omitempty"`

	// If set skip dangling shortcut files.
	SkipDanglingShortcuts *string `json:"skipDanglingShortcuts,omitempty"`

	// Skip google documents in all listings.
	SkipGdocs *string `json:"skipGdocs,omitempty"`

	// If set skip shortcut files.
	SkipShortcuts *string `json:"skipShortcuts,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// Only show files that are starred.
	StarredOnly *string `json:"starredOnly,omitempty"`

	// Make download limit errors be fatal.
	StopOnDownloadLimit *string `json:"stopOnDownloadLimit,omitempty"`

	// Make upload limit errors be fatal.
	StopOnUploadLimit *string `json:"stopOnUploadLimit,omitempty"`

	// ID of the Shared Drive (Team Drive).
	TeamDrive string `json:"teamDrive,omitempty"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`

	// Only show files that are in the trash.
	TrashedOnly *string `json:"trashedOnly,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// Use file created date instead of modified date.
	UseCreatedDate *string `json:"useCreatedDate,omitempty"`

	// Use date file was shared instead of modified date.
	UseSharedDate *string `json:"useSharedDate,omitempty"`

	// Send files to the trash instead of deleting permanently.
	UseTrash *string `json:"useTrash,omitempty"`

	// If Object's are greater, use drive v2 API to download.
	V2DownloadMinSize *string `json:"v2DownloadMinSize,omitempty"`
}

// Validate validates this datasource drive request
func (m *DatasourceDriveRequest) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceDriveRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceDriveRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceDriveRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceDriveRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource drive request based on the context it is used
func (m *DatasourceDriveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceDriveRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceDriveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceDriveRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceDriveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
