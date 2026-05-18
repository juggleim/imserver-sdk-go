package juggleimsdk

import "net/http"

type BindDevice struct {
	UserId        string `json:"user_id"`
	DeviceId      string `json:"device_id"`
	Platform      string `json:"platform"`
	DeviceCompany string `json:"device_company"`
	DeviceModel   string `json:"device_model"`
	CreatedTime   int64  `json:"created_time"`
}

type BindDevicesResp struct {
	Items []*BindDevice `json:"items"`
}

func (sdk *JuggleIMSdk) AddBindDevice(device BindDevice) (ApiCode, string, error) {
	urlPath := "/apigateway/binddevices/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, device, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelBindDevice(device BindDevice) (ApiCode, string, error) {
	urlPath := "/apigateway/binddevices/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, device, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryBindDevices(userId string) (*BindDevicesResp, ApiCode, string, error) {
	urlPath := "/apigateway/binddevices/query?user_id=" + userId
	resp := &BindDevicesResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
