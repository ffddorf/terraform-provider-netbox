// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersUsersPartialUpdateReader is a Reader for the UsersUsersPartialUpdate structure.
type UsersUsersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersPartialUpdateOK creates a UsersUsersPartialUpdateOK with default headers values
func NewUsersUsersPartialUpdateOK() *UsersUsersPartialUpdateOK {
	return &UsersUsersPartialUpdateOK{}
}

/*UsersUsersPartialUpdateOK handles this case with default header values.

UsersUsersPartialUpdateOK users users partial update o k
*/
type UsersUsersPartialUpdateOK struct {
	Payload *models.User
}

func (o *UsersUsersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/users/{id}/][%d] usersUsersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersPartialUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersPartialUpdateDefault creates a UsersUsersPartialUpdateDefault with default headers values
func NewUsersUsersPartialUpdateDefault(code int) *UsersUsersPartialUpdateDefault {
	return &UsersUsersPartialUpdateDefault{
		_statusCode: code,
	}
}

/*UsersUsersPartialUpdateDefault handles this case with default header values.

UsersUsersPartialUpdateDefault users users partial update default
*/
type UsersUsersPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users partial update default response
func (o *UsersUsersPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersUsersPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/users/{id}/][%d] users_users_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
