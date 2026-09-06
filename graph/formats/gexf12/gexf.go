package gexf12

import (
	"encoding/xml"
	"time"
)

type Content struct {
	XMLName xml.Name `xml:"http://www.gexf.net/1.2draft gexf"`
	Meta    *Meta    `xml:"meta,omitempty"`
	Graph   Graph    `xml:"graph"`

	Version string `xml:"version,attr"`
	Variant string `xml:"variant,attr,omitempty"`
}

type Meta struct {
	Creator      string    `xml:"creator,omitempty"`
	Keywords     string    `xml:"keywords,omitempty"`
	Description  string    `xml:"description,omitempty"`
	LastModified time.Time `xml:"lastmodifieddate,attr,omitempty"`
}

func (t *Meta) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Meta) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

type Graph struct {
	Attributes []Attributes `xml:"attributes"`
	Nodes      Nodes        `xml:"nodes"`
	Edges      Edges        `xml:"edges"`

	TimeFormat string `xml:"timeformat,attr,omitempty"`
	Start      string `xml:"start,attr,omitempty"`
	StartOpen  string `xml:"startopen,attr,omitempty"`
	End        string `xml:"end,attr,omitempty"`
	EndOpen    string `xml:"endopen,attr,omitempty"`

	DefaultEdgeType string `xml:"defaultedgetype,attr,omitempty"`

	IDType string `xml:"idtype,attr,omitempty"`

	Mode string `xml:"mode,attr,omitempty"`
}

type Attributes struct {
	Attributes []Attribute `xml:"attribute,omitempty"`

	Class string `xml:"class,attr"`

	Mode      string `xml:"mode,attr,omitempty"`
	Start     string `xml:"start,attr,omitempty"`
	StartOpen string `xml:"startopen,attr,omitempty"`
	End       string `xml:"end,attr,omitempty"`
	EndOpen   string `xml:"endopen,attr,omitempty"`
}

type Attribute struct {
	ID    string `xml:"id,attr"`
	Title string `xml:"title,attr"`

	Type    string `xml:"type,attr"`
	Default string `xml:"default,omitempty"`
	Options string `xml:"options,omitempty"`
}

type Nodes struct {
	Count int    `xml:"count,attr,omitempty"`
	Nodes []Node `xml:"node,omitempty"`
}

type Node struct {
	ID        string     `xml:"id,attr,omitempty"`
	Label     string     `xml:"label,attr,omitempty"`
	AttValues *AttValues `xml:"attvalues"`
	Spells    *Spells    `xml:"spells"`
	Nodes     *Nodes     `xml:"nodes"`
	Edges     *Edges     `xml:"edges"`
	ParentID  string     `xml:"pid,attr,omitempty"`
	Parents   *Parents   `xml:"parents"`
	Color     *Color     `xml:"http://www.gexf.net/1.2draft/viz color"`
	Position  *Position  `xml:"http://www.gexf.net/1.2draft/viz position"`
	Size      *Size      `xml:"http://www.gexf.net/1.2draft/viz size"`
	Shape     *NodeShape `xml:"http://www.gexf.net/1.2draft/viz shape"`
	Start     string     `xml:"start,attr,omitempty"`
	StartOpen string     `xml:"startopen,attr,omitempty"`
	End       string     `xml:"end,attr,omitempty"`
	EndOpen   string     `xml:"endopen,attr,omitempty"`
}

type NodeShape struct {
	Spells *Spells `xml:"spells,omitempty"`

	Shape     string `xml:"value,attr"`
	URI       string `xml:"uri,attr,omitempty"`
	Start     string `xml:"start,attr,omitempty"`
	StartOpen string `xml:"startopen,attr,omitempty"`
	End       string `xml:"end,attr,omitempty"`
	EndOpen   string `xml:"endopen,attr,omitempty"`
}

type Color struct {
	Spells    *Spells `xml:"spells,omitempty"`
	R         byte    `xml:"r,attr"`
	G         byte    `xml:"g,attr"`
	B         byte    `xml:"b,attr"`
	A         float64 `xml:"a,attr,omitempty"`
	Start     string  `xml:"start,attr,omitempty"`
	StartOpen string  `xml:"startopen,attr,omitempty"`
	End       string  `xml:"end,attr,omitempty"`
	EndOpen   string  `xml:"endopen,attr,omitempty"`
}

