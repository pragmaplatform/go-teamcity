// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ServeVersionReader is a Reader for the ServeVersion structure.
type ServeVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeVersionOK creates a ServeVersionOK with default headers values
func NewServeVersionOK() *ServeVersionOK {
	return &ServeVersionOK{}
}

/*ServeVersionOK handles this case with default header values.

successful operation
*/
type ServeVersionOK struct {
	Payload string
}

func (o *ServeVersionOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/version][%d] serveVersionOK  %+v", 200, o.Payload)
}

func (o *ServeVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
