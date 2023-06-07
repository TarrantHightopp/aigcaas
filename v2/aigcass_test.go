package aigcaas

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

var client = NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)

func TestClient_CetusMix(t *testing.T) {
	client.ApiName = ApiNameTEXT2IMG
	client.Mode = Polling
	req := &CommonText2ImgRequest{Prompt: "1girl,death,reaper,weapon"}
	var res *CommonResponse
	var err error
	res, err = client.CetusMixText2Img(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("request id --> ", res.AigcaasRequestId)
	fmt.Println("res --> ", res)
	// 49cd820a-34a1-4686-9498-d3bb3bacfcf6
}

func TestClient_AsyncRequestId(t *testing.T) {
	var response *http.Response
	var err error
	if response, err = client.AsyncRequestId(`49cd820a-34a1-4686-9498-d3bb3bacfcf6`); err != nil {
		panic(err)
	}
	var res AsyncResponse
	var byteInfo = make([]byte, 0)
	if byteInfo, err = io.ReadAll(response.Body); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(byteInfo, &res); err != nil {
		panic(err)
	}

	if res.Error != "" {
		t.Error("err --> ", res.Error)
	}

	ddd, _ := base64.StdEncoding.DecodeString(res.Images[0]) //成图片文件并把文件写入到buffer
	err = os.WriteFile("./output.jpg", ddd, 0666)            //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		panic(err)
	}
}
func TestImageFileToBase64(t *testing.T) {
	//获取本地文件
	file, err := os.Open("./output.jpg")
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := io.ReadAll(file)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	result := base64.StdEncoding.EncodeToString(imgByte)
	fmt.Println(result)
}
