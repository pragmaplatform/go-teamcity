// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// GetFeaturesReader is a Reader for the GetFeatures structure.
type GetFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeaturesOK creates a GetFeaturesOK with default headers values
func NewGetFeaturesOK() *GetFeaturesOK {
	return &GetFeaturesOK{}
}

/*GetFeaturesOK handles this case with default header values.

successful operation
*/
type GetFeaturesOK struct {
	Payload *models.Features
}

func (o *GetFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/features][%d] getFeaturesOK  %+v", 200, o.Payload)
}

func (o *GetFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Features)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
