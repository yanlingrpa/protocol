package basic

import "sort"

/*
* Rect 表示一个矩形区域，定义了左上角坐标和宽高
 */
type Rect struct {
	/*
	* X 表示矩形左上角的 X 坐标
	 */
	X int `json:"x"`
	/*
	* Y 表示矩形左上角的 Y 坐标
	 */
	Y int `json:"y"`
	/*
	* Width 表示矩形宽度
	 */
	Width int `json:"width"`
	/*
	* Height 表示矩形高度
	 */
	Height int `json:"height"`
}

/*
* Position 返回矩形的左上角位置
 */
func (r Rect) Position() Point {
	return Point{
		X: r.X,
		Y: r.Y,
	}
}

/*
* Size 返回矩形的宽高信息
 */
func (r Rect) Size() Size {
	return Size{
		Width:  r.Width,
		Height: r.Height,
	}
}

/*
* Compare 比较两个矩形的大小关系，返回 -1 表示小于，0 表示等于，1 表示大于
* 比较优先级：面积 > X 坐标 > Y 坐标 > 宽度 > 高度
 */
func (r Rect) Compare(another Rect) int {
	if r.Size().Area() != another.Size().Area() {
		if r.Size().Area() < another.Size().Area() {
			return -1
		}
		return 1
	}
	if r.X != another.X {
		if r.X < another.X {
			return -1
		}
		return 1
	}
	if r.Y != another.Y {
		if r.Y < another.Y {
			return -1
		}
		return 1
	}
	if r.Width != another.Width {
		if r.Width < another.Width {
			return -1
		}
		return 1
	}
	if r.Height != another.Height {
		if r.Height < another.Height {
			return -1
		}
		return 1
	}
	return 0
}

/*
* SubRect 从矩形内部提取一个子矩形，指定子矩形的相对位置和大小
* 自动检查边界，防止越界
 */
func (r Rect) SubRect(x, y, width, height int) Rect {
	if x < 0 {
		x = 0
	}
	if x > r.Width-1 {
		x = r.Width - 1
	}
	if y < 0 {
		y = 0
	}
	if y > r.Height-1 {
		y = r.Height - 1
	}
	width = min(width, r.Width-x)
	height = min(height, r.Height-y)
	return Rect{
		X:      r.X + x,
		Y:      r.Y + y,
		Width:  width,
		Height: height,
	}
}

/*
* LeftPercent 从左边保留指定百分比的宽度，返回左侧矩形
 */
func (r Rect) LeftPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.LeftPixel(int(float32(r.Width) * percent))
}

/*
* LeftPixel 从左边保留指定像素宽度，返回左侧矩形
 */
func (r Rect) LeftPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Width {
		pixel = r.Width
	}
	return Rect{
		X:      r.X,
		Y:      r.Y,
		Width:  pixel,
		Height: r.Height,
	}
}

/*
* RightPercent 从右边保留指定百分比的宽度，返回右侧矩形
 */
func (r Rect) RightPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.RightPixel(int(float32(r.Width) * percent))
}

/*
* RightPixel 从右边保留指定像素宽度，返回右侧矩形
 */
func (r Rect) RightPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Width {
		pixel = r.Width
	}
	return Rect{
		X:      r.X + r.Width - pixel,
		Y:      r.Y,
		Width:  pixel,
		Height: r.Height,
	}
}

/*
* HeaderPercent 从顶部保留指定百分比的高度，返回顶部矩形
 */
func (r Rect) HeaderPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.HeaderPixel(int(float32(r.Height) * percent))
}

/*
* HeaderPixel 从顶部保留指定像素高度，返回顶部矩形
 */
func (r Rect) HeaderPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Height {
		pixel = r.Height
	}
	return Rect{
		X:      r.X,
		Y:      r.Y,
		Width:  r.Width,
		Height: pixel,
	}
}

/*
* FooterPercent 从底部保留指定百分比的高度，返回底部矩形
 */
func (r Rect) FooterPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.FooterPixel(int(float32(r.Height) * percent))
}

/*
* FooterPixel 从底部保留指定像素高度，返回底部矩形
 */
func (r Rect) FooterPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Height {
		pixel = r.Height
	}
	return Rect{
		X:      r.X,
		Y:      r.Y + r.Height - pixel,
		Width:  r.Width,
		Height: pixel,
	}
}

/*
* IsOverlapping 检查两个矩形是否重叠
 */
func (r Rect) IsOverlapping(another Rect) bool {
	if r.X > another.X+another.Width || another.X > r.X+r.Width {
		return false
	}
	if r.Y > another.Y+another.Height || another.Y > r.Y+r.Height {
		return false
	}
	return true
}

/*
* IsEmpty 检查矩形是否为空（宽度或高度为 0 或负数）
 */
func (r Rect) IsEmpty() bool {
	return r.Width <= 0 || r.Height <= 0
}

