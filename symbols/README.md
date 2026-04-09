# RPA Script API Reference

## 📋 Overview

This document lists all standard library and business API interfaces available to RPA scripts. All interfaces have been security reviewed to ensure scripts cannot access the local filesystem (except through controlled interfaces) or directly access the network (provided through controlled HTTP client).

**Version**: 1.0  
**Last Updated**: 2025-11-07

---

## 🔒 Security Notice

### ✅ What Scripts Can Do
- String, number, date, and other data processing
- JSON/XML/CSV data format conversion
- Regular expression matching
- Image processing (in memory)
- Access HTTP, filesystem, and local storage through controlled interfaces
- Browser automation operations
- GUI window automation operations

### ❌ What Scripts Cannot Do
- Direct access to local filesystem (must use `LocalFilesystem` interface)
- Direct network connections (must use `HttpClient` interface)
- Execute system commands
- Access environment variables or system information (except through `DeviceInfo` interface)

---

## 📦 Standard Library Packages

### 1. String and Text Processing

#### `strings` - String Operations
Common functions: `Contains`, `Split`, `Join`, `Replace`, `Trim`, `ToUpper`, `ToLower`, etc.

#### `strconv` - String Conversion
Common functions: `Atoi`, `Itoa`, `ParseInt`, `ParseFloat`, `FormatInt`, `FormatFloat`, etc.

#### `regexp` - Regular Expressions
Common functions: `Compile`, `Match`, `MatchString`, `FindString`, `ReplaceAllString`, etc.

#### `unicode` - Unicode Character Properties
Common functions: `IsLetter`, `IsDigit`, `IsSpace`, `ToUpper`, `ToLower`, etc.

#### `unicode/utf8` - UTF-8 Encoding
Common functions: `RuneCount`, `Valid`, `DecodeRune`, `EncodeRune`, etc.

---

### 2. Data Encoding/Decoding

#### `encoding/json` - JSON Encoding/Decoding
Common functions: `Marshal`, `Unmarshal`, `MarshalIndent`, etc.

#### `encoding/xml` - XML Encoding/Decoding
Common functions: `Marshal`, `Unmarshal`, etc.

#### `encoding/csv` - CSV Processing
For reading and writing CSV format data

#### `encoding/base64` - Base64 Encoding/Decoding
Common functions: `StdEncoding.EncodeToString`, `StdEncoding.DecodeString`, etc.

#### `encoding/hex` - Hexadecimal Encoding/Decoding
Common functions: `EncodeToString`, `DecodeString`, etc.

---

### 3. Cryptography and Hashing

#### `crypto/md5` - MD5 Hash
Common functions: `New`, `Sum`, etc.

#### `crypto/sha256` - SHA256 Hash
Common functions: `New`, `Sum256`, etc.

#### `hash` - Hash Interface
Provides generic hash interface

---

### 4. Mathematics

#### `math` - Math Functions
Common functions: `Abs`, `Ceil`, `Floor`, `Max`, `Min`, `Pow`, `Sqrt`, `Sin`, `Cos`, etc.

#### `math/big` - Big Integers and High Precision Calculation
Common types: `Int`, `Float`, `Rat`  
For scenarios requiring high precision or very large numbers

#### `math/rand` - Random Number Generation
Common functions: `Intn`, `Float64`, `Shuffle`, `Seed`, etc.  
**Note**: Pseudo-random number generator, suitable for general scenarios

---

### 5. Data Structures and Sorting

#### `sort` - Sorting
Common functions: `Ints`, `Strings`, `Float64s`, `Slice`, `SliceStable`, etc.

#### `container/heap` - Heap
For implementing priority queues and similar data structures

#### `container/list` - Doubly Linked List
Provides linked list data structure

#### `container/ring` - Circular List
Provides circular linked list data structure

---

### 6. Bytes and Buffering

