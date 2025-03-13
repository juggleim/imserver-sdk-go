package juggleimsdk

import "net/http"

type BotInfo struct {
	BotId     string            `json:"bot_id"`
	Nickname  string            `json:"nickname"`
	Portrait  string            `json:"portrait"`
	BotType   *int              `json:"bot_type"`
	BotConf   string            `json:"bot_conf"`
	Webhook   string            `json:"webhook"`
	ExtFields map[string]string `json:"ext_fields"`
}

func (sdk *JuggleIMSdk) AddBot(bot BotInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/bots/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, bot, nil)
	return code, traceId, err
}
