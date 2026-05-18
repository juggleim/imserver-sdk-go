package juggleimsdk

import (
	"fmt"
	"net/http"
)

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
	Settings      *GroupSettings    `json:"settings"`
}

type GroupSettings struct {
	HideGrpMsg          *int64 `json:"hide_grp_msg"`
	GrpMsgSecondLimiter *int64 `json:"grp_msg_second_limiter"`
	GrpMsgMinuteLimiter *int64 `json:"grp_msg_minute_limiter"`
	GrpMsgHourLimiter   *int64 `json:"grp_msg_hour_limiter"`
}

type GroupMemberSettings struct {
	HideGrpMsg *int64 `json:"hide_grp_msg"`
}

type SetGroupSettingReq struct {
	GroupId  string         `json:"group_id"`
	Settings *GroupSettings `json:"settings"`
}

type SetGroupMemberSettingReq struct {
	GroupId  string               `json:"group_id"`
	MemberId string               `json:"member_id"`
	Settings *GroupMemberSettings `json:"settings"`
}

type GroupMemberUpdateReq struct {
	GroupId        string            `json:"group_id"`
	MemberId       string            `json:"member_id"`
	GrpDisplayName string            `json:"grp_display_name"`
	ExtFields      map[string]string `json:"ext_fields"`
}

type GroupMemberAllowReq struct {
	GroupId   string   `json:"group_id"`
	MemberIds []string `json:"member_ids"`
	IsAllow   int      `json:"is_allow"`
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

func (sdk *JuggleIMSdk) QryGroupInfo(grpId string) (*GroupInfo, ApiCode, string, error) {
	urlPath := "/apigateway/groups/info?group_id=" + grpId
	resp := &GroupInfo{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

type GroupMembers struct {
	Items  []*GroupMember `json:"items"`
	Offset string         `json:"offset"`
}

type GroupMember struct {
	MemberId       string            `json:"member_id"`
	IsMute         int               `json:"is_mute"`
	IsAllow        int               `json:"is_allow"`
	GrpDisplayName string            `json:"grp_display_name"`
	ExtFields      map[string]string `json:"ext_fields"`
}

func (sdk *JuggleIMSdk) QryGroupMembers(grpId string, limit int, offset string) (*GroupMembers, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/groups/members/query?group_id=%s&limit=%d&offset=%s", grpId, limit, offset)
	resp := &GroupMembers{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) GroupMembersByIds(groupMembers GroupMembersReq) (*GroupMembers, ApiCode, string, error) {
	urlPath := "/apigateway/groups/members/querybyids"
	resp := &GroupMembers{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, groupMembers, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) GroupMemberUpdate(req GroupMemberUpdateReq) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/members/update"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SetGroupMembersAllow(grpId string, isAllow int, memberIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/groupmemberallow/set"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &GroupMemberAllowReq{
		GroupId:   grpId,
		IsAllow:   isAllow,
		MemberIds: memberIds,
	}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SetGroupMemberSettings(req SetGroupMemberSettingReq) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/members/settings/set"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

type GroupMuteReq struct {
	GroupId string `json:"group_id"`
	IsMute  int    `json:"is_mute"`
}

func (sdk *JuggleIMSdk) SetGroupMute(grpId string, isMute int) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/groupmute/set"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &GroupMuteReq{
		GroupId: grpId,
		IsMute:  isMute,
	}, nil)
	return code, traceId, err
}

type GroupMembersMuteReq struct {
	GroupId    string   `json:"group_id"`
	MemberIds  []string `json:"member_ids"`
	IsMute     int      `json:"is_mute"`
	MuteMinute int      `json:"mute_minute"`
}

func (sdk *JuggleIMSdk) SetGroupMembersMute(grpId string, isMute int, memberIds []string) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/groupmembermute/set"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, &GroupMembersMuteReq{
		GroupId:   grpId,
		IsMute:    isMute,
		MemberIds: memberIds,
	}, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) SetGroupSettings(req SetGroupSettingReq) (ApiCode, string, error) {
	urlPath := "/apigateway/groups/settings/set"
	return sdk.HttpCall(http.MethodPost, urlPath, req, nil)
}

func (sdk *JuggleIMSdk) GetGroupSettings(grpId string) (*GroupInfo, ApiCode, string, error) {
	urlPath := "/apigateway/groups/settings/get?group_id=" + grpId
	resp := &GroupInfo{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
