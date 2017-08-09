package itunes_search

import (
	"os"
	"fmt"
	"time"
	"math"
	"sort"
	"strconv"
	"strings"
	"encoding/json"
	"text/template"
)

import "github.com/PuerkitoBio/goquery"

/**************************************************************
* struct:	App
**************************************************************/

// App represent iTunes application entity
// Some fields, like Platforms, InAppPurchase SiblingApps RelatedApps SupportSite & Reviews
// could only be fetched from iTunes page. a parser adjust for CN Store is provided
type App struct {
	ID               int64
	Name             string
	URL              string
	Icon             string
	Kind             string
	Version          string
	BundleID         string
	AuthorID         int64
	AuthorName       string
	AuthorURL        string
	VendorName       string
	VendorURL        string
	Copyright        string // reserved
	GenreID          int64
	GenreName        string
	Genres           []string
	GenreIDs         []int64
	Icon60           string
	Icon100          string
	Price            int64 // Since all price is ￥<int> or $<x.99>, use int rather than float
	Currency         string
	System           string
	Features         []string
	Devices          []string
	Languages        []string
	Platforms        []string // reserved
	Rating           string
	Reasons          []string
	Size             int64
	CntRating        int64
	AvgRating        float64
	CntRatingCurrent int64
	AvgRatingCurrent float64
	VppDevice        bool
	GameCenter       bool
	Screenshots      []string
	InAppPurchase    []string // reserved
	SiblingApps      []int64  // reserved
	RelatedApps      []int64  // reserved
	SupportSites     string   // reserved
	Reviews          string   // reserved
	Description      string
	ReleaseNotes     string
	ReleaseTime      time.Time
	PublishTime      time.Time
	CrawledTime      time.Time
}

/**************************************************************
* util:	App print auxiliary
**************************************************************/
const appTemplateStr = `
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ [{{.Kind}}] {{.ID}} {{.BundleID}} {{.Name}} {{.Version}}
┃ {{.URL}}
┃ {{.Icon}}
┃ Price: {{.Price}} {{.Currency}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Provider:
┃	{{.AuthorID}} {{.AuthorName}}  {{.AuthorURL}}
┃	{{.VendorName}} {{.VendorURL}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	{{.GenreID}} {{.GenreName}}
┃	{{.GenreIDs}} {{.Genres}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon60 : {{.Icon60}}
┃ Icon100: {{.Icon100}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃	System   : {{.System}}
┃	Features : {{.Features}}
┃	Devices  : {{.Devices}}
┃	Languages: {{.Languages}}
┃	Platforms: {{.Platforms}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating : {{.Rating}}
┃ Reasons: {{.Reasons}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Size: {{.Size}} VppDevice: {{.VppDevice}} GameCenter:{{.GameCenter}}
┃ Rating(Current)  : {{.AvgRatingCurrent}} / {{.CntRatingCurrent}}
┃ Rating(Historic) : {{.AvgRating}} / {{.CntRating}}
┃ Sibling Apps : {{.SiblingApps}}
┃ Related Apps : {{.RelatedApps}}
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Screenshots:{{range .Screenshots}}
	{{.}}{{end}}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Purchase:{{range .InAppPurchase}}
	{{.}}{{end}}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Reviews: {{.Reviews}}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Description:
{{.Description}}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Support Sites:
{{.SupportSites}}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  ReleaseNotes:
{{.ReleaseNotes}}
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Release: {{.ReleaseTime}}
┃ Publish: {{.PublishTime}}
┃ Crawled: {{.CrawledTime}}
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━`

// iosAppTmpl is text template for printing
var appTmpl, _ = template.New("app").Parse(appTemplateStr)

// Entry_Print will print human-readable entry
func (self *App) Print() {
	if err := appTmpl.Execute(os.Stdout, self); err != nil {
		fmt.Println(err.Error())
	}
	return
}

/**************************************************************
* constructor:	Entry.ToApp()
* 	Create App from Entry
**************************************************************/

