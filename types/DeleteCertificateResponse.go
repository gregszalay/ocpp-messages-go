// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DeleteCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DeleteCertificateStatusEnumType, v)
	}
	*j = DeleteCertificateStatusEnumType(v)
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
func (j *DeleteCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DeleteCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DeleteCertificateStatusEnumType_1, v)
	}
	*j = DeleteCertificateStatusEnumType_1(v)
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

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

type DeleteCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status DeleteCertificateStatusEnumType_1 `json:"status" yaml:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty" yaml:"statusInfo,omitempty"`
}

type DeleteCertificateStatusEnumType string

const DeleteCertificateStatusEnumTypeAccepted DeleteCertificateStatusEnumType = "Accepted"
const DeleteCertificateStatusEnumTypeFailed DeleteCertificateStatusEnumType = "Failed"
const DeleteCertificateStatusEnumTypeNotFound DeleteCertificateStatusEnumType = "NotFound"

type DeleteCertificateStatusEnumType_1 string

const DeleteCertificateStatusEnumType_1_Accepted DeleteCertificateStatusEnumType_1 = "Accepted"
const DeleteCertificateStatusEnumType_1_Failed DeleteCertificateStatusEnumType_1 = "Failed"
const DeleteCertificateStatusEnumType_1_NotFound DeleteCertificateStatusEnumType_1 = "NotFound"

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

var enumValues_DeleteCertificateStatusEnumType = []interface{}{
	"Accepted",
	"Failed",
	"NotFound",
}
var enumValues_DeleteCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"Failed",
	"NotFound",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status in DeleteCertificateResponseJson: required")
	}
	type Plain DeleteCertificateResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DeleteCertificateResponseJson(plain)
	return nil
}