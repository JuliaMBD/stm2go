package stm2go

import (
	"io"
)

const srcdir string = "src"

type GoPkgSource struct {
	Fullpkgname string `json: "fullpkgname"`
	Domain      string `json: "domain"`
	Pkgname     string `json: "pkgname"`
}

type GoSTMSource struct {
	Id      string
	states  map[*State]struct{}      // set for states to be used in hasState?
	ts      map[*State][]*Transition // local transitions
	inexts  map[*State][]*Transition // inbound external transitions
	outexts map[*State][]*Transition // outbound external transitions
	initial *State                   // initial state
	pkg     *GoPkgSource             // pakcage
	root    bool                     // indicater the STM is root or not
}

// A utility function to make a unique order for states
func (g *GoSTMSource) sortState() []*State {
	ss := make([]*State, 0, len(g.states))
	for k := range g.states {
		ss = append(ss, k)
	}
	return ss
}

// A function to create GoSource
//   out: io.Writer (e.g. os.Stdout)
//   domain: domain name for package
//   pkgname: package name
func NewGoPkgSource(domain string, pkgname string) *GoPkgSource {
	return &GoPkgSource{
		Fullpkgname: domain + "/" + pkgname,
		Domain:      domain,
		Pkgname:     pkgname,
	}
}

// A function to create GoSTMSource
//   out: io.Writer
//   name: name of STM
//   ss: Slice of States
//   states: Set for states
//   ts: Slice of Transitions
//   ex: Slice of ExTransitions
//   initial: Initial State
//   pkg: GoSource for Pkg
//   root: Indicator whether the stm is root or not
func NewGoSTMSource(name string,
	ss []*State, ts []*Transition, ex []*Transition,
	initial *State, pkg *GoPkgSource, root bool) *GoSTMSource {
	states := make(map[*State]struct{})
	for _, s := range ss {
		states[s] = struct{}{}
	}
	inex := make(map[*State][]*Transition)
	outex := make(map[*State][]*Transition)
	for _, t := range ex {
		if _, ok := states[t.Src]; ok {
			if _, ok := outex[t.Src]; ok {
				outex[t.Src] = append(outex[t.Src], t)
			} else {
				outex[t.Src] = []*Transition{t}
			}
		} else if _, ok := states[t.Dest]; ok {
			if _, ok := inex[t.Dest]; ok {
				inex[t.Dest] = append(inex[t.Dest], t)
			} else {
				inex[t.Dest] = []*Transition{t}
			}
		}
	}
	return &GoSTMSource{
		Id:      name,
		states:  states,
		ts:      makeTransitionMap(ts),
		inexts:  inex,
		outexts: outex,
		initial: initial,
		pkg:     pkg,
		root:    root,
	}
}

// A function to make map between a parent state and a state machine
func NewGoSTMMap(pkg *GoPkgSource, stms map[string]*StateMachine, states map[string]*State) ([]*GoSTMSource, map[*State][]*GoSTMSource, *State) {
	stmap := make([]*GoSTMSource, 0)
	sttree := make(map[*State][]*GoSTMSource)
	root := &State{Name: "Root"}
	for k, s := range stms {
		st := NewGoSTMSource(k, s.States, s.Transitions, s.ExTransitions, s.Initial, pkg, false)
		stmap = append(stmap, st)
		if p, ok := states[s.Parent]; ok {
			if _, ok := sttree[p]; ok {
				sttree[p] = append(sttree[p], st)
			} else {
				sttree[p] = []*GoSTMSource{st}
			}
		} else {
			if _, ok := sttree[root]; ok {
				sttree[root] = append(sttree[root], st)
			} else {
				sttree[root] = []*GoSTMSource{st}
			}
		}
	}
	for i := range sttree[root] {
		sttree[root][i].root = true
	}
	return stmap, sttree, root
}

type Writer struct {
	out io.Writer
}

func NewWriter(out io.Writer) *Writer {
	return &Writer{
		out: out,
	}
}

// A function to write source
func (w *Writer) writeln(s string) {
	w.out.Write([]byte(s + "\n"))
}

// A function to make the map for transitions
func makeTransitionMap(ts []*Transition) map[*State][]*Transition {
	m := make(map[*State][]*Transition)
	for _, t := range ts {
		if _, ok := m[t.Src]; ok {
			m[t.Src] = append(m[t.Src], t)
		} else {
			m[t.Src] = []*Transition{t}
		}
	}
	return m
}

// A function to generate the string of header
func (g *GoSTMSource) BaseHeader(w *Writer) {
	w.writeln("// This file was generated by a program.")
	w.writeln("// Please do not edit this file directly.")
	w.writeln("package " + g.pkg.Pkgname + "\n")
	w.writeln("import (\n// package names to be imported\n)\n")
}

