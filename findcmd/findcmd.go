package findcmd

import (
	"io/ioutil"
	"log"
	"regexp"
)

func doFindCommand(path string, regex string, chanResults chan string) {

	r, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("error: failed to regexp compile")
		return
	}

	recurseFindCommand(path, r, chanResults)

	chanResults <- ""
}

func recurseFindCommand(path string, r *regexp.Regexp, chanResults chan string) {

	filesinfo, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	// results := make([]string, 0)

	for _, fileinfo := range filesinfo {
		// log.Printf("inspecting item  = %s", fileinfo.Name())
		if fileinfo.IsDir() {
			// log.Printf("path=%s is a directory", path+fileinfo.Name())
			recurseFindCommand(path+fileinfo.Name()+"/", r, chanResults)

		} else {

			// log.Printf("matching on the file=%s", path+fileinfo.Name())
			if r.Match([]byte(fileinfo.Name())) {
				chanResults <- path + fileinfo.Name()
			}
		}
	}
}
