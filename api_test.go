package itunes_search

import (
	"testing"
)

// Simplest search example
func TestSearch(t *testing.T) {
	res, _ := Search([]string{"Hello", "World"}).
		Country(US).App().Limit(5).Results()

	for _, r := range res {
		r.Print()
	}
}

// Simplest lookup example
func TestLookup(t *testing.T) {
	res, _ := Lookup().ID(989673964).Result()
	res.Print()
}

// App Specific API
func TestEntry_DetailCN(t *testing.T) {
	res, _ := Lookup().ID(989673964).Result()
	res.Detail(CN).Print()
}

// Test US Store
func TestEntry_DetailUS(t *testing.T) {
	res, _ := Lookup().Country(US).ID(414478124).Result()
	res.Detail(US).Print()
}

func TestLookupCNStoreByiTunesID(t *testing.T) {
	testCase := []struct {
		ID         int64
		ExpectName string
	}{
		{414478124, "微信"},
		{534453594, "保卫萝卜1"},
		{529479190, "部落冲突 (Clash of Clans)"},
		{510940882, "找你妹"},
	}

	for _, c := range testCase {
		if res, err := Lookup().ID(c.ID).Result(); err != nil {
			t.Error(err)
		} else {
			// res.Print()
			if res.TrackName != c.ExpectName {
				t.Errorf("expect name of id %d is %s, got %s", c.ID, c.ExpectName, res.TrackName)
			}
		}
	}
}

func TestLookupCNAppByBundleID(t *testing.T) {
	testCase := []struct {
		BundleID   string
		ExpectName string
	}{
		{"com.tencent.xin", "微信"},
		{"cairot", "保卫萝卜1"},
		{"com.supercell.magic", "部落冲突 (Clash of Clans)"},
		{"com.funship.smarteye", "找你妹"},
	}

	for _, c := range testCase {
		if res, err := Lookup().BundleID(c.BundleID).Country(CN).Result(); err != nil {
			t.Error(err)
		} else {
			// res.Print()
			if res.TrackName != c.ExpectName {
				t.Errorf("expect name of bundleID %s is %s, got %s", c.BundleID, c.ExpectName, res.TrackName)
			}
		}
	}
}
