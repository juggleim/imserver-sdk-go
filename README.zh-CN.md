<div align="center">

# JuggleIM Go 服务端 SDK

**使用轻量、直观的 Go SDK 构建可靠的即时通讯后端。**

一个服务端客户端即可覆盖用户、会话、消息、群组、聊天室、历史消息、内容治理、机器人、动态和推送等能力。

[![Go Reference](https://pkg.go.dev/badge/github.com/juggleim/imserver-sdk-go.svg)](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go)
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/juggleim/imserver-sdk-go?style=social)](https://github.com/juggleim/imserver-sdk-go/stargazers)

[English](README.md) · [简体中文](README.zh-CN.md) · [API 参考](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go) · [反馈问题](https://github.com/juggleim/imserver-sdk-go/issues)

</div>

---

`imserver-sdk-go` 是用于后端服务对接 JuggleIM Server API 的 Go SDK。它统一处理鉴权和 HTTP 请求细节，并通过类型明确的方法覆盖常见的即时通讯业务。

> 这是一个**服务端 SDK**。请仅在可信的后端环境中保存 App Secret，切勿将其内置到 Web、桌面或移动客户端中。

## 为什么选择它？

- **即时通讯能力覆盖广**：覆盖用户、会话、消息、群组和聊天室的完整生命周期。
- **符合 Go 使用习惯**：提供类型明确的请求、响应模型以及统一的返回值。
- **便于排查问题**：每次请求都会返回 API 状态码和 Trace ID。
- **依赖精简**：仅包含一个直接的第三方依赖。
- **适合多种后端场景**：可用于 API 服务、异步任务、机器人、运营后台和数据迁移工具。

如果这个 SDK 对你有帮助，欢迎[点亮 Star](https://github.com/juggleim/imserver-sdk-go)。你的支持能让更多 Go 开发者发现并参与完善这个项目。⭐

## 功能范围

| 模块 | 能力 |
| --- | --- |
| 用户 | 注册、更新、查询、设置、在线状态、封禁、踢出和黑名单 |
| 消息 | 单聊、系统、群聊、聊天室、公共频道、广播、群定向和流式消息 |
| 会话 | 查询、新增、删除、清未读、置顶、免打扰、普通标签和全局标签 |
| 群组 | 创建、更新、解散、成员管理、权限、设置和禁言 |
| 聊天室 | 生命周期、成员检查、禁言/封禁/白名单和自定义属性 |
| 历史消息 | 查询、撤回、删除、清理、修改和导入 |
| 社交 | 好友、动态、评论和表态 |
| 运营管理 | 机器人、设备绑定、敏感词、全局禁言和连接密钥 |
| 推送 | 用户标签和基于标签的定向推送 |
| 公共频道 | 生命周期、订阅、成员管理和消息发送 |

## 安装

当前模块要求 Go 1.23 或更高版本。

```bash
go get github.com/juggleim/imserver-sdk-go
```

## 快速开始

使用 JuggleIM 应用的凭证创建 SDK 客户端：

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

发送一条单聊消息：

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

最小可运行示例位于 [`examples/main.go`](examples/main.go)，完整的导出类型和方法可查看 [pkg.go.dev](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go)。

## 返回值约定

如果接口包含业务数据，SDK 方法会先返回响应对象，随后返回三个通用值：

```go
response, code, traceID, err := sdk.Register(user)
```

| 返回值 | 含义 |
| --- | --- |
| `response` | 类型明确的响应数据，仅在接口存在响应内容时返回 |
| `code` | JuggleIM API 状态码，`0` 表示成功 |
| `traceID` | 客户端生成的请求 ID，可用于关联应用日志 |
| `err` | 网络、解码、不支持的方法或服务端 API 错误 |

请始终检查 `err`，并在业务日志中记录 `code` 和 `traceID`，以便快速定位问题。

## 全局会话标签

为会话中的所有参与者设置共享标签：

```go
tags := []string{"support", "priority"}
code, traceID, err := sdk.SetGlobalConverTags(juggleimsdk.GlobalConverTagsReq{
	ConverId:         "group-1001",
	ChannelType:      2,
	GlobalConverTags: &tags,
})
```

传入空切片的指针可以清空全部全局标签。将 `GlobalConverTags` 保持为 `nil` 会省略该字段，并被设置接口拒绝。

## 参与贡献

欢迎提交问题、文档改进、示例、测试以及新的 API 能力。

1. Fork 仓库并创建职责单一的分支。
2. 完成修改，对变更的 Go 文件运行 `gofmt`。
3. 运行 `go test ./...`。
4. 创建 Pull Request，并说明问题和解决方案。

Bug 和功能建议可以提交到 [GitHub Issues](https://github.com/juggleim/imserver-sdk-go/issues) 或 [Gitee Issues](https://gitee.com/juggleim/imserver-sdk-go/issues)。

## 开源协议

本项目基于 [Apache License 2.0](LICENSE) 开源。

---

<div align="center">

**为构建实时通讯产品的 Go 团队而生。**

[⭐ 在 GitHub 点亮 Star](https://github.com/juggleim/imserver-sdk-go) · [Gitee 镜像](https://gitee.com/juggleim/imserver-sdk-go) · [API 参考](https://pkg.go.dev/github.com/juggleim/imserver-sdk-go)

</div>
