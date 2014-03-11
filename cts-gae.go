// +build appengine
// Connexionz Transportation System
package cts

import (
	"io/ioutil"
	"net/url"

	"appengine"
	"appengine/urlfetch" //Need to use this in-place of http.Get()
)

type CTS struct {
	baseURL string
	context appengine.Context
}

func New(context appengine.Context, baseURL string) *CTS {
	c := new(CTS)
	c.baseURL = baseURL
	c.context = context

	return c
}

//
//Convinence method
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
	client := urlfetch.Client(c.context)
	resp, err := client.Get(c.baseURL + "rtt/public/utility/file.aspx?" + params.Encode())
	if err != nil {
		return nil, err
	}

	//Read response
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
