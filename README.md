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

Simple & handy tools for iTunes **App** Search & Lookup, source code : [`bin/itunes.go`](bin/itunes.go) 

### Usage:

```
Usage of itunes:
  -b string
    	bundleID for lookup eg:com.tencent.xin
  -c string
    	restrict to country. default:CN (default "CN")
  -d	fetch extra details. default:disabled
  -i string
    	id for lookup. eg:414478124
  -l	show result in list format
  -n int
    	number of result size. 1~200,default:50 (default 10)
  -s string
    	searching keyword eg:HelloWorld
```

### Options

- Available language code could refer to [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) , Specify by `-c`


-  `-d` options provides extra info, requires an extra RoundTrip to corresponding iTunes Store.
- Actually you can fetch basic entry & extra info from two different country.
- `-n` will limit result size, default 50, range from 1 to 200. but fixed to 1 when provide `id` or `bundleID`
- `-l` will tabulate result with `iTunesID, BundleID, AppName, Version` only, only available on search
- `-s` provides searching keywords, you may refer  [iTunes Search API Document](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/)  for more detail
- `-i` means lookup iTunes entity by `iTunesID`, which is a serial like `414478124`。
- `-d` means lookup iTunes entity by `BundleID`, which looks like a domain name: `com.MaoTian.XiXiGame`

### Example

#### Search keywords

```
# search keyword `credit` in US Store, return 20 records, list format
itunes -s credit -c US -n 20 -l
```

```
------------------------------------------------------------------------------------------
iTunesID  |BundleID                                  | Name & Ver
------------------------------------------------------------------------------------------
519817714  com.creditkarma.mobile                     Credit Karma: Credit Scores, Reports & Alerts 4.12.2
370811491  org.navyfederal.nfcuforiphone              Navy Federal Credit Union 5.10
382617920  com.viber                                  Viber Messenger – Text & Call 7.3
476718980  com.creditsesame.mobile.ios.finance        Credit Sesame - Instant Credit Score & Alerts 3.1.0
298867247  com.chase                                  Chase Mobile® 2.623
284847138  com.bankofamerica.BofA                     Bank of America - Mobile Banking 7.4.12
324389392  com.intuit.GoPayment                       QuickBooks GoPayment: POS Credit Card Reader 8.4.0
407558537  com.capitalone.enterprisemobilebanking     Capital One Mobile 5.22.0
1008234539 com.capitalone.credittracker2              Capital One CreditWise - Credit score and report 1.5.0
1128712763 com.creditonebank.mobile                   Credit One Bank Mobile 1.12
711923939  com.squareup.cash                          Square Cash - Send and Receive Money 2.17.2
602710567  com.BillGuard                              Prosper Daily - Money Tracking, Free Credit Score 4.8.3
300238550  com.mint.internal                          Mint: Personal Finance, Budget, Bills & Money 5.13.1
1087101090 com.experian.experianapp                   Experian - Credit Report 1.8.4
335393788  com.squareup.square                        Square Point of Sale - POS System (Register) 4.69
570315854  com.jaredallen.repost                      Repost for Instagram 3.2.5
301724680  com.citigroup.citimobile                   Citi Mobile® 8.7.0
404066296  com.bancard.payanywhere                    PayAnywhere - Point of Sale 5.3.2
965030252  com.experian.app                           Credit Tracker – Members Only 1.4.7
338010821  com.discoverfinancial.mobile               Discover – Mobile Banking and Finance 8.7.0
------------------------------------------------------------------------------------------
```

#### Lookup App by iTunes Track ID (ID)

