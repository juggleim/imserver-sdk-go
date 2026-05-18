package juggleimsdk

import "net/http"

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
