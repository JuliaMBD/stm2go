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

const config_file string = "stm2go.json"
const model_file string = "models.json"
const default_packagename string = "github.com/example/mypackage"
const default_xmlname string = "model.drawio"

// A function to check whether a file exists or not
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func usage() {
	msg := `usage: stm2go <command> [<args>]

commands: (command help: stm2go command -h)
  init      Initialize directories.
  generate  Generate Go source files (base and impl) from draw.io/diagrams.net file.
  main      Generate a template of main.go from draw.io/diagrams.net file.
  example   Generate an example of draw.io/diagrams.net file.
  help      Display this message.`

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
		case "main":
			cmdmain(args)
		case "example":
			cmdexample(args)
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
	var external bool
	flag.StringVar(&srcdir, "d", "src", "Name of src directory. Default is src.")
	flag.BoolVar(&force, "f", false, "Flag to overwrite Go sources excluding main file.")
	flag.BoolVar(&external, "e", false, "Flag to allow external transitions.")
	flag.CommandLine.Parse(args)

	// Read config
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

	// Read XML
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

	// Parse XML
	stms, states := stm2go.Parse(xml)
	stmap, sttree, root := stm2go.NewGoSTMMap(pkg, stms, states)

	// Define STM names
	if !fileExists(model_file) {
		fmt.Println("Create model file: " + model_file)
		tmp := make(map[string]string)
		b, err := json.Marshal(tmp)
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			os.Exit(1)
		}
		os.WriteFile(model_file, b, 0644)
	}
	if raw, err := os.ReadFile(model_file); err == nil {
		json.Unmarshal(raw, &names)
	} else {
		fmt.Println("Fail to read " + model_file)
		os.Exit(1)
	}
	for parent, stms := range sttree {
		for _, s := range stms {
			if _, ok := names[s.Id]; !ok {
				names[s.Id] = parent.Name + "Stm" + strconv.Itoa(len(names))
			}
		}
	}
	b, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		os.Exit(1)
	}
	os.WriteFile(model_file, b, 0644)

	// work
	var fn string

	// generate directory
	if err := os.MkdirAll(srcdir, 0777); err != nil {
		fmt.Println("Fail to create the directory")
		os.Exit(1)
	}

	// generate common file
	fn = srcdir + "/common.go"
	if force == true || !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
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
		fn = srcdir + "/" + names[s.Id] + "_base.go"
		if force == true || !fileExists(fn) {
			if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
				w := stm2go.NewWriter(f)
				s.BaseHeader(w)
				s.BaseStateDefinition(w, names)
				s.BaseStateInitialize(w, names)
				if external {
					s.BaseTransDefinitionWithExternal(w, names)
				} else {
					s.BaseTransDefinition(w, names)
				}
				s.UpdateDefinition(w, sttree, names)
				f.Close()
			} else {
				fmt.Println("Fail to create " + fn)
				os.Exit(1)
			}
		} else {
			fmt.Println("File " + fn + " exists. Skip creating.")
		}

		// impl
		fn = srcdir + "/" + names[s.Id] + "_impl.go"
		if force == true || !fileExists(fn) {
			if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
				w := stm2go.NewWriter(f)
				s.ImplHeader(w)
				if external {
					s.ImplFunctionsWithExternal(w, sttree, names)
				} else {
					s.ImplFunctions(w, sttree, names)
				}
				f.Close()
			} else {
				fmt.Println("Fail to create " + fn)
				os.Exit(1)
			}
		} else {
			fmt.Println("File " + fn + " exists. Skip creating.")
		}
	}

	// generate test file
	fn = srcdir + "/" + pkg.Pkgname + "_test.go"
	if force == true || !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
			w := stm2go.NewWriter(f)
			pkg.TestGen(w, sttree[root], names)
			f.Close()
		} else {
			fmt.Println("Fail to create " + fn)
			os.Exit(1)
		}
	} else {
		fmt.Println("File " + fn + " exists. Skip creating.")
	}

	fmt.Println("Code generation done.")
}

