package stm2go

import (
	"fmt"
	"os"
	"testing"
)

func TestMakeTransitionMap(t *testing.T) {
	ss := []*State{
		&State{"A"},
		&State{"B"},
	}
	ts := []*Transition{
		&Transition{
			Src:   ss[0],
			Dest:  ss[1],
			Event: &Event{"EventA"},
		},
	}
	m := makeTransitionMap(ts)
	fmt.Println(ss)
	fmt.Println(ts)
	fmt.Println(m)
}

func TestGenHeader(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	gs.baseHeader()
}

func TestGenStateDefinition(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	states := []*State{
		&State{"A"},
		&State{"B"},
	}
	gs.baseStateDefinition("stm1", states)
}

func TestGenInitState(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	states := []*State{
		&State{"A"},
		&State{"B"},
	}
	gs.baseStateInitialize("stm1", states[0])
}

func TestGenTrans(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	states := []*State{
		&State{"A"},
		&State{"B"},
	}
	trans := map[*State][]*Transition{
		states[0]: []*Transition{
			&Transition{
				Src:   states[0],
				Dest:  states[1],
				Event: &Event{"EventA"},
			},
		},
	}
	gs.baseTransDefinition("stm1", states, trans)
}

func TestGenSTM(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	states := []*State{
		&State{"A"},
		&State{"B"},
	}
	trans := map[*State][]*Transition{
		states[0]: []*Transition{
			&Transition{
				Src:   states[0],
				Dest:  states[1],
				Event: &Event{"EventA"},
			},
		},
	}
	gs.baseHeader()
	gs.baseStateDefinition("stm1", states)
	gs.baseStateInitialize("stm1", states[0])
	gs.baseTransDefinition("stm1", states, trans)
}

func TestGenCommon(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	gs.common()
}

func TestGenFunc(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	states := []*State{
		&State{"A"},
		&State{"B"},
	}
	trans := map[*State][]*Transition{
		states[0]: []*Transition{
			&Transition{
				Src:   states[0],
				Dest:  states[1],
				Event: &Event{"EventA"},
			},
		},
	}
	gs.implFunctions(states, trans)
}

func TestGenTest(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	gs.testGen("stm1")
}

func TestGenMain(t *testing.T) {
	gs := NewGoSource(os.Stdout, "com.github/example", "test")
	gs.testMain("test")
}
