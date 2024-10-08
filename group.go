package juggleimsdk

import "net/http"

type GroupMembersReq struct {
	GroupId       string   `json:"group_id"`
	GroupName     string   `json:"group_name"`
	GroupPortrait string   `json:"group_portrait"`
	MemberIds     []string `json:"member_ids"`
}

func (sdk *JuggleIMSdk) CreateGroup(groupMembers GroupMembersReq) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, groupMembers, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) GroupAddMembers(groupMembers GroupMembersReq) (ApiCode, string, error) {
	return sdk.CreateGroup(groupMembers)
}

func (sdk *JuggleIMSdk) GroupDelMembers(groupMembers GroupMembersReq) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/members/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, groupMembers, nil)
	return code, traceId, err
}

type GroupInfo struct {
	GroupId       string            `json:"group_id"`
	GroupName     string            `json:"group_name"`
	GroupPortrait string            `json:"group_portrait"`
	IsMute        int               `json:"is_mute"`
	UpdatedTime   int64             `json:"updated_time"`
	ExtFields     map[string]string `json:"ext_fields"`
}

func (sdk *JuggleIMSdk) DissolveGroup(groupId string) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &GroupInfo{
		GroupId: groupId,
	}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) UpdateGroup(grpInfo GroupInfo) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/update"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &grpInfo, nil)
	return code, traceId, err
}
