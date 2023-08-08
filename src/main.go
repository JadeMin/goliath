package main

import (
	os "os"
	path "path/filepath"

	arger "github.com/alexflint/go-arg"
);

func clean() {
	var goPath string = os.Getenv("GOPATH");

	if len(goPath) == 0 {
		println("WARNING: GOPATH is not set!");
	} else {
		var modPath string = path.Join(goPath, "\\pkg\\mod");
		if err := os.RemoveAll(modPath); err != nil {
			println(err);
		}
	}
};



type Args struct {
	Command		string	`arg:"positional,required" help:"Command to execute"`
	//WorkingDir	string	`arg:"-d" help:"Working directory"`
};

func main() {
	var args Args;
	//args.WorkingDir = "./src";
	arger.MustParse(&args);

	if args.Command == "clean" {
		clean();
	} else {
		println("Unknown command: " + args.Command);
	}
};