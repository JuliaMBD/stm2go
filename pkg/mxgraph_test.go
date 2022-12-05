package stm2go

import (
	"fmt"
	"testing"
)

// type TestParser struct{}

// func (p *TestParser) GetNode(c *mxCell, s map[string]string) (Component, bool) {
// 	result := Component{
// 		"id":       c.Id,
// 		"parent":   c.Parent,
// 		"value":    c.Value,
// 		"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
// 	}
// 	for k, v := range s {
// 		result[k] = v
// 	}
// 	return result, true
// }

// func (p *TestParser) GetArc(c *mxCell, s map[string]string) (Component, bool) {
// 	result := Component{
// 		"id":     c.Id,
// 		"parent": c.Parent,
// 		"value":  c.Value,
// 		"source": c.Source,
// 		"target": c.Target,
// 	}
// 	for k, v := range s {
// 		result[k] = v
// 	}
// 	return result, true
// }

func TestDecode01(t *testing.T) {
	s := []byte(`5VhRb5swEP41eewUoJDmsUnTblInRcqkrY9uuIJbwyHjJNBfv3MwEOK1zbYmbZSn5D7ujP1992FMzxsnxY1kWfwdQxA9tx8WPe+q57pOfzigH42UFRI4QQVEkocmqQVm/BnqSoMueAh5J1EhCsWzLjjHNIW56mBMSlx10x5QdO+asQgsYDZnwkZ/8lDFFXrhDlr8K/Aoru/sBMPqSsLqZLOSPGYhrjYgb9LzxhJRVf+SYgxCk1fzUtVdv3C1mZiEVO1ScFte3/ojPHviNwWbPH5Llrx/5pm5qbJeMIS0fhOiVDFGmDIxadGRxEUagh61T1Gbc4uYEegQ+AhKlUZMtlBIUKwSYa5CwdUvXf7FN9GdGUz/vyo2g7IOUiXLjSId3tXj6aAtW0d1XbU+vagXaTNQjgs5h1e4qtuPyQjUK3luIy65AjABmg/VSRBM8WV3Hsy0Z9TktQrSHyPiXwhqxl0ysTB3Sm2JhSD7aClXMVcwy9h63StycFcolmeVpx54oQUfkcsU4ynIDrtLkAqK1/m1+agLzo0/zAPCGZh41drNM1C84bQae3cGg2O3RP9wlnB3tIT/n5ZYl15KycqNhAx5qvKNkacaaDvLc/1uZ3n+VnNUI7at0kzt37vHtfxntVO3Wd6woERFFGFK4fCdHNfsrQ0vtuP8PzjO2Zfj6oE/g+Wcz245f0fLDT9yF/LfdsGn2oS8bUt89CbkODZhx2WJA76YDXe0xAs9cBhLDI9vY2iOLh+2Mdi76Q861HGKLfJo2WqLIcj5M7tfJ2iGdL/nVeuvnyuCR5q9ObGjnyUjzR2nk9+luZDwMFz7yrxo0G38Uc+/IuSBCzFGgfoRlGKqk3Il8Qm2wK6ie5DoYjeF3L0p5NkK5SCXnLr2RBXqd9863cCWKDioROeWRNP5Ilc0vDxRjYKtg8HAtyQaHFQi+31pql0EklZ9mhJ51meBvWlEYfsVrjoHtt8yvclv`)
	expected := `<mxGraphModel dx="1097" dy="616" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0"><root><mxCell id="0"/><mxCell id="1" parent="0"/><mxCell id="LyFL5Bo-kiGxaEjImvi0-3" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-1" target="LyFL5Bo-kiGxaEjImvi0-2"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-1" value="n" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1"><mxGeometry x="140" y="170" width="30" height="30" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-6" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-2" target="LyFL5Bo-kiGxaEjImvi0-5"><mxGeometry relative="1" as="geometry"><Array as="points"><mxPoint x="325" y="135"/></Array></mxGeometry></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-2" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1"><mxGeometry x="210" y="130" width="50" height="10" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-10" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-5" target="LyFL5Bo-kiGxaEjImvi0-9"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-5" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1"><mxGeometry x="310" y="170" width="30" height="30" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-9" target="LyFL5Bo-kiGxaEjImvi0-1"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-9" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1"><mxGeometry x="210" y="220" width="50" height="10" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-12" value="Tarrival" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="210" y="80" width="50" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-13" value="Tservice" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="205" y="260" width="60" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-14" value="Pcustomer" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="65" y="175" width="70" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-15" value="Pservered" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="340" y="175" width="70" height="20" as="geometry"/></mxCell></root></mxGraphModel>`
	w, err := decode(s)
	if err != nil {
		fmt.Println("error")
	}
	if expected != string(w) {
		t.Error("Error in decode test.")
	}
}

