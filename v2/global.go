package aigcaas

const (
	_        int = iota
	TEXT2IMG     // 文生图
	IMG2IMG      // 图生图
)

const (
	Default int = iota // 同步模式
	Polling            // 轮训模式
	Notice             // 通知模式
)

const (
	ApiNameTEXT2IMG = "text2img"
	ApiNameIMG2IMG  = "img2img"
)

var URL = `https://api.aigcaas.cn/v2/application`

var AsyncRequestIdURL = `https://api.aigcaas.cn/v2/async`

const (
	ResponseSuccessCode = 200
)
