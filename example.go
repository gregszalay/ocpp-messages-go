package main

import (
	"github.com/gregszalay/ocpp-messages-go/examples"
	"github.com/sanity-io/litter"
)

func main() {
	litter.Config.StripPackageNames = true

	// CALL example
	call_json_string := examples.Create_call_json_string()
	examples.Run_call_unmarshal_example(call_json_string)

	// CALLRESULT example
	callresult_json_string := examples.Create_callresult_json_string()
	examples.Run_callresult_unmarshal_example(callresult_json_string)

	// CALLERROR
	callerror_json_string := examples.Create_callerror_json_string()
	examples.Run_callerror_unmarshal_example(callerror_json_string)

}
