package eth

type Response struct {
	MessageTitle string `json:"MessageTitle"` //the title use in typed sign
	Signature    string `json:"Signature"`
}
