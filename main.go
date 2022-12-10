package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	stm2go "com.github/JuliaMBD/stm2go/src"
)

func usage() {
	msg := `usage: stm2go <command> [<args>]

commands: (command help: stm2go command -h)
  gen     Generate Go source files
  help    Display this message`

	fmt.Println(msg)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	} else {
		mode := os.Args[1]
		args := os.Args[2:]
		switch mode {
		case "gen":
			cmdgen(args)
		case "help":
			usage()
		default:
			usage()
		}
	}
}

func cmdgen(args []string) {
	infile := flag.String("f", "", "State machine file (XML file of draw.io/diagrams.net")
	fullpkgname := flag.String("p", "com.github/example/package", "Full package name. Default is \"com.github/example/package\"")
	flag.CommandLine.Parse(args)

	pkgname := filepath.Base(*fullpkgname)
	domain := filepath.Dir(*fullpkgname)

	var xml []byte
	if *infile != "" {
		if b, err := os.ReadFile(*infile); err == nil {
			xml = b
		} else {
			panic(err)
		}
	} else {
		if b, err := io.ReadAll(os.Stdin); err == nil {
			xml = b
		} else {
			panic(err)
		}
	}

	stm2go.Generate(domain, pkgname, xml)
}
