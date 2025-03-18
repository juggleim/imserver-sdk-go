package juggleimsdk

import (
	"fmt"
	"net/http"
)

type ChannelType int

const (
	ChannelType_Private ChannelType = 1
	ChannelType_Group   ChannelType = 2
)

type HisMsgs struct {
	Msgs       []*HisMsg `json:"msgs"`
	IsFinished bool      `json:"is_finished"`
}
type HisMsg struct {
	SenderId    string `json:"sender_id"`
	TargetId    string `json:"target_id"`
	ReceiverId  string `json:"receiver_id"`
	ChannelType int32  `json:"channel_type"`
	MsgId       string `json:"msg_id"`
	MsgTime     int64  `json:"msg_time"`
	MsgType     string `json:"msg_type"`
	MsgContent  string `json:"msg_content"`

	IsStorage *bool `json:"is_storage,omitempty"`
	IsCount   *bool `json:"is_count,omitempty"`
}

func (sdk *JuggleIMSdk) QryHisMsgs(userId string, targetId string, channelType ChannelType) (*HisMsgs, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/hismsgs/query?channel_type=%d&from_id=%s&target_id=%s", channelType, userId, targetId)
	resp := &HisMsgs{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
