package extention

import "github.com/yanlingrpa/protocol/basic"

type OcrWord struct {
	Text       string         `json:"text"`       // 识别的文本内容
	Confidence float64        `json:"confidence"` // 置信度
	Box        [4]basic.Point `json:"box"`        // 边界框坐标
}

type OcrResult struct {
	Width  int       `json:"width"`  // 图像宽度
	Height int       `json:"height"` // 图像高度
	Words  []OcrWord `json:"words"`  // 识别的单词列表
}

type VisionExtension interface {
	Ocr(image []byte) (*OcrResult, error)                         // 对图片进行OCR识别
	Locate(image []byte, text string) ([]basic.Rect, error)       // 对图片进行文本定位，返回文本区域的边界框列表
	Detect(image []byte, prompt, question string) (string, error) // 对图片进行视觉理解，返回json结果
	Read(image []byte, prompt, question string) (string, error)   // 对图片进行视觉阅读，返回json结果
}
