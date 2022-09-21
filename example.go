package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationRequest"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	"github.com/sanity-io/litter"
)

func createExampleJSONString() []byte {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	boot_req := BootNotificationRequest.BootNotificationRequestJson{
		Reason: "PowerUp",
		ChargingStation: BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	// All OCPP messages need to be wrapped into a CALL, CALLRESULT or CALLERROR type
	call_wrapper := wrappers.CALL{
		MessageTypeId: wrappers.CALL_TYPE,
		MessageId:     uuid.NewString(),
		Action:        "BootNotification",
		Payload:       boot_req,
	}

	return call_wrapper.Marshal()
}

func main() {

	example_message := createExampleJSONString()
	fmt.Println("*******************************")
	fmt.Println("Example OCPP message json:")
	fmt.Println("*******************************")
	fmt.Printf("%s\n", example_message)
	fmt.Println("*******************************")

	// Unmarshal CALL message
	var call wrappers.CALL
	call_unmarshal_err := call.UnmarshalJSON(example_message)
	if call_unmarshal_err != nil {
		fmt.Printf("Failed OCPP message json unmarshal. Error: %s", call_unmarshal_err)
	} else {
		fmt.Println("Example CALL message unmarshalled using the types:")
		fmt.Println("*******************************")
		litter.Dump(call)
	}
	fmt.Println("*******************************")

	// Re-marshal payload only
	re_marshalled_payload, re_marshall_err := json.MarshalIndent(call.Payload, "", " ")
	if re_marshall_err != nil {
		fmt.Println(re_marshall_err)
	}

	// Unmarshal payload
	var req BootNotificationRequest.BootNotificationRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(re_marshalled_payload)
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed OCPP message json unmarshal. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Example payload unmarshalled using the types:")
		fmt.Println("*******************************")
		litter.Dump(req)
	}
	fmt.Println("*******************************")

}
