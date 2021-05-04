package goerror

import (
	"net/http"
)

const (
	ErrCodeUnexpected       = "ERR_UNEXPECTED"
	ErrCodeNetBuild         = "ERR_NET_BUILD"
	ErrCodeNetConnect       = "ERR_NET_CONNECT"
	ErrCodeValidation       = "ERR_VALIDATION"
	ErrCodeFormatting       = "ERR_FORMATTING"
	ErrCodeDataRead         = "ERR_DATA_READ"
	ErrCodeDataWrite        = "ERR_DATA_WRITE"
	ErrCodeNoResult         = "ERR_NO_RESULT"
	ErrCodeUnauthorized     = "ERR_UNAUTHORIZED"
	ErrCodeExpired          = "ERR_EXPIRED"
	ErrCodeForbidden        = "ERR_FORBIDDEN"
	ErrCodeTooManyRequest   = "ERR_REQUEST_LIMIT"
	ErrCodeDataIncomplete   = "ERR_DATA_INCOMPLETE"
	ErrCodeEncryptData      = "ERR_ENCRYPT_DATA"
	ErrCodeInternalServer   = "ERR_INTERNAL_SERVER"
	ErrCodeNotFound         = "ERR_NOT_FOUND"
	ErrCodeMethodNotAllowed = "ERR_METHOD_NOT_ALLOWED"
)

var mapper = map[string]string{
	ErrCodeUnexpected:       "unexpected error occurred while processing request",
	ErrCodeNetBuild:         "failed to build connection to data source",
	ErrCodeNetConnect:       "failed to establish connection to data source",
	ErrCodeValidation:       "request contains invalid data",
	ErrCodeFormatting:       "an error occurred while formatting data",
	ErrCodeDataRead:         "failed to read data from data provider",
	ErrCodeDataWrite:        "failed to persist data into provider",
	ErrCodeNoResult:         "no result found match criteria",
	ErrCodeUnauthorized:     "unauthorized access",
	ErrCodeForbidden:        "forbidden access",
	ErrCodeExpired:          "expired pemission",
	ErrCodeTooManyRequest:   "request limit exceeded",
	ErrCodeDataIncomplete:   "stored data incomplete",
	ErrCodeEncryptData:      "failed to encrypting data",
	ErrCodeInternalServer:   "oops something went wrong",
	ErrCodeNotFound:         "404 not found",
	ErrCodeMethodNotAllowed: "405 method not allowed error",
}

var restmapper = map[string]int{
	ErrCodeUnexpected:       http.StatusInternalServerError,
	ErrCodeNetBuild:         http.StatusBadGateway,
	ErrCodeNetConnect:       http.StatusBadGateway,
	ErrCodeValidation:       http.StatusBadRequest,
	ErrCodeFormatting:       http.StatusBadRequest,
	ErrCodeDataRead:         http.StatusBadRequest,
	ErrCodeDataWrite:        http.StatusBadRequest,
	ErrCodeNoResult:         http.StatusNoContent,
	ErrCodeUnauthorized:     http.StatusUnauthorized,
	ErrCodeForbidden:        http.StatusForbidden,
	ErrCodeExpired:          http.StatusForbidden,
	ErrCodeTooManyRequest:   http.StatusTooManyRequests,
	ErrCodeDataIncomplete:   http.StatusPartialContent,
	ErrCodeEncryptData:      http.StatusBadRequest,
	ErrCodeInternalServer:   http.StatusInternalServerError,
	ErrCodeNotFound:         http.StatusNotFound,
	ErrCodeMethodNotAllowed: http.StatusMethodNotAllowed,
}

//Message retrieve error messages from given error code
func Message(code string) string {
	if s, ok := mapper[code]; ok {
		return s
	}
	return mapper[ErrCodeUnexpected]
}

//Messages retrieve all registered mapping error messages
func Messages() map[string]string {
	return mapper
}

//Message retrieve http code from given error code
func RestCode(code string) int {
	if s, ok := restmapper[code]; ok {
		return s
	}
	return restmapper[ErrCodeUnexpected]
}

//Messages retrieve all registered mapping error http code
func RestCodes() map[string]int {
	return restmapper
}