#### `bytes` - Byte Operations
Common functions: `Contains`, `Split`, `Join`, `Replace`, `Trim`, etc.  
Similar to `strings` but operates on byte slices

#### `bufio` - Buffered I/O
Provides buffered readers and writers

#### `io` - I/O Interfaces
Provides basic interfaces like `Reader`, `Writer`, etc.  
**Note**: Only provides interfaces, cannot directly access files

---

### 7. Formatting and Error Handling

#### `fmt` - Formatted Output
Common functions: `Sprintf`, `Printf`, `Errorf`, etc.

#### `errors` - Error Handling
Common functions: `New`, `Is`, `As`, `Unwrap`, etc.

---

### 8. Paths and URLs

#### `path` - Path Operations (POSIX style)
Common functions: `Base`, `Dir`, `Ext`, `Join`, `Clean`, etc.  
**Note**: String operations only, does not access filesystem

#### `path/filepath` - File Path Operations
Common functions: `Base`, `Dir`, `Ext`, `Join`, `Clean`, `Split`, etc.  
**Note**: Functions like `Walk`, `Glob`, `Abs` that access filesystem have been removed

#### `net/url` - URL Parsing
Common functions: `Parse`, `ParseQuery`, `QueryEscape`, `JoinPath`, etc.  
**Note**: Only parses URL strings, does not access network

---

### 9. Time and Context

#### `time` - Time Handling
Common functions: `Now`, `Parse`, `Format`, `Sleep`, `Since`, `Until`, etc.  
Common types: `Time`, `Duration`, `Timer`, `Ticker`

#### `context` - Context Management
Common functions: `Background`, `TODO`, `WithCancel`, `WithTimeout`, `WithValue`, etc.  
For controlling timeouts and cancellations

---

### 10. Image Processing

#### `image` - Image Structures
Common functions: `Decode`, `NewRGBA`, `Rect`, `Pt`, etc.  
Common types: `Image`, `RGBA`, `Gray`, `Point`, `Rectangle`

#### `image/color` - Color Models
Common types: `RGBA`, `Gray`, `Alpha`, etc.  
Common functions: `RGBToYCbCr`, `YCbCrToRGB`, etc.

#### `image/png` - PNG Images
Common functions: `Decode`, `Encode`

#### `image/jpeg` - JPEG Images
Common functions: `Decode`, `Encode`

#### `image/gif` - GIF Images
Common functions: `Decode`, `DecodeAll`, `Encode`, `EncodeAll`

---

### 11. HTML and Templates

#### `html` - HTML Escaping
Common functions: `EscapeString`, `UnescapeString`

#### `text/template` - Text Templates
Common functions: `New`, `Parse` (string parsing only)  
**Note**: File access functions like `ParseFiles` have been removed

---

### 12. Concurrency Control

#### `sync` - Synchronization Primitives
Common types: `Mutex`, `RWMutex`, `WaitGroup`, `Once`, `Pool`, `Map`  
For concurrency safety

#### `sync/atomic` - Atomic Operations
Provides atomic-level numeric operations

---

## 🏢 Business API Packages

### 1. `basic` - Basic Data Structures

#### Data Types
- `Point` - Coordinate point (X, Y)
- `FPoint` - Floating point coordinate
- `Size` - Size (Width, Height)
- `Rect` - Rectangle area
- `OcrText` - OCR recognized text
- `OcrResult` - OCR recognition result
- `Task` - Task interface
- `DispatchTaskData` - Task dispatch data

#### Utility Functions
- `MaxAreaRect` - Get maximum area rectangle
- `MinAreaRect` - Get minimum area rectangle
- `MergeOverlappingRectangles` - Merge overlapping rectangles
- `MergeGroupRectangles` - Merge rectangles by group
- `MergeAllRectangles` - Merge all rectangles

---

### 2. `browser` - Browser Automation

#### `BrowserWindow` - Browser Window
- `CurrentPage()` - Get current page
- `DefaultPage()` - Get default page
- `NewTabPage(url)` - Open new tab
- `IDTabPage(id)` - Get tab by ID

