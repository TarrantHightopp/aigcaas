package aigcaas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

var client = NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)

func TestClient_CetusMix(t *testing.T) {
	client.ApiName = ApiNameTEXT2IMG
	req := &CommonRequest{Prompt: "1girl,death,reaper,weapon"}
	var res *CommonResponse
	var err error
	res, err = client.CetusMix(req)
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
