package main

import (
	"fmt"

	juggleimsdk "github.com/juggleim/imserver-sdk-go"
)

func main() {
	imsdk := juggleimsdk.NewJuggleIMSdk("appkey", "appsecret", "http://127.0.0.1:8082")

	// resp, code, trace, err := imsdk.Register(juggleimsdk.User{
	// 	UserId: "userid1",
	// })
	resp, code, trace, err := imsdk.QryHisMsgs("userid1", "groupid1", juggleimsdk.ChannelType_Group, 0, 10, false)
	fmt.Println(resp, code, trace, err)
	fmt.Println(juggleimsdk.ToJson(resp))
	fmt.Println(len(resp.Msgs))
}
