package aigcaas

import (
	"fmt"
	"net/http"
)

// 国风3 GuoFeng3
// 国风3 GuoFeng3 基于模型 Other 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/118.html

const GuoFengApplicationName = `3Guofeng3_v33`

func (c *Client) guoFeng3(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, GuoFengApplicationName, c.ApiName)
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

func (c *Client) GuoFeng3Text2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.guoFeng3(commonText2ImgRequest)
}

func (c *Client) GuoFeng3Img2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.guoFeng3(commonImg2ImgRequest)
}
