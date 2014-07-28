#twigo

twigo is a go library wrapped around Twilio's REST APIs for initiating calls and sending SMS.

#Supported REST apis

##SMS:
```
1. POST /2010-04-01/Accounts/{AccountSid}/Messages.json
```
##Voice:
```
1. POST /2010-04-01/Accounts/{AccountSid}/Calls.json
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
        client, err := twigo.NewClient(<ACCOUNT_SID>, <AUTH_TOKEN>, <TWILIO_NUMBER>)

		if err != nil {
			fmt.Println(err)
		}

        // Create message
        message := &twigo.SMS{To:"+18001234567",Body:"Hello,World!"}

        // Send a Text
        twilio_response,err := client.Text(message)

		if err != nil {
			fmt.Println(err):
		}
}
```
