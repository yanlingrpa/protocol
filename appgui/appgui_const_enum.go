package appgui

/*
* AppKey represents a hardware or system key on a mobile device.
 */
type AppKey string

const (
	// Navigation keys
	// 返回键。
	KeyBack AppKey = "back"
	// 主页键。
	KeyHome AppKey = "home"
	// 菜单键。
	KeyMenu AppKey = "menu"
	// 最近任务键。
	KeyRecent AppKey = "recent"
	// Esc 退出键。
	KeyEscape AppKey = "escape"
	// 接听电话键。
	KeyCall AppKey = "call"
	// 挂断电话键。
	KeyEndCall AppKey = "end_call"
	// 相机键。
	KeyCamera AppKey = "camera"
	// 清除键。
	KeyClear AppKey = "clear"

	// Direction and navigation keys
	// 方向上键。
	KeyDpadUp AppKey = "dpad_up"
	// 方向下键。
	KeyDpadDown AppKey = "dpad_down"
	// 方向左键。
	KeyDpadLeft AppKey = "dpad_left"
	// 方向右键。
	KeyDpadRight AppKey = "dpad_right"
	// 方向确认键。
	KeyDpadCenter AppKey = "dpad_center"
	// 向上翻页键。
	KeyPageUp AppKey = "page_up"
	// 向下翻页键。
	KeyPageDown AppKey = "page_down"
	// 文本移动到开头键。
	KeyMoveHome AppKey = "move_home"
	// 文本移动到结尾键。
	KeyMoveEnd AppKey = "move_end"
	// 插入键。
	KeyInsert AppKey = "insert"

	// Volume keys
	// 音量增加键。
	KeyVolumeUp AppKey = "volume_up"
	// 音量降低键。
	KeyVolumeDown AppKey = "volume_down"
	// 静音键。
	KeyVolumeMute AppKey = "volume_mute"

	// Power key
	// 电源键。
	KeyPower AppKey = "power"

	// Soft keyboard keys
	// 回车键。
	KeyEnter AppKey = "enter"
	// 删除键。
	KeyDelete AppKey = "delete"
	// 制表键。
	KeyTab AppKey = "tab"
	// 搜索键。
	KeySearch AppKey = "search"
	// 输入法完成键，映射为回车键。
	KeyDone AppKey = "done"
	// 输入法下一项键，映射为制表键。
	KeyNext AppKey = "next"
	// 输入法上一项键，没有对应的单一 Android keycode。
	KeyPrevious AppKey = "previous"
	// 空格键。
	KeySpace AppKey = "space"
	// 退格键。
	KeyBackspace AppKey = "backspace"

	// Modifier keys
	// Shift 修饰键。
	KeyShift AppKey = "shift"
	// Ctrl 修饰键。
	KeyCtrl AppKey = "ctrl"
	// Alt 修饰键。
	KeyAlt AppKey = "alt"
	// Meta 修饰键。
	KeyMeta AppKey = "meta"
	// 大小写锁定键。
	KeyCapsLock AppKey = "caps_lock"
	// 逗号键。
	KeyComma AppKey = "comma"
	// 句号键。
	KeyPeriod AppKey = "period"
	// 耳机接听键。
	KeyHeadsetHook AppKey = "headset_hook"
	// 通知键。
	KeyNotification AppKey = "notification"
	// 文件浏览器键。
	KeyExplorer AppKey = "explorer"
	// 邮件键。
	KeyEnvelope AppKey = "envelope"
	// 相机对焦键。
	KeyFocus AppKey = "focus"

	// Number keys
	// 数字 0 键。
	Key0 AppKey = "0"
	// 数字 1 键。
	Key1 AppKey = "1"
	// 数字 2 键。
	Key2 AppKey = "2"
	// 数字 3 键。
	Key3 AppKey = "3"
	// 数字 4 键。
	Key4 AppKey = "4"
	// 数字 5 键。
	Key5 AppKey = "5"
	// 数字 6 键。
	Key6 AppKey = "6"
	// 数字 7 键。
	Key7 AppKey = "7"
	// 数字 8 键。
	Key8 AppKey = "8"
	// 数字 9 键。
	Key9 AppKey = "9"

	// Letter keys
	// 字母 A 键。
	KeyA AppKey = "a"
	// 字母 B 键。
	KeyB AppKey = "b"
	// 字母 C 键。
	KeyC AppKey = "c"
	// 字母 D 键。
	KeyD AppKey = "d"
	// 字母 E 键。
	KeyE AppKey = "e"
	// 字母 F 键。
	KeyF AppKey = "f"
	// 字母 G 键。
	KeyG AppKey = "g"
	// 字母 H 键。
	KeyH AppKey = "h"
	// 字母 I 键。
	KeyI AppKey = "i"
	// 字母 J 键。
	KeyJ AppKey = "j"
	// 字母 K 键。
	KeyK AppKey = "k"
	// 字母 L 键。
	KeyL AppKey = "l"
	// 字母 M 键。
	KeyM AppKey = "m"
	// 字母 N 键。
	KeyN AppKey = "n"
	// 字母 O 键。
	KeyO AppKey = "o"
	// 字母 P 键。
	KeyP AppKey = "p"
	// 字母 Q 键。
	KeyQ AppKey = "q"
	// 字母 R 键。
	KeyR AppKey = "r"
	// 字母 S 键。
	KeyS AppKey = "s"
	// 字母 T 键。
	KeyT AppKey = "t"
	// 字母 U 键。
	KeyU AppKey = "u"
	// 字母 V 键。
	KeyV AppKey = "v"
	// 字母 W 键。
	KeyW AppKey = "w"
	// 字母 X 键。
	KeyX AppKey = "x"
	// 字母 Y 键。
	KeyY AppKey = "y"
	// 字母 Z 键。
	KeyZ AppKey = "z"

	// Function keys
	// 功能键 F1。
	KeyF1 AppKey = "f1"
	// 功能键 F2。
	KeyF2 AppKey = "f2"
	// 功能键 F3。
	KeyF3 AppKey = "f3"
	// 功能键 F4。
	KeyF4 AppKey = "f4"
	// 功能键 F5。
	KeyF5 AppKey = "f5"
	// 功能键 F6。
	KeyF6 AppKey = "f6"
	// 功能键 F7。
	KeyF7 AppKey = "f7"
	// 功能键 F8。
	KeyF8 AppKey = "f8"
	// 功能键 F9。
	KeyF9 AppKey = "f9"
	// 功能键 F10。
	KeyF10 AppKey = "f10"
	// 功能键 F11。
	KeyF11 AppKey = "f11"
	// 功能键 F12。
	KeyF12 AppKey = "f12"

	// Media keys
	// 媒体播放/暂停键。
	KeyMediaPlayPause AppKey = "media_play_pause"
	// 媒体停止键。
	KeyMediaStop AppKey = "media_stop"
	// 媒体下一曲键。
	KeyMediaNext AppKey = "media_next"
	// 媒体上一曲键。
	KeyMediaPrevious AppKey = "media_previous"
	// 媒体快退键。
	KeyMediaRewind AppKey = "media_rewind"
	// 媒体快进键。
	KeyMediaFastForward AppKey = "media_fast_forward"
	// 媒体静音键。
	KeyMediaMute AppKey = "media_mute"
	// 媒体录音键。
	KeyMediaRecord AppKey = "media_record"
)

