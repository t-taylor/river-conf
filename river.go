package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cfg := config{
		modKey: "Super",
	}
	cfg.r()
}

type config struct {
	modKey string
}

func (c *config) r(args ...string) {
	cmd := exec.Command("riverctl", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("riverctl %v : %v", args, err)
		fmt.Printf("%v", string(out))
		os.Exit(2)
	}
}
