package aigcaas

import (
	"fmt"
	"testing"
)

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