```
# lookup by iTunesID 414478124 in JAPANESE store ,show extra detail
itunes -i 414478124 -c JP -d
```

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ [software] 414478124 com.tencent.xin WeChat 6.5.13
┃ https://itunes.apple.com/jp/app/wechat/id414478124?mt=8&uo=4
┃ http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/512x512bb.jpg
┃ Price: 0 JPY
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Provider:
┃	614694882 WeChat  https://itunes.apple.com/jp/developer/wechat/id614694882?uo=4
┃	Tencent Technology (Shenzhen) Company Limited http://www.wechat.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6005 Social Networking
┃	[6005 6007] [ソーシャルネットワーキング 仕事効率化]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon60 : http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/60x60bb.jpg
┃ Icon100: http://is2.mzstatic.com/image/thumb/Purple128/v4/db/c6/be/dbc6beea-e606-8548-9017-d2cfa3e6c300/source/100x100bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃	System   : 8.0
┃	Features : [iosUniversal]
┃	Devices  : [iPad23G-iPad23G iPad2Wifi-iPad2Wifi iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74 iPadAir-iPadAir iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadAirCellular-iPadAirCellular iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadMini4G-iPadMini4G iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPadPro-iPadPro iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPadProCellular-iPadProCellular iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone4S-iPhone4S iPhone5-iPhone5 iPhone5c-iPhone5c iPhone5s-iPhone5s iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPhoneSE-iPhoneSE iPodTouchFifthGen-iPodTouchFifthGen iPodTouchSixthGen-iPodTouchSixthGen]
┃	Languages: [AR DE EN ES FR HE HI ID IT JA KO MS PL PT RU TH TR VI ZH ZH ZH]
┃	Platforms: [iPad iPhone iPod iWatch]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating : 12+
┃ Reasons: [まれ/軽度な性的表現またはヌード]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Size: 207158272 VppDevice: true GameCenter:false
┃ Rating(Current)  : 4 / 8
┃ Rating(Historic) : 4 / 5777
┃ Sibling Apps : []
┃ Related Apps : []
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Screenshots:
	http://a2.mzstatic.com/jp/r30/Purple117/v4/3d/bf/07/3dbf07b5-6708-a6ed-6296-896e937291b9/sc1024x768.jpeg
	http://a2.mzstatic.com/jp/r30/Purple117/v4/6a/95/2c/6a952c86-fb2f-57d8-17c7-317039b93c95/sc1024x768.jpeg
	http://a2.mzstatic.com/jp/r30/Purple117/v4/ae/a2/27/aea22749-0085-fb95-d1c8-7aa46a91aa9d/screen696x696.jpeg
	http://a2.mzstatic.com/jp/r30/Purple117/v4/dc/fc/fd/dcfcfd4a-7b93-acbf-836d-7850df1b26b2/screen390x390.jpeg
	http://a2.mzstatic.com/jp/r30/Purple117/v4/dd/a8/bf/dda8bf60-a523-a0d8-3d44-50d1deb126b6/screen696x696.jpeg
	http://a2.mzstatic.com/jp/r30/Purple127/v4/3b/f2/bb/3bf2bb21-47cc-12b0-4c09-2abcabc1975f/screen390x390.jpeg
	http://a3.mzstatic.com/jp/r30/Purple127/v4/1d/f0/6f/1df06fcf-b1b3-7654-0d6c-11c2b55cc349/screen696x696.jpeg
	http://a3.mzstatic.com/jp/r30/Purple127/v4/ee/5b/c8/ee5bc868-d213-b4c1-3fae-b15a61e1bb91/screen696x696.jpeg
	http://a4.mzstatic.com/jp/r30/Purple117/v4/3b/8f/ac/3b8facee-4f81-f04d-622d-71a10f1e18e7/sc1024x768.jpeg
	http://a4.mzstatic.com/jp/r30/Purple117/v4/46/e1/89/46e18979-d354-4579-50c6-70ef576ad26d/sc1024x768.jpeg
	http://a4.mzstatic.com/jp/r30/Purple127/v4/3b/dc/34/3bdc3426-2612-a89a-e2d3-2925e2e8e212/screen696x696.jpeg
	http://a4.mzstatic.com/jp/r30/Purple127/v4/4d/c9/2e/4dc92ec5-eac5-058d-d68c-760472ead004/screen390x390.jpeg
	http://a5.mzstatic.com/jp/r30/Purple117/v4/96/e4/64/96e464fe-b7bb-ef72-0098-d05b282a7646/screen390x390.jpeg
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Purchase:
	1:¥120:Doraemon
	2:¥120:Cat's Melody
	3:¥120:Molang
	4:¥120:Winnie the Pooh
	5:¥120:Onigiri&Friends
	6:¥120:SpongeBob
	7:¥120:DURURU
	8:¥120:Crayon Shin Chan
	9:¥120:Hello, Crayon
	10:¥120:Oni & Bugi
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Reviews: [
    [
        "じゃん・ピエール",
        "星 2 つ",
        "LINEにある「既読」表示をつけてほしい",
        "相手がメッセージ読んだかどうかわからなくて不便"
    ],
    [
        "常用登机人",
        "星 1 つ",
        "指纹解锁",
        "应该配有指纹解锁功能！"
    ]
]
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Description:
WeChatはメッセージと通話のアプリで、家族や友人と簡単にどこにいてもつながることができます。テキスト(SMS/MMS)、音声と動画呼び出し、モーメンツ、写真の共有とゲームのすべてがひとつになったコミュニケーションアプリです。

