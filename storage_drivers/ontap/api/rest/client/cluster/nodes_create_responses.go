// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NodesCreateReader is a Reader for the NodesCreate structure.
type NodesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewNodesCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesCreateAccepted creates a NodesCreateAccepted with default headers values
func NewNodesCreateAccepted() *NodesCreateAccepted {
	return &NodesCreateAccepted{}
}

/*
NodesCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NodesCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this nodes create accepted response has a 2xx status code
func (o *NodesCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes create accepted response has a 3xx status code
func (o *NodesCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes create accepted response has a 4xx status code
func (o *NodesCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes create accepted response has a 5xx status code
func (o *NodesCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes create accepted response a status code equal to that given
func (o *NodesCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *NodesCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateAccepted  %+v", 202, o.Payload)
}

func (o *NodesCreateAccepted) String() string {
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateAccepted  %+v", 202, o.Payload)
}

func (o *NodesCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NodesCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesCreateDefault creates a NodesCreateDefault with default headers values
func NewNodesCreateDefault(code int) *NodesCreateDefault {
	return &NodesCreateDefault{
		_statusCode: code,
	}
}

/*
	NodesCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262245 | The value provided was invalid. |
| 1179795 | A node being added is already in the cluster. |
| 1179813 | Fields set for one node must be set for all nodes. |
| 1179817 | The IP address, subnet mask, and gateway must all be provided for cluster manangement interface. |
| 1179818 | The IP address and gateway must be of the same family. |
| 1179821 | An IP address and subnet mask conflicts with an existing entry. |
| 131727360 | A node cannot be added to the cluster. This is a generic code, see response message for details. |
*/
type NodesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nodes create default response
func (o *NodesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this nodes create default response has a 2xx status code
func (o *NodesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nodes create default response has a 3xx status code
func (o *NodesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nodes create default response has a 4xx status code
func (o *NodesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nodes create default response has a 5xx status code
func (o *NodesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nodes create default response a status code equal to that given
func (o *NodesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NodesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodes_create default  %+v", o._statusCode, o.Payload)
}

func (o *NodesCreateDefault) String() string {
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodes_create default  %+v", o._statusCode, o.Payload)
}

func (o *NodesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
