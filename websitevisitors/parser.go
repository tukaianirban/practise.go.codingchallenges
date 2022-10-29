package websitevisitors

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IPAddress struct {
	parts []int
}

// returns -1 when this is less than other; +1 when this is more then other
// return 0 when both are same
func (ipaddr IPAddress) compare(otherIP IPAddress) int {

	for i := range ipaddr.parts {
		if ipaddr.parts[i] < otherIP.parts[i] {
			return -1
		}

		if ipaddr.parts[i] > otherIP.parts[i] {
			return 1
		}
	}

	return 0
}

func (ipaddr IPAddress) isValid() bool {
	if ipaddr.parts == nil {
		return false
	}

	for _, part := range ipaddr.parts {
		if part < 0 || part > 255 {
			return false
		}
	}

	return true
}

func (ipaddr IPAddress) toString() string {
	str := ""
	for _, part := range ipaddr.parts {
		str = str + strconv.Itoa(part) + "."
	}
	return str[:len(str)-1]
}

func NewIPAddress(value string) (IPAddress, error) {

	parts := strings.Split(value, ".")
	if len(parts) != 4 {
		return IPAddress{}, fmt.Errorf("not an IP")
	}

	numbers := make([]int, 4)

	for i, part := range parts {

		p, err := strconv.Atoi(part)
		if err != nil || (p < 0 || p > 255) {
			return IPAddress{}, fmt.Errorf("not an IP")
		}
		numbers[i] = p
	}

	return IPAddress{
		parts: numbers,
	}, nil
}

// break the line into space-separated strings
// find the first string that matches an IP addr. format
func fetchIPFromLogsLine(line string) (IPAddress, error) {

	parts := strings.Split(line, " ")
	if len(parts) == 0 {
		return IPAddress{}, fmt.Errorf("no IP found")
	}

	for _, part := range parts {

		if ipaddr, err := NewIPAddress(part); err == nil {
			return ipaddr, nil
		}
	}

	return IPAddress{}, fmt.Errorf("no IP found")
}

func getAllIPAddresses(filename string) ([]IPAddress, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ipaddresses := make([]IPAddress, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if ipaddr, err := fetchIPFromLogsLine(line); err == nil {
			ipaddresses = append(ipaddresses, ipaddr)
		}
	}

	// deliberately skipping the error handling scanner.Err() at the end as it should not matter
	return ipaddresses, nil
}

func GetNPopularVisitors(nCount int, filename string) []string {

	allIPAddresses, err := getAllIPAddresses(filename)
	if err != nil {
		return []string{}
	}

	// uniqueIPAddresses := make([]IPAddress, 0)

	// create occurences map
	occurences := make(map[string]int)

	uniqueFound := true
	for uniqueFound {

		uniqueFound = false
		uniqueIP := ""
		count := 0
		for _, ipaddr := range allIPAddresses {

			if _, ok := occurences[ipaddr.toString()]; ok {
				continue
			}

			if uniqueIP == "" {
				uniqueIP = ipaddr.toString()
				uniqueFound = true
				count = 1
			} else if ipaddr.toString() == uniqueIP {
				count = count + 1
			}
		}

		if uniqueIP != "" {
			occurences[uniqueIP] = count
		}
	}

	// log.Print("occurences map:")
	// for ip, count := range occurences {
	// 	log.Printf("IP=%s count=%d", ip, count)
	// }

	// sort based on the occurences map
	respList := make([]string, nCount)
	maxCount := len(allIPAddresses) // the highest count possible is if all hits are from single IP

	for inx := 0; inx < nCount; inx++ {

		// find the highest number less than the previous largest
		highestCount := 0
		highestCountIP := ""
		for ip, count := range occurences {

			if count >= maxCount {
				continue
			}
			if count > highestCount {
				highestCount = count
				highestCountIP = ip
			}
		}

		if highestCount == 0 && highestCountIP == "" {
			break
		}

		// log.Printf("new popular ip=%s count=%d", highestCountIP, highestCount)
		maxCount = highestCount
		respList[inx] = highestCountIP
	}

	return respList
}
