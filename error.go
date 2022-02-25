package goauth

import "errors"

var ErrNotImplement = errors.New("not implement yet")
var ErrInvalidInformation = errors.New("ivalid information")
var ErrInvalidSignature = errors.New("invalid signature")
var ErrClientNotFound = errors.New("client not found")
var ErrAccountNotFound = errors.New("account not found")

var ErrMempoolNotFound = errors.New("mempool not found")
var ErrSessionExpire = errors.New("session is expired")
