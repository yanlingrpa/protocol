package osgui

type Keyboard string

const (
	// Letter keys
	KeyA Keyboard = "a"
	KeyB Keyboard = "b"
	KeyC Keyboard = "c"
	KeyD Keyboard = "d"
	KeyE Keyboard = "e"
	KeyF Keyboard = "f"
	KeyG Keyboard = "g"
	KeyH Keyboard = "h"
	KeyI Keyboard = "i"
	KeyJ Keyboard = "j"
	KeyK Keyboard = "k"
	KeyL Keyboard = "l"
	KeyM Keyboard = "m"
	KeyN Keyboard = "n"
	KeyO Keyboard = "o"
	KeyP Keyboard = "p"
	KeyQ Keyboard = "q"
	KeyR Keyboard = "r"
	KeyS Keyboard = "s"
	KeyT Keyboard = "t"
	KeyU Keyboard = "u"
	KeyV Keyboard = "v"
	KeyW Keyboard = "w"
	KeyX Keyboard = "x"
	KeyY Keyboard = "y"
	KeyZ Keyboard = "z"

	CapA Keyboard = "A"
	CapB Keyboard = "B"
	CapC Keyboard = "C"
	CapD Keyboard = "D"
	CapE Keyboard = "E"
	CapF Keyboard = "F"
	CapG Keyboard = "G"
	CapH Keyboard = "H"
	CapI Keyboard = "I"
	CapJ Keyboard = "J"
	CapK Keyboard = "K"
	CapL Keyboard = "L"
	CapM Keyboard = "M"
	CapN Keyboard = "N"
	CapO Keyboard = "O"
	CapP Keyboard = "P"
	CapQ Keyboard = "Q"
	CapR Keyboard = "R"
	CapS Keyboard = "S"
	CapT Keyboard = "T"
	CapU Keyboard = "U"
	CapV Keyboard = "V"
	CapW Keyboard = "W"
	CapX Keyboard = "X"
	CapY Keyboard = "Y"
	CapZ Keyboard = "Z"

	// Number keys
	Key0 Keyboard = "0"
	Key1 Keyboard = "1"
	Key2 Keyboard = "2"
	Key3 Keyboard = "3"
	Key4 Keyboard = "4"
	Key5 Keyboard = "5"
	Key6 Keyboard = "6"
	Key7 Keyboard = "7"
	Key8 Keyboard = "8"
	Key9 Keyboard = "9"

	// Backspace backspace key string
	Backspace Keyboard = "backspace"
	Delete    Keyboard = "delete"
	Enter     Keyboard = "enter"
	Tab       Keyboard = "tab"
	Esc       Keyboard = "esc"
	Escape    Keyboard = "escape"
	Up        Keyboard = "up"    // Up arrow key
	Down      Keyboard = "down"  // Down arrow key
	Right     Keyboard = "right" // Right arrow key
	Left      Keyboard = "left"  // Left arrow key
	Home      Keyboard = "home"
	End       Keyboard = "end"
	Pageup    Keyboard = "pageup"
	Pagedown  Keyboard = "pagedown"

	F1  Keyboard = "f1"
	F2  Keyboard = "f2"
	F3  Keyboard = "f3"
	F4  Keyboard = "f4"
	F5  Keyboard = "f5"
	F6  Keyboard = "f6"
	F7  Keyboard = "f7"
	F8  Keyboard = "f8"
	F9  Keyboard = "f9"
	F10 Keyboard = "f10"
	F11 Keyboard = "f11"
	F12 Keyboard = "f12"
	F13 Keyboard = "f13"
	F14 Keyboard = "f14"
	F15 Keyboard = "f15"
	F16 Keyboard = "f16"
	F17 Keyboard = "f17"
	F18 Keyboard = "f18"
	F19 Keyboard = "f19"
	F20 Keyboard = "f20"
	F21 Keyboard = "f21"
	F22 Keyboard = "f22"
	F23 Keyboard = "f23"
	F24 Keyboard = "f24"

	Cmd  Keyboard = "cmd"  // is the "win" key for windows
	Lcmd Keyboard = "lcmd" // left command
	Rcmd Keyboard = "rcmd" // right command
	// "command"
	Alt     Keyboard = "alt"
	Lalt    Keyboard = "lalt" // left alt
	Ralt    Keyboard = "ralt" // right alt
	Ctrl    Keyboard = "ctrl"
	Lctrl   Keyboard = "lctrl" // left ctrl
	Rctrl   Keyboard = "rctrl" // right ctrl
	Control Keyboard = "control"
	Shift   Keyboard = "shift"
	Lshift  Keyboard = "lshift" // left shift
	Rshift  Keyboard = "rshift" // right shift
	// "right_shift"
	Capslock    Keyboard = "capslock"
	Space       Keyboard = "space"
	Print       Keyboard = "print"
	Printscreen Keyboard = "printscreen" // No Mac support
	Insert      Keyboard = "insert"
	Menu        Keyboard = "menu" // Windows only

	AudioMute    Keyboard = "audio_mute"     // Mute the volume
	AudioVolDown Keyboard = "audio_vol_down" // Lower the volume
	AudioVolUp   Keyboard = "audio_vol_up"   // Increase the volume
	AudioPlay    Keyboard = "audio_play"
	AudioStop    Keyboard = "audio_stop"
	AudioPause   Keyboard = "audio_pause"
	AudioPrev    Keyboard = "audio_prev"    // Previous Track
	AudioNext    Keyboard = "audio_next"    // Next Track
	AudioRewind  Keyboard = "audio_rewind"  // Linux only
	AudioForward Keyboard = "audio_forward" // Linux only
	AudioRepeat  Keyboard = "audio_repeat"  //  Linux only
	AudioRandom  Keyboard = "audio_random"  //  Linux only

	Num0    Keyboard = "num0" // numpad 0
	Num1    Keyboard = "num1"
	Num2    Keyboard = "num2"
	Num3    Keyboard = "num3"
	Num4    Keyboard = "num4"
	Num5    Keyboard = "num5"
	Num6    Keyboard = "num6"
	Num7    Keyboard = "num7"
	Num8    Keyboard = "num9"
	NumLock Keyboard = "num_lock"

	NumDecimal Keyboard = "num."
	NumPlus    Keyboard = "num+"
	NumMinus   Keyboard = "num-"
	NumMul     Keyboard = "num*"
	NumDiv     Keyboard = "num/"
	NumClear   Keyboard = "num_clear"
	NumEnter   Keyboard = "num_enter"
	NumEqual   Keyboard = "num_equal"

	LightsMonUp     Keyboard = "lights_mon_up"     // Turn up monitor brightness			No Windows support
	LightsMonDown   Keyboard = "lights_mon_down"   // Turn down monitor brightness		No Windows support
	LightsKbdToggle Keyboard = "lights_kbd_toggle" // Toggle Keyboard backlight on/off		No Windows support
	LightsKbdUp     Keyboard = "lights_kbd_up"     // Turn up Keyboard backlight brightness	No Windows support
	LightsKbdDown   Keyboard = "lights_kbd_down"
)

func (k Keyboard) String() string {
	return string(k)
}
