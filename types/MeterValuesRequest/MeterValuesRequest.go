// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package MeterValuesRequest

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType_1, v)
	}
	*j = LocationEnumType_1(v)
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

type LocationEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType_1, v)
	}
	*j = ReadingContextEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType, v)
	}
	*j = LocationEnumType(v)
	return nil
}

const LocationEnumTypeBody LocationEnumType = "Body"
const LocationEnumTypeCable LocationEnumType = "Cable"
const LocationEnumTypeEV LocationEnumType = "EV"
const LocationEnumTypeInlet LocationEnumType = "Inlet"
const LocationEnumTypeOutlet LocationEnumType = "Outlet"

type MeasurandEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType_1, v)
	}
	*j = PhaseEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType, v)
	}
	*j = MeasurandEnumType(v)
	return nil
}

const MeasurandEnumTypeCurrentExport MeasurandEnumType = "Current.Export"
const MeasurandEnumTypeCurrentImport MeasurandEnumType = "Current.Import"
const MeasurandEnumTypeCurrentOffered MeasurandEnumType = "Current.Offered"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

const MeasurandEnumTypeEnergyActiveImportRegister MeasurandEnumType = "Energy.Active.Import.Register"
const MeasurandEnumTypeEnergyActiveExportRegister MeasurandEnumType = "Energy.Active.Export.Register"
const MeasurandEnumTypeEnergyReactiveImportRegister MeasurandEnumType = "Energy.Reactive.Import.Register"
const MeasurandEnumTypeEnergyActiveExportInterval MeasurandEnumType = "Energy.Active.Export.Interval"
const MeasurandEnumTypeEnergyActiveImportInterval MeasurandEnumType = "Energy.Active.Import.Interval"
const MeasurandEnumTypeEnergyActiveNet MeasurandEnumType = "Energy.Active.Net"
const MeasurandEnumTypeEnergyReactiveExportInterval MeasurandEnumType = "Energy.Reactive.Export.Interval"
const MeasurandEnumTypeEnergyReactiveImportInterval MeasurandEnumType = "Energy.Reactive.Import.Interval"
const MeasurandEnumTypeEnergyReactiveNet MeasurandEnumType = "Energy.Reactive.Net"
const MeasurandEnumTypeEnergyApparentNet MeasurandEnumType = "Energy.Apparent.Net"
const MeasurandEnumTypeEnergyApparentImport MeasurandEnumType = "Energy.Apparent.Import"
const MeasurandEnumTypeEnergyApparentExport MeasurandEnumType = "Energy.Apparent.Export"
const MeasurandEnumTypeFrequency MeasurandEnumType = "Frequency"
const MeasurandEnumTypePowerActiveExport MeasurandEnumType = "Power.Active.Export"
const MeasurandEnumTypePowerActiveImport MeasurandEnumType = "Power.Active.Import"
const MeasurandEnumTypePowerFactor MeasurandEnumType = "Power.Factor"
const MeasurandEnumTypePowerOffered MeasurandEnumType = "Power.Offered"
const MeasurandEnumTypePowerReactiveExport MeasurandEnumType = "Power.Reactive.Export"
const MeasurandEnumTypePowerReactiveImport MeasurandEnumType = "Power.Reactive.Import"
const MeasurandEnumTypeSoC MeasurandEnumType = "SoC"
const MeasurandEnumTypeEnergyReactiveExportRegister MeasurandEnumType = "Energy.Reactive.Export.Register"

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeterValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["sampledValue"]; !ok || v == nil {
		return fmt.Errorf("field sampledValue in MeterValueType: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp in MeterValueType: required")
	}
	type Plain MeterValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.SampledValue) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "sampledValue", 1)
	}
	*j = MeterValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SampledValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in SampledValueType: required")
	}
	type Plain SampledValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SampledValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType, v)
	}
	*j = ReadingContextEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnitOfMeasureType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain UnitOfMeasureType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["multiplier"]; !ok || v == nil {
		plain.Multiplier = 0
	}
	if v, ok := raw["unit"]; !ok || v == nil {
		plain.Unit = "Wh"
	}
	*j = UnitOfMeasureType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignedMeterValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["encodingMethod"]; !ok || v == nil {
		return fmt.Errorf("field encodingMethod in SignedMeterValueType: required")
	}
	if v, ok := raw["publicKey"]; !ok || v == nil {
		return fmt.Errorf("field publicKey in SignedMeterValueType: required")
	}
	if v, ok := raw["signedMeterData"]; !ok || v == nil {
		return fmt.Errorf("field signedMeterData in SignedMeterValueType: required")
	}
	if v, ok := raw["signingMethod"]; !ok || v == nil {
		return fmt.Errorf("field signingMethod in SignedMeterValueType: required")
	}
	type Plain SignedMeterValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignedMeterValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType, v)
	}
	*j = PhaseEnumType(v)
	return nil
}

