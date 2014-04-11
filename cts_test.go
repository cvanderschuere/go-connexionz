//Connexionz Transit System
package cts

import "testing"
import "fmt"

const baseURL = "http://www.corvallistransit.com/"

func TestAllPlatforms(t *testing.T) {
	c := New(baseURL)

	p, err := c.Platforms()

	fmt.Println(p[0])

	if err != nil || len(p) == 0 {
		t.Error(err, p)
	}
}

func TestAllPlatformsEtas(t *testing.T) {
	c := New(baseURL)

	p, err := c.Platforms()

	if err != nil || len(p) == 0 {
		t.Error(err, p)
	}

	for _, platform := range p {
		_, err := c.ETA(platform)

		if err != nil {
			t.Error(err, platform)
		}
	}

}

func TestPatterns(t *testing.T) {
	c := New(baseURL)
	r, err := c.Patterns()

	if err != nil || len(r) == 0 {
		t.Error(err, r)
	}

	//fmt.Println(r[0].Destination[0].Patterns[0].Polyline)
}

func TestAllPlatformGroups(t *testing.T) {
	c := New(baseURL)

	p, err := c.PlatformGroups()

	if err != nil || len(p) == 0 {
		t.Error(err, p)
	}

}

func TestMasterSchedule(t *testing.T) {
	c := New(baseURL)

	s, err := c.MasterSchedules()

	if err != nil || s == nil {
		t.Error(err, s)
	}

}

func TestScheduleDetail(t *testing.T) {
	c := New(baseURL)

	s, err := c.ServiceSchedules("Weekday")

	if err != nil || s == nil {
		t.Error(err, s)
	}

}
