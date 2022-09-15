package main

import (
	"flag"
	"fmt"

	mjson "github.com/eyedeekay/manifest-json-version/lib"
)

var manifestFile = flag.String("mf", "", "Path to the manifet file to extract info from")

func main() {
	flag.Parse()
	text := mjson.ManifestInfo(*manifestFile)
	fmt.Printf(text)
}
