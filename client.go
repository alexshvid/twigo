package twigo

func NewClient(account_sid, auth_token, number string) (*Client, error) {

	c := &Client{AccountSid:account_sid,AuthToken:auth_token,Number:number}

	err := Validate(*c)

	if err != nil {
		return nil,err
	}

	return c, nil
}

func (c *Client) Text(msg_sms *SMS) (*TwilioResponse, *TwilioError, error) {

	err := Validate(*msg_sms)

	if err != nil {
		return nil, nil, err
	}

	resp, twil_err := Send(c, msg_sms)
	return resp, twil_err, nil
}

func (c *Client) Call(msg_voice *Voice) (*TwilioResponse, *TwilioError, error) {

	err := Validate(*msg_voice)

	if err != nil {
		return nil, nil, err
	}

	resp, twil_err := Send(c, msg_voice)
	return resp, twil_err, nil
}
