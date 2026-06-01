package main

import (
	"fmt"

	juggleimsdk "github.com/juggleim/imserver-sdk-go"
)

func main() {
	imsdk := juggleimsdk.NewJuggleIMSdk("appkey", "appsecret", "https://api.juggle.im")

	resp, code, trace, err := imsdk.Register(juggleimsdk.User{
		UserId: "userid1",
	})
	fmt.Println(resp, code, trace, err)
}
