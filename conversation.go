package juggleimsdk

import (
	"fmt"
	"net/http"
)

type Conversation struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
	SubChannel  string `json:"sub_channel"`
	Time        int64  `json:"time"`
}

type Conversations struct {
	UserId     string          `json:"user_id,omitempty"`
	Items      []*Conversation `json:"items"`
	IsFinished bool            `json:"is_finished"`
}

type UndisturbConverItem struct {
	TargetId      string `json:"target_id"`
	ChannelType   int    `json:"channel_type"`
	UndisturbType int32  `json:"undisturb_type"`
}

type UndisturbConversReq struct {
	UserId string                 `json:"user_id"`
	Items  []*UndisturbConverItem `json:"items"`
}

type TopConverReqItem struct {
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
	IsTop       bool   `json:"is_top"`
}

type TopConversReq struct {
	UserId string              `json:"user_id"`
	Items  []*TopConverReqItem `json:"items"`
}

type TagConversReq struct {
	UserId  string          `json:"user_id"`
	Tag     string          `json:"tag"`
	TagName string          `json:"tag_name"`
	Convers []*Conversation `json:"convers"`
}

func (sdk *JuggleIMSdk) ClearUnread(convers *Conversations) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/clearunread"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, convers, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryGlobalConvers(startTime int64, count int, targetId *string, channelType *int32) (*Conversations, ApiCode, string, error) {
	return sdk.QryGlobalConversWithExclude(startTime, count, targetId, channelType, nil)
}

func (sdk *JuggleIMSdk) QryGlobalConversWithExclude(startTime int64, count int, targetId *string, channelType *int32, excludeUserIds []string) (*Conversations, ApiCode, string, error) {
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
	for _, uid := range excludeUserIds {
		urlPath = urlPath + fmt.Sprintf("&exclude_user_id=%s", uid)
	}
	resp := &Conversations{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddConversation(conver Conversation) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, conver, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelConversation(convers Conversations) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, convers, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) UndisturbConvers(req UndisturbConversReq) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/undisturb"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) TopConversations(req TopConversReq) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/top"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryConvers(userId string, startTime int64, count int, order int) (*Conversations, ApiCode, string, error) {
	if count < 0 || count > 100 {
		count = 100
	}
	urlPath := fmt.Sprintf("/apigateway/convers/query?user_id=%s&start=%d&count=%d&order=%d", userId, startTime, count, order)
	resp := &Conversations{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) TagConvers(req TagConversReq) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/tags/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) UnTagConvers(req TagConversReq) (ApiCode, string, error) {
	urlPath := "/apigateway/convers/tags/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}