const LocationEnumType_1_Body LocationEnumType_1 = "Body"

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType_1, v)
	}
	*j = MeasurandEnumType_1(v)
	return nil
}

const LocationEnumType_1_Inlet LocationEnumType_1 = "Inlet"
const LocationEnumType_1_EV LocationEnumType_1 = "EV"
const LocationEnumType_1_Cable LocationEnumType_1 = "Cable"

type LocationEnumType_1 string

const LocationEnumType_1_Outlet LocationEnumType_1 = "Outlet"
const MeasurandEnumTypeVoltage MeasurandEnumType = "Voltage"

type MeasurandEnumType_1 string

const MeasurandEnumType_1_CurrentExport MeasurandEnumType_1 = "Current.Export"
const MeasurandEnumType_1_CurrentImport MeasurandEnumType_1 = "Current.Import"
const MeasurandEnumType_1_CurrentOffered MeasurandEnumType_1 = "Current.Offered"
const MeasurandEnumType_1_EnergyActiveExportInterval MeasurandEnumType_1 = "Energy.Active.Export.Interval"
const MeasurandEnumType_1_EnergyActiveExportRegister MeasurandEnumType_1 = "Energy.Active.Export.Register"
const MeasurandEnumType_1_EnergyActiveImportInterval MeasurandEnumType_1 = "Energy.Active.Import.Interval"
const MeasurandEnumType_1_EnergyActiveImportRegister MeasurandEnumType_1 = "Energy.Active.Import.Register"
const MeasurandEnumType_1_EnergyActiveNet MeasurandEnumType_1 = "Energy.Active.Net"
const MeasurandEnumType_1_EnergyApparentExport MeasurandEnumType_1 = "Energy.Apparent.Export"
const MeasurandEnumType_1_EnergyApparentImport MeasurandEnumType_1 = "Energy.Apparent.Import"
const MeasurandEnumType_1_EnergyApparentNet MeasurandEnumType_1 = "Energy.Apparent.Net"
const MeasurandEnumType_1_EnergyReactiveExportInterval MeasurandEnumType_1 = "Energy.Reactive.Export.Interval"
const MeasurandEnumType_1_EnergyReactiveExportRegister MeasurandEnumType_1 = "Energy.Reactive.Export.Register"
const MeasurandEnumType_1_EnergyReactiveImportInterval MeasurandEnumType_1 = "Energy.Reactive.Import.Interval"
const MeasurandEnumType_1_EnergyReactiveImportRegister MeasurandEnumType_1 = "Energy.Reactive.Import.Register"
const MeasurandEnumType_1_EnergyReactiveNet MeasurandEnumType_1 = "Energy.Reactive.Net"
const MeasurandEnumType_1_Frequency MeasurandEnumType_1 = "Frequency"
const MeasurandEnumType_1_PowerActiveExport MeasurandEnumType_1 = "Power.Active.Export"
const MeasurandEnumType_1_PowerActiveImport MeasurandEnumType_1 = "Power.Active.Import"
const MeasurandEnumType_1_PowerFactor MeasurandEnumType_1 = "Power.Factor"
const MeasurandEnumType_1_PowerOffered MeasurandEnumType_1 = "Power.Offered"
const MeasurandEnumType_1_PowerReactiveExport MeasurandEnumType_1 = "Power.Reactive.Export"
const MeasurandEnumType_1_PowerReactiveImport MeasurandEnumType_1 = "Power.Reactive.Import"
const MeasurandEnumType_1_SoC MeasurandEnumType_1 = "SoC"
const MeasurandEnumType_1_Voltage MeasurandEnumType_1 = "Voltage"

