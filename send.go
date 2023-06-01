package aigcaas

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Send 发送请求
func (c *Client) Send(url, body string) (*http.Response, error) {
	var err error
	var request *http.Request
	if request, err = http.NewRequest("POST", url, strings.NewReader(body)); err != nil {
		return nil, err
	}
	c.InitRequest(request)
	switch c.Mode {
	case Polling: // 论序模式
		request.Header.Add("Aigcaas-Async", "True")
	case Notice: // 通知模式
		request.Header.Add("Aigcaas-Async", "True")
		request.Header.Add("Aigcaas-Async-Callback", c.CallbackUrl)
	}
	var client = http.Client{}
	return client.Do(request)
}

// InitRequest 初始化请求头
func (c *Client) InitRequest(request *http.Request) {
	var timestamp = time.Now().Unix()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := rand.Intn(10000)
	sign := strconv.FormatInt(timestamp, 10) + c.SecretKey + strconv.Itoa(nonce)
	hash := sha256.New()
	hash.Write([]byte(sign))
	var token = hex.EncodeToString(hash.Sum(nil))
	request.Header.Add("SecretID", c.SecretId)
	request.Header.Add("Nonce", strconv.Itoa(nonce))
	request.Header.Add("Token", token)
	request.Header.Add("Timestamp", strconv.FormatInt(timestamp, 10))
	request.Header.Add("Content-Type", "application/json")
}

// CheckToken 异步通知时候验证token
func (c *Client) CheckToken(token string, timestamp string, nonce string) bool {
	var sign = timestamp + c.SecretKey + nonce
	var hash = sha256.New()
	hash.Write([]byte(sign))
	var localToken = hex.EncodeToString(hash.Sum(nil))
	return localToken == token
}
