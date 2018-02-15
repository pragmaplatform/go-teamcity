// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserFieldReader is a Reader for the DeleteUserField structure.
type DeleteUserFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteUserFieldDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteUserFieldDefault creates a DeleteUserFieldDefault with default headers values
func NewDeleteUserFieldDefault(code int) *DeleteUserFieldDefault {
	return &DeleteUserFieldDefault{
		_statusCode: code,
	}
}

/*DeleteUserFieldDefault handles this case with default header values.

successful operation
*/
type DeleteUserFieldDefault struct {
	_statusCode int
}

// Code gets the status code for the delete user field default response
func (o *DeleteUserFieldDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserFieldDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/users/{userLocator}/{field}][%d] deleteUserField default ", o._statusCode)
}

func (o *DeleteUserFieldDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