type Edges struct {
	Count int    `xml:"count,attr,omitempty"`
	Edges []Edge `xml:"edge,omitempty"`
}

type Edge struct {
	ID        string     `xml:"id,attr,omitempty"`
	AttValues *AttValues `xml:"attvalues"`
	Spells    *Spells    `xml:"spells"`
	Color     *Color     `xml:"http://www.gexf.net/1.2draft/viz color"`
	Thickness *Thickness `xml:"http://www.gexf.net/1.2draft/viz thickness"`
	Shape     *Edgeshape `xml:"http://www.gexf.net/1.2draft/viz shape"`
	Start     string     `xml:"start,attr,omitempty"`
	StartOpen string     `xml:"startopen,attr,omitempty"`
	End       string     `xml:"end,attr,omitempty"`
	EndOpen   string     `xml:"endopen,attr,omitempty"`

	Type   string  `xml:"type,attr,omitempty"`
	Label  string  `xml:"label,attr,omitempty"`
	Source string  `xml:"source,attr"`
	Target string  `xml:"target,attr"`
	Weight float64 `xml:"weight,attr,omitempty"`
}

type AttValues struct {
	AttValues []AttValue `xml:"attvalue,omitempty"`
}

type AttValue struct {
	For       string `xml:"for,attr"`
	Value     string `xml:"value,attr"`
	Start     string `xml:"start,attr,omitempty"`
	StartOpen string `xml:"startopen,attr,omitempty"`
	End       string `xml:"end,attr,omitempty"`
	EndOpen   string `xml:"endopen,attr,omitempty"`
}

type Edgeshape struct {
	Shape     string  `xml:"value,attr"`
	Spells    *Spells `xml:"spells,omitempty"`
	Start     string  `xml:"start,attr,omitempty"`
	StartOpen string  `xml:"startopen,attr,omitempty"`
	End       string  `xml:"end,attr,omitempty"`
	EndOpen   string  `xml:"endopen,attr,omitempty"`
}

type Parents struct {
	Parents []Parent `xml:"parent,omitempty"`
}

type Parent struct {
	For string `xml:"for,attr"`
}

type Position struct {
	X         float64 `xml:"x,attr"`
	Y         float64 `xml:"y,attr"`
	Z         float64 `xml:"z,attr"`
	Spells    *Spells `xml:"spells,omitempty"`
	Start     string  `xml:"start,attr,omitempty"`
	StartOpen string  `xml:"startopen,attr,omitempty"`
	End       string  `xml:"end,attr,omitempty"`
	EndOpen   string  `xml:"endopen,attr,omitempty"`
}

type Size struct {
	Value     float64 `xml:"value,attr"`
	Spells    *Spells `xml:"http://www.gexf.net/1.2draft/viz spells,omitempty"`
	Start     string  `xml:"start,attr,omitempty"`
	StartOpen string  `xml:"startopen,attr,omitempty"`
	End       string  `xml:"end,attr,omitempty"`
	EndOpen   string  `xml:"endopen,attr,omitempty"`
}

type Thickness struct {
	Value     float64 `xml:"value,attr"`
	Spells    *Spells `xml:"http://www.gexf.net/1.2draft/viz spells,omitempty"`
	Start     string  `xml:"start,attr,omitempty"`
	StartOpen string  `xml:"startopen,attr,omitempty"`
	End       string  `xml:"end,attr,omitempty"`
	EndOpen   string  `xml:"endopen,attr,omitempty"`
}

type Spells struct {
	Spells []Spell `xml:"spell"`
}

type Spell struct {
	Start     string `xml:"start,attr,omitempty"`
	StartOpen string `xml:"startopen,attr,omitempty"`
	End       string `xml:"end,attr,omitempty"`
	EndOpen   string `xml:"endopen,attr,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t xsdDate) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	_ = "STUB: not implemented"
	return *new(xml.Attr), nil
}

func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
