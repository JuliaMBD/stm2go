package sm2go

import (
	"fmt"
	"io/ioutil"
	"os"

	// "os"
	"testing"
)

/*
func TestWriteStopWatch(t *testing.T) {
	// jsonファイル読み込みのテスト

	// state.json読み込み
	states, err := ioutil.ReadFile("./stop_watch/stop_watch.json")
	if err != nil {
		panic(err.Error())
	}

	var sm StateMent
	err = json.Unmarshal(states, &sm)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("------------------------------　output.go　-----------------------------------")
	write_package()
	write_enum(sm.States)
	write_event(sm.Transitions)
	write_init(sm.Initial)
	fmt.Println("------------------------------　output_edit.go　-----------------------------------")
	write_package_edit()
	write_func(sm.States, sm.Events)
	write_main()
}
*/

// 図の情報（xmlファイル）の解析情報を表示するテスト
func TestSM2GoEasy(t *testing.T) {
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

	result := Parse(data)
	for k, v := range result {
		fmt.Println("key = ", k)

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

// func TestSM2GoCom(t *testing.T) {
// 	data := []byte(`
//   <mxGraphModel dx="2132" dy="903" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
//   <root>
//     <mxCell id="0" />
//     <mxCell id="1" parent="0" />
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="3" edge="1">
//       <mxGeometry relative="1" as="geometry">
//         <mxPoint x="-350" y="115" as="targetPoint" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="time3sec" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
//       <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
//         <mxPoint x="-19" y="-14" as="offset" />
//       </mxGeometry>
//     </mxCell>
//     <object label="On" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
//       <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
//         <mxGeometry x="-590" y="75" width="130" height="80" as="geometry" />
//       </mxCell>
//     </object>
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=0.5;entryY=1;entryDx=0;entryDy=0;exitX=0.5;exitY=1;exitDx=0;exitDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
//       <mxGeometry relative="1" as="geometry">
//         <mxPoint x="-360" y="260" as="targetPoint" />
//         <mxPoint x="-320" y="356" as="sourcePoint" />
//         <Array as="points">
//           <mxPoint x="200" y="260" />
//           <mxPoint x="-525" y="260" />
//         </Array>
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="correctKey" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
//       <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
//         <mxPoint x="-16" y="29" as="offset" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
//       <mxGeometry relative="1" as="geometry" />
//     </mxCell>
//     <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
//       <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
//         <mxGeometry x="-490" y="-120" width="80" height="100" as="geometry" />
//       </mxCell>
//     </object>
//     <object label="Off" type="state" id="3">
//       <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="1" vertex="1">
//         <mxGeometry x="-320" y="20" width="630" height="190" as="geometry">
//           <mxRectangle x="-320" y="20" width="50" height="23" as="alternateBounds" />
//         </mxGeometry>
//       </mxCell>
//     </object>
//     <object label="OffEntered" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
//       <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
//         <mxGeometry x="440" y="60" width="160" height="85" as="geometry" />
//       </mxCell>
//     </object>
//     <object label="OffEmpty" type="state" id="m3naWr25p_J3Rivu2WGc-2">
//       <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="3" vertex="1">
//         <mxGeometry x="150" y="60" width="160" height="85" as="geometry" />
//       </mxCell>
//     </object>
//     <mxCell id="HdLz7Pnjk3gmC13WeZwF-1" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.25;exitDx=0;exitDy=0;entryX=0;entryY=0.25;entryDx=0;entryDy=0;" parent="3" source="m3naWr25p_J3Rivu2WGc-2" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
//       <mxGeometry relative="1" as="geometry">
//         <mxPoint x="580" y="81.5" as="targetPoint" />
//         <mxPoint x="540" y="280" as="sourcePoint" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="HdLz7Pnjk3gmC13WeZwF-2" value="keyFailed" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-1" vertex="1" connectable="0">
//       <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
//         <mxPoint x="8" y="-13" as="offset" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="HdLz7Pnjk3gmC13WeZwF-3" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;entryX=1;entryY=0.75;entryDx=0;entryDy=0;exitX=0;exitY=0.75;exitDx=0;exitDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-12" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
//       <mxGeometry relative="1" as="geometry">
//         <mxPoint x="450" y="280" as="targetPoint" />
//         <mxPoint x="580" y="123.5" as="sourcePoint" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="HdLz7Pnjk3gmC13WeZwF-4" value="keyEntered" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="HdLz7Pnjk3gmC13WeZwF-3" vertex="1" connectable="0">
//       <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
//         <mxPoint x="-8" y="20" as="offset" />
//       </mxGeometry>
//     </mxCell>
//     <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
//       <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="3" vertex="1">
//         <mxGeometry x="30" y="87.5" width="30" height="30" as="geometry" />
//       </mxCell>
//     </object>
//     <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="3" source="EBAIrEwSQ_sO8G7dM4pI-9" target="m3naWr25p_J3Rivu2WGc-2" edge="1">
//       <mxGeometry relative="1" as="geometry">
//         <Array as="points">
//           <mxPoint x="100" y="102.5" />
//           <mxPoint x="100" y="102.5" />
//         </Array>
//         <mxPoint x="489.9999999999998" y="-20" as="targetPoint" />
//       </mxGeometry>
//     </mxCell>
//     <mxCell id="6" style="edgeStyle=none;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="5" target="3" edge="1">
//       <mxGeometry relative="1" as="geometry" />
//     </mxCell>
//     <object label="" type="initialstate" id="5">
//       <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
//         <mxGeometry x="-20" y="-100" width="30" height="30" as="geometry" />
//       </mxCell>
//     </object>
//   </root>
// </mxGraphModel>
// `)

// 	result := Parse(data)
// 	for _, v := range result {
// 		for _, s := range v.States {
// 			fmt.Println("v.States", s)
// 		}

// 		for _, t := range v.Transitions {
// 			fmt.Println("v.Transitions", *t)

// 			// for _, item := range t {
// 			// 	fmt.Println("v.Transitions", *item)
// 			// }
// 		}

// 		fmt.Println("v.Initial", v.Initial)

// 		for _, e := range v.Events {
// 			fmt.Println("v.Events", e)
// 		}
// 	}
// }

// ソースコードをファイルに書き込むテスト(stop_watch)
func TestSM2Go02(t *testing.T) {

	// 入力ファイル（diagrams.netにて作成）の読み込み
	// var result map[string]*StateMachine

	count := 0

	var oline []string
	var oeline []string
	var otline []string
	var omline []string
	var osline []string

	m := map[string]string{"1": "top", "2": "child", "3": "grandchild"}
	// デフォルトで先頭文字は小文字
	// 10階層以上ある場合は，「階層が深すぎること」を通知．
	// m := map[string]string{"1": "one", "2": "two", "3": "three", ...}
	// 基本的には先頭小文字に
	//

	if xml, err := ioutil.ReadFile("keypad_easy.drawio"); err == nil {
		result := Parse(xml)

		for key, v := range result {
			base, err := os.Create(fmt.Sprintf("testfile/%s_base.go", m[key]))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer base.Close()

			impl, err := os.Create(fmt.Sprintf("testfile/%s_impl.go", m[key]))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer impl.Close()

			test, err := os.Create(fmt.Sprintf("testfile/%s_test.go", m[key]))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer test.Close()

			oline, oeline, otline, omline, osline = WriteAll(xml, "modelTest", key, v, m, count)

			// package_base.go
			for _, o := range oline {
				b := []byte(o)
				_, err := base.Write(b)
				if err != nil {
					panic(err)
				}
			}

			// package_impl.go
			for _, oe := range oeline {
				b := []byte(oe)
				_, err := impl.Write(b)
				if err != nil {
					panic(err)
				}
			}

			// package_test.go
			for _, ot := range otline {
				b := []byte(ot)
				_, err := test.Write(b)
				if err != nil {
					panic(err)
				}
			}

			count++
		}

	} else {
		panic(err)
	}

	// sm2goディレクトリ生成
	if err := os.Mkdir("testfile/sm2go", 0777); err != nil {
		fmt.Println(err)
	}

	// sm2go.go生成
	sm, err := os.Create("testfile/sm2go/sm2go.go")
	if err != nil {
		panic(err)
	}
	defer sm.Close()

	// mainディレクトリ生成
	if err := os.Mkdir("testfile/main", 0777); err != nil {
		fmt.Println(err)
	}

	// main.go生成
	main, err := os.Create("testfile/main/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer main.Close()

	// sm2go.go
	for _, os := range osline {
		b := []byte(os)
		_, err := sm.Write(b)
		if err != nil {
			panic(err)
		}
	}

	// main.go
	for _, om := range omline {
		b := []byte(om)
		_, err := main.Write(b)
		if err != nil {
			panic(err)
		}
	}
}

// ソースコードをファイルに書き込むテスト(ON_OFF)
/*
func TestSM2Go03(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1670" dy="994" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-9" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-630" y="100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="push_button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="ON" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-510" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <object label="OFF" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;startSize=23;" parent="1" vertex="1">
        <mxGeometry x="-290" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-410" y="-120" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="i0cF6CM7KJtXVy2GyI8c-1" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.75;exitDx=0;exitDy=0;entryX=1;entryY=0.75;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry">
        <mxPoint x="-370" y="125" as="sourcePoint" />
        <mxPoint x="-370" y="160" as="targetPoint" />
      </mxGeometry>
    </mxCell>
    <mxCell id="i0cF6CM7KJtXVy2GyI8c-2" value="pull_button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="i0cF6CM7KJtXVy2GyI8c-1" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="9" y="19" as="offset" />
      </mxGeometry>
    </mxCell>
  </root>
</mxGraphModel>

`)

	f, err := os.Create("stop_watch/output/output_test/output.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fe, err := os.Create("stop_watch/output/output_test/output_edit.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fe.Close()

	result := Parse(data)
	for _, v := range result {
		// ------------------------------　output.go　-----------------------------------
		write_package(f)
		write_enum(f, v.States)
		write_event(f, v.Transitions)
		write_init(f, v.Initial)
		// ------------------------------　output_edit.go　-----------------------------------
		write_package_edit(fe)
		write_func(fe, v.States, v.Events)
		write_main(fe)
	}
}
*/
