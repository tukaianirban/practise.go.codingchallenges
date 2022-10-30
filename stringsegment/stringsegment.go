package stringsegment

import "strings"

var dictionary = []string{}

func findSegments(masterWord string, dictWords []string) []string {

	dictionary = dictWords

	isSegmantable, segments := segmentMasterString([]string{}, masterWord)
	if !isSegmantable {
		return nil
	}

	return segments
}

func segmentMasterString(alreadyScoped []string, remainingMaster string) (bool, []string) {

	// check if any of the dictionary words can be extracted from remainingMaster
	for _, dword := range dictionary {

		idx := strings.Index(remainingMaster, dword)
		if idx < 0 {
			continue
		}

		if len(dword) == len(remainingMaster) {

			// exact match
			return true, append(alreadyScoped, dword)
		}

		flag, result := segmentMasterString(append(alreadyScoped, dword), remainingMaster[(idx+len(dword)):])
		if flag {
			return true, result
		}
	}

	return false, alreadyScoped
}
