package basic

/*
* Point represents a 2D integer coordinate point.
 */
type Point struct {
	/*
	* X is the horizontal coordinate.
	 */
	X int `json:"x"`
	/*
	* Y is the vertical coordinate.
	 */
	Y int `json:"y"`
}

/*
* FPoint represents a 2D floating-point coordinate point.
 */
type FPoint struct {
	/*
	* X coordinate of the element, in pixels.
	 */
	X float64 `json:"x"`
	/*
	* Y coordinate of the element, in pixels.
	 */
	Y float64 `json:"y"`
}

/*
* Size represents width and height.
 */
type Size struct {
	/*
	* Width represents the width.
	 */
	Width int `json:"width"`
	/*
	* Height represents the height.
	 */
	Height int `json:"height"`
}

/*
* Area returns the area of the size.
 */
func (s Size) Area() int {
	return s.Width * s.Height
}

/*
* OcrText represents a single OCR text recognition result.
 */
type OcrText struct {
	/*
	* Text is the recognized text content.
	 */
	Text string `json:"text"`
	/*
	* Rect is the rectangular area where the text is located.
	 */
	Rect Rect `json:"rect"`
	/*
	* Confidence is the recognition confidence score.
	 */
	Confidence float64 `json:"confidence"`
}

/*
* OcrResult represents the OCR recognition result of the entire image.
 */
type OcrResult struct {
	/*
	* Width is the image width.
	 */
	Width int `json:"width"`
	/*
	* Height is the image height.
	 */
	Height int `json:"height"`
	/*
	* Texts is the list of recognized text items.
	 */
	Texts []OcrText `json:"texts"`
}
