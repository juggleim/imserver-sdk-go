package juggleimsdk

import (
	"fmt"
	"net/http"
	"strings"
)

type UserSettingKey string

const (
	UserSettingKey_Language  UserSettingKey = "language"
	UserSettingKey_Undisturb UserSettingKey = "undisturb"
)

type User struct {
	UserId       string                    `json:"user_id"`
	Nickname     string                    `json:"nickname"`
	UserPortrait string                    `json:"user_portrait"`
	ExtFields    map[string]string         `json:"ext_fields"`
	Settings     map[UserSettingKey]string `json:"settings"`
}

type UserRegResp struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func (sdk *JuggleIMSdk) Register(user User) (*UserRegResp, ApiCode, string, error) {
	urlPath := "/apigateway/users/register"
	resp := &UserRegResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, user, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) UpdateUser(user User) (ApiCode, string, error) {
	urlPath := "/apigateway/users/update"
	code, tranceId, err := sdk.HttpCall(http.MethodPost, urlPath, user, nil)
	return code, tranceId, err
}

func (sdk *JuggleIMSdk) QryUserInfo(userId string) (*User, ApiCode, string, error) {
	urlPath := "/apigateway/users/info?user_id=" + userId
	resp := &User{}
	code, tranceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, tranceId, err
}

func (sdk *JuggleIMSdk) SetUserSettings(user User) (ApiCode, string, error) {
	urlPath := "/apigateway/users/settings/set"
	code, tranceId, err := sdk.HttpCall(http.MethodPost, urlPath, user, nil)
	return code, tranceId, err
}

func (sdk *JuggleIMSdk) GetUserSettings(userId string) (*User, ApiCode, string, error) {
	urlPath := "/apigateway/users/settings/get?user_id=" + userId
	resp := &User{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

type UserOnlineStatusReq struct {
	UserIds []string `json:"user_ids"`
}

type UserOnlineStatusResp struct {
	Items []*UserOnlineStatusItem `json:"items"`
}

type UserOnlineStatusItem struct {
	UserId   string `json:"user_id"`
	IsOnline bool   `json:"is_online"`
}

func (sdk *JuggleIMSdk) QryUserOnlineStatus(userIds []string) (*UserOnlineStatusResp, ApiCode, string, error) {
	urlPath := "/apigateway/users/onlinestatus/query"
	resp := &UserOnlineStatusResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &UserOnlineStatusReq{
		UserIds: userIds,
	}, resp)
	return resp, code, traceId, err
}

type BanUser struct {
	UserId        string `json:"user_id"`
	CreatedTime   int64  `json:"created_time"`
	EndTime       int64  `json:"end_time"`
	EndTimeOffset int64  `json:"end_time_offset"`
	ScopeKey      string `json:"scope_key"`
	ScopeValue    string `json:"scope_value"`
	Ext           string `json:"ext,omitempty"`
}

type BanUsers struct {
	Items  []*BanUser `json:"items"`
	Offset string     `json:"offset"`
}

func (sdk *JuggleIMSdk) BanUsers(banUsers *BanUsers) (ApiCode, string, error) {
	urlPath := "/apigateway/users/banusers/ban"
	code, tranceId, err := sdk.HttpCall(http.MethodPost, urlPath, banUsers, nil)
	return code, tranceId, err
}

func (sdk *JuggleIMSdk) UnBanUsers(banUsers *BanUsers) (ApiCode, string, error) {
	urlPath := "/apigateway/users/banusers/unban"
	code, tranceId, err := sdk.HttpCall(http.MethodPost, urlPath, banUsers, nil)
	return code, tranceId, err
}

func (sdk *JuggleIMSdk) QryBanUsers(limit int, offset string) (*BanUsers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/users/banusers/query?limit=%d&offset=%s", limit, offset)
	resp := &BanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err

}

func (sdk *JuggleIMSdk) QryBanUsersByUserIds(userIds []string) (*BanUsers, ApiCode, string, error) {
	if len(userIds) <= 0 {
		return nil, ApiCode_Success, "", nil
	}
	urlPath := "/apigateway/users/banusers/query"
	idParams := []string{}
	for _, userId := range userIds {
		idParams = append(idParams, fmt.Sprintf("user_id=%s", userId))
	}
	urlPath = fmt.Sprintf("%s?%s", urlPath, strings.Join(idParams, "&"))
	resp := &BanUsers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
