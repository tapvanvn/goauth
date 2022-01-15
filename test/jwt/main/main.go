package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tapvanvn/goauth/jwt"
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
	"github.com/tapvanvn/godbengine/engine/adapter"
	"golang.org/x/oauth2/jws"
)

func startEngine(eng *engine.Engine) {

	localMem := &adapter.LocalMemDB{}
	localMem.Init("")
	eng.Init(localMem, nil, nil)
}
func main() {
	engines.InitEngineFunc = startEngine
	eng := engines.GetEngine()

	reader := rand.Reader
	bitSize := 2048

	privateKey, err := rsa.GenerateKey(reader, bitSize)

	if err != nil {

		panic(err)
	}
	refreshTokenLifeTime := time.Second * 3
	tokenLifeTime := time.Second * 2
	client, err := jwt.NewClient(eng, "jwt", tokenLifeTime, refreshTokenLifeTime, privateKey)

	if err != nil {

		panic(err)
	}

	jwtString, err := client.BeginSession("acc_10", nil)
	if err != nil {

		panic(err)
	}
	fmt.Println("jwtString", jwtString.GetSessionID())

	if success, err := client.Verify(jwtString, nil, nil); err == nil {
		if success {
			fmt.Println("verify success")
		} else {
			fmt.Println("verify fail")
			os.Exit(1)
		}
	} else {
		panic(err)
	}
	time.Sleep(time.Second * 2)
	claim, err := jws.Decode(string(jwtString.GetSessionID()))
	parts := strings.Split(claim.Aud, ".")
	refreshToken := parts[len(parts)-1]
	if jwtString2, err := client.RenewSession(refreshToken); err == nil {

		fmt.Println("new jwt:", jwtString2.GetSessionID())
	} else {
		panic(err)
	}
	time.Sleep(time.Second * 4)
	if jwtString3, err := client.RenewSession(refreshToken); err == nil {
		fmt.Println("[wrong] new jwt:", jwtString3.GetSessionID())
	} else {
		fmt.Println("[right] cannot renew when exprite")
	}
}
