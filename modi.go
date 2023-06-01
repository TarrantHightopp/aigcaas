package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 迪士尼风格扩散生成模型
// 本模型自 Stable Diffusion 1.5 微调而来，微调数据来自某著名动画工作室的电影截图

const (
	ModiApplicationName string = "modi"
	ModiApiName         string = "modi"
)

func (c *Client) modi(req *CommonStableDiffusionRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, ModiApplicationName, ModiApiName)
	return c.Send(url, string(byteInfo))
}

// ModiSync 同步
func (c *Client) ModiSync(req *CommonStableDiffusionRequest) (commonStableDiffusionResponse *CommonStableDiffusionResponse, err error) {
	var response *http.Response
	if response, err = c.modi(req); err != nil {
		return nil, err
	}
	return c.StableDiffusionSync(response)
}

// ModiAsync 异步
func (c *Client) ModiAsync(req *CommonStableDiffusionRequest) (string, error) {
	var err error
	var response *http.Response
	if response, err = c.modi(req); err != nil {
		return "", err
	}
	return c.StableDiffusionAsync(response)
}
