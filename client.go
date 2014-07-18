package twigo

func NewClient(account_sid, auth_token, number string) (*Client, error) {

	c := &Client{AccountSid:account_sid,AuthToken:auth_token,Number:number}
	return c, nil
}

func (c *Client) Text(msg_sms *SMS) (*TwilioResponse, *TwilioError) {

	resp, twil_err := Send(c, msg_sms)
	return resp, twil_err
}

func (c *Client) Call(msg_voice *Voice) (*TwilioResponse, *TwilioError) {

	resp, twil_err := Send(c, msg_voice)
	return resp, twil_err
}
