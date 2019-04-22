// Build.gradle parser written in Go
// github . com /konradit/android-go

package main

import (
	"bufio"
	"strings"

	"flag"
	"github.com/apsdehal/go-logger"
	"io/ioutil"
	"os"
)

// utils

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func loadEntries(data string) {

	log, err := logger.New("loadEntries", 1, os.Stdout)
	isErr(err)

	// to do
	//dependencies := make(map[string][]string)
	android := make(map[string][]string)
	var settingAndroid = false
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(data)))
	for scanner.Scan() {
		// new line
		splitLine := strings.Split(scanner.Text(), " ")
		if splitLine[0] == "android" {
			// is obj
			log.NoticeF("[android] %s", splitLine[0])
			settingAndroid = true
		}

		if settingAndroid == true {
			if scanner.Text() == "}" {
				log.Warning("Close!")
				settingAndroid = false
			} else {
				log.NoticeF("[0] %s", scanner.Text())

				android[splitLine[0]] = splitLine
				log.Notice(android[splitLine[0]][0])
			}
		}
	}

}
func main() {

	// House keeping
	log, err := logger.New("Main", 1, os.Stdout)
	isErr(err)
	log.Info("gradle_parser")

	// Parse file name
	fileLocation := flag.String("i", "build.gradle", "Build.gradle file location")
	flag.Parse()
	if isFlagPassed("i") {
		log.NoticeF("Parsing filename... %s", *fileLocation)

		dat, err := ioutil.ReadFile(*fileLocation)
		isErr(err)
		loadEntries(string(dat))
	} else {
		log.Error("No build.gradle passed!")
	}

}
