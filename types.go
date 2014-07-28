package twigo

const (
	BASE_URL = "https://api.twilio.com/"
	API_VERSION = "2010-04-01"
	SMS_ENDPOINT = "/Messages.json"
	VOICE_ENDPOINT = "/Calls.json"
)

type Client struct {
	AccountSid string `required`
	AuthToken  string `required`
	Number 	   string `required`
}

type SMS struct {
	To   string `required`
	Body string `required`
}

type CALL struct {
	To  	 	     string `required` 
	Url	 	     string `required`
	Method 	 	     string `optional`
	FallbackUrl	     string `optional`
	FallbackMethod	     string `optional`
	StatusCallback 	     string `optional`
	StatusCallbackMethod string `optional`
	SendDigits	     string `optional`
	IfMachine	     string `optional`
	Timeout		     string `optional`
	Record		     string `optional`
}

// Reponses
type TwilioError struct {
	Status 	int  	`json:"status"`
	Message string  `json:"message"`
	Code 	int 	`json:"code"`
	Info	string	`json:"more_info"`
}

type MessageResponse struct {
	Sid             string  `json:"sid"`
        DateCreated     string  `json:"date_created"`
        DateUpdate      string  `json:"date_updated"`
        DateSent        string  `json:"date_sent"`
        AccountSid      string  `json:"account_sid"`
        To              string  `json:"to"`
        From            string  `json:"from"`
        Status          string  `json:"status"`
	Price           string  `json:"price"`
	Price_Unit	string	`json:"price_unit"`
	Direction	string  `json:"direction"`
	Uri		string	`json:"uri"`
}

type SMSResponse struct {
	MessageResponse
	NumSegments string   `json:"num_segments"`
        ErrorCode   string   `json:"error_code"`
        ErrorMsg    string   `json:"error_message"`
	Body        string   `json:"body"`
	Status      string   `json:"status"`
        ApiVersion  string   `json:"api_version"`
}

type CALLResponse struct {
	MessageResponse
	ParentCallSid 	string 	 `json:"parent_call_sid"`
	FormattedTo	string	 `json:"to_formatted"`
	FormattedFrom   string   `json:"from_formatted"`
	PhoneNumberSid  string   `json:"phone_number_sid"`
	StartTime	string	 `json:"start_time"`
	EndTime		string	 `json:"end_time"`
	Duration	string	 `json:"duration"`
	AnsweredBy	string	 `json:"answered_by"`
	ApiVersion	string	 `json:"api_version"`
	ForwardedFrom   string	 `json:"forwarded_from"`
	CallerName	string	 `json:"caller_name"`
}
