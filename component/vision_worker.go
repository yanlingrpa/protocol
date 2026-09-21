package component

import "yanlingrpa.com/yanling/protocol/basic"

type VisionWorker interface {
	/*
	* 对图像执行 OCR 识别。
	* image: 图像的字节数组。
	* confidence: OCR 结果的最低置信度阈值；低于该阈值的结果将被过滤掉。
	* min: 要识别的文本最小尺寸；小于该尺寸的文本将被忽略。
	* max: 要识别的文本最大尺寸；大于该尺寸的文本将被忽略。
	* return: OCR 结果，包含识别出的文本及对应位置信息。
	 */
	Ocr(image []byte, confidence float64, min, max basic.Size) (*basic.OcrResult, error)

	/*
	* 对图像执行视觉定位。
	* image: 图像的字节数组。
	* texts: 要定位的文本列表，可包含一个或多个文本值。
	* confidence: 定位结果的最低置信度阈值；低于该阈值的结果将被过滤掉。
	* min: 要定位的文本最小尺寸；小于该尺寸的文本将被忽略。
	* max: 要定位的文本最大尺寸；大于该尺寸的文本将被忽略。
	* return: 定位结果，包含所有匹配文本的位置信息，
	* 按从左到右、从上到下排序；未找到的文本将从结果中排除。
	 */
	Locate(image []byte, texts []string, confidence float64, min, max basic.Size) ([]basic.Rect, error)

	/*
	* 对图像执行视觉检测。
	* image: 图像的字节数组。
	* description: 目标检测说明，清晰描述目标及要求，
	* confidence: 检测结果的最低置信度阈值；低于该阈值的结果将被过滤掉。
	* min: 要检测目标的最小尺寸；小于该尺寸的目标将被忽略。
	* max: 要检测目标的最大尺寸；大于该尺寸的目标将被忽略。
	* 例如："检测图像中的所有人脸，并返回每张脸的年龄和性别"。
	* schema: 检测结果的 JSON Schema 定义，描述结构和字段，
	* 以便调用方解析和使用返回数据。
	* return: 检测结果，包含根据说明从图像中检测到的位置和属性信息。
	 */
	Detect(image []byte, description string, confidence float64, min, max basic.Size) ([]basic.Rect, error)

	/*
	* 对图像执行视觉读取。
	* image: 图像的字节数组。
	* description: 阅读说明，明确描述要读取的文本内容，
	* 例如："读取图像中的所有文本，并返回每段文本的内容和位置"。
	* confidence: 阅读结果的最低置信度阈值；低于该阈值的结果将被过滤掉。
	* min: 要读取的文本最小尺寸；小于该尺寸的文本将被忽略。
	* max: 要读取的文本最大尺寸；大于该尺寸的文本将被忽略。
	* 例如每段文本的起始和结束位置。
	* return: 阅读结果，包含根据说明从图像中读取的文本内容和位置信息。
	 */
	Read(image []byte, description string, confidence float64, min, max basic.Size) (*basic.OcrResult, error)
}