// A function to generate Enum for states
func (g *GoSTMSource) BaseStateDefinition(w *Writer, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	w.writeln("type " + stm + "State uint8")
	w.writeln("const (")
	for i, s := range ss {
		if i == 0 {
			w.writeln(stm + s.Name + " " + stm + "State = iota")
		} else {
			w.writeln(stm + s.Name)
		}
	}
	w.writeln(")\n")
	w.writeln("var " + stm + "Eod Eod")
	w.writeln("var " + stm + "CurrentState " + stm + "State")
	w.writeln("var " + stm + "NextState " + stm + "State\n")
}

// A function to generate init function
func (g *GoSTMSource) BaseStateInitialize(w *Writer, names map[string]string) {
	stm := names[g.Id]
	i := stm + g.initial.Name
	w.writeln("func init() {")
	w.writeln(stm + "Initialize()")
	w.writeln("}\n")
	w.writeln("func " + stm + "Initialize() {")
	w.writeln(stm + "Eod = Entry")
	w.writeln(stm + "CurrentState = " + i)
	w.writeln(stm + "NextState = " + i)
	w.writeln("}\n")
}

// A function to generate base for a given STM
func (g *GoSTMSource) BaseTransDefinition(w *Writer, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	ts := g.ts
	if g.root {
		w.writeln("func Entry" + stm + "Task() {")
		w.writeln(stm + "Task()")
		w.writeln(stm + "Update()")
		w.writeln("}\n")
	}
	w.writeln("func " + stm + "Task() {")
	w.writeln("switch " + stm + "CurrentState {")
	for _, s := range ss {
		w.writeln("case " + stm + s.Name + ":")
		w.writeln("if " + stm + "Eod == Entry {")
		w.writeln(stm + s.Name + "Entry()")
		w.writeln(stm + "Eod = Do")
		w.writeln("}")
		w.writeln("if " + stm + "Eod == Do {")
		w.writeln(stm + s.Name + "Do()")
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("if " + tr + "Cond() {")
			w.writeln(tr + "Action()")
			w.writeln(stm + "NextState = " + stm + t.Dest.Name)
			w.writeln(stm + "Eod = Exit")
			w.writeln("}")
		}
		w.writeln("}")
		w.writeln("if " + stm + "Eod == Exit {")
		w.writeln(stm + s.Name + "Exit()")
		w.writeln(stm + "Eod = Entry")
		w.writeln("}")
	}
	w.writeln("}")
	w.writeln("}\n")
}

// A function to generate base for a given STM
func (g *GoSTMSource) BaseTransDefinitionWithExternal(w *Writer, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	ts := g.ts
	outts := g.outexts
	if g.root {
		w.writeln("func Entry" + stm + "Task() {")
		w.writeln(stm + "Task()")
		w.writeln(stm + "Update()")
		w.writeln("}\n")
	}
	w.writeln("func " + stm + "Task() {")
	w.writeln("switch " + stm + "CurrentState {")
	for _, s := range ss {
		w.writeln("case " + stm + s.Name + ":")
		w.writeln("if " + stm + "Eod == Entry {")
		w.writeln(stm + s.Name + "Entry()")
		w.writeln(stm + "Eod = Do")
		w.writeln("}")
		w.writeln("if " + stm + "Eod == Do {")
		w.writeln(stm + s.Name + "Do()")
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("if " + tr + "Cond() {")
			w.writeln(tr + "Action()")
			w.writeln(stm + "NextState = " + stm + t.Dest.Name)
			w.writeln(stm + "Eod = Exit")
			w.writeln("}")
		}
		w.writeln("}")
		for _, t := range outts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("// external")
			w.writeln("if " + tr + "Cond() {")
			w.writeln(tr + "Action()")
			w.writeln(stm + "NextState = " + stm + g.initial.Name)
			w.writeln(stm + "Eod = Exit")
			stm2 := names[t.Dest.Stm.Parent]
			w.writeln(stm2 + "NextState = " + stm2 + t.Dest.Name)
			w.writeln(stm2 + "Eod = Exit")
			w.writeln("}")
		}
		w.writeln("if " + stm + "Eod == Exit {")
		w.writeln(stm + s.Name + "Exit()")
		w.writeln(stm + "Eod = Entry")
		w.writeln("}")
	}
	w.writeln("}")
	w.writeln("}\n")
}

