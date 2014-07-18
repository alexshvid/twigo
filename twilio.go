package twigo

import (
	"net/url"
	"net/http"
	"bytes"
	"fmt"
	"strings"
	"encoding/json"
	"io/ioutil"
)

func Send(client *Client, msg interface{}) (*TwilioResponse, *TwilioError) {

	var twurl *string
	var err error

	params := url.Values{}

 	if m, ok := msg.(*SMS); ok {
		twurl, _ = GetRestUrl(client,SMS_ENDPOINT)
		params.Set("From",client.Number)
		params.Set("To",m.To)
		params.Set("Body",m.Body)

	} else if _, ok := msg.(*Voice); ok {
		twurl, _ = GetRestUrl(client,VOICE_ENDPOINT)
	} else {
		fmt.Errorf("Uknown message type")
		return nil,nil
	}

	if err != nil {
		fmt.Errorf("%s",err)
		return nil,nil
	}

	req, _ := http.NewRequest("POST",*twurl,strings.NewReader(params.Encode()))
	req.SetBasicAuth(client.AccountSid,client.AuthToken)
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	req.Header.Add("Accept","application/json")

	http_client := &http.Client{}
	resp, _ := http_client.Do(req)
	
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	var twilio_error *TwilioError
	var twilio_response *TwilioResponse

	if resp.StatusCode != http.StatusCreated {
		twilio_error = &TwilioError{}
		json.Unmarshal(resp_body,twilio_error)
	} else {
		twilio_response = &TwilioResponse{}
		json.Unmarshal(resp_body,twilio_response)
	}

	return twilio_response,twilio_error
}

func GetRestUrl(client *Client, end_point string) (*string, error) {
	var buffer bytes.Buffer
	var url string

	buffer.WriteString(BASE_URL)
	buffer.WriteString(API_VERSION)
	buffer.WriteString("/Accounts/")
	buffer.WriteString(client.AccountSid)
	buffer.WriteString(end_point)

	url = buffer.String()
	return &url, nil
}
