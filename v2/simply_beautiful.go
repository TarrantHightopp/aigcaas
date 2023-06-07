package aigcaas

import (
	"fmt"
	"net/http"
)

// Simply Beautiful
// Simply Beautiful 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/94.html

const SimplyBeautifulApplicationName = `simplyBeautiful_v10`

func (c *Client) simplyBeautiful(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, SimplyBeautifulApplicationName, c.ApiName)
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

func (c *Client) SimplyBeautifulText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.simplyBeautiful(commonText2ImgRequest)
}

func (c *Client) SimplyBeautifulImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.simplyBeautiful(commonImg2ImgRequest)
}
