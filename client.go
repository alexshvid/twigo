package twigo

func NewClient(account_sid, auth_token, number string) (*Client, error) {

	c := &Client{AccountSid:account_sid,AuthToken:auth_token,Number:number}

	err := Validate(*c)

	if err != nil {
		return nil,err
	}

	return c, nil
}

func (c *Client) Text(msg_sms *SMS) (*SMSResponse, error) {

	err := Validate(*msg_sms)

	if err != nil {
		return nil, err
	}

	smsResponse := &SMSResponse{}

	err = Send(c, msg_sms, smsResponse)
	return smsResponse, err 
}

func (c *Client) Call(msg_voice *CALL) (*CALLResponse, error) {

	err := Validate(*msg_voice)

	if err != nil {
		return nil, err
	}

	callResponse := &CALLResponse{}

	err = Send(c, msg_voice, callResponse)
	return callResponse, err
}
