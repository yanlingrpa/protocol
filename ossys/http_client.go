package ossys

/*
* HttpClient 定义 HTTP 请求与域名级配置能力
* 该接口支持 Cookie、Header 管理以及常见请求与上传下载操作
 */
type HttpClient interface {
	/*
	* SetDomainCookies 设置指定域名的 Cookie
	* cookies 使用键值对表示 Cookie 名称和值
	 */
	SetDomainCookies(domain string, cookies map[string]string) error

	/*
	* GetDomainCookies 获取指定域名的 Cookie
	* 返回域名下已保存的 Cookie 键值对
	 */
	GetDomainCookies(domain string) (map[string]string, error)

	/*
	* DeleteDomainCookie 删除指定域名下的 Cookie
	* 可一次删除多个 Cookie 名称
	 */
	DeleteDomainCookie(domain string, cookieNames ...string) error

	/*
	* SetDomainHeaders 设置指定域名的默认 Header
	* headers 使用键值对表示请求头
	 */
	SetDomainHeaders(domain string, headers map[string]string) error

	/*
	* GetDomainHeaders 获取指定域名的默认 Header
	* 返回域名下已配置的请求头键值对
	 */
	GetDomainHeaders(domain string) (map[string]string, error)

	/*
	* DeleteDomainHeader 删除指定域名下的 Header
	* 可一次删除多个 Header 名称
	 */
	DeleteDomainHeader(domain string, headerNames ...string) error

	/*
	* Get 发送 HTTP GET 请求
	* headers 参数用于覆盖或补充请求头
	 */
	Get(url string, headers map[string]string) ([]byte, error)

	/*
	* Post 发送 HTTP POST 原始数据请求
	* body 参数为请求体字节数组
	 */
	Post(url string, body []byte, headers map[string]string) ([]byte, error)

	/*
	* PostForm 发送 HTTP 表单请求
	* params 参数将以表单格式提交
	 */
	PostForm(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* PostJson 发送 HTTP JSON 请求
	* params 参数将编码为 JSON 作为请求体
	 */
	PostJson(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* DownloadFile 下载文件并返回本地路径
	* headers 参数用于附加下载请求头
	 */
	DownloadFile(url string, headers map[string]string) (string, error)

	/*
	* UploadFile 上传本地文件
	* filePath 为本地文件路径，params 为附加表单参数
	 */
	UploadFile(url string, filePath string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* UploadData 上传内存中的文件数据
	* fileName 为上传文件名，fileData 为上传内容
	 */
	UploadData(url string, fileName string, fileData []byte, params map[string]any, headers map[string]string) ([]byte, error)
}
