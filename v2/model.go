package aigcaas

// CommonText2ImgRequest  文生图通用的请求参数
type CommonText2ImgRequest struct {
	EnableHr                          bool          `json:"enable_hr,,omitempty"`
	DenoisingStrength                 int           `json:"denoising_strength,,omitempty"`
	FirstphaseWidth                   int           `json:"firstphase_width,,omitempty"`
	FirstphaseHeight                  int           `json:"firstphase_height,,omitempty"`
	HrScale                           int           `json:"hr_scale,,omitempty"`
	HrUpscaler                        string        `json:"hr_upscaler,,omitempty"`
	HrSecondPassSteps                 int           `json:"hr_second_pass_steps,,omitempty"`
	HrResizeX                         int           `json:"hr_resize_x,,omitempty"`
	HrResizeY                         int           `json:"hr_resize_y,,omitempty"`
	HrSamplerName                     string        `json:"hr_sampler_name,,omitempty"`
	HrPrompt                          string        `json:"hr_prompt,,omitempty"`
	HrNegativePrompt                  string        `json:"hr_negative_prompt,,omitempty"`
	Prompt                            string        `json:"prompt,,omitempty"`
	Styles                            []string      `json:"styles,,omitempty"`
	Seed                              int           `json:"seed,,omitempty"`
	Subseed                           int           `json:"subseed,,omitempty"`
	SubseedStrength                   int           `json:"subseed_strength,,omitempty"`
	SeedResizeFromH                   int           `json:"seed_resize_from_h,,omitempty"`
	SeedResizeFromW                   int           `json:"seed_resize_from_w,,omitempty"`
	SamplerName                       string        `json:"sampler_name,,omitempty"`
	BatchSize                         int           `json:"batch_size,,omitempty"`
	NIter                             int           `json:"n_iter,,omitempty"`
	Steps                             int           `json:"steps,,omitempty"`
	CfgScale                          int           `json:"cfg_scale,,omitempty"`
	Width                             int           `json:"width,,omitempty"`
	Height                            int           `json:"height,,omitempty"`
	RestoreFaces                      bool          `json:"restore_faces,,omitempty"`
	Tiling                            bool          `json:"tiling,,omitempty"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples,,omitempty"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid,,omitempty"`
	NegativePrompt                    string        `json:"negative_prompt,,omitempty"`
	Eta                               int           `json:"eta,,omitempty"`
	SMinUncond                        int           `json:"s_min_uncond,,omitempty"`
	SChurn                            int           `json:"s_churn,,omitempty"`
	STmax                             int           `json:"s_tmax,,omitempty"`
	STmin                             int           `json:"s_tmin,,omitempty"`
	SNoise                            int           `json:"s_noise,,omitempty"`
	OverrideSettings                  struct{}      `json:"override_settings,,omitempty"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards,,omitempty"`
	ScriptArgs                        []interface{} `json:"script_args,,omitempty"`
	SamplerIndex                      string        `json:"sampler_index,,omitempty"`
	ScriptName                        string        `json:"script_name,,omitempty"`
	SendImages                        bool          `json:"send_images,,omitempty"`
	SaveImages                        bool          `json:"save_images,,omitempty"`
	AlwaysonScripts                   struct{}      `json:"alwayson_scripts,,omitempty"`
}

