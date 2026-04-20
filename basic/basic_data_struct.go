package basic

/*
* Point 表示一个二维整数坐标点
 */
type Point struct {
	/*
	* X 表示横坐标
	 */
	X int `json:"x"`
	/*
	* Y 表示纵坐标
	 */
	Y int `json:"y"`
}

/*
* FPoint 表示一个二维浮点坐标点
 */
type FPoint struct {
	/*
	* 元素的 X 坐标，以像素为单位
	 */
	X float64 `json:"x"`
	/*
	* 元素的 Y 坐标，以像素为单位
	 */
	Y float64 `json:"y"`
}

/*
* Size 表示宽度和高度
 */
type Size struct {
	/*
	* Width 表示宽度
	 */
	Width int `json:"width"`
	/*
	* Height 表示高度
	 */
	Height int `json:"height"`
}

/*
* Area 返回尺寸的面积
 */
func (s Size) Area() int {
	return s.Width * s.Height
}

/*
* OcrText 表示单条 OCR 文本识别结果
 */
type OcrText struct {
	/*
	* Text 表示识别出的文本内容
	 */
	Text string `json:"text"`
	/*
	* Rect 表示文本所在的矩形区域
	 */
	Rect Rect `json:"rect"`
	/*
	* Confidence 表示识别置信度
	 */
	Confidence float64 `json:"confidence"`
}

/*
* OcrResult 表示整张图像的 OCR 识别结果
 */
type OcrResult struct {
	/*
	* Width 表示图像宽度
	 */
	Width int `json:"width"`
	/*
	* Height 表示图像高度
	 */
	Height int `json:"height"`
	/*
	* Texts 表示识别出的文本列表
	 */
	Texts []OcrText `json:"texts"`
}
