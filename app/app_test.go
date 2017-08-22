package app

import "testing"
import . "github.com/Vonng/go-itunes-search"
import "github.com/go-pg/pg"

// App Specific API
func TestNewApp(t *testing.T) {
	entry, _ := Lookup().ID(989673964).Result()
	NewDetailedApp(entry, "CN").Print()
}

// Test US Store
func TestEntry_DetailUS(t *testing.T) {
	entry, _ := Lookup().Country(US).ID(1061097588).Result()
	NewDetailedApp(entry, "US").Print()
}

func TestApp_Save(t *testing.T) {
	Pg := pg.Connect(&pg.Options{
		Addr:     ":5432",
		Database: "haha",
		User:     "haha",
	})

	idList := []int64{
		989673964,
		1110193350,
		1110195252,
		1110194837,
	}

	for _, id := range idList {
		app, err := NewAppByID(id, "CN")
		if err != nil {
			t.Error(err)
		}
		err = app.Save(Pg)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestApp_Save2(t *testing.T) {

}
