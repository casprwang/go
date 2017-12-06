package main

import (
	"fmt"
	"os"
	"runtime"
)

func UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func main() {
	dir := UserHomeDir()
	fmt.Println(dir)
}
