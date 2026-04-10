package osgui

import (
	"fmt"
	"strconv"
	"strings"

	"yanlingrpa.com/yanling/protocol/basic"
)

// type GuiWindow struct {
// 	key            string     // 标题关键词
// 	hwnd           uintptr    // 窗口句柄
// 	title          string     // 窗口标题
// 	processId      uint32     // 进程ID
// 	processName    string     // 进程名称
// 	executablePath string     // 可执行文件路径
// 	threadId       uint32     // 线程ID
// 	dpi            uint32     // DPI
// 	screenBounds   basic.Rect // 窗口在屏幕上的位置和大小
// }

// type Color struct {
// 	R uint8 `json:"r"`
// 	G uint8 `json:"g"`
// 	B uint8 `json:"b"`
// 	A uint8 `json:"a"`
// }

// func HexColor(hex string) Color {
// 	// 支持 #RRGGBB, #RRGGBBAA, 0xRRGGBB, 0xRRGGBBAA 四种写法
// 	var r, g, b, a uint8 = 0, 0, 0, 255
// 	var s string

// 	// 去除前缀
// 	if len(hex) >= 2 && (hex[:2] == "0x" || hex[:2] == "0X") {
// 		s = hex[2:]
// 	} else if len(hex) >= 1 && hex[0] == '#' {
// 		s = hex[1:]
// 	} else {
// 		s = hex
// 	}

// 	switch len(s) {
// 	case 6:
// 		// RRGGBB
// 		fmt.Sscanf(s, "%02x%02x%02x", &r, &g, &b)
// 	case 8:
// 		// RRGGBBAA
// 		fmt.Sscanf(s, "%02x%02x%02x%02x", &r, &g, &b, &a)
// 	default:
// 		// 不支持的格式，返回全0
// 		return Color{R: 0, G: 0, B: 0, A: 255}
// 	}
// 	return Color{R: r, G: g, B: b, A: a}
// }

func _hex_color_to_gray_color_value(hex string) (uint8, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, fmt.Errorf("invalid hex color")
	}
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
	return gray, nil
}

type TextField struct {
	Text            string `json:"content"`
	FontSize        int    `json:"font_size"`
	FontColor       string `json:"font_color"`
	BackgroundColor uint8  `json:"background_color"`
	FontBold        bool   `json:"font_bold"`
	FontFamily      string `json:"font_family"`
}

func NewTextField(text string, fontSize int, fontColor string, backgroundColor uint8, fontBold bool) TextField {
	return TextField{
		Text:            text,
		FontSize:        fontSize,
		FontColor:       fontColor,
		BackgroundColor: backgroundColor,
		FontBold:        fontBold,
		FontFamily:      "",
	}
}

type CardField struct {
	OuterColor string     `json:"outer_color"`
	InnerColor string     `json:"inner_color"`
	MinSize    basic.Size `json:"min_size"`
	MaxSize    basic.Size `json:"max_size"`
}

func (s CardField) OuterGrayVaue() uint8 {
	gray, err := _hex_color_to_gray_color_value(s.OuterColor)
	if err != nil {
		return 0
	}
	return gray
}

func (s CardField) InnerGrayValue() uint8 {
	gray, err := _hex_color_to_gray_color_value(s.InnerColor)
	if err != nil {
		return 0
	}
	return gray
}

func (s CardField) Threshold() float32 {
	outerGray := s.OuterGrayVaue()
	innerGray := s.InnerGrayValue()
	if outerGray > innerGray {
		return min(float32((outerGray+innerGray)/2), float32(outerGray-1))
	} else if outerGray < innerGray {
		return min(float32((innerGray+outerGray)/2), float32(innerGray-1))
	} else {
		return 0
	}
}

type ShapeField struct {
	ShapeType GraphicShape `json:"shape_type"`
	MinDist   float64      `json:"min_dist"`
	MinSize   basic.Size   `json:"min_size"`
	MaxSize   basic.Size   `json:"max_size"`
}
