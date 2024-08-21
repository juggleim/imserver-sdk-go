package juggleimsdk

import "net/http"

type Message struct {
	SenderId       string   `json:"sender_id"`
	TargetIds      []string `json:"target_ids"`
	MsgType        string   `json:"msg_type"`
	MsgContent     string   `json:"msg_content"`
	IsStorage      *bool    `json:"is_storage"`
	IsCount        *bool    `json:"is_count"`
	IsNotifySender *bool    `json:"is_notify_sender"`
	IsState        *bool    `json:"is_state"`
	IsCmd          *bool    `json:"is_cmd"`
}

func (sdk *JuggleIMSdk) SendSystemMsg(msg Message) (ApiCode, string, error) {
	url := sdk.ApiUrl + "apigateway/messages/system/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, url, msg, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SendGroupMsg(msg Message) (ApiCode, string, error) {
	url := sdk.ApiUrl + "apigateway/messages/group/send"
	code, traceId, err := sdk.HttpCall(http.MethodPost, url, msg, nil)
	return code, traceId, err
}
