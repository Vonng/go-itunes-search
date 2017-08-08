package main

import (
	"os"
	"fmt"
	"regexp"
)
import (
	. "github.com/Vonng/go-itunes-search"
	"strings"
)

var idPattern = regexp.MustCompile(`(\d+)`)

func Usage() {
	fmt.Println(`Usage:
	./itunes [search] term1 term2 ...
	./itunes lookup <idType> <idValue>
	./itunes lookup <iTunesID>|<BundleID>

	eg: 414478124 com.tencent.xin WeChat`)
	os.Exit(-1)
}

func HandleSearch(args []string) {
	res, err := Search(args).Country(CN).App().Limit(10).Results()
	if err != nil {
		fmt.Printf("error when search %s: %s\n", args, err.Error())
		os.Exit(-1)
	}

	fmt.Println("%d result returned.", len(res))
	for _, r := range res {
		//r.Detail().Print()
		r.ToApp().Print()
	}
}

func HandleLookup(args []string) {
	var idType, idValue string
	if args == nil || len(args) < 1 {
		Usage()
	}

	if len(args) == 2 {
		idType, idValue = args[0], strings.TrimSpace(args[1])
	} else {

		if match := idPattern.FindStringSubmatch(args[0]); match != nil {
			idType, idValue = ITunesID, args[0]
		} else {
			idType, idValue = BundleID, args[0]
		}
	}

	res, err := Lookup().SetParam(idType, idValue).Country(CN).App().Limit(1).Result()
	if err != nil {
		fmt.Printf("error when looup %s:%s: %s\n", idType, idValue, err.Error())
		os.Exit(-1)
	}

	res.Detail().Print()
}

func main() {
	nargs := len(os.Args)
	if nargs < 2 {
		Usage()
	}
	switch os.Args[1] {
	case "lookup":
		HandleLookup(os.Args[2:])
	case "search":
		HandleSearch(os.Args[2:])
	default:
		HandleSearch(os.Args[1:])
	}
}