// CommonImg2imgRequest 图生图
type CommonImg2imgRequest struct {
	InitImages                        []string      `json:"init_images"`
	ResizeMode                        int           `json:"resize_mode"`
	DenoisingStrength                 float64       `json:"denoising_strength"`
	ImageCfgScale                     int           `json:"image_cfg_scale"`
	Mask                              string        `json:"mask"`
	MaskBlur                          int           `json:"mask_blur"`
	InpaintingFill                    int           `json:"inpainting_fill"`
	InpaintFullRes                    bool          `json:"inpaint_full_res"`
	InpaintFullResPadding             int           `json:"inpaint_full_res_padding"`
	InpaintingMaskInvert              int           `json:"inpainting_mask_invert"`
	InitialNoiseMultiplier            int           `json:"initial_noise_multiplier"`
	Prompt                            string        `json:"prompt"`
	Styles                            []string      `json:"styles"`
	Seed                              int           `json:"seed"`
	Subseed                           int           `json:"subseed"`
	SubseedStrength                   int           `json:"subseed_strength"`
	SeedResizeFromH                   int           `json:"seed_resize_from_h"`
	SeedResizeFromW                   int           `json:"seed_resize_from_w"`
	SamplerName                       string        `json:"sampler_name"`
	BatchSize                         int           `json:"batch_size"`
	NIter                             int           `json:"n_iter"`
	Steps                             int           `json:"steps"`
	CfgScale                          int           `json:"cfg_scale"`
	Width                             int           `json:"width"`
	Height                            int           `json:"height"`
	RestoreFaces                      bool          `json:"restore_faces"`
	Tiling                            bool          `json:"tiling"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid"`
	NegativePrompt                    string        `json:"negative_prompt"`
	Eta                               int           `json:"eta"`
	SMinUncond                        int           `json:"s_min_uncond"`
	SChurn                            int           `json:"s_churn"`
	STmax                             int           `json:"s_tmax"`
	STmin                             int           `json:"s_tmin"`
	SNoise                            int           `json:"s_noise"`
	OverrideSettings                  struct{}      `json:"override_settings"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
	ScriptArgs                        []interface{} `json:"script_args"`
	SamplerIndex                      string        `json:"sampler_index"`
	IncludeInitImages                 bool          `json:"include_init_images"`
	ScriptName                        string        `json:"script_name"`
	SendImages                        bool          `json:"send_images"`
	SaveImages                        bool          `json:"save_images"`
	AlwaysonScripts                   struct{}      `json:"alwayson_scripts"`
}

// CommonResponse 通用的响应参数
// 主要：请求成功只会响应 Aigcaas-Request-Id 参数
type CommonResponse struct {
	AigcaasRequestId string `json:"Aigcaas-Request-Id"`
	ErrorCommonResponse
}

// ErrorCommonResponse 通用的错误响应参数
type ErrorCommonResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Help    string `json:"help"`
}

// AsyncResponse 请求结果的响应参数
type AsyncResponse struct {
	Images     []string `json:"images"`
	Parameters struct {
		EnableHr          bool        `json:"enable_hr"`
		DenoisingStrength int         `json:"denoising_strength"`
		FirstphaseWidth   int         `json:"firstphase_width"`
		FirstphaseHeight  int         `json:"firstphase_height"`
		HrScale           float64     `json:"hr_scale"`
		HrUpscaler        interface{} `json:"hr_upscaler"`
		HrSecondPassSteps int         `json:"hr_second_pass_steps"`
		HrResizeX         int         `json:"hr_resize_x"`
		HrResizeY         int         `json:"hr_resize_y"`
		HrSamplerName     interface{} `json:"hr_sampler_name"`
		HrPrompt          string      `json:"hr_prompt"`
		HrNegativePrompt  string      `json:"hr_negative_prompt"`
		Prompt            string      `json:"prompt"`
		Styles            interface{} `json:"styles"`
		Seed              int         `json:"seed"`
		Subseed           int         `json:"subseed"`
		SubseedStrength   int         `json:"subseed_strength"`
		SeedResizeFromH   int         `json:"seed_resize_from_h"`
		SeedResizeFromW   int         `json:"seed_resize_from_w"`
		SamplerName       interface{} `json:"sampler_name"`
		BatchSize         int         `json:"batch_size"`
		NIter             int         `json:"n_iter"`
		Steps             int         `json:"steps"`
		CfgScale          float64     `json:"cfg_scale"`
		Width             int         `json:"width"`
		Height            int         `json:"height"`
		RestoreFaces      bool        `json:"restore_faces"`
		Tiling            bool        `json:"tiling"`
		DoNotSaveSamples  bool        `json:"do_not_save_samples"`
		DoNotSaveGrid     bool        `json:"do_not_save_grid"`
		NegativePrompt    interface{} `json:"negative_prompt"`
		Eta               interface{} `json:"eta"`
		SMinUncond        float64     `json:"s_min_uncond"`
		SChurn            float64     `json:"s_churn"`
		STmax             interface{} `json:"s_tmax"`
		STmin             float64     `json:"s_tmin"`
		SNoise            float64     `json:"s_noise"`
		OverrideSettings  struct {
		} `json:"override_settings"`
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
	ErrorCommonResponse
}

// AsyncInfoResponse info响应参数
type AsyncInfoResponse struct {
	Prompt                        string        `json:"prompt"`
	AllPrompts                    []string      `json:"all_prompts"`
	NegativePrompt                string        `json:"negative_prompt"`
	AllNegativePrompts            []string      `json:"all_negative_prompts"`
	Seed                          int           `json:"seed"`
	AllSeeds                      []int         `json:"all_seeds"`
	Subseed                       int           `json:"subseed"`
	AllSubseeds                   []int         `json:"all_subseeds"`
	SubseedStrength               int           `json:"subseed_strength"`
	Width                         int           `json:"width"`
	Height                        int           `json:"height"`
	SamplerName                   string        `json:"sampler_name"`
	CfgScale                      float64       `json:"cfg_scale"`
	Steps                         int           `json:"steps"`
	BatchSize                     int           `json:"batch_size"`
	RestoreFaces                  bool          `json:"restore_faces"`
	FaceRestorationModel          interface{}   `json:"face_restoration_model"`
	SdModelHash                   string        `json:"sd_model_hash"`
	SeedResizeFromW               int           `json:"seed_resize_from_w"`
	SeedResizeFromH               int           `json:"seed_resize_from_h"`
	DenoisingStrength             int           `json:"denoising_strength"`
	ExtraGenerationParams         struct{}      `json:"extra_generation_params"`
	IndexOfFirstImage             int           `json:"index_of_first_image"`
	Infotexts                     []string      `json:"infotexts"`
	Styles                        []interface{} `json:"styles"`
	JobTimestamp                  string        `json:"job_timestamp"`
	ClipSkip                      int           `json:"clip_skip"`
	IsUsingInpaintingConditioning bool          `json:"is_using_inpainting_conditioning"`
}