// Meter_ Value
// urn:x-oca:ocpp:uid:2:233265
// Collection of one or more sampled values in MeterValuesRequest and
// TransactionEvent. All sampled values in a MeterValue are sampled at the same
// point in time.
//
type MeterValueType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// SampledValue corresponds to the JSON schema field "sampledValue".
	SampledValue []SampledValueType `json:"sampledValue" yaml:"sampledValue"`

	// Meter_ Value. Timestamp. Date_ Time
	// urn:x-oca:ocpp:uid:1:569259
	// Timestamp for measured value(s).
	//
	Timestamp string `json:"timestamp" yaml:"timestamp"`
}

// Request_ Body
// urn:x-enexis:ecdm:uid:2:234744
//
type MeterValuesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Request_ Body. EVSEID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:571101
	// This contains a number (&gt;0) designating an EVSE of the Charging Station.
	// ‘0’ (zero) is used to designate the main power meter.
	//
	EvseId int `json:"evseId" yaml:"evseId"`

	// MeterValue corresponds to the JSON schema field "meterValue".
	MeterValue []MeterValueType `json:"meterValue" yaml:"meterValue"`
}

type PhaseEnumType string

const PhaseEnumTypeL1 PhaseEnumType = "L1"
const PhaseEnumTypeL1L2 PhaseEnumType = "L1-L2"
const PhaseEnumTypeL1N PhaseEnumType = "L1-N"
const PhaseEnumTypeL2 PhaseEnumType = "L2"
const PhaseEnumTypeL2L3 PhaseEnumType = "L2-L3"
const PhaseEnumTypeL2N PhaseEnumType = "L2-N"
const PhaseEnumTypeL3 PhaseEnumType = "L3"
const PhaseEnumTypeL3L1 PhaseEnumType = "L3-L1"
const PhaseEnumTypeL3N PhaseEnumType = "L3-N"
const PhaseEnumTypeN PhaseEnumType = "N"

type PhaseEnumType_1 string

const PhaseEnumType_1_L1 PhaseEnumType_1 = "L1"
const PhaseEnumType_1_L1L2 PhaseEnumType_1 = "L1-L2"
const PhaseEnumType_1_L1N PhaseEnumType_1 = "L1-N"
const PhaseEnumType_1_L2 PhaseEnumType_1 = "L2"
const PhaseEnumType_1_L2L3 PhaseEnumType_1 = "L2-L3"
const PhaseEnumType_1_L2N PhaseEnumType_1 = "L2-N"
const PhaseEnumType_1_L3 PhaseEnumType_1 = "L3"
const PhaseEnumType_1_L3L1 PhaseEnumType_1 = "L3-L1"
const PhaseEnumType_1_L3N PhaseEnumType_1 = "L3-N"
const PhaseEnumType_1_N PhaseEnumType_1 = "N"

type ReadingContextEnumType string

const ReadingContextEnumTypeInterruptionBegin ReadingContextEnumType = "Interruption.Begin"
const ReadingContextEnumTypeInterruptionEnd ReadingContextEnumType = "Interruption.End"
const ReadingContextEnumTypeOther ReadingContextEnumType = "Other"
const ReadingContextEnumTypeSampleClock ReadingContextEnumType = "Sample.Clock"
const ReadingContextEnumTypeSamplePeriodic ReadingContextEnumType = "Sample.Periodic"
const ReadingContextEnumTypeTransactionBegin ReadingContextEnumType = "Transaction.Begin"
const ReadingContextEnumTypeTransactionEnd ReadingContextEnumType = "Transaction.End"
const ReadingContextEnumTypeTrigger ReadingContextEnumType = "Trigger"

type ReadingContextEnumType_1 string

