package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 基于 Stable Diffusion 模型的API公用方法

// StableDiffusionSync 同步
func (c *Client) StableDiffusionSync(response *http.Response) (commonStableDiffusionResponse *CommonStableDiffusionResponse, err error) {
	defer func() {
		_ = response.Body.Close()
	}()
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyByte, &commonStableDiffusionResponse); err != nil {
		return nil, err
	}
	return commonStableDiffusionResponse, nil
}

// StableDiffusionAsync 异步
func (c *Client) StableDiffusionAsync(response *http.Response) (string, error) {
	var err error
	defer func() {
		_ = response.Body.Close()
	}()
	return response.Header.Get("Aigcaas-Request-Id"), err
}

// AsyncRequestId 根据 request id 获取内容
func (c *Client) AsyncRequestId(requestId string) (*AsyncRequestIdResponse, error) {
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
	defer func() {
		_ = response.Body.Close
	}()
	var bodyInfo = make([]byte, 0)
	if bodyInfo, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	fmt.Println(string(bodyInfo))
	var asyncRequestIdResponse AsyncRequestIdResponse
	if err = json.Unmarshal(bodyInfo, &asyncRequestIdResponse); err != nil {
		return nil, err
	}
	return &asyncRequestIdResponse, nil
}
