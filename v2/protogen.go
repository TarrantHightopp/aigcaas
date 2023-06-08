package aigcaas

import (
	"fmt"
	"net/http"
)

const ProtogenApplicationName = `protogenV22Anime_protogenV22`

func (c *Client) protogen(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, ProtogenApplicationName, c.ApiName)
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

func (c *Client) ProtogenText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.protogen(commonText2ImgRequest)
}

func (c *Client) ProtogenImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.protogen(commonImg2ImgRequest)
}
