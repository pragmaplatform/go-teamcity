// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProjectsReader is a Reader for the DeleteProjects structure.
type DeleteProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteProjectsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteProjectsDefault creates a DeleteProjectsDefault with default headers values
func NewDeleteProjectsDefault(code int) *DeleteProjectsDefault {
	return &DeleteProjectsDefault{
		_statusCode: code,
	}
}

/*DeleteProjectsDefault handles this case with default header values.

successful operation
*/
type DeleteProjectsDefault struct {
	_statusCode int
}

// Code gets the status code for the delete projects default response
func (o *DeleteProjectsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectsDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agentPools/{agentPoolLocator}/projects][%d] deleteProjects default ", o._statusCode)
}

func (o *DeleteProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
