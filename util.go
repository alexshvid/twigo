package twigo

const (
	BASE_URL = "https://api.twilio.com/"
	API_VERSION = "2010-04-01"
	SMS_ENDPOINT = "/SMS/Messages.json"
	VOICE_ENDPOINT = "/Calls.json"
)

type Client struct {
	AccountSid, AuthToken, Number string
}

type SMS struct {
	To, Body string
}

type Voice struct {
	To, Url string
}

type TwilioError struct {
	Status int  `json:"status"`
	Message string `json:"message"`
	Code int `json:"code"`
}

type TwilioResponse struct {
	Sid         string   `json:"sid"`
	DateCreated string   `json:"date_created"`
	DateUpdate  string   `json:"date_updated"`
	DateSent    string   `json:"date_sent"`
	AccountSid  string   `json:"account_sid"`
	NumSegments string   `json:"num_segments"`
	ErrorCode   string   `json:"error_code"`
	ErrorMsg    string   `json:"error_message"`
	To          string   `json:"to"`
	From        string   `json:"from"`
	Body        string   `json:"body"`
	Status      string   `json:"status"`
	Direction   string   `json:"direction"`
	ApiVersion  string   `json:"api_version"`
	Price       *float32 `json:"price,omitempty"`
	Url         string   `json:"uri"`
}
