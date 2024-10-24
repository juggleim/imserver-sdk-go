package juggleimsdk

import (
	"fmt"
	"net/http"
)

type ChatroomInfo struct {
	ChatId      string            `json:"chat_id"`
	ChatName    string            `json:"chat_name"`
	Members     []*ChatroomMember `json:"members"`
	Atts        []*ChatroomAtt    `json:"atts"`
	MemberCount int32             `json:"member_count"`
	IsMute      int               `json:"is_mute"`
}

type ChatroomMember struct {
	MemberId   string `json:"member_id"`
	MemberName string `json:"member_name"`
	AddedTime  int64  `json:"added_time"`
}

type ChatroomAtt struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	AttTime int64  `json:"att_time"`
	UserId  string `json:"user_id"`

	IsForce *bool `json:"is_force,omitempty"`
}

func (sdk *JuggleIMSdk) CreateChatroom(chat ChatroomInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/create"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, chat, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DestroyChatroom(chat ChatroomInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/destroy"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, chat, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryChatroomInfo(chatId string, withMembers, withAtts bool, order, count int32) (*ChatroomInfo, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/chatrooms/info?chat_id=%s&with_members=%v&with_atts=%v&order=%d&count=%d", chatId, withMembers, withAtts, order, count)
	resp := &ChatroomInfo{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SetChatroomMute(chatId string, isMute bool) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/chrmmute/set"
	mute := 0
	if isMute {
		mute = 1
	}
	chat := &ChatroomInfo{
		ChatId: chatId,
		IsMute: mute,
	}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, chat, nil)
	return code, traceId, err
}