// A function to generate base for a given STM
func (g *GoSTMSource) UpdateDefinition(w *Writer, sttree map[*State][]*GoSTMSource, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	w.writeln("func " + stm + "Update() {")
	w.writeln("switch " + stm + "CurrentState {")
	for _, s := range ss {
		w.writeln("case " + stm + s.Name + ":")
		if stms, ok := sttree[s]; ok {
			for _, stm := range stms {
				w.writeln(names[stm.Id] + "Update() // Call the update for " + names[stm.Id])
			}
		}
	}
	w.writeln("}")
	w.writeln(stm + "CurrentState = " + stm + "NextState")
	w.writeln("}\n")
}

// A function to generate the string of header for impl
func (g *GoSTMSource) ImplHeader(w *Writer) {
	w.writeln("// This file was generated by a program.")
	w.writeln("package " + g.pkg.Pkgname)
	w.writeln("import (\n// package names to be imported\n)")
}

// A function to generate template functions
func (g *GoSTMSource) ImplFunctions(w *Writer, sttree map[*State][]*GoSTMSource, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	ts := g.ts
	for _, s := range ss {
		w.writeln("///////////////////////////////////////////////")
		w.writeln("// functions for State " + stm + s.Name)
		w.writeln("///////////////////////////////////////////////\n")
		w.writeln("func " + stm + s.Name + "Entry() {")
		if stms, ok := sttree[s]; ok {
			for _, stm := range stms {
				w.writeln(names[stm.Id] + "Initialize() // Call the initialize for " + names[stm.Id])
			}
		}
		w.writeln("if debug {")
		w.writeln("logger.Println(\"Entering State " + stm + s.Name + "\")")
		w.writeln("}")
		if s.Entry == "" {
			w.writeln("// Please write an enter process for State " + stm + s.Name)
		} else {
			w.writeln(s.Entry)
		}
		w.writeln("}\n")
		w.writeln("func " + stm + s.Name + "Do() {")
		if stms, ok := sttree[s]; ok {
			for _, stm := range stms {
				w.writeln(names[stm.Id] + "Task() // Call the task for " + names[stm.Id])
			}
		}
		if s.Do == "" {
			w.writeln("// Please write a do process for State " + stm + s.Name)
		} else {
			w.writeln(s.Do)
		}
		w.writeln("}\n")
		w.writeln("func " + stm + s.Name + "Exit() {")
		w.writeln("if debug {")
		w.writeln("logger.Println(\"Leaving State " + stm + s.Name + "\")")
		w.writeln("}")
		if s.Exit == "" {
			w.writeln("// Please write an exit process for State " + stm + s.Name)
		} else {
			w.writeln(s.Exit)
		}
		w.writeln("}\n")
	}
	w.writeln("///////////////////////////////////////////////")
	w.writeln("// Conditions")
	w.writeln("///////////////////////////////////////////////\n")
	for _, s := range ss {
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("func " + tr + "Cond() bool {")
			if t.Event.Cond == "" {
				w.writeln("// Please edit the condition")
				w.writeln("return true")
			} else {
				w.writeln("return " + t.Event.Cond)
			}
			w.writeln("}\n")
		}
	}
	w.writeln("///////////////////////////////////////////////")
	w.writeln("// Actions")
	w.writeln("///////////////////////////////////////////////\n")
	for _, s := range ss {
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("func " + tr + "Action() {")
			if t.Event.Action == "" {
				w.writeln("// Please edit the action when " + t.Event.Name + " occurs at State " + s.Name)
			} else {
				w.writeln(t.Event.Action)
			}
			w.writeln("}\n")
		}
	}
}

