package appgui

/*
* AppKey represents a hardware or system key on a mobile device.
 */
type AppKey string

const (
	// Navigation keys
	KeyBack   AppKey = "back"
	KeyHome   AppKey = "home"
	KeyMenu   AppKey = "menu"
	KeyRecent AppKey = "recent"

	// Volume keys
	KeyVolumeUp   AppKey = "volume_up"
	KeyVolumeDown AppKey = "volume_down"
	KeyVolumeMute AppKey = "volume_mute"

	// Power key
	KeyPower AppKey = "power"

	// Soft keyboard keys
	KeyEnter     AppKey = "enter"
	KeyDelete    AppKey = "delete"
	KeyTab       AppKey = "tab"
	KeySearch    AppKey = "search"
	KeyDone      AppKey = "done"
	KeyNext      AppKey = "next"
	KeyPrevious  AppKey = "previous"
	KeySpace     AppKey = "space"
	KeyBackspace AppKey = "backspace"
)
