<div align="center">

# JuggleIM Server SDK for Go

**Build reliable messaging backends in Go with a small, straightforward SDK.**

Users, conversations, messages, groups, chatrooms, history, moderation, bots, moments, push, and more—through one server-side client.

[![Go Reference](https://pkg.go.dev/badge/github.com/juggleim/imserver-sdk-go.svg)](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go)
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/juggleim/imserver-sdk-go?style=social)](https://github.com/juggleim/imserver-sdk-go/stargazers)

[English](README.md) · [简体中文](README.zh-CN.md) · [API Reference](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go) · [Report an issue](https://github.com/juggleim/imserver-sdk-go/issues)

</div>

---

`imserver-sdk-go` is a Go client for integrating backend services with the JuggleIM Server API. It keeps authentication and HTTP details in one place and exposes typed methods for common instant-messaging workflows.

> This is a **server-side SDK**. Keep your App Secret on trusted backend infrastructure and never ship it in a web, desktop, or mobile client.

### Why use it?

- **Broad IM coverage** — manage the complete lifecycle of users, conversations, messages, groups, and chatrooms.
- **Go-native API** — typed request and response models with predictable return values.
- **Operational visibility** — every request returns an API code and trace ID for troubleshooting.
- **Small dependency surface** — only one direct third-party dependency.
- **Backend ready** — use the SDK from API services, workers, bots, admin systems, or migration tools.

If this SDK saves you time, please [give the project a Star](https://github.com/juggleim/imserver-sdk-go). It helps more Go developers discover and improve it. ⭐

### Feature coverage

| Area | Capabilities |
| --- | --- |
| Users | Register, update, query, settings, online status, bans, kicks, and blocks |
| Messages | Private, system, group, chatroom, public-channel, broadcast, group-cast, and streaming messages |
| Conversations | Query, add, delete, clear unread, pin, mute, local tags, and global tags |
| Groups | Create, update, dissolve, membership, permissions, settings, and mute controls |
| Chatrooms | Lifecycle, member checks, mute/ban/allow lists, and custom attributes |
| History | Query, recall, delete, clean, modify, and import historical messages |
| Social | Friends, moments, comments, and reactions |
| Operations | Bots, device binding, sensitive words, global mute lists, and connection keys |
| Push | User tags and tag-targeted push notifications |
| Public channels | Lifecycle, subscriptions, members, and messaging |

### Installation

The module currently targets Go 1.23 or later.

```bash
go get github.com/juggleim/imserver-sdk-go
```

### Quick start

Create an SDK client with credentials from your JuggleIM application:

```go
package main

import (
	"fmt"
	"log"

	juggleimsdk "github.com/juggleim/imserver-sdk-go"
)

func main() {
	sdk := juggleimsdk.NewJuggleIMSdk(
		"your-app-key",
		"your-app-secret",
		"https://api.juggle.im",
	)

	user, code, traceID, err := sdk.Register(juggleimsdk.User{
		UserId:   "alice",
		Nickname: "Alice",
	})
	if err != nil {
		log.Fatalf("register user: code=%d trace_id=%s err=%v", code, traceID, err)
	}

	fmt.Printf("registered user=%s\n", user.UserId)
}
```

Send a private message:

```go
code, traceID, err := sdk.SendPrivateMsg(juggleimsdk.Message{
	SenderId:   "alice",
	ReceiverId: "bob",
	MsgType:    "jg:text",
	MsgContent: `{"content":"Hello from Go!"}`,
})
if err != nil {
	log.Printf("send message: code=%d trace_id=%s err=%v", code, traceID, err)
}
```

See [`examples/main.go`](examples/main.go) for a minimal runnable example and [pkg.go.dev](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go) for all exported types and methods.

### Response convention

SDK methods return business data when applicable, followed by three common values:

```go
response, code, traceID, err := sdk.Register(user)
```

| Value | Meaning |
| --- | --- |
| `response` | Typed response data; only returned by APIs that have response content |
| `code` | JuggleIM API result code; `0` means success |
| `traceID` | Client-generated request ID, useful for correlating entries in application logs |
| `err` | Transport, decoding, unsupported-method, or server API error |

Always check `err`. Include `code` and `traceID` in your application logs so failures can be diagnosed quickly.

### Global conversation tags

Set tags shared by every participant in a conversation:

```go
tags := []string{"support", "priority"}
code, traceID, err := sdk.SetGlobalConverTags(juggleimsdk.GlobalConverTagsReq{
	ConverId:         "group-1001",
	ChannelType:      2,
	GlobalConverTags: &tags,
})
```

Pass a pointer to an empty slice to clear all global tags. Leaving `GlobalConverTags` as `nil` omits the field and is rejected by the set endpoint.

### Contributing

Issues, documentation improvements, examples, tests, and API coverage are welcome.

1. Fork the repository and create a focused branch.
2. Make your change and run `gofmt` on modified Go files.
3. Run `go test ./...`.
4. Open a pull request describing the problem and the solution.

For bugs and feature requests, use [GitHub Issues](https://github.com/juggleim/imserver-sdk-go/issues) or [Gitee Issues](https://gitee.com/juggleim/imserver-sdk-go/issues).

### License

Licensed under the [Apache License 2.0](LICENSE).

---

<div align="center">

**Built for Go teams shipping real-time communication.**

[⭐ Star on GitHub](https://github.com/juggleim/imserver-sdk-go) · [Gitee Mirror](https://gitee.com/juggleim/imserver-sdk-go) · [API Reference](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go)

</div>
