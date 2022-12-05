package stm2go

import (
	"log"
	"os"
)

var logger *log.Logger

type StateMachine struct {
	Parent      string
	States      []*State
	Events      []*Event
	Transitions []*Transition
	Initial     *State
}

func NewStateMachine(parent string) *StateMachine {
	return &StateMachine{
		Parent:      parent,
		States:      make([]*State, 0),
		Events:      make([]*Event, 0),
		Transitions: make([]*Transition, 0),
		Initial:     nil,
	}
}

type State struct {
	Name string
}

type Event struct {
	Name string
}

type Transition struct {
	Src   *State
	Dest  *State
	Event *Event
}

func Parse(data []byte) map[string]*StateMachine {
	logger = log.New(os.Stdout, "[Parse] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	// logger.SetOutput(io.Discard)
	states := make(map[string]*State)
	mxstates := make(map[string]*MxElement)
	events := make(map[string]*Event)
	edges := make(map[string]*Event)
	istates := make(map[string]*MxElement)
	transitions := make([]*MxElement, 0)
	elems, err := GetGraphModel(data)
	if err != nil {
		logger.Panic(err)
	}
	sm := make(map[string]*StateMachine)
	for i, x := range elems {
		switch x.Type {
		case "state":
			logger.Printf("Detect a state: %s-%s %s\n", x.Parent, x.Id, x.Value)
			p := x.Parent
			if _, ok := sm[p]; ok == false {
				sm[p] = NewStateMachine(p)
			}
			s := &State{x.Value}
			sm[p].States = append(sm[p].States, s)
			states[x.Id] = s
			mxstates[x.Id] = &elems[i]
		case "initialstate":
			logger.Printf("Detect an initial state: %s-%s %s\n", x.Parent, x.Id, x.Value)
			istates[x.Id] = &elems[i]
		case "vertex":
			switch x.getNode() {
			case "edgeLabel":
				logger.Printf("Detect an event: %s\n", x.Value)
				if e, ok := events[x.Value]; ok {
					edges[x.Parent] = e
				} else {
					e := &Event{x.Value}
					events[x.Value] = e
					edges[x.Parent] = e
				}
			default:
			}
		case "edge":
			transitions = append(transitions, &elems[i])
		default:
		}
	}

	// inner struct
	type key struct {
		s *StateMachine
		e *Event
	}
	smevents := make(map[key]struct{})

	for _, x := range transitions {
		src, ok1 := mxstates[x.Properties["source"]]
		dest, ok2 := mxstates[x.Properties["target"]]
		isrc, ok3 := istates[x.Properties["source"]]
		if ok1 && ok2 && src.Parent == dest.Parent && src.Parent == x.Parent {
			s := sm[src.Parent]
			t := &Transition{
				Src:   states[src.Id],
				Dest:  states[dest.Id],
				Event: edges[x.Id],
			}
			smevents[key{s, edges[x.Id]}] = struct{}{}
			s.Transitions = append(s.Transitions, t)
		} else if ok3 && ok2 && isrc.Parent == dest.Parent && isrc.Parent == x.Parent {
			// set initital
			s := sm[isrc.Parent]
			s.Initial = states[dest.Id]
		}
	}
	for k, _ := range smevents {
		k.s.Events = append(k.s.Events, k.e)
	}
	return sm
}

func (c *MxElement) getNode() string {
	s := c.Properties
	_, edgeLabel := s["edgeLabel"]

	switch {
	case edgeLabel == true:
		logger.Printf("mxCell %s is detected as EdgeLabel", c.Id)
		c.Type = "edgeLabel"
	default:
		logger.Println("Cannot detect mxCell as a node: id=", c.Id)
		c.Type = ""
	}
	return c.Type
}
