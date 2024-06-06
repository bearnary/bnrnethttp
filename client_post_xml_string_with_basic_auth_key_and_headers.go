package klnethttp

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *defaultClient) PostXMLStringWithBasicAuthKeyAndHeaders(url string, basicAuthenKey string, headers *map[string]string, xmlBodyReq string, resp interface{}) (*ResponseMetadata, error) {

	start := time.Now()
	rc := c.client.R().
		SetHeader("Content-Type", "text/xml").
		SetBody(xmlBodyReq)
	if headers != nil {
		for k, v := range *headers {
			rc = rc.SetHeader(k, v)
		}
	}
	authVal := fmt.Sprintf("Basic %s", basicAuthenKey)
	rc.SetHeader("Authorization", authVal)

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
