// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "reflect"
import "encoding/json"

type BootNotificationRequestJson struct {
	// ChargingStation corresponds to the JSON schema field "chargingStation".
	ChargingStation ChargingStationType `json:"chargingStation" yaml:"chargingStation"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Reason corresponds to the JSON schema field "reason".
	Reason BootReasonEnumType_1 `json:"reason" yaml:"reason"`
}

type BootReasonEnumType string

const BootReasonEnumTypeApplicationReset BootReasonEnumType = "ApplicationReset"
const BootReasonEnumTypeFirmwareUpdate BootReasonEnumType = "FirmwareUpdate"
const BootReasonEnumTypeLocalReset BootReasonEnumType = "LocalReset"
const BootReasonEnumTypePowerUp BootReasonEnumType = "PowerUp"
const BootReasonEnumTypeRemoteReset BootReasonEnumType = "RemoteReset"
const BootReasonEnumTypeScheduledReset BootReasonEnumType = "ScheduledReset"
const BootReasonEnumTypeTriggered BootReasonEnumType = "Triggered"
const BootReasonEnumTypeUnknown BootReasonEnumType = "Unknown"
const BootReasonEnumTypeWatchdog BootReasonEnumType = "Watchdog"

type BootReasonEnumType_1 string

const BootReasonEnumType_1_ApplicationReset BootReasonEnumType_1 = "ApplicationReset"
const BootReasonEnumType_1_FirmwareUpdate BootReasonEnumType_1 = "FirmwareUpdate"
const BootReasonEnumType_1_LocalReset BootReasonEnumType_1 = "LocalReset"
const BootReasonEnumType_1_PowerUp BootReasonEnumType_1 = "PowerUp"
const BootReasonEnumType_1_RemoteReset BootReasonEnumType_1 = "RemoteReset"
const BootReasonEnumType_1_ScheduledReset BootReasonEnumType_1 = "ScheduledReset"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingStationType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["model"]; !ok || v == nil {
		return fmt.Errorf("field model in ChargingStationType: required")
	}
	if v, ok := raw["vendorName"]; !ok || v == nil {
		return fmt.Errorf("field vendorName in ChargingStationType: required")
	}
	type Plain ChargingStationType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingStationType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootReasonEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BootReasonEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BootReasonEnumType_1, v)
	}
	*j = BootReasonEnumType_1(v)
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
func (j *BootReasonEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BootReasonEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BootReasonEnumType, v)
	}
	*j = BootReasonEnumType(v)
	return nil
}

const BootReasonEnumType_1_Triggered BootReasonEnumType_1 = "Triggered"
const BootReasonEnumType_1_Unknown BootReasonEnumType_1 = "Unknown"
const BootReasonEnumType_1_Watchdog BootReasonEnumType_1 = "Watchdog"

// Charge_ Point
// urn:x-oca:ocpp:uid:2:233122
// The physical system where an Electrical Vehicle (EV) can be charged.
//
type ChargingStationType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// This contains the firmware version of the Charging Station.
	//
	//
	FirmwareVersion *string `json:"firmwareVersion,omitempty" yaml:"firmwareVersion,omitempty"`

	// Device. Model. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569325
	// Defines the model of the device.
	//
	Model string `json:"model" yaml:"model"`

	// Modem corresponds to the JSON schema field "modem".
	Modem *ModemType `json:"modem,omitempty" yaml:"modem,omitempty"`

	// Device. Serial_ Number. Serial_ Number
	// urn:x-oca:ocpp:uid:1:569324
	// Vendor-specific device identifier.
	//
	SerialNumber *string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`

	// Identifies the vendor (not necessarily in a unique manner).
	//
	VendorName string `json:"vendorName" yaml:"vendorName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingStation"]; !ok || v == nil {
		return fmt.Errorf("field chargingStation in BootNotificationRequestJson: required")
	}
	if v, ok := raw["reason"]; !ok || v == nil {
		return fmt.Errorf("field reason in BootNotificationRequestJson: required")
	}
	type Plain BootNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BootNotificationRequestJson(plain)
	return nil
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

// Wireless_ Communication_ Module
// urn:x-oca:ocpp:uid:2:233306
// Defines parameters required for initiating and maintaining wireless
// communication with other devices.
//
type ModemType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Wireless_ Communication_ Module. ICCID. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569327
	// This contains the ICCID of the modem’s SIM card.
	//
	Iccid *string `json:"iccid,omitempty" yaml:"iccid,omitempty"`

	// Wireless_ Communication_ Module. IMSI. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569328
	// This contains the IMSI of the modem’s SIM card.
	//
	Imsi *string `json:"imsi,omitempty" yaml:"imsi,omitempty"`
}

var enumValues_BootReasonEnumType = []interface{}{
	"ApplicationReset",
	"FirmwareUpdate",
	"LocalReset",
	"PowerUp",
	"RemoteReset",
	"ScheduledReset",
	"Triggered",
	"Unknown",
	"Watchdog",
}
var enumValues_BootReasonEnumType_1 = []interface{}{
	"ApplicationReset",
	"FirmwareUpdate",
	"LocalReset",
	"PowerUp",
	"RemoteReset",
	"ScheduledReset",
	"Triggered",
	"Unknown",
	"Watchdog",
}
