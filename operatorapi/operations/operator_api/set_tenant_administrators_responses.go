// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operator_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// SetTenantAdministratorsNoContentCode is the HTTP code returned for type SetTenantAdministratorsNoContent
const SetTenantAdministratorsNoContentCode int = 204

/*
SetTenantAdministratorsNoContent A successful response.

swagger:response setTenantAdministratorsNoContent
*/
type SetTenantAdministratorsNoContent struct {
}

// NewSetTenantAdministratorsNoContent creates SetTenantAdministratorsNoContent with default headers values
func NewSetTenantAdministratorsNoContent() *SetTenantAdministratorsNoContent {

	return &SetTenantAdministratorsNoContent{}
}

// WriteResponse to the client
func (o *SetTenantAdministratorsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
SetTenantAdministratorsDefault Generic error response.

swagger:response setTenantAdministratorsDefault
*/
type SetTenantAdministratorsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetTenantAdministratorsDefault creates SetTenantAdministratorsDefault with default headers values
func NewSetTenantAdministratorsDefault(code int) *SetTenantAdministratorsDefault {
	if code <= 0 {
		code = 500
	}

	return &SetTenantAdministratorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set tenant administrators default response
func (o *SetTenantAdministratorsDefault) WithStatusCode(code int) *SetTenantAdministratorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set tenant administrators default response
func (o *SetTenantAdministratorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the set tenant administrators default response
func (o *SetTenantAdministratorsDefault) WithPayload(payload *models.Error) *SetTenantAdministratorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set tenant administrators default response
func (o *SetTenantAdministratorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetTenantAdministratorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
