// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType, v)
	}
	*j = AuthorizationStatusEnumType(v)
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

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken" yaml:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type" yaml:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken in AdditionalInfoType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in AdditionalInfoType: required")
	}
	type Plain AdditionalInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType(plain)
	return nil
}

// Contains the identifier to use for authorization.
//
type AuthorizationData struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType `json:"idToken" yaml:"idToken"`

	// IdTokenInfo corresponds to the JSON schema field "idTokenInfo".
	IdTokenInfo *IdTokenInfoType `json:"idTokenInfo,omitempty" yaml:"idTokenInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateEnumType_1, v)
	}
	*j = UpdateEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType, v)
	}
	*j = IdTokenEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status in IdTokenInfoType: required")
	}
	type Plain IdTokenInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.EvseId) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "evseId", 1)
	}
	*j = IdTokenInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_1, v)
	}
	*j = AuthorizationStatusEnumType_1(v)
	return nil
}

const AuthorizationStatusEnumTypeUnknown AuthorizationStatusEnumType = "Unknown"
const AuthorizationStatusEnumTypeNotAtThisTime AuthorizationStatusEnumType = "NotAtThisTime"

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateEnumType, v)
	}
	*j = UpdateEnumType(v)
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

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_1, v)
	}
	*j = IdTokenEnumType_1(v)
	return nil
}

const AuthorizationStatusEnumTypeNotAtThisLocation AuthorizationStatusEnumType = "NotAtThisLocation"
const AuthorizationStatusEnumTypeNotAllowedTypeEVSE AuthorizationStatusEnumType = "NotAllowedTypeEVSE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken in IdTokenType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in IdTokenType: required")
	}
	type Plain IdTokenType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.AdditionalInfo) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "additionalInfo", 1)
	}
	*j = IdTokenType(plain)
	return nil
}

const AuthorizationStatusEnumTypeNoCredit AuthorizationStatusEnumType = "NoCredit"
const AuthorizationStatusEnumTypeInvalid AuthorizationStatusEnumType = "Invalid"

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

const AuthorizationStatusEnumTypeExpired AuthorizationStatusEnumType = "Expired"

type AuthorizationStatusEnumType string

const AuthorizationStatusEnumTypeConcurrentTx AuthorizationStatusEnumType = "ConcurrentTx"
const AuthorizationStatusEnumTypeBlocked AuthorizationStatusEnumType = "Blocked"
const AuthorizationStatusEnumTypeAccepted AuthorizationStatusEnumType = "Accepted"

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

type AuthorizationStatusEnumType_1 string

const AuthorizationStatusEnumType_1_Accepted AuthorizationStatusEnumType_1 = "Accepted"
const AuthorizationStatusEnumType_1_Blocked AuthorizationStatusEnumType_1 = "Blocked"
const AuthorizationStatusEnumType_1_ConcurrentTx AuthorizationStatusEnumType_1 = "ConcurrentTx"
const AuthorizationStatusEnumType_1_Expired AuthorizationStatusEnumType_1 = "Expired"
const AuthorizationStatusEnumType_1_Invalid AuthorizationStatusEnumType_1 = "Invalid"
const AuthorizationStatusEnumType_1_NoCredit AuthorizationStatusEnumType_1 = "NoCredit"
const AuthorizationStatusEnumType_1_NotAllowedTypeEVSE AuthorizationStatusEnumType_1 = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumType_1_NotAtThisLocation AuthorizationStatusEnumType_1 = "NotAtThisLocation"
const AuthorizationStatusEnumType_1_NotAtThisTime AuthorizationStatusEnumType_1 = "NotAtThisTime"
const AuthorizationStatusEnumType_1_Unknown AuthorizationStatusEnumType_1 = "Unknown"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

type IdTokenEnumType string

const IdTokenEnumTypeCentral IdTokenEnumType = "Central"
const IdTokenEnumTypeEMAID IdTokenEnumType = "eMAID"
const IdTokenEnumTypeISO14443 IdTokenEnumType = "ISO14443"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationData) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken in AuthorizationData: required")
	}
	type Plain AuthorizationData
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuthorizationData(plain)
	return nil
}

const IdTokenEnumTypeISO15693 IdTokenEnumType = "ISO15693"
const IdTokenEnumTypeKeyCode IdTokenEnumType = "KeyCode"
const IdTokenEnumTypeLocal IdTokenEnumType = "Local"
const IdTokenEnumTypeMacAddress IdTokenEnumType = "MacAddress"
const IdTokenEnumTypeNoAuthorization IdTokenEnumType = "NoAuthorization"

type IdTokenEnumType_1 string

const IdTokenEnumType_1_Central IdTokenEnumType_1 = "Central"
const IdTokenEnumType_1_EMAID IdTokenEnumType_1 = "eMAID"
const IdTokenEnumType_1_ISO14443 IdTokenEnumType_1 = "ISO14443"
const IdTokenEnumType_1_ISO15693 IdTokenEnumType_1 = "ISO15693"
const IdTokenEnumType_1_KeyCode IdTokenEnumType_1 = "KeyCode"
const IdTokenEnumType_1_Local IdTokenEnumType_1 = "Local"
const IdTokenEnumType_1_MacAddress IdTokenEnumType_1 = "MacAddress"
const IdTokenEnumType_1_NoAuthorization IdTokenEnumType_1 = "NoAuthorization"

// ID_ Token
// urn:x-oca:ocpp:uid:2:233247
// Contains status information about an identifier.
// It is advised to not stop charging for a token that expires during charging, as
// ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the
// status has no end date.
//
type IdTokenInfoType struct {
	// ID_ Token. Expiry. Date_ Time
	// urn:x-oca:ocpp:uid:1:569373
	// Date and Time after which the token must be considered invalid.
	//
	CacheExpiryDateTime *string `json:"cacheExpiryDateTime,omitempty" yaml:"cacheExpiryDateTime,omitempty"`

	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules
	// this one.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty" yaml:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Only used when the IdToken is only valid for one or more specific EVSEs, not
	// for the entire Charging Station.
	//
	//
	EvseId []int `json:"evseId,omitempty" yaml:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType `json:"groupIdToken,omitempty" yaml:"groupIdToken,omitempty"`

	// ID_ Token. Language1. Language_ Code
	// urn:x-oca:ocpp:uid:1:569374
	// Preferred user interface language of identifier user. Contains a language code
	// as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	//
	Language1 *string `json:"language1,omitempty" yaml:"language1,omitempty"`

	// ID_ Token. Language2. Language_ Code
	// urn:x-oca:ocpp:uid:1:569375
	// Second preferred user interface language of identifier user. Don’t use when
	// language1 is omitted, has to be different from language1. Contains a language
	// code as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language2 *string `json:"language2,omitempty" yaml:"language2,omitempty"`

	// PersonalMessage corresponds to the JSON schema field "personalMessage".
	PersonalMessage *MessageContentType `json:"personalMessage,omitempty" yaml:"personalMessage,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status AuthorizationStatusEnumType `json:"status" yaml:"status"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken" yaml:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType `json:"type" yaml:"type"`
}

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

