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

func (sdk *JuggleIMSdk) QryHisMsgs(userId string, targetId string, channelType ChannelType, startTime int64, count int, isPositive bool) (*HisMsgs, ApiCode, string, error) {
	if count < 0 || count > 50 {
		count = 50
	}
	order := 0
	if isPositive {
		order = 1
	}
	urlPath := fmt.Sprintf("/apigateway/hismsgs/query?channel_type=%d&from_id=%s&target_id=%s&count=%d&order=%d&start=%d", channelType, userId, targetId, count, order, startTime)
	resp := &HisMsgs{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

type RecallMsgReq struct {
	FromId      string            `json:"from_id"`
	TargetId    string            `json:"target_id"`
	ChannelType int32             `json:"channel_type"`
	MsgId       string            `json:"msg_id"`
	MsgTime     int64             `json:"msg_time"`
	Exts        map[string]string `json:"exts"`
}

func (sdk *JuggleIMSdk) RecallMsg(recall *RecallMsgReq) (ApiCode, string, error) {
	urlPath := "/apigateway/hismsgs/recall"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, recall, nil)
	return code, traceId, err
}

type DelMsgsReq struct {
	FromId      string       `json:"from_id"`
	TargetId    string       `json:"target_id"`
	ChannelType int32        `json:"channel_type"`
	DelScope    int          `json:"del_scope"`
	Msgs        []*SimpleMsg `json:"msgs"`
}

type SimpleMsg struct {
	MsgId        string `json:"msg_id"`
	MsgTime      int64  `json:"msg_time"`
	MsgReadIndex int64  `json:"msg_read_index"`
}

func (sdk *JuggleIMSdk) DelMsgs(delMsgs *DelMsgsReq) (ApiCode, string, error) {
	urlPath := "/apigateway/hismsgs/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, delMsgs, nil)
	return code, traceId, err
}
