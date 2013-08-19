package main

import (
	"encoding/json"
	"flag"
	"github.com/sillyotter/gbsearch"
	"log"
)

var isbn = flag.String("isbn", "0394758269", "ISBN number")

func main() {
	flag.Parse()
	opt := gbsearch.DefaultOptions()
	opt.SetLanguageCode("en")
	opt.SetPrintType(gbsearch.All)
	res, _ := gbsearch.ISBNSearch(*isbn, opt)
	out, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(out))
}