#### `BrowserTabPage` - Browser Tab
- `Activate()` - Activate tab
- `Destroy()` - Close tab
- `SaveCookies()` / `LoadCookies()` - Cookie management
- `SaveLocalStorage()` / `LoadLocalStorage()` - LocalStorage management
- `ClearCookies()` / `ClearLocalStorage()` - Clear data
- `WaitForNewTab()` - Wait for new tab

#### `BrowserFramePage` - Page/Frame
- `GetURL()` / `GetTitle()` / `GetDomain()` - Get page info
- `QuerySelector(css)` - CSS selector query
- `QuerySelectorAll(css)` - Query all matching elements
- `QueryXPath(xpath)` / `QueryXPathAll(xpath)` - XPath query
- `WaitSelector(css, timeout)` - Wait for element
- `Evaluate(script)` - Execute JavaScript
- `Reload()` - Refresh page

#### `BrowserElement` - Page Element
- `Click()` / `DoubleClick()` / `RightClick()` - Click operations
- `Hover()` / `MouseMove()` - Mouse movements
- `Input(text)` / `WriteText(text)` - Input text
- `Focus()` / `Blur()` - Focus control
- `Text()` / `Html()` - Get text/HTML
- `GetAttribute(name)` / `SetAttribute(name, value)` - Attribute operations
- `GetProperty(name)` / `SetProperty(name, value)` - Property operations
- `Visible()` / `Disabled()` / `Interactable()` - State queries
- `WaitVisible()` / `WaitEnabled()` - Wait for states
- `ScrollIntoView()` - Scroll into view
- `SelectByText()` / `SelectByCss()` - Dropdown selection
- `SetFiles(paths)` - File upload

---

### 3. `osgui` - GUI Window Automation

#### `GuiWindow` - GUI Window
- `GetWindowTitle()` - Get window title
- `GetWindowRect()` / `GetClientRect()` - Get window area
- `GetMonitor()` - Get monitor
- `Activate()` / `DeActivate()` - Activate/deactivate window
- `Close()` - Close window
- `MoveTo(x, y)` / `ResizeTo(w, h)` - Move/resize
- `Snapshot(rect)` - Take screenshot
- `BodyLocator()` / `RectLocator(rect)` - Create locators
- `PressKeys(keys)` - Key operations
- `ReadClipboard()` / `WriteClipboard(text)` - Clipboard operations
- `GetWindowCursorPos()` / `GetWindowCaretPos()` - Get cursor position
- `TransToScreen(point)` / `TransFromScreen(point)` - Coordinate conversion

#### `Locator` - Element Locator
- `Click(point)` / `DoubleClick(point)` / `RightClick(point)` - Click
- `MouseMove(point)` - Mouse movement
- `Focus()` - Focus
- `Snapshot(rect)` - Screenshot
- `Ocr(rect)` - OCR recognition
- `ReadText()` / `WriteText(text)` / `ClearText()` - Text operations
- `IsEditing()` / `WaitForEditing()` - Editing state
- `ScrollVertical(amount)` / `ScrollHorizontal(amount)` - Scrolling
- `ImageLocator(template, similarity)` - Image matching
- `ImageLocators(template, similarity)` - Multiple image matching
- `TextLocator(text)` / `TextLocators(text)` - Text locating
- `ShapeLocator(shape)` / `ShapeLocators(shape)` - Shape locating
- `SubLocator(rect)` - Create sub-locator
- `TransToScreen()` / `TransFromScreen()` - Coordinate conversion

#### Keyboard Constants
Provides all common key constants: `Enter`, `Esc`, `Ctrl`, `Alt`, `Shift`, `F1`-`F24`, `KeyA`-`KeyZ`, `Key0`-`Key9`, etc.

