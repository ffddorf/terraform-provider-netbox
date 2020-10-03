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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimManufacturersPartialUpdateReader is a Reader for the DcimManufacturersPartialUpdate structure.
type DcimManufacturersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersPartialUpdateOK creates a DcimManufacturersPartialUpdateOK with default headers values
func NewDcimManufacturersPartialUpdateOK() *DcimManufacturersPartialUpdateOK {
	return &DcimManufacturersPartialUpdateOK{}
}

/*DcimManufacturersPartialUpdateOK handles this case with default header values.

DcimManufacturersPartialUpdateOK dcim manufacturers partial update o k
*/
type DcimManufacturersPartialUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersPartialUpdateDefault creates a DcimManufacturersPartialUpdateDefault with default headers values
func NewDcimManufacturersPartialUpdateDefault(code int) *DcimManufacturersPartialUpdateDefault {
	return &DcimManufacturersPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimManufacturersPartialUpdateDefault handles this case with default header values.

DcimManufacturersPartialUpdateDefault dcim manufacturers partial update default
*/
type DcimManufacturersPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers partial update default response
func (o *DcimManufacturersPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcim_manufacturers_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
