package juggleimsdk

import (
	"fmt"
	"net/http"
)

type PublicChannelInfo struct {
	ChannelId       string `json:"channel_id"`
	ChannelName     string `json:"channel_name"`
	ChannelPortrait string `json:"channel_portrait"`
	CreatorId       string `json:"creator_id"`
	CreatedTime     int64  `json:"created_time"`
	UpdatedTime     int64  `json:"updated_time"`
}

type PublicChannelMemberIds struct {
	ChannelId string   `json:"channel_id"`
	MemberIds []string `json:"member_ids"`
}

type PublicChannelMember struct {
	MemberId    string `json:"member_id"`
	CreatedTime int64  `json:"created_time"`
}

type PublicChannelMembers struct {
	Members []*PublicChannelMember `json:"members"`
	Offset  string                 `json:"offset"`
	Limit   int64                  `json:"limit"`
}

func (sdk *JuggleIMSdk) CreatePublicChannel(channel PublicChannelInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/publicchannel/create"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, channel, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) UpdatePublicChannel(channel PublicChannelInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/publicchannel/update"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, channel, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DestroyPublicChannel(channel PublicChannelInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/publicchannel/destroy"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, channel, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SubscribePublicChannel(req PublicChannelMemberIds) (ApiCode, string, error) {
	urlPath := "/apigateway/publicchannel/subscribe"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) UnsubscribePublicChannel(req PublicChannelMemberIds) (ApiCode, string, error) {
	urlPath := "/apigateway/publicchannel/unsubscribe"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryPublicChannelMembers(channelId, offset string, limit int) (*PublicChannelMembers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/publicchannel/members/list?channel_id=%s&offset=%s&limit=%d", channelId, offset, limit)
	resp := &PublicChannelMembers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
