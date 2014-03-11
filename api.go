package cts

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

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
				Latitude:  val2,
				Longitude: val1,
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
