package main

import (
	"fmt"
	"os"
	"os/exec"
)

func (c *config) setupLayouts() {
	c.rDef("default-layout", "rivertile")

	execChild("rivertile", "-view-padding", "6", "-outer-padding", "6")
}

func execChild(cmdStr string, args ...string) {
	cmd := exec.Command(cmdStr, args...)
	if err := cmd.Start(); err != nil {
		fmt.Printf("%s %v : %v", cmdStr, args, err)
		os.Exit(2)
	}
}
