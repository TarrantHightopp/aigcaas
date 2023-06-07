package aigcaas

import (
	"fmt"
	"net/http"
)

//【Checkpoint】YesMix
//【Checkpoint】YesMix 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/89.html

const CheckpointYesmixApplicationName = `CheckpointYesmix_v16Original`

func (c *Client) checkpointYesmix(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, CheckpointYesmixApplicationName, c.ApiName)
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

func (c *Client) CheckpointYesmixText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.checkpointYesmix(commonText2ImgRequest)
}

func (c *Client) CheckpointYesmixImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.checkpointYesmix(commonImg2ImgRequest)
}
