package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	stm2go "github.com/JuliaMBD/stm2go/src"
)

var (
	pkg    *stm2go.GoPkgSource
	srcdir string
	names  map[string]string
)

const config_file string = "package.json"
const default_packagename string = "github.com/example/mypackage"
const default_xmlname string = "model.xml"

func init() {
	names = make(map[string]string)
	// default model names (to be modified)
	for i := 1; i < 100; i++ {
		names[strconv.Itoa(i)] = "stm" + strconv.Itoa(i)
	}
}

// A function to check whether a file exists or not
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func usage() {
	msg := `usage: stm2go <command> [<args>]

commands: (command help: stm2go command -h)
  init      Initialize directories
  generate  Generate Go source files
  help      Display this message`

	fmt.Println(msg)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	} else {
		mode := os.Args[1]
		args := os.Args[2:]
		switch mode {
		case "init":
			cmdinit(args)
		case "generate":
			cmdgen(args)
		case "help":
			usage()
		default:
			usage()
		}
	}
}

func cmdinit(args []string) {
	flag.CommandLine.Parse(args)
	fullpkgname := flag.Arg(0)
	if fullpkgname == "" {
		fmt.Println("Use the default package name '" + default_packagename + "'")
		fullpkgname = default_packagename
	}

	pkgname := filepath.Base(fullpkgname)
	domain := filepath.Dir(fullpkgname)
	pkg = stm2go.NewGoPkgSource(domain, pkgname)

	if !fileExists(config_file) {
		if f, err := json.MarshalIndent(pkg, "", "  "); err == nil {
			os.WriteFile(config_file, f, 0644)
		} else {
			fmt.Println("Fail to create JSON")
			os.Exit(1)
		}
	} else {
		fmt.Println(config_file + " exists. Skip creating.")
	}
}

func cmdgen(args []string) {
	var force bool
	var mainfn string
	flag.StringVar(&srcdir, "d", "src", "Name of src directory. Default is src.")
	flag.BoolVar(&force, "f", false, "Flag to overwrite Go sources excluding main file.")
	flag.StringVar(&mainfn, "main", "main.go", "Name of main file. Default is main.go.")
	flag.CommandLine.Parse(args)

	if !fileExists(config_file) {
		fmt.Println("Did not find " + config_file + ".")
		fmt.Println("Probably this is not a directory for stm2go.")
		os.Exit(1)
	}

	if raw, err := os.ReadFile(config_file); err == nil {
		json.Unmarshal(raw, &pkg)
	} else {
		fmt.Println("Fail to read " + config_file)
		os.Exit(1)
	}

	infile := flag.Arg(0)
	if infile == "" {
		fmt.Println("The name of xml file is not detected. Use the default xmlfile '" + default_xmlname + "'")
		infile = default_packagename
	}

	if !fileExists(infile) {
		fmt.Println("Did not find XML file: " + infile)
		os.Exit(1)
	}

	xml, err := os.ReadFile(infile)
	if err != nil {
		fmt.Println("Fail to read XML file: " + infile)
		os.Exit(1)
	}

	stms, states := stm2go.Parse(xml)
	stmap, sttree, root := stm2go.NewGoSTMMap(pkg, names, stms, states)

	var fn string
	var entryname string
	if s, ok := sttree[root]; ok && len(s) == 1 {
		entryname = s[0].Name
	} else {
		fmt.Println("Error: There are two entry points")
		os.Exit(1)
	}

	// generate directory
	if err := os.MkdirAll(srcdir, 0777); err != nil {
		fmt.Println("Fail to create the directory")
		os.Exit(1)
	}

	// generate common file
	fn = srcdir + "/common.go"
	if force == true || !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664); err == nil {
			w := stm2go.NewWriter(f)
			pkg.Common(w)
			f.Close()
		} else {
			fmt.Println("Fail to create " + fn)
			os.Exit(1)
		}
	} else {
		fmt.Println("File " + fn + " exists. Skip creating.")
	}

	// generate base and impl files
	for _, s := range stmap {
		// base
		fn = srcdir + "/" + s.Name + "_base.go"
		if force == true || !fileExists(fn) {
			if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664); err == nil {
				w := stm2go.NewWriter(f)
				s.BaseHeader(w)
				s.BaseStateDefinition(w)
				s.BaseStateInitialize(w)
				s.BaseTransDefinition(w)
				f.Close()
			} else {
				fmt.Println("Fail to create " + fn)
				os.Exit(1)
			}
		} else {
			fmt.Println("File " + fn + " exists. Skip creating.")
		}

		// impl
		fn = srcdir + "/" + s.Name + "_impl.go"
		if force == true || !fileExists(fn) {
			if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664); err == nil {
				w := stm2go.NewWriter(f)
				s.ImplHeader(w)
				s.ImplFunctions(w, sttree)
				f.Close()
			} else {
				fmt.Println("Fail to create " + fn)
				os.Exit(1)
			}
		} else {
			fmt.Println("File " + fn + " exists. Skip creating.")
		}
	}

	// generate test and main files
	fn = srcdir + "/" + pkg.Pkgname + "_test.go"
	if force == true || !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664); err == nil {
			w := stm2go.NewWriter(f)
			pkg.TestGen(w, entryname)
			f.Close()
		} else {
			fmt.Println("Fail to create " + fn)
			os.Exit(1)
		}
	} else {
		fmt.Println("File " + fn + " exists. Skip creating.")
	}

	fn = mainfn
	if !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664); err == nil {
			w := stm2go.NewWriter(f)
			pkg.GenMain(w, entryname)
			f.Close()
		} else {
			fmt.Println("Fail to create " + fn)
			os.Exit(1)

		}
	} else {
		fmt.Println("File " + fn + " exists. Skip creating.")
	}

	fmt.Println("Code generation done.")

	fmt.Println("Please execute:")
	fmt.Println("    go mod init " + pkg.Fullpkgname + "; go mod tidy")
}
