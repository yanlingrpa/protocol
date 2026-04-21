package extension

import "yanlingrpa.com/yanling/protocol/basic"

type VisionExtension interface {
	/*
	* Performs OCR recognition on an image.
	* image: Byte array of the image.
	* return: OCR result containing recognized text and corresponding position information.
	 */
	Ocr(image []byte) (*basic.OcrResult, error)

	/*
	* Performs visual localization on an image.
	* image: Byte array of the image.
	* texts: List of texts to locate; can include one or more text values.
	* return: Localization result containing position information of all matched texts,
	* sorted from left to right and top to bottom; texts not found are excluded from the result.
	 */
	Locate(image []byte, texts ...string) ([]basic.Rect, error)

	/*
	* Performs visual detection on an image.
	* image: Byte array of the image.
	* instruction: Detection instruction that clearly describes the target and requirements,
	* for example, "Detect all faces in the image and return the age and gender of each face".
	* schema: JSON Schema definition of the detection result, describing structure and fields
	* so callers can parse and use the returned data.
	* return: Detection result conforming to the schema, usually a JSON string containing
	* information detected from the image according to the instruction.
	 */
	Detect(image []byte, instruction, schema string) (string, error)

	/*
	* Performs visual reading on an image.
	* image: Byte array of the image.
	* instruction: Reading instruction that clearly describes what text to read,
	* for example, "Read all text in the image and return the content and position of each text".
	* schema: JSON Schema definition of the reading result, describing the result structure,
	* such as start and end positions for each text segment.
	* return: Reading result conforming to the schema, usually a JSON string containing
	* information read from the image according to the instruction.
	 */
	Read(image []byte, instruction, schema string) (string, error)
}