WECHATを使う利点:
• マルチメディアメッセージ: 動画、画像、テキストそれから音声メッセージを送信できます。
• グループチャットと呼び出し: 最大500人までのグループが作成でき、グループ動画呼び出しには最大9人が参加できます。
• 無料音声&動画呼び出し: 高画質で世界中のどこにいても無料の呼び出しができます。
• WECHAT OUT通話: 世界中の固定電話や携帯電話に低料金で通話ができます(一部の地域のみ)。
• ステッカーギャラリー: 数百もの無料で、楽しい、動くステッカーのお気に入りのアニメや映画で自分自身を表現できます。
• モーメンツ: 個人のフォトストリームで、最高のモーメンツを共有しましょう。
• よりよい個人情報保護: WeChatは最高レベルのプライバシーを提供します。 TRUSTeから認定された唯一のメッセージアプリです。
• 新しい友人に会う: 「友人を探す」、「近くにいる人」、「シェイク」を利用して新しい友人に会いましょう。
• リアルタイムの場所: あなたがどこにいるのかを伝える代わりに、リアルタイムの場所の共有をするだけで済みます。
• 言語サポート: 20の言語に翻訳されています。他の言語へメッセージの翻訳もできます。
• WeRun-WeChat: 「WeRun-WeChat」公式アカウントを通じてHealthKitデータと友人とスコアを競うチャレンジにアクセス
• さらに: デスクトップアプリ、カスタム壁紙、カスタム通知、公式アカウントも使用できます。"
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Support Sites:
{
    "WeChat Web サイト": "http://www.wechat.com",
    "WeChat のサポート": "http://www.wechat.com"
}
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  ReleaseNotes:
最新のアップデート:
- 送信前に動画を編集
- モーメンツ投稿へのいいねやコメントのアラートを無効にする

