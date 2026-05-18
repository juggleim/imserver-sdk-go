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
	EndTime    int64  `json:"end_time"`
}

type ChrmBanUserReq struct {
	ChatId        string   `json:"chat_id"`
	MemberIds     []string `json:"member_ids"`
	EndTime       int64    `json:"end_time"`
	EndTimeOffset int64    `json:"end_time_offset"`
}

type ChrmBanUsers struct {
	ChatId  string            `json:"chat_id"`
	Members []*ChatroomMember `json:"members"`
	Offset  string            `json:"offset"`
}

type ChrmMemberIds struct {
	ChatId    string   `json:"chat_id"`
	MemberIds []string `json:"member_ids"`
}

type ChrmMemberExistItem struct {
	MemberId string `json:"member_id"`
	Exist    bool   `json:"exist"`
}

type ChrmMembersExistResp struct {
	Items []*ChrmMemberExistItem `json:"items"`
}

type ChatroomAtts struct {
	FromId string         `json:"from_id"`
	ChatId string         `json:"chat_id"`
	Atts   []*ChatroomAtt `json:"atts"`
}

type ChatroomAttResp struct {
	Key     string `json:"key"`
	Code    int32  `json:"code"`
	AttTime int64  `json:"att_time"`
}

type ChatroomAttsResp struct {
	Atts []*ChatroomAttResp `json:"atts"`
}

type ChatroomAttsReq struct {
	ChatId  string   `json:"chat_id"`
	AttKeys []string `json:"att_keys"`
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

func (sdk *JuggleIMSdk) ChrmMembersExists(req ChrmMemberIds) (*ChrmMembersExistResp, ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/members/exist"
	resp := &ChrmMembersExistResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddChrmMuteMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/mutemembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelChrmMuteMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/mutemembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryChrmMuteMembers(chatId, offset string, limit int) (*ChrmBanUsers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/chatrooms/mutemembers/query?chat_id=%s&offset=%s&limit=%d", chatId, offset, limit)
	resp := &ChrmBanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddChrmGlobalMuteMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/globalmutemembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelChrmGlobalMuteMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/globalmutemembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryChrmGlobalMuteMembers(offset string, limit int) (*ChrmBanUsers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/chatrooms/globalmutemembers/query?offset=%s&limit=%d", offset, limit)
	resp := &ChrmBanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddChrmBanMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/banmembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelChrmBanMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/banmembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryChrmBanMembers(chatId, offset string, limit int) (*ChrmBanUsers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/chatrooms/banmembers/query?chat_id=%s&offset=%s&limit=%d", chatId, offset, limit)
	resp := &ChrmBanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddChrmAllowMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/allowmembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelChrmAllowMembers(req ChrmBanUserReq) (ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/allowmembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryChrmAllowMembers(chatId, offset string, limit int) (*ChrmBanUsers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/chatrooms/allowmembers/query?chat_id=%s&offset=%s&limit=%d", chatId, offset, limit)
	resp := &ChrmBanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddChrmAtts(req ChatroomAtts) (*ChatroomAttsResp, ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/atts/add"
	resp := &ChatroomAttsResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) DelChrmAtts(req ChatroomAtts) (*ChatroomAttsResp, ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/atts/del"
	resp := &ChatroomAttsResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) QryChrmAtts(req ChatroomAttsReq) (*ChatroomInfo, ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/atts/qry"
	resp := &ChatroomInfo{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) ListChrmAtts(chatId string) (*ChatroomInfo, ApiCode, string, error) {
	urlPath := "/apigateway/chatrooms/atts/list?chat_id=" + chatId
	resp := &ChatroomInfo{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
