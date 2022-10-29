/*
Slice / Array based FIFO style cache of strings
Slice grows until preset capacity, and then eviction logic kicks in
For every eviction, it shifts all items back 1 position, and therefore runs more compute time
compared to LinkedList approach
*/
package lrucachelist

type CacheList struct {
	capacity  int
	itemsList []string
}

func (c *CacheList) Add(element string) {

	if len(c.itemsList) < c.capacity {
		c.itemsList = append(c.itemsList, element)
		return
	}

	// cache is at capacity
	// log.Printf("evicting item=%s", c.itemsList[0])
	for inx := 1; inx < len(c.itemsList); inx++ {
		c.itemsList[inx-1] = c.itemsList[inx]
	}
	c.itemsList[len(c.itemsList)-1] = element
}

func (c *CacheList) GetAll() []string {
	return c.itemsList
}

func (c *CacheList) Size() int {
	return len(c.itemsList)
}

func NewLRUCacheList(cap int) *CacheList {
	return &CacheList{
		capacity:  cap,
		itemsList: make([]string, 0),
	}
}
