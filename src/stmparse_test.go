package stm2go

import (
	"fmt"
	"testing"
)

func TestSTMParse(t *testing.T) {
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

	result, _ := Parse(data)
	for k, v := range result {
		fmt.Println("key = ", k)
		fmt.Println("parent = ", v.Parent)

		for _, s := range v.States {
			fmt.Println("v.States", s)
		}

		for _, t := range v.Transitions {
			fmt.Println("v.Transitions", t)
		}

		fmt.Println("v.Initial", v.Initial)

		for _, e := range v.Events {
			fmt.Println("v.Events", e)
		}
	}

}

func TestSTMParse2(t *testing.T) {
	data := []byte(`
  <mxGraphModel dx="2132" dy="903" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
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
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=0.5;entryY=1;entryDx=0;entryDy=0;exitX=0.5;exitY=1;exitDx=0;exitDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-360" y="260" as="targetPoint" />
        <mxPoint x="-320" y="356" as="sourcePoint" />
        <Array as="points">
          <mxPoint x="200" y="260" />
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

	result, _ := Parse(data)
	for _, v := range result {
		for _, s := range v.States {
			fmt.Println("v.States", s)
		}

		for _, t := range v.Transitions {
			fmt.Println("v.Transitions", *t)

			// for _, item := range t {
			// 	fmt.Println("v.Transitions", *item)
			// }
		}

		fmt.Println("v.Initial", v.Initial)

		for _, e := range v.Events {
			fmt.Println("v.Events", e)
		}
	}
}

func TestSTMParse3(t *testing.T) {
	data := []byte(`
  <mxGraphModel dx="1746" dy="1046" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-9" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-600" y="100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="TurnOn" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOff" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-510" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.75;exitDx=0;exitDy=0;entryX=1;entryY=0.75;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-280" y="260" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="TurnOff" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-9" y="14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOn" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;container=1;" parent="1" vertex="1">
        <mxGeometry x="-250" y="-170" width="380" height="250" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="6" style="edgeStyle=none;html=1;exitX=1;exitY=0.25;exitDx=0;exitDy=0;entryX=0;entryY=0.75;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-12" source="2" target="3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="12" value="Off" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="6" vertex="1" connectable="0">
      <mxGeometry x="-0.0511" y="-1" relative="1" as="geometry">
        <mxPoint x="1" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOn6" type="state" id="2">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="110" y="85" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="7" style="edgeStyle=none;html=1;exitX=0.25;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-12" source="3" target="2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="13" value="On" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="7" vertex="1" connectable="0">
      <mxGeometry x="0.1875" relative="1" as="geometry">
        <mxPoint x="1" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOff7" type="state" id="3">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=20;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="280" y="120" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="11" style="edgeStyle=none;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-12" source="8" target="2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="8">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="30" y="105" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="This is a note." type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-400" y="-80" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>
`)

	result, _ := Parse(data)
	for _, v := range result {
		fmt.Println(v.Parent)
		for _, s := range v.States {
			fmt.Println("v.States", s)
		}

		for _, t := range v.Transitions {
			fmt.Println("v.Transitions", *t)

			// for _, item := range t {
			// 	fmt.Println("v.Transitions", *item)
			// }
		}

		fmt.Println("v.Initial", v.Initial)

		for _, e := range v.Events {
			fmt.Println("v.Events", e)
		}
	}
}

func TestSTMParse4(t *testing.T) {
	data := []byte(`
  <mxGraphModel dx="1514" dy="881" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-9" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-600" y="100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="16" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="TurnOn" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOff" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-500" y="40" width="270" height="100" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="19" style="edgeStyle=none;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-10" source="15" target="16" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="20" value="Push" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="19" vertex="1" connectable="0">
      <mxGeometry x="-0.3333" relative="1" as="geometry">
        <mxPoint x="7" y="15" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Release" type="state" id="15">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="EBAIrEwSQ_sO8G7dM4pI-10" vertex="1">
        <mxGeometry x="70" y="30" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <object label="Push" type="state" id="16">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=20;" parent="EBAIrEwSQ_sO8G7dM4pI-10" vertex="1">
        <mxGeometry x="180" y="30" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="18" style="edgeStyle=none;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-10" source="17" target="15" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="17">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="EBAIrEwSQ_sO8G7dM4pI-10" vertex="1">
        <mxGeometry x="10" y="40" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.25;exitY=1;exitDx=0;exitDy=0;entryX=0.75;entryY=1;entryDx=0;entryDy=0;" parent="1" source="3" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-280" y="260" as="targetPoint" />
        <Array as="points">
          <mxPoint x="88" y="200" />
          <mxPoint x="-297" y="200" />
        </Array>
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="TurnOff" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint x="-9" y="14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="LedOn" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;container=1;" parent="1" vertex="1">
        <mxGeometry x="-120" y="20" width="280" height="120" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="6" style="edgeStyle=none;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-12" source="2" target="3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="12" value="Push" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="6" vertex="1" connectable="0">
      <mxGeometry x="-0.0511" y="-1" relative="1" as="geometry">
        <mxPoint x="1" y="14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Release" type="state" id="2">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="80" y="50" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <object label="Push" type="state" id="3">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=20;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="190" y="50" width="70" height="50" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="11" style="edgeStyle=none;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="EBAIrEwSQ_sO8G7dM4pI-12" source="8" target="2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="8">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="EBAIrEwSQ_sO8G7dM4pI-12" vertex="1">
        <mxGeometry x="20" y="60" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="This is a note." type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-400" y="-80" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>
`)

	result, _ := Parse(data)
	for _, v := range result {
		fmt.Println("STM:", v.Parent)
		for _, s := range v.States {
			fmt.Println("v.States", s)
		}

		for _, t := range v.Transitions {
			fmt.Println("v.Transitions", *t)
		}

		fmt.Println("v.Initial", v.Initial)

		for _, e := range v.Events {
			fmt.Println("v.Events", e)
		}
	}
}
