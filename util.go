package golangds

type container struct {
	store []*interface{}
}

// Empty func
func (c *container) Empty() bool {
	return len(c.store) == 0
}

// Size func
func (c *container) Size() uint {
	return uint(len(c.store))
}

// Clear func
func (c *container) Clear() {
	c.store = nil
}
