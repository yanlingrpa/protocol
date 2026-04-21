package ossys

/*
* HttpClient defines capabilities for HTTP requests and domain-level configuration.
* This interface supports cookie/header management and common request/upload/download operations.
 */
type HttpClient interface {
	/*
	* SetDomainCookies sets cookies for the specified domain.
	* cookies uses key-value pairs to represent cookie names and values.
	 */
	SetDomainCookies(domain string, cookies map[string]string) error

	/*
	* GetDomainCookies gets cookies for the specified domain.
	* Returns the saved cookie key-value pairs under that domain.
	 */
	GetDomainCookies(domain string) (map[string]string, error)

	/*
	* DeleteDomainCookie deletes cookies under the specified domain.
	* Multiple cookie names can be deleted at once.
	 */
	DeleteDomainCookie(domain string, cookieNames ...string) error

	/*
	* SetDomainHeaders sets default headers for the specified domain.
	* headers uses key-value pairs to represent request headers.
	 */
	SetDomainHeaders(domain string, headers map[string]string) error

	/*
	* GetDomainHeaders gets default headers for the specified domain.
	* Returns configured request-header key-value pairs under that domain.
	 */
	GetDomainHeaders(domain string) (map[string]string, error)

	/*
	* DeleteDomainHeader deletes headers under the specified domain.
	* Multiple header names can be deleted at once.
	 */
	DeleteDomainHeader(domain string, headerNames ...string) error

	/*
	* Get sends an HTTP GET request.
	* The headers argument is used to override or add request headers.
	 */
	Get(url string, headers map[string]string) ([]byte, error)

	/*
	* Post sends an HTTP POST request with raw data.
	* The body argument is the request body byte array.
	 */
	Post(url string, body []byte, headers map[string]string) ([]byte, error)

	/*
	* PostForm sends an HTTP form request.
	* The params argument is submitted in form format.
	 */
	PostForm(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* PostJson sends an HTTP JSON request.
	* The params argument is encoded as JSON as the request body.
	 */
	PostJson(url string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* DownloadFile downloads a file and returns the local path.
	* The headers argument is used to attach request headers for download.
	 */
	DownloadFile(url string, headers map[string]string) (string, error)

	/*
	* UploadFile uploads a local file.
	* filePath is the local file path, and params are additional form parameters.
	 */
	UploadFile(url string, filePath string, params map[string]any, headers map[string]string) ([]byte, error)

	/*
	* UploadData uploads file data from memory.
	* fileName is the upload file name, and fileData is the upload content.
	 */
	UploadData(url string, fileName string, fileData []byte, params map[string]any, headers map[string]string) ([]byte, error)
}
