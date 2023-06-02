package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 图像去噪与去模糊化（NAFNet）
// 适用于智能手机拍摄的带噪图片，通过该应用，可以进行图像的去燥与去模糊操作。

var NafNetApplicationName string = "mafnet_denoise"

const (
	NafNetDeblur  = "deblur"  // 图拍去模糊
	NafNetDenoise = "denoise" // 图像去噪
)

func (c *Client) nafNet(req *NafNetRequest, apiName string) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, NafNetApplicationName, apiName)
	return c.Send(url, string(byteInfo))
}

// NafNetSync 同步调用
func (c *Client) NafNetSync(req *NafNetRequest, apiName string) (*CommonResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.nafNet(req, apiName); err != nil {
		return nil, err
	}
	return c.Sync(response)
}
