package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("TERMINAL", "foot")
	sessionDesktop := os.Getenv("XDG_SESSION_DESKTOP")

	mk := "Super"
	// if running in embeded dwm
	if sessionDesktop == "dwm" {
		// set mod key to alt
		mk = "Alt"
	}
	cfg := &config{
		modKey:      mk,
		powerModKey: "Shift",
		altModKey:   "Ctrl",

		nextKey:  "J",
		prevKey:  "K",
		leftKey:  "H",
		rightKey: "L",
	}

	cfg.bindManage()
	cfg.bindStart()
	cfg.setupLayouts()
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
	mapArgs := make([]string, len(args)+1)
	mapArgs[0] = "map"

	for i, v := range args {
		mapArgs[i+1] = v
	}

	c.rDef(mapArgs...)
}

func (c *config) modPower() string {
	return c.modKey + "+" + c.powerModKey
}

func (c *config) modAlt() string {
	return c.modKey + "+" + c.altModKey
}
