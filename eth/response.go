package eth

type Response struct {
	MessageTitle  string `json:"MessageTitle"` //the title use in typed sign
	VerifyMessage string `json:"VerifyMessage"`
	Signature     string `json:"Signature"`
}
