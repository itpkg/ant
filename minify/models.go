package minify

import "encoding/xml"

//HTML html tag
type HTML struct {
	XMLName xml.Name `xml:"html"`
	Head    Head     `xml:"head"`
	Body    string   `xml:"body"`
}

//Head head tag
type Head struct {
	Title string
}
