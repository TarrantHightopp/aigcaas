package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 图像去色带（RealESRGAN）
// 适用于因色彩精度不够导致色带的图像，能将色带去除。

const (
	RrdbApplicationName string = "rrdb"
	RrdbApiName         string = "default"
)

func (c *Client) rrdb(req *RrdbRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, RrdbApplicationName, RrdbApiName)
	return c.Send(url, string(byteInfo))
}

// RrdbSync 同步调用
func (c *Client) RrdbSync(req *RrdbRequest) (*CommonResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.rrdb(req); err != nil {
		return nil, err
	}
	return c.Sync(response)
}
