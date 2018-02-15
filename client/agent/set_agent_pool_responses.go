// Code generated by go-swagger; DO NOT EDIT.

package agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// SetAgentPoolReader is a Reader for the SetAgentPool structure.
type SetAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAgentPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAgentPoolOK creates a SetAgentPoolOK with default headers values
func NewSetAgentPoolOK() *SetAgentPoolOK {
	return &SetAgentPoolOK{}
}

/*SetAgentPoolOK handles this case with default header values.

successful operation
*/
type SetAgentPoolOK struct {
	Payload *models.AgentPool
}

func (o *SetAgentPoolOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/agents/{agentLocator}/pool][%d] setAgentPoolOK  %+v", 200, o.Payload)
}

func (o *SetAgentPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
