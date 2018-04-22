package main

import "os"
import "os/exec"

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	if len(os.Args) > 1 {
		pwd = os.Args[1]
	}
	cmd := exec.Command("/Applications/GitKraken.app/Contents/MacOS/GitKraken", "-p", pwd)
	cmd.Start()
}
