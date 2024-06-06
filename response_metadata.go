package klnethttp

import (
	"time"
)

type ResponseMetadata struct {
	Url          string
	StatusCode   int
	Method       string
	Host         string
	ResponseTime time.Duration
}
