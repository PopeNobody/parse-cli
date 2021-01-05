package parsecli

import (
	"net/http"
	"net/url"
  "fmt"

	"github.com/facebookgo/parse"
	"github.com/facebookgo/stackerr"
  "github.com/davecgh/go-spew/spew"
)

// ParseAPIClient is the http client used by parse-cli
type ParseAPIClient struct {
	APIClient *parse.Client
  e *Env
}

func NewParseAPIClient(e *Env) (*ParseAPIClient, error) {
	baseURL, err := url.Parse(e.Server)
	if err != nil {
		return nil, stackerr.Newf("invalid server URL %q: %s", e.Server, err)
	}
  client := &ParseAPIClient{
		APIClient: &parse.Client{
			BaseURL: baseURL,
		},
	}
  client.e=e
  return client,nil
}

func (c *ParseAPIClient) appendCommonHeaders(header http.Header) http.Header {
	if header == nil {
		header = make(http.Header)
	}
	header.Add("User-Agent", UserAgent)
	return header
}

// Get performs a GET method call on the given url and unmarshal response into
// result.
func (c *ParseAPIClient) Get(u *url.URL, result interface{}) (*http.Response, error) {
  var req *http.Request  = &http.Request{Method: "GET", URL: u}
  spew.Dump(req)
	return c.Do( req, nil, result)
}

// Post performs a POST method call on the given url with the given body and
// unmarshal response into result.
func (c *ParseAPIClient) Post(u *url.URL, body, result interface{}) (*http.Response, error) {
	return c.Do(&http.Request{Method: "POST", URL: u}, body, result)
}

// Put performs a PUT method call on the given url with the given body and
// unmarshal response into result.
func (c *ParseAPIClient) Put(u *url.URL, body, result interface{}) (*http.Response, error) {
  fmt.Println("PUT:",u)
	return c.Do(&http.Request{Method: "PUT", URL: u}, body, result)
}

// Delete performs a DELETE method call on the given url and unmarshal response
// into result.
func (c *ParseAPIClient) Delete(u *url.URL, result interface{}) (*http.Response, error) {
  fmt.Print("DELETE: ",u)
	return c.Do(&http.Request{Method: "DELETE", URL: u}, nil, result)
}

// RoundTrip is a wrapper for parse.Client.RoundTrip
func (c *ParseAPIClient) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header = c.appendCommonHeaders(req.Header)
  fmt.Print("RoundTrip1: req.Header:")
  spew.Dump(req.Header)

	req.Header = c.appendCommonHeaders(req.Header)

  fmt.Print("RoundTrip2: req.Header:")
  spew.Dump(req.Header)
	return c.APIClient.RoundTrip(req)
}

// Do is a wrapper for parse.Client.Do
func (c *ParseAPIClient) Do(req *http.Request, body, result interface{}) (*http.Response, error) {
	req.Header = c.appendCommonHeaders(req.Header)
  fmt.Print("DO: req.Header: ")
  spew.Dump(req.Header)
	return c.APIClient.Do(req, body, result)
}

// WithCredentials is a wrapper for parse.Client.WithCredentials
func (c *ParseAPIClient) WithCredentials(cr parse.Credentials) *ParseAPIClient {
	c.APIClient = c.APIClient.WithCredentials(cr)
  fmt.Print("WithCredentials2: cr: ")
  spew.Dump(cr)
	return c
}
