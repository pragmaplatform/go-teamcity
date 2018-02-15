// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ReplaceTagsReader is a Reader for the ReplaceTags structure.
type ReplaceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceTagsOK creates a ReplaceTagsOK with default headers values
func NewReplaceTagsOK() *ReplaceTagsOK {
	return &ReplaceTagsOK{}
}

/*ReplaceTagsOK handles this case with default header values.

successful operation
*/
type ReplaceTagsOK struct {
	Payload *models.Tags
}

func (o *ReplaceTagsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/builds/{buildLocator}/tags][%d] replaceTagsOK  %+v", 200, o.Payload)
}

func (o *ReplaceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
