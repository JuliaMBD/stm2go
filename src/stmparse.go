package stm2go

import (
	"log"
	"os"
)

var logger *log.Logger

type StateMachine struct {
	Parent        string
	States        []*State
	Events        []*Event
	Transitions   []*Transition
	ExTransitions []*Transition
	Initial       *State
}

func NewStateMachine(parent string) *StateMachine {
	return &StateMachine{
		Parent:        parent,
		States:        make([]*State, 0),
		Events:        make([]*Event, 0),
		Transitions:   make([]*Transition, 0),
		ExTransitions: make([]*Transition, 0),
		Initial:       nil,
	}
}

type State struct {
	Name  string
	Stm   *StateMachine
	Entry string
	Do    string
	Exit  string
}

type Event struct {
	Name   string
	Cond   string
	Action string
}

type Transition struct {
	Src   *State
	Dest  *State
	Event *Event
}

func Parse(data []byte) (map[string]*StateMachine, map[string]*State) {
	logger = log.New(os.Stdout, "[Parse] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	// logger.SetOutput(io.Discard)
	states := make(map[string]*State)
	mxstates := make(map[string]*MxElement)
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
			s := &State{
				Name: x.Value,
				Stm:  sm[p],
			}
			if v, ok := x.Properties["entry"]; ok {
				s.Entry = v
			}
			if v, ok := x.Properties["do"]; ok {
				s.Do = v
			}
			if v, ok := x.Properties["exit"]; ok {
				s.Exit = v
			}
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
				e := &Event{
					Name: x.Value,
				}
				edges[x.Parent] = e
			default:
			}
		case "edge":
			transitions = append(transitions, &elems[i])
			if x.Value != "" {
				logger.Printf("Detect an event: %s\n", x.Value)
				e := &Event{
					Name: x.Value,
				}
				edges[x.Id] = e
			}
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
			if v, ok := x.Properties["guard"]; ok {
				t.Event.Cond = v
			}
			if v, ok := x.Properties["action"]; ok {
				t.Event.Action = v
			}
			s.Transitions = append(s.Transitions, t)
		} else if ok1 && ok2 && src.Parent != dest.Parent {
			logger.Printf("Find an external transition accossing STMs between %s and %s\n", src.Parent, dest.Parent)
			s1 := sm[src.Parent]
			s2 := sm[dest.Parent]
			t := &Transition{
				Src:   states[src.Id],
				Dest:  states[dest.Id],
				Event: edges[x.Id],
			}
			smevents[key{s1, edges[x.Id]}] = struct{}{}
			s1.ExTransitions = append(s1.ExTransitions, t)
			s2.ExTransitions = append(s2.ExTransitions, t)
		} else if ok3 && ok2 && isrc.Parent == dest.Parent && isrc.Parent == x.Parent {
			// set initital
			s := sm[isrc.Parent]
			s.Initial = states[dest.Id]
		} else {
			logger.Println("Ignoring a transition ", x)
		}
	}
	for k, _ := range smevents {
		k.s.Events = append(k.s.Events, k.e)
	}
	return sm, states
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
