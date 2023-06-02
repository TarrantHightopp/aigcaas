package aigcaas

// CommonRequest 共通的StableDiffusion模型请求参数
type CommonRequest struct {
	Text              string  `json:"text,omitempty"`                // 内容
	Height            int     `json:"height,omitempty"`              // 高度
	Width             int     `json:"width,omitempty"`               // 宽度
	NegativePrompt    string  `json:"negative_prompt,omitempty"`     // 反面描述内容
	GuidanceScale     int     `json:"guidance_scale,omitempty"`      // 控制图片与提示词的相关程度
	Seed              int     `json:"seed,omitempty"`                // 控制生成图片是否可复现
	NumInferenceSteps int     `json:"num_inference_steps,omitempty"` // 推理步数 推理过程给图想去噪音，数值越大效果越好，速度越慢，本平台限制(5,40)
	Eta               float32 `json:"eta,omitempty"`                 // 噪声种子（取值范围0.0-1.0）
}

// CommonResponse 共通的StableDiffusion模型响应参数
type CommonResponse struct {
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

// AsyncRequestIdResponse 异步请求request id 响应参数
type AsyncRequestIdResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Help    string `json:"help"`
	Data    string `json:"data"`
}

// AltDiffusionRequest AltDiffusion 请求参数
type AltDiffusionRequest struct {
	Prompt            string  `json:"prompt"`                         // 内容
	Height            int     `json:"height,omitempty"`               // 高度
	Width             int     `json:"width,omitempty"`                // 宽度
	NegativePrompt    string  `json:"negative_prompt,omitempty"`      // 反面描述内容
	GuidanceScale     int     `json:"guidance_scale,omitempty"`       // 控制图片与提示词的相关程度
	Seed              int     `json:"seed,omitempty"`                 // 控制生成图片是否可复现
	NumInferenceSteps int     `json:"num_inference_steps,omitempty"`  // 推理步数 推理过程给图想去噪音，数值越大效果越好，速度越慢，本平台限制(5,40)
	Eta               float32 `json:"eta,omitempty"`                  // 噪声种子（取值范围0.0-1.0）
	EnableHr          bool    `json:"enable_hr,omitempty"`            // 高清修复（Hires. fix） 可以提升图片的画质，但是会耗费更多VRAM
	DenoisingStrength float32 `json:"denoising_strength,omitempty"`   // 降噪强度  降噪强度，取值范围是0-1，默认设置0.75。取值越大，说明图片变化越大，0表示图片几乎不变，1表示可能严重偏离原图。一般将该参数设置在0.6~0.8范围；
	HrScale           float32 `json:"hr_scale,omitempty"`             // 缩放系数 缩放系数；默认为2
	HrUpscaler        string  `json:"hr_upscaler,omitempty"`          // 要使用的放大器 高清修复模式下要使用的放大器
	HrSecondPassSteps float32 `json:"hr_second_pass_steps,omitempty"` // 高清修复步数
	HrResizeX         float32 `json:"hr_resize_x,omitempty"`          // 高清修复制定尺寸 x
	HrResizeY         float32 `json:"hr_resize_y,omitempty"`          // 高清修复制定尺寸 y
	HrPrompt          string  `json:"hr_prompt,omitempty"`            // 高清修复提示词
	HrNegativePrompt  string  `json:"hr_negative_prompt,omitempty"`   // 高清修复反面描述内容
	Tiling            string  `json:"tiling,omitempty"`               // 可平埔 生成可以平铺的周期性图像，默认为false
	ScriptName        string  `json:"script_name,omitempty"`          // 脚本/命令名称 取值有prompt matrix（会生出一个表格图片，用于比对不同提示词生图的效果）, prompts from file or textbox（从写好提示词的文件生成图片）, x/y/z plot（用于比对不同提示词、采样方法、CFG Scale、种子码的组合所生图的效果），通常与脚本/命令参数（script_args）搭配使用
	ScriptArgs        string  `json:"script_args,omitempty"`          // 脚本/命令参数 一般用于 lora模型或其他插件参数，例如：[0, true, true, "LoRA", "dingzhenlora_v1(fa7c1732cc95)", 1, 1]
	SamplerIndex      string  `json:"sampler_index,omitempty"`        // 采样方法，默认为Euler，其他可选参数：Euler a, Euler, LMS, Heun, DPM2, DPM2 a, DPM++ 2S a, DPM++ 2M, DPM++ SDE, DPM++ 2M SDE, DPM fast, DPM adaptive, LMS Karras, DPM2 Karras, DPM2 a Karras, DPM++ 2S a Karras, DPM++ 2M Karras, DPM++ SDE Karras, DPM++ 2M SDE Karras, DDIM, PLMS, UniPC
	Steps             int     `json:"steps,omitempty"`                // 生成步数 生成步数，默认50
	Niter             int     `json:"n_iter,omitempty"`               // 生成批次 生成批次，默认为1；
	RestoreFaces      bool    `json:"restore_faces,omitempty"`        // 恢复面部缺陷 应用了额外模型，该模型可以恢复面部缺陷，默认为 false
	Subseed           int     `json:"subseed,omitempty"`              // 附加种子值 附加种子值，测试更多种子码变化之用
	SubseedStrength   int     `json:"subseed_strength,omitempty"`     // 种子和variation seed之间的差值程度，如果生成了两张图，可以通过第一个图为seed，第二个图为subseed，然后设置subseed strength在0-1之间的强度变化，那么会生成逐渐过渡的两个图。
	SeedResizeFromW   int     `json:"seed_resize_from_w,omitempty"`   // 对应 UI 界面的 Resize seed from width；默认为-1
	SeedResizeFromH   int     `json:"seed_resize_from_h,omitempty"`   // 对应 UI 界面的 Resize seed from height；默认为-1
	CfgScale          int     `json:"cfg_scale,omitempty"`            // AI生图与你给的提示词的相关度，数值越高越会按照你说的内容下去生图。控制模型和prompt的匹配程度，1.忽视prompt， 3.更有创意，7，在prompt和freedom中取得平衡，15，遵守prompt，30，严格按照prompt
	Styles            string  `json:"styles,omitempty"`               // 风格
}

