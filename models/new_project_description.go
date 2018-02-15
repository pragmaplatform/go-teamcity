// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewProjectDescription new project description
// swagger:model newProjectDescription
type NewProjectDescription struct {

	// build types ids map
	BuildTypesIdsMap *Properties `json:"buildTypesIdsMap,omitempty"`

	// copy all associated settings
	CopyAllAssociatedSettings *bool `json:"copyAllAssociatedSettings,omitempty" xml:"copyAllAssociatedSettings"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// parent project
	ParentProject *Project `json:"parentProject,omitempty"`

	// projects ids map
	ProjectsIdsMap *Properties `json:"projectsIdsMap,omitempty"`

	// source project
	SourceProject *Project `json:"sourceProject,omitempty"`

	// source project locator
	SourceProjectLocator string `json:"sourceProjectLocator,omitempty" xml:"sourceProjectLocator"`

	// vcs roots ids map
	VcsRootsIdsMap *Properties `json:"vcsRootsIdsMap,omitempty"`
}

// Validate validates this new project description
func (m *NewProjectDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildTypesIdsMap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParentProject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProjectsIdsMap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceProject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcsRootsIdsMap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewProjectDescription) validateBuildTypesIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildTypesIdsMap) { // not required
		return nil
	}

	if m.BuildTypesIdsMap != nil {

		if err := m.BuildTypesIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypesIdsMap")
			}
			return err
		}
	}

	return nil
}

func (m *NewProjectDescription) validateParentProject(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentProject) { // not required
		return nil
	}

	if m.ParentProject != nil {

		if err := m.ParentProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentProject")
			}
			return err
		}
	}

	return nil
}

func (m *NewProjectDescription) validateProjectsIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectsIdsMap) { // not required
		return nil
	}

	if m.ProjectsIdsMap != nil {

		if err := m.ProjectsIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projectsIdsMap")
			}
			return err
		}
	}

	return nil
}

func (m *NewProjectDescription) validateSourceProject(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceProject) { // not required
		return nil
	}

	if m.SourceProject != nil {

		if err := m.SourceProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceProject")
			}
			return err
		}
	}

	return nil
}

func (m *NewProjectDescription) validateVcsRootsIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsRootsIdsMap) { // not required
		return nil
	}

	if m.VcsRootsIdsMap != nil {

		if err := m.VcsRootsIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRootsIdsMap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewProjectDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewProjectDescription) UnmarshalBinary(b []byte) error {
	var res NewProjectDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
