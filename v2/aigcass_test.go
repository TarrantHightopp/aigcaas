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
	if response, err = client.AsyncRequestId(`78f4d20f-2468-4a14-91ea-5c3bd909e9d0`); err != nil {
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
	err = os.WriteFile("./Meinamix.jpg", ddd, 0666)          //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		panic(err)
	}
}

func TestClient_DarkSushiMixText2Img(t *testing.T) {
	req := &CommonText2ImgRequest{Prompt: "1girl, (colorful),(finely detailed beautiful eyes and detailed face),cinematic lighting,bust shot,extremely detailed CG unity 8k wallpaper,white hair,solo,smile,intricate skirt,((flying petal)),(Flowery meadow)\nsky, cloudy_sky, building, moonlight, moon, night, (dark theme:1.3), light, fantasy", Seed: 248131369}
	var res *CommonResponse
	var err error
	res, err = client.DarkSushiMixText2Img(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("request id --> ", res.AigcaasRequestId)
	fmt.Println("res --> ", res)
}

func TestClient_GuoFeng3Text2Img(t *testing.T) {
	req := &CommonText2ImgRequest{
		Prompt:         "best quality, masterpiece, highres, 1girl,blush,(seductive smile:0.8),star-shaped pupils,china hanfu,hair ornament,necklace, jewelry,Beautiful face,upon_body, tyndall effect,photorealistic, dark studio, rim lighting, two tone lighting,(high detailed skin:1.2), 8k uhd, dslr, soft lighting, high quality, volumetric lighting, candid, Photograph, high resolution, 4k, 8k, Bokeh",
		Seed:           3556647833,
		Steps:          30,
		SamplerName:    "Euler a",
		CfgScale:       9,
		NegativePrompt: `(((simple background))),monochrome ,lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, lowres, bad anatomy, bad hands, text, error, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, ugly,pregnant,vore,duplicate,morbid,mut ilated,tran nsexual, hermaphrodite,long neck,mutated hands,poorly drawn hands,poorly drawn face,mutation,deformed,blurry,bad anatomy,bad proportions,malformed limbs,extra limbs,cloned face,disfigured,gross proportions, (((missing arms))),(((missing legs))), (((extra arms))),(((extra legs))),pubic hair, plump,bad legs,error legs,username,blurry,bad fee`,
	}
	var res *CommonResponse
	var err error
	res, err = client.GuoFeng3Text2Img(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("request id --> ", res.AigcaasRequestId)
	fmt.Println("res --> ", res)
}

func TestClient_MeinamixText2Img(t *testing.T) {
	req := &CommonText2ImgRequest{
		Prompt:         "RAW photo, (close up photo:1.2) of gigachad, wearing shorts, posing for the camera,in times square, (solo:1.1), 8k uhd, dslr, high quality, film grain, Fujifilm XT3 <lora:Gigachadv1:0.7>",
		NegativePrompt: "(deformed iris, deformed pupils, semi-realistic, cgi, 3d, render, sketch, cartoon, drawing, anime:1.4), text, close up, cropped, out of frame, worst quality, low quality, jpeg artifacts, ugly, duplicate, morbid, mutilated, extra fingers, mutated hands, poorly drawn hands, poorly drawn face, mutation, deformed, blurry, dehydrated, bad anatomy, bad proportions, extra limbs, cloned face, disfigured, gross proportions, malformed limbs, missing arms, missing legs, extra arms, extra legs, fused fingers, too many fingers, long neck",
		SamplerIndex:   "DPM++ 2M Karras",
	}
	var res *CommonResponse
	var err error
	res, err = client.ProtogenText2Img(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("request id --> ", res.AigcaasRequestId)
	fmt.Println("res --> ", res)
}
