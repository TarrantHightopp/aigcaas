package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 语音识别（Paraformer）
// 可以输入语音，输出对应的文字信息。Paraformer是达摩院语音团队提出的一种高效的非自回归端到端语音识别框架。
// 只支持 同步调用
const (
	ParaformerApplicationName string = "paraformer"
	ParaformerApiName         string = "default"
)

func (c *Client) paraformer(req *ParaformerRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, ParaformerApplicationName, ParaformerApiName)
	return c.Send(url, string(byteInfo))
}

func (c *Client) ParaformerSync(req *ParaformerRequest) (*ParaformerResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.paraformer(req); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var paraformerResponse ParaformerResponse
	if err = json.Unmarshal(bodyByte, &paraformerResponse); err != nil {
		return nil, err
	}
	return &paraformerResponse, err
}
