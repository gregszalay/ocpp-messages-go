// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package SetVariablesResponse

import "fmt"
import "reflect"
import "encoding/json"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetVariableStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetVariableStatusEnumType, v)
	}
	*j = SetVariableStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetVariableStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetVariableStatusEnumType_1, v)
	}
	*j = SetVariableStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType, v)
	}
	*j = AttributeEnumType(v)
	return nil
}

const AttributeEnumTypeActual AttributeEnumType = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["attributeStatus"]; !ok || v == nil {
		return fmt.Errorf("field attributeStatus in SetVariableResultType: required")
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component in SetVariableResultType: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable in SetVariableResultType: required")
	}
	type Plain SetVariableResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariableResultType(plain)
	return nil
}

type AttributeEnumType string

const AttributeEnumTypeMaxSet AttributeEnumType = "MaxSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in VariableType: required")
	}
	type Plain VariableType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId in CustomDataType: required")
	}
	type Plain CustomDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_1, v)
	}
	*j = AttributeEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in EVSEType: required")
	}
	type Plain EVSEType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode in StatusInfoType: required")
	}
	type Plain StatusInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in ComponentType: required")
	}
	type Plain ComponentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType(plain)
	return nil
}

const AttributeEnumTypeMinSet AttributeEnumType = "MinSet"
const AttributeEnumTypeTarget AttributeEnumType = "Target"

type AttributeEnumType_1 string

const AttributeEnumType_1_Actual AttributeEnumType_1 = "Actual"
const AttributeEnumType_1_MaxSet AttributeEnumType_1 = "MaxSet"
const AttributeEnumType_1_MinSet AttributeEnumType_1 = "MinSet"
const AttributeEnumType_1_Target AttributeEnumType_1 = "Target"

// A physical or logical component
//
type ComponentType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty" yaml:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id" yaml:"id"`
}

type SetVariableResultType struct {
	// AttributeStatus corresponds to the JSON schema field "attributeStatus".
	AttributeStatus SetVariableStatusEnumType `json:"attributeStatus" yaml:"attributeStatus"`

	// AttributeStatusInfo corresponds to the JSON schema field "attributeStatusInfo".
	AttributeStatusInfo *StatusInfoType `json:"attributeStatusInfo,omitempty" yaml:"attributeStatusInfo,omitempty"`

	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType_1 `json:"attributeType,omitempty" yaml:"attributeType,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component" yaml:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType `json:"variable" yaml:"variable"`
}

type SetVariableStatusEnumType string

const SetVariableStatusEnumTypeAccepted SetVariableStatusEnumType = "Accepted"
const SetVariableStatusEnumTypeNotSupportedAttributeType SetVariableStatusEnumType = "NotSupportedAttributeType"
const SetVariableStatusEnumTypeRebootRequired SetVariableStatusEnumType = "RebootRequired"
const SetVariableStatusEnumTypeRejected SetVariableStatusEnumType = "Rejected"
const SetVariableStatusEnumTypeUnknownComponent SetVariableStatusEnumType = "UnknownComponent"
const SetVariableStatusEnumTypeUnknownVariable SetVariableStatusEnumType = "UnknownVariable"

type SetVariableStatusEnumType_1 string

const SetVariableStatusEnumType_1_Accepted SetVariableStatusEnumType_1 = "Accepted"
const SetVariableStatusEnumType_1_NotSupportedAttributeType SetVariableStatusEnumType_1 = "NotSupportedAttributeType"
const SetVariableStatusEnumType_1_RebootRequired SetVariableStatusEnumType_1 = "RebootRequired"
const SetVariableStatusEnumType_1_Rejected SetVariableStatusEnumType_1 = "Rejected"
const SetVariableStatusEnumType_1_UnknownComponent SetVariableStatusEnumType_1 = "UnknownComponent"
const SetVariableStatusEnumType_1_UnknownVariable SetVariableStatusEnumType_1 = "UnknownVariable"

type SetVariablesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// SetVariableResult corresponds to the JSON schema field "setVariableResult".
	SetVariableResult []SetVariableResultType `json:"setVariableResult" yaml:"setVariableResult"`
}

// Element providing more information about the status.
//
type StatusInfoType struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode" yaml:"reasonCode"`
}

// Reference key to a component-variable.
//
type VariableType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name"`
}

var enumValues_AttributeEnumType = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_1 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_SetVariableStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
	"RebootRequired",
}
var enumValues_SetVariableStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
	"RebootRequired",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariablesResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setVariableResult"]; !ok || v == nil {
		return fmt.Errorf("field setVariableResult in SetVariablesResponseJson: required")
	}
	type Plain SetVariablesResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.SetVariableResult) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "setVariableResult", 1)
	}
	*j = SetVariablesResponseJson(plain)
	return nil
}
