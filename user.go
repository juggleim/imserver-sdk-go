package juggleimsdk

import (
	"net/http"
)

type User struct {
	UserId       string            `json:"user_id"`
	Nickname     string            `json:"nickname"`
	UserPortrait string            `json:"user_portrait"`
	ExtFields    map[string]string `json:"ext_fields"`
	Settings     map[string]string `json:"settings"`
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
	code, tranceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, tranceId, err
}
