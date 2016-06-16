package main

import (
	"os"
	"flag"
	"log"
	"fmt"
	"io/ioutil"
	"github.com/soywod/archive"
	"github.com/soywod/file64"
)

func main() {
	var source, destZip, destGo, pkg, funcName string

	flag.StringVar(&source, "d", "", "Input directory to compress")
	flag.StringVar(&destZip, "o", "./archive.zip", "Output ZIP path + name")
	flag.StringVar(&destGo, "g", "./archive.go", "Output Go file path + name")
	flag.StringVar(&pkg, "p", "main", "Package name of the Go file")
	flag.StringVar(&funcName, "f", "GetArchive", "Function name that will return the base64 archive data")
	flag.Parse()

	if source == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := archive.CompressFolder(source, destZip)
	defer os.Remove(destZip)
	if err != nil {
		log.Fatal(err)
	}

	b64code, err := file64.Encode(destZip)
	if err != nil {
		log.Fatal(err)
	}

	goCode := fmt.Sprintf("package %s\nfunc %s() string {return %q}", pkg, funcName, b64code)
	
	if err := ioutil.WriteFile(destGo, []byte(goCode), 0660); err != nil {
		log.Fatal(err)
	}
}
