package juggleimsdk

import (
	"fmt"
	"net/http"
)

type UserIds struct {
	UserIds []string `json:"user_ids"`
}

type MuteUser struct {
	UserId      string `json:"user_id"`
	CreatedTime int64  `json:"created_time"`
}

type QryMuteUsersResp struct {
	Items  []*MuteUser `json:"items"`
	Offset string      `json:"offset"`
}

func (sdk *JuggleIMSdk) AddPrivateGlobalMuteMembers(userIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/private/globalmutemembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserIds{UserIds: userIds}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelPrivateGlobalMuteMembers(userIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/private/globalmutemembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserIds{UserIds: userIds}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryPrivateGlobalMuteMembers(limit int, offset string) (*QryMuteUsersResp, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/private/globalmutemembers/query?limit=%d&offset=%s", limit, offset)
	resp := &QryMuteUsersResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddGroupGlobalMuteMembers(userIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/group/globalmutemembers/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserIds{UserIds: userIds}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelGroupGlobalMuteMembers(userIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/group/globalmutemembers/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserIds{UserIds: userIds}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryGroupGlobalMuteMembers(limit int, offset string) (*QryMuteUsersResp, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/group/globalmutemembers/query?limit=%d&offset=%s", limit, offset)
	resp := &QryMuteUsersResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
