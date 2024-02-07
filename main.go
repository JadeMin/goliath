package main

import (
	fmt "fmt"
	os "os"
	path "path/filepath"

	arger "github.com/alexflint/go-arg"
)

func clean() {
	goRoot := os.Getenv("GOROOT")

	if len(goRoot) == 0 {
		println("WARNING: GOROOT is not set!")
	} else {
		modPath := path.Join(goRoot, "/pkg/mod")
		err := os.RemoveAll(modPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}



type Args struct {
	Command			string		`arg:"positional,required" help:"Command to execute"`
	//WorkingDir	string		`arg:"-d" help:"Working directory"`
}

func main() {
	var args Args
	arger.MustParse(&args)

	if args.Command == "clean" {
		clean()
	} else {
		fmt.Printf("Unknown command: %s\n", args.Command)
	}
}