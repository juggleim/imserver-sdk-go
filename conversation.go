package juggleimsdk

import "net/http"

type Conversation struct {
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
}

type Conversations struct {
	UserId string          `json:"user_id"`
	Items  []*Conversation `json:"items"`
}

func (sdk *JuggleIMSdk) ClearUnread(convers *Conversations) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/clearunread"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, convers, nil)
	return code, traceId, err
}
