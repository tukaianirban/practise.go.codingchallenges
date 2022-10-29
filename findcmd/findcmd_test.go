package findcmd

import "testing"

func TestDoFindCommand(t *testing.T) {

	chanResults := make(chan string, 10)
	go doFindCommand("/Users/anirbanmukherjee/programming/react-native/", ".*(.js)$", chanResults)
	t.Logf("results for .js files in path: ")

	for result := range chanResults {
		t.Logf(result)

		if result == "" {
			break
		}
	}
}
