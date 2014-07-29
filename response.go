package twigo

import (
	"encoding/json"
)

type TwilioResponse interface {
	Decode([]byte) (error)
}

func (smsResponse *SMSResponse) Decode(respBody []byte) (error) {

        err := json.Unmarshal(respBody,smsResponse)

	if err != nil {
		return err
	}

	return nil
}

func (callResponse *CALLResponse) Decode(respBody []byte) (error) {

        err := json.Unmarshal(respBody,callResponse)

	if err != nil {
		return err
	}

	return nil
}
