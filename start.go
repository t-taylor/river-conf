package main

// bindStart binds keys to programs to exec
func (c *config) bindStart() {
	n := "normal"

	c.rMap(n, c.modKey, "Q", "close")

	c.rMap(n, c.modKey, "M", "spawn", "foot")
	// run menu
	c.rMap(n, c.modKey, "O", "spawn", "amenu -r")
	// pass menu
	c.rMap(n, c.modKey, "G", "spawn", "amenu -p")

	c.rMap(n, c.modKey, "F", "spawn", "qutebrowser")
}
