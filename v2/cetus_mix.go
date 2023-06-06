package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Cetus-Mix
// Cetus-Mix 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/119.html

const CetusMixApplicationName = `cetusMix_v4`

func (c *Client) CetusMix(commonRequest *CommonRequest) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, CetusMixApplicationName, c.ApiName)
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(commonRequest); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = c.Send(url, string(byteInfo)); err != nil {
		return nil, err
	}
	if err = c.ReadResponseBody(&commonResponse, response); err != nil {
		return nil, err
	}
	return commonResponse, err
}
