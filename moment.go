package juggleimsdk

import (
	"fmt"
	"net/http"
)

type MomentContent struct {
	Text   string         `json:"text"`
	Medias []*MomentMedia `json:"medias"`
}

type MomentMedia struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	SnapshotUrl string `json:"snapshot_url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Duration    int    `json:"duration"`
}

type Moment struct {
	UserId       string            `json:"user_id,omitempty"`
	MomentId     string            `json:"moment_id"`
	Title        string            `json:"title"`
	Content      *MomentContent    `json:"content"`
	ContentBrief string            `json:"content_brief"`
	MomentTime   int64             `json:"moment_time"`
	UpdatedTime  int64             `json:"updated_time"`
	UserInfo     *User             `json:"user_info"`
	TopComments  []*Comment        `json:"top_comments"`
	Reactions    []*MomentReaction `json:"reactions"`
}

type AddMomentResp struct {
	MomentId   string `json:"moment_id"`
	MomentTime int64  `json:"moment_time"`
}

type MomentIds struct {
	UserId    string   `json:"user_id"`
	MomentIds []string `json:"moment_ids"`
}

type Moments struct {
	Items      []*Moment `json:"items"`
	IsFinished bool      `json:"is_finished"`
}

type CommentContent struct {
	Text string `json:"text"`
}

type Comment struct {
	UserId          string          `json:"user_id,omitempty"`
	CommentId       string          `json:"comment_id"`
	MomentId        string          `json:"moment_id"`
	ParentCommentId string          `json:"parent_comment_id"`
	Content         *CommentContent `json:"content"`
	CommentTime     int64           `json:"comment_time"`
	UpdatedTime     int64           `json:"updated_time"`
	SeqNo           int64           `json:"seq_no"`
	ParentUserInfo  *User           `json:"parent_user_info"`
	UserInfo        *User           `json:"user_info"`
}

type CommentResp struct {
	CommentId      string `json:"comment_id"`
	CommentTime    int64  `json:"comment_time"`
	SeqNo          int64  `json:"seq_no"`
	UserInfo       *User  `json:"user_info"`
	ParentUserInfo *User  `json:"parent_user_info"`
}

type CommentIds struct {
	UserId     string   `json:"user_id"`
	MomentId   string   `json:"moment_id"`
	CommentIds []string `json:"comment_ids"`
}

type Comments struct {
	Items      []*Comment `json:"items"`
	IsFinished bool       `json:"is_finished"`
}

type MomentReaction struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
	UserInfo  *User  `json:"user_info,omitempty"`
}

type ReactionReq struct {
	UserId   string          `json:"user_id"`
	MomentId string          `json:"moment_id"`
	Reaction *MomentReaction `json:"reaction"`
}

type MomentReactions struct {
	Items      []*MomentReaction `json:"items"`
	IsFinished bool              `json:"is_finished"`
}

func (sdk *JuggleIMSdk) AddMoment(moment Moment) (*AddMomentResp, ApiCode, string, error) {
	urlPath := "/apigateway/moments/add"
	resp := &AddMomentResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, moment, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) UpdateMoment(moment Moment) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/update"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, moment, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelMoment(req MomentIds) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) MomentInfo(momentId string) (*Moment, ApiCode, string, error) {
	urlPath := "/apigateway/moments/info?moment_id=" + momentId
	resp := &Moment{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) QryMoments(userId string, startTime int64, limit int, order int) (*Moments, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/moments/list?user_id=%s&start=%d&limit=%d&order=%d", userId, startTime, limit, order)
	resp := &Moments{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddReaction(req ReactionReq) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/reactions/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelReaction(req ReactionReq) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/reactions/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryReactions(momentId string, startTime int64, limit int, order int) (*MomentReactions, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/moments/reactions/list?moment_id=%s&start=%d&limit=%d&order=%d", momentId, startTime, limit, order)
	resp := &MomentReactions{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddComment(comment Comment) (*CommentResp, ApiCode, string, error) {
	urlPath := "/apigateway/moments/comments/add"
	resp := &CommentResp{}
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, comment, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) UpdateComment(comment Comment) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/comments/update"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, comment, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DelComment(req CommentIds) (ApiCode, string, error) {
	urlPath := "/apigateway/moments/comments/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) QryComments(momentId string, startTime int64, limit int, order int) (*Comments, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/moments/comments/list?moment_id=%s&start=%d&limit=%d&order=%d", momentId, startTime, limit, order)
	resp := &Comments{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}
