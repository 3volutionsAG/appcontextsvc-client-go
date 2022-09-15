// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: remote_connection_svc.proto

package appcontextsvc_client

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _remote_connection_svc_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListRemoteConnectionsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRemoteConnectionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRemoteConnectionsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ListRemoteConnectionsResponseMultiError, or nil if none found.
func (m *ListRemoteConnectionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRemoteConnectionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRemoteConnections() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRemoteConnectionsResponseValidationError{
						field:  fmt.Sprintf("RemoteConnections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRemoteConnectionsResponseValidationError{
						field:  fmt.Sprintf("RemoteConnections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRemoteConnectionsResponseValidationError{
					field:  fmt.Sprintf("RemoteConnections[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PageNumber

	// no validation rules for PageSize

	// no validation rules for TotalCount

	if len(errors) > 0 {
		return ListRemoteConnectionsResponseMultiError(errors)
	}

	return nil
}

// ListRemoteConnectionsResponseMultiError is an error wrapping multiple
// validation errors returned by ListRemoteConnectionsResponse.ValidateAll()
// if the designated constraints aren't met.
type ListRemoteConnectionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRemoteConnectionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRemoteConnectionsResponseMultiError) AllErrors() []error { return m }

// ListRemoteConnectionsResponseValidationError is the validation error
// returned by ListRemoteConnectionsResponse.Validate if the designated
// constraints aren't met.
type ListRemoteConnectionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRemoteConnectionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRemoteConnectionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRemoteConnectionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRemoteConnectionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRemoteConnectionsResponseValidationError) ErrorName() string {
	return "ListRemoteConnectionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRemoteConnectionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRemoteConnectionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRemoteConnectionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRemoteConnectionsResponseValidationError{}

// Validate checks the field values on GetRemoteConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRemoteConnectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRemoteConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRemoteConnectionRequestMultiError, or nil if none found.
func (m *GetRemoteConnectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRemoteConnectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = GetRemoteConnectionRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GetRemoteConnectionRequestMultiError(errors)
	}

	return nil
}

func (m *GetRemoteConnectionRequest) _validateUuid(uuid string) error {
	if matched := _remote_connection_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetRemoteConnectionRequestMultiError is an error wrapping multiple
// validation errors returned by GetRemoteConnectionRequest.ValidateAll() if
// the designated constraints aren't met.
type GetRemoteConnectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRemoteConnectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRemoteConnectionRequestMultiError) AllErrors() []error { return m }

// GetRemoteConnectionRequestValidationError is the validation error returned
// by GetRemoteConnectionRequest.Validate if the designated constraints aren't met.
type GetRemoteConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRemoteConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRemoteConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRemoteConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRemoteConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRemoteConnectionRequestValidationError) ErrorName() string {
	return "GetRemoteConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetRemoteConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRemoteConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRemoteConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRemoteConnectionRequestValidationError{}

// Validate checks the field values on CreateRemoteConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRemoteConnectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRemoteConnectionRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateRemoteConnectionRequestMultiError, or nil if none found.
func (m *CreateRemoteConnectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRemoteConnectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRemoteConnection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRemoteConnectionRequestValidationError{
					field:  "RemoteConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRemoteConnectionRequestValidationError{
					field:  "RemoteConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRemoteConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRemoteConnectionRequestValidationError{
				field:  "RemoteConnection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRemoteConnectionRequestMultiError(errors)
	}

	return nil
}

// CreateRemoteConnectionRequestMultiError is an error wrapping multiple
// validation errors returned by CreateRemoteConnectionRequest.ValidateAll()
// if the designated constraints aren't met.
type CreateRemoteConnectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRemoteConnectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRemoteConnectionRequestMultiError) AllErrors() []error { return m }

// CreateRemoteConnectionRequestValidationError is the validation error
// returned by CreateRemoteConnectionRequest.Validate if the designated
// constraints aren't met.
type CreateRemoteConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRemoteConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRemoteConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRemoteConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRemoteConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRemoteConnectionRequestValidationError) ErrorName() string {
	return "CreateRemoteConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRemoteConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRemoteConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRemoteConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRemoteConnectionRequestValidationError{}

// Validate checks the field values on UpdateRemoteConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRemoteConnectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRemoteConnectionRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateRemoteConnectionRequestMultiError, or nil if none found.
func (m *UpdateRemoteConnectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRemoteConnectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRemoteConnection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRemoteConnectionRequestValidationError{
					field:  "RemoteConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRemoteConnectionRequestValidationError{
					field:  "RemoteConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRemoteConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRemoteConnectionRequestValidationError{
				field:  "RemoteConnection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRemoteConnectionRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRemoteConnectionRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRemoteConnectionRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRemoteConnectionRequestMultiError(errors)
	}

	return nil
}

// UpdateRemoteConnectionRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateRemoteConnectionRequest.ValidateAll()
// if the designated constraints aren't met.
type UpdateRemoteConnectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRemoteConnectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRemoteConnectionRequestMultiError) AllErrors() []error { return m }

// UpdateRemoteConnectionRequestValidationError is the validation error
// returned by UpdateRemoteConnectionRequest.Validate if the designated
// constraints aren't met.
type UpdateRemoteConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRemoteConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRemoteConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRemoteConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRemoteConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRemoteConnectionRequestValidationError) ErrorName() string {
	return "UpdateRemoteConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRemoteConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRemoteConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRemoteConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRemoteConnectionRequestValidationError{}

// Validate checks the field values on DeleteRemoteConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRemoteConnectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRemoteConnectionRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeleteRemoteConnectionRequestMultiError, or nil if none found.
func (m *DeleteRemoteConnectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRemoteConnectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = DeleteRemoteConnectionRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return DeleteRemoteConnectionRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteRemoteConnectionRequest) _validateUuid(uuid string) error {
	if matched := _remote_connection_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteRemoteConnectionRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteRemoteConnectionRequest.ValidateAll()
// if the designated constraints aren't met.
type DeleteRemoteConnectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRemoteConnectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRemoteConnectionRequestMultiError) AllErrors() []error { return m }

// DeleteRemoteConnectionRequestValidationError is the validation error
// returned by DeleteRemoteConnectionRequest.Validate if the designated
// constraints aren't met.
type DeleteRemoteConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRemoteConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRemoteConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRemoteConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRemoteConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRemoteConnectionRequestValidationError) ErrorName() string {
	return "DeleteRemoteConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRemoteConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRemoteConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRemoteConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRemoteConnectionRequestValidationError{}