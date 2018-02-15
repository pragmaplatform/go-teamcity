// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePoolProjectReader is a Reader for the DeletePoolProject structure.
type DeletePoolProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePoolProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePoolProjectDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePoolProjectDefault creates a DeletePoolProjectDefault with default headers values
func NewDeletePoolProjectDefault(code int) *DeletePoolProjectDefault {
	return &DeletePoolProjectDefault{
		_statusCode: code,
	}
}

/*DeletePoolProjectDefault handles this case with default header values.

successful operation
*/
type DeletePoolProjectDefault struct {
	_statusCode int
}

// Code gets the status code for the delete pool project default response
func (o *DeletePoolProjectDefault) Code() int {
	return o._statusCode
}

func (o *DeletePoolProjectDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agentPools/{agentPoolLocator}/projects/{projectLocator}][%d] deletePoolProject default ", o._statusCode)
}

func (o *DeletePoolProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
