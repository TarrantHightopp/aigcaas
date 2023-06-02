package aigcaas

import (
	"fmt"
	"testing"
)

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
