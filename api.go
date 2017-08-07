package itunes_search

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"encoding/json"
)

/**************************************************************
* Structure Define
**************************************************************/

// iTunesResult represent iTunes response outer most structure
type iTunesResult struct {
	ResultCount int         `json:"resultCount"`
	Results     []Entry     `json:"results"`
}

// Params holds iTunes API params.
// See following url for more details:
// https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/
type Params struct {
	url.Values
	endpoint string
}

/**************************************************************
* Enter Point
**************************************************************/

// Lookup begins the API chain
func Lookup() Params {
	return Params{make(url.Values), LookupURL}.Country(CN)
}

// Search begins the API chain with a series of search terms
func Search(terms []string) Params {
	return Params{make(url.Values), SearchURL}.Country(CN).Terms(terms)
}

// SearchOne begins the API chain with one term
func SearchOne(term string) Params {
	return Params{make(url.Values), SearchURL}.Country(CN).Term(term)
}

/**************************************************************
* Chain Method
**************************************************************/

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

func (self Params) ID(id int64) Params {
	self.Values.Set(ITunesID, strconv.FormatInt(id, 10))
	return self
}

func (self Params) BundleID(bundleID string) Params {
	self.Values.Set(BundleID, bundleID)
	return self
}

// App: restrict to application
func (self Params) App() Params {
	self.Values.Set("media", Software)
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

/**************************************************************
* End Point
**************************************************************/

// Results will finally do the request
func (self Params) Results() ([]Entry, error) {
	res, err := http.Get(self.endpoint + self.Encode())
	if err != nil {
		return nil, err
	}

	lr := new(iTunesResult)
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
func (self Params) Result() (*Entry, error) {
	if entries, err := self.Results(); err != nil {
		return nil, err
	} else {
		return &(entries[0]), nil
	}
}
