package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/0x75960/liner"
)

var inqPath = flag.String("i", "", "input file path")
var dbPath = flag.String("db", "", "database file path")
var pickNotFound = flag.Bool("v", false, "pick up not founded on database")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			"ILD\n\tinquiry each lines to line-based database\n\nUsage of this:\n\t%s -i [file to inquiry] -db [database file]\n\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *inqPath == "" {
		log.Fatalln("-i option required. see --help")
	}

	if *dbPath == "" {
		log.Fatalln("-db option required. see --help")
	}

	fl, err := os.Open(*dbPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer fl.Close()

	db := map[string]bool{}

	for line := range liner.LinesIn(fl) {
		db[line] = true
	}

	fs, err := os.Open(*inqPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer fs.Close()

	for line := range liner.LinesIn(fs) {
		if db[line] != *pickNotFound {
			fmt.Println(line)
		}
	}

}
