package bnrnethttp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/structs"
)

func (c *defaultClient) PostUrlEncodedWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error) {

	fd := structs.Map(req)

	fdm := MapInterfaceToMapString(fd)

	start := time.Now()
	rc := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(fdm)
	if headers != nil {
		for k, v := range *headers {
			rc = rc.SetHeader(k, v)
		}
	}

	r, err := rc.Post(url)
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
