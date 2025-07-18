// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebsiteGroup website group
// Example: isResource
//
// swagger:model WebsiteGroup
type WebsiteGroup struct {

	// The description of the group
	// Example: Amazon web and ping checks
	Description string `json:"description,omitempty"`

	// true: alerting is disabled for the websites in the group
	// false: alerting is enabled for the websites in the group
	// If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group
	// Example: false
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// The full path of the group
	// Read Only: true
	FullPath string `json:"fullPath,omitempty"`

	// has websites disabled
	// Read Only: true
	HasWebsitesDisabled *bool `json:"hasWebsitesDisabled,omitempty"`

	// The Id of the group
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the group
	// Example: Amazon Website Checks
	// Required: true
	Name *string `json:"name"`

	// The number of direct website groups in this group (exlcuding those in subgroups)
	// Read Only: true
	NumOfDirectSubGroups int32 `json:"numOfDirectSubGroups,omitempty"`

	// num of direct websites
	// Read Only: true
	NumOfDirectWebsites int32 `json:"numOfDirectWebsites,omitempty"`

	// num of websites
	// Read Only: true
	NumOfWebsites int32 `json:"numOfWebsites,omitempty"`

	// The Id of the parent group. If parentId=1 then the group exists under the root  group
	// Example: 1
	ParentID int32 `json:"parentId,omitempty"`

	// properties
	Properties []*NameAndValue `json:"properties"`

	// true: monitoring is disabled for the websites in the group
	// false: monitoring is enabled for the websites in the group
	// If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group
	// Example: false
	StopMonitoring bool `json:"stopMonitoring,omitempty"`

	// An object that indicates the websites locations.
	// eg. {'all': false, smgId:[1,2,3], collectorIds:[14,16]}
	TestLocation *WebsiteLocation `json:"testLocation,omitempty"`

	// The permission level of the user that made the API request. Acceptable values are: write, read, ack
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this website group
func (m *WebsiteGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {
		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {
			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebsiteGroup) validateTestLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.TestLocation) { // not required
		return nil
	}

	if m.TestLocation != nil {
		if err := m.TestLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("testLocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this website group based on the context it is used
func (m *WebsiteGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHasWebsitesDisabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfDirectSubGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfDirectWebsites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumOfWebsites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTestLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteGroup) contextValidateFullPath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fullPath", "body", string(m.FullPath)); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateHasWebsitesDisabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hasWebsitesDisabled", "body", m.HasWebsitesDisabled); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateNumOfDirectSubGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfDirectSubGroups", "body", int32(m.NumOfDirectSubGroups)); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateNumOfDirectWebsites(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfDirectWebsites", "body", int32(m.NumOfDirectWebsites)); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateNumOfWebsites(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numOfWebsites", "body", int32(m.NumOfWebsites)); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteGroup) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Properties); i++ {

		if m.Properties[i] != nil {
			if err := m.Properties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebsiteGroup) contextValidateTestLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.TestLocation != nil {
		if err := m.TestLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("testLocation")
			}
			return err
		}
	}

	return nil
}

func (m *WebsiteGroup) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteGroup) UnmarshalBinary(b []byte) error {
	var res WebsiteGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
