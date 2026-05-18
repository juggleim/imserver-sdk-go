package juggleimsdk

import "net/http"

type BotConf struct {
	BotId    string `json:"bot_id,omitempty"`
	Url      string `json:"url,omitempty"`
	ApiKey   string `json:"api_key,omitempty"`
	IsStream bool   `json:"is_stream"`
}

type BotSettings struct {
	OnlyMentioned bool `json:"only_mentioned"`
}

type BotInfo struct {
	BotId       string            `json:"bot_id"`
	Nickname    string            `json:"nickname"`
	Portrait    string            `json:"portrait"`
	BotConf     *BotConf          `json:"bot_conf,omitempty"`
	BotSettings *BotSettings      `json:"bot_settings,omitempty"`
	ExtFields   map[string]string `json:"ext_fields"`
	UpdatedTime int64             `json:"updated_time"`
}

func (sdk *JuggleIMSdk) RegisterBot(bot BotInfo) (*UserRegResp, ApiCode, string, error) {
	urlPath := "/apigateway/bots/register"
	resp := &UserRegResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, bot, resp)
	return resp, code, traceId, err
}