// AltDiffusionSyncResponse 异步响应
type AltDiffusionSyncResponse struct {
	Images     []string `json:"images"`
	Parameters struct {
		EnableHr                          bool          `json:"enable_hr"`
		DenoisingStrength                 int           `json:"denoising_strength"`
		FirstphaseWidth                   int           `json:"firstphase_width"`
		FirstphaseHeight                  int           `json:"firstphase_height"`
		HrScale                           float64       `json:"hr_scale"`
		HrUpscaler                        interface{}   `json:"hr_upscaler"`
		HrSecondPassSteps                 int           `json:"hr_second_pass_steps"`
		HrResizeX                         int           `json:"hr_resize_x"`
		HrResizeY                         int           `json:"hr_resize_y"`
		HrSamplerName                     interface{}   `json:"hr_sampler_name"`
		HrPrompt                          string        `json:"hr_prompt"`
		HrNegativePrompt                  string        `json:"hr_negative_prompt"`
		Prompt                            string        `json:"prompt"`
		Styles                            interface{}   `json:"styles"`
		Seed                              int           `json:"seed"`
		Subseed                           int           `json:"subseed"`
		SubseedStrength                   int           `json:"subseed_strength"`
		SeedResizeFromH                   int           `json:"seed_resize_from_h"`
		SeedResizeFromW                   int           `json:"seed_resize_from_w"`
		SamplerName                       interface{}   `json:"sampler_name"`
		BatchSize                         int           `json:"batch_size"`
		NIter                             int           `json:"n_iter"`
		Steps                             int           `json:"steps"`
		CfgScale                          float64       `json:"cfg_scale"`
		Width                             int           `json:"width"`
		Height                            int           `json:"height"`
		RestoreFaces                      bool          `json:"restore_faces"`
		Tiling                            bool          `json:"tiling"`
		DoNotSaveSamples                  bool          `json:"do_not_save_samples"`
		DoNotSaveGrid                     bool          `json:"do_not_save_grid"`
		NegativePrompt                    interface{}   `json:"negative_prompt"`
		Eta                               interface{}   `json:"eta"`
		SMinUncond                        float64       `json:"s_min_uncond"`
		SChurn                            float64       `json:"s_churn"`
		STmax                             interface{}   `json:"s_tmax"`
		STmin                             float64       `json:"s_tmin"`
		SNoise                            float64       `json:"s_noise"`
		OverrideSettings                  interface{}   `json:"override_settings"`
		OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
		ScriptArgs                        []interface{} `json:"script_args"`
		SamplerIndex                      string        `json:"sampler_index"`
		ScriptName                        interface{}   `json:"script_name"`
		SendImages                        bool          `json:"send_images"`
		SaveImages                        bool          `json:"save_images"`
		AlwaysonScripts                   struct {
		} `json:"alwayson_scripts"`
	} `json:"parameters"`
	Info string `json:"info"`
}

