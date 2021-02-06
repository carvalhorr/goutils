package http

import (
	"io"
	h "net/http"
	"net/url"
)

type Client interface {
	CloseIdleConnections()
	Do(req *h.Request) (*h.Response, error)
	Get(url string) (resp *h.Response, err error)
	Head(url string) (resp *h.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *h.Response, err error)
	PostForm(url string, data url.Values) (resp *h.Response, err error)
}