// A function to generate template functions
func (g *GoSTMSource) ImplFunctionsWithExternal(w *Writer, sttree map[*State][]*GoSTMSource, names map[string]string) {
	stm := names[g.Id]
	ss := g.sortState()
	ts := g.ts
	outts := g.outexts
	for _, s := range ss {
		w.writeln("///////////////////////////////////////////////")
		w.writeln("// functions for State " + stm + s.Name)
		w.writeln("///////////////////////////////////////////////\n")
		w.writeln("func " + stm + s.Name + "Entry() {")
		if stms, ok := sttree[s]; ok {
			for _, stm := range stms {
				w.writeln(names[stm.Id] + "Initialize() // Call the initialize for " + names[stm.Id])
			}
		}
		w.writeln("if debug {")
		w.writeln("logger.Println(\"Entering State " + stm + s.Name + "\")")
		w.writeln("}")
		w.writeln("// Please write an enter process for State " + stm + s.Name)
		w.writeln("}\n")
		w.writeln("func " + stm + s.Name + "Do() {")
		if stms, ok := sttree[s]; ok {
			for _, stm := range stms {
				w.writeln(names[stm.Id] + "Task() // Call the task for " + names[stm.Id])
			}
		}
		w.writeln("// Please write a do process for State " + stm + s.Name)
		w.writeln("}\n")
		w.writeln("func " + stm + s.Name + "Exit() {")
		w.writeln("if debug {")
		w.writeln("logger.Println(\"Leaving State " + stm + s.Name + "\")")
		w.writeln("}")
		w.writeln("// Please write an exit process for State " + stm + s.Name)
		w.writeln("}\n")
	}
	w.writeln("///////////////////////////////////////////////")
	w.writeln("// Conditions")
	w.writeln("///////////////////////////////////////////////\n")
	for _, s := range ss {
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("func " + tr + "Cond() bool {")
			w.writeln("// Please edit the condition")
			w.writeln("return true")
			w.writeln("}\n")
		}
		for _, t := range outts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("// external (outbound)")
			w.writeln("func " + tr + "Cond() bool {")
			w.writeln("// Please edit the condition")
			w.writeln("return true")
			w.writeln("}\n")
		}
	}
	w.writeln("///////////////////////////////////////////////")
	w.writeln("// Actions")
	w.writeln("///////////////////////////////////////////////\n")
	for _, s := range ss {
		for _, t := range ts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("func " + tr + "Action() {")
			w.writeln("// Please edit the action when " + t.Event.Name + " occurs at State " + s.Name)
			w.writeln("}\n")
		}
		for _, t := range outts[s] {
			tr := stm + s.Name + t.Event.Name
			w.writeln("// external")
			w.writeln("func " + tr + "Action() {")
			w.writeln("// Please edit the action when " + t.Event.Name + " occurs at State " + s.Name)
			w.writeln("}\n")
		}
	}
}

// A function to generate Common
func (g *GoPkgSource) Common(w *Writer) {
	w.writeln("// This file was generated by a program.")
	w.writeln("package " + g.Pkgname + "\n")
	w.writeln("import (\n// package names to be imported\n)\n")
	w.writeln("type Eod uint8\n")
	w.writeln("const (")
	w.writeln("Entry Eod = iota")
	w.writeln("Do")
	w.writeln("Exit")
	w.writeln(")\n")
	w.writeln("type DebugLogger interface {")
	w.writeln("Println(...interface{})")
	w.writeln("}\n")
	w.writeln("var logger DebugLogger\n")
	w.writeln("func ConfigureLogger(d DebugLogger) {\nlogger = d\n}\n")
	w.writeln("const (\ndebug = true\n)\n")
	w.writeln("// func InputDevice() {\n// Please edit to get status of input devices\n// }\n")
	w.writeln("// func OutputDevice() {\n// Please edit to put status of output devices\n// }\n")
}

// A function to generate an example of test code
func (g *GoPkgSource) TestGen(w *Writer, stms []*GoSTMSource, names map[string]string) {
	w.writeln("// This file was generated by a program.")
	w.writeln("package " + g.Pkgname + "\n")
	w.writeln("import (\n\"log\"\n\"testing\"\n\"time\"\n stm2go \"github.com/JuliaMBD/stm2go/testing\"\n)\n")
	w.writeln("type TestLogger struct{}\n")
	w.writeln("func (l TestLogger) Println(s ...interface{}) {")
	w.writeln("log.Println(s...)")
	w.writeln("}\n")
	w.writeln("func " + "TestExample(t *testing.T) {")
	w.writeln("ConfigureLogger(TestLogger{})\n")
	w.writeln("env := stm2go.NewTestEnv()\n")
	w.writeln("env.Add(stm2go.Continue, func() {")
	w.writeln("for {")
	w.writeln("// InputDevice()")
	for _, stm := range stms {
		w.writeln("Entry" + names[stm.Id] + "Task()")
	}
	w.writeln("// OutputDevice()")
	w.writeln("time.Sleep(10 * time.Millisecond)")
	w.writeln("}")
	w.writeln("})")
	w.writeln("env.Add(stm2go.Done, func() {")
	w.writeln("time.Sleep(100 * time.Millisecond) // 10 times runs")
	w.writeln("})")
	w.writeln("env.Set(1)")
	w.writeln("env.Go()")
	w.writeln("}")
}

// A function to generate an example of main
func (g *GoPkgSource) GenMain(w *Writer, stms []*GoSTMSource, names map[string]string) {
	w.writeln("package main\n")
	w.writeln("import (\"time\"\n" + g.Pkgname + " \"" + g.Fullpkgname + "/" + srcdir + "\"\n)\n")
	w.writeln("func main() {")
	w.writeln("for {")
	w.writeln("// " + g.Pkgname + ".InputDevice()")
	for _, stm := range stms {
		w.writeln(g.Pkgname + ".Entry" + names[stm.Id] + "Task()")
	}
	w.writeln("// " + g.Pkgname + ".OutputDevice()")
	w.writeln("time.Sleep(time.Millisecond * 10)")
	w.writeln("}")
	w.writeln("}")
}
