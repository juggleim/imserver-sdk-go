package juggleimsdk

import "net/http"

type FriendIds struct {
	UserId    string   `json:"user_id"`
	FriendIds []string `json:"friend_ids"`
}

func (sdk *JuggleIMSdk) AddFriends(friendIds FriendIds) (ApiCode, string, error) {
	urlPath := "/apigateway/friends/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, friendIds, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelFriends(friendIds FriendIds) (ApiCode, string, error) {
	urlPath := "/apigateway/friends/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, friendIds, nil)
	return code, traceId, err
}
