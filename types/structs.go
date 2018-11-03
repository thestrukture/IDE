package types

type myDemoObject Dex

type FSCs struct {
	Path string
	Hide bool
	Form Forms
}
type Dex struct {
	Misc string
	Text string
	Link string
}
type SoftUser struct {
	Username         string
	Email            string
	Password         []byte
	Apps             []App
	Docker           string
	TrialEnd         int64
	StripeID, FLogin string
}
type USettings struct {
	LastPaid string
	Email    string
	StripeID string
}
type App struct {
	Type             string
	Name             string
	PublicName       string
	Css              []string
	Groups           []string
	Passed, Running  bool
	LatestBuild, Pid string
}
type TemplateEdits struct {
	SavesTo, PKG, PreviewLink, ID, Mime string
	File                                []byte
	Settings                            RPut
}
type WebRootEdits struct {
	SavesTo, Type, PreviewLink, ID, PKG string
	Faas                                bool
	File                                []byte
}
type TEditor struct {
	PKG, Type, LType string
	CreateForm       RPut
}
type Navbars struct {
	Mode string
	ID   string
}
type SModal struct {
	Title   string
	Body    string
	Color   string
	Buttons []SButton
	Form    Forms
}
type Forms struct {
	Link    string
	Inputs  []Inputs
	Buttons []SButton
	CTA     string
	Class   string
}
type SButton struct {
	Text  string
	Class string
	Link  string
}
type STab struct {
	Buttons []SButton
}
type DForm struct {
	Text, Link string
}
type Alertbs struct {
	Type     string
	Text     string
	Redirect string
}
type Inputs struct {
	Misc    string
	Text    string
	Name    string
	Type    string
	Options []string
	Value   string
}
type Aput struct {
	Link, Param, Value string
}
type RPut struct {
	Link     string
	DLink    string
	Inputs   []Inputs
	Count    string
	ListLink string
}
type SSWAL struct {
	Title, Type, Text string
}
type SPackageEdit struct {
	Type, Mainf, Shutdown, Initf, Sessionf                  string
	IType, Package, Port, Key, Name, Ffpage, Erpage, Domain Aput
	Css                                                     RPut
	Imports                                                 []RPut
	Variables                                               []RPut
	CssFiles                                                []RPut
	CreateVar                                               RPut
	CreateImport                                            RPut
	TName                                                   string
}
type DebugObj struct {
	PKG, Id, Username, RawLog, Time string
	Bugs                            []DebugNode
}
type DebugNode struct {
	Action, Line, CTA string
}
type PkgItem struct {
	ID       string    `json:"id"`
	Icon     string    `json:"icon"`
	Text     string    `json:"text"`
	Children []PkgItem `json:"children"`
	isXml    bool      `json:"isXml"`
	Parent   string    `json:"parent"`
	Link     string    `json:"link"`
	Type     string    `json:"type"`
	DType    string    `json:"dtype"`
	RType    string    `json:"rtype"`
	NType    string    `json:"ntype"`
	MType    string    `json:"mtype"`
	CType    string    `json:"ctype"`
	AppID    string    `json:"appid"`
}
type SROC struct {
	Name      string
	CompLog   []byte
	Build     bool
	Time, Pid string
}
type VHuf struct {
	Type, PKG string
	Edata     []byte
}
