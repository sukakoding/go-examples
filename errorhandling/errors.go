package main

import "net/http"

const (
	codeUnknown uint32 = iota + 1000
	codeNotFound
	codeTimeOut
	codeBadRequest
	codeInternalServer
	codeUnauthorized
	codeInvalidUserPassword
	codeDuplicatedEmail
	codeFileSizeExceedLimit
	codeNotSupported
	codeInvalidCurriculum
	codeUserHasActivated
	codeRoleNotFound
)

var (
	// ErrInternalServer will throw if any the Internal Server Error happen
	ErrInternalServer = NewProblem(codeNotFound, http.StatusInternalServerError, "Something happened wrong during generating response")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = NewProblem(codeNotFound, http.StatusNotFound, "Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = NewProblem(codeNotFound, http.StatusConflict, "Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = NewProblem(codeBadRequest, http.StatusBadRequest, "Given Param is not valid")
	// ErrAuthenticationFailed is returned when request has not authenticated user
	ErrAuthenticationFailed = NewProblem(codeNotFound, http.StatusNotFound, "Authentication failed")
	// ErrAccessDenied is returned when user has not permission
	ErrAccessDenied = NewProblem(codeNotFound, http.StatusNotFound, "You need permission to perform this action")
	// ErrUserNotFound is returned when user not found
	ErrUserNotFound = NewProblem(codeNotFound, http.StatusNotFound, "user not found")
	// ErrUserPasswordResetTokenNotFound is returned when user password reset token not found
	ErrUserPasswordResetTokenNotFound = NewProblem(codeNotFound, http.StatusNotFound, "user password reset token not found")
	// ErrRoleNotFound is returned when role not found
	ErrRoleNotFound = NewProblem(codeRoleNotFound, http.StatusNotFound, "role not found")
	// ErrUserHasActivated ...
	ErrUserHasActivated = NewProblem(codeUserHasActivated, http.StatusMethodNotAllowed, "User has activated")
	// ErrInvalidCurriculum ...
	ErrInvalidCurriculum = NewProblem(codeInvalidCurriculum, http.StatusBadRequest, "Invalid Curriculum")
)

// NewErrBadRequest ...
func NewErrBadRequest(message string) error {
	return NewProblem(codeBadRequest, http.StatusBadRequest, message)
}

// NewErrNotFound ...
func NewErrNotFound(message string) error {
	return NewProblem(codeNotFound, http.StatusNotFound, message)
}