func TestMxgraph01(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1257" dy="789" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=1;" parent="1" vertex="1">
      <mxGeometry x="10" y="30" width="110" height="110" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;fillColor=#000000;" vertex="1" parent="tUs784MvHepMPJK3DbLS-1">
      <mxGeometry x="30" y="30" width="20" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-19" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-18">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-2" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
      <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-3" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
      <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-12" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-6" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" vertex="1" connectable="0" parent="1">
      <mxGeometry x="145" y="150" width="40" height="80" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="tttttt" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-18" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;dashed=1;" vertex="1" parent="1">
      <mxGeometry x="145" y="30" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
      <mxGeometry x="310" y="320" width="40" height="20" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	result, err := getCells(data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestMxgraph02(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1286" dy="809" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=1;" parent="1" vertex="1">
      <mxGeometry x="10" y="30" width="110" height="110" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-1" vertex="1">
      <mxGeometry x="30" y="30" width="20" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-2" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
      <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-3" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
      <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-12" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-6" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" parent="1" vertex="1" connectable="0">
      <mxGeometry x="145" y="150" width="40" height="80" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="tttttt" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="60" width="40" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-18" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;dashed=1;" parent="1" vertex="1">
      <mxGeometry x="145" y="30" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-19" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=oval;endFill=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-18" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}

func TestMxgraph03(t *testing.T) {
	data := []byte(`
<mxfile host="app.diagrams.net" modified="2021-08-10T02:22:24.531Z" agent="5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36" etag="EpJn6j2i2rT-mE96lq05" version="14.9.4" type="device"><diagram id="xAHSQMcz1_Kjpg2oic0b" name="Page-1">5VhRb5swEP41eewUoJDmsUnTblInRcqkrY9uuIJbwyHjJNBfv3MwEOK1zbYmbZSn5D7ujP1992FMzxsnxY1kWfwdQxA9tx8WPe+q57pOfzigH42UFRI4QQVEkocmqQVm/BnqSoMueAh5J1EhCsWzLjjHNIW56mBMSlx10x5QdO+asQgsYDZnwkZ/8lDFFXrhDlr8K/Aoru/sBMPqSsLqZLOSPGYhrjYgb9LzxhJRVf+SYgxCk1fzUtVdv3C1mZiEVO1ScFte3/ojPHviNwWbPH5Llrx/5pm5qbJeMIS0fhOiVDFGmDIxadGRxEUagh61T1Gbc4uYEegQ+AhKlUZMtlBIUKwSYa5CwdUvXf7FN9GdGUz/vyo2g7IOUiXLjSId3tXj6aAtW0d1XbU+vagXaTNQjgs5h1e4qtuPyQjUK3luIy65AjABmg/VSRBM8WV3Hsy0Z9TktQrSHyPiXwhqxl0ysTB3Sm2JhSD7aClXMVcwy9h63StycFcolmeVpx54oQUfkcsU4ynIDrtLkAqK1/m1+agLzo0/zAPCGZh41drNM1C84bQae3cGg2O3RP9wlnB3tIT/n5ZYl15KycqNhAx5qvKNkacaaDvLc/1uZ3n+VnNUI7at0kzt37vHtfxntVO3Wd6woERFFGFK4fCdHNfsrQ0vtuP8PzjO2Zfj6oE/g+Wcz245f0fLDT9yF/LfdsGn2oS8bUt89CbkODZhx2WJA76YDXe0xAs9cBhLDI9vY2iOLh+2Mdi76Q861HGKLfJo2WqLIcj5M7tfJ2iGdL/nVeuvnyuCR5q9ObGjnyUjzR2nk9+luZDwMFz7yrxo0G38Uc+/IuSBCzFGgfoRlGKqk3Il8Qm2wK6ie5DoYjeF3L0p5NkK5SCXnLr2RBXqd9863cCWKDioROeWRNP5Ilc0vDxRjYKtg8HAtyQaHFQi+31pql0EklZ9mhJ51meBvWlEYfsVrjoHtt8yvclv</diagram></mxfile>
`)
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}

func TestMxgraph04(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1097" dy="616" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
      <root>
        <mxCell id="0" />
        <mxCell id="1" parent="0" />
        <mxCell id="2" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" source="3" target="5" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="3" value="n" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1">
          <mxGeometry x="140" y="170" width="30" height="30" as="geometry" />
        </mxCell>
        <mxCell id="4" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" source="5" target="7" parent="1">
          <mxGeometry relative="1" as="geometry">
            <Array as="points">
              <mxPoint x="325" y="135" />
            </Array>
          </mxGeometry>
        </mxCell>
        <mxCell id="5" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1">
          <mxGeometry x="210" y="130" width="50" height="10" as="geometry" />
        </mxCell>
        <mxCell id="6" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" source="7" target="9" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1">
          <mxGeometry x="310" y="170" width="30" height="30" as="geometry" />
        </mxCell>
        <mxCell id="8" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" source="9" target="3" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="9" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1">
          <mxGeometry x="210" y="220" width="50" height="10" as="geometry" />
        </mxCell>
        <mxCell id="10" value="Tarrival" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="210" y="80" width="50" height="20" as="geometry" />
        </mxCell>
        <mxCell id="11" value="Tservice" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="205" y="260" width="60" height="20" as="geometry" />
        </mxCell>
        <mxCell id="12" value="Pcustomer" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="65" y="175" width="70" height="20" as="geometry" />
        </mxCell>
        <mxCell id="13" value="Pservered" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="340" y="175" width="70" height="20" as="geometry" />
        </mxCell>
      </root>
    </mxGraphModel>
    `)
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}

func TestMxgraph05(t *testing.T) {
	data := []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<mxfile><diagram><mxGraphModel dx="1097" dy="616" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
<root>
        <mxCell id="0" />
        <mxCell id="1" parent="0" />
        <mxCell id="2" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" source="3" target="5" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="3" value="n" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1">
          <mxGeometry x="140" y="170" width="30" height="30" as="geometry" />
        </mxCell>
        <mxCell id="4" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" source="5" target="7" parent="1">
          <mxGeometry relative="1" as="geometry">
            <Array as="points">
              <mxPoint x="325" y="135" />
            </Array>
          </mxGeometry>
        </mxCell>
        <mxCell id="5" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1">
          <mxGeometry x="210" y="130" width="50" height="10" as="geometry" />
        </mxCell>
        <mxCell id="6" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" source="7" target="9" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1">
          <mxGeometry x="310" y="170" width="30" height="30" as="geometry" />
        </mxCell>
        <mxCell id="8" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" source="9" target="3" parent="1">
          <mxGeometry relative="1" as="geometry" />
        </mxCell>
        <mxCell id="9" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1">
          <mxGeometry x="210" y="220" width="50" height="10" as="geometry" />
        </mxCell>
        <mxCell id="10" value="Tarrival" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="210" y="80" width="50" height="20" as="geometry" />
        </mxCell>
        <mxCell id="11" value="Tservice" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="205" y="260" width="60" height="20" as="geometry" />
        </mxCell>
        <mxCell id="12" value="Pcustomer" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="65" y="175" width="70" height="20" as="geometry" />
        </mxCell>
        <mxCell id="13" value="Pservered" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
          <mxGeometry x="340" y="175" width="70" height="20" as="geometry" />
        </mxCell>
      </root>
    </mxGraphModel>
  </diagram>
</mxfile>
    `)
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}

func TestMxgraph06(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1865" dy="753" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="1" type="place" id="tUs784MvHepMPJK3DbLS-1">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" parent="1" vertex="1">
        <mxGeometry x="20" y="40" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="place" id="tUs784MvHepMPJK3DbLS-2">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
        <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="exp" rate="10" id="tUs784MvHepMPJK3DbLS-3">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
        <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-12" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-6" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-3" value="#Pnormal" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="tUs784MvHepMPJK3DbLS-12" vertex="1" connectable="0">
      <mxGeometry x="-0.2596" relative="1" as="geometry">
        <mxPoint x="-1" y="13" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" parent="1" vertex="1" connectable="0">
      <mxGeometry x="145" y="150" width="95" height="85" as="geometry" />
    </mxCell>
    <object label="" type="gen" guard="#Pnormal &gt;= 2" dist="unif" id="tUs784MvHepMPJK3DbLS-6">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
        <mxGeometry y="20" width="50" height="10" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry x="5" y="50" width="40" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="imm" id="tUs784MvHepMPJK3DbLS-18">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;strokeWidth=2;" parent="1" vertex="1">
        <mxGeometry x="145" y="30" width="50" height="2" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-19" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=oval;endFill=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-18" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-1" value="Pnormal" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="-5" y="10" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-2" value="Pfailure" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="270" y="120" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-6" value="trepair" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="180" width="50" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-7" value="Tfail" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="120" y="110" width="40" height="20" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
    `)
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}
