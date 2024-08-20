package main

import (
	"fmt"
	"juggleimsdk"
)

func main() {
	imsdk := juggleimsdk.NewJuggleIMSdk("{APPKEY}", "{SECRET}", "https://api.juggleim.com")
	fmt.Println(imsdk.Register(juggleimsdk.User{
		UserId:   "userid1",
		Nickname: "user1",
	}))
}