// AltDiffusionSyncInfoResponse  AltDiffusionSyncResponse中info的结构体
type AltDiffusionSyncInfoResponse struct {
	Prompt                string      `json:"prompt"`
	AllPrompts            []string    `json:"all_prompts"`
	NegativePrompt        string      `json:"negative_prompt"`
	AllNegativePrompts    []string    `json:"all_negative_prompts"`
	Seed                  int         `json:"seed"`
	AllSeeds              []int       `json:"all_seeds"`
	Subseed               int         `json:"subseed"`
	AllSubseeds           []int       `json:"all_subseeds"`
	SubseedStrength       int         `json:"subseed_strength"`
	Width                 int         `json:"width"`
	Height                int         `json:"height"`
	SamplerName           string      `json:"sampler_name"`
	CfgScale              float64     `json:"cfg_scale"`
	Steps                 int         `json:"steps"`
	BatchSize             int         `json:"batch_size"`
	RestoreFaces          bool        `json:"restore_faces"`
	FaceRestorationModel  interface{} `json:"face_restoration_model"`
	SdModelHash           string      `json:"sd_model_hash"`
	SeedResizeFromW       int         `json:"seed_resize_from_w"`
	SeedResizeFromH       int         `json:"seed_resize_from_h"`
	DenoisingStrength     int         `json:"denoising_strength"`
	ExtraGenerationParams struct {
	} `json:"extra_generation_params"`
	IndexOfFirstImage             int           `json:"index_of_first_image"`
	Infotexts                     []string      `json:"infotexts"`
	Styles                        []interface{} `json:"styles"`
	JobTimestamp                  string        `json:"job_timestamp"`
	ClipSkip                      int           `json:"clip_skip"`
	IsUsingInpaintingConditioning bool          `json:"is_using_inpainting_conditioning"`
}

// DCTNetRequest 请求参数
type DCTNetRequest struct {
	ImageUrl    string `json:"image_url,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
}

// Visualglm6bRequest Visualglm6b请求参数
type Visualglm6bRequest struct {
	Text    string `json:"text"`
	History string `json:"history,omitempty"`
	Image   string `json:"image"`
}

// Visualglm6bResponse Visualglm6b响应参数
type Visualglm6bResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Help    string `json:"help"`
	Data    struct {
		Response string     `json:"response"`
		History  [][]string `json:"history"`
	} `json:"data"`
}

// FacefuRequest facefu 请求参数
type FacefuRequest struct {
	TemplatePathUrl    string `json:"template_path_url"`    // 原图图片地址
	UserPathUrl        string `json:"user_path_url"`        // 即将融合头像图片地址
	UserPathBase64     string `json:"user_path_base64"`     // 即将融合头像图片Base64
	TemplatePathBase64 string `json:"template_path_base64"` // 原图图片Base64
}
