package ossys

import "time"

/*
* HttpRequestOptions 定义请求级控制选项。
* 该结构体用于配置单次请求的超时和重试行为。
 */
type HttpRequestOptions struct {
	/*
	* Timeout 限制请求的最大持续时间。
	* 零值表示使用实现默认值。
	 */
	Timeout time.Duration

	/*
	* RetryCount 定义在瞬时错误发生时重试的次数。
	* 零表示不重试。
	 */
	RetryCount int

	/*
	* RetryInterval 定义重试之间的间隔时间。
	* 零值表示使用实现默认值。
	 */
	RetryInterval time.Duration

	/*
	* FollowRedirect 控制是否应自动跟随重定向。
	* 若为 true，则由实现处理重定向。
	 */
	FollowRedirect bool
}

/*
* HttpRequest 定义通用 HTTP 请求模型。
* 该结构体支持任意 HTTP 方法和每次请求的配置。
 */
type HttpRequest struct {
	/*
	* Method 是 HTTP 方法，例如 GET、POST、PUT、PATCH、DELETE 或 HEAD。
	 */
	Method string

	/*
	* URL 是目标请求地址。
	 */
	URL string

	/*
	* Query 定义 URL 查询参数。
	 */
	Query map[string]string

	/*
	* Headers 定义请求头。
	 */
	Headers map[string]string

	/*
	* Body 是原始请求体字节数据。
	 */
	Body []byte

	/*
	* Form 定义表单提交字段。
	 */
	Form map[string]any

	/*
	* Json 定义 JSON 请求体中的字段值。
	 */
	Json map[string]any

	/*
	* Options 控制当前请求的超时和重试行为。
	 */
	Options HttpRequestOptions
}

/*
* HttpResponse 定义规范化的 HTTP 响应模型。
* 该结构体包含状态码、响应头和响应体数据。
 */
type HttpResponse struct {
	/*
	* StatusCode 是服务器返回的 HTTP 状态码。
	 */
	StatusCode int

	/*
	* Headers 包含响应头。
	 */
	Headers map[string]string

	/*
	* Body 是原始响应体字节数据。
	 */
	Body []byte

	/*
	* FinalURL 是重定向后的最终 URL。
	 */
	FinalURL string
}

/*
* HttpClient 定义 HTTP 请求和域级配置的能力接口。
* 该接口支持域级默认值、通用请求执行以及常见便捷操作。
 */
type HttpClient interface {
	/*
	* SetDomainCookies 为指定域设置 cookies。
	* cookies 使用键值对表示 cookie 名称和对应值。
	 */
	SetDomainCookies(domain string, cookies map[string]string) error

	/*
	* GetDomainCookies 获取指定域下的 cookies。
	* 返回该域下已保存的 cookie 键值对。
	 */
	GetDomainCookies(domain string) (map[string]string, error)

	/*
	* DeleteDomainCookie 删除指定域下的 cookies。
	* 可一次删除多个 cookie 名称。
	 */
	DeleteDomainCookie(domain string, cookieNames ...string) error

	/*
	* SetDomainHeaders 为指定域设置默认请求头。
	* headers 使用键值对表示请求头名称和值。
	 */
	SetDomainHeaders(domain string, headers map[string]string) error

	/*
	* GetDomainHeaders 获取指定域下的默认请求头。
	* 返回该域下已配置的请求头键值对。
	 */
	GetDomainHeaders(domain string) (map[string]string, error)

	/*
	* DeleteDomainHeader 删除指定域下的请求头。
	* 可一次删除多个请求头名称。
	 */
	DeleteDomainHeader(domain string, headerNames ...string) error

	/*
	* Request 发送通用 HTTP 请求。
	* 支持任意 HTTP 方法，并返回结构化的响应元数据。
	 */
	Request(req HttpRequest) (HttpResponse, error)

	/*
	* Get 发送 HTTP GET 请求。
	* headers 参数用于覆盖或新增请求头。
	 */
	Get(url string, headers map[string]string) ([]byte, error)

	/*
	* Post 发送带原始数据的 HTTP POST 请求。
	* body 参数是请求体字节数组。
	 */
	Post(url string, body []byte, headers map[string]string) ([]byte, error)

	/*
	* PostForm 发送 HTTP 表单请求。
	* params 参数以表单格式提交。
	 */
	PostForm(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* PostJson 发送 HTTP JSON 请求。
	* params 参数会被编码为 JSON 作为请求体。
	 */
	PostJson(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* DownloadFile 下载文件并返回本地路径。
	* headers 参数用于附加下载请求头。
	 */
	DownloadFile(url string, headers map[string]string) (string, error)

	/*
	* DownloadFileTo 下载文件到指定本地路径。
	* options 控制请求超时和重试行为。
	 */
	DownloadFileTo(url string, savePath string, headers map[string]string, options HttpRequestOptions) (string, error)

	/*
	* UploadFile 上传本地文件。
	* filePath 是本地文件路径，params 是附加的表单参数。
	 */
	UploadFile(url string, filePath string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* UploadData 上传内存中的文件数据。
	* fileName 是上传文件名，fileData 是上传内容。
	 */
	UploadData(url string, fileName string, fileData []byte, params map[string]any, headers map[string]string) ([]byte, error)
}
