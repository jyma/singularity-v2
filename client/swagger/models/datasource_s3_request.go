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

// DatasourceS3Request datasource s3 request
//
// swagger:model datasource.S3Request
type DatasourceS3Request struct {

	// AWS Access Key ID.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Canned ACL used when creating buckets and storing or copying objects.
	ACL string `json:"acl,omitempty"`

	// Canned ACL used when creating buckets.
	BucketACL string `json:"bucketAcl,omitempty"`

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Cutoff for switching to multipart copy.
	CopyCutoff *string `json:"copyCutoff,omitempty"`

	// If set this will decompress gzip encoded objects.
	Decompress *string `json:"decompress,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// Don't store MD5 checksum with object metadata.
	DisableChecksum *string `json:"disableChecksum,omitempty"`

	// Disable usage of http2 for S3 backends.
	DisableHttp2 *string `json:"disableHttp2,omitempty"`

	// Custom endpoint for downloads.
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for S3 API.
	Endpoint string `json:"endpoint,omitempty"`

	// Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	EnvAuth *string `json:"envAuth,omitempty"`

	// If true use path style access if false use virtual hosted style.
	ForcePathStyle *string `json:"forcePathStyle,omitempty"`

	// If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	LeavePartsOnError *string `json:"leavePartsOnError,omitempty"`

	// Size of listing chunk (response list for each ListObject S3 request).
	ListChunk *string `json:"listChunk,omitempty"`

	// Whether to url encode listings: true/false/unset
	ListURLEncode *string `json:"listUrlEncode,omitempty"`

	// Version of ListObjects to use: 1,2 or 0 for auto.
	ListVersion *string `json:"listVersion,omitempty"`

	// Location constraint - must be set to match the Region.
	LocationConstraint string `json:"locationConstraint,omitempty"`

	// Maximum number of parts in a multipart upload.
	MaxUploadParts *string `json:"maxUploadParts,omitempty"`

	// How often internal memory buffer pools will be flushed.
	MemoryPoolFlushTime *string `json:"memoryPoolFlushTime,omitempty"`

	// Whether to use mmap buffers in internal memory pool.
	MemoryPoolUseMmap *string `json:"memoryPoolUseMmap,omitempty"`

	// Set this if the backend might gzip objects.
	MightGzip *string `json:"mightGzip,omitempty"`

	// If set, don't attempt to check the bucket exists or create it.
	NoCheckBucket *string `json:"noCheckBucket,omitempty"`

	// If set, don't HEAD uploaded objects to check integrity.
	NoHead *string `json:"noHead,omitempty"`

	// If set, do not do HEAD before GET when getting objects.
	NoHeadObject *string `json:"noHeadObject,omitempty"`

	// Suppress setting and reading of system metadata
	NoSystemMetadata *string `json:"noSystemMetadata,omitempty"`

	// Profile to use in the shared credentials file.
	Profile string `json:"profile,omitempty"`

	// Choose your S3 provider.
	Provider string `json:"provider,omitempty"`

	// Region to connect to.
	Region string `json:"region,omitempty"`

	// Enables requester pays option when interacting with S3 bucket.
	RequesterPays *string `json:"requesterPays,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// AWS Secret Access Key (password).
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// The server-side encryption algorithm used when storing this object in S3.
	ServerSideEncryption string `json:"serverSideEncryption,omitempty"`

	// An AWS session token.
	SessionToken string `json:"sessionToken,omitempty"`

	// Path to the shared credentials file.
	SharedCredentialsFile string `json:"sharedCredentialsFile,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm,omitempty"`

	// To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKey string `json:"sseCustomerKey,omitempty"`

	// If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyBase64 string `json:"sseCustomerKeyBase64,omitempty"`

	// If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	SseCustomerKeyMd5 string `json:"sseCustomerKeyMd5,omitempty"`

	// If using KMS ID you must provide the ARN of Key.
	SseKmsKeyID string `json:"sseKmsKeyId,omitempty"`

	// The storage class to use when storing new objects in S3.
	StorageClass string `json:"storageClass,omitempty"`

	// Endpoint for STS.
	StsEndpoint string `json:"stsEndpoint,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *string `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// If true use the AWS S3 accelerated endpoint.
	UseAccelerateEndpoint *string `json:"useAccelerateEndpoint,omitempty"`

	// Whether to use ETag in multipart uploads for verification
	UseMultipartEtag *string `json:"useMultipartEtag,omitempty"`

	// Whether to use a presigned request or PutObject for single part uploads
	UsePresignedRequest *string `json:"usePresignedRequest,omitempty"`

	// If true use v2 authentication.
	V2Auth *string `json:"v2Auth,omitempty"`

	// Show file versions as they were at the specified time.
	VersionAt *string `json:"versionAt,omitempty"`

	// Include old versions in directory listings.
	Versions *string `json:"versions,omitempty"`
}

// Validate validates this datasource s3 request
func (m *DatasourceS3Request) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceS3Request) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceS3Request) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceS3Request) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceS3Request) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource s3 request based on the context it is used
func (m *DatasourceS3Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceS3Request) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceS3Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceS3Request) UnmarshalBinary(b []byte) error {
	var res DatasourceS3Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
