package bnrnethttp

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
)

func (c *defaultClient) PostMultipartFormDataWithAuthToken(url string, accessToken string, forms interface{}, files interface{}, resp interface{}) error {

	rc := c.client.R().SetAuthToken(accessToken).
		SetHeader("Content-Type", "multipart/form-data")

	if forms != nil {
		fd := structs.Map(forms)
		fdm := MapInterfaceToMapString(fd)
		rc = rc.SetFormData(fdm)
	}
	if files != nil {
		fid := structs.Map(files)
		fidm := MapInterfaceToMapString(fid)
		rc = rc.SetFiles(fidm)
	}

	r, err := rc.Post(url)
	if err != nil {
		return fmt.Errorf("error calling http status code: %v err: %v", r.StatusCode(), err.Error())
	}

	err = json.Unmarshal(r.Body(), resp)
	if err != nil {
		return fmt.Errorf("error parsing body status code: %v response: %v body: %v", r.StatusCode(),
			err.Error(), string(r.Body()))
	}

	return nil
}
