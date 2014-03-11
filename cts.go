// Connexionz Transportation System
package cts

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
// Information
//

func (c *CTS) Platforms() ([]*Platform, error) {

	//Make request
	resp, err := c.xmlResponseForMethod("Platform.xml", nil)
	if err != nil {
		return nil, err
	}

	q := &PlatformQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	return q.Platforms, nil
}

func (c *CTS) PlatformGroups() ([]*PlatformGroup, error) {

	//Make request
	resp, err := c.xmlResponseForMethod("PlatformGroup.xml", nil)
	if err != nil {
		return nil, err
	}

	q := &PlatformGroupQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	return q.PlatformGroups, nil
}

func (c *CTS) Patterns() ([]*Route, error) {
	resp, err := c.xmlResponseForMethod("RoutePattern.xml", nil)
	if err != nil {
		return nil, err
	}

	q := &PatternQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	for _, r := range q.Project.Routes {
		p := r.Destination[0].Patterns[0]
		elems := strings.Fields(p.Mif)

		p.Polyline = make([]*Coordinate, (len(elems)-44)/2)
		for i := 42; i < len(elems)-2; i += 2 {
			val1, _ := strconv.ParseFloat(elems[i], 64)
			val2, _ := strconv.ParseFloat(elems[i+1], 64)

			coord := &Coordinate{
				Latitude:  val1,
				Longitude: val2,
			}

			p.Polyline[(i-42)/2] = coord
		}

	}

	return q.Project.Routes, nil
}

//Find the Estimated Time of Arrivals (ETA) for the given platform
func (c *CTS) ETA(p *Platform) ([]*Route, error) {
	if p.Tag == "" && p.Number == "" {
		return nil, errors.New("Need a valid platform tag or platform number")
	}

	resp, err := c.xmlResponseForMethod("RoutePositionET.xml", map[string]string{
		"PlatformNo":  p.Number,
		"PlatformTag": p.Tag,
	})

	if err != nil {
		return nil, err
	}

	q := &PlatformQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	//Check if invalid platform
	if len(q.Platforms) == 0 {
		return nil, errors.New("Invalid platform")
	}

	return q.Platforms[0].Routes, nil
}

func (c *CTS) MasterSchedules() ([]*Project, error) {
	//Make request
	resp, err := c.xmlResponseForMethod("ScheduleMaster.xml", nil)
	if err != nil {
		return nil, err
	}

	q := &ScheduleQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	return q.Projects, nil

}

func (c *CTS) ServiceSchedules(serviceName string) ([]*Project, error) {
	//Make request
	resp, err := c.xmlResponseForMethod("ScheduleDetail.xml", map[string]string{"ServiceName": serviceName})
	if err != nil {
		return nil, err
	}

	q := &ScheduleQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp, q)
	if err != nil {
		return nil, err
	}

	return q.Projects, nil
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

//
// Types
//

type Content struct {
	Expires         string `xml:"Expires,attr"`
	MaxArrivalScope string `xml:"MaxArrivalScope,attr"`
}

type PlatformQuery struct {
	Content   *Content    `xml:"Content"`
	Platforms []*Platform `xml:"Platform"`
}

type Platform struct {
	Tag      string      `xml:"PlatformTag,attr"`
	Number   string      `xml:"PlatformNo,attr"`
	Name     string      `xml:"Name,attr"`
	Bearing  string      `xml:"BearingToRoad,attr"`
	RoadName string      `xml:"RoadName,attr"`
	Location *Coordinate `xml:"Position"`

	ScheduleAdheranceTimepoint bool `xml:"ScheduleAdheranceTimepoint,attr"`

	Routes []*Route `xml:"Route"` //Only populated by after calling ETA()
}

type Coordinate struct {
	Latitude  float64 `xml:"Lat,attr"`
	Longitude float64 `xml:"Long,attr"`
}

type PlatformGroupQuery struct {
	Content        *Content         `xml:"Content"`
	PlatformGroups []*PlatformGroup `xml:"PlatformGroup"`
}

type PlatformGroup struct {
	Name      string      `xml:",attr"`
	Platforms []*Platform `xml:"Platform"` // Does no fill in Location or BearingToRoad
}

type ScheduleQuery struct {
	Content  Content    `xml:"Content"`
	Projects []*Project `xml:"Project"`
}

type Project struct {
	ID       string    `xml:"ProjectID,attr"`
	Name     string    `xml:"Name,attr"`
	Schedule *Schedule `xml:"Schedule"`
	Routes   []*Route  `xml:"Route"`
}

type Schedule struct {
	ValidFrom string   `xml:"ValidFrom,attr"`
	Routes    []*Route `xml:"Route"`
}

type Route struct {
	Number string            `xml:"RouteNo,attr"`
	Name   string            `xml:"Name,attr"`
	Group  *DestinationGroup `xml:"DestinationGroup"`

	Destination []*Destination `xml:"Destination"`
}

type DestinationGroup struct {
	Name     string     `xml:"Name,attr"`
	ID       string     `xml:"ID,attr"`
	Services []*Service `xml:"Service"`
}
type Destination struct {
	Name     string     `xml:"Name,attr"`
	Trip     *Trip      `xml:"Trip"`
	Patterns []*Pattern `xml:"Pattern"`
}
type Trip struct {
	ETA string `xml:"ETA,attr"`
}
type Service struct {
	Name string `xml:"Name,attr"`
}

type PatternQuery struct {
	Content *Content `xml:"Content"`
	Project *Project `xml:"Project"`
}

type Pattern struct {
	Tag       string `xml: "RouteTag"`
	Name      string `xml: "Name"`
	Length    string `xml: "Length"`
	Direction string `xml: "Direction"`
	Schedule  string `xml: "Schedule"`

	Mid string `xml: "Mid,chardata"`
	Mif string `xml: "Mif,chardata"`

	Polyline  []*Coordinate
	Platforms []*Platform `xml: "Platform"`
}
