package app

import "github.com/go-pg/pg"
import . "github.com/Vonng/go-itunes-search"

func (app *App) Save(db *pg.DB) error {
	_, err := db.Model(app).
		OnConflict("(id) DO UPDATE").
		Set("name= ?name").
		Set("url= ?url").
		Set("icon= ?icon").
		Set("kind= ?kind").
		Set("version= ?version").
		Set("bundle_id= ?bundle_id").
		Set("author_id= ?author_id").
		Set("author_name= ?author_name").
		Set("author_url= ?author_url").
		Set("vendor_name= ?vendor_name").
		Set("vendor_url= ?vendor_url").
		Set("copyright= ?copyright").
		Set("genre_id= ?genre_id").
		Set("genre_id_list= ?genre_id_list").
		Set("genre= ?genre").
		Set("genre_list= ?genre_list").
		Set("icon60= ?icon60").
		Set("icon100= ?icon100").
		Set("price= ?price").
		Set("currency= ?currency").
		Set("system= ?system").
		Set("features= ?features").
		Set("devices= ?devices").
		Set("languages= ?languages").
		Set("platforms= ?platforms").
		Set("rating= ?rating").
		Set("reasons= ?reasons").
		Set("size= ?size").
		Set("cnt_rating= ?cnt_rating").
		Set("avg_rating= ?avg_rating").
		Set("cnt_rating_current= ?cnt_rating_current").
		Set("avg_rating_current= ?avg_rating_current").
		Set("vpp_device= ?vpp_device").
		Set("game_center= ?game_center").
		Set("screenshots= ?screenshots").
		Set("in_app_purchase= ?in_app_purchase").
		Set("sibling_apps= ?sibling_apps").
		Set("related_apps= ?related_apps").
		Set("support_sites= ?support_sites").
		Set("reviews= ?reviews").
		Set("extra= ?extra").
		Set("description= ?description").
		Set("release_notes= ?release_notes").
		Set("release_time= ?release_time").
		Set("publish_time= ?publish_time").
		Set("crawled_time= ?crawled_time").
		Insert()
	return err
}

func NewAppByID(id int64, country string) (*App, error) {
	entry, err := Lookup().ID(id).Country(country).App().Result()
	if err != nil {
		return nil, err
	}
	return NewDetailedApp(entry, country), nil
}

func NewAppByBundleID(bundleID string, country string) (*App, error) {
	entry, err := Lookup().BundleID(bundleID).Country(country).App().Result()
	if err != nil {
		return nil, err
	}
	return NewDetailedApp(entry, country), nil
}
