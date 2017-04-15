package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	colorize "github.com/sabhiram/go-colorize"
)

// Version & Revision
var (
	Version  string
	Revision string
)

func init() {
	Version = "0.0.1"
	Revision = "0000000"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(colorize.ColorString("Error: Invalid Usage", "red"))
		return
	}

	res, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	re := regexp.MustCompile("(<a[^>]*>[^<]+</a>)")

	colorizedBody := re.ReplaceAllString(string(body), "<red>${1}</red>")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", colorize.Colorize(colorizedBody))
}
