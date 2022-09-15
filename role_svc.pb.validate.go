// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: role_svc.proto

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
var _role_svc_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListRolesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRolesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRolesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRolesResponseMultiError, or nil if none found.
func (m *ListRolesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRolesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRolesResponseValidationError{
						field:  fmt.Sprintf("Roles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRolesResponseValidationError{
						field:  fmt.Sprintf("Roles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRolesResponseValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
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
		return ListRolesResponseMultiError(errors)
	}

	return nil
}

// ListRolesResponseMultiError is an error wrapping multiple validation errors
// returned by ListRolesResponse.ValidateAll() if the designated constraints
// aren't met.
type ListRolesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRolesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRolesResponseMultiError) AllErrors() []error { return m }

// ListRolesResponseValidationError is the validation error returned by
// ListRolesResponse.Validate if the designated constraints aren't met.
type ListRolesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRolesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRolesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRolesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRolesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRolesResponseValidationError) ErrorName() string {
	return "ListRolesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRolesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRolesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRolesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRolesResponseValidationError{}

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRoleRequestMultiError,
// or nil if none found.
func (m *GetRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = GetRoleRequestValidationError{
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
		return GetRoleRequestMultiError(errors)
	}

	return nil
}

func (m *GetRoleRequest) _validateUuid(uuid string) error {
	if matched := _role_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetRoleRequestMultiError is an error wrapping multiple validation errors
// returned by GetRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleRequestMultiError) AllErrors() []error { return m }

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on CreateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleRequestMultiError, or nil if none found.
func (m *CreateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRole()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRoleRequestValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRoleRequestValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRoleRequestValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRoleRequestMultiError(errors)
	}

	return nil
}

// CreateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleRequestMultiError) AllErrors() []error { return m }

// CreateRoleRequestValidationError is the validation error returned by
// CreateRoleRequest.Validate if the designated constraints aren't met.
type CreateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleRequestValidationError) ErrorName() string {
	return "CreateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleRequestValidationError{}

// Validate checks the field values on UpdateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleRequestMultiError, or nil if none found.
func (m *UpdateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRole()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRoleRequestValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRoleRequestValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRoleRequestValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRoleRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRoleRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRoleRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRoleRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleRequestMultiError) AllErrors() []error { return m }

// UpdateRoleRequestValidationError is the validation error returned by
// UpdateRoleRequest.Validate if the designated constraints aren't met.
type UpdateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleRequestValidationError) ErrorName() string {
	return "UpdateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleRequestValidationError{}

// Validate checks the field values on DeleteRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleRequestMultiError, or nil if none found.
func (m *DeleteRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = DeleteRoleRequestValidationError{
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
		return DeleteRoleRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteRoleRequest) _validateUuid(uuid string) error {
	if matched := _role_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteRoleRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleRequestMultiError) AllErrors() []error { return m }

// DeleteRoleRequestValidationError is the validation error returned by
// DeleteRoleRequest.Validate if the designated constraints aren't met.
type DeleteRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleRequestValidationError) ErrorName() string {
	return "DeleteRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleRequestValidationError{}