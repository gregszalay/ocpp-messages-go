// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

type CertificateHashDataType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm" yaml:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash" yaml:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash" yaml:"issuerNameHash"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber" yaml:"serialNumber"`
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
func (j *HashAlgorithmEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_1, v)
	}
	*j = HashAlgorithmEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateHashDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm in CertificateHashDataType: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash in CertificateHashDataType: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash in CertificateHashDataType: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber in CertificateHashDataType: required")
	}
	type Plain CertificateHashDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateHashDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType, v)
	}
	*j = HashAlgorithmEnumType(v)
	return nil
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

type DeleteCertificateRequestJson struct {
	// CertificateHashData corresponds to the JSON schema field "certificateHashData".
	CertificateHashData CertificateHashDataType `json:"certificateHashData" yaml:"certificateHashData"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`
}

type HashAlgorithmEnumType string

const HashAlgorithmEnumTypeSHA256 HashAlgorithmEnumType = "SHA256"
const HashAlgorithmEnumTypeSHA384 HashAlgorithmEnumType = "SHA384"
const HashAlgorithmEnumTypeSHA512 HashAlgorithmEnumType = "SHA512"

type HashAlgorithmEnumType_1 string

const HashAlgorithmEnumType_1_SHA256 HashAlgorithmEnumType_1 = "SHA256"
const HashAlgorithmEnumType_1_SHA384 HashAlgorithmEnumType_1 = "SHA384"
const HashAlgorithmEnumType_1_SHA512 HashAlgorithmEnumType_1 = "SHA512"

var enumValues_HashAlgorithmEnumType = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_1 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["certificateHashData"]; !ok || v == nil {
		return fmt.Errorf("field certificateHashData in DeleteCertificateRequestJson: required")
	}
	type Plain DeleteCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DeleteCertificateRequestJson(plain)
	return nil
}
