package twigo

import (
	"encoding/json"
)

type TwilioResponse interface {
	Decode([]byte) 
}

func (smsResponse *SMSResponse) Decode(respBody []byte) {

        json.Unmarshal(respBody,smsResponse)
}

func (callResponse *CALLResponse) Decode(respBody []byte) {

        json.Unmarshal(respBody,callResponse)
}