#### Graphic Shape Constants
- `GraphicShape_Circle` - Circle
- `GraphicShape_Rectangle` - Rectangle
- `GraphicShape_Triangle` - Triangle
- `GraphicShape_Star` - Star
- etc.

---

### 4. `ossys` - System Services

#### `DeviceInfo` - Device Information
- `DeviceId()` / `DeviceName()` - Device identification
- `GetComputerName()` / `GetUserName()` - Computer/user names
- `OS()` / `OSVersion()` - Operating system info
- `NumLogicCPU()` - CPU core count
- `GetGpuMemoryMB()` / `HasNvidiaGPU()` - GPU information
- `GetMonitors()` / `GetPrimaryMonitor()` - Monitor information

#### `MonitorInfo` - Monitor Information
- `GetBounds()` / `GetWorkArea()` - Screen areas
- `GetDPI()` - DPI information
- `IsPrimary()` - Is primary monitor

#### `HttpClient` - HTTP Client (Controlled)
- `Get(url, headers)` - GET request
- `Post(url, data, headers)` - POST request
- `PostJson(url, jsonData, headers)` - POST JSON
- `PostForm(url, formData, headers)` - POST form
- `DownloadFile(url, savePath)` - Download file
- `UploadFile(url, filePath, headers)` - Upload file
- `UploadData(url, data, filename, headers)` - Upload data
- `SetDomainCookies(domain, cookies)` - Set cookies
- `GetDomainCookies(domain)` - Get cookies
- `SetDomainHeaders(domain, headers)` - Set headers
- `GetDomainHeaders(domain)` - Get headers

#### `LocalFilesystem` - Local Filesystem (Controlled)
- `JoinDataPath(paths...)` - Join data directory path
- `PathExists(path)` - Check path existence
- `IsFile(path)` / `IsDir(path)` - Check file/directory
- `ReadFile(path)` - Read file
- `WriteFile(path, data)` - Write file
- `CopyFile(src, dst)` - Copy file
- `Rename(old, new)` - Rename
- `Remove(path)` / `RemoveAll(path)` - Delete
- `MkdirAll(path)` - Create directory
- `CreateTmpFile()` - Create temporary file

#### `LocalStorage` - Local Storage (Key-Value)
- `Get(key)` / `Set(key, value)` - Read/write
- `SetEx(key, value, ttl)` - Write with expiration
- `MGet(keys...)` / `MSet(kvPairs...)` - Batch operations
- `Del(keys...)` - Delete
- `Keys(pattern)` - Query keys

#### `ScriptLogger` - Logging
- `Debug(message)` - Debug log
- `Info(message)` - Info log
- `Warn(message)` - Warning log
- `Error(message)` - Error log

---

### 5. `robot` - Script Execution Framework

#### `ProjectExecutor` - Project Executor
- `GetID()` - Get project ID
- `GetVariable(key)` - Get project variable
- `IsInterrupts()` - Check for interruption
- `SleepRandom(min, max)` - Random sleep
- `LockScreen()` / `UnlockScreen()` - Lock/unlock screen
- `DeviceInfo()` - Get device info
- `HttpClient()` - Get HTTP client
- `FileSystem()` - Get filesystem
- `Storage()` - Get local storage
- `Logger()` - Get logger
- `DispatchTask(taskData)` - Dispatch task

#### Script Types
- `BrowserScript` - Browser script interface
- `GuiScript` - GUI script interface
- `GoScript` - Go script interface
- `WxAppScript` - WeChat mini program script interface

#### Operator Type Constants
- `BrowserOperator` - Browser operator
- `GuiOperator` - GUI operator
- `WxAppOperator` - WeChat mini program operator

---

### 6. `wxapp` - WeChat Mini Program Automation

#### `WxAppWindow` - WeChat Mini Program Window
Provides WeChat mini program automation capabilities (specific methods exposed based on implementation)

---

## 📖 Usage Examples

