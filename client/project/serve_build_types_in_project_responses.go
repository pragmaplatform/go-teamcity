// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ServeBuildTypesInProjectReader is a Reader for the ServeBuildTypesInProject structure.
type ServeBuildTypesInProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildTypesInProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeBuildTypesInProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildTypesInProjectOK creates a ServeBuildTypesInProjectOK with default headers values
func NewServeBuildTypesInProjectOK() *ServeBuildTypesInProjectOK {
	return &ServeBuildTypesInProjectOK{}
}

/*ServeBuildTypesInProjectOK handles this case with default header values.

successful operation
*/
type ServeBuildTypesInProjectOK struct {
	Payload *models.BuildTypes
}

func (o *ServeBuildTypesInProjectOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/buildTypes][%d] serveBuildTypesInProjectOK  %+v", 200, o.Payload)
}

func (o *ServeBuildTypesInProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
