package stm2go

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

type mxFile struct {
	XMLName     xml.Name `xml:"mxfile"`
	Diagram     []byte   `xml:"diagram"`
	Description []byte   `xml:",innerxml"`
}

type mxDiagram struct {
	XMLName     xml.Name `xml:"diagram"`
	GraphModel  []byte   `xml:"mxGraphModel"`
	Description []byte   `xml:",innerxml"`
}

type mxGraphModel struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    mxCells  `xml:"root"`
}

type mxCells struct {
	XMLName xml.Name   `xml:"root"`
	Cells   []mxCell   `xml:"mxCell"`
	Objects []mxObject `xml:"object"`
}

type mxObject struct {
	XMLName xml.Name   `xml:"object"`
	Label   string     `xml:"label,attr"`
	Id      string     `xml:"id,attr"`
	Type    string     `xml:"type,attr"`
	AnyAttr []xml.Attr `xml:",any,attr"`
	Cell    mxCell     `xml:"mxCell"`
}

type mxCell struct {
	XMLName  xml.Name   `xml:"mxCell"`
	Id       string     `xml:"id,attr"`
	Parent   string     `xml:"parent,attr"`
	Value    string     `xml:"value,attr"`
	Style    string     `xml:"style,attr"`
	Vertex   bool       `xml:"vertex,attr"`
	Edge     bool       `xml:"edge,attr"`
	Source   string     `xml:"source,attr"`
	Target   string     `xml:"target,attr"`
	Geometry mxGeometry `xml:"mxGeometry"`
}

type mxGeometry struct {
	XMLName xml.Name `xml:"mxGeometry"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
	Width   float64  `xml:"width,attr"`
	Height  float64  `xml:"height,attr"`
}

type MxElement struct {
	Id         string
	Type       string
	Parent     string
	Value      string
	Properties map[string]string
	Geometry   mxGeometry
}

func decode(data []byte) ([]byte, error) {
	var b []byte
	var s string
	var err error
	b, err = base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	r := flate.NewReader(bytes.NewReader(b))
	defer r.Close()
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s, err = url.QueryUnescape(string(b))
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func getCells(data []byte) ([]MxElement, error) {
	model := mxGraphModel{}
	elems := make([]MxElement, 0)
	if err := xml.Unmarshal(data, &model); err == nil {
		for _, x := range model.Root.Objects {
			elem := MxElement{
				Id:         "",
				Type:       "",
				Parent:     "",
				Value:      "",
				Properties: make(map[string]string),
				Geometry:   mxGeometry{},
			}
			elem.Id = x.Id
			elem.Type = x.Type
			elem.Parent = x.Cell.Parent
			elem.Value = x.Label
			elem.Geometry = x.Cell.Geometry
			for _, v := range x.AnyAttr {
				elem.Properties[v.Name.Local] = v.Value
			}
			if x.Type == "" && x.Cell.Vertex {
				elem.Type = "vertex"
			} else if x.Type == "" && x.Cell.Edge {
				elem.Type = "edge"
				elem.Properties["source"] = x.Cell.Source
				elem.Properties["target"] = x.Cell.Target
			}
			x.Cell.getStyle(elem.Properties)
			elems = append(elems, elem)
		}
		for _, x := range model.Root.Cells {
			elem := MxElement{
				Id:         "",
				Type:       "",
				Parent:     "",
				Value:      "",
				Properties: make(map[string]string),
				Geometry:   mxGeometry{},
			}
			elem.Id = x.Id
			elem.Parent = x.Parent
			elem.Value = x.Value
			elem.Geometry = x.Geometry
			if x.Vertex {
				elem.Type = "vertex"
			} else if x.Edge {
				elem.Type = "edge"
				elem.Properties["source"] = x.Source
				elem.Properties["target"] = x.Target
			}
			x.getStyle(elem.Properties)
			elems = append(elems, elem)
		}
		return elems, nil
	} else {
		return nil, err
	}
}

func (c *mxCell) getStyle(properties map[string]string) {
	for _, x := range strings.Split(c.Style, ";") {
		e := strings.Split(x, "=")
		if e[0] != "" {
			switch len(e) {
			case 1:
				properties[e[0]] = ""
			case 2:
				properties[e[0]] = e[1]
			default:
				log.Panic("Error: Style in Cell")
			}
		}
	}
}

func GetGraphModel(data []byte) ([]MxElement, error) {
	model := mxFile{}
	if err := xml.Unmarshal(data, &model); err == nil {
		diagram := mxDiagram{}
		if s, err := decode(model.Diagram); err == nil {
			return getCells(s)
		} else if err := xml.Unmarshal(model.Description, &diagram); err == nil {
			return getCells(diagram.Description)
		} else {
			return nil, err
		}
	} else {
		return getCells(data)
	}
}
