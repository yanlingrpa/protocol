package basic

import "time"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FPoint struct {
	// 元素的X坐标，以像素为单位。
	X float64 `json:"x"`
	// 元素的Y坐标，以像素为单位。
	Y float64 `json:"y"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (s Size) Area() int {
	return s.Width * s.Height
}

type Task interface {
	GetGuid() string
	GetProjectId() string
	GetScriptId() string
	GetParams() map[string]any
	GetPayload() string
	GetRetryCount() int
	GetMaxRetries() int
	GetExpiredAt() time.Time
}

type DispatchTaskData struct {
	ProjectId  string         `json:"project_id"`  // 目标项目ID
	ScriptId   string         `json:"script_id"`   // 目标脚本ID
	DeviceId   string         `json:"device_id"`   // 目标设备ID, self表示本设备, 空表示不限设备, 其他表示指定设备
	Params     map[string]any `json:"params"`      // 任务参数
	Payload    string         `json:"payload"`     // 任务负载
	ExecTime   time.Time      `json:"exec_time"`   // 任务计划执行时间
	ExpiredAt  time.Time      `json:"expired_at"`  // 任务过期时间
	MaxRetries int            `json:"max_retries"` // 最大重试次数
}

type OcrText struct {
	Text       string  `json:"text"`       // 识别的文本内容
	Rect       Rect    `json:"rect"`       // 文本所在的矩形区域
	Confidence float64 `json:"confidence"` // 识别置信度
}

type OcrResult struct {
	Width  int       `json:"width"`  // 图像宽度
	Height int       `json:"height"` // 图像高度
	Texts  []OcrText `json:"texts"`  // 识别的文本列表
}
