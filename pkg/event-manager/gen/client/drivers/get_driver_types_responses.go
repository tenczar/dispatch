///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriverTypesReader is a Reader for the GetDriverTypes structure.
type GetDriverTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriverTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDriverTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDriverTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDriverTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDriverTypesOK creates a GetDriverTypesOK with default headers values
func NewGetDriverTypesOK() *GetDriverTypesOK {
	return &GetDriverTypesOK{}
}

/*GetDriverTypesOK handles this case with default header values.

Successful operation
*/
type GetDriverTypesOK struct {
	Payload []*v1.EventDriverType
}

func (o *GetDriverTypesOK) Error() string {
	return fmt.Sprintf("[GET /drivertypes][%d] getDriverTypesOK  %+v", 200, o.Payload)
}

func (o *GetDriverTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypesUnauthorized creates a GetDriverTypesUnauthorized with default headers values
func NewGetDriverTypesUnauthorized() *GetDriverTypesUnauthorized {
	return &GetDriverTypesUnauthorized{}
}

/*GetDriverTypesUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetDriverTypesUnauthorized struct {
	Payload *v1.Error
}

func (o *GetDriverTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /drivertypes][%d] getDriverTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDriverTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypesForbidden creates a GetDriverTypesForbidden with default headers values
func NewGetDriverTypesForbidden() *GetDriverTypesForbidden {
	return &GetDriverTypesForbidden{}
}

/*GetDriverTypesForbidden handles this case with default header values.

access to this resource is forbidden
*/
type GetDriverTypesForbidden struct {
	Payload *v1.Error
}

func (o *GetDriverTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /drivertypes][%d] getDriverTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetDriverTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverTypesDefault creates a GetDriverTypesDefault with default headers values
func NewGetDriverTypesDefault(code int) *GetDriverTypesDefault {
	return &GetDriverTypesDefault{
		_statusCode: code,
	}
}

/*GetDriverTypesDefault handles this case with default header values.

Unknown error
*/
type GetDriverTypesDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get driver types default response
func (o *GetDriverTypesDefault) Code() int {
	return o._statusCode
}

func (o *GetDriverTypesDefault) Error() string {
	return fmt.Sprintf("[GET /drivertypes][%d] getDriverTypes default  %+v", o._statusCode, o.Payload)
}

func (o *GetDriverTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
