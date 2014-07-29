package twigo

import (
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
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

func Send(client *Client, request TwilioRequest, response TwilioResponse) (error) {
       	var twilio_url string
	var url_params url.Values	

        twilio_url = request.GetUrl(client.AccountSid)
	url_params = request.Headers()
        url_params.Set("From",client.Number)
	
	req, err := http.NewRequest("POST",twilio_url,strings.NewReader(url_params.Encode()))

	if err != nil {
		return err
	}

	req.SetBasicAuth(client.AccountSid,client.AuthToken)
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	req.Header.Add("Accept","application/json")

	http_client := &http.Client{}
        resp, err := http_client.Do(req)

	if err != nil {
		return err
	}

        defer resp.Body.Close()

        resp_body, _ := ioutil.ReadAll(resp.Body)

	var twilio_error TwilioError
        
	if resp.StatusCode != http.StatusCreated {
                json.Unmarshal(resp_body,&twilio_error)
		err = twilioErrToError(&twilio_error)
		return err
	} else {
		err = response.Decode(resp_body)

		if err != nil {
			return err
		}
	}

        return nil
}

func twilioErrToError(twilioError *TwilioError) error {
	return fmt.Errorf("Status: %d, Message: %s",twilioError.Status,twilioError.Message)	
}

func Validate(params interface{}) (error) {

        var object interface{}

        switch t := params.(type) {
                case Client, SMS, CALL:
                        object = t
                default:
                        return fmt.Errorf("Assertion Error -- Unknown Type")
        }

        for k := 0; k < reflect.ValueOf(object).NumField(); k++ {
                field := reflect.ValueOf(object).Type().Field(k)
                value := reflect.ValueOf(object).Field(k).String()

                if value == "" && field.Tag == "required" {
                        return fmt.Errorf("Missing Required Field: %s", field.Name)
                }
        }

        return nil
}
