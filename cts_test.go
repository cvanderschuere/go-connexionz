package cts

import (
	"testing"
)

const baseURL = "http://www.corvallistransit.com/"

func TestAllPlatforms(t *testing.T) {
	c := New(baseURL)

	p, err := c.Platforms()

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
