package main

import(
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: parseNmap <input>")
		return
	}

	filename := args[1]

	content, err := ioutil.ReadFile(filename)
	check(err)

	// Parse XML
	var result Result
	err = xml.Unmarshal(content, &result)
	check(err)

	// Prepare output filename
	ext := path.Ext(filename)
	var output string
	if ext != "" {
		output = strings.Replace(filename, ext, ".csv", -1)
	} else {
		output = filename + ".csv"
	}

	f, err := os.Create(output)
	check(err)
	defer f.Close()

	f.WriteString("IP,Accuracy,Name\n")

	total := 0
	for _, host := range result.Host {
		if host.Status.State == "up" {
			for _, osmatch := range host.Os.OsMatch {
				if strings.Contains(osmatch.Name, "Windows") {
					f.WriteString(getIP(host) + ",")
					f.WriteString(osmatch.Accuracy + "%, " + strings.Replace(osmatch.Name, ",", "/", -1) + "\n")
					total += 1
					break
				}
			}
		}
	}
	f.WriteString(fmt.Sprintf("%d windows server scanned.\n", total))
	f.Sync()
}

func getIP(host Host) string {
	for _, addr := range host.Address {
		if addr.AddrType == "ipv4" {
			return addr.Addr
		}
	}
	return ""
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}