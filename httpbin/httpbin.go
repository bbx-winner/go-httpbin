package httpbin

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// Default configuration values
const (
	DefaultMaxBodySize int64 = 1024 * 1024
	DefaultMaxDuration       = 10 * time.Second
)

const (
	jsonContentType = "application/json; encoding=utf-8"
	htmlContentType = "text/html; charset=utf-8"
)

type headersResponse struct {
	Headers http.Header `json:"headers"`
}

type headersResponseCompat struct {
	Headers map[string]interface{} `json:"headers"`
}

func (hr headersResponse) MarshalJSON() ([]byte, error) {
	hrCompat := headersResponseCompat{
		Headers: make(map[string]interface{}),
	}

	for k, headers := range hr.Headers {
		if len(headers) != 1 {
			hrCompat.Headers[k] = headers
		} else {
			hrCompat.Headers[k] = headers[0]
		}
	}
	return json.Marshal(hrCompat)
}

type ipResponse struct {
	Origin string `json:"origin"`
}

type userAgentResponse struct {
	UserAgent string `json:"user-agent"`
}

type getResponse struct {
	Args    url.Values  `json:"args"`
	Headers http.Header `json:"headers"`
	Origin  string      `json:"origin"`
	URL     string      `json:"url"`
}

type getResponseCompat struct {
	Args    map[string]interface{} `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
}

func (gr getResponse) MarshalJSON() ([]byte, error) {
	grCompat := getResponseCompat{
		Args:    make(map[string]interface{}),
		Headers: make(map[string]interface{}),
		Origin:  gr.Origin,
		URL:     gr.URL,
	}

	for k, args := range gr.Args {
		if len(args) != 1 {
			grCompat.Args[k] = args
		} else {
			grCompat.Args[k] = args[0]
		}
	}
	for k, headers := range gr.Headers {
		if len(headers) != 1 {
			grCompat.Headers[k] = headers
		} else {
			grCompat.Headers[k] = headers[0]
		}
	}
	return json.Marshal(grCompat)
}

// A generic response for any incoming request that might contain a body
type bodyResponse struct {
	Args    url.Values  `json:"args"`
	Headers http.Header `json:"headers"`
	Origin  string      `json:"origin"`
	URL     string      `json:"url"`

	Data  string              `json:"data"`
	Files map[string][]string `json:"files"`
	Form  map[string][]string `json:"form"`
	JSON  interface{}         `json:"json"`
}

type bodyResponseCompat struct {
	Args    map[string]interface{} `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`

	Data  string                 `json:"data"`
	Files map[string]interface{} `json:"files"`
	Form  map[string]interface{} `json:"form"`
	JSON  interface{}            `json:"json"`
}

func (br bodyResponse) MarshalJSON() ([]byte, error) {
	brCompat := bodyResponseCompat{
		Args:    make(map[string]interface{}),
		Headers: make(map[string]interface{}),
		Origin:  br.Origin,
		URL:     br.URL,
		Data:    br.Data,
		Files:   make(map[string]interface{}),
		Form:    make(map[string]interface{}),
		JSON:    br.JSON,
	}

	for k, args := range br.Args {
		if len(args) != 1 {
			brCompat.Args[k] = args
		} else {
			brCompat.Args[k] = args[0]
		}
	}
	for k, headers := range br.Headers {
		if len(headers) != 1 {
			brCompat.Headers[k] = headers
		} else {
			brCompat.Headers[k] = headers[0]
		}
	}
	for k, files := range br.Files {
		if len(files) != 1 {
			brCompat.Files[k] = files
		} else {
			brCompat.Files[k] = files[0]
		}
	}
	for k, forms := range br.Form {
		if len(forms) != 1 {
			brCompat.Form[k] = forms
		} else {
			brCompat.Form[k] = forms[0]
		}
	}
	return json.Marshal(brCompat)
}

type cookiesResponse map[string]string

type authResponse struct {
	Authorized bool   `json:"authorized"`
	User       string `json:"user"`
}

type gzipResponse struct {
	Headers http.Header `json:"headers"`
	Origin  string      `json:"origin"`
	Gzipped bool        `json:"gzipped"`
}

type gzipResponseCompat struct {
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	Gzipped bool                   `json:"gzipped"`
}

func (gr gzipResponse) MarshalJSON() ([]byte, error) {
	grCompat := gzipResponseCompat{
		Headers: make(map[string]interface{}),
		Origin:  gr.Origin,
		Gzipped: gr.Gzipped,
	}

	for k, headers := range gr.Headers {
		if len(headers) != 1 {
			grCompat.Headers[k] = headers
		} else {
			grCompat.Headers[k] = headers[0]
		}
	}
	return json.Marshal(grCompat)
}

type deflateResponse struct {
	Headers  http.Header `json:"headers"`
	Origin   string      `json:"origin"`
	Deflated bool        `json:"deflated"`
}

type deflateResponseCompat struct {
	Headers  map[string]interface{} `json:"headers"`
	Origin   string                 `json:"origin"`
	Deflated bool                   `json:"deflated"`
}

func (dr deflateResponse) MarshalJSON() ([]byte, error) {
	drCompat := deflateResponseCompat{
		Headers:  make(map[string]interface{}),
		Origin:   dr.Origin,
		Deflated: dr.Deflated,
	}

	for k, headers := range dr.Headers {
		if len(headers) != 1 {
			drCompat.Headers[k] = headers
		} else {
			drCompat.Headers[k] = headers[0]
		}
	}
	return json.Marshal(drCompat)
}

// An actual stream response body will be made up of one or more of these
// structs, encoded as JSON and separated by newlines
type streamResponse struct {
	ID      int         `json:"id"`
	Args    url.Values  `json:"args"`
	Headers http.Header `json:"headers"`
	Origin  string      `json:"origin"`
	URL     string      `json:"url"`
}

type streamResponseCompat struct {
	ID      int                    `json:"id"`
	Args    map[string]interface{} `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
}

