package wrappers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CALLERROR struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId uint32
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// This field must contain a string from the RPC Framework Error Codes table.
	ErrorCode string
	// Should be filled in if possible, otherwise a clear empty string "".
	ErrorDescription string
	// This JSON object describes error details in an undefined way. If there are no error details you MUST fill in an empty object {}.
	ErrorDetails string
}

func (j *CALLERROR) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	err := json.Unmarshal([]byte(b), &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	if len(raw) != 5 {
		return errors.New("invalid CALL ERROR message length")
	}
	//MessageTypeId integer This is a Message Type Number which is used to identify the type of the message.
	message_type_id, message_type_id_err := raw[0].(float64)
	if message_type_id_err {
		return fmt.Errorf("CALL data[0] is not a number")
	}
	//MessageId string[36] This must be the exact same id that is in the call request so that the recipient can match request and result.
	message_id, message_id_err := raw[1].(string)
	if message_id_err || message_id == "" {
		return fmt.Errorf("CALL data[1] is not a string")
	}
	// ErrorCode string This field must contain a string from the RPC Framework Error Codes table.
	error_code, error_code_err := raw[2].(string)
	if error_code_err || error_code == "" {
		return fmt.Errorf("CALLERROR data[2] is not a string")
	}
	//ErrorDescription string[255] Should be filled in if possible, otherwise a clear empty string "".
	error_description, error_description_err := raw[3].(string)
	if error_description_err || error_description == "" {
		return fmt.Errorf("CALLERROR data[3] is not a string")
	}
	//ErrorDetails JSON This JSON object describes error details in an undefined way. If there are no error details you
	//MUST fill in an empty object {}.
	error_details, error_details_err := raw[4].(string)
	if error_details_err || error_details == "" {
		return fmt.Errorf("CALLERROR data[4] is not a string")
	}
	*j = CALLERROR{
		MessageTypeId:    uint32(message_type_id),
		MessageId:        message_id,
		ErrorCode:        error_code,
		ErrorDescription: error_description,
		ErrorDetails:     error_details,
	}
	return nil
}

func (c *CALLERROR) Marshal() []byte {
	message_array := [...]interface{}{CALLERROR_TYPE, c.MessageId, c.ErrorCode, c.ErrorDescription, c.ErrorDetails}
	message_array_json, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALLERROR message: %s\n", err)
		return []byte("")
	}
	return message_array_json
}

func (c *CALLERROR) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}

type rpc_framework_error_code string

// Payload for Action is syntactically incorrect
const FormatViolation rpc_framework_error_code = "FormatViolation"

// Any other error not covered by the more specific error codes in this table
const GenericError rpc_framework_error_code = "GenericError"

// An internal error occurred and the receiver was not able to process the requested Action successfully
const InternalError rpc_framework_error_code = "InternalError"

// A message with an Message Type Number received that is not supported by this implementation.
const MessageTypeNotSupported rpc_framework_error_code = "MessageTypeNotSupported"

// Requested Action is not known by receiver
const NotImplemented rpc_framework_error_code = "NotImplemented"

// Requested Action is recognized but not supported by the receiver
const NotSupported rpc_framework_error_code = "NotSupported"

// Payload for Action is syntactically correct but at least one of the fields violates occurrence constraints
const OccurrenceConstraintViolation rpc_framework_error_code = "OccurrenceConstraintViolation"

// Payload is syntactically correct but at least one field contains an invalid value
const PropertyConstraintViolation rpc_framework_error_code = "PropertyConstraintViolation"

// Payload for Action is not conform the PDU structure
const ProtocolError rpc_framework_error_code = "ProtocolError"

// Content of the call is not a valid RPC Request, for example: MessageId could not be read.
const RpcFrameworkError rpc_framework_error_code = "RpcFrameworkError"

// During the processing of Action a security issue occurred preventing receiver from completing the Action successfully
const SecurityError rpc_framework_error_code = "SecurityError"

// Payload for Action is syntactically correct but at least one of the fields violates data type constraints (e.g. "somestring": 12)
const TypeConstraintViolation rpc_framework_error_code = "TypeConstraintViolation"
