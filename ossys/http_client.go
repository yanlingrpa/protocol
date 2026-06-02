package ossys

import "time"

/*
* HttpRequestOptions defines request-level control options.
* This struct is used to configure timeout and retry behavior for a single request.
 */
type HttpRequestOptions struct {
	/*
	* Timeout limits the maximum duration of a request.
	* Zero value means using implementation default.
	 */
	Timeout time.Duration

	/*
	* RetryCount defines how many times to retry when transient errors occur.
	* Zero means no retry.
	 */
	RetryCount int

	/*
	* RetryInterval defines the interval between retries.
	* Zero value means using implementation default.
	 */
	RetryInterval time.Duration

	/*
	* FollowRedirect controls whether redirects should be followed automatically.
	* If true, redirects are handled by the implementation.
	 */
	FollowRedirect bool
}

/*
* HttpRequest defines a general HTTP request model.
* This struct supports arbitrary methods and per-request configuration.
 */
type HttpRequest struct {
	/*
	* Method is the HTTP method, such as GET, POST, PUT, PATCH, DELETE, or HEAD.
	 */
	Method string

	/*
	* URL is the target request URL.
	 */
	URL string

	/*
	* Query defines URL query parameters.
	 */
	Query map[string]string

	/*
	* Headers defines request headers.
	 */
	Headers map[string]string

	/*
	* Body is the raw request body bytes.
	 */
	Body []byte

	/*
	* Form defines form fields for form submissions.
	 */
	Form map[string]any

	/*
	* Json defines JSON payload values for JSON requests.
	 */
	Json map[string]any

	/*
	* Options controls timeout and retry behavior for this request.
	 */
	Options HttpRequestOptions
}

/*
* HttpResponse defines a normalized HTTP response model.
* This struct contains status code, headers, and body data.
 */
type HttpResponse struct {
	/*
	* StatusCode is the HTTP status code returned by the server.
	 */
	StatusCode int

	/*
	* Headers contains response headers.
	 */
	Headers map[string]string

	/*
	* Body is the raw response body bytes.
	 */
	Body []byte

	/*
	* FinalURL is the final URL after redirects.
	 */
	FinalURL string
}

/*
* HttpClient defines capabilities for HTTP requests and domain-level configuration.
* This interface supports domain-level defaults, generic request execution, and common convenience operations.
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
	* Request sends a generalized HTTP request.
	* Supports arbitrary HTTP methods and returns structured response metadata.
	 */
	Request(req HttpRequest) (HttpResponse, error)

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
	* DownloadFileTo downloads a file to the specified local path.
	* options controls request timeout and retry behavior.
	 */
	DownloadFileTo(url string, savePath string, headers map[string]string, options HttpRequestOptions) (string, error)

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
