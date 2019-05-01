// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DataPumpCollectorAttribute data pump collector attribute
// swagger:model DataPumpCollectorAttribute
type DataPumpCollectorAttribute struct {
	DataPumpCollectorAttributeAllOf1
}

// Name gets the name of this subtype
func (m *DataPumpCollectorAttribute) Name() string {
	return "datapump"
}

// SetName sets the name of this subtype
func (m *DataPumpCollectorAttribute) SetName(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DataPumpCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {
		DataPumpCollectorAttributeAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DataPumpCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.DataPumpCollectorAttributeAllOf1 = data.DataPumpCollectorAttributeAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DataPumpCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		DataPumpCollectorAttributeAllOf1
	}{

		DataPumpCollectorAttributeAllOf1: m.DataPumpCollectorAttributeAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this data pump collector attribute
func (m *DataPumpCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DataPumpCollectorAttributeAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DataPumpCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPumpCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res DataPumpCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataPumpCollectorAttributeAllOf1 data pump collector attribute all of1
// swagger:model DataPumpCollectorAttributeAllOf1
type DataPumpCollectorAttributeAllOf1 interface{}