func cmdmain(args []string) {
	var force bool
	flag.BoolVar(&force, "f", false, "Flag to overwrite Go sources excluding main file.")
	flag.CommandLine.Parse(args)

	// Read config
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

	// Read XML
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

	// Parse XML
	stms, states := stm2go.Parse(xml)
	_, sttree, root := stm2go.NewGoSTMMap(pkg, stms, states)

	fn := "main.go"
	if force == true || !fileExists(fn) {
		if f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
			w := stm2go.NewWriter(f)
			pkg.GenMain(w, sttree[root], names)
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

func cmdexample(args []string) {
	var outfile string
	flag.StringVar(&outfile, "o", "example.drawio", "Name of the example of draw.io/diagras.net file.")
	flag.CommandLine.Parse(args)

	drawio := []byte(`<mxfile host="app.diagrams.net" modified="2022-12-13T05:07:33.635Z" agent="5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36" etag="_CkyxOyHdMj36E9PywLI" version="20.5.3" type="google"><diagram id="3Aa49lHglVZ3ABnJkFck" name="Page 1">7VjbbuIwEP0aHkG5lNtjobRbqRVdUWm3TytDnMStsSPHFOjXr+1MLgYKoUvZrrRIQObYHjszc47jNPzhfHUjUBLf8wDThucEq4Z/1fA8t9/rqj+NrAFxOm6GRIIEgJXAhLzhvCOgCxLg1OooOaeSJDY444zhmbQwJARf2t1CTu1ZExTBjE4JTGaI4q1uP0gg4wzted0S/4ZJFOczu51+1jJHeWdwnMYo4MsK5I8a/lBwLrOr+WqIqY5eHpds3PU7rcXCBGayzoDR4PJWjJaT77/Sce+mG9xfJLdN9wIWJ9f5HeNABQBMLmTMI84QHZXoQPAFC7B26yir7HPHeaJAV4HPWMo1ZBMtJFdQLOcUWvGKyJ+V6yftqtUG62oFno2xzg0mxfpn1aiM0mY5zFj5uO0wQeRSvhAzvCc2kEiJRITlvhhCgnXgKjNAFm4wn2O1INVBYIokebUrC0GBRkW/MofqAtJoTD591iXuORRNMc2GwRrXCXgljEiCaCqRxFnb3vz3KyuFeqmWAqWKZzrly5hIPEmQCdhScd1OKEqTjHwhWenCGEQUpSlkIH3BchaDERJKh5xyYWbwHfPZl6ZXLCRe7Y0rtDY7DmSiEBuwlyVzfYDiCmlz7IhUKBOycQzZ2v/JdohEh9nmnYltNXPagzJFdAFTPS4EG7Odqb7LiGtRh5KIqeuZChlWrBjogidqA7qEhjkJgqwScEre0NT408FOOGHS3E970Ghf7Qz/3jrcolaxjcIs1k61i3JOy3dd1yZdZtWOPvh+0DdT4bLbt7wWG1XugodhqupkM33FEuvq5x0OxmG4qaJ15dOthmdbP9MlmVPEbBq7h+X0XfU8hUq2XVslu+0tkXR3qWTvXCrZ+Zsq6Vgq2a0pk64tk91P1Ml8xz//U8k7TC0KA+rJ62wUSrZUGHYMX+vVS0GBDQXOWf01JbhzKgluOi2nl3P4tBpsS/BnKjD7sAB75xVgtS4hQUQ8/0SK7PW/tiIXJ9gvoMgHBFkZD1gQFQZN5I1n2fLx9clS5d0ajVlwqV8gKJNxU0IKuSY6kqY9QGlc1NMfCPo/ccx8jEmqy0T/IPVlXOLWJmk1WIez+w+daYyMQ+PtIFOnaPYSmVobLyQlOk8GD5B4GatRRJp8tpz2EafS0HxOxO6LjVNps7d9Ku3tYHdxev04vZVZvl/K1Ll8TeePfgM=</diagram></mxfile>`)

	os.WriteFile(outfile, drawio, 0644)
}
