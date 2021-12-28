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

// Device device
//
// swagger:model Device
type Device struct {

	// The Auto Balanced Collector Group id. 0 means not monitored by ABCG
	// Example: 0
	AutoBalancedCollectorGroupID int32 `json:"autoBalancedCollectorGroupId,omitempty"`

	// Any auto properties assigned to the device
	// Read Only: true
	AutoProperties []*NameAndValue `json:"autoProperties"`

	// The time, in epoch seconds format, that properties were first discovered for this device
	// Read Only: true
	AutoPropsAssignedOn int64 `json:"autoPropsAssignedOn,omitempty"`

	// The time, in epoch seconds, that auto properties last ran and updated the properties table for this device
	// Read Only: true
	AutoPropsUpdatedOn int64 `json:"autoPropsUpdatedOn,omitempty"`

	// The AWS instance state (if applicable): 1 indicates that the instance is running, 2 indicates that the instance is stopped and 3 the instance is terminated
	// Read Only: true
	AwsState int32 `json:"awsState,omitempty"`

	// The GCP instance state (if applicable): 1 indicates that the instance is running, 2 indicates that the instance is stopped and 3 the instance is terminated.
	// Read Only: true
	AzureState int32 `json:"azureState,omitempty"`

	// The description/name of the collector for this device
	// Read Only: true
	CollectorDescription string `json:"collectorDescription,omitempty"`

	// The time, in epoch seconds format, that the device was added to your LogicMonitor account
	// Read Only: true
	CreatedOn int64 `json:"createdOn,omitempty"`

	// The id of the collector currently monitoring the device and discovering instances
	// Example: 1
	CurrentCollectorID int32 `json:"currentCollectorId,omitempty"`

	// Any non-system properties (aside from system.categories) defined for this device
	CustomProperties []*NameAndValue `json:"customProperties"`

	// The time in milliseconds that the device has been dead for, or since the AWS device was filtered out
	// Read Only: true
	DeletedTimeInMs int64 `json:"deletedTimeInMs,omitempty"`

	// The device description
	// Example: This is a Cisco Router
	Description string `json:"description,omitempty"`

	// The type of device: 0 indicates a regular device, 2 indicates an AWS device, 4 indicates an Azure device
	// Example: 0
	DeviceType int32 `json:"deviceType,omitempty"`

	// Indicates whether alerting is disabled (true) or enabled (false) for this device
	// Example: true
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// The display name of the device
	// Example: Cisco Router
	// Required: true
	DisplayName *string `json:"displayName"`

	// Indicates whether Netflow is enabled (true) or disabled (false) for the device
	// Example: true
	EnableNetflow bool `json:"enableNetflow,omitempty"`

	// The Azure instance state (if applicable): 1 indicates that the instance is running, 2 indicates that the instance is stopped and 3 the instance is terminated.
	// Read Only: true
	GcpState int32 `json:"gcpState,omitempty"`

	// The Id(s) of the groups the device is in, where multiple group ids are comma separated
	// Example: 16,4,3
	HostGroupIds string `json:"hostGroupIds,omitempty"`

	// The status of this device, where possible statuses are normal, dead and dead-collector
	// Read Only: true
	HostStatus string `json:"hostStatus,omitempty"`

	// The Id of the device
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// Any properties inherit from parents
	// Read Only: true
	InheritedProperties []*NameAndValue `json:"inheritedProperties"`

	// The last time, in epoch seconds, that the device received Netflow data
	// Read Only: true
	LastDataTime int64 `json:"lastDataTime,omitempty"`

	// The last time, in epoch seconds, that raw Netflow data was reported
	// Read Only: true
	LastRawdataTime int64 `json:"lastRawdataTime,omitempty"`

	// The URL link associated with the device
	// Example: www.ciscorouter.com
	Link string `json:"link,omitempty"`

	// The host name or IP address of the device
	// Example: collector.localhost
	// Required: true
	Name *string `json:"name"`

	// The description/name of the netflow collector for this device
	// Read Only: true
	NetflowCollectorDescription string `json:"netflowCollectorDescription,omitempty"`

	// The id of the Collector Group associated with the device's netflow collector
	// Read Only: true
	NetflowCollectorGroupID int32 `json:"netflowCollectorGroupId,omitempty"`

	// The name of the Collector Group associated with the device's netflow collector
	// Read Only: true
	NetflowCollectorGroupName string `json:"netflowCollectorGroupName,omitempty"`

	// The Id of the netflow collector associated with the device
	// Example: 1
	NetflowCollectorID int32 `json:"netflowCollectorId,omitempty"`

	// The id of the Collector Group associated with the device's preferred collector
	// Read Only: true
	PreferredCollectorGroupID int32 `json:"preferredCollectorGroupId,omitempty"`

	// The name of the Collector Group associated with the device's preferred collector
	// Read Only: true
	PreferredCollectorGroupName string `json:"preferredCollectorGroupName,omitempty"`

	// The Id of the preferred collector assigned to monitor the device
	// Example: 2
	// Required: true
	PreferredCollectorID *int32 `json:"preferredCollectorId"`

	// The Id of the AWS EC2 instance related to this device, if one exists in the LogicMonitor account. This value defaults to -1, which indicates that there are no related devices
	// Example: -1
	RelatedDeviceID int32 `json:"relatedDeviceId,omitempty"`

	// The Id of the netscan configuration which was used to discover this device. 0 indicates that the device was not discovered by a scan
	// Example: 0
	// Read Only: true
	ScanConfigID int32 `json:"scanConfigId,omitempty"`

	// Any system properties (aside from system.categories) defined for this device
	// Read Only: true
	SystemProperties []*NameAndValue `json:"systemProperties"`

	// The number of milliseconds until the device will be automatically deleted from your LogicMonitor account (a value of zero indicates that a future delete time/date has not been scheduled)
	// Read Only: true
	ToDeleteTimeInMs int64 `json:"toDeleteTimeInMs,omitempty"`

	// The uptime of the device in seconds. This value will always be the largest value reported by the following datasources:
	// Host Uptime-
	// SNMPUptime-
	// SNMP_Engine_Uptime-
	// WinSystemUptime-
	// NimbleUptime-
	// Read Only: true
	UpTimeInSeconds int64 `json:"upTimeInSeconds,omitempty"`

	// The time, in epoch seconds format, that the device was last updated
	// Read Only: true
	UpdatedOn int64 `json:"updatedOn,omitempty"`

	// The read and/or write permissions for this device that are granted to the user who made the API request
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritedProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredCollectorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateAutoProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.AutoProperties); i++ {
		if swag.IsZero(m.AutoProperties[i]) { // not required
			continue
		}

		if m.AutoProperties[i] != nil {
			if err := m.AutoProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateCustomProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomProperties); i++ {
		if swag.IsZero(m.CustomProperties[i]) { // not required
			continue
		}

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateInheritedProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritedProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.InheritedProperties); i++ {
		if swag.IsZero(m.InheritedProperties[i]) { // not required
			continue
		}

		if m.InheritedProperties[i] != nil {
			if err := m.InheritedProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Device) validatePreferredCollectorID(formats strfmt.Registry) error {

	if err := validate.Required("preferredCollectorId", "body", m.PreferredCollectorID); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateSystemProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemProperties); i++ {
		if swag.IsZero(m.SystemProperties[i]) { // not required
			continue
		}

		if m.SystemProperties[i] != nil {
			if err := m.SystemProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this device based on the context it is used
func (m *Device) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoPropsAssignedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoPropsUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectorDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletedTimeInMs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritedProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastDataTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastRawdataTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetflowCollectorDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetflowCollectorGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetflowCollectorGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredCollectorGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredCollectorGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanConfigID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToDeleteTimeInMs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpTimeInSeconds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedOn(ctx, formats); err != nil {
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

func (m *Device) contextValidateAutoProperties(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "autoProperties", "body", []*NameAndValue(m.AutoProperties)); err != nil {
		return err
	}

	for i := 0; i < len(m.AutoProperties); i++ {

		if m.AutoProperties[i] != nil {
			if err := m.AutoProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("autoProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) contextValidateAutoPropsAssignedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "autoPropsAssignedOn", "body", int64(m.AutoPropsAssignedOn)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateAutoPropsUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "autoPropsUpdatedOn", "body", int64(m.AutoPropsUpdatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateAwsState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "awsState", "body", int32(m.AwsState)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateAzureState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "azureState", "body", int32(m.AzureState)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateCollectorDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectorDescription", "body", string(m.CollectorDescription)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateCreatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdOn", "body", int64(m.CreatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateCustomProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomProperties); i++ {

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) contextValidateDeletedTimeInMs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deletedTimeInMs", "body", int64(m.DeletedTimeInMs)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateGcpState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "gcpState", "body", int32(m.GcpState)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateHostStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hostStatus", "body", string(m.HostStatus)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateInheritedProperties(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inheritedProperties", "body", []*NameAndValue(m.InheritedProperties)); err != nil {
		return err
	}

	for i := 0; i < len(m.InheritedProperties); i++ {

		if m.InheritedProperties[i] != nil {
			if err := m.InheritedProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) contextValidateLastDataTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastDataTime", "body", int64(m.LastDataTime)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateLastRawdataTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastRawdataTime", "body", int64(m.LastRawdataTime)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateNetflowCollectorDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "netflowCollectorDescription", "body", string(m.NetflowCollectorDescription)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateNetflowCollectorGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "netflowCollectorGroupId", "body", int32(m.NetflowCollectorGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateNetflowCollectorGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "netflowCollectorGroupName", "body", string(m.NetflowCollectorGroupName)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidatePreferredCollectorGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "preferredCollectorGroupId", "body", int32(m.PreferredCollectorGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidatePreferredCollectorGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "preferredCollectorGroupName", "body", string(m.PreferredCollectorGroupName)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateScanConfigID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scanConfigId", "body", int32(m.ScanConfigID)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateSystemProperties(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "systemProperties", "body", []*NameAndValue(m.SystemProperties)); err != nil {
		return err
	}

	for i := 0; i < len(m.SystemProperties); i++ {

		if m.SystemProperties[i] != nil {
			if err := m.SystemProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) contextValidateToDeleteTimeInMs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "toDeleteTimeInMs", "body", int64(m.ToDeleteTimeInMs)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateUpTimeInSeconds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "upTimeInSeconds", "body", int64(m.UpTimeInSeconds)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedOn", "body", int64(m.UpdatedOn)); err != nil {
		return err
	}

	return nil
}

func (m *Device) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}