// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRequestDetailsReader is a Reader for the PostRequestDetails structure.
type PostRequestDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRequestDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRequestDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRequestDetailsOK creates a PostRequestDetailsOK with default headers values
func NewPostRequestDetailsOK() *PostRequestDetailsOK {
	return &PostRequestDetailsOK{}
}

/*PostRequestDetailsOK handles this case with default header values.

successful operation
*/
type PostRequestDetailsOK struct {
	Payload string
}

func (o *PostRequestDetailsOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/debug/currentRequest/details{extra}][%d] postRequestDetailsOK  %+v", 200, o.Payload)
}

func (o *PostRequestDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
