package aigcaas

import (
	"fmt"
	"net/http"
)

// Cetus-Mix
// Cetus-Mix 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/119.html

const CetusMixApplicationName = `cetusMix_v4`

func (c *Client) CetusMix(commonRequest *CommonRequest) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, CetusMixApplicationName, c.ApiName)
	var response *http.Response
	if response, err = c.Send(url, commonRequest); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if err = c.ReadResponseBody(&commonResponse, response); err != nil {
		return nil, err
	}
	return commonResponse, nil
}
