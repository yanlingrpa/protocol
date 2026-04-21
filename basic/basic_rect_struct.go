package basic

import "sort"

/*
* Rect represents a rectangular region, defined by the top-left coordinate and width/height.
 */
type Rect struct {
	/*
	* X is the X coordinate of the rectangle's top-left corner.
	 */
	X int `json:"x"`
	/*
	* Y is the Y coordinate of the rectangle's top-left corner.
	 */
	Y int `json:"y"`
	/*
	* Width represents the rectangle width.
	 */
	Width int `json:"width"`
	/*
	* Height represents the rectangle height.
	 */
	Height int `json:"height"`
}

/*
* Position returns the top-left position of the rectangle.
 */
func (r Rect) Position() Point {
	return Point{
		X: r.X,
		Y: r.Y,
	}
}

/*
* Size returns the width and height of the rectangle.
 */
func (r Rect) Size() Size {
	return Size{
		Width:  r.Width,
		Height: r.Height,
	}
}

/*
* Compare compares two rectangles and returns -1 for less than, 0 for equal, and 1 for greater than.
* Comparison priority: area > X coordinate > Y coordinate > width > height.
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
* SubRect extracts a sub-rectangle from within the rectangle using relative position and size.
* It automatically checks boundaries to prevent out-of-range values.
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
* LeftPercent keeps a given percentage of width from the left side and returns that rectangle.
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
* LeftPixel keeps a given pixel width from the left side and returns that rectangle.
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
* RightPercent keeps a given percentage of width from the right side and returns that rectangle.
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
* RightPixel keeps a given pixel width from the right side and returns that rectangle.
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
* HeaderPercent keeps a given percentage of height from the top and returns that rectangle.
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
* HeaderPixel keeps a given pixel height from the top and returns that rectangle.
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
* FooterPercent keeps a given percentage of height from the bottom and returns that rectangle.
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
* FooterPixel keeps a given pixel height from the bottom and returns that rectangle.
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
* IsOverlapping checks whether two rectangles overlap.
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
* IsEmpty checks whether the rectangle is empty (width or height is 0 or negative).
 */
func (r Rect) IsEmpty() bool {
	return r.Width <= 0 || r.Height <= 0
}

/*
* Merge combines two rectangles into the smallest rectangle containing both.
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
* Intersect computes the intersection of two rectangles.
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
* CenterPoint returns the center point of the rectangle.
 */
func (r Rect) CenterPoint() Point {
	return Point{
		X: r.X + r.Width/2,
		Y: r.Y + r.Height/2,
	}
}

/*
* Contains checks whether point (x, y) lies inside the rectangle.
 */
func (r Rect) Contains(x, y int) bool {
	return x >= r.X && x < r.X+r.Width && y >= r.Y && y < r.Y+r.Height
}

/*
* ContainsPoint checks whether point p lies inside the rectangle.
 */
func (r Rect) ContainsPoint(p Point) bool {
	return r.Contains(p.X, p.Y)
}

/*
* MergeOverlappingRectangles merges overlapping rectangles.
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
* MergeAllRectangles merges all rectangles into the smallest rectangle containing all of them.
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
* MinAreaRect returns the rectangle with the smallest area.
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
* MaxAreaRect returns the rectangle with the largest area.
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
* MergeGroupRectangles merges rectangles from multiple groups in all possible combinations.
* It returns all possible merged results, sorted by area from small to large.
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
	* Deduplicate using a map.
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
