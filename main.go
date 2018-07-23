package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var result = 0
	if len(os.Args) < 2 {
		help()
	}

	if len(os.Args) >= 2 {
		path, err := exec.LookPath(fmt.Sprintf("git-hook-%s", os.Args[1]))
		if err != nil {
			fmt.Printf("error: %s\n", err)
			result = 1
		} else {
			cmd := exec.Command(path)
			cmd.Args = os.Args[1:]
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin

			if err := cmd.Start(); err != nil {
				fmt.Println(fmt.Sprintf("error: unable to start [%s]: %s", path, err.Error()))
				result = 1
			}
			if err := cmd.Wait(); err != nil {
				result = 1
			}
		}
	}
	os.Exit(result)
}

func help() {
	fmt.Println("Usage git hook <hook-name>")
}
