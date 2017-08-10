package main

import (
	"os"
	"fmt"
	"flag"
	"strconv"
)

import . "github.com/Vonng/go-itunes-search"
import . "github.com/Vonng/go-itunes-search/app"
import "github.com/olekukonko/tablewriter"

// flags
var (
	itunesID string
	bundleID string
	keywords string
	nResult  int    = 10
	country  string = CN
	detail   bool   = false
	list     bool
)

func HandleSearch(keywords string) {
	res, err := SearchOne(keywords).Country(country).App().Limit(nResult).Results()
	if err != nil {
		fmt.Printf("error when search %s: %s\n", keywords, err.Error())
		os.Exit(-1)
	}

	fmt.Printf("%d result returned.\n", len(res))
	if list {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"iTunesID", "BundleID", "Name", "Ver"})
		for _, entry := range res {
			table.Append([]string{
				strconv.FormatInt(entry.TrackID, 10),
				entry.BundleID,
				entry.TrackName,
				entry.Version,
			})
		}
		table.Render()
	} else if detail {
		for _, entry := range res {
			NewDetailedApp(&entry, country).Print()
		}
	} else {
		for _, entry := range res {
			NewDetailedApp(&entry, country).Print()
		}
	}
}

func HandleLookup(idType, idValue string) {
	entry, err := Lookup().SetParam(idType, idValue).Country(country).App().Limit(1).Result()
	if err != nil {
		fmt.Printf("error when looup %s:%s: %s\n", idType, idValue, err.Error())
		os.Exit(-1)
	}
	if detail {
		NewDetailedApp(entry, country).Print()
	} else {
		NewApp(entry).Print()
	}
}

func main() {
	flag.StringVar(&itunesID, "i", "", "id for lookup. eg:414478124")
	flag.StringVar(&bundleID, "b", "", "bundleID for lookup eg:com.tencent.xin")
	flag.StringVar(&keywords, "s", "", "searching keyword eg:HelloWorld")
	flag.IntVar(&nResult, "n", 10, "number of result size. 1~200,default:50")
	flag.StringVar(&country, "c", CN, "restrict to country. default:CN")
	flag.BoolVar(&detail, "d", false, "fetch extra details. default:disabled")
	flag.BoolVar(&list, "l", false, "show result in list format")
	flag.Parse()

	switch {
	// iTunes id is provided, using lookup api
	case itunesID != "":
		HandleLookup(ITunesID, itunesID)
	case bundleID != "":
		HandleLookup(BundleID, bundleID)
	case keywords != "":
		HandleSearch(keywords)
	default:
		flag.Usage()
	}
}
