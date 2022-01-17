package goauth

import "errors"

var ErrNotImplement = errors.New("Not implement yet")
var ErrInvalidInfomation = errors.New("Invalid infomation")
var ErrInvalidSignature = errors.New("Invalid signature")
var ErrClientNotFound = errors.New("Client not found")
var ErrAccountNotFound = errors.New("Account not found")

var ErrMempoolNotFound = errors.New("Mempool not found")
var ErrSessionExpire = errors.New("Session is expired")
