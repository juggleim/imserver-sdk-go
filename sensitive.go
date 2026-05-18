package juggleimsdk

import (
	"fmt"
	"net/http"
)

type SensitiveWord struct {
	Id       string `json:"id"`
	Word     string `json:"word"`
	WordType int    `json:"word_type"`
}

type SensitiveWords struct {
	Items      []*SensitiveWord `json:"items"`
	Total      int32            `json:"total"`
	IsFinished bool             `json:"is_finished"`
}

type DelSensitiveWordsReq struct {
	Words []string `json:"words"`
}

func (sdk *JuggleIMSdk) QrySensitiveWords(size, page int, word string, wordType int) (*SensitiveWords, ApiCode, string, error) {
	urlPath := fmt.Sprintf("/apigateway/sensitivewords/list?size=%d&page=%d&word=%s&word_type=%d", size, page, word, wordType)
	resp := &SensitiveWords{}
	code, traceId, err := sdk.HttpCall(http.MethodGet, urlPath, nil, resp)
	return resp, code, traceId, err
}

func (sdk *JuggleIMSdk) AddSensitiveWords(words SensitiveWords) (ApiCode, string, error) {
	urlPath := "/apigateway/sensitivewords/add"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, words, nil)
	return code, traceId, err
}

func (sdk *JuggleIMSdk) DeleteSensitiveWords(req DelSensitiveWordsReq) (ApiCode, string, error) {
	urlPath := "/apigateway/sensitivewords/del"
	code, traceId, err := sdk.HttpCall(http.MethodPost, urlPath, req, nil)
	return code, traceId, err
}
