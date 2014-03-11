go-connexionz
=============

A go wrapper for Connexionz APIs

Two known examples are :
1. [Corvallis, OR](http://www.corvallistransit.com/)
2. [San Clarita, CA](http://businfo.santa-clarita.com/)

__You can use these two URLS as base URLS__

# Documentation

[![GoDoc](https://godoc.org/github.com/cvanderschuere/go-connexionz?status.png)](https://godoc.org/github.com/cvanderschuere/go-connexionz)

The best documentation can be found on [GoDoc](http://godoc.org/github.com/cvanderschuere/go-connexionz)


# Usage

go-connexionz is built to work as a normal go packages as well as a package for Google App Engine

All documentation is for the non-GAE users.

On GAE, the only change is:

* Google App Engine
```go
  New(context appengine.Context, baseURL string) *CTS {
```
* Non-Google App Engine
```go
New(baseURL string) *CTS
```

## Example
```go
  // Create new client on GAE
  c := cts.New(context,baseURL)

  //Platforms
  p, err := c.Platforms()

  //Platform Groups
  g, err := c.PlatformGroups()

  //Master Schedule
  ms, err := c.MasterSchedules()

  //Route Patterns
  rp, err := c.Patterns()

  //Route ETAs
  platform := &cts.Platform{
		Tag: "360", //Must give either tag or number
	}
  r, err := c.ETA(platform)

```
