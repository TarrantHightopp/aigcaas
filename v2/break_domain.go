package aigcaas

import (
	"fmt"
	"net/http"
)

// BreakDomain
// BreakDomain 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/116.html

const BreakDomainApplicationName = `breakdomain_M2150`

func (c *Client) breakDomain(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, BreakDomainApplicationName, c.ApiName)
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

func (c *Client) BreakDomainText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.breakDomain(commonText2ImgRequest)
}

func (c *Client) BreakDomainImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.breakDomain(commonImg2ImgRequest)
}
