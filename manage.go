package main

func (c *config) bindManage() {
	n := "normal"

	c.rMap(n, c.modKey, c.nextKey, "focus-view", "next")
	c.rMap(n, c.modKey, c.prevKey, "focus-view", "previous")

	c.rMap(n, c.modPower(), c.nextKey, "swap", "next")
	c.rMap(n, c.modPower(), c.prevKey, "swap", "previous")

	c.rMap(n, c.modKey, "N", "focus-output", "next")
	// c.rMap(n, c.modKey, c.leftKey, "focus-output", "previous")

	c.rMap(n, c.modPower(), "N", "send-to-output", "next")
	// c.rMap(n, c.modPower(), c.leftKey, "send-to-output", "previous")

	c.rMap(n, c.modKey, "V", "zoom")

	c.rMap(n, c.modKey, "Period", "send-layout-cmd", "rivertile", "main-ratio +0.05")
	c.rMap(n, c.modKey, "Comma", "send-layout-cmd", "rivertile", "main-ratio -0.05")

	c.rMap(n, c.modKey, "A", "send-layout-cmd", "rivertile", "main-count +1")
	c.rMap(n, c.modKey, "X", "send-layout-cmd", "rivertile", "main-count -1")

	// float manage
	c.rDef("map-pointer", n, c.modKey, "BTN_LEFT", "move-view")
	c.rDef("map-pointer", n, c.modKey, "BTN_RIGHT", "resize-view")

	c.rMap(n, c.modKey, "P", "toggle-float")
}