最近のアップデート:
- ファイル、チャット履歴の中の写真やリンクの検索
- グループ所有者向けにグループメンバの参加方法の情報
- 友人に送信する前に、選択した写真のプレビューと編集が可能
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Release: 2017-08-07 06:33:21 +0000 UTC
┃ Publish: 2011-01-21 01:32:15 +0000 UTC
┃ Crawled: 2017-08-08 16:38:40.823925746 +0800 CST m=+0.493901121
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```



#### Lookup App by BundleID

```
# lookup app by bundleID: com.tencent.smoba in default Store(CN) without extra detail
itunes -b com.tencent.smoba
```

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ [software] 989673964 com.tencent.smoba 王者荣耀 1.20.1.21
┃ https://itunes.apple.com/cn/app/%E7%8E%8B%E8%80%85%E8%8D%A3%E8%80%80/id989673964?mt=8&uo=4
┃ http://is1.mzstatic.com/image/thumb/Purple128/v4/a5/bb/e5/a5bbe51f-5ba0-575d-de94-da6f65deb3a2/source/512x512bb.jpg
┃ Price: 0 CNY
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Provider:
┃	446324237 Tencent Mobile Games  https://itunes.apple.com/cn/developer/tencent-mobile-games/id446324237?uo=4
┃	Shenzhen Tencent Computer Systems Company Limited
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6014 Games
┃	[6014 7001 7017] [游戏 动作游戏 策略游戏]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon60 : http://is1.mzstatic.com/image/thumb/Purple128/v4/a5/bb/e5/a5bbe51f-5ba0-575d-de94-da6f65deb3a2/source/60x60bb.jpg
┃ Icon100: http://is1.mzstatic.com/image/thumb/Purple128/v4/a5/bb/e5/a5bbe51f-5ba0-575d-de94-da6f65deb3a2/source/100x100bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃	System   : 7.0
┃	Features : [iosUniversal]
┃	Devices  : [iPad23G-iPad23G iPad2Wifi-iPad2Wifi iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74 iPadAir-iPadAir iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadAirCellular-iPadAirCellular iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadMini4G-iPadMini4G iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPadPro-iPadPro iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPadProCellular-iPadProCellular iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone4-iPhone4 iPhone4S-iPhone4S iPhone5-iPhone5 iPhone5c-iPhone5c iPhone5s-iPhone5s iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPhoneSE-iPhoneSE iPodTouchFifthGen-iPodTouchFifthGen iPodTouchSixthGen-iPodTouchSixthGen]
┃	Languages: [ZH]
┃	Platforms: []
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating : 17+
┃ Reasons: [频繁/强烈的色情内容或裸露 偶尔/轻微的成人/性暗示题材 偶尔/轻微的卡通或幻想暴力 偶尔/轻微的现实暴力]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Size: 844592128 VppDevice: true GameCenter:false
┃ Rating(Current)  : 4.5 / 53714
┃ Rating(Historic) : 4.5 / 1217125
┃ Sibling Apps : []
┃ Related Apps : []
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Screenshots:
	http://a2.mzstatic.com/us/r30/Purple117/v4/46/cf/28/46cf282f-e711-560e-bbb4-0162321635aa/sc1024x768.jpeg
	http://a2.mzstatic.com/us/r30/Purple117/v4/ce/c6/f4/cec6f474-9d2d-eff5-106f-69189e86fd92/screen696x696.jpeg
	http://a2.mzstatic.com/us/r30/Purple117/v4/f5/9d/d4/f59dd48d-5ec1-afcd-8350-fef52a09adb4/sc1024x768.jpeg
	http://a3.mzstatic.com/us/r30/Purple117/v4/39/0e/af/390eaf8b-3a8a-b973-b5c0-fd9c26f1004a/sc1024x768.jpeg
	http://a3.mzstatic.com/us/r30/Purple127/v4/22/58/83/2258835d-e229-b560-2be8-bdeaa62f7798/screen696x696.jpeg
	http://a3.mzstatic.com/us/r30/Purple127/v4/5c/4e/77/5c4e7732-5aba-aaa9-f69b-90445304bc6d/screen696x696.jpeg
	http://a3.mzstatic.com/us/r30/Purple127/v4/8d/ad/f6/8dadf6b8-dc8b-db6b-894b-18fb404825a0/screen696x696.jpeg
	http://a4.mzstatic.com/us/r30/Purple117/v4/79/b5/ca/79b5ca24-4325-e689-5ff4-739bbdc7b77b/sc1024x768.jpeg
	http://a4.mzstatic.com/us/r30/Purple117/v4/c3/a2/22/c3a22263-fd6e-f6f9-9b68-bb2471ac5e4f/sc1024x768.jpeg
	http://a5.mzstatic.com/us/r30/Purple127/v4/cd/7c/f5/cd7cf587-a786-c34d-24f0-1c75bf021551/screen696x696.jpeg
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Purchase:
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Reviews:
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Description:
王者荣耀：无处不团，两亿好友聚好玩



【游戏介绍】
《王者荣耀》是腾讯第一5V5英雄公平对战手游，腾讯最新MOBA手游大作！5V5王者峡谷、5V5深渊大乱斗、以及3V3、1V1等多样模式一键体验，热血竞技尽享快感！海量英雄随心选择，精妙配合默契作战！10秒实时跨区匹配，与好友组队登顶最强王者！操作简单易上手，一血、五杀、超神，极致还原经典体验！实力操作公平对战，回归MOBA初心！
赶快加入《王者荣耀》，随时开启你的激情团战！



【游戏特色】
1、5V5！越塔强杀！超神！
5V5经典地图，三路推塔，呈现原汁原味的对战体验。英雄策略搭配，组建最强阵容，默契配合极限666！



2、深渊大乱斗！随机英雄一路团战！
5V5大乱斗，即刻激情团战！随机盲选英雄，全团杀中路，冲突一触即发！一条路，全神装，血战到底！



3、随时开团！10分钟爽一把！
适合手机的MOBA游戏，10分钟享受极致竞技体验。迂回作战，手脑配合，一战到底！人多，速来！



4、公平竞技！好玩不坑拼实力！
凭实力carry全场，靠技术决定胜负。不做英雄养成，不设体力，还你最初的游戏乐趣！



5、指尖超神！精妙走位秀操作！
微操改变战局！手速流？意识流？看我精妙走位，力压群雄，打出钻石操作！收割，连杀超神！



【特别说明】
在游戏《王者荣耀》中，用户登录时可以选择“与QQ好友玩/与微信好友玩/游客登录”，三种登录方式在iOS设备上的游戏数据不互通（包括等级、钻石、金币等）。用户在游戏中购买的游戏代币“点券”仅限在本应用中使用。腾讯的虚拟货币，比如Q币、Q点无法在本应用中使用。



【联系我们】
如果您喜欢我们的游戏，欢迎随时给我们评价、留言。
官方网站：http://pvp.qq.com
官方微信：heromoba
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  Support Sites:

┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
  ReleaseNotes:
【全新内容】
（1）万物有灵-鬼谷子：战国争鸣系列收官！人生导师鬼谷子将作为辅助登陆王者峡谷，隐身技能玩法再次升级，他的大招能让队友集体隐身！迷之可怕！
（2）S8赛季开启：排位赛赛季皮肤将开启红色系列，只要在当前赛季的排位赛对战中获胜10场或以上，就能立刻获得赛季皮肤奖励！全新段位至尊星耀加入，该段位会存在于钻石段位与王者段位之间。
（3）新玩法-无限乱斗：娱乐模式开启暑期狂潮！每隔两分钟，地图上会刷新不同的地图BUFF（如所有英雄输出增加30%等），更爽快更激烈的战斗模式，约起来！

【更多优化】
（1）社交系统优化，包括选将界面显示英雄名称，好友备注名显示等。
（2）战场体验优化，包括单人训练关优化、地图血量提示功能等。
（3）其他优化，包括小秘书成长历程优化、赠送皮肤体验优化等。

【修复内容】
优化了部分机型卡顿和闪退问题。
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Release: 2017-07-21 03:22:42 +0000 UTC
┃ Publish: 2015-10-28 03:44:09 +0000 UTC
┃ Crawled: 2017-08-08 16:39:57.117723845 +0800 CST m=+0.344626643
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```




