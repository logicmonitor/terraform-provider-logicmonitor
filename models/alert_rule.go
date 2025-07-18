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

// AlertRule alert rule
// Example: isResource
//
// swagger:model AlertRule
type AlertRule struct {

	// The datapoint the alert rule is configured to match
	// Example: *
	// Required: true
	Datapoint *string `json:"datapoint"`

	// The datasource the alert rule is configured to match
	// Example: Port-
	// Required: true
	Datasource *string `json:"datasource"`

	// The device groups and service groups the alert rule is configured to match
	// Example: [ \"Devices by Type\"]
	// Required: true
	// Unique: true
	DeviceGroups []string `json:"deviceGroups"`

	// The device names and service names the alert rule is configured to match
	// Example: [\"Cisco Router\"]
	// Required: true
	// Unique: true
	Devices []string `json:"devices"`

	// The escalation chain associated with the alert rule
	// Read Only: true
	EscalatingChain interface{} `json:"escalatingChain,omitempty"`

	// The id of the escalation chain associated with the alert rule
	// Example: 12
	// Required: true
	EscalatingChainID *int32 `json:"escalatingChainId"`

	// The escalation interval associated with the alert rule, in minutes
	// Example: 15
	// Required: true
	EscalationInterval *int32 `json:"escalationInterval"`

	// The Id of the alert rule
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The instance the alert rule is configured to match
	// Example: *
	// Required: true
	Instance *string `json:"instance"`

	// The alert severity levels the alert rule is configured to match. Acceptable values are: All, Warn, Error, Critical
	// Example: Warn
	LevelStr string `json:"levelStr,omitempty"`

	// The name of the alert rule
	// Example: Warning
	// Required: true
	Name *string `json:"name"`

	// The priority associated with the alert rule
	// Example: 100
	// Required: true
	Priority *int32 `json:"priority"`

	// The resource property filters list
	ResourceProperties []*DeviceProperty `json:"resourceProperties"`

	// Whether or not send anomaly suppressed alert
	// Example: true
	SendAnomalySuppressedAlert bool `json:"sendAnomalySuppressedAlert,omitempty"`

	// Whether or not status notifications for acknowledgements and SDTs should be sent to the alert rule
	// Example: true
	SuppressAlertAckSdt bool `json:"suppressAlertAckSdt,omitempty"`

	// Whether or not alert clear notifications should be sent to the alert rule
	// Example: true
	SuppressAlertClear bool `json:"suppressAlertClear,omitempty"`
}

// Validate validates this alert rule
func (m *AlertRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatapoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatasource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalatingChainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalationInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertRule) validateDatapoint(formats strfmt.Registry) error {

	if err := validate.Required("datapoint", "body", m.Datapoint); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Required("datasource", "body", m.Datasource); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateDeviceGroups(formats strfmt.Registry) error {

	if err := validate.Required("deviceGroups", "body", m.DeviceGroups); err != nil {
		return err
	}

	if err := validate.UniqueItems("deviceGroups", "body", m.DeviceGroups); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("devices", "body", m.Devices); err != nil {
		return err
	}

	if err := validate.UniqueItems("devices", "body", m.Devices); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateEscalatingChainID(formats strfmt.Registry) error {

	if err := validate.Required("escalatingChainId", "body", m.EscalatingChainID); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateEscalationInterval(formats strfmt.Registry) error {

	if err := validate.Required("escalationInterval", "body", m.EscalationInterval); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateInstance(formats strfmt.Registry) error {

	if err := validate.Required("instance", "body", m.Instance); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) validateResourceProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceProperties); i++ {
		if swag.IsZero(m.ResourceProperties[i]) { // not required
			continue
		}

		if m.ResourceProperties[i] != nil {
			if err := m.ResourceProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this alert rule based on the context it is used
func (m *AlertRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertRule) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AlertRule) contextValidateResourceProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceProperties); i++ {

		if m.ResourceProperties[i] != nil {
			if err := m.ResourceProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertRule) UnmarshalBinary(b []byte) error {
	var res AlertRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