func (k AppKey) AndroidKeyCode() int {
	switch k {
	case KeyBack:
		return 4
	case KeyHome:
		return 3
	case KeyMenu:
		return 82
	case KeyRecent:
		return 187
	case KeyEscape:
		return 111
	case KeyCall:
		return 5
	case KeyEndCall:
		return 6
	case KeyCamera:
		return 27
	case KeyClear:
		return 28
	case KeyDpadUp:
		return 19
	case KeyDpadDown:
		return 20
	case KeyDpadLeft:
		return 21
	case KeyDpadRight:
		return 22
	case KeyDpadCenter:
		return 23
	case KeyPageUp:
		return 92
	case KeyPageDown:
		return 93
	case KeyMoveHome:
		return 122
	case KeyMoveEnd:
		return 123
	case KeyInsert:
		return 124
	case KeyVolumeUp:
		return 24
	case KeyVolumeDown:
		return 25
	case KeyVolumeMute:
		return 164
	case KeyPower:
		return 26
	case KeyEnter:
		return 66
	case KeyDelete:
		return 67
	case KeyTab:
		return 61
	case KeySearch:
		return 84
	case KeyDone:
		return 66
	case KeyNext:
		return 61
	case KeyPrevious:
		return 0
	case KeySpace:
		return 62
	case KeyBackspace:
		return 67
	case KeyShift:
		return 59
	case KeyCtrl:
		return 113
	case KeyAlt:
		return 57
	case KeyMeta:
		return 117
	case KeyCapsLock:
		return 115
	case KeyComma:
		return 55
	case KeyPeriod:
		return 56
	case KeyHeadsetHook:
		return 79
	case KeyNotification:
		return 83
	case KeyExplorer:
		return 64
	case KeyEnvelope:
		return 65
	case KeyFocus:
		return 80
	case Key0:
		return 7
	case Key1:
		return 8
	case Key2:
		return 9
	case Key3:
		return 10
	case Key4:
		return 11
	case Key5:
		return 12
	case Key6:
		return 13
	case Key7:
		return 14
	case Key8:
		return 15
	case Key9:
		return 16
	case KeyA:
		return 29
	case KeyB:
		return 30
	case KeyC:
		return 31
	case KeyD:
		return 32
	case KeyE:
		return 33
	case KeyF:
		return 34
	case KeyG:
		return 35
	case KeyH:
		return 36
	case KeyI:
		return 37
	case KeyJ:
		return 38
	case KeyK:
		return 39
	case KeyL:
		return 40
	case KeyM:
		return 41
	case KeyN:
		return 42
	case KeyO:
		return 43
	case KeyP:
		return 44
	case KeyQ:
		return 45
	case KeyR:
		return 46
	case KeyS:
		return 47
	case KeyT:
		return 48
	case KeyU:
		return 49
	case KeyV:
		return 50
	case KeyW:
		return 51
	case KeyX:
		return 52
	case KeyY:
		return 53
	case KeyZ:
		return 54
	case KeyF1:
		return 131
	case KeyF2:
		return 132
	case KeyF3:
		return 133
	case KeyF4:
		return 134
	case KeyF5:
		return 135
	case KeyF6:
		return 136
	case KeyF7:
		return 137
	case KeyF8:
		return 138
	case KeyF9:
		return 139
	case KeyF10:
		return 140
	case KeyF11:
		return 141
	case KeyF12:
		return 142
	case KeyMediaPlayPause:
		return 85
	case KeyMediaStop:
		return 86
	case KeyMediaNext:
		return 87
	case KeyMediaPrevious:
		return 88
	case KeyMediaRewind:
		return 89
	case KeyMediaFastForward:
		return 90
	case KeyMediaMute:
		return 91
	case KeyMediaRecord:
		return 130
	default:
		return 0
	}
}
