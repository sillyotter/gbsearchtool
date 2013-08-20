package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sillyotter/gbsearch"
	"strings"
)

var (
	searchType   string
	languageCode string
	printType    string
	filter       string
	maxResults   int
	startIndex   int
	orderBy      string
	projection   string
	onlyEPub     bool
)

func init() {
	flag.StringVar(&searchType, "searchType", "unknown", "Set the type of search (title|author|publisher|subject|isbn|lccn|oclc)")
	flag.StringVar(&languageCode, "languageCode", "", "Set the two character language code (en|fr|...)")
	flag.StringVar(&printType, "mediaType", "", "Set the media type (all|book|magazine)")
	flag.StringVar(&filter, "contentType", "", "(any|full|partial|free-ebooks|paid-ebooks|ebooks)")
	flag.IntVar(&maxResults, "maxResults", -1, "Set maximum number of results to return <= 40")
	flag.IntVar(&startIndex, "startIndex", -1, "Set starting offset for fetch")
	flag.StringVar(&orderBy, "orderBy", "", "Sorty by (newest|relevance)")
	flag.StringVar(&projection, "resultType", "", "(full|lite)")
	flag.BoolVar(&onlyEPub, "onlyEPub", false, "Only search downloadable")
}

func determinePrintType(val string) gbsearch.PrintType {
	switch strings.ToLower(val) {
	case "all":
		return gbsearch.All
	case "book":
		return gbsearch.Books
	case "magazine":
		return gbsearch.Magazines
	}
	return gbsearch.UnknownPrintType
}

func determineFilterType(val string) gbsearch.FilterType {
	switch strings.ToLower(val) {
	case "full":
		return gbsearch.FullText
	case "partial":
		return gbsearch.PartialText
	case "free-ebooks":
		return gbsearch.FreeEbooks
	case "paid-ebooks":
		return gbsearch.PaidEbooks
	case "ebooks":
		return gbsearch.EBooks
	}
	return gbsearch.UnknownFilterType
}

func determineOrderBy(val string) gbsearch.OrderType {
	switch strings.ToLower(val) {
	case "relevance":
		return gbsearch.Relevance
	case "newest":
		return gbsearch.Newest
	}
	return gbsearch.UnknownOrderByType
}

func determineProjection(val string) gbsearch.ProjectionType {
	switch strings.ToLower(val) {
	case "full":
		return gbsearch.FullResults
	case "lite":
		return gbsearch.Lite
	}
	return gbsearch.UnknownProjectionType
}

func main() {

	flag.Parse()
	opt := gbsearch.DefaultOptions()
	opt.SetLanguageCode(languageCode)
	opt.OnlyFindEPubDownloads(onlyEPub)
	opt.SetStartIndex(startIndex)
	opt.SetMaxResults(maxResults)
	opt.SetPrintType(determinePrintType(printType))
	opt.SetFilter(determineFilterType(filter))
	opt.SetOrderBy(determineOrderBy(orderBy))
	opt.SetProjection(determineProjection(projection))

	var st gbsearch.SearchType = gbsearch.UnknownSearchType

	switch strings.ToLower(searchType) {
	case "title":
		st = gbsearch.InTitle
	case "author":
		st = gbsearch.InAuthor
	case "publisher":
		st = gbsearch.InPublisher
	case "subject":
		st = gbsearch.Subject
	case "isbn":
		st = gbsearch.ISBN
	case "lccn":
		st = gbsearch.LCCN
	case "oclc":
		st = gbsearch.OCLC
	}

	for _, elem := range flag.Args() {
		res, _ := gbsearch.Search(st, elem, opt)
		out, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(out))
	}
}