// Entry_ToApp will generate a app record from entry
func (self *Entry) ToApp() (app *App) {
	app = new(App)
	app.ID = self.TrackID
	app.Name = self.TrackName
	app.URL = self.TrackViewURL

	app.Icon = self.ArtworkURL512
	app.Kind = self.Kind
	app.Version = self.Version
	app.BundleID = self.BundleID

	app.AuthorID = self.ArtistID
	app.AuthorName = self.ArtistName
	app.AuthorURL = self.ArtistViewURL
	app.VendorName = self.SellerName
	app.VendorURL = self.SellerURL

	app.GenreID = self.PrimaryGenreID
	app.GenreName = self.PrimaryGenreName
	app.Genres = self.Genres
	app.GenreIDs = stringSliceToInt(self.GenreIDs)

	app.Icon60 = self.ArtworkURL60
	app.Icon100 = self.ArtworkURL100
	app.Price = int64(math.Ceil(self.Price))
	app.Currency = self.Currency

	app.System = self.MinimumOsVersion
	app.Features = self.Features
	app.Devices = self.SupportedDevices
	app.Languages = self.LanguageCodesISO2A

	app.Rating = self.TrackContentRating
	if app.Rating == "" {
		app.Rating = self.ContentAdvisoryRating
	}
	app.Reasons = self.Advisories

	app.Size, _ = strconv.ParseInt(self.FileSizeBytes, 10, 64)
	app.CntRating = self.UserRatingCount
	app.AvgRating = self.AverageUserRating
	app.CntRatingCurrent = self.UserRatingCountForCurrentVersion
	app.AvgRatingCurrent = self.AverageUserRatingForCurrentVersion
	app.VppDevice = self.IsVppDeviceBasedLicensingEnabled
	app.GameCenter = self.IsGameCenterEnabled

	app.Screenshots = merge(self.ScreenshotURLs, self.AppletvScreenshotURLs, self.IpadScreenshotURLs)

	// Reserved fields: these fields should be fetched from iTunes page
	// app.Copyright
	// app.Platforms
	// app.InAppPurchase
	// app.SiblingApps
	// app.RelatedApps
	// app.SupportSites
	// app.Reviews

	app.Description = self.Description
	app.ReleaseNotes = self.ReleaseNotes
	app.ReleaseTime, _ = time.Parse(time.RFC3339, self.CurrentVersionReleaseDate)
	app.PublishTime, _ = time.Parse(time.RFC3339, self.ReleaseDate)
	app.CrawledTime = time.Now()

	sort.Strings(app.Devices)
	sort.Strings(app.Languages)

	return app
}

func (self *Entry) Detail(country string) (app *App) {
	app = self.ToApp()
	app.ParseExtras(country)
	return
}

/**************************************************************
* ParseExtras:	Fill APP's missing fields
**************************************************************/

