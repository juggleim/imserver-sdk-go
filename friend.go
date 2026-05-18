package juggleimsdk

import (
	"fmt"
	"net/http"
)

type FriendIds struct {
	UserId    string        `json:"user_id"`
	FriendIds []string      `json:"friend_ids"`
	Friends   []*FriendItem `json:"friends"`
}

type FriendItem struct {
	UserId      string `json:"user_id,omitempty"`
	FriendId    string `json:"friend_id"`
	DisplayName string `json:"display_name"`
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

type FriendsResp struct {
	Items  []*FriendItem `json:"items"`
	Offset string        `json:"offset"`
}

func (sdk *JuggleIMSdk) QryFriends(userId string, limit int, offset string, order int) (*FriendsResp, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/friends/query?user_id=%s&limit=%d&offset=%s&order=%d", userId, limit, offset, order)
	resp := &FriendsResp{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) SetFriendDisplayName(item FriendItem) (ApiCode, string, error) {
	urlPath := "/apigateway/friends/setdisplayname"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, item, nil)
	return code, traceId, err
}
