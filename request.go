package twigo

import ( 
	"encoding/json"
	"net/url"
)

type TwilioRequest interface {
	GetUrl(string) string
	Headers() url.Values 
}

type TwilioResponse interface {
	Decode([]byte) interface{}
}

func (smsResponse *SMSResponse) Decode(respBody []byte) interface{} {

	json.Unmarshal(respBody,smsResponse)
	return smsResponse 
}

func (callResponse *CALLResponse) Decode(respBody []byte) interface{} {

	json.Unmarshal(respBody,callResponse)
	return callResponse
}

func (sms *SMS) GetUrl(acctSid string) string {
	return GetRestUrl(acctSid, SMS_ENDPOINT)
}

func (sms *SMS) Headers() url.Values {
	params := url.Values{}
	
	params.Set("To",sms.To)
	params.Set("Body",sms.Body)

	return params
}

func (call *CALL) GetUrl(acctSid string) string {
	return GetRestUrl(acctSid, VOICE_ENDPOINT)
}

func (call *CALL) Headers() url.Values {
	params := url.Values{}

	params.Set("To",call.To)
	params.Set("Url",call.Url)

	return params
} 
