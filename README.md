# iTunes Search API for Golang

`go-itunes-search` is a golang wrapper for [iTunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/#lookup) 

it also provides a naive binary tool for quick access



## Reference

See [iTunes Search API Document](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/) for more details.



## Install

### SDK:

```bash
go get github.com/Vonng/go-itunes-search
```

### Binary Util

```
cd $GOPATH/src/github.com/Vonng/go-itunes-search/bin
make install
```



## Binary Util Usage

Simple & handy tools, source code : [`bin/itunes.go`](bin/itunes.go) 

provides a sample: a binary tool that take iTunesID or BundleID as input,  print app's detail as output.

Usage:

```
Usage:
	./itunes [search] term1 term2 ...
	./itunes lookup <idType> <idValue>
	./itunes lookup <iTunesID>|<BundleID>

Example:
# search one keyword
itunes search 王者荣耀

# search multi keyword and omit keyword `search`
itunes Hello World Programer

# lookup for more detail from iTunes page (reviews,in app purchase, etc...)
itunes lookup 414478124
itunes lookup com.tencent.xin
```

Search may return multiple record while lookup will only produce one or zero result. result will print in human-readable format like following:

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

By default I'm using chinese iTunes Store for extra attributes, and restrict results media in software. you may change it in source code.





## API Usage

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

### Lookup Example

if you know something could used to identify a track, then lookup API may be a better approach.
Instead of specifying `term`, you need something like `iTunesID (track_id)`, `BundleID`(app only), `AMG ID`, etc…。And when using lookup API, there could only be one or zero entry being returned. So the API chain may end with `Result` rather than `Results`

Here's how it works, these lookups may return same results:

```go
Lookup().ID(414478124).Country(CN).Result()
Lookup().BundleID("com.tencent.xin").Result()
Lookup().Set(BundleID, "com.tencent.xin").Result()
```

## Other examle

other small examples

```go
SearchOne("世界").Country(CN).Entities([]string{Movie,Music}).Results()
SearchOne("Love").Media(Music).Entity(MusicArtist).Limit(5).Results()
Lookup().ID(529479190).Result()
Lookup().Country(GB).ID(529479190).Result()
Lookup().BundleID("com.supercell.magic").Result()
```

check [`api_test.go`](api_test.go) for more details & examples.



## Important Structure

### `Entry`

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

 `Entry.Print` print itself in human-friendly format.

`Entry` is naive mapping to iTunes's response structure.  While iTunes orgnaize everything just like music track, for those user who only interested in Application rather than music. Structure  `App` provides a more precise representation of iOS applications.

 

## App specific API

Transform `Entry` to `App` is easy,  guaranteed to success. 

```go
entry, _ := Lookup().ID(414478124).Result()
app := entry.ToApp()
```

`App` rename some fields of `Entity`, Add reserve some extra fields that can only be fetched from page rather than itunes Search API.

```go
type App struct {
	ID               int64  `sql:",pk"`
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
	Price            int64  `sql:",notnull"`
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
```

* more friendly & shorter field names

* reserved fields for extra info (can be fetch from iTunes Store Page rather than API)

* parsing method for extra fields

* New print template

* Mapping to database table.

  ​



Parsing extra reserved field is as simple as :

```go
app.ParseExtras()
```

Extra info will fill into app's structure, and a returning err won't affect infos that entry already have.

It can be write in short form:

```go
# Equvilent to app := entry.ToApp() + app.ParseExtras()
app := entry.Detail()
```

It will fetch & parse extra fields from iTunes Store. currently I'm using CN AppStore by default

* `Copyright` is fetch from left stack on iTunes page.
* `Screenshots` from iTunes page will merge & dedupe with API's results. and do not keep difference of screenshots type: iPad, iPhone, iMessage, AppleTV, etc...

- `Platforms` is infered from badge, label, screenshots, device list, etc...
- `InAppPurchase` shows selling item represented as a tri-tuple `<rank,price,title>`
- `SiblingApps` list iTunesID of apps provides by same developer shows in the app page
- `RelatedApps` list iTunesID of apps recommend by apple in this app's page.
- `SupportSites` show at bottom of description. k-v json object with title as key, url as value.
- `Reviews` is a quad-tuple represent customer comments: `<user,rating,title,content>`
- Remove `CensoredName` , `FormattedPrice` , `TrackContentRating`  because of redundant





## License

WTFPL