package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	Vitgpt2ApplicationName string = "vitgpt2"
	Vitgpt2ApiName         string = "desc"
)

func (c *Client) vitgpt2(req *Vitgpt2Request) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, Vitgpt2ApplicationName, Vitgpt2ApiName)
	return c.Send(url, string(byteInfo))
}

func (c *Client) Vitgpt2Sync(req *Vitgpt2Request) (*Vitgpt2Response, error) {
	var err error
	var response *http.Response
	if response, err = c.vitgpt2(req); err != nil {
		return nil, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var vitgpt2Response Vitgpt2Response
	if err = json.Unmarshal(bodyByte, &vitgpt2Response); err != nil {
		return nil, err
	}
	return &vitgpt2Response, err
}
