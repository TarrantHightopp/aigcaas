package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 胶片质感扩散生成模型
// 生成胶片质感的图片，利用 dreambooth 方法微调 Stable Diffusion 1.5 而来，数据集为胶片摄影数据集

const (
	AnalogApplicationName string = "analog"
	AnalogApiName         string = "analog"
)

func (c *Client) analog(req *CommonStableDiffusionRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, AnalogApplicationName, AnalogApiName)
	return c.Send(url, string(byteInfo))
}

// AnalogSync 同步
func (c *Client) AnalogSync(req *CommonStableDiffusionRequest) (commonStableDiffusionResponse *CommonStableDiffusionResponse, err error) {
	var response *http.Response
	if response, err = c.analog(req); err != nil {
		return nil, err
	}
	return c.StableDiffusionSync(response)
}

// AnalogAsync 异步
func (c *Client) AnalogAsync(req *CommonStableDiffusionRequest) (string, error) {
	var err error
	var response *http.Response
	if response, err = c.analog(req); err != nil {
		return "", err
	}
	return c.StableDiffusionAsync(response)
}
