package main

import "os"

func main() {
	environment, _ := ReadDir(os.Args[1])
	os.Exit(RunCmd(os.Args[2:], environment))
}
