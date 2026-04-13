package basic

import "sort"

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (r Rect) Position() Point {
	return Point{
		X: r.X,
		Y: r.Y,
	}
}

func (r Rect) Size() Size {
	return Size{
		Width:  r.Width,
		Height: r.Height,
	}
}

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

func (r Rect) LeftPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.LeftPixel(int(float32(r.Width) * percent))
}

func (r Rect) ExLeftPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.ExLeftPixel(int(float32(r.Width) * percent))
}

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

func (r Rect) ExLeftPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Width {
		pixel = r.Width
	}
	return Rect{
		X:      r.X + pixel,
		Y:      r.Y,
		Width:  r.Width - pixel,
		Height: r.Height,
	}
}

func (r Rect) RightPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.RightPixel(int(float32(r.Width) * percent))
}

func (r Rect) ExRightPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.ExRightPixel(int(float32(r.Width) * percent))
}

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

func (r Rect) ExRightPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Width {
		pixel = r.Width
	}
	return Rect{
		X:      r.X,
		Y:      r.Y,
		Width:  r.Width - pixel,
		Height: r.Height,
	}
}

func (r Rect) CenterPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.CenterPixel(int(float32(r.Width) * percent))
}

func (r Rect) CenterPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Width {
		pixel = r.Width
	}
	return Rect{
		X:      r.X + (r.Width-pixel)/2,
		Y:      r.Y,
		Width:  pixel,
		Height: r.Height,
	}
}

func (r Rect) HeaderPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.HeaderPixel(int(float32(r.Height) * percent))
}

func (r Rect) ExHeaderPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.ExHeaderPixel(int(float32(r.Height) * percent))
}

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

func (r Rect) ExHeaderPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Height {
		pixel = r.Height
	}
	return Rect{
		X:      r.X,
		Y:      r.Y + pixel,
		Width:  r.Width,
		Height: r.Height - pixel,
	}
}

func (r Rect) FooterPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.FooterPixel(int(float32(r.Height) * percent))
}

func (r Rect) ExFooterPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.ExFooterPixel(int(float32(r.Height) * percent))
}

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

func (r Rect) ExFooterPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Height {
		pixel = r.Height
	}
	return Rect{
		X:      r.X,
		Y:      r.Y,
		Width:  r.Width,
		Height: r.Height - pixel,
	}
}

func (r Rect) MainPercent(percent float32) Rect {
	if percent < 0 {
		percent = 0
	} else if percent > 1 {
		percent = 1
	}
	return r.MainPixel(int(float32(r.Height) * percent))
}

func (r Rect) MainPixel(pixel int) Rect {
	if pixel < 0 {
		pixel = 0
	} else if pixel > r.Height {
		pixel = r.Height
	}
	return Rect{
		X:      r.X,
		Y:      r.Y + (r.Height-pixel)/2,
		Width:  r.Width,
		Height: pixel,
	}
}

func (r Rect) IsOverlapping(another Rect) bool {
	if r.X > another.X+another.Width || another.X > r.X+r.Width {
		return false
	}
	if r.Y > another.Y+another.Height || another.Y > r.Y+r.Height {
		return false
	}
	return true
}

func (r Rect) IsEmpty() bool {
	return r.Width <= 0 || r.Height <= 0
}

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

func (r Rect) Center() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

func (r Rect) CenterPoint() Point {
	return Point{
		X: r.X + r.Width/2,
		Y: r.Y + r.Height/2,
	}
}

func (r Rect) Contains(x, y int) bool {
	return x >= r.X && x < r.X+r.Width && y >= r.Y && y < r.Y+r.Height
}

func (r Rect) ContainsPoint(p Point) bool {
	return r.Contains(p.X, p.Y)
}

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

	// 用 map 去重
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
