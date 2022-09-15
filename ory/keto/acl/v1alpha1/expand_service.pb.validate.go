// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ory/keto/acl/v1alpha1/expand_service.proto

package acl

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

// Validate checks the field values on ExpandRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExpandRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExpandRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExpandRequestMultiError, or
// nil if none found.
func (m *ExpandRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExpandRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExpandRequestValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExpandRequestValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandRequestValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MaxDepth

	// no validation rules for Snaptoken

	if len(errors) > 0 {
		return ExpandRequestMultiError(errors)
	}

	return nil
}

// ExpandRequestMultiError is an error wrapping multiple validation errors
// returned by ExpandRequest.ValidateAll() if the designated constraints
// aren't met.
type ExpandRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExpandRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExpandRequestMultiError) AllErrors() []error { return m }

// ExpandRequestValidationError is the validation error returned by
// ExpandRequest.Validate if the designated constraints aren't met.
type ExpandRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExpandRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExpandRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExpandRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExpandRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExpandRequestValidationError) ErrorName() string { return "ExpandRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExpandRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExpandRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExpandRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExpandRequestValidationError{}

// Validate checks the field values on ExpandResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExpandResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExpandResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExpandResponseMultiError,
// or nil if none found.
func (m *ExpandResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExpandResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTree()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExpandResponseValidationError{
					field:  "Tree",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExpandResponseValidationError{
					field:  "Tree",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTree()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandResponseValidationError{
				field:  "Tree",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ExpandResponseMultiError(errors)
	}

	return nil
}

// ExpandResponseMultiError is an error wrapping multiple validation errors
// returned by ExpandResponse.ValidateAll() if the designated constraints
// aren't met.
type ExpandResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExpandResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExpandResponseMultiError) AllErrors() []error { return m }

// ExpandResponseValidationError is the validation error returned by
// ExpandResponse.Validate if the designated constraints aren't met.
type ExpandResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExpandResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExpandResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExpandResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExpandResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExpandResponseValidationError) ErrorName() string { return "ExpandResponseValidationError" }

// Error satisfies the builtin error interface
func (e ExpandResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExpandResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExpandResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExpandResponseValidationError{}

// Validate checks the field values on SubjectTree with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubjectTree) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectTree with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubjectTreeMultiError, or
// nil if none found.
func (m *SubjectTree) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectTree) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeType

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubjectTreeValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubjectTreeValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubjectTreeValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubjectTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubjectTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubjectTreeValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SubjectTreeMultiError(errors)
	}

	return nil
}

// SubjectTreeMultiError is an error wrapping multiple validation errors
// returned by SubjectTree.ValidateAll() if the designated constraints aren't met.
type SubjectTreeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectTreeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectTreeMultiError) AllErrors() []error { return m }

// SubjectTreeValidationError is the validation error returned by
// SubjectTree.Validate if the designated constraints aren't met.
type SubjectTreeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectTreeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectTreeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectTreeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectTreeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectTreeValidationError) ErrorName() string { return "SubjectTreeValidationError" }

// Error satisfies the builtin error interface
func (e SubjectTreeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectTree.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectTreeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectTreeValidationError{}
