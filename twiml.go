package twigo

import "encoding/xml"

type Response struct {
	XMLName  xml.Name  `xml:"Response"`
	Verbs    []Verb
}

type Verb interface{}

type Say struct {
	XMLName  xml.Name  `xml:"Say"`
	Voice    string    `xml:"voice,attr,omitempty"`
	Language string    `xml:"language,attr,omitempty"`
	Loop     int64     `xml:"loop,attr,omitempty"`
	Text     string    `xml:",innerxml"`
}

type Play struct {
 	XMLName  xml.Name  `xml:"Play"`
	Loop     int64     `xml:"loop,attr,omitempty"`
	Url      string    `xml:",innerxml"`
}

type Pause struct {
	XMLName  xml.Name  `xml:"Pause"`
	Length   int64     `xml:"length,attr,omitempty"`
}

type Reject struct {
	XMLName  xml.Name  `xml:"Reject"`
	Reason   string    `xml:"reason,attr,omitempty"`
}

func (resp *Response) Marshal(verbs...Verb) ([]byte,error) {

	for _,v := range verbs {
    		resp.Verbs = append(resp.Verbs,v)
  	}

	XmlBytes, err := xml.MarshalIndent(resp,"","  ")

	if err != nil {
 		return nil, err
  	}

	twiml := xml.Header + string(XmlBytes)
  	return []byte(twiml), nil
}
