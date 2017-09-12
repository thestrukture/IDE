package core

import (
	"encoding/xml"
)

/*
	type server struct {
    XMLName    xml.Name `xml:"server"`
    Port string `xml:"port"`
    Package  string `xml:"package"`
}
	GoS Xml Strucs
	Links are used to import methods
*/
/*
	Struct for Go method handling
*/
type gosArch struct {
	interfacedMethods []string
	methodlimits      []string
	keeplocal         []string
	webmethods        []string
	structs           []string
	objects           []string
}

type gos struct {
	XMLName          xml.Name          `xml:"gos"`
	Port             string            `xml:"port"`
	Domain           string          `xml:"domain"`
	Debug            string            `xml:"trace"`
	Output           string            `xml:"output"`
	ErrorPage        string            `xml:"error"`
	NPage            string            `xml:"not_found"`
	Type             string            `xml:"deploy"`
	Main             string            `xml:"main"`
	Variables        []GlobalVariables `xml:"var"`
	WriteOut         bool
	Export           string    `xml:"export"`
	Key              string    `xml:"key"`
	Session          string    `xml:"session"`
	Template_path    string    `xml:"templatePath"`
	Web_root         string    `xml:"webroot"`
	Package          string    `xml:"package"`
	Web              string    `xml:"web"`
	Tmpl             string    `xml:"tmpl"`
	RootImports      []Import  `xml:"import"`
	Init_Func        string    `xml:"init"`
	Header           Header    `xml:"header"`
	Methods          Methods   `xml:"methods"`
	Timers           Timers    `xml:"timers"`
	Templates        Templates `xml:"templates"`
	Endpoints        Endpoints `xml:"endpoints"`
	FolderRoot, Name string
	Prod             bool
}


type Pgos struct {
	XMLName          xml.Name          `xml:"gos"`
	Port             string            `xml:"port"`
	Domain           string          `xml:"domain"`
	Debug            string            `xml:"trace"`
	Output           string            `xml:"output"`
	ErrorPage        string            `xml:"error"`
	NPage            string            `xml:"not_found"`
	Type             string            `xml:"deploy"`
	Main             string            `xml:"main"`
	Variables        []GlobalVariables `xml:"var"`
	WriteOut         bool
	Export           string    `xml:"export"`
	Key              string    `xml:"key"`
	Session          string    `xml:"session"`
	Template_path    string    `xml:"templatePath"`
	Web_root         string    `xml:"webroot"`
	Package          string    `xml:"package"`
	Web              string    `xml:"web"`
	Tmpl             string    `xml:"tmpl"`
	RootImports      []Import  `xml:"import"`
	Init_Func        string    `xml:"init"`
	Header           Header    `xml:"header"`
	Methods          Methods   `xml:"methods"`
	Timers           Timers    `xml:"timers"`
	Templates        Templates `xml:"templates"`
	Endpoints        Endpoints `xml:"endpoints"`
	FolderRoot, Name string
	Prod             bool
}
/*
type Pgos struct {
	XMLName       xml.Name          `xml:"gos"`
	Port          string            `xml:"port"`
	Output        string            `xml:"output"`
	Type          string            `xml:"deploy"`
	Main          string            `xml:"main"`
	Variables     []GlobalVariables `xml:"var"`
	WriteOut      bool
	Export        string    `xml:"export"`
	Key           string    `xml:"key"`
	Session       string    `xml:"session"`
	Template_path string    `xml:"templatePath"`
	Web_root      string    `xml:"webroot"`
	Package       string    `xml:"package"`
	Web           string    `xml:"web"`
	Tmpl          string    `xml:"tmpl"`
	RootImports   []Import  `xml:"import"`
	Init_Func     string    `xml:"init"`
	Header        Header    `xml:"header"`
	Methods       Methods   `xml:"methods"`
	Timers        Timers    `xml:"timers"`
	Templates     Templates `xml:"templates"`
	Endpoints     Endpoints `xml:"endpoints"`
	FolderRoot    string
} */

type GlobalVariables struct {
	XMLName xml.Name `xml:"var"`
	Name    string   `xml:",innerxml"`
	Type    string   `xml:"type,attr"`
}

type Error struct {
	reason string
	code   int
}

/*
	Root Types to GoS xml File
*/
type gosConfig struct {
	template_path string
	web_root      string
}
type Import struct {
	XMLName  xml.Name `xml:"import"`
	Src      string   `xml:"src,attr"`
	Download string   `xml:"fetch,attr"`
}

type Header struct {
	XMLName xml.Name `xml:"header"`
	Structs []Struct `xml:"struct"`
	Objects []Object `xml:"object"`
}

type Methods struct {
	XMLName xml.Name `xml:"methods"`
	Methods []Method `xml:"method"`
}

type Timers struct {
	XMLName xml.Name `xml:"timers"`
	Timers  []Timer  `xml:"timer"`
}

type Templates struct {
	XMLName   xml.Name   `xml:"templates"`
	Templates []Template `xml:"template"`
}

type Endpoints struct {
	XMLName   xml.Name   `xml:"endpoints"`
	Endpoints []Endpoint `xml:"end"`
}

/*
	Nested values within GoS root file
*/

type VGos struct {
	XMLName xml.Name `xml:"gos"`
	Objects []Object `xml:"object"`
	Structs []Struct `xml:"struct"`
	Methods []Method `xml:"method"`
}

type Struct struct {
	XMLName    xml.Name `xml:"struct"`
	Name       string   `xml:"name,attr"`
	Attributes string   `xml:",innerxml"`
}

type Object struct {
	XMLName xml.Name `xml:"object"`
	Name    string   `xml:"name,attr"`
	Templ   string   `xml:"struct,attr"`
	Methods string   `xml:",innerxml"`
}

type Method struct {
	XMLName    xml.Name `xml:"method"`
	Method     string   `xml:",innerxml"`
	Name       string   `xml:"name,attr"`
	Variables  string   `xml:"var,attr"`
	Limit      string   `xml:"limit,attr"`
	Object     string   `xml:"object,attr"`
	Autoface   string   `xml:"autoface,attr"`
	Keeplocal  string   `xml:"keep-local,attr"`
	Returntype string   `xml:"return,attr"`
}

type Timer struct {
	XMLName  xml.Name `xml:"timer"`
	Method   string   `xml:",innerxml"`
	Interval string   `xml:"interval,attr"`
	Name     string   `xml:"name,attr"`
	Unit     string   `xml:"unit,attr"`
}

type Template struct {
	XMLName      xml.Name `xml:"template"`
	Name         string   `xml:"name,attr"`
	TemplateFile string   `xml:"tmpl,attr"`
	Bundle       string   `xml:"bundle,attr"`
	Struct       string   `xml:"struct,attr"`
	ForcePath    bool
}

type Endpoint struct {
	XMLName xml.Name `xml:"end"`
	Path    string   `xml:"path,attr"`
	Method  string   `xml:",innerxml"`
	Type    string   `xml:"type,attr"`
	Id 		string	 `xml:"id,attr"`
}
