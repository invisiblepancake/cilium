// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetIdentityReader is a Reader for the GetIdentity structure.
type GetIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 520:
		result := NewGetIdentityUnreachable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 521:
		result := NewGetIdentityInvalidStorageFormat()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /identity] GetIdentity", response, response.Code())
	}
}

// NewGetIdentityOK creates a GetIdentityOK with default headers values
func NewGetIdentityOK() *GetIdentityOK {
	return &GetIdentityOK{}
}

/*
GetIdentityOK describes a response with status code 200, with default header values.

Success
*/
type GetIdentityOK struct {
	Payload []*models.Identity
}

// IsSuccess returns true when this get identity o k response has a 2xx status code
func (o *GetIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get identity o k response has a 3xx status code
func (o *GetIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity o k response has a 4xx status code
func (o *GetIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity o k response has a 5xx status code
func (o *GetIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity o k response a status code equal to that given
func (o *GetIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get identity o k response
func (o *GetIdentityOK) Code() int {
	return 200
}

func (o *GetIdentityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityOK %s", 200, payload)
}

func (o *GetIdentityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityOK %s", 200, payload)
}

func (o *GetIdentityOK) GetPayload() []*models.Identity {
	return o.Payload
}

func (o *GetIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityNotFound creates a GetIdentityNotFound with default headers values
func NewGetIdentityNotFound() *GetIdentityNotFound {
	return &GetIdentityNotFound{}
}

/*
GetIdentityNotFound describes a response with status code 404, with default header values.

Identities with provided parameters not found
*/
type GetIdentityNotFound struct {
}

// IsSuccess returns true when this get identity not found response has a 2xx status code
func (o *GetIdentityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity not found response has a 3xx status code
func (o *GetIdentityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity not found response has a 4xx status code
func (o *GetIdentityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity not found response has a 5xx status code
func (o *GetIdentityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity not found response a status code equal to that given
func (o *GetIdentityNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get identity not found response
func (o *GetIdentityNotFound) Code() int {
	return 404
}

func (o *GetIdentityNotFound) Error() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityNotFound", 404)
}

func (o *GetIdentityNotFound) String() string {
	return fmt.Sprintf("[GET /identity][%d] getIdentityNotFound", 404)
}

func (o *GetIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIdentityUnreachable creates a GetIdentityUnreachable with default headers values
func NewGetIdentityUnreachable() *GetIdentityUnreachable {
	return &GetIdentityUnreachable{}
}

/*
GetIdentityUnreachable describes a response with status code 520, with default header values.

Identity storage unreachable. Likely a network problem.
*/
type GetIdentityUnreachable struct {
	Payload models.Error
}

// IsSuccess returns true when this get identity unreachable response has a 2xx status code
func (o *GetIdentityUnreachable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity unreachable response has a 3xx status code
func (o *GetIdentityUnreachable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity unreachable response has a 4xx status code
func (o *GetIdentityUnreachable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity unreachable response has a 5xx status code
func (o *GetIdentityUnreachable) IsServerError() bool {
	return true
}

// IsCode returns true when this get identity unreachable response a status code equal to that given
func (o *GetIdentityUnreachable) IsCode(code int) bool {
	return code == 520
}

// Code gets the status code for the get identity unreachable response
func (o *GetIdentityUnreachable) Code() int {
	return 520
}

func (o *GetIdentityUnreachable) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityUnreachable %s", 520, payload)
}

func (o *GetIdentityUnreachable) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityUnreachable %s", 520, payload)
}

func (o *GetIdentityUnreachable) GetPayload() models.Error {
	return o.Payload
}

func (o *GetIdentityUnreachable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityInvalidStorageFormat creates a GetIdentityInvalidStorageFormat with default headers values
func NewGetIdentityInvalidStorageFormat() *GetIdentityInvalidStorageFormat {
	return &GetIdentityInvalidStorageFormat{}
}

/*
GetIdentityInvalidStorageFormat describes a response with status code 521, with default header values.

Invalid identity format in storage
*/
type GetIdentityInvalidStorageFormat struct {
	Payload models.Error
}

// IsSuccess returns true when this get identity invalid storage format response has a 2xx status code
func (o *GetIdentityInvalidStorageFormat) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity invalid storage format response has a 3xx status code
func (o *GetIdentityInvalidStorageFormat) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity invalid storage format response has a 4xx status code
func (o *GetIdentityInvalidStorageFormat) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity invalid storage format response has a 5xx status code
func (o *GetIdentityInvalidStorageFormat) IsServerError() bool {
	return true
}

// IsCode returns true when this get identity invalid storage format response a status code equal to that given
func (o *GetIdentityInvalidStorageFormat) IsCode(code int) bool {
	return code == 521
}

// Code gets the status code for the get identity invalid storage format response
func (o *GetIdentityInvalidStorageFormat) Code() int {
	return 521
}

func (o *GetIdentityInvalidStorageFormat) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityInvalidStorageFormat %s", 521, payload)
}

func (o *GetIdentityInvalidStorageFormat) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identity][%d] getIdentityInvalidStorageFormat %s", 521, payload)
}

func (o *GetIdentityInvalidStorageFormat) GetPayload() models.Error {
	return o.Payload
}

func (o *GetIdentityInvalidStorageFormat) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
