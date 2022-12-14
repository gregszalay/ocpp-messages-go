// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package CostUpdatedRequest

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

type CostUpdatedRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Current total cost, based on the information known by the CSMS, of the
	// transaction including taxes. In the currency configured with the configuration
	// Variable: [&lt;&lt;configkey-currency, Currency&gt;&gt;]
	//
	//
	TotalCost float64 `json:"totalCost" yaml:"totalCost"`

	// Transaction Id of the transaction the current cost are asked for.
	//
	//
	TransactionId string `json:"transactionId" yaml:"transactionId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostUpdatedRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["totalCost"]; !ok || v == nil {
		return fmt.Errorf("field totalCost in CostUpdatedRequestJson: required")
	}
	if v, ok := raw["transactionId"]; !ok || v == nil {
		return fmt.Errorf("field transactionId in CostUpdatedRequestJson: required")
	}
	type Plain CostUpdatedRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostUpdatedRequestJson(plain)
	return nil
}
