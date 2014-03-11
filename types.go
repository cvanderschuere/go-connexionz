package cts

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

	Mid string `xml: "Mid,chardata" json:"-"`
	Mif string `xml: "Mif,chardata" json:"-"`

	Polyline  []*Coordinate
	Platforms []*Platform `xml: "Platform"`
}
