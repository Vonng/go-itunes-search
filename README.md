# iTunes Search API for Golang



## Reference

[iTunes Search API Document](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/)

## Install

```bash
go get github.com/Vonng/go-itunes-search
```

## Usage

```go
import . "github.com/Vonng/go-itunes-search"
```

### Search Example
```go
func TestSearch(t *testing.T) {
	res, err := Search().
  				Media(MediaSoftware).
  				Country(US).
  				Term("Hello").
  				Limit(2).
  				Results()
  
	if err != nil {
		t.Error(err)
	}

	for _, r := range res {
		r.Print()
	}
}
```

that will produce output like:

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ iTunes Track [ software / software ]
┃	508231856 Zello Walkie Talkie (Zello Walkie Talkie)
┃	[com.zello.client.main]  https://itunes.apple.com/us/app/zello-walkie-talkie/id508231856?mt=8&uo=4
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Artist:
┃	508231859 Zello  https://itunes.apple.com/us/developer/zello/id508231859?uo=4
┃	Zello Inc http://zello.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6005 Social Networking
┃	[6005 6000] [Social Networking Business]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon:
┃ 	60:	http://is2.mzstatic.com/image/thumb/Purple128/v4/15/01/2a/15012ae5-f5fb-8f57-43ff-b352599b3b07/source/60x60bb.jpg
┃ 	100:http://is2.mzstatic.com/image/thumb/Purple128/v4/15/01/2a/15012ae5-f5fb-8f57-43ff-b352599b3b07/source/100x100bb.jpg
┃ 	512:http://is2.mzstatic.com/image/thumb/Purple128/v4/15/01/2a/15012ae5-f5fb-8f57-43ff-b352599b3b07/source/512x512bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Price:	0 USD Free
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃ Languages:	[AR BG ZH HR NL EN FR HE ID IT JA NB PL PT RU ZH ES SV TR]
┃ Features:		[]
┃ Devices:		[iPad2Wifi-iPad2Wifi iPad23G-iPad23G iPhone4S-iPhone4S iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone5-iPhone5 iPodTouchFifthGen-iPodTouchFifthGen iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini4G-iPadMini4G iPhone5c-iPhone5c iPhone5s-iPhone5s iPadAir-iPadAir iPadAirCellular-iPadAirCellular iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPodTouchSixthGen-iPodTouchSixthGen iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadPro-iPadPro iPadProCellular-iPadProCellular iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPhoneSE-iPhoneSE iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74]
┃ SystemRequirement:	8.0
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating:
┃ TrackContentRating:		12+
┃ ContentAdvisoryRating: 	12+
┃ Reason:					[Infrequent/Mild Sexual Content and Nudity Infrequent/Mild Profanity or Crude Humor]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Ranking:
┃	Current:	4.5	3324
┃	Historic:	3	2
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Screenshots:
┃ Urls:	[http://a2.mzstatic.com/us/r30/Purple1/v4/bb/b3/76/bbb37650-043b-3cf6-967e-7d8730a4f7d6/screen696x696.jpeg http://a1.mzstatic.com/us/r30/Purple3/v4/c9/f0/5c/c9f05c30-6400-50af-c0a6-c3449f27e5e8/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple3/v4/bc/81/dd/bc81dd21-5c69-12b2-eb56-7c5f21c7cf7b/screen696x696.jpeg http://a1.mzstatic.com/us/r30/Purple1/v4/28/94/78/289478ad-6f8c-4f7b-f522-954885426c14/screen696x696.jpeg]
┃ Ipad:	[]
┃ TV:	[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ GameCenter Enabled              | false
┃ VppDeviceBasedLicensingEnabled  | true
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ FileSizeBytes                   | 39574528
┃ Version                         | 3.40
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Description:
Zello is the place for free live private and public conversations.

Join millions of people who use Zello instead of texting.  You can use it one-on-one with a friend, for a live group call with your family or soccer team.  The Zello app can even replace 2-way radios at work.

Zello is the only place for live open group communication – old school CB Radio style. Create a live Zello channel for your forum or customers, or enjoy conversations from across the globe.

+ Free live voice over any network or Wi-Fi connection
+ See who’s available or busy
+ Send photos to friends instantly
+ Replay messages later, even if your phone was off
+ Cross-platform
+ Free with no ads
+ Won’t spam your friends
+ Lets you delete your account
+ Supports Apple Watch as remote control (no audio support yet)
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
ReleaseNotes:
- Improved Send location function
- Faster sign in for accounts with large contact lists
- Bugfixes
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ ReleaseDate                     | 2012-03-29T20:21:28Z
┃ CurrentVersionReleaseDate       | 2017-07-31T18:27:33Z
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ iTunes Track [ software / software ]
┃	582654048 Sonic Dash (Sonic Dash)
┃	[com.sega.sonicdash]  https://itunes.apple.com/us/app/sonic-dash/id582654048?mt=8&uo=4
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Artist:
┃	281966698 SEGA  https://itunes.apple.com/us/developer/sega/id281966698?mt=8&uo=4
┃	Sega America http://www.sega.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6014 Games
┃	[6014 7001 7003] [Games Action Arcade]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon:
┃ 	60:	http://is4.mzstatic.com/image/thumb/Purple127/v4/63/17/22/631722ec-a856-d18e-2375-cdc6c328fdb4/source/60x60bb.jpg
┃ 	100:http://is4.mzstatic.com/image/thumb/Purple127/v4/63/17/22/631722ec-a856-d18e-2375-cdc6c328fdb4/source/100x100bb.jpg
┃ 	512:http://is4.mzstatic.com/image/thumb/Purple127/v4/63/17/22/631722ec-a856-d18e-2375-cdc6c328fdb4/source/512x512bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Price:	0 USD Free
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃ Languages:	[EN FR DE IT PT RU ES]
┃ Features:		[iosUniversal]
┃ Devices:		[iPad2Wifi-iPad2Wifi iPad23G-iPad23G iPhone4S-iPhone4S iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone5-iPhone5 iPodTouchFifthGen-iPodTouchFifthGen iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini4G-iPadMini4G iPhone5c-iPhone5c iPhone5s-iPhone5s iPadAir-iPadAir iPadAirCellular-iPadAirCellular iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPodTouchSixthGen-iPodTouchSixthGen iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadPro-iPadPro iPadProCellular-iPadProCellular iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPhoneSE-iPhoneSE iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74]
┃ SystemRequirement:	8.0
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating:
┃ TrackContentRating:		4+
┃ ContentAdvisoryRating: 	4+
┃ Reason:					[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Ranking:
┃	Current:	4.5	420269
┃	Historic:	4.5	1857
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Screenshots:
┃ Urls:	[http://a2.mzstatic.com/us/r30/Purple91/v4/5b/d5/5b/5bd55bb2-096d-7e11-a526-322e9940a014/screen696x696.jpeg http://a5.mzstatic.com/us/r30/Purple111/v4/6b/5b/13/6b5b1338-8d88-d9e7-269d-6ed589619381/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple122/v4/61/9b/df/619bdf77-0c13-3fe0-68c0-dba7b706ea68/screen696x696.jpeg http://a3.mzstatic.com/us/r30/Purple111/v4/3b/b0/88/3bb0880c-2c9a-c339-b5ea-55f66ea2d288/screen696x696.jpeg http://a3.mzstatic.com/us/r30/Purple122/v4/7b/87/67/7b876735-eb53-05c6-e5e0-03a0e2786ef6/screen696x696.jpeg]
┃ Ipad:	[http://a4.mzstatic.com/us/r30/Purple111/v4/48/2d/cb/482dcb48-d047-9ea2-9e29-c9c6eb7dccdb/sc1024x768.jpeg http://a2.mzstatic.com/us/r30/Purple111/v4/d5/f1/36/d5f136b1-57b7-93c9-985d-b16216f40c62/sc1024x768.jpeg http://a5.mzstatic.com/us/r30/Purple122/v4/ec/f0/20/ecf02097-95b9-8ae3-c545-011421b35c9a/sc1024x768.jpeg http://a2.mzstatic.com/us/r30/Purple111/v4/6f/fd/db/6ffddb4a-4180-df70-e88b-2fa5a8155383/sc1024x768.jpeg http://a3.mzstatic.com/us/r30/Purple111/v4/2e/f5/4f/2ef54f89-4df0-b26a-b37b-dd55f4a25ce3/sc1024x768.jpeg]
┃ TV:	[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ GameCenter Enabled              | false
┃ VppDeviceBasedLicensingEnabled  | true
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ FileSizeBytes                   | 179712000
┃ Version                         | 3.7.3
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Description:
"A perfect blend of Temple Run-style gaming and the lightspeed nature of Sonic. Amazing visuals. It feels as fast, fun and frantic as you've been wanting Sonic to for a long time." - 5/5 Gamezebo

4/5 - Common Sense Media 

How far can the world’s fastest hedgehog run?

Play as Sonic the Hedgehog as you dash, jump and spin your way across stunning 3D environments.  Swipe your way over and under challenging obstacles in this fast and frenzied endless running game for iPad, iPad mini, iPhone & iPod touch.

SONIC…
The world famous Sonic the Hedgehog stars in his first endless running game – how far can you go?

…DASH!
Unleash Sonic’s incredible dash move that allows you to run at insane speed and destroy everything in your path!

AMAZING ABILITIES
Utilise Sonic’s powers to dodge hazards, jump over barriers and speed around loop de loops.  Plus defeat enemies using Sonic’s devastating homing attack!

STUNNING GRAPHICS
Sonic’s beautifully detailed world comes to life on mobile and tablet – never has an endless runner looked so good!

MULTIPLE CHARACTERS
Choose to play as one of Sonic’s friends, including Tails, Shadow and Knuckles.

EPIC BOSS BATTLES
Face off against two of Sonic's biggest rivals, the always scheming and cunning Dr. Eggman and the devastatingly deadly Zazz from Sonic Lost World! Use all of Sonic's agility and speed to take down these villains before it's too late!

POWERUPS
Unlock, win or buy ingenious power-ups to help you run further. Including head starts, shields, ring magnets and unique score boosters!

KEEP ON RUNNING
Get more rewards the more you play! Level up your score multiplier by completing unique missions, or win amazing prizes including Red Star Rings & additional characters by completing Daily Challenges and playing the Daily Spin.

SOCIALLY CONNECTED
Challenge your friends on the leader boards or invite your friends through Facebook to prove who the best speed runner is…

COMING SOON
We’re working hard to bring you future FREE updates!

Sonic Dash supports iPhone 4 and higher, iPad 2 and higher, and iPod touch v5 and higher. 
PLEASE NOTE: iPod Touch 4th generation devices are currently not supported.

- - - - -
Privacy Policy: http://www.sega.com/mprivacy
Terms of Use: http://www.sega.com/terms

This game may include "Interest Based Ads" ​(please see http://www.sega.com/mprivacy#3IBADiscolure​ for more information)​​ and may collect "Precise Location Data" ​(please see http://www.sega.com/mprivacy#5LocationDataDisclosure​ for more information)
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
ReleaseNotes:
It's time to go fast all over again with improved stability and bug fixes!
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ ReleaseDate                     | 2013-03-07T08:00:00Z
┃ CurrentVersionReleaseDate       | 2017-06-22T16:04:16Z
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Lookup API

if you know something could used to identify a track, then lookup API may be a better approach.
Instead of specifying `term`, you need something like `iTunesID (track_id)`, `bundleID`(app only), `AMG ID`, etc...

Here's how it works:
```go

res, _ := Lookup().ID(414478124).Country(CN)
res.Print()
```

and it produce:

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┃ iTunes Track [ software / software ]
┃	414478124 微信 (微信)
┃	[com.tencent.xin]  https://itunes.apple.com/cn/app/%E5%BE%AE%E4%BF%A1/id414478124?mt=8&uo=4
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Artist:
┃	614694882 WeChat  https://itunes.apple.com/cn/developer/wechat/id614694882?uo=4
┃	Tencent Technology (Shenzhen) Company Limited http://weixin.qq.com
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Genre:
┃	6005 Social Networking
┃	[6005 6007] [社交 效率]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Icon:
┃ 	60:	http://is5.mzstatic.com/image/thumb/Purple128/v4/80/43/01/8043016a-7f84-9276-32f2-4e1ec0a8f3c6/source/60x60bb.jpg
┃ 	100:http://is5.mzstatic.com/image/thumb/Purple128/v4/80/43/01/8043016a-7f84-9276-32f2-4e1ec0a8f3c6/source/100x100bb.jpg
┃ 	512:http://is5.mzstatic.com/image/thumb/Purple128/v4/80/43/01/8043016a-7f84-9276-32f2-4e1ec0a8f3c6/source/512x512bb.jpg
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Price:	0 CNY 免费
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Environ:
┃ Languages:	[AR ZH EN FR DE HE HI ID IT JA KO MS PL PT RU ZH ES TH ZH TR VI]
┃ Features:		[iosUniversal]
┃ Devices:		[iPad2Wifi-iPad2Wifi iPad23G-iPad23G iPhone4S-iPhone4S iPadThirdGen-iPadThirdGen iPadThirdGen4G-iPadThirdGen4G iPhone5-iPhone5 iPodTouchFifthGen-iPodTouchFifthGen iPadFourthGen-iPadFourthGen iPadFourthGen4G-iPadFourthGen4G iPadMini-iPadMini iPadMini4G-iPadMini4G iPhone5c-iPhone5c iPhone5s-iPhone5s iPadAir-iPadAir iPadAirCellular-iPadAirCellular iPadMiniRetina-iPadMiniRetina iPadMiniRetinaCellular-iPadMiniRetinaCellular iPhone6-iPhone6 iPhone6Plus-iPhone6Plus iPadAir2-iPadAir2 iPadAir2Cellular-iPadAir2Cellular iPadMini3-iPadMini3 iPadMini3Cellular-iPadMini3Cellular iPodTouchSixthGen-iPodTouchSixthGen iPhone6s-iPhone6s iPhone6sPlus-iPhone6sPlus iPadMini4-iPadMini4 iPadMini4Cellular-iPadMini4Cellular iPadPro-iPadPro iPadProCellular-iPadProCellular iPadPro97-iPadPro97 iPadPro97Cellular-iPadPro97Cellular iPhoneSE-iPhoneSE iPhone7-iPhone7 iPhone7Plus-iPhone7Plus iPad611-iPad611 iPad612-iPad612 iPad71-iPad71 iPad72-iPad72 iPad73-iPad73 iPad74-iPad74]
┃ SystemRequirement:	8.0
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Rating:
┃ TrackContentRating:		12+
┃ ContentAdvisoryRating: 	12+
┃ Reason:					[偶尔/轻微的色情内容或裸露]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Ranking:
┃	Current:	4	809103
┃	Historic:	4.5	48392
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ Screenshots:
┃ Urls:	[http://a3.mzstatic.com/us/r30/Purple128/v4/4a/55/4f/4a554fb5-38c9-0a09-b8d8-e86bb2bb961d/screen696x696.jpeg http://a5.mzstatic.com/us/r30/Purple118/v4/d7/e1/dc/d7e1dc99-98c2-7752-509b-47dc7fea69d9/screen696x696.jpeg http://a2.mzstatic.com/us/r30/Purple128/v4/c4/8f/e0/c48fe02b-22de-c5ab-f6e6-7e43bd09e96e/screen696x696.jpeg http://a3.mzstatic.com/us/r30/Purple118/v4/14/c7/c4/14c7c42d-7731-3b55-cdb7-8b4e86611eb4/screen696x696.jpeg http://a4.mzstatic.com/us/r30/Purple128/v4/6c/04/49/6c04495a-9c16-650b-5285-bb7c4d7bac9b/screen696x696.jpeg]
┃ Ipad:	[http://a4.mzstatic.com/us/r30/Purple118/v4/dc/67/a3/dc67a38a-a174-fd60-acc9-c0841b1cb88c/sc1024x768.jpeg http://a2.mzstatic.com/us/r30/Purple118/v4/09/d1/a4/09d1a48b-2aef-9263-316a-76be9e8df3ea/sc1024x768.jpeg http://a1.mzstatic.com/us/r30/Purple128/v4/56/4a/77/564a7717-e82f-8f77-1e56-e8b5274d0df1/sc1024x768.jpeg]
┃ TV:	[]
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ GameCenter Enabled              | false
┃ VppDeviceBasedLicensingEnabled  | true
┣┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ FileSizeBytes                   | 197679104
┃ Version                         | 6.5.12
┗┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
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
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
ReleaseNotes:
本次更新
- 可以对视频进行编辑。
- 可以设置某条朋友圈的互动不再通知。
- 修复部分用户无法收到新消息提醒的问题。

最近更新
- 可在微信实验室体验正在探索的功能。
- 聊天中查找聊天内容时，可以查找文件、图片、链接。
- 群主可在群成员信息页中，了解对方是如何加入群聊的。
- 选择图片时，可便捷地调整并预览已选择的内容。
┏┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
┃ ReleaseDate                     | 2011-01-21T01:32:15Z
┃ CurrentVersionReleaseDate       | 2017-07-12T02:28:11Z
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

that can also achieve via `Lookup().CNAPP(id)` in short.

check [`api_test.go`](api_test.go) for more examples.



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

