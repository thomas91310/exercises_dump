package main

import "fmt"

//LRUCache struct
type LRUCache struct {
	elements map[string]*Node
	capacity int
	orders   LinkedList
}

func (c *LRUCache) getValueForKey(key string) string {
	if len(c.elements) == c.capacity {
		toRemove := c.orders.root.key
		c.orders.root = c.orders.root.next
		delete(c.elements, toRemove)
	}
	node, exists := c.elements[key]
	if exists {
		node.prev.next = node.next
		node.prev = c.orders.end
		node.next = nil
		c.orders.end.next = node
		return node.value
	}
	newNode := NewNode(key, c.getValueFromDB(key), c.orders.end, nil)
	if len(c.elements) == 0 {
		c.orders.root = newNode
		c.orders.end = newNode
	}
	c.elements[key] = newNode
	c.elements[key].prev = c.orders.end
	c.elements[key].next = nil
	c.orders.end.next = newNode
	c.orders.end = newNode
	return newNode.value
}

func (c *LRUCache) getValueFromDB(key string) string {
	return fmt.Sprintf("value for %v", key)
}

func main() {
	capacity := 5
	c := LRUCache{
		elements: make(map[string]*Node, capacity),
		capacity: capacity,
	}
	c.getValueForKey("A")
	c.getValueForKey("B")
	c.getValueForKey("C")
	c.getValueForKey("D")
	c.getValueForKey("E")
	c.getValueForKey("A")
	c.getValueForKey("F")
	c.orders.print()
}
