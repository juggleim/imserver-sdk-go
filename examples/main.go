package main

import (
	"fmt"

	juggleimsdk "github.com/juggleim/imserver-sdk-go"
)

func main() {
	imsdk := juggleimsdk.NewJuggleIMSdk("appkey", "appsecret", "http://127.0.0.1:9001")

	resp, code, trace, err := imsdk.Register(juggleimsdk.User{
		UserId: "userid1",
	})
	fmt.Println(resp, code, trace, err)
}
