package apple

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewClient(agent string, teamID string, serviceID string, keyID string, privateKeyString []byte) (*Client, error) {
	block, _ := pem.Decode(privateKeyString)
	if block == nil {
		return nil, errors.New("empty block after decoding")
	}

	privatekey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	client := &Client{
		TeamID:     teamID,
		ServiceID:  serviceID,
		KeyID:      keyID,
		PrivateKey: privatekey,
		Agent:      agent,
	}
	return client, nil
}

type Client struct {
	KeyID      string
	TeamID     string
	ServiceID  string
	Agent      string
	PrivateKey interface{}
}

func (client *Client) PrintDebug() {
	fmt.Println("Client:", client.Agent)
	fmt.Println("\tTeamID:", client.TeamID)
	fmt.Println("\tServiceID:", client.ServiceID)
	fmt.Println("\tKeyID:", client.KeyID)
}
func (client *Client) GenerateSecret(durationSecond int64) (string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		Issuer:    client.TeamID,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Duration(durationSecond) * time.Second).Unix(),
		Audience:  "https://appleid.apple.com",
		Subject:   client.ServiceID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["alg"] = "ES256"
	token.Header["kid"] = client.KeyID

	return token.SignedString(client.PrivateKey)
}

func (client *Client) ValidateCode(code string) (*ValidateResponse, error) {

	secret, err := client.GenerateSecret(300)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("client_id", client.ServiceID)
	data.Set("client_secret", secret)
	data.Set("code", code)

	//data.Set("redirect_uri", "https://newcontinent-team.com/apple/signin")

	data.Set("grant_type", "authorization_code")
	url := "https://appleid.apple.com/auth/token"

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, strings.NewReader(data.Encode()))
	if err != nil {

		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", client.Agent)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	validateResponse := &ValidateResponse{}
	err = json.Unmarshal(resdata, res)
	if err != nil {
		return nil, err
	}

	if validateResponse.Error != "" {

		return validateResponse, errors.New(validateResponse.Error)
	}

	return validateResponse, nil
}
