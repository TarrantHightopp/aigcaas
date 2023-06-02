package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BAAI/AltDiffusion
// 我们的版本在中英文对齐方面表现非常出色，是目前市面上开源的最强版本，保留了原版stable diffusion的大部分能力，并且在某些例子上比有着比原版模型更出色的能力。
// 当前接口只支持异步调用

const (
	AltDiffusionApplicationName string = "altdiffusion"
	AltDiffusionApiName         string = "default"
)

func (c *Client) altDiffusion(req *AltDiffusionRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, AltDiffusionApplicationName, AltDiffusionApiName)
	return c.Send(url, string(byteInfo))
}

// AltDiffusionAsync 异步
func (c *Client) AltDiffusionAsync(req *AltDiffusionRequest) (string, error) {
	var err error
	var response *http.Response
	if response, err = c.altDiffusion(req); err != nil {
		return "", err
	}
	return c.StableDiffusionAsync(response)
}

func (c *Client) AltDiffusionAsyncRequestId(requestId string) (*AltDiffusionSyncResponse, error) {
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
	var altDiffusionSyncResponse AltDiffusionSyncResponse
	if err = json.Unmarshal(bodyInfo, &altDiffusionSyncResponse); err != nil {
		return nil, err
	}
	return &altDiffusionSyncResponse, nil
}
