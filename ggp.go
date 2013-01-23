// ggp stands for Go Get Packages.
// You can find Go packages from http://godoc.org
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// PkgData has path, desc.
//   path: a package repository path
//   desc: the package description
type PkgData struct {
	path string
	desc string
}

// getPackageData returns pakcage data.
// You can get package infomation from http://godoc.org
// and refer to http://godoc.org/-/about
func getPackageData(query string) []*PkgData {
	query = fmt.Sprintf("http://godoc.org/?q=%v", query)
	req, err := http.NewRequest("GET", query, nil)

	if nil != err {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("Accept", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)

	if nil != err {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println(err)
		return nil
	}

	buff := string(body)
	lines := strings.Split(buff, "\n")
	size := len(lines)
	var pkgData = make([]*PkgData, size)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) < 0 {
			break
		}

		data := strings.SplitN(line, " ", 2)
		if len(data) != 2 {
			continue
		}

		pkgData[i] = &PkgData{data[0], data[1]}
	}

	return pkgData
}

func usage() {
	flag.PrintDefaults()
	os.Exit(1)
}

func parseOpts() (query string) {
	flag.Usage = usage
	flag.StringVar(&query, "q", "", "query")
	flag.Parse()

	args := flag.Args()
	if len(args) != 0 {
		flag.Usage()
	}

	return
}

func searchPackage(query string) {
	pkgData := getPackageData(query)
	fmt.Println("package\tdescription")
	fmt.Println("=============================")

	for _, pkg := range pkgData {
		if nil == pkg {
			break
		}

		fmt.Printf("* %v\n\t- %v\n\n", pkg.path, pkg.desc)
	}
}

func main() {
	query := parseOpts()
	searchPackage(query)
}
