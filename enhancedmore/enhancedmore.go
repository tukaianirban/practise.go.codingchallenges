/**
Problem definition: From all files in a directory, print all lines from all files that contain a matching keyword(s)
You are provided a buffer of size K within which the printed lines must be fitted.
If the user presses ENTER, get the next K lines matching the conditions above, and fit them into the buffer
**/
package enhancedmore

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

type ListNode struct {
	Val  string
	Next *ListNode
}

type Buffer struct {
	head     *ListNode
	tail     *ListNode
	PageSize int
	Length   int
}

func createBuffer(size int) *Buffer {
	return &Buffer{
		head:     nil,
		tail:     nil,
		PageSize: size,
		Length:   0,
	}
}

// reset the buffer to make space for adding a new page
func (s *Buffer) setNextPage() {

	s.head = nil
	s.tail = nil
	s.Length = 0
}

// returns a bool on whether the line could be fitted into the buffer or not
func (s *Buffer) AddLineToBuffer(line string) bool {

	if s.Length == s.PageSize {
		log.Print("buffer is already full")
		return false
	}

	if s.Length == 0 {
		s.head = &ListNode{
			Val:  line,
			Next: nil,
		}

		s.tail = s.head
		s.Length = 1
		return true
	}

	s.tail.Next = &ListNode{
		Val:  line,
		Next: nil,
	}

	s.tail = s.tail.Next
	s.Length++

	return true
}

func (s *Buffer) GetBuffer() []string {

	result := make([]string, 0)
	movingPtr := s.head
	for movingPtr != nil {
		result = append(result, movingPtr.Val)
		movingPtr = movingPtr.Next
	}

	return result
}

var chMoveForwards = make(map[string]chan bool)
var buffer *Buffer = nil

func findMatchingLinesInFolder(path string, substring string, bufferSize int) {

	buffer = createBuffer(bufferSize)

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("error reading folder, reason=%s", err.Error())
		return
	}

	var wg sync.WaitGroup

	for _, fileInfo := range fileInfos {

		if fileInfo.IsDir() {
			continue
		}

		chMoveForward := make(chan bool)
		wg.Add(1)
		go readFromFile(path+"/"+fileInfo.Name(), substring, chMoveForward, &wg)
	}

	wg.Wait()
}

func setNextPage() {

	if buffer == nil {
		return
	}

	buffer.setNextPage()

	for _, chMoveForward := range chMoveForwards {
		chMoveForward <- true
	}
}

func getBufferDump() []string {

	if buffer == nil {
		return nil
	}

	return buffer.GetBuffer()
}

func readFromFile(path string, substring string, chMoveForward chan bool, wg *sync.WaitGroup) {

	defer wg.Done()
	defer delete(chMoveForwards, path)

	file, err := os.Open(path)
	if err != nil {
		log.Printf("error opening file=%s reason=%s", path, err.Error())
		return
	}
	defer file.Close()

	chMoveForwards[path] = chMoveForward

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		if !strings.Contains(line, substring) {
			continue
		}

		isAddSuccess := buffer.AddLineToBuffer(path + " :::: " + line)

		if !isAddSuccess {

			flag := <-chMoveForward
			if !flag {
				return
			}
		}
	}
}
