package juggleimsdk

import (
	"fmt"
	"net/http"
	"strings"
)

type UserTag struct {
	UserID string   `json:"user_id"`
	Tags   []string `json:"tags"`
}

type UserTagsPayload struct {
	UserTags []UserTag `json:"user_tags"`
}

type PushCondition struct {
	TagsAnd []string `json:"tags_and"`
	TagsOr  []string `json:"tags_or"`
}

type PushMsgBody struct {
	MsgType    string `json:"msg_type"`
	MsgContent string `json:"msg_content"`
}

type PushNotification struct {
	Title    string `json:"title"`
	PushText string `json:"push_text"`
}

type PushPayload struct {
	FromUserId   string            `json:"from_user_id"`
	Condition    PushCondition     `json:"condition"`
	MsgBody      *PushMsgBody      `json:"msg_body,omitempty"`
	Notification *PushNotification `json:"notification,omitempty"`
}

type PushResp struct {
	PushId string `json:"push_id"`
}

func (sdk *JuggleIMSdk) QryUserTags(userIds []string) (*UserTagsPayload, ApiCode, string, error) {
	if len(userIds) == 0 {
		return &UserTagsPayload{}, ApiCode_Success, "", nil
	}
	params := make([]string, len(userIds))
	for i, uid := range userIds {
		params[i] = fmt.Sprintf("user_id=%s", uid)
	}
	urlPath := "/apigateway/usertags/query?" + strings.Join(params, "&")
	resp := &UserTagsPayload{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddUserTags(req UserTagsPayload) (ApiCode, string, error) {
	urlPath := "/apigateway/usertags/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelUserTags(req UserTagsPayload) (ApiCode, string, error) {
	urlPath := "/apigateway/usertags/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) ClearUserTags(userIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/usertags/clear"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserIds{UserIds: userIds}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) PushWithTags(req PushPayload) (*PushResp, ApiCode, string, error) {
	urlPath := "/apigateway/push"
	resp := &PushResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, resp)
	return resp, code, traceId, err
}
