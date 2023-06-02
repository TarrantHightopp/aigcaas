package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// VisualGLM-6B 是一个开源的，支持图像、中文和英文的多模态对话语言模型。
// 只支持 同步调用

var (
	Visualglm6bApplicationName string = "visualglm6b"
	Visualglm6bApiName         string = "sixb"
)

func (c *Client) visualglm6b(req *Visualglm6bRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, Visualglm6bApplicationName, Visualglm6bApiName)
	return c.Send(url, string(byteInfo))
}

func (c *Client) Visualglm6bSync(req *Visualglm6bRequest) (*Visualglm6bResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.visualglm6b(req); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var visualglm6bResponse Visualglm6bResponse
	if err = json.Unmarshal(bodyByte, &visualglm6bResponse); err != nil {
		return nil, err
	}
	return &visualglm6bResponse, err
}
