package cts

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/twpayne/gopolyline/polyline"
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

		points := make([]float64, len(elems)-44)
		for i := 42; i < len(elems)-2; i += 2 {
			lon, _ := strconv.ParseFloat(elems[i], 64)
			lat, _ := strconv.ParseFloat(elems[i+1], 64)

			points[i-42] = lat
			points[(i+1)-42] = lon
		}

		p.Polyline = polyline.Encode(points, 2)

	}

	return q.Project.Routes, nil
}

//Find the Estimated Time of Arrivals (ETA) for the given platform
func (c *CTS) ETA(p *Platform) ([]*Route, error) {
	if p.Tag <= 0 && p.Number <= 0 {
		return nil, errors.New("Need a valid platform tag or platform number")
	}

	// Make map of information -- platform number before tag
	param := make(map[string]string)
	if p.Number != 0 {
		param["PlatformNo"] = strconv.FormatInt(p.Number, 10)
	} else {
		param["PlatformTag"] = strconv.FormatInt(p.Tag, 10)
	}

	resp, err := c.xmlResponseForMethod("RoutePositionET.xml", param)

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
