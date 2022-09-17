// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType, v)
	}
	*j = MessageStateEnumType(v)
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

// A physical or logical component
//
type ComponentType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in ComponentType: required")
	}
	type Plain ComponentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType(plain)
	return nil
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_1, v)
	}
	*j = MessageStateEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType, v)
	}
	*j = MessageFormatEnumType(v)
	return nil
}

const MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
const MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
const MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
const MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content" yaml:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType `json:"format" yaml:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty" yaml:"language,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content in MessageContentType: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format in MessageContentType: required")
	}
	type Plain MessageContentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType(plain)
	return nil
}

type MessageFormatEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_1, v)
	}
	*j = MessagePriorityEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_1, v)
	}
	*j = MessageFormatEnumType_1(v)
	return nil
}

const MessageFormatEnumType_1_ASCII MessageFormatEnumType_1 = "ASCII"
const MessageFormatEnumType_1_HTML MessageFormatEnumType_1 = "HTML"
const MessageFormatEnumType_1_URI MessageFormatEnumType_1 = "URI"
const MessageFormatEnumType_1_UTF8 MessageFormatEnumType_1 = "UTF8"

type MessagePriorityEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in MessageInfoType: required")
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message in MessageInfoType: required")
	}
	if v, ok := raw["priority"]; !ok || v == nil {
		return fmt.Errorf("field priority in MessageInfoType: required")
	}
	type Plain MessageInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType, v)
	}
	*j = MessagePriorityEnumType(v)
	return nil
}

const MessagePriorityEnumTypeAlwaysFront MessagePriorityEnumType = "AlwaysFront"

type MessageFormatEnumType string

// Message_ Info
// urn:x-enexis:ecdm:uid:2:233264
// Contains message details, for a message to be displayed on a Charging Station.
//
type MessageInfoType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Display corresponds to the JSON schema field "display".
	Display *ComponentType `json:"display,omitempty" yaml:"display,omitempty"`

	// Message_ Info. End. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569257
	// Until what date-time should this message be shown, after this date/time this
	// message SHALL be removed.
	//
	EndDateTime *string `json:"endDateTime,omitempty" yaml:"endDateTime,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Master resource identifier, unique within an exchange context. It is defined
	// within the OCPP context as a positive Integer value (greater or equal to zero).
	//
	Id int `json:"id" yaml:"id"`

	// Message corresponds to the JSON schema field "message".
	Message MessageContentType `json:"message" yaml:"message"`

	// Priority corresponds to the JSON schema field "priority".
	Priority MessagePriorityEnumType `json:"priority" yaml:"priority"`

	// Message_ Info. Start. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569256
	// From what date-time should this message be shown. If omitted: directly.
	//
	StartDateTime *string `json:"startDateTime,omitempty" yaml:"startDateTime,omitempty"`

	// State corresponds to the JSON schema field "state".
	State *MessageStateEnumType `json:"state,omitempty" yaml:"state,omitempty"`

	// During which transaction shall this message be shown.
	// Message SHALL be removed by the Charging Station after transaction has
	// ended.
	//
	TransactionId *string `json:"transactionId,omitempty" yaml:"transactionId,omitempty"`
}

const MessagePriorityEnumTypeInFront MessagePriorityEnumType = "InFront"
const MessagePriorityEnumTypeNormalCycle MessagePriorityEnumType = "NormalCycle"

type MessagePriorityEnumType_1 string

const MessagePriorityEnumType_1_AlwaysFront MessagePriorityEnumType_1 = "AlwaysFront"
const MessagePriorityEnumType_1_InFront MessagePriorityEnumType_1 = "InFront"
const MessagePriorityEnumType_1_NormalCycle MessagePriorityEnumType_1 = "NormalCycle"

type MessageStateEnumType string

const MessageStateEnumTypeCharging MessageStateEnumType = "Charging"
const MessageStateEnumTypeFaulted MessageStateEnumType = "Faulted"
const MessageStateEnumTypeIdle MessageStateEnumType = "Idle"
const MessageStateEnumTypeUnavailable MessageStateEnumType = "Unavailable"

type MessageStateEnumType_1 string

const MessageStateEnumType_1_Charging MessageStateEnumType_1 = "Charging"
const MessageStateEnumType_1_Faulted MessageStateEnumType_1 = "Faulted"
const MessageStateEnumType_1_Idle MessageStateEnumType_1 = "Idle"
const MessageStateEnumType_1_Unavailable MessageStateEnumType_1 = "Unavailable"

type SetDisplayMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Message corresponds to the JSON schema field "message".
	Message MessageInfoType `json:"message" yaml:"message"`
}

var enumValues_MessageFormatEnumType = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_1 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessagePriorityEnumType = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_1 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessageStateEnumType = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_1 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetDisplayMessageRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message in SetDisplayMessageRequestJson: required")
	}
	type Plain SetDisplayMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetDisplayMessageRequestJson(plain)
	return nil
}
