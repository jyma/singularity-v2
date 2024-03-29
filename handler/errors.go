package handler

import (
	"fmt"

	"github.com/pkg/errors"
)

type InvalidParameterError struct {
	Err error
}

func (e InvalidParameterError) Unwrap() error {
	return e.Err
}
func (e InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter: %s", e.Err.Error())
}

func NewInvalidParameterErr(err string) InvalidParameterError {
	return InvalidParameterError{
		Err: errors.New(err),
	}
}

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Unwrap() error {
	return e.Err
}
func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.Err.Error())
}

type DuplicateRecordError struct {
	Err error
}

func (e DuplicateRecordError) Unwrap() error {
	return e.Err
}

func (e DuplicateRecordError) Error() string {
	return fmt.Sprintf("duplicate record: %s", e.Err.Error())
}

func NewDuplicateRecordError(err string) DuplicateRecordError {
	return DuplicateRecordError{
		Err: errors.New(err),
	}
}
