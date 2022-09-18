package main

import (
	"encoding/json"
	"fmt"

	"github.com/chargerevolution/ocpp-go/types/BootNotificationRequest"
)

func main() {

	// Create example OCPP message object (e.g. BootNotificationRequest)
	boot_req := BootNotificationRequest.BootNotificationRequestJson{
		Reason: "Unknown",
		ChargingStation: BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	// Create JSON string from object
	boot_req_json_string, marshal_err := json.Marshal(boot_req)
	if marshal_err != nil {
		fmt.Printf("Failed OCPP message json unmarshal. Error: %s", marshal_err)
	}
	// Print result
	fmt.Printf("Example OCPP message successfully marshalled.\nJSON:\n%s\n", boot_req_json_string)

	// Unmarshal JSON string to OCPP object using the types
	var req BootNotificationRequest.BootNotificationRequestJson
	unmarshal_err := req.UnmarshalJSON([]byte(boot_req_json_string))
	if unmarshal_err != nil {
		fmt.Printf("Failed OCPP message json unmarshal. Error: %s", unmarshal_err)
	}
	// Print result
	fmt.Printf("Example OCPP message successfully unmarshalled.\nObject:\n%+v\n", req)

}
