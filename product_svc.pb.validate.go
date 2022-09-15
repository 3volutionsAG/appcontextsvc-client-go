// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product_svc.proto

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
var _product_svc_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateSkuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateSkuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSkuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSkuRequestMultiError, or nil if none found.
func (m *CreateSkuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSkuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSku()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSkuRequestValidationError{
					field:  "Sku",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSkuRequestValidationError{
					field:  "Sku",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSku()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSkuRequestValidationError{
				field:  "Sku",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateSkuRequestMultiError(errors)
	}

	return nil
}

// CreateSkuRequestMultiError is an error wrapping multiple validation errors
// returned by CreateSkuRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateSkuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSkuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSkuRequestMultiError) AllErrors() []error { return m }

// CreateSkuRequestValidationError is the validation error returned by
// CreateSkuRequest.Validate if the designated constraints aren't met.
type CreateSkuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSkuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSkuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSkuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSkuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSkuRequestValidationError) ErrorName() string { return "CreateSkuRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateSkuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSkuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSkuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSkuRequestValidationError{}

// Validate checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductsResponseMultiError, or nil if none found.
func (m *ListProductsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProducts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProductsResponseValidationError{
					field:  fmt.Sprintf("Products[%v]", idx),
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
		return ListProductsResponseMultiError(errors)
	}

	return nil
}

// ListProductsResponseMultiError is an error wrapping multiple validation
// errors returned by ListProductsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListProductsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductsResponseMultiError) AllErrors() []error { return m }

// ListProductsResponseValidationError is the validation error returned by
// ListProductsResponse.Validate if the designated constraints aren't met.
type ListProductsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductsResponseValidationError) ErrorName() string {
	return "ListProductsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductsResponseValidationError{}

// Validate checks the field values on ListSkuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSkuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSkuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSkuResponseMultiError, or nil if none found.
func (m *ListSkuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSkuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSkus() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSkuResponseValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSkuResponseValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSkuResponseValidationError{
					field:  fmt.Sprintf("Skus[%v]", idx),
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
		return ListSkuResponseMultiError(errors)
	}

	return nil
}

// ListSkuResponseMultiError is an error wrapping multiple validation errors
// returned by ListSkuResponse.ValidateAll() if the designated constraints
// aren't met.
type ListSkuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSkuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSkuResponseMultiError) AllErrors() []error { return m }

// ListSkuResponseValidationError is the validation error returned by
// ListSkuResponse.Validate if the designated constraints aren't met.
type ListSkuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkuResponseValidationError) ErrorName() string { return "ListSkuResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListSkuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkuResponseValidationError{}

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductRequestMultiError, or nil if none found.
func (m *GetProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetProductRequestValidationError{
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
		return GetProductRequestMultiError(errors)
	}

	return nil
}

func (m *GetProductRequest) _validateUuid(uuid string) error {
	if matched := _product_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetProductRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductRequestMultiError) AllErrors() []error { return m }

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRequestMultiError, or nil if none found.
func (m *CreateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductRequestValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateProductRequestMultiError(errors)
	}

	return nil
}

// CreateProductRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRequestMultiError) AllErrors() []error { return m }

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProductRequestMultiError, or nil if none found.
func (m *UpdateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProductRequestValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateProductRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateProductRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProductRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateProductRequestMultiError(errors)
	}

	return nil
}

// UpdateProductRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProductRequestMultiError) AllErrors() []error { return m }

// UpdateProductRequestValidationError is the validation error returned by
// UpdateProductRequest.Validate if the designated constraints aren't met.
type UpdateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductRequestValidationError) ErrorName() string {
	return "UpdateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductRequestValidationError{}

// Validate checks the field values on UpdateSkuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateSkuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSkuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSkuRequestMultiError, or nil if none found.
func (m *UpdateSkuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSkuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSku()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSkuRequestValidationError{
					field:  "Sku",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSkuRequestValidationError{
					field:  "Sku",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSku()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSkuRequestValidationError{
				field:  "Sku",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSkuRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSkuRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSkuRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSkuRequestMultiError(errors)
	}

	return nil
}

// UpdateSkuRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateSkuRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateSkuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSkuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSkuRequestMultiError) AllErrors() []error { return m }

// UpdateSkuRequestValidationError is the validation error returned by
// UpdateSkuRequest.Validate if the designated constraints aren't met.
type UpdateSkuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSkuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSkuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSkuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSkuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSkuRequestValidationError) ErrorName() string { return "UpdateSkuRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateSkuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSkuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSkuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSkuRequestValidationError{}

// Validate checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductRequestMultiError, or nil if none found.
func (m *DeleteProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = DeleteProductRequestValidationError{
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
		return DeleteProductRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteProductRequest) _validateUuid(uuid string) error {
	if matched := _product_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteProductRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProductRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductRequestMultiError) AllErrors() []error { return m }

// DeleteProductRequestValidationError is the validation error returned by
// DeleteProductRequest.Validate if the designated constraints aren't met.
type DeleteProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductRequestValidationError) ErrorName() string {
	return "DeleteProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductRequestValidationError{}

// Validate checks the field values on DeleteSkuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteSkuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSkuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSkuRequestMultiError, or nil if none found.
func (m *DeleteSkuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSkuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = DeleteSkuRequestValidationError{
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
		return DeleteSkuRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteSkuRequest) _validateUuid(uuid string) error {
	if matched := _product_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteSkuRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteSkuRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteSkuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSkuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSkuRequestMultiError) AllErrors() []error { return m }

// DeleteSkuRequestValidationError is the validation error returned by
// DeleteSkuRequest.Validate if the designated constraints aren't met.
type DeleteSkuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSkuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSkuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSkuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSkuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSkuRequestValidationError) ErrorName() string { return "DeleteSkuRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteSkuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSkuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSkuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSkuRequestValidationError{}

// Validate checks the field values on GetSkuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetSkuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSkuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetSkuRequestMultiError, or
// nil if none found.
func (m *GetSkuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSkuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetSkuRequestValidationError{
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
		return GetSkuRequestMultiError(errors)
	}

	return nil
}

func (m *GetSkuRequest) _validateUuid(uuid string) error {
	if matched := _product_svc_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetSkuRequestMultiError is an error wrapping multiple validation errors
// returned by GetSkuRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSkuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSkuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSkuRequestMultiError) AllErrors() []error { return m }

// GetSkuRequestValidationError is the validation error returned by
// GetSkuRequest.Validate if the designated constraints aren't met.
type GetSkuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSkuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSkuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSkuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSkuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSkuRequestValidationError) ErrorName() string { return "GetSkuRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetSkuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSkuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSkuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSkuRequestValidationError{}
