package app

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
import . "github.com/Vonng/go-itunes-search"

/**************************************************************
* struct:	App
**************************************************************/

// App represent iTunes application entity
// Some fields, like Platforms, InAppPurchase SiblingApps RelatedApps SupportSite & Reviews
// could only be fetched from iTunes page. a parser adjust for CN Store is provided
type App struct {
	ID               int64        `sql:",pk"`
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
	Copyright        string
	GenreID          int64
	GenreIDList      []int64    `pg:",array"`
	Genre            string
	GenreList        []string   `pg:",array"`
	Icon60           string
	Icon100          string
	Price            int64      `sql:",notnull"`
	Currency         string
	System           string
	Features         []string   `pg:",array"`
	Devices          []string   `pg:",array"`
	Languages        []string   `pg:",array"`
	Platforms        []string   `pg:",array"`
	Rating           string
	Reasons          []string   `pg:",array"`
	Size             int64
	CntRating        int64
	AvgRating        float64
	CntRatingCurrent int64
	AvgRatingCurrent float64
	VppDevice        bool
	GameCenter       bool
	Screenshots      []string   `pg:",array"`
	InAppPurchase    []string   `pg:",array"`
	SiblingApps      []int64    `pg:",array"`
	RelatedApps      []int64    `pg:",array"`
	SupportSites     string
	Reviews          string
	Description      string
	ReleaseNotes     string
	Extra            string
	ReleaseTime      time.Time
	PublishTime      time.Time
	CrawledTime      time.Time
	tableName        struct{}    `sql:"apple"`
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
┃	{{.VendorName}} {{.Copyright}}
┃	{{.VendorURL}}
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	{{.GenreID}} {{.GenreIDList}}
┃	{{.Genre}} {{.GenreList}}
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
* constructor:	NewApp(*Entry) *App
* 	Create App from Entry
**************************************************************/

// NewApp create app from entry and guaranteed to success
func NewApp(entry *Entry) (app *App) {
	app = new(App)
	app.ID = entry.TrackID
	app.Name = entry.TrackName
	app.URL = entry.TrackViewURL

	app.Icon = entry.ArtworkURL512
	app.Kind = entry.Kind
	app.Version = entry.Version
	app.BundleID = entry.BundleID

	app.AuthorID = entry.ArtistID
	app.AuthorName = entry.ArtistName
	app.AuthorURL = entry.ArtistViewURL
	app.VendorName = entry.SellerName
	app.VendorURL = entry.SellerURL

	app.GenreID = entry.PrimaryGenreID
	app.GenreIDList = stringSliceToInt(entry.GenreIDs)
	app.Genre = entry.PrimaryGenreName
	app.GenreList = entry.Genres

	app.Icon60 = entry.ArtworkURL60
	app.Icon100 = entry.ArtworkURL100
	app.Price = int64(math.Ceil(entry.Price))
	app.Currency = entry.Currency

	app.System = entry.MinimumOsVersion
	app.Features = entry.Features
	app.Devices = entry.SupportedDevices
	app.Languages = entry.LanguageCodesISO2A

	app.Rating = entry.TrackContentRating
	if app.Rating == "" {
		app.Rating = entry.ContentAdvisoryRating
	}
	app.Reasons = entry.Advisories

	app.Size, _ = strconv.ParseInt(entry.FileSizeBytes, 10, 64)
	app.CntRating = entry.UserRatingCount
	app.AvgRating = entry.AverageUserRating
	app.CntRatingCurrent = entry.UserRatingCountForCurrentVersion
	app.AvgRatingCurrent = entry.AverageUserRatingForCurrentVersion
	app.VppDevice = entry.IsVppDeviceBasedLicensingEnabled
	app.GameCenter = entry.IsGameCenterEnabled

	app.Screenshots = merge(entry.ScreenshotURLs, entry.AppletvScreenshotURLs, entry.IpadScreenshotURLs)

	// Reserved fields: these fields should be fetched from iTunes page
	// app.Copyright
	// app.Platforms
	// app.InAppPurchase
	// app.SiblingApps
	// app.RelatedApps
	// app.SupportSites
	// app.Reviews

	app.Description = entry.Description
	app.ReleaseNotes = entry.ReleaseNotes
	app.ReleaseTime, _ = time.Parse(time.RFC3339, entry.CurrentVersionReleaseDate)
	app.PublishTime, _ = time.Parse(time.RFC3339, entry.ReleaseDate)
	app.CrawledTime = time.Now()

	sort.Strings(app.Devices)
	sort.Strings(app.Languages)

	return app
}

// NewDetailedApp will parse extra info while omit error
func NewDetailedApp(entry *Entry, country string) (app *App) {
	app = NewApp(entry)
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
