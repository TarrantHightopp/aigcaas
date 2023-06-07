package aigcaas

import (
	"fmt"
	"net/http"
)

// Realisian
// Realisian 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/115.html

const RealisianApplicationName = `realisian_v40`

func (c *Client) realisian(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, RealisianApplicationName, c.ApiName)
	var response *http.Response
	if response, err = c.Send(url, req); err != nil {
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

func (c *Client) RealisianText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.realisian(commonText2ImgRequest)
}

func (c *Client) RealisianImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.realisian(commonImg2ImgRequest)
}
