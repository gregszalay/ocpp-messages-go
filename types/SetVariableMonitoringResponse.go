// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetMonitoringStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetMonitoringStatusEnumType, v)
	}
	*j = SetMonitoringStatusEnumType(v)
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

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetMonitoringStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetMonitoringStatusEnumType_1, v)
	}
	*j = SetMonitoringStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType, v)
	}
	*j = MonitorEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component in SetMonitoringResultType: required")
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity in SetMonitoringResultType: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status in SetMonitoringResultType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in SetMonitoringResultType: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable in SetMonitoringResultType: required")
	}
	type Plain SetMonitoringResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringResultType(plain)
	return nil
}

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
func (j *MonitorEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_1, v)
	}
	*j = MonitorEnumType_1(v)
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

type MonitorEnumType string

const MonitorEnumTypeDelta MonitorEnumType = "Delta"
const MonitorEnumTypeLowerThreshold MonitorEnumType = "LowerThreshold"
const MonitorEnumTypePeriodic MonitorEnumType = "Periodic"
const MonitorEnumTypePeriodicClockAligned MonitorEnumType = "PeriodicClockAligned"
const MonitorEnumTypeUpperThreshold MonitorEnumType = "UpperThreshold"

type MonitorEnumType_1 string

const MonitorEnumType_1_Delta MonitorEnumType_1 = "Delta"
const MonitorEnumType_1_LowerThreshold MonitorEnumType_1 = "LowerThreshold"
const MonitorEnumType_1_Periodic MonitorEnumType_1 = "Periodic"
const MonitorEnumType_1_PeriodicClockAligned MonitorEnumType_1 = "PeriodicClockAligned"
const MonitorEnumType_1_UpperThreshold MonitorEnumType_1 = "UpperThreshold"

// Class to hold result of SetVariableMonitoring request.
//
type SetMonitoringResultType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component" yaml:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Id given to the VariableMonitor by the Charging Station. The Id is only
	// returned when status is accepted. Installed VariableMonitors should have unique
	// id's but the id's of removed Installed monitors should have unique id's but the
	// id's of removed monitors MAY be reused.
	//
	Id *int `json:"id,omitempty" yaml:"id,omitempty"`

	// The severity that will be assigned to an event that is triggered by this
	// monitor. The severity range is 0-9, with 0 as the highest and 9 as the lowest
	// severity level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	//
	Severity int `json:"severity" yaml:"severity"`

	// Status corresponds to the JSON schema field "status".
	Status SetMonitoringStatusEnumType `json:"status" yaml:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty" yaml:"statusInfo,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type MonitorEnumType_1 `json:"type" yaml:"type"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType `json:"variable" yaml:"variable"`
}

type SetMonitoringStatusEnumType string

const SetMonitoringStatusEnumTypeAccepted SetMonitoringStatusEnumType = "Accepted"
const SetMonitoringStatusEnumTypeDuplicate SetMonitoringStatusEnumType = "Duplicate"
const SetMonitoringStatusEnumTypeRejected SetMonitoringStatusEnumType = "Rejected"
const SetMonitoringStatusEnumTypeUnknownComponent SetMonitoringStatusEnumType = "UnknownComponent"
const SetMonitoringStatusEnumTypeUnknownVariable SetMonitoringStatusEnumType = "UnknownVariable"
const SetMonitoringStatusEnumTypeUnsupportedMonitorType SetMonitoringStatusEnumType = "UnsupportedMonitorType"

type SetMonitoringStatusEnumType_1 string

const SetMonitoringStatusEnumType_1_Accepted SetMonitoringStatusEnumType_1 = "Accepted"
const SetMonitoringStatusEnumType_1_Duplicate SetMonitoringStatusEnumType_1 = "Duplicate"
const SetMonitoringStatusEnumType_1_Rejected SetMonitoringStatusEnumType_1 = "Rejected"
const SetMonitoringStatusEnumType_1_UnknownComponent SetMonitoringStatusEnumType_1 = "UnknownComponent"
const SetMonitoringStatusEnumType_1_UnknownVariable SetMonitoringStatusEnumType_1 = "UnknownVariable"
const SetMonitoringStatusEnumType_1_UnsupportedMonitorType SetMonitoringStatusEnumType_1 = "UnsupportedMonitorType"

type SetVariableMonitoringResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// SetMonitoringResult corresponds to the JSON schema field "setMonitoringResult".
	SetMonitoringResult []SetMonitoringResultType `json:"setMonitoringResult" yaml:"setMonitoringResult"`
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

var enumValues_MonitorEnumType = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_1 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_SetMonitoringStatusEnumType = []interface{}{
	"Accepted",
	"UnknownComponent",
	"UnknownVariable",
	"UnsupportedMonitorType",
	"Rejected",
	"Duplicate",
}
var enumValues_SetMonitoringStatusEnumType_1 = []interface{}{
	"Accepted",
	"UnknownComponent",
	"UnknownVariable",
	"UnsupportedMonitorType",
	"Rejected",
	"Duplicate",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableMonitoringResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setMonitoringResult"]; !ok || v == nil {
		return fmt.Errorf("field setMonitoringResult in SetVariableMonitoringResponseJson: required")
	}
	type Plain SetVariableMonitoringResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.SetMonitoringResult) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "setMonitoringResult", 1)
	}
	*j = SetVariableMonitoringResponseJson(plain)
	return nil
}
