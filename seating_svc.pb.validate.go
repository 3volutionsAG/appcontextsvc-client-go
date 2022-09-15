// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: seating_svc.proto

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
var _seating_svc_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListSeatingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSeatingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSeatingsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSeatingsResponseMultiError, or nil if none found.
func (m *ListSeatingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSeatingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSeatings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSeatingsResponseValidationError{
						field:  fmt.Sprintf("Seatings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSeatingsResponseValidationError{
						field:  fmt.Sprintf("Seatings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSeatingsResponseValidationError{
					field:  fmt.Sprintf("Seatings[%v]", idx),
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
		return ListSeatingsResponseMultiError(errors)
	}

	return nil
}

// ListSeatingsResponseMultiError is an error wrapping multiple validation
// errors returned by ListSeatingsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListSeatingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSeatingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSeatingsResponseMultiError) AllErrors() []error { return m }

// ListSeatingsResponseValidationError is the validation error returned by
// ListSeatingsResponse.Validate if the designated constraints aren't met.
type ListSeatingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSeatingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSeatingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSeatingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSeatingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSeatingsResponseValidationError) ErrorName() string {
	return "ListSeatingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSeatingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSeatingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSeatingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSeatingsResponseValidationError{}

// Validate checks the field values on GetSeatingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSeatingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSeatingRequestMultiError, or nil if none found.
func (m *GetSeatingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSeatingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetSeatingRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSeatingRequestMultiError(errors)
	}

	return nil
}

func (m *GetSeatingRequest) _validateUuid(uuid string) error {
	if matched := _seating_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetSeatingRequestMultiError is an error wrapping multiple validation errors
// returned by GetSeatingRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSeatingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSeatingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSeatingRequestMultiError) AllErrors() []error { return m }

// GetSeatingRequestValidationError is the validation error returned by
// GetSeatingRequest.Validate if the designated constraints aren't met.
type GetSeatingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSeatingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSeatingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSeatingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSeatingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSeatingRequestValidationError) ErrorName() string {
	return "GetSeatingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSeatingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSeatingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSeatingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSeatingRequestValidationError{}

// Validate checks the field values on CreateSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSeatingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSeatingRequestMultiError, or nil if none found.
func (m *CreateSeatingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSeatingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSeating()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSeatingRequestValidationError{
					field:  "Seating",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSeatingRequestValidationError{
					field:  "Seating",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSeating()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSeatingRequestValidationError{
				field:  "Seating",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateSeatingRequestMultiError(errors)
	}

	return nil
}

// CreateSeatingRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSeatingRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSeatingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSeatingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSeatingRequestMultiError) AllErrors() []error { return m }

// CreateSeatingRequestValidationError is the validation error returned by
// CreateSeatingRequest.Validate if the designated constraints aren't met.
type CreateSeatingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSeatingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSeatingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSeatingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSeatingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSeatingRequestValidationError) ErrorName() string {
	return "CreateSeatingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSeatingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSeatingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSeatingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSeatingRequestValidationError{}

// Validate checks the field values on UpdateSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSeatingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSeatingRequestMultiError, or nil if none found.
func (m *UpdateSeatingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSeatingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSeating()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSeatingRequestValidationError{
					field:  "Seating",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSeatingRequestValidationError{
					field:  "Seating",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSeating()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSeatingRequestValidationError{
				field:  "Seating",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSeatingRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSeatingRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSeatingRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSeatingRequestMultiError(errors)
	}

	return nil
}

// UpdateSeatingRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSeatingRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSeatingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSeatingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSeatingRequestMultiError) AllErrors() []error { return m }

// UpdateSeatingRequestValidationError is the validation error returned by
// UpdateSeatingRequest.Validate if the designated constraints aren't met.
type UpdateSeatingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSeatingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSeatingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSeatingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSeatingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSeatingRequestValidationError) ErrorName() string {
	return "UpdateSeatingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSeatingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSeatingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSeatingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSeatingRequestValidationError{}

// Validate checks the field values on DeleteSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSeatingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSeatingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSeatingRequestMultiError, or nil if none found.
func (m *DeleteSeatingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSeatingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = DeleteSeatingRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteSeatingRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteSeatingRequest) _validateUuid(uuid string) error {
	if matched := _seating_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteSeatingRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteSeatingRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteSeatingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSeatingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSeatingRequestMultiError) AllErrors() []error { return m }

// DeleteSeatingRequestValidationError is the validation error returned by
// DeleteSeatingRequest.Validate if the designated constraints aren't met.
type DeleteSeatingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSeatingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSeatingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSeatingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSeatingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSeatingRequestValidationError) ErrorName() string {
	return "DeleteSeatingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSeatingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSeatingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSeatingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSeatingRequestValidationError{}
