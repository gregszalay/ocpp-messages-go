// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package ChangeAvailabilityResponse

import "fmt"
import "reflect"
import "encoding/json"

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

type ChangeAvailabilityResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ChangeAvailabilityStatusEnumType_1 `json:"status" yaml:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty" yaml:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChangeAvailabilityStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChangeAvailabilityStatusEnumType, v)
	}
	*j = ChangeAvailabilityStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChangeAvailabilityStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChangeAvailabilityStatusEnumType_1, v)
	}
	*j = ChangeAvailabilityStatusEnumType_1(v)
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

type ChangeAvailabilityStatusEnumType string

const ChangeAvailabilityStatusEnumTypeAccepted ChangeAvailabilityStatusEnumType = "Accepted"
const ChangeAvailabilityStatusEnumTypeRejected ChangeAvailabilityStatusEnumType = "Rejected"
const ChangeAvailabilityStatusEnumTypeScheduled ChangeAvailabilityStatusEnumType = "Scheduled"

type ChangeAvailabilityStatusEnumType_1 string

const ChangeAvailabilityStatusEnumType_1_Accepted ChangeAvailabilityStatusEnumType_1 = "Accepted"
const ChangeAvailabilityStatusEnumType_1_Rejected ChangeAvailabilityStatusEnumType_1 = "Rejected"
const ChangeAvailabilityStatusEnumType_1_Scheduled ChangeAvailabilityStatusEnumType_1 = "Scheduled"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
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

var enumValues_ChangeAvailabilityStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}
var enumValues_ChangeAvailabilityStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status in ChangeAvailabilityResponseJson: required")
	}
	type Plain ChangeAvailabilityResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChangeAvailabilityResponseJson(plain)
	return nil
}
