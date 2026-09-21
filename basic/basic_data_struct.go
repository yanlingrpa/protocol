package basic

/*
* Point 表示二维整数坐标点。
 */
type Point struct {
	/*
	* X 是横坐标。
	 */
	X int `json:"x"`
	/*
	* Y 是纵坐标。
	 */
	Y int `json:"y"`
}

/*
* FPoint 表示二维浮点坐标点。
 */
type FPoint struct {
	/*
	* X 是元素的横坐标，单位为像素。
	 */
	X float64 `json:"x"`
	/*
	* Y 是元素的纵坐标，单位为像素。
	 */
	Y float64 `json:"y"`
}

/*
* Size 表示宽度和高度。
 */
type Size struct {
	/*
	* Width 表示宽度。
	 */
	Width int `json:"width"`
	/*
	* Height 表示高度。
	 */
	Height int `json:"height"`
}

/*
* Area 返回该尺寸的面积。
 */
func (s Size) Area() int {
	return s.Width * s.Height
}

/*
* OcrText 表示单条 OCR 文本识别结果。
 */
type OcrText struct {
	/*
	* Text 是识别出的文本内容。
	 */
	Text string `json:"text"`
	/*
	* Rect 是文本所在的矩形区域。
	 */
	Rect Rect `json:"rect"`
	/*
	* Confidence 是识别置信度分数。
	 */
	Confidence float64 `json:"confidence"`
}

/*
* OcrResult 表示整张图片的 OCR 识别结果。
 */
type OcrResult struct {
	/*
	* Width 是图片宽度。
	 */
	Width int `json:"width"`
	/*
	* Height 是图片高度。
	 */
	Height int `json:"height"`
	/*
	* Texts 是识别出的文本列表。
	 */
	Texts []OcrText `json:"texts"`
}
