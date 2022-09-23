package wrappers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CALLRESULT struct {
	// This is a Message Type Number which is used to identify the type of the message.
	MessageTypeId uint32
	// This must be the exact same id that is in the call request so that the recipient can match request and result.
	MessageId string
	// JSON Payload of the action, see: JSON Payload for more information.
	Payload interface{}
}

func (j *CALLRESULT) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	err := json.Unmarshal([]byte(b), &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	if len(raw) != 3 {
		return fmt.Errorf("invalid CALLRESULT message length")
	}
	// MessageTypeId integer This is a Message Type Number which is used to identify the type of the message.
	message_type_id, message_type_id_err := raw[0].(float64)
	if !message_type_id_err {
		return fmt.Errorf("CALLRESULT data[0] is not a number")
	}
	// MessageId string[36] This is a unique identifier that will be used to match request and result.
	message_id, message_id_err := raw[1].(string)
	if !message_id_err || message_id == "" {
		return fmt.Errorf("CALLRESULT data[1] is not a string")
	}
	// Payload
	payload, payload_err := raw[2].(map[string]interface{})
	if !payload_err || payload == nil {
		return fmt.Errorf("CALLRESULT data[2] is not a map[string]interface{}")
	}
	*j = CALLRESULT{
		MessageTypeId: uint32(message_type_id),
		MessageId:     message_id,
		Payload:       payload,
	}
	return nil
}

func (c *CALLRESULT) Marshal() []byte {
	message_array := [...]interface{}{CALLRESULT_TYPE, c.MessageId, c.Payload}
	message_array_json, err := json.Marshal(message_array)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALLRESULT message: %s\n", err)
		return []byte("")
	}
	return message_array_json
}

func (c *CALLRESULT) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}

func (c *CALLRESULT) GetPayloadAsJSON() []byte {
	if c == nil {
		fmt.Println("CALLRESULT object is empty")
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
