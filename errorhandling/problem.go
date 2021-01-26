package main

import (
	"errors"
	"fmt"
)

type (
	// Problem ...
	Problem struct {
		code    uint32
		status  int
		message string
	}
)

// NewProblem return a new Problem.
func NewProblem(code uint32, status int, message string) Problem {
	return Problem{
		code:    code,
		status:  status,
		message: message,
	}
}

func (p Problem) Error() string {
	return p.message
}

// MarshalJSON implements json.Marshaler.
func (p Problem) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("{\"code\": %d, \"message\": \"%s\"}", p.code, p.message)
	return []byte(json), nil
}

// Code ...
func (p Problem) Code() uint32 {
	return p.code
}

// Status ...
func (p Problem) Status() int {
	return p.status
}

// Message ...
func (p Problem) Message() string {
	return p.message
}

// Is implement errors.Is method
func (p Problem) Is(err error) bool {
	var Problem Problem
	if !errors.As(err, &Problem) {
		return false
	}
	return Problem.code == p.code
}

func (p Problem) String() string {
	return fmt.Sprintf("code: %d, status: %d, message: %s", p.code, p.status, p.message)
}
