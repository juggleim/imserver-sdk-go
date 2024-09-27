package juggleimsdk

import "net/http"

type ChatroomInfo struct {
	ChatId   string `json:"chat_id"`
	ChatName string `json:"chat_name"`
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
