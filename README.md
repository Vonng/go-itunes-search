# iTunes Search API for Golang

`go-itunes-search` is a golang wrapper for [iTunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/#lookup) 

## Reference

See [iTunes Search API Document](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/) for more details.



## Install

```bash
go get github.com/Vonng/go-itunes-search
```



## Usage

### import

```go
import . "github.com/Vonng/go-itunes-search"
```

### Search Example

search with keyword `Hello` & `World`，US AppStore，Restrict media type to `Software`，At more 5 result.

```go
func TestSearch(t *testing.T) {
	res, _ := Search([]string{"Hello", "World"}).
		Country(US).App().Limit(5).Results()

	for _, r := range res {
		r.Print()
	}
}
```

result is fetched via `.Results` or `.Result`，the latter assert only one result is returned so it returns `*Entry` rather than `[]Entry`。

### Lookup API

if you know something could used to identify a track, then lookup API may be a better approach.
Instead of specifying `term`, you need something like `iTunesID (track_id)`, `BundleID`(app only), `AMG ID`, etc…。And when using lookup API, there could only be one or zero entry being returned. So the API chain may end with `Result` rather than `Results`

Here's how it works, these lookups may return same results:

```go
Lookup().ID(414478124).Country(CN).Result()
Lookup().BundleID("com.tencent.xin").Result()
Lookup().Set(BundleID, "com.tencent.xin").Result()
```

### Entry

`Entry` contains many fields: 

```
type Entry struct {
	TrackID                            int64    `json:"trackId"` // Track
	TrackName                          string   `json:"trackName"`
	TrackCensoredName                  string   `json:"trackCensoredName"`
	TrackViewUrl                       string   `json:"trackViewUrl"`
	BundleID                           string   `json:"bundleId"` // App bundle
	ArtistID                           int64    `json:"artistId"` // Artist
	ArtistName                         string   `json:"artistName"`
	ArtistViewUrl                      string   `json:"artistViewUrl"`
	SellerName                         string   `json:"sellerName"` // Seller
	SellerUrl                          string   `json:"sellerUrl"`
	PrimaryGenreID                     int64    `json:"primaryGenreId"` // Genre
	PrimaryGenreName                   string   `json:"primaryGenreName"`
	Genres                             []string `json:"genres"`
	GenreIDs                           []string `json:"genreIds"`
	ArtworkUrl60                       string   `json:"artworkUrl60"` // Icon
	ArtworkUrl100                      string   `json:"artworkUrl100"`
	ArtworkUrl512                      string   `json:"artworkUrl512"`
	Price                              float64  `json:"price"` // Price
	Currency                           string   `json:"currency"`
	FormattedPrice                     string   `json:"formattedPrice"`
	LanguageCodesISO2A                 []string `json:"languageCodesISO2A"` // Platform
	Features                           []string `json:"features"`
	SupportedDevices                   []string `json:"supportedDevices"`
	MinimumOsVersion                   string   `json:"minimumOsVersion"`
	TrackContentRating                 string   `json:"trackContentRating"`
	ContentAdvisoryRating              string   `json:"contentAdvisoryRating"` // Rating
	Advisories                         []string `json:"advisories"`
	UserRatingCount                    int64    `json:"userRatingCount"` // Ranking
	AverageUserRating                  float64  `json:"averageUserRating"`
	UserRatingCountForCurrentVersion   int64    `json:"userRatingCountForCurrentVersion"`
	AverageUserRatingForCurrentVersion float64  `json:"averageUserRatingForCurrentVersion"`
	Kind                               string   `json:"kind"` // Type
	WrapperType                        string   `json:"wrapperType"`
	ScreenshotUrls                     []string `json:"screenshotUrls"` // Screenshots
	IpadScreenshotUrls                 []string `json:"ipadScreenshotUrls"`
	AppletvScreenshotUrls              []string `json:"appletvScreenshotUrls"`
	IsGameCenterEnabled                bool     `json:"isGameCenterEnabled"` // Flags
	IsVppDeviceBasedLicensingEnabled   bool     `json:"isVppDeviceBasedLicensingEnabled"`
	FileSizeBytes                      string   `json:"fileSizeBytes"` // Attribute
	Version                            string   `json:"version"`
	Description                        string   `json:"description"`
	ReleaseNotes                       string   `json:"releaseNotes"`
	ReleaseDate                        string   `json:"releaseDate"`
	CurrentVersionReleaseDate          string   `json:"currentVersionReleaseDate"`
}
```

 `Entry.Print` provides an util for pretty-print. formatted as :

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ iTunes Track [ software / software ]
┃	332615624 Bible - Daily Reading & Study Bible by Olive Tree (Bible - Daily Reading & Study Bible by Olive Tree)
┃	[com.olivetree.BR-Free]  https://itunes.apple.com/us/app/bible-daily-reading-study-bible-by-olive-tree/id332615624?mt=8&uo=4
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Artist:
┃	444535393 HarperCollins Christian Publishing, Inc.  https://itunes.apple.com/us/developer/harpercollins-christian-publishing-inc/id444535393?uo=4
┃	HarperCollins Christian Publishing, Inc. http://www.olivetree.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6006 Reference
┃	[6006 6018] [Reference Books]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon:
┃ 	60:	http://is4.mzstatic.com/image/thumb/Purple117/v4/62/1e/71/621e71b6-3238-914d-b219-37ebb9227411/source/60x60bb.jpg
┃ 	100:http://is4.mzstatic.com/image/thumb/Purple117/v4/62/1e/71/621e71b6-3238-914d-b219-37ebb9227411/source/100x100bb.jpg
┃ 	512:http://is4.mzstatic.com/image/thumb/Purple117/v4/62/1e/71/621e71b6-3238-914d-b219-37ebb9227411/source/512x512bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Price:	0 USD Free
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃ Languages:	[EN DE PT ES]
┃ Features:		[iosUniversal]
┃ Devices:		[iPad2Wifi-iPad2Wifi iPad23G-iPad23G iPhone4S-iPhone4S iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone5-iPhone5 iPodTouchFifthGen-iPodTouchFifthGen iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini4G-iPadMini4G iPhone5c-iPhone5c iPhone5s-iPhone5s iPadAir-iPadAir iPadAirCellular-iPadAirCellular iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPodTouchSixthGen-iPodTouchSixthGen iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadPro-iPadPro iPadProCellular-iPadProCellular iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPhoneSE-iPhoneSE iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74]
┃ SystemRequirement:	9.0
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating:
┃ TrackContentRating:		4+
┃ ContentAdvisoryRating: 	4+
┃ Reason:					[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Ranking:
┃	Current:	4.5	51772
┃	Historic:	5	3157
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Screenshots:
┃ Urls:	[http://a3.mzstatic.com/us/r30/Purple71/v4/f3/ae/c3/f3aec3f4-904d-1278-3082-93d22cd69928/screen696x696.jpeg http://a1.mzstatic.com/us/r30/Purple62/v4/c9/40/fe/c940fefb-f43e-b18d-d3b0-6d1dcd4e436f/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple62/v4/d9/a5/8d/d9a58d76-f488-ac10-13a3-8a25e9026a1a/screen696x696.jpeg http://a3.mzstatic.com/us/r30/Purple71/v4/1d/a9/e7/1da9e785-1429-a59e-cb16-b18683fd8d33/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple62/v4/ac/1b/5f/ac1b5f5e-5462-d7e8-6c97-a9e46251174e/screen696x696.jpeg]
┃ Ipad:	[http://a5.mzstatic.com/us/r30/Purple71/v4/46/8f/aa/468faa35-47ea-81cf-f96c-1c34d63f7ed8/sc1024x768.jpeg http://a4.mzstatic.com/us/r30/Purple62/v4/b4/6e/bd/b46ebddc-dd3f-b0b1-e78e-85e1b03c8beb/sc1024x768.jpeg http://a3.mzstatic.com/us/r30/Purple62/v4/98/aa/f9/98aaf94e-f0e7-7000-4110-26e775070470/sc1024x768.jpeg http://a4.mzstatic.com/us/r30/Purple71/v4/0f/32/e1/0f32e1c0-574c-9624-6d39-8ec40f66c0ab/sc1024x768.jpeg http://a1.mzstatic.com/us/r30/Purple71/v4/0e/60/a3/0e60a3f4-0d74-1f67-523a-f5bfc65d3c9c/sc1024x768.jpeg]
┃ TV:	[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ GameCenter Enabled              | false
┃ VppDeviceBasedLicensingEnabled  | true
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ FileSizeBytes                   | 124377088
┃ Version                         | 6.2.7
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Description:
Bible by Olive Tree is the best free Bible app for reading and studying God’s Holy Word and comes with great translations like the NIV, ESV, KJV, NKJV, and more!

Do more than just read your Bible - take notes, highlights, and save passages and sync to all your devices. This free Bible Study App features a powerful Resource Guide that links your Bible text with outstanding study Bibles, maps, commentaries, and more for an in-depth Bible study experience. Start a Bible reading plan and Bible will track your progress as you read through Scripture. Our unique split window allows you to create your own customized parallel Bible to easily compare Bible translations. Install now and explore over 100 more free titles with full offline functionality!

In addition to the New International Version (NIV), King James Version (KJV), English Standard Version (ESV), New King James Version bible that work offline, you can also download dozens of free study resources. Even more translations and great study resources are also available for purchase in-app.


OFFLINE BIBLE STUDY
Read and study whether you’re connected or not. Your library, notes, highlights and all of the app features are stored on your device so that you have full functionality when you are offline or in airplane mode.

CLOUD SYNC
Sync your Bible study resources, highlights, notes, save passages, and book ribbons between any devices with Bible.

POWERFUL RESOURCE GUIDE
With our one-of-a-kind Resource Guide, perform powerful searches through your entire Library of Bibles, Bible commentaries, Bible dictionaries, and more.

SIDE-BY-SIDE STUDY
The split window feature allows you to create your own customized parallel Bible for translation comparison, view your study notes while you read, or follow along with a commentary while you study Scripture.

IMMERSIVE BIBLE STUDY
• Remove distractions by opening your books and Bibles in full screen and immerse yourself in Scripture. 
• Night theme for easier reading in lowlight. 

PERSONAL BIBLE STUDY
• Highlight words and passages
• Take your own personalized notes
• Save your favorite passages 
• Tag anything to find it quickly later
• Leave a book ribbon on a page in order to pick up where you left off
• Select and copy text from any Bible or book in your Library

DAILY READING PLANS
• Free downloadable reading plans on various topics, books of the Bible, or specific biblical characters
• Sync your reading plan across your devices with Bible+
• Plans vary in length with options as short as 5 days to as long as three years!

SOCIAL BIBLE STUDY
Instantly share the Bible with your friends from inside the app. Tap on a verse to share it through Twitter, Facebook or via email.

OTHER BIBLE STUDY RESOURCES AVAILABLE FOR PURCHASE IN-APP:

• The Message, Amplified Bible, New American Standard Bible (NASB), New Living Translation (NLT), New Revised Standard Version (NRSV), and over 100 more!
• Best-selling study Bibles: ESV Study Bible, NLT Study Bible, NIV Study Notes, NKJV Study Notes, Life Application Study Bible, Reformation Study Bible Notes
• Word Study Bibles with Strong’s Numbers in NIV, KJV, ESV, NKJV, HCSB and NASB Bible translations
• Commentaries and Study Tools: Vine’s Expository Dictionary, Expositor’s Bible Commentary; Olive Tree Bible Maps, Bible Knowledge Commentary, Zondervan Atlas of the Bible
• Interlinear Bibles: Easily compare the Original Languages of the Bible with ESV, KJV, and NKJV Bible translations.
• Harmony of the Gospels: Read through the life of Jesus chronologically with our unique Gospel harmonies.
• Original language Bibles: Greek New Testament: NA28 & UBS-4; Hebrew Old Testament: BHS; Greek Old Testament: Septuaginta, LXX
• Non-English Bibles including Spanish, Portuguese, German and more: Reina-Valera, Almeida Revista e Atualizada, Dios Habla Hoy, Luther Bibel 1984, Louis Segond, Indonesian Bible

AND MANY MORE!
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
ReleaseNotes:
Thanks for using the Olive Tree Bible App!

In this release we added the ability to swipe to mark messages as read in our message center. We also fixed a number of bugs that were trying to get in the way of your reading and study. 

If you enjoy the app please leave us a review. It really means a lot!
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ ReleaseDate                     | 2009-10-02T22:47:42Z
┃ CurrentVersionReleaseDate       | 2017-07-10T16:31:08Z
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

check [`api_test.go`](api_test.go) for more details & examples.



## API

```
// public
Lookup()
Search()

// public method
.Add
.Set
.Get
.Del
.Encode
.Term
.Terms
.Country
.Entity
.Entites
.AddEntities
.Media
.AddMedia
.Limit
.ID
.BundleID
.CNAPP
.USAPP

// fetch Result
.Result     // fetch one (if there is one)
.Results    // fetch all
```

