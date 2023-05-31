package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	AnalogApplicationName string = "analog"
	AnalogApiName         string = "analog"
)

func (c *Client) Analog(req *CommonStableDiffusionRequest) (*CommonStableDiffusionResponse, error) {
	var err error
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, AnalogApplicationName, AnalogApiName)
	var response *http.Response
	if response, err = c.Send(url, string(byteInfo)); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var commonStableDiffusionResponse CommonStableDiffusionResponse
	if err = json.Unmarshal(bodyByte, &commonStableDiffusionResponse); err != nil {
		return nil, err
	}

	return &commonStableDiffusionResponse, nil
}