/*
* Merge 将两个矩形合并为包含两者的最小矩形
 */
func (r Rect) Merge(another Rect) Rect {
	if r.IsEmpty() {
		return another
	}
	if another.IsEmpty() {
		return r
	}
	return Rect{
		X:      min(r.X, another.X),
		Y:      min(r.Y, another.Y),
		Width:  max(r.X+r.Width, another.X+another.Width) - min(r.X, another.X),
		Height: max(r.Y+r.Height, another.Y+another.Height) - min(r.Y, another.Y),
	}
}

/*
* Intersect 计算两个矩形的交集
 */
func (r Rect) Intersect(another Rect) Rect {
	if r.IsEmpty() || another.IsEmpty() {
		return Rect{}
	}
	return Rect{
		X:      max(r.X, another.X),
		Y:      max(r.Y, another.Y),
		Width:  min(r.X+r.Width, another.X+another.Width) - max(r.X, another.X),
		Height: min(r.Y+r.Height, another.Y+another.Height) - max(r.Y, another.Y),
	}
}

/*
* CenterPoint 返回矩形的中心点
 */
func (r Rect) CenterPoint() Point {
	return Point{
		X: r.X + r.Width/2,
		Y: r.Y + r.Height/2,
	}
}

/*
* Contains 检查点 (x, y) 是否在矩形内部
 */
func (r Rect) Contains(x, y int) bool {
	return x >= r.X && x < r.X+r.Width && y >= r.Y && y < r.Y+r.Height
}

/*
* ContainsPoint 检查点 p 是否在矩形内部
 */
func (r Rect) ContainsPoint(p Point) bool {
	return r.Contains(p.X, p.Y)
}

/*
* MergeOverlappingRectangles 将重叠的矩形合并为一个
 */
func MergeOverlappingRectangles(rectangles []Rect) []Rect {
	if len(rectangles) == 0 {
		return nil
	}
	merged := []Rect{rectangles[0]}
	for _, rect := range rectangles[1:] {
		overlapped := false
		for i, mergedRect := range merged {
			if mergedRect.IsOverlapping(rect) {
				merged[i] = mergedRect.Merge(rect)
				overlapped = true
				break
			}
		}
		if !overlapped {
			merged = append(merged, rect)
		}
	}
	return merged
}

/*
* MergeAllRectangles 将所有矩形合并为一个包含所有矩形的最小矩形
 */
func MergeAllRectangles(rectangles []Rect) Rect {
	if len(rectangles) == 0 {
		return Rect{}
	}
	merged := rectangles[0]
	for _, rect := range rectangles[1:] {
		merged = merged.Merge(rect)
	}
	return merged
}

/*
* MinAreaRect 返回面积最小的矩形
 */
func MinAreaRect(rectangles []Rect) Rect {
	if len(rectangles) == 0 {
		return Rect{}
	}
	minIdx := 0
	minArea := rectangles[0].Size().Area()
	for i := 1; i < len(rectangles); i++ {
		area := rectangles[i].Size().Area()
		if area < minArea {
			minArea = area
			minIdx = i
		}
	}
	return rectangles[minIdx]
}

/*
* MaxAreaRect 返回面积最大的矩形
 */
func MaxAreaRect(rectangles []Rect) Rect {
	if len(rectangles) == 0 {
		return Rect{}
	}
	maxIdx := 0
	maxArea := rectangles[0].Size().Area()
	for i := 1; i < len(rectangles); i++ {
		area := rectangles[i].Size().Area()
		if area > maxArea {
			maxArea = area
			maxIdx = i
		}
	}
	return rectangles[maxIdx]
}

/*
* MergeGroupRectangles 将多个矩形组中的矩形进行各种组合方式合并
* 返回所有可能的合并结果，按面积从小到大排序
 */
func MergeGroupRectangles(groupRects ...[]Rect) []Rect {
	if len(groupRects) == 0 {
		return nil
	}
	var results []Rect
	var dfs func(idx int, picked []Rect)
	dfs = func(idx int, picked []Rect) {
		if idx == len(groupRects) {
			merged := MergeAllRectangles(picked)
			results = append(results, merged)
			return
		}
		for _, r := range groupRects[idx] {
			dfs(idx+1, append(picked, r))
		}
	}
	dfs(0, nil)

	/*
	* 用 map 去重
	 */
	uniqueMap := make(map[Rect]struct{})
	uniqueResults := make([]Rect, 0, len(results))
	for _, r := range results {
		if _, exists := uniqueMap[r]; !exists {
			uniqueMap[r] = struct{}{}
			uniqueResults = append(uniqueResults, r)
		}
	}

	sort.Slice(uniqueResults, func(i, j int) bool {
		return uniqueResults[i].Compare(uniqueResults[j]) < 0
	})

	return uniqueResults
}
