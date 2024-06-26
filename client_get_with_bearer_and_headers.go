package bnrnethttp

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *defaultClient) GetJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error) {

	start := time.Now()
	rc := c.client.R().
		SetAuthToken(bearerToken).
		SetBody(req)
	if headers != nil {
		for k, v := range *headers {
			rc = rc.SetHeader(k, v)
		}
	}

	r, err := rc.Get(url)
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
