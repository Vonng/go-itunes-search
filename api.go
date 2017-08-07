package itunes_search

import (
	"net/http"
	"net/url"
	"encoding/json"
	"strconv"
	"strings"
)

type lookupResult struct {
	ResultCount int         `json:"resultCount"`
	Results     []Entry    `json:"results"`
}

type Params struct {
	url.Values
	endpoint string
}

func (self Params) Term(term string) Params {
	self.Values.Set("term", term)
	return self
}

func (self Params) Terms(terms []string) Params {
	self.Values.Set("term", strings.Join(terms, "+"))
	return self
}

func (self Params) Country(country string) Params {
	self.Values.Set("country", country)
	return self
}

func (self Params) Entity(entity string) Params {
	self.Values.Set("entity", entity)
	return self
}

func (self Params) Entities(entities []string) Params {
	self.Values["entity"] = entities
	return self
}

func (self Params) AddEntity(entity string) Params {
	self.Values.Add("entity", entity)
	return self
}

func (self Params) Media(media string) Params {
	self.Values.Set("media", media)
	return self
}

func (self Params) Medias(medias []string) Params {
	self.Values["media"] = medias
	return self
}

func (self Params) AddMedia(media string) Params {
	self.Values.Add("media", media)
	return self
}

func (self Params) Limit(n int) Params {
	if n > 200 {
		n = 200
	}

	if n < 1 {
		n = 1
	}
	self.Values.Set("limit", strconv.Itoa(n))
	return self
}

func (self Params) ID(id int64) Params {
	self.Values.Set(KeyITunesID, strconv.FormatInt(id, 10))
	return self
}

func (self Params) BundleID(bundleID string) Params {
	self.Values.Set(KeyBundleID, bundleID)
	return self
}

// CNAPP: set search & lookup constraint: CN & software
func (self Params) CNAPP(id int64) Params {
	self.Values.Set(KeyITunesID, strconv.FormatInt(id, 10))
	self.Values.Set("media", MediaSoftware)
	self.Values.Set("country", CN)
	return self
}

// USAPP: set search & lookup constraint: US & software
func (self Params) USAPP(id int64) Params {
	self.Values.Set(KeyITunesID, strconv.FormatInt(id, 10))
	self.Values.Set("media", MediaSoftware)
	return self
}

// Results will finally do the request
func (self Params) Results() ([]Entry, error) {
	res, err := http.Get(self.endpoint + self.Encode())
	if err != nil {
		return nil, err
	}

	lr := new(lookupResult)
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(lr); err != nil {
		return nil, err
	}

	if lr.ResultCount == 0 || lr.Results == nil || len(lr.Results) == 0 {
		return nil, ErrNotFound
	}

	return lr.Results, nil
}

// Result assert there's one result
func (params Params) Result() (*Entry, error) {
	res, err := http.Get(LookupURL + params.Encode())
	if err != nil {
		return nil, err
	}

	lr := new(lookupResult)
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(lr); err != nil {
		return nil, err
	}

	if lr.ResultCount == 0 || lr.Results == nil || len(lr.Results) == 0 {
		return nil, ErrNotFound
	}

	return &(lr.Results[0]), nil
}

// Lookup begins the API chain
func Lookup() Params {
	return Params{make(url.Values), LookupURL}
}

// Search begins the API chain
func Search() Params {
	return Params{make(url.Values), SearchURL}
}