type MessageFormatEnumType string

const MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
const MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
const MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
const MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"

type MessageFormatEnumType_1 string

const MessageFormatEnumType_1_ASCII MessageFormatEnumType_1 = "ASCII"
const MessageFormatEnumType_1_HTML MessageFormatEnumType_1 = "HTML"
const MessageFormatEnumType_1_URI MessageFormatEnumType_1 = "URI"
const MessageFormatEnumType_1_UTF8 MessageFormatEnumType_1 = "UTF8"

type SendLocalListRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// LocalAuthorizationList corresponds to the JSON schema field
	// "localAuthorizationList".
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty" yaml:"localAuthorizationList,omitempty"`

	// UpdateType corresponds to the JSON schema field "updateType".
	UpdateType UpdateEnumType_1 `json:"updateType" yaml:"updateType"`

	// In case of a full update this is the version number of the full list. In case
	// of a differential update it is the version number of the list after the update
	// has been applied.
	//
	VersionNumber int `json:"versionNumber" yaml:"versionNumber"`
}

type UpdateEnumType string

const UpdateEnumTypeDifferential UpdateEnumType = "Differential"
const UpdateEnumTypeFull UpdateEnumType = "Full"

type UpdateEnumType_1 string

const UpdateEnumType_1_Differential UpdateEnumType_1 = "Differential"
const UpdateEnumType_1_Full UpdateEnumType_1 = "Full"

var enumValues_AuthorizationStatusEnumType = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_1 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_IdTokenEnumType = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_1 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
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
var enumValues_UpdateEnumType = []interface{}{
	"Differential",
	"Full",
}
var enumValues_UpdateEnumType_1 = []interface{}{
	"Differential",
	"Full",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SendLocalListRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["updateType"]; !ok || v == nil {
		return fmt.Errorf("field updateType in SendLocalListRequestJson: required")
	}
	if v, ok := raw["versionNumber"]; !ok || v == nil {
		return fmt.Errorf("field versionNumber in SendLocalListRequestJson: required")
	}
	type Plain SendLocalListRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.LocalAuthorizationList) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "localAuthorizationList", 1)
	}
	*j = SendLocalListRequestJson(plain)
	return nil
}