## Package Usage

package provides itunes-search-api wrapper of Golang.

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



## Entry

`Entry` is naive mapping to iTunes's response structure.  While iTunes orgnaize everything just like music track, for those user who only interested in Application rather than music. Structure  `App` provides a more precise representation of iOS applications.

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



## App

`App` is an enhanced version of `Entity`, you may choose it instead of using Entry directly.

It provides some additinoal features:

- more friendly & shorter field names.
- fetch extra fields like `InAppPurchase`,`Reviews`,`RelatedApps`,`SupportedSites`, etc...
- Parser to fetch extra fields from corresponding country's store. 
- ORM Mapping to RDS table. 
- Pretty print template

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



### Entry2App

Transform from `Entry` to `App` is simple just using `entry.ToApp()`，it will always success with guarantee.

```go
entry, _ := Lookup().ID(414478124).Result()
app := entry.ToApp()
```

### Parsing Extra App Info


Parsing extra reserved field is simple as :

```go
app.ParseExtras(US)
```

Following fields may change during parsing, and won't change if parse failed

* `Copyright` is fetch from left stack on iTunes page.
* `Screenshots` from iTunes page will merge & dedupe with API's results. and do not keep difference of screenshots type: iPad, iPhone, iMessage, AppleTV, etc...

- `Platforms` is infered from badge, label, screenshots, device list, etc...
- `InAppPurchase` shows selling item represented as a tri-tuple `<rank,price,title>`
- `SiblingApps` list iTunesID of apps provides by same developer shows in the app page
- `RelatedApps` list iTunesID of apps recommend by apple in this app's page.
- `SupportSites` show at bottom of description. k-v json object with title as key, url as value.
- `Reviews` is a quad-tuple represent customer comments: `<user,rating,title,content>`
- Remove `CensoredName` , `FormattedPrice` 
- Merge `TrackContentRating` & `ContentAdvisoryRating` to `Rating`





## License

WTFPL