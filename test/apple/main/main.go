package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goauth/apple"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	TeamID           string `json:"TeamID"`
	ServiceID        string `json:"ServiceID"`
	KeyID            string `json:"KeyID"`
	PrivateKeyBase64 string `json:"PrivateKeyBase64"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("test.sh <config>")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyData, err := base64.RawStdEncoding.DecodeString(config.PrivateKeyBase64)
	if err != nil {
		log.Fatal(err)
	}
	client, err := apple.NewClient("test", config.TeamID, config.ServiceID, config.KeyID, privateKeyData)
	if err != nil {
		log.Fatal("create client fail", err)
	}
	client.PrintDebug()
	//token := "eyJraWQiOiJlWGF1bm1MIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLm5ld2NvbnRpbmVudC10ZWFtLnRlc3QuYXBwbGVzaWduaW4iLCJleHAiOjE2MzUxODMxNzUsImlhdCI6MTYzNTA5Njc3NSwic3ViIjoiMDAxMzQ4LjU5MDk0M2Y4MjA5NTQ1MTZhMGU3N2QyMGI3ODQyNWQyLjA4MjkiLCJjX2hhc2giOiJ6OFNQaWZNVWtGWk82MG55dWhiVzNnIiwiZW1haWwiOiJwNHhkcnhwZm1xQHByaXZhdGVyZWxheS5hcHBsZWlkLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjoidHJ1ZSIsImlzX3ByaXZhdGVfZW1haWwiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjM1MDk2Nzc1LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.HkuowNO4QBFM3WiKhFXuXM5BUfqHSAhrl5lVdv0tBiJgMx0LARM2zHhEcEe7UiVCdqtFZgwolpSunvtnuqMaFgNLycf6HkGrU_EwSxAzvjP3x44ieHEy7eB_Yh7i3cIVtL1k1S6fsmzFQZwYKGomxYgjxNl02_SyHZaUXf7CT_hw4uDtomYgLpi6W2opu9f8X2w7rxYFRotsALmhhIxFiX3uBj8j9Cg6I6s8FKKDpL4tOrKlaP1LktxWhyb52EpVJT2iPuWIt2WjIqzV-xhPZXBhBc4YMiQR4CcLcCy90rUogp5wJB7y6aiU9pW-tZ-1hZe969J40fUhPrzoHwW7-A"
	token := "c0410595558ce4ee3b4218b728e899c6a.0.rrtuy.YGwfUcCgiiSL5KRdKT-rOQ"
	err = client.ValidateToken(token)

	if err != nil {

		log.Fatal("validate fail", err)
	}
}
