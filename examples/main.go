package main

import (
	"fmt"

	juggleimsdk "github.com/juggleim/imserver-sdk-go"
)

func main() {
	imsdk := juggleimsdk.NewJuggleIMSdk("nwm6fxqt2aeebhb7", "gcvnz2t7405ub4xc", "https://api.juggle.im")

	resp, code, trace, err := imsdk.Register(juggleimsdk.User{
		UserId: "userid1",
	})
	fmt.Println(resp, code, trace, err)

	resp, code, trace, err = imsdk.RegisterBot(juggleimsdk.BotInfo{
		BotId:    "botid1",
		Nickname: "botnickname1",
		Portrait: "botportrait1",
	})
	fmt.Println(resp, code, trace, err)
}
