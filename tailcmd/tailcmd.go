package tailcmd

import (
	"bufio"
	"os"

	"practise.go.codingchallenges/lrucachelist"
	// "practice_scripting/lrucache"
)

func GetNLines(filename string, numberOfLines int) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// cache := lrucache.NewCache(numberOfLines)
	cache := lrucachelist.NewLRUCacheList(numberOfLines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cache.Add(line)
	}

	return cache.GetAll(), scanner.Err()
}
