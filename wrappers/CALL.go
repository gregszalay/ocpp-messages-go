package wrappers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CALL struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId uint32
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// The name of the remote procedure or action. This field SHALL contain a case-sensitive string.
	// The field SHALL contain the OCPP Message name without the "Request" suffix. For example: For
	// a "BootNotificationRequest", this field shall be set to "BootNotification".
	Action string
	// JSON Payload of the action, see: JSON Payload for more information.
	Payload interface{}
}

func (j *CALL) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	err := json.Unmarshal([]byte(b), &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	if len(raw) != 4 {
		return fmt.Errorf("invalid CALL message length")
	}
	// MessageTypeId integer This is a Message Type Number which is used to identify the type of the message.
	message_type_id, err_message_type_id := raw[0].(float64)
	if !err_message_type_id {
		return fmt.Errorf("CALL data[0] is not a number")
	}
	// MessageId string[36] This is a unique identifier that will be used to match request and result.
	message_id, err_message_id := raw[1].(string)
	if !err_message_id || message_id == "" {
		return fmt.Errorf("CALL data[1] is not a string")
	}
	// Action string The name of the remote procedure or action. This field SHALL contain a case-sensitive string.
	// The field SHALL contain the OCPP Message name without the "Request" suffix. For example: For a "BootNotificationRequest", this field shall be set to "BootNotification".
	action, err_action := raw[2].(string)
	if !err_action || action == "" {
		return fmt.Errorf("CALL data[2] is not a string")
	}
	// Payload
	payload, err_payload := raw[3].(map[string]interface{})
	if !err_payload || payload == nil {
		return fmt.Errorf("CALL data[3] is not a map[string]interface{}")
	}
	*j = CALL{
		MessageTypeId: uint32(message_type_id),
		MessageId:     message_id,
		Action:        action,
		Payload:       payload,
	}
	return nil
}

func (c *CALL) Marshal() []byte {
	message_array := [...]interface{}{CALL_TYPE, c.MessageId, c.Action, c.Payload}
	result, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Could not marshal CALL message: %s\n", err)
		return []byte("")
	}
	return result
}

func (c *CALL) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}

func (c *CALL) GetPayloadAsJSON() []byte {
	if c == nil {
		fmt.Println("CALL object is empty")
		return []byte("")
	}
	// Re-marshal payload only
	re_marshalled_payload, re_marshall_err := json.MarshalIndent(c.Payload, "", " ")
	if re_marshall_err != nil {
		fmt.Println("Failed to remarshall CALL pazload to json")
		return []byte("")
	}

	return re_marshalled_payload
}
