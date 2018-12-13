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
	"github.com/go-openapi/validate"
)

// SDKScriptCollectorAttribute s d k script collector attribute
// swagger:model SDKScriptCollectorAttribute
type SDKScriptCollectorAttribute struct {

	// groovy script
	// Required: true
	GroovyScript *string `json:"groovyScript"`

	// sdk name
	// Required: true
	SdkName *string `json:"sdkName"`

	// sdk version
	// Required: true
	SdkVersion *string `json:"sdkVersion"`
}

// Name gets the name of this subtype
func (m *SDKScriptCollectorAttribute) Name() string {
	return "sdkscript"
}

// SetName sets the name of this subtype
func (m *SDKScriptCollectorAttribute) SetName(val string) {

}

// GroovyScript gets the groovy script of this subtype

// SdkName gets the sdk name of this subtype

// SdkVersion gets the sdk version of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SDKScriptCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// groovy script
		// Required: true
		GroovyScript *string `json:"groovyScript"`

		// sdk name
		// Required: true
		SdkName *string `json:"sdkName"`

		// sdk version
		// Required: true
		SdkVersion *string `json:"sdkVersion"`
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

	var result SDKScriptCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.GroovyScript = data.GroovyScript

	result.SdkName = data.SdkName

	result.SdkVersion = data.SdkVersion

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SDKScriptCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// groovy script
		// Required: true
		GroovyScript *string `json:"groovyScript"`

		// sdk name
		// Required: true
		SdkName *string `json:"sdkName"`

		// sdk version
		// Required: true
		SdkVersion *string `json:"sdkVersion"`
	}{

		GroovyScript: m.GroovyScript,

		SdkName: m.SdkName,

		SdkVersion: m.SdkVersion,
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

// Validate validates this s d k script collector attribute
func (m *SDKScriptCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroovyScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSdkName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSdkVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDKScriptCollectorAttribute) validateGroovyScript(formats strfmt.Registry) error {

	if err := validate.Required("groovyScript", "body", m.GroovyScript); err != nil {
		return err
	}

	return nil
}

func (m *SDKScriptCollectorAttribute) validateSdkName(formats strfmt.Registry) error {

	if err := validate.Required("sdkName", "body", m.SdkName); err != nil {
		return err
	}

	return nil
}

func (m *SDKScriptCollectorAttribute) validateSdkVersion(formats strfmt.Registry) error {

	if err := validate.Required("sdkVersion", "body", m.SdkVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDKScriptCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDKScriptCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res SDKScriptCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
