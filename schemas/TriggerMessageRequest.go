// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
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

type MessageTriggerEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageTriggerEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageTriggerEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageTriggerEnumType_1, v)
	}
	*j = MessageTriggerEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageTriggerEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageTriggerEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageTriggerEnumType, v)
	}
	*j = MessageTriggerEnumType(v)
	return nil
}

const MessageTriggerEnumTypeBootNotification MessageTriggerEnumType = "BootNotification"
const MessageTriggerEnumTypeFirmwareStatusNotification MessageTriggerEnumType = "FirmwareStatusNotification"
const MessageTriggerEnumTypeHeartbeat MessageTriggerEnumType = "Heartbeat"
const MessageTriggerEnumTypeLogStatusNotification MessageTriggerEnumType = "LogStatusNotification"
const MessageTriggerEnumTypeMeterValues MessageTriggerEnumType = "MeterValues"
const MessageTriggerEnumTypePublishFirmwareStatusNotification MessageTriggerEnumType = "PublishFirmwareStatusNotification"
const MessageTriggerEnumTypeSignChargingStationCertificate MessageTriggerEnumType = "SignChargingStationCertificate"
const MessageTriggerEnumTypeSignCombinedCertificate MessageTriggerEnumType = "SignCombinedCertificate"
const MessageTriggerEnumTypeSignV2GCertificate MessageTriggerEnumType = "SignV2GCertificate"
const MessageTriggerEnumTypeStatusNotification MessageTriggerEnumType = "StatusNotification"
const MessageTriggerEnumTypeTransactionEvent MessageTriggerEnumType = "TransactionEvent"

type MessageTriggerEnumType_1 string

const MessageTriggerEnumType_1_BootNotification MessageTriggerEnumType_1 = "BootNotification"
const MessageTriggerEnumType_1_FirmwareStatusNotification MessageTriggerEnumType_1 = "FirmwareStatusNotification"
const MessageTriggerEnumType_1_Heartbeat MessageTriggerEnumType_1 = "Heartbeat"
const MessageTriggerEnumType_1_LogStatusNotification MessageTriggerEnumType_1 = "LogStatusNotification"
const MessageTriggerEnumType_1_MeterValues MessageTriggerEnumType_1 = "MeterValues"
const MessageTriggerEnumType_1_PublishFirmwareStatusNotification MessageTriggerEnumType_1 = "PublishFirmwareStatusNotification"
const MessageTriggerEnumType_1_SignChargingStationCertificate MessageTriggerEnumType_1 = "SignChargingStationCertificate"
const MessageTriggerEnumType_1_SignCombinedCertificate MessageTriggerEnumType_1 = "SignCombinedCertificate"
const MessageTriggerEnumType_1_SignV2GCertificate MessageTriggerEnumType_1 = "SignV2GCertificate"
const MessageTriggerEnumType_1_StatusNotification MessageTriggerEnumType_1 = "StatusNotification"
const MessageTriggerEnumType_1_TransactionEvent MessageTriggerEnumType_1 = "TransactionEvent"

type TriggerMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty"`

	// RequestedMessage corresponds to the JSON schema field "requestedMessage".
	RequestedMessage MessageTriggerEnumType_1 `json:"requestedMessage" yaml:"requestedMessage"`
}

var enumValues_MessageTriggerEnumType = []interface{}{
	"BootNotification",
	"LogStatusNotification",
	"FirmwareStatusNotification",
	"Heartbeat",
	"MeterValues",
	"SignChargingStationCertificate",
	"SignV2GCertificate",
	"StatusNotification",
	"TransactionEvent",
	"SignCombinedCertificate",
	"PublishFirmwareStatusNotification",
}
var enumValues_MessageTriggerEnumType_1 = []interface{}{
	"BootNotification",
	"LogStatusNotification",
	"FirmwareStatusNotification",
	"Heartbeat",
	"MeterValues",
	"SignChargingStationCertificate",
	"SignV2GCertificate",
	"StatusNotification",
	"TransactionEvent",
	"SignCombinedCertificate",
	"PublishFirmwareStatusNotification",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestedMessage"]; !ok || v == nil {
		return fmt.Errorf("field requestedMessage in TriggerMessageRequestJson: required")
	}
	type Plain TriggerMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TriggerMessageRequestJson(plain)
	return nil
}
