package aigcaas

import (
	"fmt"
	"net/http"
)

// BreakDomainRealistic 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// aigcaas.cn/product/detail/83.html

const BreakDomainRealisticApplicationName = `oxigien2ProSD21Hires_v3Epic`

func (c *Client) breakDomainRealistic(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, BreakDomainRealisticApplicationName, c.ApiName)
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

func (c *Client) BreakDomainRealisticText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.breakDomainRealistic(commonText2ImgRequest)
}

func (c *Client) BreakDomainRealisticImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.breakDomainRealistic(commonImg2ImgRequest)
}
