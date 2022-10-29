package lrucache

import "testing"

func TestCache(t *testing.T) {

	cache := NewCache(5)
	t.Logf("cache init done...")

	cache.Add("1")
	cache.Add("2")
	cache.Add("3")
	cache.Add("4")
	cache.Add("5")

	t.Logf("cache dump = %#v", cache.GetAll())

	cache.Add("6")
	t.Logf("cache dump after adding 6 = %#v", cache.GetAll())

	cache.Add("7")
	t.Logf("cache dump after adding 7 = %#v", cache.GetAll())
}
