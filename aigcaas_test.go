package aigcaas

import (
	"encoding/json"
	"fmt"
	"testing"
)

var client = NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)

func TestClient_AnalogSync(t *testing.T) {
	commonStableDiffusionRequest := CommonRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonResponse
	if commonStableDiffusionResponse, err = client.AnalogSync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonStableDiffusionResponse)
}

func TestClient_AnalogAsync(t *testing.T) {
	var err error
	commonStableDiffusionRequest := CommonRequest{
		Text: "dog",
	}
	client.Mode = Polling
	var requestId string
	if requestId, err = client.AnalogAsync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}

	fmt.Println(requestId)
}

func TestClient_AsyncRequestId(t *testing.T) {
	var err error
	var resp *AsyncRequestIdResponse
	if resp, err = client.AsyncRequestId(`47CE03F2-0A1E-59DF-B7E6-963903FE0BB9`); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println("message --> ", resp.Message)
	fmt.Println("status --> ", resp.Status)
	fmt.Println("data -->", resp.Data)
	fmt.Println("help --> ", resp.Help)
}

func TestClient_AltDiffusionSync(t *testing.T) {
	var err error
	altDiffusionRequest := &AltDiffusionRequest{
		Prompt: "黑暗精灵公主，非常详细，幻想，非常详细，数字绘画，概念艺术，敏锐的焦点，插图",
	}
	var requestId string
	if requestId, err = client.AltDiffusionAsync(altDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(requestId)
}

func TestClient_AltDiffusionAsyncRequestId(t *testing.T) {
	var err error
	var altDiffusionSyncResponse *AltDiffusionSyncResponse
	if altDiffusionSyncResponse, err = client.AltDiffusionAsyncRequestId(`3aa771a2-dfc2-4996-8b5b-776c86cd3384`); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(altDiffusionSyncResponse)
	var altDiffusionSyncInfoResponse AltDiffusionSyncInfoResponse
	if err = json.Unmarshal([]byte(altDiffusionSyncResponse.Info), &altDiffusionSyncInfoResponse); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(altDiffusionSyncInfoResponse)
}

// 连续语义增强机器翻译（CSANMT）
// 通过输入中文，可以获得英文翻译（CSANMT连续语义增强机器翻译-中英-通用领域）

func TestClient_CsanmtSync(t *testing.T) {
	csanmtRequest := CsanmtRequest{InputSequence: "你好，请问你是谁"}
	var csanmtResponse *CsanmtResponse
	var err error
	if csanmtResponse, err = client.CsanmtSync(&csanmtRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(csanmtResponse)
}

func TestClient_DctNetSync(t *testing.T) {
	var err error
	dCTNetRequest := DCTNetRequest{
		ImageUrl:    "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685184985115EseK2F6MkAPO09vQ.png?Expires=1685698590&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=S%2F6lGXunqDDL%2FDI0p1Qd6VeoWlQ%3D",
		ImageBase64: "",
	}
	var commonResponse *CommonResponse
	if commonResponse, err = client.DctNetSync(&dCTNetRequest, DCTNetApiNameCompound); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonResponse)
}

// 迪士尼风格扩散生成模型
// 本模型自 Stable Diffusion 1.5 微调而来，微调数据来自某著名动画工作室的电影截图

func TestClient_ModiAsync(t *testing.T) {
	commonStableDiffusionRequest := CommonRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonResponse
	if commonStableDiffusionResponse, err = client.ModiSync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err -> %v", err)
	}

	fmt.Println(commonStableDiffusionResponse)
}

func TestClient_ModiSync(t *testing.T) {
	commonStableDiffusionRequest := CommonRequest{
		Text: "dog",
	}
	var err error
	var requestId string
	if requestId, err = client.ModiAsync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err -> %v", err)
	}

	fmt.Println(requestId)
}

// 图像去噪与去模糊化（NAFNet）
// 适用于智能手机拍摄的带噪图片，通过该应用，可以进行图像的去燥与去模糊操作。

func TestClient_NafNetSync(t *testing.T) {
	var err error
	var commonResponse *CommonResponse
	nafNetRequest := &NafNetRequest{
		ImageUrl: "https://x-ai.oss-cn-guangzhou.aliyuncs.com/images/uploads/2023-05-31/bcf892f4334cf4526156a1957854ee6f.png?Expires=1685725215&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=0Gwd75Yd9mvZrSvcV7cmxb7MNhM%3D",
	}
	if commonResponse, err = client.NafNetSync(nafNetRequest, NafNetDeblur); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonResponse)
}

// 语音识别（Paraformer）
// 可以输入语音，输出对应的文字信息。Paraformer是达摩院语音团队提出的一种高效的非自回归端到端语音识别框架。
// 只支持 同步调用

func TestClient_ParaformerSync(t *testing.T) {
	var err error
	paraformerRequest := ParaformerRequest{AudioUrl: "https://isv-data.oss-cn-hangzhou.aliyuncs.com/ics/MaaS/ASR/test_audio/asr_example_zh.wav"}
	var paraformerResponse *ParaformerResponse
	if paraformerResponse, err = client.ParaformerSync(&paraformerRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(paraformerResponse)
}

// 图像去色带（RealESRGAN）
// 适用于因色彩精度不够导致色带的图像，能将色带去除。

func TestClient_RrdbSync(t *testing.T) {
	var err error
	var resp *CommonResponse
	req := &RrdbRequest{
		ImageUrl: "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685184807947dLCWr4BDegD7wvRf.png?Expires=1685725887&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=fh0f8VdWYT18ygTv%2B%2B6iNCb%2BYPc%3D",
	}
	if resp, err = client.RrdbSync(req); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(resp)
	fmt.Println("status --> ", resp.Status)
}

// VisualGLM-6B 是一个开源的，支持图像、中文和英文的多模态对话语言模型。

func TestClient_Visualglm6bAsync(t *testing.T) {
	// var err error
	visualglm6bRequest := &Visualglm6bRequest{
		Text:    "描述图片内容",
		History: "",
		Image:   "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685164392164Z0tt7lH32KAHNZjv.png?Expires=1685699444&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=gN8uMhhBXwALfhEcIypf49lMxnI%3D",
	}
	var err error
	var visualglm6bResponse *Visualglm6bResponse
	if visualglm6bResponse, err = client.Visualglm6bSync(visualglm6bRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(visualglm6bResponse)
}

func TestClient_Visualglm6bSync(t *testing.T) {
	var err error
	var resp *Vitgpt2Response
	req := Vitgpt2Request{
		ImageUrl: "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685163473641xIbXnSPMvVWxF6ee.png?Expires=1685726382&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=ujHff%2Fa7YUcw8EVchWD3tHHlIoQ%3D",
	}
	if resp, err = client.Vitgpt2Sync(&req); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(resp)
}

func TestClient_OcrSync(t *testing.T) {
	req := OrcRequest{
		Lang:     "ch",
		ImageUrl: "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-18/v3Gs21aigf48FZbi2E8d.png?Expires=1685774160&OSSAccessKeyId=TMP.3Kk9JYAgxS2AAF7r2Fpwdxht5AUaKyuCeZ15Ft2DPcn8Cgof95t26AXzs2fmCam4D3xXDCSgrqJQgT8PrptkGAJkw9ALyR&Signature=A0n3%2FXbDbOJIb5t66LhXyghpQNc%3D",
	}
	client.OcrSync(&req, OcrSimple)
}

type T struct {
	Status string              `json:"status"`
	Data   [][][][]interface{} `json:"data"`
}

type T2 struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
