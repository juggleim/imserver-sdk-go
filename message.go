package juggleimsdk

import "net/http"

type Message struct {
	SenderId       string       `json:"sender_id"`
	TargetIds      []string     `json:"target_ids"`
	MsgType        string       `json:"msg_type"`
	MsgContent     string       `json:"msg_content"`
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
