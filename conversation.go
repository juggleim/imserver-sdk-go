package juggleimsdk

import (
	"fmt"
	"net/http"
)

type Conversation struct {
	UserId      string `json:"user_id"`
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
	Time        int64  `json:"time"`
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

func (sdk *JuggleIMSdk) QryGlobalConvers(startTime int64, count int, targetId *string, channelType *int32) (*Conversations, ApiCode, string, error) {
	if count < 0 || count > 50 {
		count = 50
	}
	urlPath := fmt.Sprintf("/apigateway/globalconvers/query?start=%d&count=%d", startTime, count)
	if targetId != nil && *targetId != "" {
		urlPath = urlPath + fmt.Sprintf("&target_id=%s", *targetId)
	}
	if channelType != nil && *channelType > 0 {
		urlPath = urlPath + fmt.Sprintf("&channel_type=%d", *channelType)
	}
	resp := &Conversations{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
