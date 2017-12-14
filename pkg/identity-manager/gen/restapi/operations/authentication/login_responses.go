///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// LoginFoundCode is the HTTP code returned for type LoginFound
const LoginFoundCode int = 302

/*LoginFound redirect

swagger:response loginFound
*/
type LoginFound struct {
	/*redirect location
	  Required: true
	*/
	Location string `json:"Location"`
}

// NewLoginFound creates LoginFound with default headers values
func NewLoginFound() *LoginFound {
	return &LoginFound{}
}

// WithLocation adds the location to the login found response
func (o *LoginFound) WithLocation(location string) *LoginFound {
	o.Location = location
	return o
}

// SetLocation sets the location to the login found response
func (o *LoginFound) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *LoginFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(302)
}

/*LoginDefault error

swagger:response loginDefault
*/
type LoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoginDefault creates LoginDefault with default headers values
func NewLoginDefault(code int) *LoginDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login default response
func (o *LoginDefault) WithStatusCode(code int) *LoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login default response
func (o *LoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the login default response
func (o *LoginDefault) WithPayload(payload *models.Error) *LoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login default response
func (o *LoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}