package juggleimsdk

import "net/http"

type AppInfo struct {
	AppType     int    `json:"app_type"`
	AppName     string `json:"app_name"`
	CreatedTime int64  `json:"created_time"`
	UpdateTime  int64  `json:"updated_time"`

	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	AppStatus int    `json:"app_status"`

	MaxUserCount int   `json:"max_user_count"`
	CurUserCount int64 `json:"cur_user_count"`

	RestrictedFields *RestrictedFields `json:"restricted_fields"`
	ConfigFields     map[string]string `json:"config_fields"`

	ExpiredTime int64 `json:"expired_time"`

	LicenseConf string `json:"license_conf,omitempty"`
}

type RestrictedFields struct {
	MaxUserCount int32 `json:"max_user_count"`
}

type ActiveAppReq struct {
	License string `json:"license"`
}

type ConnectSignKeysReq struct {
	SignKeys []string `json:"sign_keys"`
}

type ConnectSignKeysResp struct {
	SignKeys []string `json:"sign_keys"`
}

func (sdk *JuggleIMSdk) AddConnectSignKeys(signKeys []string) (ApiCode, string, error) {
	urlPath := "/apigateway/connectsignkeys/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &ConnectSignKeysReq{SignKeys: signKeys}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryConnectSignKeys() (*ConnectSignKeysResp, ApiCode, string, error) {
	urlPath := "/apigateway/connectsignkeys/query"
	resp := &ConnectSignKeysResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddAppConnectSignKeys(signKeys []string) (ApiCode, string, error) {
	urlPath := "/apigateway/apps/connectsignkeys/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &ConnectSignKeysReq{SignKeys: signKeys}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryAppConnectSignKeys() (*ConnectSignKeysResp, ApiCode, string, error) {
	urlPath := "/apigateway/apps/connectsignkeys/query"
	resp := &ConnectSignKeysResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
