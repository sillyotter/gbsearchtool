package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/sillyotter/gbsearch"
	"log"
	"os"
)

var (
	ErrUsageNotFound    = errors.New("Usage section not found")
	ErrMoreThanOneUsage = errors.New("Usage section found more than once")
)

func parseSection(section string, doc string) ([]string, error) {
	return make([]string, 0), nil
}

func parsePattern(usage string, options []string) error {
	return nil
}

func parseDefaults(doc string) ([]string, error) {
	return make([]string, 0), nil
}

func formalUsage(section string) string {
	return ""
}

func docopt(doc string, argv []string, help bool, version string, optionsFirst bool) (map[string]string, error) {
	args := argv[1:]
	log.Println(args)

	usage_sections, err := parseSection("usage:", doc)
	if err != nil {
		return nil, err
	}

	if len(usage_sections) == 0 {
		return nil, ErrUsageNotFound
	}

	if len(usage_sections) > 1 {
		return nil, ErrMoreThanOneUsage
	}

	options, err := parseDefaults(doc)
	parsePattern(formalUsage(usage_sections[0]), options)

	return make(map[string]string), nil
}

var isbn = flag.String("isbn", "0394758269", "ISBN number")

func main() {
	doc := `
Usage:
    my_program tcp <host> <port> [--timeout=<seconds>]
    my_program serial <port> [--baud=<n>] [--timeout=<seconds>]
    my_program (-h | --help | --version)

Options:
    -h, --help  Show this screen and exit.
    --baud=<n>  Baudrate [default: 9600]`

	x, _ := docopt(doc, os.Args, true, "", false)
	log.Println(x["tcp"])

	flag.Parse()
	opt := gbsearch.DefaultOptions()
	opt.SetLanguageCode("en")
	opt.SetPrintType(gbsearch.All)
	res, _ := gbsearch.ISBNSearch(*isbn, opt)
	out, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(out))
}
