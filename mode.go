package aigcaas

const (
	Default int = iota // 同步模式
	Polling            // 轮训模式
	Notice             // 通知模式
)
