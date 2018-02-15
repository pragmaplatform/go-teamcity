// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// GetParentChangesReader is a Reader for the GetParentChanges structure.
type GetParentChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParentChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParentChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParentChangesOK creates a GetParentChangesOK with default headers values
func NewGetParentChangesOK() *GetParentChangesOK {
	return &GetParentChangesOK{}
}

/*GetParentChangesOK handles this case with default header values.

successful operation
*/
type GetParentChangesOK struct {
	Payload *models.Changes
}

func (o *GetParentChangesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}/parentChanges][%d] getParentChangesOK  %+v", 200, o.Payload)
}

func (o *GetParentChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Changes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
