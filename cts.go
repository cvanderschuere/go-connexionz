package cts

import(
	"net/http"
	"encoding/xml"
	"net/url"
	"io/ioutil"
)


type CTS struct{
	baseURL string

}

func New(baseURL string)(*CTS){
	c := new(CTS)
	c.baseURL = baseURL

	return c
}

//
// Information
//

type PlatformQuery struct {
	Platforms []Platform `xml:"Platform"`
}

type Platform struct{
	PlatformTag string `xml:",attr"`
	PlatformNo string `xml:",attr"`
	Name string `xml:",attr"`
	BearingToRoad string `xml:",attr"`
	RoadName string `xml:",attr"`
	Location Coordinate `xml:"Position"`
}

type Coordinate struct{
	Latitude string `xml:"Lat,attr"`
	Longitude string `xml:"Long,attr"`
}

func (c *CTS) Platforms()([]Platform,error){
	
	//Make request
	resp,err := c.xmlResponseForMethod("Platform.xml",nil)
	if err != nil{
		return nil,err
	}

	q := &PlatformQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp,q)
	if err != nil{
		return nil,err
	}

	return q.Platforms,nil
}

type PlatformGroupQuery struct {
	PlatformGroups []PlatformGroup `xml:"PlatformGroup"`
}

type PlatformGroup struct{
	Name string `xml:",attr"`
	Platform Platform `xml:"Platform"` // Does no fill in Location or BearingToRoad
}

func (c *CTS) PlatformGroups()([]PlatformGroup,error){
	
	//Make request
	resp,err := c.xmlResponseForMethod("PlatformGroup.xml",nil)
	if err != nil{
		return nil,err
	}

	q := &PlatformGroupQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp,q)
	if err != nil{
		return nil,err
	}

	return q.PlatformGroups,nil
}

type ScheduleQuery struct {
	Projects []Project `xml:"Project"`
}

type Project struct{
	ID string `xml:"ProjectID,attr"`
	Name string `xml:"Name,attr"`
	Schedule Schedule `xml:"Schedule"`
}

type Schedule struct{
	ValidFrom string `xml:"ValidFrom,attr"`
	Routes []Route `xml:"Route"`
}

type Route struct{
	Number string `xml:"RouteNo,attr"`
	Name string `xml:"Name,attr"`
	Group DestinationGroup `xml:"DestinationGroup"`
}

type DestinationGroup struct{
	Name string `xml:"Name,attr"`
	ID string `xml:"ID,attr"`
	Services []Service `xml:"Service"`
}

type Service struct{
	Name string `xml:"Name,attr"`
}


func (c *CTS) MasterSchedules()([]Project,error){
	//Make request
	resp,err := c.xmlResponseForMethod("ScheduleDetail.xml",nil)
	if err != nil{
		return nil,err
	}

	q := &ScheduleQuery{}

	//Unmarshal XML
	err = xml.Unmarshal(resp,q)
	if err != nil{
		return nil,err
	}

	return q.Projects,nil

}


//
//Convinence methods
//

func (c *CTS) xmlResponseForMethod(method string,options map[string]string)([]byte,error){
	//Convert paramaters
	params := url.Values{}
	for key,val := range options{
		params.Add(key,val)
	}

	params.Add("contenttype","SQLXML")
	params.Set("Name",method)

	//Perform GET request
	resp, err := http.Get(c.baseURL+"rtt/public/utility/file.aspx?"+params.Encode())
	if err != nil {
		return nil,err
	}

	//Read response
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}