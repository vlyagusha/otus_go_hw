package main

import "os"

func main() {
	environment, err := ReadDir(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	os.Exit(RunCmd(os.Args[2:], environment))
}
