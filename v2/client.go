package aigcaas

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	SecretKey   string `json:"secret_key"`
	SecretId    string `json:"secret_id"`
	Mode        int    `json:"mode"`
	CallbackUrl string `json:"callback_url"`
	ApiName     string `json:"api_name"`
}

func NewClient(secretKey, secretId string) *Client {
	return &Client{
		SecretKey: secretKey,
		SecretId:  secretId,
	}
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

func (c *Client) AsyncRequestId(requestId string) (*http.Response, error) {
	var err error
	var request *http.Request
	if request, err = http.NewRequest("GET", AsyncRequestIdURL, nil); err != nil {
		return nil, err
	}
	c.InitRequest(request)
	request.Header.Add("RequestId", requestId)
	var client = http.Client{}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, err
	}
	return response, err
}

func (c *Client) ReadResponseBody(data interface{}, response *http.Response) (err error) {
	var body = make([]byte, 0)
	if body, err = io.ReadAll(response.Body); err != nil {
		fmt.Println("read err --> ", err)
		return err
	}
	fmt.Println(string(body))
	return json.Unmarshal(body, &data)
}