func (sr streamResponse) MarshalJSON() ([]byte, error) {
	srCompat := streamResponseCompat{
		ID:      sr.ID,
		Args:    make(map[string]interface{}),
		Headers: make(map[string]interface{}),
		Origin:  sr.Origin,
		URL:     sr.URL,
	}

	for k, args := range sr.Args {
		if len(args) != 1 {
			srCompat.Args[k] = args
		} else {
			srCompat.Args[k] = args[0]
		}
	}
	for k, headers := range sr.Headers {
		if len(headers) != 1 {
			srCompat.Headers[k] = headers
		} else {
			srCompat.Headers[k] = headers[0]
		}
	}
	return json.Marshal(srCompat)
}

type uuidResponse struct {
	UUID string `json:"uuid"`
}

type bearerResponse struct {
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token"`
}

// HTTPBin contains the business logic
type HTTPBin struct {
	// Max size of an incoming request generated response body, in bytes
	MaxBodySize int64

	// Max duration of a request, for those requests that allow user control
	// over timing (e.g. /delay)
	MaxDuration time.Duration

	// Observer called with the result of each handled request
	Observer Observer

	// Default parameter values
	DefaultParams DefaultParams
}

// DefaultParams defines default parameter values
type DefaultParams struct {
	DripDuration time.Duration
	DripDelay    time.Duration
	DripNumBytes int64
}

// DefaultDefaultParams defines the DefaultParams that are used by default. In
// general, these should match the original httpbin.org's defaults.
var DefaultDefaultParams = DefaultParams{
	DripDuration: 2 * time.Second,
	DripDelay:    2 * time.Second,
	DripNumBytes: 10,
}

