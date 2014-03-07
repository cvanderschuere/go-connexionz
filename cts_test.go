package cts

import(
	"testing"
)

func TestAllPlatforms(t *testing.T){
	c := New("http://www.corvallistransit.com/")

	p,err := c.Platforms()

	if err != nil || len(p) == 0{
		t.Error(err,p)
	}

}

func TestAllPlatformGroups(t *testing.T){
	c := New("http://www.corvallistransit.com/")

	p,err := c.PlatformGroups()

	if err != nil || len(p) == 0{
		t.Error(err,p)
	}

}

func TestMasterSchedule(t *testing.T){
	c := New("http://www.corvallistransit.com/")

	s,err := c.MasterSchedules()

	if err != nil || s == nil{
		t.Error(err,s)
	}

	t.Log(s)
}