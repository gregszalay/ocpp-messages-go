// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package RequestStopTransactionRequest

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

type RequestStopTransactionRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// The identifier of the transaction which the Charging Station is requested to
	// stop.
	//
	TransactionId string `json:"transactionId" yaml:"transactionId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStopTransactionRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["transactionId"]; !ok || v == nil {
		return fmt.Errorf("field transactionId in RequestStopTransactionRequestJson: required")
	}
	type Plain RequestStopTransactionRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStopTransactionRequestJson(plain)
	return nil
}