// ParseExtras will fill reserved fields by fetching & parsing iTunes Store
// It's good but not necessary. By default using chinese store
func (app *App) ParseExtras(country string) error {
	if app == nil || app.ID == 0 {
		return ErrParseFailed
	}
	if country == "" {
		country = "cn"
	}

	url := fmt.Sprintf("https://itunes.apple.com/%s/app/id%d",
		strings.ToLower(country),
		app.ID,
	)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	// quick selector
	left := doc.Find(`#left-stack`)
	mid := doc.Find(`div.center-stack`)
	ul := left.Find(`ul.list:first-of-type`)

	// app.Screenshots
	app.Screenshots = merge(app.Screenshots,
		getAttrList(mid.Find(`img[itemprop="screenshot"]`), "src"))

	// app.Copyright
	app.Copyright = getText(ul.Find("li.copyright"))

	// app.Platforms
	platform := make(map[string]bool, 6)

	// infer iPhone,iPad,iPod support from Device
	deviceList := strings.Join(app.Devices, ",")
	platform["iPad"] = strings.Contains(deviceList, "iPad")
	platform["iPhone"] = strings.Contains(deviceList, "iPhone")
	platform["iPod"] = strings.Contains(deviceList, "iPod")
	platform["macOS"] = strings.Contains(app.System, "macOS") || strings.Contains(app.System, "OS X")
	platform["AppleTV"] = strings.Contains(app.System, "tvOS")

	// infer iPhone, iPad support from left stack badge
	if PadnPhone := getText(left.Find("div.fat-binary-blurb span:last-of-type")); PadnPhone != "" {
		if strings.Contains(PadnPhone, "iPhone") {
			platform["iPhone"] = true
		}
		if strings.Contains(PadnPhone, "iPad") {
			platform["iPad"] = true
		}
	}

	// infer iWatch supoort from iWatch badge or label
	if t := getText(left.Find("div.works-on-apple-watch span:last-of-type")); t != "" {
		platform["iWatch"] = true
	} else if len(left.Find("span.works-on-apple-watch-badge").Nodes) > 0 {
		platform["iWatch"] = true
	}

	// infer AppleTV from left stack label
	if !platform["AppleTv"] {
		left.Find("div.application span.label").Map(func(ind int, s *goquery.Selection) string {
			if s.Text() == "Apple TV: " {
				platform["AppleTV"] = true
			}
			return ""
		})
	}

	// infer iMessage support from left stack label
	if len(left.Find("offers-i-message-app-badge").Nodes) > 0 {
		platform["iMessage"] = true
	}

	// sort platforms
	var platformList []string
	for k, ok := range platform {
		if ok {
			platformList = append(platformList, k)
		}
	}
	sort.Strings(platformList)
	app.Platforms = platformList

	// app.InAppPurchase
	app.InAppPurchase = left.Find("div.in-app-purchases ol.list li").Map(func(ind int, s *goquery.Selection) string {
		itemTitle := s.Find("span.in-app-title").Text()
		itemPrice := s.Find("span.in-app-price").Text()
		return fmt.Sprintf("%d:%s:%s", ind+1, itemPrice, itemTitle)
	})

	// app.SiblingApps
	app.SiblingApps = stringSliceToInt(getAttrList(left.Find("div.more-by > ul.list > li > div"), "adam-id"))

	// app.RelatedApps
	app.RelatedApps = stringSliceToInt(getAttrList(mid.Find("div.lockup.application.small"), "adam-id"))

	// app.SupportSites
	support := make(map[string]string, 0)
	mid.Find(`a.see-all[rel="nofollow"]`).Map(func(ind int, s *goquery.Selection) string {
		support[s.Text()], _ = s.Attr("href")
		return ""
	})
	if body, err := json.MarshalIndent(support, "", "    "); err == nil {
		if sb := string(body); sb != "" && sb != "null" {
			app.SupportSites = sb
		}
	}
	if app.SupportSites == "" {
		app.SupportSites = "{}"
	}

	// app.Reviews:	quad-tuple for `<user,rating,title,content>`
	// more detailed comment could be fetched from
	// https://itunes.apple.com/cn/rss/customerreviews/id=<appid>/sortBy=mostRecent/json
	var reviews [][4]string
	mid.Find("div.customer-review").Map(func(ind int, s *goquery.Selection) string {
		title := getText(s.Find("span.customerReviewTitle"))
		rating := getAttr(s.Find("div.rating"), "aria-label")
		user := getText(s.Find("span.user-info"))
		if piece := strings.Split(user, "\n"); len(piece) > 1 {
			user = strings.TrimSpace(piece[len(piece)-1])
		}
		content := getRichText(s.Find("p.content"))
		if user != "" {
			reviews = append(reviews, [4]string{user, rating, title, content})
		}
		return ""
	})
	if body, err := json.MarshalIndent(reviews, "", "    "); err == nil {
		if sb := string(body); sb != "" && sb != "null" {
			app.Reviews = sb
		}
	}
	if app.Reviews == "" {
		app.Reviews = "[]"
	}

	return nil
}

/**************************************************************
* Auxiliary Functions
**************************************************************/
// stringSliceToInt will transform []string to []bigint
func stringSliceToInt(s []string) []int64 {
	var ilist []int64
	for _, str := range s {
		if num, err := strconv.ParseInt(str, 10, 64); num != 0 && err == nil {
			ilist = append(ilist, num)
		}
	}
	return ilist
}

// getText will extract text from selector and trim space
func getText(selection *goquery.Selection) (s string) {
	return strings.TrimSpace(selection.Text())
}

// getAttr will extract attr according attrName from selector and trim space
func getAttr(selection *goquery.Selection, attrName string) (s string) {
	s, _ = selection.Attr(attrName)
	return strings.TrimSpace(s)
}

// getRichText handles multiline text
func getRichText(selection *goquery.Selection) (s string) {
	if s, err := selection.Html(); s != "" && err == nil {
		s = strings.Replace(s, "<br>", "\n", -1)
		s = strings.Replace(s, "<br/>", "\n", -1)
		s = strings.TrimSpace(s)
		return s
	}
	return
}

// getAttrList will fetch a list of attr of selectors
func getAttrList(selection *goquery.Selection, attrName string) []string {
	res := selection.Map(func(ind int, s *goquery.Selection) string {
		attr, _ := s.Attr(attrName)
		return attr
	})
	return removeEmpty(res)
}

// removeEmpty remove empty string from a string slice
func removeEmpty(input []string) (output []string) {
	for _, str := range input {
		if str != "" {
			output = append(output, str)
		}
	}
	return
}

// merge will merge two string slice & dedupe it
func merge(source ... []string) ([]string) {
	m := make(map[string]struct{}, len(source)*10)
	for _, list := range source {
		for _, item := range list {
			m[item] = struct{}{}
		}
	}
	dst := make([]string, len(m))
	cnt := 0
	for k, _ := range m {
		dst[cnt] = k
		cnt += 1
	}
	sort.Strings(dst)
	return dst
}
