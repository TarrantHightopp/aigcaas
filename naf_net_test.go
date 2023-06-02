package aigcaas

import (
	"fmt"
	"testing"
)

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
