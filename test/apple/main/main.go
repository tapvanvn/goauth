package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tapvanvn/goauth/apple"
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

	code := "c0410595558ce4ee3b4218b728e899c6a.0.rrtuy.YGwfUcCgiiSL5KRdKT-rOQ"
	_, err = client.ValidateCode(code)

	if err != nil {

		log.Fatal("validate fail", err)
	}
}
