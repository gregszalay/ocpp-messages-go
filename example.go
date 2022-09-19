package main

import (
	"encoding/json"
	"fmt"

	"github.com/gregszalay/ocpp-go/types/BootNotificationRequest"
	"github.com/sanity-io/litter"
)

func createExampleJSONString() string {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	boot_req := BootNotificationRequest.BootNotificationRequestJson{
		Reason: "Unknown",
		ChargingStation: BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	// Create JSON string from example object
	boot_req_json_string, marshal_err := json.MarshalIndent(boot_req, "", "  ")
	if marshal_err != nil {
		fmt.Printf("Failed OCPP message json marshal. Error: %s", marshal_err)
	}
	return string(boot_req_json_string)
}

func main() {

	example_message := createExampleJSONString()
	fmt.Println("*******************************")
	fmt.Println("Example OCPP message json:")
	fmt.Println("*******************************")
	fmt.Printf("%s\n", example_message)
	fmt.Println("*******************************")

	// Unmarshal JSON string to see the types in action
	var req BootNotificationRequest.BootNotificationRequestJson
	unmarshal_err := req.UnmarshalJSON([]byte(example_message))
	if unmarshal_err != nil {
		fmt.Printf("Failed OCPP message json unmarshal. Error: %s", unmarshal_err)
	} else {
		fmt.Println("Example message unmarshalled using the types:")
		fmt.Println("*******************************")
		litter.Dump(req)
	}
	fmt.Println("*******************************")

}
