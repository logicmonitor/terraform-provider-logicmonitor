// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Widget widget
//
// swagger:model Widget
type Widget struct {

	// The id of the dashboard the widget belongs to
	// Example: 1
	// Required: true
	DashboardID *int32 `json:"dashboardId"`

	// The description of the widget
	// Example: Devices By Type
	Description string `json:"description,omitempty"`

	// The Id of the widget
	ID int32 `json:"id,omitempty"`

	// The refresh interval of the widget, in minutes
	// Example: 5
	Interval int32 `json:"interval,omitempty"`

	// The user that last updated the widget
	// Read Only: true
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// The time that corresponds to when the widget was last updated, in epoch format
	// Read Only: true
	LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

	// The name of the widget
	// Example: Test
	// Required: true
	Name *string `json:"name"`

	// The color scheme of the widget. Options are: borderPurple | borderGray | borderBlue | solidPurple | solidGray | solidBlue | simplePurple | simpleBlue | simpleGray | newBorderGray | newBorderBlue | newBorderDarkBlue | newSolidGray | newSolidBlue | newSolidDarkBlue | newSimpleGray | newSimpleBlue |newSimpleDarkBlue
	// Example: newBorderBlue
	Theme string `json:"theme,omitempty"`

	// The default timescale of the widget
	Timescale string `json:"timescale,omitempty"`

	// alert | batchjob | flash | gmap | ngraph | ograph | cgraph | sgraph | netflowgraph | groupNetflowGraph | netflow | groupNetflow | html | bigNumber | gauge | pieChart | table | dynamicTable | deviceSLA | text | statsd | deviceStatus | serviceAlert | noc | websiteOverview | websiteOverallStatus | websiteIndividualStatus | websiteSLA
	// Example: bigNumber
	// Required: true
	Type *string `json:"type"`

	// The permission level of the user who last modified the widget
	// Example: write
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this widget
func (m *Widget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Widget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID); err != nil {
		return err
	}

	return nil
}

func (m *Widget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Widget) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this widget based on the context it is used
func (m *Widget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedOn(ctx, formats); err != nil {
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

func (m *Widget) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedBy", "body", string(m.LastUpdatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *Widget) contextValidateLastUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedOn", "body", int64(m.LastUpdatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *Widget) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Widget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Widget) UnmarshalBinary(b []byte) error {
	var res Widget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
