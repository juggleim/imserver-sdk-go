package juggleimsdk

import "net/http"

type Message struct {
	SenderId       string       `json:"sender_id"`
	TargetId       string       `json:"target_id"`
	ReceiverId     string       `json:"receiver_id"`
	TargetIds      []string     `json:"target_ids"`
	ToUserIds      []string     `json:"to_user_ids"`
	MsgType        string       `json:"msg_type"`
	MsgContent     string       `json:"msg_content"`
	LifeTime          int64        `json:"life_time"`
	LifeTimeAfterRead int64        `json:"life_time_after_read"`
	IsStorage      *bool        `json:"is_storage"`
	IsCount        *bool        `json:"is_count"`
	IsNotifySender *bool        `json:"is_notify_sender"`
	IsState        *bool        `json:"is_state"`
	IsCmd          *bool        `json:"is_cmd"`
	MentionInfo    *MentionInfo `json:"mention_info,omitempty"`
	ReferMsg       *ReferMsg    `json:"refer_msg,omitempty"`
	PushData       *PushData    `json:"push_data,omitempty"`

	MsgId *string `json:"msg_id,omitempty"`
}
type MentionType string

const (
	MentionType_All        MentionType = "mention_all"
	MentionType_Someone    MentionType = "mention_someone"
	MentionType_AllSomeone MentionType = "mention_all_someone"
)

type MentionInfo struct {
	MentionType   MentionType `json:"mention_type"`
	TargetUsers   []*User     `json:"target_users"`
	TargetUserIds []string    `json:"target_user_ids"`
}

type ReferMsg struct {
	MsgId       string `json:"msg_id"`
	SenderId    string `json:"sender_id"`
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
	MsgType     string `json:"msg_type"`
	MsgTime     int64  `json:"msg_time"`
	MsgContent  string `json:"msg_content"`
}

type PushData struct {
	PushTitle string `json:"push_title"`
	PushText  string `json:"push_text"`
	PushExtra string `json:"push_extra"`
	PushLevel int    `json:"push_level"`
}

func (sdk *JuggleIMSdk) SendPrivateMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/private/send"
	return sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
}

func (sdk *JuggleIMSdk) SendSystemMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/system/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SendGroupMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/group/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SendChatroomMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/chatroom/send"
	return sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
}

func (sdk *JuggleIMSdk) SendChatroomBrdMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/chatroom/broadcast"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
	return code, traceId, err
}

type TargetConver struct {
	TargetId    string `json:"target_id"`
	ChannelType int    `json:"channel_type"`
}

type SendGrpCastMsgReq struct {
	SenderId      string          `json:"sender_id"`
	TargetId      string          `json:"target_id"`
	MsgType       string          `json:"msg_type"`
	MsgContent    string          `json:"msg_content"`
	TargetConvers []*TargetConver `json:"target_convers"`
}

func (sdk *JuggleIMSdk) SendGroupCastMsg(req SendGrpCastMsgReq) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/groupcast/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

type SendBrdCastMsgReq struct {
	SenderId   string `json:"sender_id"`
	MsgType    string `json:"msg_type"`
	MsgContent string `json:"msg_content"`
	IsStorage  *bool  `json:"is_storage"`
}

func (sdk *JuggleIMSdk) SendBroadCastMsg(req SendBrdCastMsgReq) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/broadcast/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SendPublicChannelMsg(msg Message) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/publicchannel/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, nil)
	return code, traceId, err
}

type MarkReadReq struct {
	UserId      string   `json:"user_id"`
	TargetId    string   `json:"target_id"`
	ChannelType int32    `json:"channel_type"`
	MsgIds      []string `json:"msg_ids"`
}

func (sdk *JuggleIMSdk) MarkRead(req MarkReadReq) (ApiCode, string, error) {
	urlPath := "/apigateway/messages/markread"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

type StreamMsg struct {
	MsgId          string `json:"msg_id"`
	FromId         string `json:"from_id"`
	TargetId       string `json:"target_id"`
	PartialContent string `json:"partial_content"`
	Seq            int    `json:"seq"`
	IsFinished     bool   `json:"is_finished"`
}

type SendMsgRespItem struct {
	TargetId string `json:"target_id,omitempty"`
	MsgId    string `json:"msg_id"`
}

func (sdk *JuggleIMSdk) SendPrivateStreamMsg(msg StreamMsg) (*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/private/stream/send"
	resp := &SendMsgRespItem{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SendPrivateMsgWithResp(msg Message) ([]*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/private/send"
	var resp []*SendMsgRespItem
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, &resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SendSystemMsgWithResp(msg Message) ([]*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/system/send"
	var resp []*SendMsgRespItem
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, &resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SendGroupMsgWithResp(msg Message) ([]*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/group/send"
	var resp []*SendMsgRespItem
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, &resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SendChatroomMsgWithResp(msg Message) ([]*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/chatroom/send"
	var resp []*SendMsgRespItem
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, &resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SendPublicChannelMsgWithResp(msg Message) ([]*SendMsgRespItem, ApiCode, string, error) {
	urlPath := "/apigateway/messages/publicchannel/send"
	var resp []*SendMsgRespItem
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, msg, &resp)
	return resp, code, traceId, err
}
