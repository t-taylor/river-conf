package main

// bindStart binds keys to programs to exec
func (c *config) bindStart() {
	c.rMap("normal", c.modKey, "M", "spawn", "foot")
	// run menu
	c.rMap("normal", c.modKey, "O", "spawn", "amenu -r")
	// pass menu
	c.rMap("normal", c.modKey, "G", "spawn", "amenu -p")

	c.rMap("normal", c.modKey, "F", "spawn", "$BROWSER")
}
