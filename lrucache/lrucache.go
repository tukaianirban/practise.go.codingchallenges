/*
Linked List based FIFO cache, which evicts entries in FIFO style
Linked List grows to max capacity as inititally set, and then starts with eviction logic
LinkedList in itself does not mutate, and therefore cache insertion and eviction is 0-delay compute time
*/
package lrucache

type CacheNode struct {
	data string
	Next *CacheNode
}

type CacheList struct {
	capacity  int
	length    int
	firstNode *CacheNode
	lastNode  *CacheNode
}

func (c *CacheList) Add(element string) {

	// log.Printf("adding element=%s", element)

	if c.firstNode == nil {
		// log.Printf("cache is empty, adding first element")
		c.firstNode = &CacheNode{
			data: element,
			Next: nil,
		}

		c.lastNode = c.firstNode
		c.length = c.length + 1

		return
	}

	if c.length >= c.capacity {
		// already at capacity; evict the first node
		// log.Printf("cache at capacity; evicting the first element=%s", c.firstNode.data)
		c.firstNode = c.firstNode.Next
		c.length = c.length - 1
	}

	c.lastNode.Next = &CacheNode{
		data: element,
		Next: nil,
	}

	c.length = c.length + 1
	c.lastNode = c.lastNode.Next

	// log.Printf("after adding, firstnode=%s lastnode=%s", c.firstNode.data, c.lastNode.data)
}

func (c *CacheList) GetAll() []string {
	ref := c.firstNode

	result := make([]string, 0)
	for ref != nil {
		result = append(result, ref.data)
		ref = ref.Next
	}

	return result
}

func (c *CacheList) GetSize() int {
	// log.Printf("length=%d capacity=%d", c.length, c.capacity)
	return c.length
}

func NewCache(cap int) *CacheList {

	return &CacheList{
		capacity:  cap,
		length:    0,
		firstNode: nil,
		lastNode:  nil,
	}
}
