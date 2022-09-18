// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package GetCertificateStatusRequest

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

type GetCertificateStatusRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// OcspRequestData corresponds to the JSON schema field "ocspRequestData".
	OcspRequestData OCSPRequestDataType `json:"ocspRequestData" yaml:"ocspRequestData"`
}

type HashAlgorithmEnumType string

const HashAlgorithmEnumTypeSHA256 HashAlgorithmEnumType = "SHA256"
const HashAlgorithmEnumTypeSHA384 HashAlgorithmEnumType = "SHA384"
const HashAlgorithmEnumTypeSHA512 HashAlgorithmEnumType = "SHA512"

type HashAlgorithmEnumType_1 string

const HashAlgorithmEnumType_1_SHA256 HashAlgorithmEnumType_1 = "SHA256"
const HashAlgorithmEnumType_1_SHA384 HashAlgorithmEnumType_1 = "SHA384"
const HashAlgorithmEnumType_1_SHA512 HashAlgorithmEnumType_1 = "SHA512"

type OCSPRequestDataType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_1 `json:"hashAlgorithm" yaml:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash" yaml:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash" yaml:"issuerNameHash"`

	// This contains the responder URL (Case insensitive).
	//
	//
	ResponderURL string `json:"responderURL" yaml:"responderURL"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber" yaml:"serialNumber"`
}

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

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCSPRequestDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm in OCSPRequestDataType: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash in OCSPRequestDataType: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash in OCSPRequestDataType: required")
	}
	if v, ok := raw["responderURL"]; !ok || v == nil {
		return fmt.Errorf("field responderURL in OCSPRequestDataType: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber in OCSPRequestDataType: required")
	}
	type Plain OCSPRequestDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = OCSPRequestDataType(plain)
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
func (j *GetCertificateStatusRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["ocspRequestData"]; !ok || v == nil {
		return fmt.Errorf("field ocspRequestData in GetCertificateStatusRequestJson: required")
	}
	type Plain GetCertificateStatusRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetCertificateStatusRequestJson(plain)
	return nil
}