// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: models/tenant.proto

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
var _tenant_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tenant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TenantMultiError, or nil if none found.
func (m *Tenant) ValidateAll() error {
	return m.validate(true)
}

func (m *Tenant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {

		if m.GetId() != "" {

			if err := m._validateUuid(m.GetId()); err != nil {
				err = TenantValidationError{
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

	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return TenantMultiError(errors)
	}

	return nil
}

func (m *Tenant) _validateUuid(uuid string) error {
	if matched := _tenant_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// TenantMultiError is an error wrapping multiple validation errors returned by
// Tenant.ValidateAll() if the designated constraints aren't met.
type TenantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantMultiError) AllErrors() []error { return m }

// TenantValidationError is the validation error returned by Tenant.Validate if
// the designated constraints aren't met.
type TenantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantValidationError) ErrorName() string { return "TenantValidationError" }

// Error satisfies the builtin error interface
func (e TenantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantValidationError{}
