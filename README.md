#twigo

twigo is a go library wrapped around Twilio's REST APIs for initiating calls and sending SMS.

#Supported REST apis

##SMS:
```
1. POST /2010-04-01/Accounts/{AccountSid}/Messages.json
2. POST /2010-04-01/Accounts/{AccountSid}/Calls.json
```

#Usage

##Send a Text
```
package main

import (
	"github.com/fibonacci1729/twigo"
	"encoding/json"
	"fmt"
	)

func main() {
   
        // Initiate new Twilio Session
        client, _ := twigo.NewClient(<ACCOUNT_SID>, <AUTH_TOKEN>, <TWILIO_NUMBER>)

        // Create message
        message := &twigo.SMS{To:"+18001234567",Body:"Hello,World!"}

	// Send a Text
        twilio_response,twilio_error := client.Text(message)

        // Pretty print Twilio Response
        fmt.Println("Twilio Response: ")
        b,_ := json.MarshalIndent(twilio_response," ","   ")
        fmt.Println(string(b))

        // Pretty print Twilio Error Response
        fmt.Println("Twilio Error: ")
        b1,_ := json.MarshalIndent(twilio_error," ","   ")
        fmt.Println(string(b1))
}
```
