package ossys

type HttpClient interface {
	SetDomainCookies(domain string, cookies map[string]string) error
	GetDomainCookies(domain string) (map[string]string, error)
	DeleteDomainCookie(domain string, cookieNames ...string) error

	SetDomainHeaders(domain string, headers map[string]string) error
	GetDomainHeaders(domain string) (map[string]string, error)
	DeleteDomainHeader(domain string, headerNames ...string) error

	Get(url string, headers map[string]string) ([]byte, error)
	Post(url string, body []byte, headers map[string]string) ([]byte, error)
	PostForm(url string, params map[string]any, headers map[string]string) ([]byte, error)
	PostJson(url string, params map[string]any, headers map[string]string) ([]byte, error)

	DownloadFile(url string, headers map[string]string) (string, error)
	UploadFile(url string, filePath string, params map[string]any, headers map[string]string) ([]byte, error)
	UploadData(url string, fileName string, fileData []byte, params map[string]any, headers map[string]string) ([]byte, error)
}
