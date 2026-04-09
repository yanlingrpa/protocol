package browser

import (
	"time"

	"github.com/yanlingrpa/protocol/basic"
)

type BrowserElement interface {
	// 让当前元素获得输入焦点，即让浏览器把光标定位到这个元素上，通常用于输入框、按钮等可交互元素。
	Focus()
	// 将当前元素滚动到浏览器窗口的可视区域内，如果它本来不在可视区域，就自动滚动页面让它显示出来
	ScrollIntoView() error
	Hover() error
	MoveMouseOut() error
	Click() error
	RightClick() error
	DoubleClick() error
	Tap() error
	Interactable() (basic.FPoint, error)
	SelectText(regex string) error
	SelectAllText() error
	Input(text string) error
	Blur() error
	SelectByText(texts []string, selected bool) error
	SelectByRegex(regexes []string, selected bool) error
	SelectByCss(selectors []string, selected bool) error
	MatchByCss(selector string) (bool, error)
	GetAttribute(name string) (string, error)
	SetAttribute(name, value string) error
	GetProperty(name string) (any, error)
	SetProperty(name string, value any) error
	Disabled() (bool, error)
	SetFiles(filePaths []string) error
	FramePage() (BrowserFramePage, error)
	ContainsElement(target BrowserElement) (bool, error)
	Text() (string, error)
	Html() (string, error)
	Visible() (bool, error)
	WaitStable(stableTime time.Duration) error
	WaitStableRAF() error
	WaitInteractable() (basic.FPoint, error)
	WaitVisible() error
	WaitEnabled() error
	WaitWritable() error
	WaitInvisible() error
	Evalute(jsCode string, arg ...any) (any, error)
	GetXPath(optimized bool) (string, error)
}
