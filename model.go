package aigcaas

type CommonStableDiffusionRequest struct {
	Text              string  `json:"text,omitempty"`                // 内容
	Height            int     `json:"height,omitempty"`              // 高度
	Width             int     `json:"width,omitempty"`               // 宽度
	NegativePrompt    string  `json:"negative_prompt,omitempty"`     // 反面描述内容
	GuidanceScale     int     `json:"guidance_scale,omitempty"`      // 控制图片与提示词的相关程度
	Seed              int     `json:"seed,omitempty"`                // 控制生成图片是否可复现
	NumInferenceSteps int     `json:"num_inference_steps,omitempty"` // 推理步数 推理过程给图想去噪音，数值越大效果越好，速度越慢，本平台限制(5,40)
	Eta               float32 `json:"eta,omitempty"`                 // 噪声种子（取值范围0.0-1.0）
}

type CommonStableDiffusionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Detail  []struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
		Ctx  struct {
			LimitValue int `json:"limit_value"`
		} `json:"ctx"`
	} `json:"detail"`
}
