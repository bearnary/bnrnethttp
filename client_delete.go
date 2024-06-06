package klnethttp

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *defaultClient) Delete(url string, resp interface{}) (*ResponseMetadata, error) {

	start := time.Now()
	r, err := c.client.R().
		Delete(url)
	if err != nil {
		return nil, fmt.Errorf("error calling http status code: %v err: %v", r.StatusCode(), err.Error())
	}

	err = json.Unmarshal(r.Body(), resp)
	if err != nil {
		return nil, fmt.Errorf("error parsing body status code: %v response: %v body: %v", r.StatusCode(),
			err.Error(), string(r.Body()))
	}

	respMetadata := c.createResponseMetadata(url, r, start)

	return respMetadata, nil
}
