package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 连续语义增强机器翻译（CSANMT）
// 通过输入中文，可以获得英文翻译（CSANMT连续语义增强机器翻译-中英-通用领域）
const (
	CsanmtApplicationName string = "csanmt"
	CsanmtApiName         string = "default"
)

func (c *Client) csanmt(req *CsanmtRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, CsanmtApplicationName, CsanmtApiName)
	return c.Send(url, string(byteInfo))
}

func (c *Client) CsanmtSync(req *CsanmtRequest) (*CsanmtResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.csanmt(req); err != nil {
		return nil, err
	}
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var csanmtResponse CsanmtResponse
	if err = json.Unmarshal(bodyByte, &csanmtResponse); err != nil {
		return nil, err
	}
	return &csanmtResponse, err
}