### Example 1: String Processing
```go
import "strings"

text := "Hello, RPA World!"
upper := strings.ToUpper(text)        // "HELLO, RPA WORLD!"
parts := strings.Split(text, ", ")    // ["Hello", "RPA World!"]
joined := strings.Join(parts, " - ")  // "Hello - RPA World!"
```

### Example 2: JSON Processing
```go
import "encoding/json"

// Serialize
data := map[string]interface{}{
    "name": "RPA Task",
    "count": 42,
}
jsonBytes, _ := json.Marshal(data)
jsonStr := string(jsonBytes)

// Deserialize
var result map[string]interface{}
json.Unmarshal(jsonBytes, &result)
```

### Example 3: HTTP Requests
```go
// Get HttpClient from ProjectExecutor
httpClient := executor.HttpClient()

// GET request
response, _ := httpClient.Get("https://api.example.com/data", nil)

// POST JSON
jsonData := `{"name": "test"}`
response, _ := httpClient.PostJson("https://api.example.com/create", jsonData, nil)

// Download file
httpClient.DownloadFile("https://example.com/file.pdf", "downloaded.pdf")
```

### Example 4: Browser Automation
```go
// Get browser window (injected by framework)
page := browserWindow.DefaultPage()

// Navigate to page
page.Evaluate(`window.location.href = "https://example.com"`)

// Query element and click
button := page.QuerySelector("#submit-button")
button.Click()

// Input text
input := page.QuerySelector("input[name='username']")
input.Input("myusername")

// Wait for element
page.WaitSelector(".success-message", 5000)
```

### Example 5: GUI Automation
```go
// Get GUI window (injected by framework)
window.Activate()

// Create locator
locator := window.BodyLocator()

// Image recognition click
imageLocator := locator.ImageLocator("button_template.png", 0.9)
imageLocator.Click(basic.Point{X: 10, Y: 10})

// OCR text recognition
ocrResult := locator.Ocr(basic.Rect{X: 100, Y: 100, Width: 200, Height: 50})
for _, text := range ocrResult.Texts {
    logger.Info("Recognized text: " + text.Text)
}

// Key press
window.PressKeys([]string{osgui.Ctrl, osgui.KeyC})  // Ctrl+C
```

### Example 6: File Operations
```go
// Get FileSystem from ProjectExecutor
fs := executor.FileSystem()

// Read file
content, _ := fs.ReadFile(fs.JoinDataPath("config.json"))

// Write file
fs.WriteFile(fs.JoinDataPath("output.txt"), []byte("Hello, RPA!"))

// Check file existence
exists := fs.PathExists(fs.JoinDataPath("data.csv"))
```

### Example 7: Local Storage
```go
// Get Storage from ProjectExecutor
storage := executor.Storage()

// Save data
storage.Set("last_run_time", "2025-11-07 10:30:00")
storage.SetEx("temp_token", "abc123", 3600)  // Expires in 1 hour

// Read data
lastRun := storage.Get("last_run_time")

// Batch operations
values := storage.MGet("key1", "key2", "key3")
```

---

## ⚠️ Important Limitations

1. **Filesystem Access Restricted**
   - Can only access files through `LocalFilesystem` interface
   - Access paths are sandboxed
   - Cannot use `os` package to directly operate files

2. **Network Access Restricted**
   - Can only access network through `HttpClient` interface
   - May have whitelist or blacklist restrictions
   - Cannot use `net` package to directly establish connections

3. **System Commands Prohibited**
   - Cannot execute external programs
   - Cannot access environment variables (unless through controlled interfaces)

4. **Reflection and Unsafe Operations Prohibited**
   - `reflect` package not exposed
   - `unsafe` package not exposed
   - Prevents bypassing security restrictions

---

## 📚 Additional Resources

- Complete API Documentation: [TBD]
- Example Script Library: [TBD]
- FAQ: [TBD]

---

**Document Version**: 1.0  
**Generated**: 2025-11-07