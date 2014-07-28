package twigo

import (
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
	"bytes"
	"fmt"
)

func GetRestUrl(acctSid, endpoint string) (string) {
        var buffer bytes.Buffer
        var url string

        buffer.WriteString(BASE_URL)
        buffer.WriteString(API_VERSION)
        buffer.WriteString("/Accounts/")
        buffer.WriteString(acctSid)
        buffer.WriteString(endpoint)

        url = buffer.String()
        return url
}

func Send(client *Client, request TwilioRequest, response TwilioResponse) (interface{}, error) {
       	var twilio_url string
	var url_params url.Values	

        twilio_url = request.GetUrl(client.AccountSid)
	url_params = request.Headers()
        url_params.Set("From",client.Number)
	
	req, _ := http.NewRequest("POST",twilio_url,strings.NewReader(url_params.Encode()))
        req.SetBasicAuth(client.AccountSid,client.AuthToken)
        req.Header.Add("Content-Type","application/x-www-form-urlencoded")
        req.Header.Add("Accept","application/json")

	http_client := &http.Client{}
        resp, _ := http_client.Do(req)

        defer resp.Body.Close()

        resp_body, _ := ioutil.ReadAll(resp.Body)

	var twilio_error TwilioError
	var twilio_response interface{}
	var err error 

        if resp.StatusCode != http.StatusCreated {
                json.Unmarshal(resp_body,&twilio_error)
		err = twilioErrToError(&twilio_error)
	} else {
        	twilio_response = response.Decode(resp_body)
	}

        return twilio_response,err
}

func twilioErrToError(twilioError *TwilioError) error {
	return fmt.Errorf("Status: %d, Message: %s",twilioError.Status,twilioError.Message)	
}