// Handler returns an http.Handler that exposes all HTTPBin endpoints
func (h *HTTPBin) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", methods(h.Index, "GET"))
	mux.HandleFunc("/forms/post", methods(h.FormsPost, "GET"))
	mux.HandleFunc("/encoding/utf8", methods(h.UTF8, "GET"))

	mux.HandleFunc("/delete", methods(h.RequestWithBody, "DELETE"))
	mux.HandleFunc("/get", methods(h.Get, "GET"))
	mux.HandleFunc("/head", methods(h.Get, "HEAD"))
	mux.HandleFunc("/patch", methods(h.RequestWithBody, "PATCH"))
	mux.HandleFunc("/post", methods(h.RequestWithBody, "POST"))
	mux.HandleFunc("/put", methods(h.RequestWithBody, "PUT"))

	mux.HandleFunc("/ip", h.IP)
	mux.HandleFunc("/user-agent", h.UserAgent)
	mux.HandleFunc("/headers", h.Headers)
	mux.HandleFunc("/response-headers", h.ResponseHeaders)

	mux.HandleFunc("/status/", h.Status)
	mux.HandleFunc("/unstable", h.Unstable)

	mux.HandleFunc("/redirect/", h.Redirect)
	mux.HandleFunc("/relative-redirect/", h.RelativeRedirect)
	mux.HandleFunc("/absolute-redirect/", h.AbsoluteRedirect)
	mux.HandleFunc("/redirect-to", h.RedirectTo)

	mux.HandleFunc("/cookies", h.Cookies)
	mux.HandleFunc("/cookies/set", h.SetCookies)
	mux.HandleFunc("/cookies/delete", h.DeleteCookies)

	mux.HandleFunc("/basic-auth/", h.BasicAuth)
	mux.HandleFunc("/hidden-basic-auth/", h.HiddenBasicAuth)
	mux.HandleFunc("/digest-auth/", h.DigestAuth)
	mux.HandleFunc("/bearer", h.Bearer)

	mux.HandleFunc("/deflate", h.Deflate)
	mux.HandleFunc("/gzip", h.Gzip)

	mux.HandleFunc("/stream/", h.Stream)
	mux.HandleFunc("/delay/", h.Delay)
	mux.HandleFunc("/drip", h.Drip)

	mux.HandleFunc("/range/", h.Range)
	mux.HandleFunc("/bytes/", h.Bytes)
	mux.HandleFunc("/stream-bytes/", h.StreamBytes)

	mux.HandleFunc("/html", h.HTML)
	mux.HandleFunc("/robots.txt", h.Robots)
	mux.HandleFunc("/deny", h.Deny)

	mux.HandleFunc("/cache", h.Cache)
	mux.HandleFunc("/cache/", h.CacheControl)
	mux.HandleFunc("/etag/", h.ETag)

	mux.HandleFunc("/links/", h.Links)

	mux.HandleFunc("/image", h.ImageAccept)
	mux.HandleFunc("/image/", h.Image)
	mux.HandleFunc("/xml", h.XML)
	mux.HandleFunc("/json", h.JSON)

	mux.HandleFunc("/uuid", h.UUID)
	mux.HandleFunc("/base64/", h.Base64)

	// existing httpbin endpoints that we do not support
	mux.HandleFunc("/brotli", notImplementedHandler)

	// Make sure our ServeMux doesn't "helpfully" redirect these invalid
	// endpoints by adding a trailing slash. See the ServeMux docs for more
	// info: https://golang.org/pkg/net/http/#ServeMux
	mux.HandleFunc("/absolute-redirect", http.NotFound)
	mux.HandleFunc("/basic-auth", http.NotFound)
	mux.HandleFunc("/delay", http.NotFound)
	mux.HandleFunc("/digest-auth", http.NotFound)
	mux.HandleFunc("/hidden-basic-auth", http.NotFound)
	mux.HandleFunc("/redirect", http.NotFound)
	mux.HandleFunc("/relative-redirect", http.NotFound)
	mux.HandleFunc("/status", http.NotFound)
	mux.HandleFunc("/stream", http.NotFound)
	mux.HandleFunc("/bytes", http.NotFound)
	mux.HandleFunc("/stream-bytes", http.NotFound)
	mux.HandleFunc("/links", http.NotFound)

	// Apply global middleware
	var handler http.Handler
	handler = mux
	handler = limitRequestSize(h.MaxBodySize, handler)
	handler = preflight(handler)
	handler = autohead(handler)
	if h.Observer != nil {
		handler = observe(h.Observer, handler)
	}

	return handler
}

// New creates a new HTTPBin instance
func New(opts ...OptionFunc) *HTTPBin {
	h := &HTTPBin{
		MaxBodySize:   DefaultMaxBodySize,
		MaxDuration:   DefaultMaxDuration,
		DefaultParams: DefaultDefaultParams,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// OptionFunc uses the "functional options" pattern to customize an HTTPBin
// instance
type OptionFunc func(*HTTPBin)

// WithDefaultParams sets the default params handlers will use
func WithDefaultParams(defaultParams DefaultParams) OptionFunc {
	return func(h *HTTPBin) {
		h.DefaultParams = defaultParams
	}
}

// WithMaxBodySize sets the maximum amount of memory
func WithMaxBodySize(m int64) OptionFunc {
	return func(h *HTTPBin) {
		h.MaxBodySize = m
	}
}

// WithMaxDuration sets the maximum amount of time httpbin may take to respond
func WithMaxDuration(d time.Duration) OptionFunc {
	return func(h *HTTPBin) {
		h.MaxDuration = d
	}
}

// WithObserver sets the request observer callback
func WithObserver(o Observer) OptionFunc {
	return func(h *HTTPBin) {
		h.Observer = o
	}
}