const ReadingContextEnumType_1_InterruptionBegin ReadingContextEnumType_1 = "Interruption.Begin"
const ReadingContextEnumType_1_InterruptionEnd ReadingContextEnumType_1 = "Interruption.End"
const ReadingContextEnumType_1_Other ReadingContextEnumType_1 = "Other"
const ReadingContextEnumType_1_SampleClock ReadingContextEnumType_1 = "Sample.Clock"
const ReadingContextEnumType_1_SamplePeriodic ReadingContextEnumType_1 = "Sample.Periodic"
const ReadingContextEnumType_1_TransactionBegin ReadingContextEnumType_1 = "Transaction.Begin"
const ReadingContextEnumType_1_TransactionEnd ReadingContextEnumType_1 = "Transaction.End"
const ReadingContextEnumType_1_Trigger ReadingContextEnumType_1 = "Trigger"

// Sampled_ Value
// urn:x-oca:ocpp:uid:2:233266
// Single sampled value in MeterValues. Each value can be accompanied by optional
// fields.
//
// To save on mobile data usage, default values of all of the optional fields are
// such that. The value without any additional fields will be interpreted, as a
// register reading of active import energy in Wh (Watt-hour) units.
//
type SampledValueType struct {
	// Context corresponds to the JSON schema field "context".
	Context *ReadingContextEnumType `json:"context,omitempty" yaml:"context,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *LocationEnumType_1 `json:"location,omitempty" yaml:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand *MeasurandEnumType_1 `json:"measurand,omitempty" yaml:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *PhaseEnumType `json:"phase,omitempty" yaml:"phase,omitempty"`

	// SignedMeterValue corresponds to the JSON schema field "signedMeterValue".
	SignedMeterValue *SignedMeterValueType `json:"signedMeterValue,omitempty" yaml:"signedMeterValue,omitempty"`

	// UnitOfMeasure corresponds to the JSON schema field "unitOfMeasure".
	UnitOfMeasure *UnitOfMeasureType `json:"unitOfMeasure,omitempty" yaml:"unitOfMeasure,omitempty"`

	// Sampled_ Value. Value. Measure
	// urn:x-oca:ocpp:uid:1:569260
	// Indicates the measured value.
	//
	//
	Value float64 `json:"value" yaml:"value"`
}

// Represent a signed version of the meter value.
//
type SignedMeterValueType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Method used to encode the meter values before applying the digital signature
	// algorithm.
	//
	EncodingMethod string `json:"encodingMethod" yaml:"encodingMethod"`

	// Base64 encoded, sending depends on configuration variable
	// _PublicKeyWithSignedMeterValue_.
	//
	PublicKey string `json:"publicKey" yaml:"publicKey"`

	// Base64 encoded, contains the signed data which might contain more then just the
	// meter value. It can contain information like timestamps, reference to a
	// customer etc.
	//
	SignedMeterData string `json:"signedMeterData" yaml:"signedMeterData"`

	// Method used to create the digital signature.
	//
	SigningMethod string `json:"signingMethod" yaml:"signingMethod"`
}

// Represents a UnitOfMeasure with a multiplier
//
type UnitOfMeasureType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Multiplier, this value represents the exponent to base 10. I.e. multiplier 3
	// means 10 raised to the 3rd power. Default is 0.
	//
	Multiplier int `json:"multiplier,omitempty" yaml:"multiplier,omitempty"`

	// Unit of the value. Default = "Wh" if the (default) measurand is an "Energy"
	// type.
	// This field SHALL use a value from the list Standardized Units of Measurements
	// in Part 2 Appendices.
	// If an applicable unit is available in that list, otherwise a "custom" unit
	// might be used.
	//
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}

var enumValues_LocationEnumType = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_LocationEnumType_1 = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_MeasurandEnumType = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}
var enumValues_MeasurandEnumType_1 = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeterValuesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId in MeterValuesRequestJson: required")
	}
	if v, ok := raw["meterValue"]; !ok || v == nil {
		return fmt.Errorf("field meterValue in MeterValuesRequestJson: required")
	}
	type Plain MeterValuesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.MeterValue) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "meterValue", 1)
	}
	*j = MeterValuesRequestJson(plain)
	return nil
}

var enumValues_PhaseEnumType = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_PhaseEnumType_1 = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_ReadingContextEnumType = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}
var enumValues_ReadingContextEnumType_1 = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}