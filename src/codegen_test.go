package stm2go

import (
	"fmt"
	"os"
	"testing"
)

func TestMakeTransitionMap(t *testing.T) {
	ss := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	ts := []*Transition{
		&Transition{
			Src:  ss[0],
			Dest: ss[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	m := makeTransitionMap(ts)
	fmt.Println(ss)
	fmt.Println(ts)
	fmt.Println(m)
}

func TestGenHeader(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	stm := NewGoSTMSource("stm1test", nil, nil, nil, nil, pkg, false)
	stm.BaseHeader(w)
}

func TestGenStateDefinition(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, nil, nil, states[0], pkg, false)
	stm.BaseStateDefinition(w, names)
}

func TestGenInitState(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, nil, nil, states[0], pkg, false)
	stm.BaseStateInitialize(w, names)
}

func TestGenTrans(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	trans := []*Transition{
		&Transition{
			Src:  states[0],
			Dest: states[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, trans, nil, states[0], pkg, false)
	stm.BaseTransDefinition(w, names)
}

func TestGenSTM(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	trans := []*Transition{
		&Transition{
			Src:  states[0],
			Dest: states[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, trans, nil, states[0], pkg, false)
	stm.BaseHeader(w)
	stm.BaseStateDefinition(w, names)
	stm.BaseStateInitialize(w, names)
	stm.BaseTransDefinition(w, names)
}

func TestGenCommon(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	pkg.Common(w)
}

func TestGenFunc(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	trans := []*Transition{
		&Transition{
			Src:  states[0],
			Dest: states[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, trans, nil, states[0], pkg, false)
	stm.ImplFunctions(w, make(map[*State][]*GoSTMSource), names)
}

func TestGenTest(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	trans := []*Transition{
		&Transition{
			Src:  states[0],
			Dest: states[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, trans, nil, states[0], pkg, false)
	pkg.TestGen(w, []*GoSTMSource{stm}, names)
}

func TestGenMain(t *testing.T) {
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/example", "test")
	states := []*State{
		&State{
			Name: "A",
		},
		&State{
			Name: "B",
		},
	}
	trans := []*Transition{
		&Transition{
			Src:  states[0],
			Dest: states[1],
			Event: &Event{
				Name: "EventA",
			},
		},
	}
	names := map[string]string{"stm1test": "stm0"}
	stm := NewGoSTMSource("stm1test", states, trans, nil, states[0], pkg, false)
	pkg.GenMain(w, []*GoSTMSource{stm}, names)
}

func TestGenAllGo(t *testing.T) {
	data := []byte(`
  <mxGraphModel dx="2037" dy="867" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="3" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-350" y="115" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="time3sec" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="On" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-590" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=0.5;entryY=1;entryDx=0;entryDy=0;exitX=0.5;exitY=1;exitDx=0;exitDy=0;" parent="1" source="3" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-360" y="260" as="targetPoint" />
        <mxPoint x="-320" y="356" as="sourcePoint" />
        <Array as="points">
          <mxPoint x="-5" y="260" />
          <mxPoint x="-525" y="260" />
        </Array>
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="correctKey" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-16" y="29" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-490" y="-120" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
    <object label="Off" type="state" id="3">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="1" vertex="1">
        <mxGeometry x="-320" y="20" width="630" height="190" as="geometry">
          <mxRectangle x="-320" y="20" width="50" height="23" as="alternateBounds" />
        </mxGeometry>
      </mxCell>
    </object>
    <object label="OffEntered" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
        <mxGeometry x="440" y="60" width="160" height="85" as="geometry" />
      </mxCell>
    </object>
    <object label="OffEmpty" type="state" id="m3naWr25p_J3Rivu2WGc-2">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
        <mxGeometry x="150" y="60" width="160" height="85" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-1" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.25;exitDx=0;exitDy=0;entryX=0;entryY=0.25;entryDx=0;entryDy=0;" parent="3" source="m3naWr25p_J3Rivu2WGc-2" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="580" y="81.5" as="targetPoint" />
        <mxPoint x="540" y="280" as="sourcePoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-2" value="keyFailed" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-1" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="8" y="-13" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-3" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=1;entryY=0.75;entryDx=0;entryDy=0;exitX=0;exitY=0.75;exitDx=0;exitDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-12" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="450" y="280" as="targetPoint" />
        <mxPoint x="580" y="123.5" as="sourcePoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-4" value="keyEntered" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-3" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-8" y="20" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="3" vertex="1">
        <mxGeometry x="30" y="87.5" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-9" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
      <mxGeometry relative="1" as="geometry">
        <Array as="points">
          <mxPoint x="100" y="102.5" />
          <mxPoint x="100" y="102.5" />
        </Array>
        <mxPoint x="489.9999999999998" y="-20" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="6" style="edgeStyle=none;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="5" target="3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="5">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-20" y="-100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>

`)

	result, s := Parse(data)
	fmt.Println(s)
	for k, _ := range result {
		fmt.Println(k)
	}
	pkg := NewGoPkgSource("github.com/JuliaMBD", "test")
	stmap, sttree, root := NewGoSTMMap(pkg, result, s)
	fmt.Println(pkg)
	fmt.Println(stmap)
	fmt.Println(sttree)
	fmt.Printf("root %p\n", root)
}

func TestGenGoSource1(t *testing.T) {
	data := []byte(`
  <mxGraphModel dx="2037" dy="867" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="3" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-350" y="115" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="time3sec" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="On" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-590" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=0.5;entryY=1;entryDx=0;entryDy=0;exitX=0.5;exitY=1;exitDx=0;exitDy=0;" parent="1" source="3" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-360" y="260" as="targetPoint" />
        <mxPoint x="-320" y="356" as="sourcePoint" />
        <Array as="points">
          <mxPoint x="-5" y="260" />
          <mxPoint x="-525" y="260" />
        </Array>
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="correctKey" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-16" y="29" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-490" y="-120" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
    <object label="Off" type="state" id="3">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="1" vertex="1">
        <mxGeometry x="-320" y="20" width="630" height="190" as="geometry">
          <mxRectangle x="-320" y="20" width="50" height="23" as="alternateBounds" />
        </mxGeometry>
      </mxCell>
    </object>
    <object label="OffEntered" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
        <mxGeometry x="440" y="60" width="160" height="85" as="geometry" />
      </mxCell>
    </object>
    <object label="OffEmpty" type="state" id="m3naWr25p_J3Rivu2WGc-2">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
        <mxGeometry x="150" y="60" width="160" height="85" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-1" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.25;exitDx=0;exitDy=0;entryX=0;entryY=0.25;entryDx=0;entryDy=0;" parent="3" source="m3naWr25p_J3Rivu2WGc-2" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="580" y="81.5" as="targetPoint" />
        <mxPoint x="540" y="280" as="sourcePoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-2" value="keyFailed" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-1" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="8" y="-13" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-3" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=1;entryY=0.75;entryDx=0;entryDy=0;exitX=0;exitY=0.75;exitDx=0;exitDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-12" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="450" y="280" as="targetPoint" />
        <mxPoint x="580" y="123.5" as="sourcePoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="HdLz7Pnjk3gmC13WeZwF-4" value="keyEntered" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-3" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-8" y="20" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="3" vertex="1">
        <mxGeometry x="30" y="87.5" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-9" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
      <mxGeometry relative="1" as="geometry">
        <Array as="points">
          <mxPoint x="100" y="102.5" />
          <mxPoint x="100" y="102.5" />
        </Array>
        <mxPoint x="489.9999999999998" y="-20" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="6" style="edgeStyle=none;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="5" target="3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="5">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-20" y="-100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>

`)

	stms, states := Parse(data)
	w := NewWriter(os.Stdout)
	pkg := NewGoPkgSource("github.com/JuliaMBD", "test")
	stmap, sttree, root := NewGoSTMMap(pkg, stms, states)

	pkg.Common(w)

	s := stmap[0]
	names := map[string]string{"1": "stm0", "3": "stm1"}

	s.BaseHeader(w)
	s.BaseStateDefinition(w, names)
	s.BaseStateInitialize(w, names)
	s.BaseTransDefinition(w, names)

	s.ImplHeader(w)
	s.ImplFunctions(w, sttree, names)

	pkg.TestGen(w, sttree[root], names)
	pkg.GenMain(w, sttree[root], names)
}
