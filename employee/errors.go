package employee

import "errors"

var (
	ErrMandatoryParamsMissing = errors.New("missing mandatory params")
	ErrUserNotFound           = errors.New("employee not found")
	ErrNoEmployeesFound       = errors.New("no employee found in the page")
	ErrInvalidLimit           = errors.New("invalid limit")
	ErrInvalidPageNumber      = errors.New("invalid page number")
)
