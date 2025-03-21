package juggleimsdk

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"math/rand"

	"github.com/google/uuid"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type ApiCode int32

var (
	ApiCode_Success          ApiCode = 0
	ApiCode_HttpTimeout      ApiCode = 1
	ApiCode_DecodeFail       ApiCode = 2
	ApiCode_NotSupportMethod ApiCode = 3
)

type ApiResp struct {
	Code ApiCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func (sdk *JuggleIMSdk) HttpCall(method, urlPath string, req interface{}, resp interface{}) (ApiCode, string, error) {
	url := sdk.ApiUrl + urlPath
	traceId := GenerateUUIDShort11()
	headers := sdk.getHeaders()
	var respBs []byte
	var err error
	var body string = ""
	if method == http.MethodPost {
		bodyBs, _ := json.Marshal(req)
		body = string(bodyBs)
	} else if method == http.MethodGet {
	} else {
		return ApiCode_NotSupportMethod, traceId, fmt.Errorf("not support method:%s", method)
	}

	respBs, err = HttpDoBytes(method, url, headers, body)
	if err != nil {
		return ApiCode_HttpTimeout, traceId, err
	}
	apiResp := &ApiResp{
		Data: resp,
	}
	err = json.Unmarshal(respBs, apiResp)
	if err != nil {
		return ApiCode_DecodeFail, traceId, err
	}
	if apiResp.Code != ApiCode_Success {
		return ApiCode(apiResp.Code), traceId, fmt.Errorf(apiResp.Msg)
	}
	if resp != nil && apiResp.Data == nil {
		return ApiCode_DecodeFail, traceId, fmt.Errorf("decode fail")
	}
	return ApiCode_Success, traceId, nil
}

func HttpDoBytes(method, url string, header map[string]string, body string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return []byte{}, err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}

	resp, err := client.Do(request)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err == nil && resp != nil && resp.Body != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}
		return respBody, nil
	}
	return []byte{}, err
}

func (sdk *JuggleIMSdk) getHeaders() map[string]string {
	nonce := fmt.Sprintf("%d", rand.Int31n(10000))
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	signature := SHA1(fmt.Sprintf("%s%s%s", sdk.Secret, nonce, timestamp))

	return map[string]string{
		"Content-Type": "application/json",
		"appkey":       sdk.Appkey,
		"nonce":        nonce,
		"timestamp":    timestamp,
		"signature":    signature,
	}
}
func GenerateUUID() uuid.UUID {
	uid := uuid.New()
	return uid
}

func GenerateUUIDBytes() []byte {
	uid, _ := uuid.NewUUID()
	return []byte(uid.String())
}

func UUIDStringByBytes(bytes []byte) (string, error) {
	uuid, err := uuid.FromBytes(bytes)
	return uuid.String(), err
}

func GenerateUUIDShort11() string {
	return ShortCut(GenerateUUIDShort22())
}
func ShortCut(str string) string {
	if len(str) > 16 {
		return str[5:16]
	}
	return ""
}
func GenerateUUIDShort22() string {
	return UUID2ShortString(GenerateUUID())
}

func UUID2ShortString(uuid uuid.UUID) string {
	mostBits := make([]byte, 8)
	leastBits := make([]byte, 8)
	for i := 0; i < 8; i++ {
		mostBits[i] = uuid[i]
	}
	for i := 8; i < 16; i++ {
		leastBits[i-8] = uuid[i]
	}
	return strings.Join([]string{toIdString(BytesToUInt64(mostBits)), toIdString(BytesToUInt64(leastBits))}, "")
}

var DIGITS64 []byte = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")

func toIdString(l uint64) string {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	var length int = 11
	var least uint64 = 63 //0x3f

	for {
		length--
		buf[length] = DIGITS64[int(l&least)]
		l = l >> 6
		if l == 0 {
			break
		}
	}
	return string(buf)
}
func BytesToUInt64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

func ToJson(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bs)
}
