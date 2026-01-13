package maps

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value;
}

func (c *Cache) Get(key string) (int, bool) {
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) Count() int {
	return len(c.data)
}

func (c *Cache) AllKeys() []string {
	keys := make([]string, 0, 10)
	for key := range c.data {
		keys = append(keys, key)
	}
 	return keys
}

func (c *Cache) RemoveBelow(limit int) {
	for key, value := range c.data {
		if value < limit {
			delete(c.data, key)
		}
	}
}
