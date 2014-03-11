// +build !appengine

// Connexionz Transportation System
package cts

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type CTS struct {
	baseURL string
}

func New(baseURL string) *CTS {
	c := new(CTS)
	c.baseURL = baseURL

	return c
}

//
//Convinence methods
//

func (c *CTS) xmlResponseForMethod(method string, options map[string]string) ([]byte, error) {
	//Convert paramaters
	params := url.Values{}
	for key, val := range options {
		params.Add(key, val)
	}

	params.Add("contenttype", "SQLXML")
	params.Set("Name", method)

	//Perform GET request
	resp, err := http.Get(c.baseURL + "rtt/public/utility/file.aspx?" + params.Encode())
	if err != nil {
		return nil, err
	}

	//Read response
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
