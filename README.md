# stm2go

A go package for code generation from state machine diagrams

## Installation

```sh
go install github.com/JuliaMBD/stm2go
```

## Usage

### Initialize

```sh
stm2go init [<fullpath of package>]
```

Intializes the current directory for stm2go, and create a JSON file (stm2go.json) to save the name of package. If you omit the fullpath, the default is `github.com/example/mypackage`.

### Draw a state machine diagram


```sh
stm2go example [-o <outfile name>]
```

Create an example of draw.io/diagrams.net file. The default is to create example.drawio that contain examples of states, initial state and edges in draw.io. By copying them, you can draw your state machine diagram.

### Generate Go sources

```sh
stm2go generate [-f] [-d <src directory>] [--main=<main file name>] [<draw.io file>]
```

Generate Go source from a given draw.io/diagrams.net file. The sources excluding the main file are generated in the src directory. The default name of src directory is `src`. If there are the files having the same names, the tool skips to overwite them. The option `-f` indicates the overwriting. The main file `main.go` is never overwritten. If you create a different name of main file, please use the option `--main`. An example of directory stracture generated by the tool is

```
mypackage/
├── example.drawio
├── stm2go.json
├── main.go
└── src
    ├── common.go
    ├── mypackage_test.go
    ├── stm1_base.go
    ├── stm1_impl.go
    ├── stm3_base.go
    └── stm3_impl.go
```

- `main.go`: A template for main
- `common.go`: A go source including common settings such as logger, debug flag and EOD
- `xxxx_base.go`: A go source to control state transitions for a STM. There is no necessary to edit this file by users.
- `xxxx_impl.go`: A template for processes on a STM. Users should edit this file.
- `xxxx_test.go`: A template for testing

### Go module

To use the generated sources, we run go module system once with the following:

```sh
go mod init <fullpath of package>
go mod tidy
```

### Testing

By default, we can check state transitions on STMs by the following commands:

```sh
cd src
go test
```
The tool provides a testing framework based on golang testing. An example of test codes is
```go
package mypackage

import (
	"log"
	"testing"
	"time"

	stm2go "github.com/JuliaMBD/stm2go/testing"
)

type DebugStruct struct{}

func (l DebugStruct) Println(s string) {
	log.Println(s)
}

func TestExample(t *testing.T) {
	ConfigureLog(DebugStruct{})

	env := stm2go.NewTestEnv()

	env.Add(stm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
			stm1Task()
		}
	})
	env.Add(stm2go.Done, func() {
		time.Sleep(100 * time.Millisecond) // 10 times runs
	})
	env.Set(1)
	env.Go()
}
```

The struct `DebugStruct` is an instance of the interface
```go
type DebugLogger interface {
	Println(string)
}
```
This is called when the state changes. `ConfigureLog` is the function to regist the DebugLogger.

The tool provides `TestEnv` to constract the testing environemnt that controls multi-threading (Goroutine). The methods of TestEnv are
- `Add`: Register a function that runs as a Goroutine. There are two types of routines; `stm2go.Continue` and `stm2go.Done`. The task with `stm2go.Continue` runs countinuously. On the other hand, the task with `stm2go.Done` has the end of process. In the testing, the test stops when all the tasks with `stm2go.Done` are done (ignores whether the task with `stm2go.Continue` ends or not).
- `Go`: Running the tasks that are registered.
- `Set`: Set timescale.
- `Sleep`: Sleep
- `Tick`: ?
- `After`: ?

Typically, realtime processing is required in embeddeing systems. The developed testing framework simulates realtime processing with Goroutine. For example, one task simulates the poling (event loop) for embeded system and another task simulates the user behavior such as pushing buttons.

## State Machine Diagram

State machine diagram is drawn with [draw.io/diagrams.net](https://www.diagrams.net/). The objects in diagrams have specified attributes to parse it.

### State

To be written

### Transition

To be written

### Initial state

To be written

### Composite states

To be written

