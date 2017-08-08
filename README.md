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
		Country(US).
		App().
		Limit(5).
		Results()

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
┃ Price: 0 USD Free
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃	Languages: | [EN DE PT ES]
┃	Features:  | [iosUniversal]
┃	Devices:   | [iPad2Wifi-iPad2Wifi iPad23G-iPad23G iPhone4S-iPhone4S iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone5-iPhone5 iPodTouchFifthGen-iPodTouchFifthGen iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini4G-iPadMini4G iPhone5c-iPhone5c iPhone5s-iPhone5s iPadAir-iPadAir iPadAirCellular-iPadAirCellular iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPodTouchSixthGen-iPodTouchSixthGen iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadPro-iPadPro iPadProCellular-iPadProCellular iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPhoneSE-iPhoneSE iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74]
┃	System:    | 9.0
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating:
┃ TrackContentRating:    | 4+
┃ ContentAdvisoryRating: | 4+
┃ RatingReason:          | []
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Ranking:
┃	Current:  | 4.5/51772
┃	Historic: | 5/3157
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Screenshots:
┃ 	Urls: [http://a3.mzstatic.com/us/r30/Purple71/v4/f3/ae/c3/f3aec3f4-904d-1278-3082-93d22cd69928/screen696x696.jpeg http://a1.mzstatic.com/us/r30/Purple62/v4/c9/40/fe/c940fefb-f43e-b18d-d3b0-6d1dcd4e436f/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple62/v4/d9/a5/8d/d9a58d76-f488-ac10-13a3-8a25e9026a1a/screen696x696.jpeg http://a3.mzstatic.com/us/r30/Purple71/v4/1d/a9/e7/1da9e785-1429-a59e-cb16-b18683fd8d33/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple62/v4/ac/1b/5f/ac1b5f5e-5462-d7e8-6c97-a9e46251174e/screen696x696.jpeg]
┃ 	Ipad: [http://a5.mzstatic.com/us/r30/Purple71/v4/46/8f/aa/468faa35-47ea-81cf-f96c-1c34d63f7ed8/sc1024x768.jpeg http://a4.mzstatic.com/us/r30/Purple62/v4/b4/6e/bd/b46ebddc-dd3f-b0b1-e78e-85e1b03c8beb/sc1024x768.jpeg http://a3.mzstatic.com/us/r30/Purple62/v4/98/aa/f9/98aaf94e-f0e7-7000-4110-26e775070470/sc1024x768.jpeg http://a4.mzstatic.com/us/r30/Purple71/v4/0f/32/e1/0f32e1c0-574c-9624-6d39-8ec40f66c0ab/sc1024x768.jpeg http://a1.mzstatic.com/us/r30/Purple71/v4/0e/60/a3/0e60a3f4-0d74-1f67-523a-f5bfc65d3c9c/sc1024x768.jpeg]
┃ 	TV:   []
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ GameCenter Enabled | false
┃ VppDevice Enabled  | true
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Size    | 124377088
┃ Version | 6.2.7
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

### other

other small examples

```go
SearchOne("世界").Country(CN).Entities([]string{Movie,Music}).Results()
SearchOne("Love").Media(Music).Entity(MusicArtist).Limit(5).Results()
Lookup().ID(529479190).Result()
Lookup().Country(GB).ID(529479190).Result()
Lookup().BundleID("com.supercell.magic").Result()
```

check [`api_test.go`](api_test.go) for more details & examples.




## App specific API

for those who only interested in Application rather than music, video, blahblah,... 
It may be a better idea using App specific API

```go
res, _ := Lookup().ID(414478124).Result()
app := res.ToApp()
```

Entry.ToApp() will transform entry to struct `App`. which have: 
* more friendly & short field names
* reserved fields for extra info (can be fetch from iTunes Store Page rather than API)
* parsing method for extra fields
* New print template



Parsing extra reserved field is as simple as :

```go
app.ParseExtras()
app.Print()
```

It will fetch & parse extra fields from iTunes Store. currently I'm using CN AppStore by default

* `Screenshots` from iTunes page will merge & dedupe with API's results.

- `Platforms` is infered from badge, label, screenshots, device list, etc...
- `InAppPurchase` shows selling item represented as a tri-tuple `<rank,price,title>`
- `SiblingApps` list iTunesID of apps provides by same developer shows in the app page
- `RelatedApps` list iTunesID of apps recommend by apple in this app's page.
- `SupportSites` show at bottom of description. k-v json object with title as key, url as value.
- `Reviews` is a quad-tuple represent customer comments: `<user,rating,title,content>`

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ [software] 414478124 com.tencent.xin 微信 6.5.13
┃ https://itunes.apple.com/cn/app/%E5%BE%AE%E4%BF%A1/id414478124?mt=8&uo=4
┃ http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/512x512bb.jpg
┃ Price: 0 CNY
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Provider:
┃       614694882 WeChat  https://itunes.apple.com/cn/developer/wechat/id614694882?uo=4
┃       Tencent Technology (Shenzhen) Company Limited http://weixin.qq.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃       6005 Social Networking
┃       [6005 6007] [社交 效率]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon60 : http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/60x60bb.jpg
┃ Icon100: http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/100x100bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃       System   : 8.0
┃       Features : [iosUniversal]
┃       Devices  : [iPad23G-iPad23G iPad2Wifi-iPad2Wifi iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74 iPadAir-iPadAir iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadAirCellular-iPadAirCellular iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadMini4G-iPadMini4G iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPadPro-iPadPro iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPadProCellular-iPadProCellular iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone4S-iPhone4S iPhone5-iPhone5 iPhone5c-iPhone5c iPhone5s-iPhone5s iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPhoneSE-iPhoneSE iPodTouchFifthGen-iPodTouchFifthGen iPodTouchSixthGen-iPodTouchSixthGen]
┃       Languages: [AR DE EN ES FR HE HI ID IT JA KO MS PL PT RU TH TR VI ZH ZH ZH]
┃       Platforms: [iPad iPhone iPod iWatch]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating : 12+
┃ Reasons: [偶尔/轻微的色情内容或裸露]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Size: 207158272 VppDevice: true GameCenter:false
┃ Rating(Current)  : 3.5 / 668
┃ Rating(Historic) : 4 / 812906
┃ Sibling Apps : []
┃ Related Apps : []
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Screenshots:
        http://a1.mzstatic.com/us/r30/Purple128/v4/e1/e4/e2/e1e4e2b0-db07-266b-7d81-b3b3486b473d/screen696x696.jpeg
        http://a2.mzstatic.com/us/r30/Purple118/v4/c6/38/4a/c6384adf-0f65-79a8-3862-791404cbef28/screen696x696.jpeg
        http://a2.mzstatic.com/us/r30/Purple128/v4/d8/56/af/d856afcd-e9cc-9936-d4f8-cd62e6d1c967/sc1024x768.jpeg
        http://a3.mzstatic.com/us/r30/Purple118/v4/90/54/7d/90547d63-3c8b-e788-9ca9-2528a7aa6ed1/screen390x390.jpeg
        http://a3.mzstatic.com/us/r30/Purple118/v4/c7/7c/27/c77c278b-385c-0208-158e-e7baf2c66031/screen696x696.jpeg
        http://a3.mzstatic.com/us/r30/Purple118/v4/e7/ec/75/e7ec75b9-f6da-5083-d923-27809ddd90c8/screen696x696.jpeg
        http://a3.mzstatic.com/us/r30/Purple118/v4/eb/bc/34/ebbc346f-033b-1b0b-0381-2f0d97709522/screen390x390.jpeg
        http://a3.mzstatic.com/us/r30/Purple128/v4/33/e2/5f/33e25f99-bca3-1747-db6b-378912b45f0c/sc1024x768.jpeg
        http://a4.mzstatic.com/us/r30/Purple118/v4/57/81/b9/5781b924-e6e0-02f2-1cbb-98afdd5e0c06/screen390x390.jpeg
        http://a4.mzstatic.com/us/r30/Purple118/v4/db/85/fc/db85fc17-dd73-b461-7430-33b501995aef/sc1024x768.jpeg
        http://a5.mzstatic.com/us/r30/Purple128/v4/02/9f/17/029f175c-7a83-5b09-9593-07deab91fd77/screen696x696.jpeg
        http://a5.mzstatic.com/us/r30/Purple128/v4/03/84/da/0384da46-9b4e-8258-b80d-e618d45d565a/screen390x390.jpeg
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Purchase:
        1:¥1.00:双拼乖巧
        2:¥6.00:邓超
        3:¥6.00:野原新之助
        4:¥6.00:小S
        5:¥6.00:Hello Kitty
        6:¥6.00:Angelababy
        7:¥6.00:哆啦A梦
        8:¥6.00:李光洙
        9:¥6.00:双重性格的喵小美
        10:¥6.00:甜甜私房猫
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Reviews: [
    [
        "mirrorz00726",
        "4星",
        "一些问题建议，望关注",
        "第一，版本更新后，还是没有看到不常联系人的模块。\n第二能否更新免打扰的微信群聊天，连红点点也没有，逼死强迫症了。\n第三，删除聊天记录有多选按钮，多选删除。\n第四，朋友圈能否分组查看。有时候代购的消息都要淹没自己的朋友了，但是代购有些讯息是需要的，并不想屏蔽。。"
    ],
    [
        "kirito011",
        "5星",
        "天天升级，就是不把错误改正",
        "腾讯果然厉害，根本不管你用户体验，二次删除这种化简为烦的东西，你不改早晚用户会因为你们的嚣张而离开"
    ],
    [
        "背后的故事、",
        "5星",
        "不稳定吗？",
        "更新之后用了一次恢复聊天记录进入安全模式，微信怎么都打不开，没办法，卸载重新登陆，聊天记录什么都没有了，心累"
    ]
]
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Description:
  微信是一款全方位的手机通讯应用，帮助你轻松连接全球好友。微信可以(通
过SMS/MMS网络)发送短信、进行视频聊天、与好友一起玩游戏，以及分享自己的
生活到朋友圈，让你感受耳目一新的移动生活方式。

  为什么要使用微信：
  • 多媒体消息：支持发送视频、图片、文本和语音消息。
  • 群聊和通话：组建高达500人的群聊和高达9人的实时视频聊天。
  • 免费语音和视频聊天：提供全球免费的高质量通话。
  • WeChat Out：超低费率拨打全球的手机或固定电话（目前仅限于部分地区）。
  • 表情商店：海量免费动态表情，包括热门卡通人物和电影，让聊天变得更生动有趣。
  • 朋友圈：与好友分享每个精彩瞬间，记录自己的生活点滴。
  • 隐私保护：严格保护用户的隐私安全，是唯一一款通过TRUSTe认证的实时通讯应用。
  • 认识新朋友：通过“雷达加朋友”、“附近的人”和“摇一摇”认识新朋友。
  • 实时位置共享：与好友分享地理位置，无需通过语言告诉对方。
  • 多语言：支持超过20种语言界面，并支持多国语言的消息翻译。
  · 微信运动，支持接入Apple Watch 及iPhone健康数据，可通过“WeRun-WeChat”公众号与好友一较高下。
  • 更多功能: 支持跨平台、聊天室墙纸自定义、消息提醒自定义和公众号服务等。
  Support Sites:
  {
    "WeChat 网站": "http://weixin.qq.com",
    "微信 支持": "http://weixin.qq.com"
}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  ReleaseNotes:
  本次更新
- 群资料页可以查看最近收到的小程序。

最近更新
- 可以对视频进行编辑。
- 可以设置某条朋友圈的互动不再通知。
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Release: 2017-08-07 06:33:21 +0000 UTC
┃ Publish: 2011-01-21 01:32:15 +0000 UTC
┃ Crawled: 2017-08-08 14:00:19.844798884 +0800 CST m=+0.574360555
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

```



## Sample & Utils

[`bin/app_lookup.go`](bin/app_lookup.go) provides a sample: a binary tool that take iTunesID or BundleID as input,  print app's detail as output.

Usage:

```
Usage:
	./itunes [search] term1 term2 ...
	./itunes lookup <idType> <idValue>
	./itunes lookup <iTunesID>|<BundleID>

	eg: 414478124 com.tencent.xin WeChat
```





## License

WTFPL