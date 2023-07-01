package main

func (c *config) bindNav() {
	c.rDef("normal", c.modKey, c.nextKey, "focus-view", "next")
	c.rDef("normal", c.modKey, c.prevKey, "focus-view", "previous")
}
