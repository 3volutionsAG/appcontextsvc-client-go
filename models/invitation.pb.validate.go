// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: models/invitation.proto

package models

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
var _invitation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on InvitationState with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *InvitationState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InvitationState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InvitationStateMultiError, or nil if none found.
func (m *InvitationState) ValidateAll() error {
	return m.validate(true)
}

func (m *InvitationState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = InvitationStateValidationError{
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

	// no validation rules for State

	if len(errors) > 0 {
		return InvitationStateMultiError(errors)
	}

	return nil
}

func (m *InvitationState) _validateUuid(uuid string) error {
	if matched := _invitation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// InvitationStateMultiError is an error wrapping multiple validation errors
// returned by InvitationState.ValidateAll() if the designated constraints
// aren't met.
type InvitationStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InvitationStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InvitationStateMultiError) AllErrors() []error { return m }

// InvitationStateValidationError is the validation error returned by
// InvitationState.Validate if the designated constraints aren't met.
type InvitationStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InvitationStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InvitationStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InvitationStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InvitationStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InvitationStateValidationError) ErrorName() string { return "InvitationStateValidationError" }

// Error satisfies the builtin error interface
func (e InvitationStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInvitationState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InvitationStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InvitationStateValidationError{}

// Validate checks the field values on Invitation with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Invitation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Invitation with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InvitationMultiError, or
// nil if none found.
func (m *Invitation) ValidateAll() error {
	return m.validate(true)
}

func (m *Invitation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = InvitationValidationError{
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

	if m.GetUserprofileId() != "" {

		if err := m._validateUuid(m.GetUserprofileId()); err != nil {
			err = InvitationValidationError{
				field:  "UserprofileId",
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
		return InvitationMultiError(errors)
	}

	return nil
}

func (m *Invitation) _validateUuid(uuid string) error {
	if matched := _invitation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// InvitationMultiError is an error wrapping multiple validation errors
// returned by Invitation.ValidateAll() if the designated constraints aren't met.
type InvitationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InvitationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InvitationMultiError) AllErrors() []error { return m }

// InvitationValidationError is the validation error returned by
// Invitation.Validate if the designated constraints aren't met.
type InvitationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InvitationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InvitationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InvitationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InvitationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InvitationValidationError) ErrorName() string { return "InvitationValidationError" }

// Error satisfies the builtin error interface
func (e InvitationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInvitation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InvitationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InvitationValidationError{}