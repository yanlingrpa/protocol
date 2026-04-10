package symbols

import (
	"reflect"

	"yanlingrpa.com/yanling/protocol/basic"
	"yanlingrpa.com/yanling/protocol/browser"
	"yanlingrpa.com/yanling/protocol/extention"
	"yanlingrpa.com/yanling/protocol/osgui"
	"yanlingrpa.com/yanling/protocol/ossys"
	"yanlingrpa.com/yanling/protocol/script"
)

var Symbols = make(map[string]map[string]reflect.Value)

func init() {
	Symbols["yanlingrpa.com/yanling/protocol/basic"] = map[string]reflect.Value{
		"DispatchTaskData":           reflect.ValueOf((*basic.DispatchTaskData)(nil)).Elem(), // Export DispatchTaskData struct
		"FPoint":                     reflect.ValueOf((*basic.FPoint)(nil)).Elem(),           // Export FPoint struct
		"MaxAreaRect":                reflect.ValueOf(basic.MaxAreaRect),                     // Export MaxAreaRect function
		"MergeAllRectangles":         reflect.ValueOf(basic.MergeAllRectangles),              // Export MergeAllRectangles function
		"MergeGroupRectangles":       reflect.ValueOf(basic.MergeGroupRectangles),            // Export MergeGroupRectangles function
		"MergeOverlappingRectangles": reflect.ValueOf(basic.MergeOverlappingRectangles),      // Export MergeOverlappingRectangles function
		"MinAreaRect":                reflect.ValueOf(basic.MinAreaRect),                     // Export MinAreaRect function
		"OcrResult":                  reflect.ValueOf((*basic.OcrResult)(nil)).Elem(),        // Export OcrResult struct
		"OcrText":                    reflect.ValueOf((*basic.OcrText)(nil)).Elem(),          // Export OcrText struct
		"Point":                      reflect.ValueOf((*basic.Point)(nil)).Elem(),            // Export Point struct
		"Rect":                       reflect.ValueOf((*basic.Rect)(nil)).Elem(),             // Export Rect struct
		"Rect.Center":                reflect.ValueOf(basic.Rect.Center),
		"Rect.CenterPercent":         reflect.ValueOf(basic.Rect.CenterPercent),
		"Rect.CenterPixel":           reflect.ValueOf(basic.Rect.CenterPixel),
		"Rect.Compare":               reflect.ValueOf(basic.Rect.Compare),
		"Rect.Contains":              reflect.ValueOf(basic.Rect.Contains),
		"Rect.ExFooterPercent":       reflect.ValueOf(basic.Rect.ExFooterPercent),
		"Rect.ExFooterPixel":         reflect.ValueOf(basic.Rect.ExFooterPixel),
		"Rect.ExHeaderPercent":       reflect.ValueOf(basic.Rect.ExHeaderPercent),
		"Rect.ExHeaderPixel":         reflect.ValueOf(basic.Rect.ExHeaderPixel),
		"Rect.ExLeftPercent":         reflect.ValueOf(basic.Rect.ExLeftPercent),
		"Rect.ExLeftPixel":           reflect.ValueOf(basic.Rect.ExLeftPixel),
		"Rect.ExRightPercent":        reflect.ValueOf(basic.Rect.ExRightPercent),
		"Rect.ExRightPixel":          reflect.ValueOf(basic.Rect.ExRightPixel),
		"Rect.FooterPercent":         reflect.ValueOf(basic.Rect.FooterPercent),
		"Rect.FooterPixel":           reflect.ValueOf(basic.Rect.FooterPixel),
		"Rect.HeaderPercent":         reflect.ValueOf(basic.Rect.HeaderPercent),
		"Rect.HeaderPixel":           reflect.ValueOf(basic.Rect.HeaderPixel),
		"Rect.Intersect":             reflect.ValueOf(basic.Rect.Intersect),
		"Rect.IsEmpty":               reflect.ValueOf(basic.Rect.IsEmpty),
		"Rect.IsOverlapping":         reflect.ValueOf(basic.Rect.IsOverlapping),
		"Rect.LeftPercent":           reflect.ValueOf(basic.Rect.LeftPercent),
		"Rect.LeftPixel":             reflect.ValueOf(basic.Rect.LeftPixel),
		"Rect.MainPercent":           reflect.ValueOf(basic.Rect.MainPercent),
		"Rect.MainPixel":             reflect.ValueOf(basic.Rect.MainPixel),
		"Rect.Merge":                 reflect.ValueOf(basic.Rect.Merge),
		"Rect.Position":              reflect.ValueOf(basic.Rect.Position),
		"Rect.RightPercent":          reflect.ValueOf(basic.Rect.RightPercent),
		"Rect.RightPixel":            reflect.ValueOf(basic.Rect.RightPixel),
		"Rect.Size":                  reflect.ValueOf(basic.Rect.Size),
		"Rect.SubRect":               reflect.ValueOf(basic.Rect.SubRect),
		"Size":                       reflect.ValueOf((*basic.Size)(nil)).Elem(), // Export Size struct
		"Size.Area":                  reflect.ValueOf(basic.Size.Area),
		"Task":                       reflect.ValueOf((*basic.Task)(nil)), // Export Task interface pointer type
		"Task.GetExpiredAt":          reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetExpiredAt"),
		"Task.GetGuid":               reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetGuid"),
		"Task.GetMaxRetries":         reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetMaxRetries"),
		"Task.GetParams":             reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetParams"),
		"Task.GetPayload":            reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetPayload"),
		"Task.GetProjectId":          reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetProjectId"),
		"Task.GetRetryCount":         reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetRetryCount"),
		"Task.GetScriptId":           reflect.ValueOf((*basic.Task)(nil)).MethodByName("GetScriptId"),
	}
	Symbols["yanlingrpa.com/yanling/protocol/browser"] = map[string]reflect.Value{
		"BrowserElement":                      reflect.ValueOf((*browser.BrowserElement)(nil)), // Export BrowserElement interface pointer type
		"BrowserElement.Blur":                 reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Blur"),
		"BrowserElement.Click":                reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Click"),
		"BrowserElement.ContainsElement":      reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("ContainsElement"),
		"BrowserElement.Disabled":             reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Disabled"),
		"BrowserElement.DoubleClick":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("DoubleClick"),
		"BrowserElement.Evalute":              reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Evalute"),
		"BrowserElement.Focus":                reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Focus"),
		"BrowserElement.FramePage":            reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("FramePage"),
		"BrowserElement.GetAttribute":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("GetAttribute"),
		"BrowserElement.GetProperty":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("GetProperty"),
		"BrowserElement.GetXPath":             reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("GetXPath"),
		"BrowserElement.Hover":                reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Hover"),
		"BrowserElement.Html":                 reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Html"),
		"BrowserElement.Input":                reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Input"),
		"BrowserElement.Interactable":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Interactable"),
		"BrowserElement.MatchByCss":           reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("MatchByCss"),
		"BrowserElement.MoveMouseOut":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("MoveMouseOut"),
		"BrowserElement.RightClick":           reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("RightClick"),
		"BrowserElement.ScrollIntoView":       reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("ScrollIntoView"),
		"BrowserElement.SelectAllText":        reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SelectAllText"),
		"BrowserElement.SelectByCss":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SelectByCss"),
		"BrowserElement.SelectByRegex":        reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SelectByRegex"),
		"BrowserElement.SelectByText":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SelectByText"),
		"BrowserElement.SelectText":           reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SelectText"),
		"BrowserElement.SetAttribute":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SetAttribute"),
		"BrowserElement.SetFiles":             reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SetFiles"),
		"BrowserElement.SetProperty":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("SetProperty"),
		"BrowserElement.Tap":                  reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Tap"),
		"BrowserElement.Text":                 reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Text"),
		"BrowserElement.Visible":              reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("Visible"),
		"BrowserElement.WaitEnabled":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitEnabled"),
		"BrowserElement.WaitInteractable":     reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitInteractable"),
		"BrowserElement.WaitInvisible":        reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitInvisible"),
		"BrowserElement.WaitStable":           reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitStable"),
		"BrowserElement.WaitStableRAF":        reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitStableRAF"),
		"BrowserElement.WaitVisible":          reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitVisible"),
		"BrowserElement.WaitWritable":         reflect.ValueOf((*browser.BrowserElement)(nil)).MethodByName("WaitWritable"),
		"BrowserFramePage":                    reflect.ValueOf((*browser.BrowserFramePage)(nil)), // Export BrowserFramePage interface pointer type
		"BrowserFramePage.Evaluate":           reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("Evaluate"),
		"BrowserFramePage.GetDomain":          reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("GetDomain"),
		"BrowserFramePage.GetID":              reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("GetID"),
		"BrowserFramePage.GetTitle":           reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("GetTitle"),
		"BrowserFramePage.GetURL":             reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("GetURL"),
		"BrowserFramePage.IsFrame":            reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("IsFrame"),
		"BrowserFramePage.QuerySelector":      reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("QuerySelector"),
		"BrowserFramePage.QuerySelectorAll":   reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("QuerySelectorAll"),
		"BrowserFramePage.QueryXPath":         reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("QueryXPath"),
		"BrowserFramePage.QueryXPathAll":      reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("QueryXPathAll"),
		"BrowserFramePage.Reload":             reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("Reload"),
		"BrowserFramePage.WaitSelector":       reflect.ValueOf((*browser.BrowserFramePage)(nil)).MethodByName("WaitSelector"),
		"BrowserTabPage":                      reflect.ValueOf((*browser.BrowserTabPage)(nil)), // Export BrowserTabPage interface pointer type
		"BrowserTabPage.Activate":             reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("Activate"),
		"BrowserTabPage.ClearAllIndexDBFiles": reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("ClearAllIndexDBFiles"),
		"BrowserTabPage.ClearCookies":         reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("ClearCookies"),
		"BrowserTabPage.ClearLocalStorage":    reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("ClearLocalStorage"),
		"BrowserTabPage.Destroy":              reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("Destroy"),
		"BrowserTabPage.LoadCookies":          reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("LoadCookies"),
		"BrowserTabPage.LoadLocalStorage":     reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("LoadLocalStorage"),
		"BrowserTabPage.SaveCookies":          reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("SaveCookies"),
		"BrowserTabPage.SaveLocalStorage":     reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("SaveLocalStorage"),
		"BrowserTabPage.WaitForNewTab":        reflect.ValueOf((*browser.BrowserTabPage)(nil)).MethodByName("WaitForNewTab"),
		"BrowserWindow":                       reflect.ValueOf((*browser.BrowserWindow)(nil)), // Export BrowserWindow interface pointer type
		"BrowserWindow.CurrentPage":           reflect.ValueOf((*browser.BrowserWindow)(nil)).MethodByName("CurrentPage"),
		"BrowserWindow.DefaultPage":           reflect.ValueOf((*browser.BrowserWindow)(nil)).MethodByName("DefaultPage"),
		"BrowserWindow.IDTabPage":             reflect.ValueOf((*browser.BrowserWindow)(nil)).MethodByName("IDTabPage"),
		"BrowserWindow.NewTabPage":            reflect.ValueOf((*browser.BrowserWindow)(nil)).MethodByName("NewTabPage"),
	}
	Symbols["yanlingrpa.com/yanling/protocol/extention"] = map[string]reflect.Value{
		"OcrResult":              reflect.ValueOf((*extention.OcrResult)(nil)).Elem(), // Export OcrResult struct
		"OcrWord":                reflect.ValueOf((*extention.OcrWord)(nil)).Elem(),   // Export OcrWord struct
		"VisionExtension":        reflect.ValueOf((*extention.VisionExtension)(nil)),  // Export VisionExtension interface pointer type
		"VisionExtension.Detect": reflect.ValueOf((*extention.VisionExtension)(nil)).MethodByName("Detect"),
		"VisionExtension.Locate": reflect.ValueOf((*extention.VisionExtension)(nil)).MethodByName("Locate"),
		"VisionExtension.Ocr":    reflect.ValueOf((*extention.VisionExtension)(nil)).MethodByName("Ocr"),
		"VisionExtension.Read":   reflect.ValueOf((*extention.VisionExtension)(nil)).MethodByName("Read"),
	}
	Symbols["yanlingrpa.com/yanling/protocol/osgui"] = map[string]reflect.Value{
		"Alt":                                 reflect.ValueOf(osgui.Alt),                      // Export Alt constant
		"AudioForward":                        reflect.ValueOf(osgui.AudioForward),             // Export AudioForward constant
		"AudioMute":                           reflect.ValueOf(osgui.AudioMute),                // Export AudioMute constant
		"AudioNext":                           reflect.ValueOf(osgui.AudioNext),                // Export AudioNext constant
		"AudioPause":                          reflect.ValueOf(osgui.AudioPause),               // Export AudioPause constant
		"AudioPlay":                           reflect.ValueOf(osgui.AudioPlay),                // Export AudioPlay constant
		"AudioPrev":                           reflect.ValueOf(osgui.AudioPrev),                // Export AudioPrev constant
		"AudioRandom":                         reflect.ValueOf(osgui.AudioRandom),              // Export AudioRandom constant
		"AudioRepeat":                         reflect.ValueOf(osgui.AudioRepeat),              // Export AudioRepeat constant
		"AudioRewind":                         reflect.ValueOf(osgui.AudioRewind),              // Export AudioRewind constant
		"AudioStop":                           reflect.ValueOf(osgui.AudioStop),                // Export AudioStop constant
		"AudioVolDown":                        reflect.ValueOf(osgui.AudioVolDown),             // Export AudioVolDown constant
		"AudioVolUp":                          reflect.ValueOf(osgui.AudioVolUp),               // Export AudioVolUp constant
		"Backspace":                           reflect.ValueOf(osgui.Backspace),                // Export Backspace constant
		"Capslock":                            reflect.ValueOf(osgui.Capslock),                 // Export Capslock constant
		"CardField":                           reflect.ValueOf((*osgui.CardField)(nil)).Elem(), // Export CardField struct
		"CardField.InnerGrayValue":            reflect.ValueOf(osgui.CardField.InnerGrayValue),
		"CardField.OuterGrayVaue":             reflect.ValueOf(osgui.CardField.OuterGrayVaue),
		"CardField.Threshold":                 reflect.ValueOf(osgui.CardField.Threshold),
		"Cmd":                                 reflect.ValueOf(osgui.Cmd),                         // Export Cmd constant
		"Control":                             reflect.ValueOf(osgui.Control),                     // Export Control constant
		"Ctrl":                                reflect.ValueOf(osgui.Ctrl),                        // Export Ctrl constant
		"Delete":                              reflect.ValueOf(osgui.Delete),                      // Export Delete constant
		"Down":                                reflect.ValueOf(osgui.Down),                        // Export Down constant
		"End":                                 reflect.ValueOf(osgui.End),                         // Export End constant
		"Enter":                               reflect.ValueOf(osgui.Enter),                       // Export Enter constant
		"Esc":                                 reflect.ValueOf(osgui.Esc),                         // Export Esc constant
		"Escape":                              reflect.ValueOf(osgui.Escape),                      // Export Escape constant
		"F1":                                  reflect.ValueOf(osgui.F1),                          // Export F1 constant
		"F10":                                 reflect.ValueOf(osgui.F10),                         // Export F10 constant
		"F11":                                 reflect.ValueOf(osgui.F11),                         // Export F11 constant
		"F12":                                 reflect.ValueOf(osgui.F12),                         // Export F12 constant
		"F13":                                 reflect.ValueOf(osgui.F13),                         // Export F13 constant
		"F14":                                 reflect.ValueOf(osgui.F14),                         // Export F14 constant
		"F15":                                 reflect.ValueOf(osgui.F15),                         // Export F15 constant
		"F16":                                 reflect.ValueOf(osgui.F16),                         // Export F16 constant
		"F17":                                 reflect.ValueOf(osgui.F17),                         // Export F17 constant
		"F18":                                 reflect.ValueOf(osgui.F18),                         // Export F18 constant
		"F19":                                 reflect.ValueOf(osgui.F19),                         // Export F19 constant
		"F2":                                  reflect.ValueOf(osgui.F2),                          // Export F2 constant
		"F20":                                 reflect.ValueOf(osgui.F20),                         // Export F20 constant
		"F21":                                 reflect.ValueOf(osgui.F21),                         // Export F21 constant
		"F22":                                 reflect.ValueOf(osgui.F22),                         // Export F22 constant
		"F23":                                 reflect.ValueOf(osgui.F23),                         // Export F23 constant
		"F24":                                 reflect.ValueOf(osgui.F24),                         // Export F24 constant
		"F3":                                  reflect.ValueOf(osgui.F3),                          // Export F3 constant
		"F4":                                  reflect.ValueOf(osgui.F4),                          // Export F4 constant
		"F5":                                  reflect.ValueOf(osgui.F5),                          // Export F5 constant
		"F6":                                  reflect.ValueOf(osgui.F6),                          // Export F6 constant
		"F7":                                  reflect.ValueOf(osgui.F7),                          // Export F7 constant
		"F8":                                  reflect.ValueOf(osgui.F8),                          // Export F8 constant
		"F9":                                  reflect.ValueOf(osgui.F9),                          // Export F9 constant
		"GraphicShape":                        reflect.ValueOf((*osgui.GraphicShape)(nil)).Elem(), // Export GraphicShape type
		"GraphicShape_ArrowDown":              reflect.ValueOf(osgui.GraphicShape_ArrowDown),      // Export GraphicShape_ArrowDown constant
		"GraphicShape_ArrowLeft":              reflect.ValueOf(osgui.GraphicShape_ArrowLeft),      // Export GraphicShape_ArrowLeft constant
		"GraphicShape_ArrowRight":             reflect.ValueOf(osgui.GraphicShape_ArrowRight),     // Export GraphicShape_ArrowRight constant
		"GraphicShape_ArrowUp":                reflect.ValueOf(osgui.GraphicShape_ArrowUp),        // Export GraphicShape_ArrowUp constant
		"GraphicShape_Circle":                 reflect.ValueOf(osgui.GraphicShape_Circle),         // Export GraphicShape_Circle constant
		"GraphicShape_CircleCross":            reflect.ValueOf(osgui.GraphicShape_CircleCross),    // Export GraphicShape_CircleCross constant
		"GraphicShape_CrossSign":              reflect.ValueOf(osgui.GraphicShape_CrossSign),      // Export GraphicShape_CrossSign constant
		"GraphicShape_Ellipse":                reflect.ValueOf(osgui.GraphicShape_Ellipse),        // Export GraphicShape_Ellipse constant
		"GraphicShape_Pentagon":               reflect.ValueOf(osgui.GraphicShape_Pentagon),       // Export GraphicShape_Pentagon constant
		"GraphicShape_PlusSign":               reflect.ValueOf(osgui.GraphicShape_PlusSign),       // Export GraphicShape_PlusSign constant
		"GraphicShape_Rectangle":              reflect.ValueOf(osgui.GraphicShape_Rectangle),      // Export GraphicShape_Rectangle constant
		"GraphicShape_Star":                   reflect.ValueOf(osgui.GraphicShape_Star),           // Export GraphicShape_Star constant
		"GraphicShape_Triangle":               reflect.ValueOf(osgui.GraphicShape_Triangle),       // Export GraphicShape_Triangle constant
		"GuiWindow":                           reflect.ValueOf((*osgui.GuiWindow)(nil)),           // Export GuiWindow interface pointer type
		"GuiWindow.Activate":                  reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("Activate"),
		"GuiWindow.BodyLocator":               reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("BodyLocator"),
		"GuiWindow.Close":                     reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("Close"),
		"GuiWindow.DeActivate":                reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("DeActivate"),
		"GuiWindow.GetClientRect":             reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetClientRect"),
		"GuiWindow.GetHwnd":                   reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetHwnd"),
		"GuiWindow.GetID":                     reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetID"),
		"GuiWindow.GetInitiatorPath":          reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetInitiatorPath"),
		"GuiWindow.GetMonitor":                reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetMonitor"),
		"GuiWindow.GetWindowCaretPos":         reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetWindowCaretPos"),
		"GuiWindow.GetWindowCursorPos":        reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetWindowCursorPos"),
		"GuiWindow.GetWindowRect":             reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetWindowRect"),
		"GuiWindow.GetWindowSchema":           reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetWindowSchema"),
		"GuiWindow.GetWindowTitle":            reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("GetWindowTitle"),
		"GuiWindow.MoveTo":                    reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("MoveTo"),
		"GuiWindow.PressKeys":                 reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("PressKeys"),
		"GuiWindow.ReadClipboard":             reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("ReadClipboard"),
		"GuiWindow.RectLocator":               reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("RectLocator"),
		"GuiWindow.ResizeTo":                  reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("ResizeTo"),
		"GuiWindow.Snapshot":                  reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("Snapshot"),
		"GuiWindow.ToMap":                     reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("ToMap"),
		"GuiWindow.TransFromScreen":           reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("TransFromScreen"),
		"GuiWindow.TransToScreen":             reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("TransToScreen"),
		"GuiWindow.WriteClipboard":            reflect.ValueOf((*osgui.GuiWindow)(nil)).MethodByName("WriteClipboard"),
		"Home":                                reflect.ValueOf(osgui.Home),                    // Export Home constant
		"Insert":                              reflect.ValueOf(osgui.Insert),                  // Export Insert constant
		"Key0":                                reflect.ValueOf(osgui.Key0),                    // Export Key0 constant
		"Key1":                                reflect.ValueOf(osgui.Key1),                    // Export Key1 constant
		"Key2":                                reflect.ValueOf(osgui.Key2),                    // Export Key2 constant
		"Key3":                                reflect.ValueOf(osgui.Key3),                    // Export Key3 constant
		"Key4":                                reflect.ValueOf(osgui.Key4),                    // Export Key4 constant
		"Key5":                                reflect.ValueOf(osgui.Key5),                    // Export Key5 constant
		"Key6":                                reflect.ValueOf(osgui.Key6),                    // Export Key6 constant
		"Key7":                                reflect.ValueOf(osgui.Key7),                    // Export Key7 constant
		"Key8":                                reflect.ValueOf(osgui.Key8),                    // Export Key8 constant
		"Key9":                                reflect.ValueOf(osgui.Key9),                    // Export Key9 constant
		"KeyA":                                reflect.ValueOf(osgui.KeyA),                    // Export KeyA constant
		"KeyB":                                reflect.ValueOf(osgui.KeyB),                    // Export KeyB constant
		"KeyC":                                reflect.ValueOf(osgui.KeyC),                    // Export KeyC constant
		"KeyD":                                reflect.ValueOf(osgui.KeyD),                    // Export KeyD constant
		"KeyE":                                reflect.ValueOf(osgui.KeyE),                    // Export KeyE constant
		"KeyF":                                reflect.ValueOf(osgui.KeyF),                    // Export KeyF constant
		"KeyG":                                reflect.ValueOf(osgui.KeyG),                    // Export KeyG constant
		"KeyH":                                reflect.ValueOf(osgui.KeyH),                    // Export KeyH constant
		"KeyI":                                reflect.ValueOf(osgui.KeyI),                    // Export KeyI constant
		"KeyJ":                                reflect.ValueOf(osgui.KeyJ),                    // Export KeyJ constant
		"KeyK":                                reflect.ValueOf(osgui.KeyK),                    // Export KeyK constant
		"KeyL":                                reflect.ValueOf(osgui.KeyL),                    // Export KeyL constant
		"KeyM":                                reflect.ValueOf(osgui.KeyM),                    // Export KeyM constant
		"KeyN":                                reflect.ValueOf(osgui.KeyN),                    // Export KeyN constant
		"KeyO":                                reflect.ValueOf(osgui.KeyO),                    // Export KeyO constant
		"KeyP":                                reflect.ValueOf(osgui.KeyP),                    // Export KeyP constant
		"KeyQ":                                reflect.ValueOf(osgui.KeyQ),                    // Export KeyQ constant
		"KeyR":                                reflect.ValueOf(osgui.KeyR),                    // Export KeyR constant
		"KeyS":                                reflect.ValueOf(osgui.KeyS),                    // Export KeyS constant
		"KeyT":                                reflect.ValueOf(osgui.KeyT),                    // Export KeyT constant
		"KeyU":                                reflect.ValueOf(osgui.KeyU),                    // Export KeyU constant
		"KeyV":                                reflect.ValueOf(osgui.KeyV),                    // Export KeyV constant
		"KeyW":                                reflect.ValueOf(osgui.KeyW),                    // Export KeyW constant
		"KeyX":                                reflect.ValueOf(osgui.KeyX),                    // Export KeyX constant
		"KeyY":                                reflect.ValueOf(osgui.KeyY),                    // Export KeyY constant
		"KeyZ":                                reflect.ValueOf(osgui.KeyZ),                    // Export KeyZ constant
		"Keyboard":                            reflect.ValueOf((*osgui.Keyboard)(nil)).Elem(), // Export Keyboard type
		"Keyboard.String":                     reflect.ValueOf(osgui.Keyboard.String),
		"Lalt":                                reflect.ValueOf(osgui.Lalt),            // Export Lalt constant
		"Lcmd":                                reflect.ValueOf(osgui.Lcmd),            // Export Lcmd constant
		"Lctrl":                               reflect.ValueOf(osgui.Lctrl),           // Export Lctrl constant
		"Left":                                reflect.ValueOf(osgui.Left),            // Export Left constant
		"LightsKbdDown":                       reflect.ValueOf(osgui.LightsKbdDown),   // Export LightsKbdDown constant
		"LightsKbdToggle":                     reflect.ValueOf(osgui.LightsKbdToggle), // Export LightsKbdToggle constant
		"LightsKbdUp":                         reflect.ValueOf(osgui.LightsKbdUp),     // Export LightsKbdUp constant
		"LightsMonDown":                       reflect.ValueOf(osgui.LightsMonDown),   // Export LightsMonDown constant
		"LightsMonUp":                         reflect.ValueOf(osgui.LightsMonUp),     // Export LightsMonUp constant
		"Locator":                             reflect.ValueOf((*osgui.Locator)(nil)), // Export Locator interface pointer type
		"Locator.ClearText":                   reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ClearText"),
		"Locator.Click":                       reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("Click"),
		"Locator.DoubleClick":                 reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("DoubleClick"),
		"Locator.Focus":                       reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("Focus"),
		"Locator.GetBodyRect":                 reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetBodyRect"),
		"Locator.GetLocatorCaretPos":          reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetLocatorCaretPos"),
		"Locator.GetLocatorCursorPos":         reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetLocatorCursorPos"),
		"Locator.GetScreenRect":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetScreenRect"),
		"Locator.GetSize":                     reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetSize"),
		"Locator.GetWindowRect":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("GetWindowRect"),
		"Locator.ImageLocator":                reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ImageLocator"),
		"Locator.ImageLocators":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ImageLocators"),
		"Locator.IsEditing":                   reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("IsEditing"),
		"Locator.IsHorizontalScroller":        reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("IsHorizontalScroller"),
		"Locator.IsVerticalScroller":          reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("IsVerticalScroller"),
		"Locator.MouseMove":                   reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("MouseMove"),
		"Locator.Ocr":                         reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("Ocr"),
		"Locator.ReadText":                    reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ReadText"),
		"Locator.RightClick":                  reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("RightClick"),
		"Locator.ScrollHorizontal":            reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ScrollHorizontal"),
		"Locator.ScrollVertical":              reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ScrollVertical"),
		"Locator.ShapeLocator":                reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ShapeLocator"),
		"Locator.ShapeLocators":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ShapeLocators"),
		"Locator.Snapshot":                    reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("Snapshot"),
		"Locator.SubLocator":                  reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("SubLocator"),
		"Locator.TextLocator":                 reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TextLocator"),
		"Locator.TextLocators":                reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TextLocators"),
		"Locator.ToMap":                       reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("ToMap"),
		"Locator.TransFromBody":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransFromBody"),
		"Locator.TransFromScreen":             reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransFromScreen"),
		"Locator.TransFromWindow":             reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransFromWindow"),
		"Locator.TransToBody":                 reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransToBody"),
		"Locator.TransToScreen":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransToScreen"),
		"Locator.TransToWindow":               reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("TransToWindow"),
		"Locator.WaitForEditing":              reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("WaitForEditing"),
		"Locator.WriteText":                   reflect.ValueOf((*osgui.Locator)(nil)).MethodByName("WriteText"),
		"Lshift":                              reflect.ValueOf(osgui.Lshift),                      // Export Lshift constant
		"Menu":                                reflect.ValueOf(osgui.Menu),                        // Export Menu constant
		"NewTextField":                        reflect.ValueOf(osgui.NewTextField),                // Export NewTextField function
		"Num0":                                reflect.ValueOf(osgui.Num0),                        // Export Num0 constant
		"Num1":                                reflect.ValueOf(osgui.Num1),                        // Export Num1 constant
		"Num2":                                reflect.ValueOf(osgui.Num2),                        // Export Num2 constant
		"Num3":                                reflect.ValueOf(osgui.Num3),                        // Export Num3 constant
		"Num4":                                reflect.ValueOf(osgui.Num4),                        // Export Num4 constant
		"Num5":                                reflect.ValueOf(osgui.Num5),                        // Export Num5 constant
		"Num6":                                reflect.ValueOf(osgui.Num6),                        // Export Num6 constant
		"Num7":                                reflect.ValueOf(osgui.Num7),                        // Export Num7 constant
		"Num8":                                reflect.ValueOf(osgui.Num8),                        // Export Num8 constant
		"NumClear":                            reflect.ValueOf(osgui.NumClear),                    // Export NumClear constant
		"NumDecimal":                          reflect.ValueOf(osgui.NumDecimal),                  // Export NumDecimal constant
		"NumDiv":                              reflect.ValueOf(osgui.NumDiv),                      // Export NumDiv constant
		"NumEnter":                            reflect.ValueOf(osgui.NumEnter),                    // Export NumEnter constant
		"NumEqual":                            reflect.ValueOf(osgui.NumEqual),                    // Export NumEqual constant
		"NumLock":                             reflect.ValueOf(osgui.NumLock),                     // Export NumLock constant
		"NumMinus":                            reflect.ValueOf(osgui.NumMinus),                    // Export NumMinus constant
		"NumMul":                              reflect.ValueOf(osgui.NumMul),                      // Export NumMul constant
		"NumPlus":                             reflect.ValueOf(osgui.NumPlus),                     // Export NumPlus constant
		"Pagedown":                            reflect.ValueOf(osgui.Pagedown),                    // Export Pagedown constant
		"Pageup":                              reflect.ValueOf(osgui.Pageup),                      // Export Pageup constant
		"Print":                               reflect.ValueOf(osgui.Print),                       // Export Print constant
		"Printscreen":                         reflect.ValueOf(osgui.Printscreen),                 // Export Printscreen constant
		"Ralt":                                reflect.ValueOf(osgui.Ralt),                        // Export Ralt constant
		"Rcmd":                                reflect.ValueOf(osgui.Rcmd),                        // Export Rcmd constant
		"Rctrl":                               reflect.ValueOf(osgui.Rctrl),                       // Export Rctrl constant
		"Right":                               reflect.ValueOf(osgui.Right),                       // Export Right constant
		"Rshift":                              reflect.ValueOf(osgui.Rshift),                      // Export Rshift constant
		"ShapeField":                          reflect.ValueOf((*osgui.ShapeField)(nil)).Elem(),   // Export ShapeField struct
		"Shift":                               reflect.ValueOf(osgui.Shift),                       // Export Shift constant
		"Space":                               reflect.ValueOf(osgui.Space),                       // Export Space constant
		"StraightLine":                        reflect.ValueOf((*osgui.StraightLine)(nil)).Elem(), // Export StraightLine struct
		"StraightLine.AngleWithHorizontal":    reflect.ValueOf(osgui.StraightLine.AngleWithHorizontal),
		"StraightLine.AngleWithHorizontalDeg": reflect.ValueOf(osgui.StraightLine.AngleWithHorizontalDeg),
		"StraightLine.Clone":                  reflect.ValueOf(osgui.StraightLine.Clone),
		"StraightLine.Length":                 reflect.ValueOf(osgui.StraightLine.Length),
		"Tab":                                 reflect.ValueOf(osgui.Tab),                      // Export Tab constant
		"TextField":                           reflect.ValueOf((*osgui.TextField)(nil)).Elem(), // Export TextField struct
		"Up":                                  reflect.ValueOf(osgui.Up),                       // Export Up constant
	}
	Symbols["yanlingrpa.com/yanling/protocol/ossys"] = map[string]reflect.Value{
		"DeviceInfo":                    reflect.ValueOf((*ossys.DeviceInfo)(nil)), // Export DeviceInfo interface pointer type
		"DeviceInfo.DeviceId":           reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("DeviceId"),
		"DeviceInfo.DeviceName":         reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("DeviceName"),
		"DeviceInfo.GetComputerName":    reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("GetComputerName"),
		"DeviceInfo.GetGpuMemoryMB":     reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("GetGpuMemoryMB"),
		"DeviceInfo.GetMonitors":        reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("GetMonitors"),
		"DeviceInfo.GetPrimaryMonitor":  reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("GetPrimaryMonitor"),
		"DeviceInfo.GetUserName":        reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("GetUserName"),
		"DeviceInfo.HasNvidiaGPU":       reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("HasNvidiaGPU"),
		"DeviceInfo.NumLogicCPU":        reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("NumLogicCPU"),
		"DeviceInfo.OS":                 reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("OS"),
		"DeviceInfo.OSVersion":          reflect.ValueOf((*ossys.DeviceInfo)(nil)).MethodByName("OSVersion"),
		"HttpClient":                    reflect.ValueOf((*ossys.HttpClient)(nil)), // Export HttpClient interface pointer type
		"HttpClient.DeleteDomainCookie": reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("DeleteDomainCookie"),
		"HttpClient.DeleteDomainHeader": reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("DeleteDomainHeader"),
		"HttpClient.DownloadFile":       reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("DownloadFile"),
		"HttpClient.Get":                reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("Get"),
		"HttpClient.GetDomainCookies":   reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("GetDomainCookies"),
		"HttpClient.GetDomainHeaders":   reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("GetDomainHeaders"),
		"HttpClient.Post":               reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("Post"),
		"HttpClient.PostForm":           reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("PostForm"),
		"HttpClient.PostJson":           reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("PostJson"),
		"HttpClient.SetDomainCookies":   reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("SetDomainCookies"),
		"HttpClient.SetDomainHeaders":   reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("SetDomainHeaders"),
		"HttpClient.UploadData":         reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("UploadData"),
		"HttpClient.UploadFile":         reflect.ValueOf((*ossys.HttpClient)(nil)).MethodByName("UploadFile"),
		"LocalFilesystem":               reflect.ValueOf((*ossys.LocalFilesystem)(nil)), // Export LocalFilesystem interface pointer type
		"LocalFilesystem.CopyFile":      reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("CopyFile"),
		"LocalFilesystem.CreateTmpFile": reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("CreateTmpFile"),
		"LocalFilesystem.IsDir":         reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("IsDir"),
		"LocalFilesystem.IsFile":        reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("IsFile"),
		"LocalFilesystem.JoinDataPath":  reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("JoinDataPath"),
		"LocalFilesystem.MkdirAll":      reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("MkdirAll"),
		"LocalFilesystem.PathExists":    reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("PathExists"),
		"LocalFilesystem.ReadFile":      reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("ReadFile"),
		"LocalFilesystem.Remove":        reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("Remove"),
		"LocalFilesystem.RemoveAll":     reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("RemoveAll"),
		"LocalFilesystem.Rename":        reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("Rename"),
		"LocalFilesystem.WriteFile":     reflect.ValueOf((*ossys.LocalFilesystem)(nil)).MethodByName("WriteFile"),
		"LocalStorage":                  reflect.ValueOf((*ossys.LocalStorage)(nil)), // Export LocalStorage interface pointer type
		"LocalStorage.Del":              reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("Del"),
		"LocalStorage.Get":              reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("Get"),
		"LocalStorage.Keys":             reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("Keys"),
		"LocalStorage.MGet":             reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("MGet"),
		"LocalStorage.MSet":             reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("MSet"),
		"LocalStorage.MSetEx":           reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("MSetEx"),
		"LocalStorage.Set":              reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("Set"),
		"LocalStorage.SetEx":            reflect.ValueOf((*ossys.LocalStorage)(nil)).MethodByName("SetEx"),
		"MonitorInfo":                   reflect.ValueOf((*ossys.MonitorInfo)(nil)), // Export MonitorInfo interface pointer type
		"MonitorInfo.GetBounds":         reflect.ValueOf((*ossys.MonitorInfo)(nil)).MethodByName("GetBounds"),
		"MonitorInfo.GetDPI":            reflect.ValueOf((*ossys.MonitorInfo)(nil)).MethodByName("GetDPI"),
		"MonitorInfo.GetWorkArea":       reflect.ValueOf((*ossys.MonitorInfo)(nil)).MethodByName("GetWorkArea"),
		"MonitorInfo.IsPrimary":         reflect.ValueOf((*ossys.MonitorInfo)(nil)).MethodByName("IsPrimary"),
		"ScriptLogger":                  reflect.ValueOf((*ossys.ScriptLogger)(nil)), // Export ScriptLogger interface pointer type
		"ScriptLogger.Debug":            reflect.ValueOf((*ossys.ScriptLogger)(nil)).MethodByName("Debug"),
		"ScriptLogger.Error":            reflect.ValueOf((*ossys.ScriptLogger)(nil)).MethodByName("Error"),
		"ScriptLogger.Info":             reflect.ValueOf((*ossys.ScriptLogger)(nil)).MethodByName("Info"),
		"ScriptLogger.Warn":             reflect.ValueOf((*ossys.ScriptLogger)(nil)).MethodByName("Warn"),
	}
	Symbols["yanlingrpa.com/yanling/protocol/script"] = map[string]reflect.Value{
		"ApiParameter":                   reflect.ValueOf((*script.ApiParameter)(nil)).Elem(),   // Export ApiParameter struct
		"Event":                          reflect.ValueOf((*script.Event)(nil)).Elem(),          // Export Event struct
		"EventHandler":                   reflect.ValueOf((*script.EventHandler)(nil)).Elem(),   // Export EventHandler type
		"ExportApi":                      reflect.ValueOf((*script.ExportApi)(nil)).Elem(),      // Export ExportApi struct
		"ExportTopic":                    reflect.ValueOf((*script.ExportTopic)(nil)).Elem(),    // Export ExportTopic struct
		"GuiApplication":                 reflect.ValueOf((*script.GuiApplication)(nil)).Elem(), // Export GuiApplication struct
		"ModuleInfo":                     reflect.ValueOf((*script.ModuleInfo)(nil)).Elem(),     // Export ModuleInfo struct
		"ModuleRuntime":                  reflect.ValueOf((*script.ModuleRuntime)(nil)),         // Export ModuleRuntime interface pointer type
		"ModuleRuntime.BrowserWindow":    reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("BrowserWindow"),
		"ModuleRuntime.CurrentSpecifier": reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("CurrentSpecifier"),
		"ModuleRuntime.DeviceInfo":       reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("DeviceInfo"),
		"ModuleRuntime.FileSystem":       reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("FileSystem"),
		"ModuleRuntime.GetCacheData":     reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("GetCacheData"),
		"ModuleRuntime.GetVariable":      reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("GetVariable"),
		"ModuleRuntime.GuiWindow":        reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("GuiWindow"),
		"ModuleRuntime.HostSpecifier":    reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("HostSpecifier"),
		"ModuleRuntime.HttpClient":       reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("HttpClient"),
		"ModuleRuntime.InvokeApi":        reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("InvokeApi"),
		"ModuleRuntime.Logger":           reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Logger"),
		"ModuleRuntime.Publish":          reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Publish"),
		"ModuleRuntime.SetCacheData":     reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("SetCacheData"),
		"ModuleRuntime.Storage":          reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Storage"),
		"ModuleRuntime.Subscribe":        reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Subscribe"),
		"ModuleRuntime.Unsubscribe":      reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Unsubscribe"),
		"ModuleRuntime.Vision":           reflect.ValueOf((*script.ModuleRuntime)(nil)).MethodByName("Vision"),
		"ParseSpecifier":                 reflect.ValueOf(script.ParseSpecifier),                // Export ParseSpecifier function
		"ScriptVariable":                 reflect.ValueOf((*script.ScriptVariable)(nil)).Elem(), // Export ScriptVariable struct
		"Specifier":                      reflect.ValueOf((*script.Specifier)(nil)).Elem(),      // Export Specifier struct
		"Specifier.Identifier":           reflect.ValueOf(script.Specifier.Identifier),
		"Specifier.String":               reflect.ValueOf(script.Specifier.String),
		"Subscriber":                     reflect.ValueOf((*script.Subscriber)(nil)), // Export Subscriber interface pointer type
		"Subscriber.GetSpecifier":        reflect.ValueOf((*script.Subscriber)(nil)).MethodByName("GetSpecifier"),
		"Subscriber.GetTopic":            reflect.ValueOf((*script.Subscriber)(nil)).MethodByName("GetTopic"),
		"Subscriber.IsActive":            reflect.ValueOf((*script.Subscriber)(nil)).MethodByName("IsActive"),
		"UrlPermission":                  reflect.ValueOf((*script.UrlPermission)(nil)).Elem(),    // Export UrlPermission struct
		"VariableBoolean":                reflect.ValueOf(script.VariableBoolean),                 // Export VariableBoolean constant
		"VariableDataType":               reflect.ValueOf((*script.VariableDataType)(nil)).Elem(), // Export VariableDataType type
		"VariableDataType.Parse":         reflect.ValueOf(script.VariableDataType.Parse),
		"VariableDataType.ToString":      reflect.ValueOf(script.VariableDataType.ToString),
		"VariableFilePath":               reflect.ValueOf(script.VariableFilePath),              // Export VariableFilePath constant
		"VariableInteger":                reflect.ValueOf(script.VariableInteger),               // Export VariableInteger constant
		"VariableJson":                   reflect.ValueOf(script.VariableJson),                  // Export VariableJson constant
		"VariableNumber":                 reflect.ValueOf(script.VariableNumber),                // Export VariableNumber constant
		"VariableString":                 reflect.ValueOf(script.VariableString),                // Export VariableString constant
		"WebApplication":                 reflect.ValueOf((*script.WebApplication)(nil)).Elem(), // Export WebApplication struct
		"YanLingScript":                  reflect.ValueOf((*script.YanLingScript)(nil)).Elem(),  // Export YanLingScript struct
	}
}
