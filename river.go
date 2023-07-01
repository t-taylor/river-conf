package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cfg := &config{
		modKey: "Super",
	}
	cfg.bindNav()
}

type config struct {
	modKey      string
	powerModKey string
	altModKey   string

	nextKey  string
	prevKey  string
	leftKey  string
	rightKey string
}

func (c *config) rDef(args ...string) {
	cmd := exec.Command("riverctl", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("riverctl %v : %v", args, err)
		fmt.Printf("%v", string(out))
		os.Exit(2)
	}
}

func (c *config) rMap(args ...string) {
	args = append([]string{"map"}, args...)
	c.rDef(args...)
}
