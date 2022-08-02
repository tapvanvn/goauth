package momo

type MiniAppAuthResponse struct {
	MiniAppUserID string `json:"MiniAppUserID"`
	PartnerUserID string `json:"PartnerUserID"`
	AuthCode      string `json:"AuthCode"`
}
