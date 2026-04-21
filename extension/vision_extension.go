package extension

import "yanlingrpa.com/yanling/protocol/basic"

type VisionExtension interface {
	/*
	* 对图片进行OCR识别
	* image: 图片的字节数组
	* return: OCR识别结果，包含识别出的文本和对应的位置信息
	 */
	Ocr(image []byte) (*basic.OcrResult, error)

	/*
	* 对图片进行视觉定位
	* image: 图片的字节数组
	* texts: 需要定位的文本列表，可以是一个或多个文本
	* return: 定位结果，包含所有文本的位置信息, 从左到右从上到下排序，如果文本未找到则不包含在结果中
	 */
	Locate(image []byte, texts ...string) ([]basic.Rect, error)

	/*
	* 对图片进行视觉识别
	* image: 图片的字节数组
	* instruction: 识别指令，需要描述清楚识别的目标和要求，例如“请识别图片中的所有人脸，并返回每个人脸的年龄和性别”
	* schema: 识别结果的json schema定义，描述了识别结果的结构和字段信息，便于调用方解析和使用识别结果
	* return: 符合schema定义的识别结果，通常是一个json字符串，包含了根据instruction从图片中识别出的信息
	 */
	Detect(image []byte, instruction, schema string) (string, error)

	/*
	* 对图片进行视觉阅读
	* image: 图片的字节数组
	* instruction: 阅读指令，需要描述清楚需要阅读的文本，例如“请阅读图片中的所有文本，并返回所有文本的内容和位置信息”
	* schema: 阅读结果的json schema定义，描述了阅读结果结构，例如每个文本的起始位置和结束位置
	* return: 符合schema定义的阅读结果，通常为一个json字符串，包含了根据instruction从图片中阅读出的信息
	 */
	Read(image []byte, instruction, schema string) (string, error)
}
