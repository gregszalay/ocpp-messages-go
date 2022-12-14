// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package UpdateFirmwareRequest

import "fmt"
import "encoding/json"

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

// Firmware
// urn:x-enexis:ecdm:uid:2:233291
// Represents a copy of the firmware that can be loaded/updated on the Charging
// Station.
//
type FirmwareType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Firmware. Install. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569462
	// Date and time at which the firmware shall be installed.
	//
	InstallDateTime *string `json:"installDateTime,omitempty" yaml:"installDateTime,omitempty"`

	// Firmware. Location. URI
	// urn:x-enexis:ecdm:uid:1:569460
	// URI defining the origin of the firmware.
	//
	Location string `json:"location" yaml:"location"`

	// Firmware. Retrieve. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569461
	// Date and time at which the firmware shall be retrieved.
	//
	RetrieveDateTime string `json:"retrieveDateTime" yaml:"retrieveDateTime"`

	// Firmware. Signature. Signature
	// urn:x-enexis:ecdm:uid:1:569464
	// Base64 encoded firmware signature.
	//
	Signature *string `json:"signature,omitempty" yaml:"signature,omitempty"`

	// Certificate with which the firmware was signed.
	// PEM encoded X.509 certificate.
	//
	SigningCertificate *string `json:"signingCertificate,omitempty" yaml:"signingCertificate,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FirmwareType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["location"]; !ok || v == nil {
		return fmt.Errorf("field location in FirmwareType: required")
	}
	if v, ok := raw["retrieveDateTime"]; !ok || v == nil {
		return fmt.Errorf("field retrieveDateTime in FirmwareType: required")
	}
	type Plain FirmwareType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FirmwareType(plain)
	return nil
}

type UpdateFirmwareRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Firmware corresponds to the JSON schema field "firmware".
	Firmware FirmwareType `json:"firmware" yaml:"firmware"`

	// The Id of this request
	//
	RequestId int `json:"requestId" yaml:"requestId"`

	// This specifies how many times Charging Station must try to download the
	// firmware before giving up. If this field is not present, it is left to Charging
	// Station to decide how many times it wants to retry.
	//
	Retries *int `json:"retries,omitempty" yaml:"retries,omitempty"`

	// The interval in seconds after which a retry may be attempted. If this field is
	// not present, it is left to Charging Station to decide how long to wait between
	// attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty" yaml:"retryInterval,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateFirmwareRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["firmware"]; !ok || v == nil {
		return fmt.Errorf("field firmware in UpdateFirmwareRequestJson: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId in UpdateFirmwareRequestJson: required")
	}
	type Plain UpdateFirmwareRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UpdateFirmwareRequestJson(plain)
	return nil
}
