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

// Datasource datasource
// Example: isResource
//
// swagger:model Datasource
type Datasource struct {

	// The Applies To for the LMModule
	// Example: system.deviceId == \"22\
	AppliesTo string `json:"appliesTo,omitempty"`

	// The data source audit version
	// Read Only: true
	AuditVersion int64 `json:"auditVersion,omitempty"`

	// auto discovery config
	AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

	// The metadata checksum for the LMModule content
	// Read Only: true
	Checksum string `json:"checksum,omitempty"`

	// The DataSource data collect interval
	// Example: 100
	// Required: true
	CollectInterval *int32 `json:"collectInterval"`

	// The  method to collect data. The values can be snmp|ping|exs|webpage|wmi|cim|datadump|dns|ipmi|jdbb|script|udp|tcp|xen
	// Example: script
	// Required: true
	CollectMethod *string `json:"collectMethod"`

	// Data collector's attributes to collector data. e.g. a ping data source has a ping collector attribute.
	//  PingCollectorAttributeV1 has two fields. the ip to ping, the data size send to ping
	// Required: true
	CollectorAttribute *CollectorAttribute `json:"collectorAttribute"`

	// The data point list
	DataPoints []*DataPoint `json:"dataPoints"`

	// The description for the LMModule
	// Example: string
	Description string `json:"description,omitempty"`

	// The data source display name
	// Example: test
	DisplayName string `json:"displayName,omitempty"`

	// Enable Auto Discovery or not when this data source has multi instance: false|true
	// Example: false
	EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

	// Enable ERI Discovery or not: false|true
	// Example: false
	EnableEriDiscovery bool `json:"enableEriDiscovery,omitempty"`

	// eri discovery config
	EriDiscoveryConfig *ScriptERIDiscoveryAttributeV2 `json:"eriDiscoveryConfig,omitempty"`

	// The DataSource data collect interval
	// Example: 10
	EriDiscoveryInterval int32 `json:"eriDiscoveryInterval,omitempty"`

	// The group the LMModule is in
	// Example: string
	Group string `json:"group,omitempty"`

	// If the DataSource has multi instance: true|false
	HasMultiInstances bool `json:"hasMultiInstances,omitempty"`

	// The ID of the LMModule
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The lineageId the LMModule belongs to
	// Read Only: true
	LineageID string `json:"lineageId,omitempty"`

	// The data source name
	// Example: datasource test
	// Required: true
	Name *string `json:"name"`

	// The DataSource payload version for custom metrics
	// Read Only: true
	PayloadVersion int32 `json:"payloadVersion,omitempty"`

	// The Tags for the LMModule
	// Example: string
	Tags string `json:"tags,omitempty"`

	// The Technical Notes for the LMModule
	// Example: string
	Technology string `json:"technology,omitempty"`

	// Use wild-value as unique identifier in case of multi instance datasource: true|false
	// Read Only: true
	UseWildValueAsUUID *bool `json:"useWildValueAsUUID,omitempty"`

	// The data source version
	// Read Only: true
	Version int64 `json:"version,omitempty"`
}

// Validate validates this datasource
func (m *Datasource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDiscoveryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectorAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEriDiscoveryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datasource) validateAutoDiscoveryConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoDiscoveryConfig) { // not required
		return nil
	}

	if m.AutoDiscoveryConfig != nil {
		if err := m.AutoDiscoveryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiscoveryConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) validateCollectInterval(formats strfmt.Registry) error {

	if err := validate.Required("collectInterval", "body", m.CollectInterval); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) validateCollectMethod(formats strfmt.Registry) error {

	if err := validate.Required("collectMethod", "body", m.CollectMethod); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) validateCollectorAttribute(formats strfmt.Registry) error {

	if err := validate.Required("collectorAttribute", "body", m.CollectorAttribute); err != nil {
		return err
	}

	if m.CollectorAttribute != nil {
		if err := m.CollectorAttribute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectorAttribute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectorAttribute")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) validateDataPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.DataPoints); i++ {
		if swag.IsZero(m.DataPoints[i]) { // not required
			continue
		}

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datasource) validateEriDiscoveryConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EriDiscoveryConfig) { // not required
		return nil
	}

	if m.EriDiscoveryConfig != nil {
		if err := m.EriDiscoveryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eriDiscoveryConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eriDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource based on the context it is used
func (m *Datasource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoDiscoveryConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecksum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorAttribute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEriDiscoveryConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineageID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayloadVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUseWildValueAsUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datasource) contextValidateAuditVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "auditVersion", "body", int64(m.AuditVersion)); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidateAutoDiscoveryConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoDiscoveryConfig != nil {
		if err := m.AutoDiscoveryConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiscoveryConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) contextValidateChecksum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "checksum", "body", string(m.Checksum)); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidateCollectorAttribute(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectorAttribute != nil {
		if err := m.CollectorAttribute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectorAttribute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectorAttribute")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) contextValidateDataPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataPoints); i++ {

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datasource) contextValidateEriDiscoveryConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EriDiscoveryConfig != nil {
		if err := m.EriDiscoveryConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eriDiscoveryConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eriDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *Datasource) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidateLineageID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lineageId", "body", string(m.LineageID)); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidatePayloadVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "payloadVersion", "body", int32(m.PayloadVersion)); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidateUseWildValueAsUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "useWildValueAsUUID", "body", m.UseWildValueAsUUID); err != nil {
		return err
	}

	return nil
}

func (m *Datasource) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", int64(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Datasource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Datasource) UnmarshalBinary(b []byte) error {
	var res Datasource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
