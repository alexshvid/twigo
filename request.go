package twigo

import ( 
	"net/url"
)

type TwilioRequest interface {
	GetUrl(string) string
	Headers() url.Values 
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
