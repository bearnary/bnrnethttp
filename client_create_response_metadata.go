package klnethttp

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func (c *defaultClient) createResponseMetadata(url string, r *resty.Response, start time.Time) *ResponseMetadata {

	resp := ResponseMetadata{
		Url:          url,
		StatusCode:   r.StatusCode(),
		ResponseTime: time.Since(start),
		Method:       r.Request.Method,
		Host:         r.Request.Header.Get("X-Forwarded-For"),
	}
	return &resp
}
