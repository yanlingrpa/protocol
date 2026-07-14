package component

import "yanlingrpa.com/yanling/protocol/basic"

type VisionWorker interface {
	/*
	* Performs OCR recognition on an image.
	* image: Byte array of the image.
	* confidence: Minimum confidence threshold for OCR results; results below this threshold are filtered out.
	* min: Minimum size of text to recognize; smaller text is ignored.
	* max: Maximum size of text to recognize; larger text is ignored.
	* return: OCR result containing recognized text and corresponding position information.
	 */
	Ocr(image []byte, confidence float64, min, max basic.Size) (*basic.OcrResult, error)

	/*
	* Performs visual localization on an image.
	* image: Byte array of the image.
	* texts: List of texts to locate; can include one or more text values.
	* confidence: Minimum confidence threshold for localization results; results below this threshold are filtered out.
	* min: Minimum size of text to locate; smaller text is ignored.
	* max: Maximum size of text to locate; larger text is ignored.
	* return: Localization result containing position information of all matched texts,
	* sorted from left to right and top to bottom; texts not found are excluded from the result.
	 */
	Locate(image []byte, texts []string, confidence float64, min, max basic.Size) ([]basic.Rect, error)

	/*
	* Performs visual detection on an image.
	* image: Byte array of the image.
	* description: Detection instruction that clearly describes the target and requirements,
	* confidence: Minimum confidence threshold for detection results; results below this threshold are filtered out.
	* min: Minimum size of the target to detect; smaller targets are ignored.
	* max: Maximum size of the target to detect; larger targets are ignored.
	* for example, "Detect all faces in the image and return the age and gender of each face".
	* schema: JSON Schema definition of the detection result, describing structure and fields
	* so callers can parse and use the returned data.
	* return: Detection result containing position and attribute
	* information detected from the image according to the instruction.
	 */
	Detect(image []byte, description string, confidence float64, min, max basic.Size) ([]basic.Rect, error)

	/*
	* Performs visual reading on an image.
	* image: Byte array of the image.
	* description: Reading instruction that clearly describes what text to read,
	* for example, "Read all text in the image and return the content and position of each text".
	* confidence: Minimum confidence threshold for reading results; results below this threshold are filtered out.
	* min: Minimum size of text to read; smaller text is ignored.
	* max: Maximum size of text to read; larger text is ignored.
	* such as start and end positions for each text segment.
	* return: Reading result containing text content and position
	* information read from the image according to the instruction.
	 */
	Read(image []byte, description string, confidence float64, min, max basic.Size) (*basic.OcrResult, error)
}
