package component

import "yanlingrpa.com/yanling/protocol/basic"

type OcrWorker interface {
	/*
	* Performs OCR recognition on an image.
	* image: Byte array of the image.
	* return: OCR result containing recognized text and corresponding position information.
	 */
	OcrImage(image []byte) (*basic.OcrResult, error)
}
