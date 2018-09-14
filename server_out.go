package main

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/cheikhshift/db"
	"github.com/cheikhshift/gos/core"
	gosweb "github.com/cheikhshift/gos/web"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/fatih/color"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var store = sessions.NewCookieStore([]byte("something-secretive-is-what-a-gorrilla-needs"))

var Prod = true

var TemplateFuncStore template.FuncMap
var templateCache = gosweb.NewTemplateCache()

func StoreNetfn() int {
	TemplateFuncStore = template.FuncMap{"a": gosweb.Netadd, "s": gosweb.Netsubs, "m": gosweb.Netmultiply, "d": gosweb.Netdivided, "js": gosweb.Netimportjs, "css": gosweb.Netimportcss, "sd": gosweb.NetsessionDelete, "sr": gosweb.NetsessionRemove, "sc": gosweb.NetsessionKey, "ss": gosweb.NetsessionSet, "sso": gosweb.NetsessionSetInt, "sgo": gosweb.NetsessionGetInt, "sg": gosweb.NetsessionGet, "form": gosweb.Formval, "eq": gosweb.Equalz, "neq": gosweb.Nequalz, "lte": gosweb.Netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "KanBan": NetKanBan, "bKanBan": NetbKanBan, "cKanBan": NetcKanBan, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "NoStruct": NetstructNoStruct, "isNoStruct": NetcastNoStruct, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf}
	return 0
}

var FuncStored = StoreNetfn()

type dbflf db.O

func renderTemplate(w http.ResponseWriter, p *gosweb.Page) {
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path : web%s.tmpl reason : %s", p.R.URL.Path, n))

			DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
			w.WriteHeader(http.StatusInternalServerError)

			pag, err := loadPage("")

			if err != nil {
				log.Println(err.Error())
				return
			}

			if pag.IsResource {
				w.Write(pag.Body)
			} else {
				pag.R = p.R
				pag.Session = p.Session
				renderTemplate(w, pag) //"

			}
		}
	}()

	// TemplateFuncStore

	if _, ok := templateCache.Get(p.R.URL.Path); !ok || !Prod {
		var tmpstr = string(p.Body)
		var localtemplate = template.New(p.R.URL.Path)

		localtemplate.Funcs(TemplateFuncStore)
		localtemplate.Parse(tmpstr)
		templateCache.Put(p.R.URL.Path, localtemplate)
	}

	outp := new(bytes.Buffer)
	err := templateCache.JGet(p.R.URL.Path).Execute(outp, p)
	if err != nil {
		log.Println(err.Error())
		DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			return
		}
		pag.R = p.R
		pag.Session = p.Session

		if pag.IsResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) // ""

		}
		return
	}

	// p.Session.Save(p.R, w)

	var outps = outp.String()
	var outpescaped = html.UnescapeString(outps)
	outp = nil
	fmt.Fprintf(w, outpescaped)

}

// Access you .gxml's end tags with
// this http.HandlerFunc.
// Use MakeHandler(http.HandlerFunc) to serve your web
// directory from memory.
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if attmpt := apiAttempt(w, r); !attmpt {
			fn(w, r)
		}

	}
}

func mResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}
func apiAttempt(w http.ResponseWriter, r *http.Request) (callmet bool) {
	var response string
	response = ""
	var session *sessions.Session
	var er error
	if session, er = store.Get(r, "session-"); er != nil {
		session, _ = store.New(r, "session-")
	}

	if strings.Contains(r.URL.Path, "/api/get") {

		me := SoftUser{Email: "Strukture user", Username: "Strukture user"}
		if r.FormValue("type") == "0" {

			mpk := []bson.M{}

			apps := getApps()

			for _, v := range apps {
				if v.Name != "" {
					gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + v.Name + "/gos.gxml")
					appCo := []PkgItem{}
					Childtm := []PkgItem{}
					for _, b := range v.Groups {
						tmpls := []PkgItem{}

						for _, tm := range gos.Templates.Templates {
							if tm.Bundle == b {
								tmpls = append(tmpls, PkgItem{Type: "5", AppID: v.Name, Icon: "fa fa-page", DType: "5&tmpl=" + b + "/" + tm.Name, Text: tm.Name, ID: v.Name + "@pkg:" + b + "/" + tm.Name})
							}
						}

						Childtm = append(Childtm, PkgItem{AppID: v.Name, Text: b, Icon: "fa fa-square", CType: "4&bundle=" + b, DType: "4&bundle=" + b, RType: "4&bundle=" + b, Children: tmpls})

					}
					appCo = append(appCo, PkgItem{AppID: v.Name, Text: "Template bundles", Icon: "fa fa-pencil-square", CType: "3", Children: Childtm})

					var folders []PkgItem
					var pkgpath = core.TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + v.Name + "/"
					if Windows {
						pkgpath = strings.Replace(pkgpath, "/", "\\", -1)
					}
					_ = filepath.Walk(pkgpath+"web", func(path string, file os.FileInfo, _ error) error {
						//fmt.Println(path)
						if file.IsDir() {
							lpathj := strings.Replace(path, pkgpath+"web", "", -1)

							var loca PkgItem = PkgItem{AppID: v.Name, Text: lpathj, Icon: "fa fa-folder", Children: []PkgItem{}}

							loca.CType = "5&path=" + lpathj
							loca.DType = "6&isDir=Yes&path=" + lpathj

							loca.MType = "6&path=" + lpathj

							files, _ := ioutil.ReadDir(path)

							for _, f := range files {
								if !f.IsDir() {
									var mjk string
									mjk = strings.Replace(path, pkgpath+"web", "", -1) + "/" + f.Name()
									if Windows {
										mjk = strings.Replace(mjk, "/", "\\", -1)
									}

									loca.Children = append(loca.Children, PkgItem{AppID: v.Name, Text: f.Name(), Icon: "fa fa-page", Type: "6", ID: v.Name + "@pkg:" + mjk, MType: "6&path=" + mjk, DType: "6&isDir=No&path=" + mjk})

								}
							}

							folders = append(folders, loca)

						}
						//fmt.Println(file,path,file.Name,file.IsDir())
						//   var loca PkgItem = PkgItem{AppID:v.Name,Text: file.Name(),Icon: "fa fa-folder"}

						return nil
					})
					var goFiles []PkgItem

					_ = filepath.Walk(pkgpath, func(path string, file os.FileInfo, _ error) error {
						//fmt.Println(path)
						if file.IsDir() {
							lpathj := strings.Replace(path, pkgpath, "", -1)

							var loca PkgItem = PkgItem{AppID: v.Name, Text: lpathj, Icon: "fa fa-circle", Children: []PkgItem{}}
							hasgo := false
							files, _ := ioutil.ReadDir(path)
							for _, f := range files {
								if !f.IsDir() && strings.Contains(f.Name(), ".go") {

									var mjk string
									mjk = strings.Replace(path, pkgpath, "", -1) + "/" + f.Name()
									if Windows {
										mjk = strings.Replace(mjk, "/", "\\", -1)
									}
									hasgo = true
									loca.Children = append(loca.Children, PkgItem{AppID: v.Name, Text: f.Name(), Icon: "fa fa-code", Type: "60", ID: v.Name + "@pkg:" + mjk, MType: "60&path=" + mjk, DType: "60&isDir=No&path=" + mjk})

								}
							}

							loca.CType = "50&path=" + lpathj
							loca.DType = "60&isDir=Yes&path=" + lpathj

							loca.MType = "60&path=" + lpathj

							if hasgo {
								goFiles = append(goFiles, loca)
							}

						}
						//fmt.Println(file,path,file.Name,file.IsDir())
						//   var loca PkgItem = PkgItem{AppID:v.Name,Text: file.Name(),Icon: "fa fa-folder"}

						return nil
					})

					appCo = append(appCo, PkgItem{AppID: v.Name, Text: "Web Resources", CType: "5&path=/", Children: folders, Icon: "fa fa-folder"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Text: "Go SRC", CType: "50&path=/", Children: goFiles, Icon: "fa fa-cube"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "16", Text: "Logs", Icon: "fa fa-list"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "18", Text: "Testing", Icon: "fa fa-flask"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "300", Text: "KanBan board", Icon: "fa fa-briefcase"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "7", Text: "Build center", Icon: "fa fa-server"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "8", Text: "Interfaces", Icon: "fa fa-share-alt"})
					//appCo = append(appCo, PkgItem{AppID:v.Name,Type:"9",Text: "Interface funcs",Icon: "fa fa-share-alt-square"} )
					appCo = append(appCo, PkgItem{Type: "10", AppID: v.Name, Text: "Template pipelines", Icon: "fa fa-exchange"})

					appCo = append(appCo, PkgItem{AppID: v.Name, Type: "11", Text: "Web services", Icon: "fa fa-circle-o-notch"})

					//appCo = append(appCo, PkgItem{AppID:v.Name,Type:"12",Text: "Timers",Icon: "fa fa-clock-o"} )

					rootel := bson.M{"dtype": "3", "text": v.Name, "type": "1", "id": v.Name, "children": appCo, "appid": v.Name, "btype": "on"}
					if v.Type == "webapp" {
						rootel["icon"] = "fa fa-globe"
					} else if v.Type == "bind" {
						rootel["icon"] = "fa fa-mobile"
					} else {
						rootel["icon"] = "fa fa-gift"
					}

					//append to children
					//add server in
					mpk = append(mpk, rootel)
				}
			}

			response = mResponse(mpk)
		} else if r.FormValue("type") == "1" {

			//get package
			sapp := NetgetApp(getApps(), r.FormValue("id"))
			prefix := "/api/put?type=0&id=" + sapp.Name
			gos, _ := core.LoadGos(os.ExpandEnv("$GOPATH") + "/src/" + sapp.Name + "/gos.gxml")

			//load gos

			//set params democss,port,key,name,type
			editor := sPackageEdit{Type: sapp.Type, TName: sapp.Name}
			editor.IType = Aput{Link: prefix, Param: "app", Value: gos.Type}

			editor.Port = Aput{Link: prefix, Param: "port", Value: gos.Port}
			editor.Key = Aput{Link: prefix, Param: "key", Value: gos.Key}
			editor.Domain = Aput{Link: prefix, Param: "domain", Value: gos.Domain}
			editor.Erpage = Aput{Link: prefix, Param: "erpage", Value: gos.ErrorPage}
			editor.Ffpage = Aput{Link: prefix, Param: "fpage", Value: gos.NPage}
			editor.Name = Aput{Link: prefix, Param: "Name", Value: sapp.Name}
			editor.Package = Aput{Link: "/api/put?type=16&pkg=" + sapp.Name, Param: "npk", Value: gos.Package}
			editor.Mainf = gos.Main
			editor.Initf = gos.Init_Func
			editor.Sessionf = gos.Session

			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "is", Type: "text", Text: "Variable type"})
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Variable name"})
			editor.CreateVar = rPut{Count: "4", Link: "/api/create?type=0&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=2&pkg=" + sapp.Name}

			varf = []Inputs{}
			varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Package path"})

			editor.CreateImport = rPut{Count: "6", Link: "/api/create?type=1&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=3&pkg=" + sapp.Name}
			varf = []Inputs{}
			varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Path to css lib"})
			editor.Css = rPut{Count: "6", Link: "/api/create?type=2&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=4&pkg=" + sapp.Name}

			response = NetbPackageEdit(editor)

		} else if r.FormValue("type") == "2" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.Variables {
				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "is", Type: "text", Text: "Variable type", Value: v.Type})
				varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Variable name", Value: v.Name})
				response = response + NetbRPUT(rPut{DLink: "/api/delete?type=0&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Count: "4", Link: "/api/act?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Inputs: varf})
			}

		} else if r.FormValue("type") == "3" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.RootImports {
				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Package path", Value: v.Src})

				response = response + NetbRPUT(rPut{DLink: "/api/delete?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Count: "6", Link: "/api/act?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Inputs: varf})
			}

		} else if r.FormValue("type") == "4" {
			sapp := NetgetApp(getApps(), r.FormValue("pkg"))

			for _, v := range sapp.Css {

				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Path to css lib", Value: v})

				response = response + NetbRPUT(rPut{DLink: "/api/delete?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v, Count: "6", Link: "/api/act?type=3&pkg=" + r.FormValue("pkg") + "&id=" + v, Inputs: varf})
			}

		} else if r.FormValue("type") == "5" {
			id := strings.Split(r.FormValue("id"), "@pkg:")
			data, _ := ioutil.ReadFile(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/tmpl/" + id[1] + ".tmpl")

			data = []byte(html.EscapeString(string(data)))
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

			template := NetgetTemplate(gos.Templates.Templates, id[1])

			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Value: template.Struct, Name: "struct", Text: "Interface to use with template"})

			response = NetbTemplateEdit(TemplateEdits{SavesTo: "tmpl/" + id[1] + ".tmpl", ID: NetRandTen(), PKG: r.FormValue("space"), Mime: "html", File: data, Settings: rPut{Link: "/api/put?type=2&id=" + id[1] + "&pkg=" + r.FormValue("space"), Inputs: varf, Count: "6"}})
		} else if r.FormValue("type") == "6" {
			id := strings.Split(r.FormValue("id"), "@pkg:")
			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/web" + id[1]
			var ftype string
			if strings.Contains(filep, ".css") {
				ftype = "css"
			} else if strings.Contains(filep, ".js") {
				ftype = "javascript"
			} else if strings.Contains(filep, ".html") {
				ftype = "html"
			} else if strings.Contains(filep, ".tmpl") {
				ftype = "html"
				//add auto complete linking
			}
			data, _ := ioutil.ReadFile(filep)
			data = []byte(html.EscapeString(string(data)))
			response = NetbWebRootEdit(WebRootEdits{SavesTo: id[1], Type: ftype, File: data, ID: NetRandTen(), PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "60" {
			id := strings.Split(r.FormValue("id"), "@pkg:")
			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + id[1]

			data, _ := ioutil.ReadFile(filep)
			data = []byte(html.EscapeString(string(data)))
			response = NetbWebRootEdittwo(WebRootEdits{SavesTo: id[1], Type: "golang", File: data, ID: NetRandTen(), PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "7" {
			sapp := NetgetApp(getApps(), r.FormValue("space"))
			response = NetbROC(sROC{Name: r.FormValue("space"), Build: sapp.Passed, Time: sapp.LatestBuild, Pid: sapp.Pid})
		} else if r.FormValue("type") == "8" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/structs.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = NetbStructEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "9" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/objects.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = NetbObjectEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "10" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/methods.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = NetbMethodEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "11" {

			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "path", Type: "text", Text: "Endpoint path"})
			kput := rPut{ListLink: "/api/get?type=13&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=7&space=" + r.FormValue("space")}
			response = NetbEndpointEditor(TEditor{CreateForm: kput, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "12" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Timer name"})
			kput := rPut{ListLink: "/api/get?type=14&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=8&space=" + r.FormValue("space")}
			response = NetbTimerEditor(TEditor{CreateForm: kput, PKG: r.FormValue("space")})
		} else if r.FormValue("type") == "13" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

			for _, v := range gos.Endpoints.Endpoints {

				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "path", Type: "text", Text: "Endpoint path", Value: v.Path})
				//varf = append(varf, Inputs{Name:"method", Type:"text",Text:"Endpoint method",Value:v.Method})
				varf = append(varf, Inputs{Name: "typ", Type: "text", Text: "Request type : GET,POST,PUT,DELETE,f,star...", Value: v.Type})

				response = response + NetbRPUT(rPut{DLink: "/api/delete?type=7&pkg=" + r.FormValue("space") + "&path=" + v.Id, Link: "/api/put?type=9&id=" + v.Id + "&pkg=" + r.FormValue("space"), Count: "12", Inputs: varf}) + addjsstr

			}

		} else if r.FormValue("type") == "13r" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.Endpoints.Endpoints {

				if v.Id == r.FormValue("id") {
					id := NetRandTen()
					response = NetbTemplateEditTwo(TemplateEdits{SavesTo: "gosforceasapi/" + r.FormValue("id") + "++()/", ID: id, PKG: r.FormValue("pkg"), Mime: "golang", File: []byte(v.Method)})
				}
			}

		} else if r.FormValue("type") == "14" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

			for _, v := range gos.Timers.Timers {

				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Timer name", Value: v.Name})
				varf = append(varf, Inputs{Name: "interval", Type: "number", Text: "Interval", Value: v.Interval})
				varf = append(varf, Inputs{Name: "unit", Type: "text", Text: "Timer refresh unit", Value: v.Unit})
				varf = append(varf, Inputs{Name: "method", Type: "text", Text: "Method to execute", Value: v.Method})
				response = response + NetbRPUT(rPut{DLink: "/api/delete?type=8&pkg=" + r.FormValue("space") + "&name=" + v.Name, Link: "/api/put?type=10&id=" + v.Name + "&pkg=" + r.FormValue("space"), Count: "2", Inputs: varf})

			}

		} else if r.FormValue("type") == "15" {

			tempx := NetbuSettings(USettings{StripeID: me.StripeID, LastPaid: "Date", Email: me.Email})
			response = NetbModal(sModal{Title: "Account settings", Body: tempx, Color: "orange"})
		} else if r.FormValue("type") == "16" {

			response = NetbDebugger(DebugObj{PKG: r.FormValue("space"), Username: ""})
		} else if r.FormValue("type") == "17" {

			var tDebugNode DebugObj

			if r.FormValue("id") == "Server" {
				tDebugNode = DebugObj{Time: "Server", Bugs: []DebugNode{}}
				gp := os.ExpandEnv("$GOPATH")
				os.Chdir(gp + "/src/" + r.FormValue("space"))
				//main.log
				rlog, err := ioutil.ReadFile("main.log")
				if err != nil {
					tDebugNode.RawLog = err.Error()
				} else {
					tDebugNode.RawLog = string(rlog)
				}
			} else {
				logs := GetLogs(r.FormValue("space"))

				for _, logg := range logs {
					if logg.Time == r.FormValue("id") {
						tDebugNode = logg
					}
				}
			}

			response = NetbDebuggerNode(tDebugNode)

		} else if r.FormValue("type") == "18" {

			response = NetbEndpointTesting(Dex{Misc: r.FormValue("space")})

		} else if r.FormValue("type") == "300" {

			response = NetbKanBan(Dex{Misc: r.FormValue("space")})

		}
		callmet = true

	}
	if r.Method == "RESET" {
		return true
	} else if !callmet && gosweb.UrlAtZ(r.URL.Path, "/api/socket") {

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		AddConnection(c)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			if len(message) != 0 {
				Broadcast(message)
			}

		}

		return true

		callmet = true
	} else if isURL := (r.URL.Path == "/api/pkg-bugs" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		bugs := GetLogs(r.FormValue("pkg"))
		sapp := NetgetApp(getApps(), r.FormValue("pkg"))
		if len(bugs) == 0 || sapp.Passed {
			response = "{}"
		} else {
			response = mResponse(bugs[0])
		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/kanban" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		pkgName := r.FormValue("pkg")
		response = mResponse(getKanBan(pkgName))

		callmet = true
	} else if isURL := (r.URL.Path == "/api/git" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		pkgName := r.FormValue("pkg")
		cmd := r.FormValue("cmd")

		if cmd == "commit" {
			mess := r.FormValue("message")
			response = mResponse(commitGit(pkgName, mess))
		}

		if cmd == "push" {
			pushGit(pkgName)
			response = mResponse(false)
		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/kanban" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		pkgName := r.FormValue("pkg")
		payload := r.FormValue("payload")

		saveKanBan(pkgName, payload)

		response = "OK"

		callmet = true
	} else if isURL := (r.URL.Path == "/api/empty" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		ClearLogs(r.FormValue("pkg"))
		response = NetbAlert(Alertbs{Type: "success", Text: "Your build logs are cleared."})

		callmet = true
	} else if isURL := (r.URL.Path == "/api/tester/" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")
		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		logfull, _ := core.RunCmdSmart("gos " + r.FormValue("mode") + " " + r.FormValue("c"))
		response = html.EscapeString(logfull)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/create" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		//me := &SoftUser{Email:"Strukture user", Username:"Strukture user"}

		if r.FormValue("type") == "0" {
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			gos.Add("var", r.FormValue("is"), r.FormValue("name"))
			// fmt.Println(gos)

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "1" {
			//import
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			gos.Add("import", "", r.FormValue("src"))
			//fmt.Println(gos)

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
		} else if r.FormValue("type") == "2" {
			//css
			apps := getApps()
			app := NetgetApp(getApps(), r.FormValue("pkg"))
			app.Css = append(app.Css, r.FormValue("src"))
			apps = NetupdateApp(getApps(), r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "3" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Bundle name"})

			response = NetbForm(Forms{Link: "/api/act?type=4&pkg=" + r.FormValue("pkg"), CTA: "Create Bundle", Class: "warning", Inputs: varf})

		} else if r.FormValue("type") == "4" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Template name"})

			response = NetbForm(Forms{Link: "/api/act?type=5&pkg=" + r.FormValue("pkg") + "&bundle=" + r.FormValue("bundle"), CTA: "Create Template file", Class: "warning", Inputs: varf})

		} else if r.FormValue("type") == "5" {
			//prefix pkg
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Text: "Path"})
			varf = append(varf, Inputs{Type: "hidden", Name: "basesix"})
			varf = append(varf, Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

			response = NetbFSC(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=6&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
		} else if r.FormValue("type") == "50" {
			//prefix pkg
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Text: "Path"})
			varf = append(varf, Inputs{Type: "hidden", Name: "basesix"})
			varf = append(varf, Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

			response = NetbFSC(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=60&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
		} else if r.FormValue("type") == "6" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

			response = NetbMV(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=7&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
		} else if r.FormValue("type") == "60" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

			response = NetbMV(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=70&pkg=" + r.FormValue("pkg") + "&folder=" + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/delete" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		if r.FormValue("type") == "0" {

			//type pkg id
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			gos.Delete("var", r.FormValue("id"))

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "101" {

			//type pkg id
			plugins := getPlugins()
			newset := []string{}

			for _, v := range plugins {
				if v != r.FormValue("pkg") {
					newset = append(newset, v)
				}
			}

			plugins = newset

			savePlugins(plugins)

		} else if r.FormValue("type") == "1" {
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			gos.Delete("import", r.FormValue("id"))

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
		} else if r.FormValue("type") == "2" {
			apps := getApps()
			app := NetgetApp(apps, r.FormValue("pkg"))
			temp := []string{}
			for _, v := range app.Css {
				if v != r.FormValue("id") {
					temp = append(temp, v)
				}
			}
			app.Css = temp
			apps = NetupdateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "3" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = NetbDelete(DForm{Text: "Are you sure you want to delete the package " + r.FormValue("pkg"), Link: "type=3&pkg=" + r.FormValue("pkg")})

			} else {
				//delete
				os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg"))

				newapps := []App{}
				apps := getApps()
				for _, v := range apps {

					if v.Name != r.FormValue("pkg") {
						newapps = append(newapps, v)
					}

				}
				saveApps(newapps)

				response = NetbAlert(Alertbs{Type: "success", Text: "Success package " + r.FormValue("pkg") + " was removed. Please reload page to close all linked resources.", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "4" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = NetbDelete(DForm{Text: "Are you sure you want to delete the bundle " + r.FormValue("bundle") + " and all of its sub templates", Link: "type=4&bundle=" + r.FormValue("bundle") + "&pkg=" + r.FormValue("pkg")})
			} else {
				//delete bundle
				apps := getApps()
				sapp := NetgetApp(apps, r.FormValue("pkg"))

				replac := []string{}

				for _, v := range sapp.Groups {

					if r.FormValue("bundle") != v {
						replac = append(replac, v)
					}

				}

				sapp.Groups = replac
				apps = NetupdateApp(apps, sapp.Name, sapp)
				saveApps(apps)
				gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
				os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/tmpl/" + r.FormValue("bundle"))
				gos.Delete("bundle", r.FormValue("bundle"))

				gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

				response = NetbAlert(Alertbs{Type: "success", Text: "Success bundle was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "5" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = NetbDelete(DForm{Text: "Are you sure you want to delete the template " + r.FormValue("tmpl"), Link: "type=5&tmpl=" + r.FormValue("tmpl") + "&pkg=" + r.FormValue("pkg")})
			} else {
				//delete

				os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/tmpl/" + r.FormValue("tmpl") + ".tmpl")

				gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
				parsedStr := strings.Split(r.FormValue("tmpl"), "/")
				gos.Delete("template", parsedStr[len(parsedStr)-1])

				gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

				response = NetbAlert(Alertbs{Type: "success", Text: "Success template " + r.FormValue("tmpl") + " was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "6" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = NetbDelete(DForm{Text: "Are you sure you want to delete the web resource at " + r.FormValue("path"), Link: "type=6&conf=do&path=" + r.FormValue("path") + "&pkg=" + r.FormValue("pkg")})

			} else {
				//delete
				if r.FormValue("isDir") == "Yes" {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/web" + r.FormValue("path"))
				} else {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/web" + r.FormValue("path"))
				}
				response = NetbAlert(Alertbs{Type: "success", Text: "Success resource at " + r.FormValue("path") + " was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "60" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = NetbDelete(DForm{Text: "Are you sure you want to delete the resource at " + r.FormValue("path"), Link: "type=60&conf=do&path=" + r.FormValue("path") + "&pkg=" + r.FormValue("pkg")})

			} else {
				//delete
				if r.FormValue("isDir") == "Yes" {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/" + r.FormValue("path"))
				} else {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/" + r.FormValue("path"))
				}
				response = NetbAlert(Alertbs{Type: "success", Text: "Success resource at " + r.FormValue("path") + " removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "7" {
			//type pkg path name
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			gos.DeleteEnd(r.FormValue("path"))

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "8" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			gos.Delete("timer", r.FormValue("name"))

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
		} else if r.FormValue("type") == "9" {

		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/rename" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		callmet = true
	} else if isURL := (r.URL.Path == "/api/new" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		if r.FormValue("type") == "0" {
			inputs := []Inputs{}
			inputs = append(inputs, Inputs{Type: "text", Name: "name", Misc: "required", Text: "Package Name"})
			inputs = append(inputs, Inputs{Type: "hidden", Name: "type", Value: "0"})
			response = NetbModal(sModal{Body: "", Title: "New Package", Color: "#ededed", Form: Forms{Link: "/api/act", CTA: "Create Package", Class: "warning btn-block", Buttons: []sButton{}, Inputs: inputs}})
		} else if r.FormValue("type") == "100" {
			inputs := []Inputs{}
			inputs = append(inputs, Inputs{Type: "text", Name: "name", Misc: "required", Text: "Plugin install path"})
			inputs = append(inputs, Inputs{Type: "hidden", Name: "type", Value: "100"})
			response = NetbModal(sModal{Body: "", Title: "PLUGINS", Color: "#ededed", Form: Forms{Link: "/api/act", CTA: "ADD", Class: "warning btn-block", Buttons: []sButton{}, Inputs: inputs}})
		} else if r.FormValue("type") == "101" {

			response = bPluginList(gosweb.NoStruct{})
		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/act" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		if r.FormValue("type") == "0" {
			apps := getApps()
			apps = append(apps, App{Type: "webapp", Name: r.FormValue("name")})

			dir := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("name") + "/gos.gxml"
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				core.RunCmdB(os.ExpandEnv("$GOPATH") + "/bin/gos make " + r.FormValue("name"))
			}
			//Users.Update(bson.M{"uid": me.UID}, me)
			saveApps(apps)
			response = NetbAlert(Alertbs{Type: "warning", Text: "Success package " + r.FormValue("name") + " was created!", Redirect: "javascript:updateTree()"})
		} else if r.FormValue("type") == "100" {
			plugins := getPlugins()
			plugins = append(plugins, r.FormValue("name"))

			//Users.Update(bson.M{"uid": me.UID}, me)

			_, err := core.RunCmdSmart("go get " + r.FormValue("name"))
			if err != nil {
				response = NetbAlert(Alertbs{Type: "warning", Text: "Error, could not find plugin.", Redirect: "#"})
			} else {
				savePlugins(plugins)
				response = NetbAlert(Alertbs{Type: "success", Text: "Success plugin " + r.FormValue("name") + " installed! Reload the page to activate plugin.", Redirect: "javascript:GetPlugins()"})
			}
		} else if r.FormValue("type") == "1" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			//update va
			gos.Update("var", r.FormValue("id"), core.GlobalVariables{Name: r.FormValue("name"), Type: r.FormValue("is")})

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			response = "Variable saved!"

		} else if r.FormValue("type") == "2" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			//update var
			gos.Update("import", r.FormValue("id"), core.Import{Src: r.FormValue("src")})

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			response = "Import saved!"

		} else if r.FormValue("type") == "3" {
			apps := getApps()
			app := NetgetApp(apps, r.FormValue("pkg"))
			temp := []string{}
			for _, v := range app.Css {
				if v != r.FormValue("id") {
					temp = append(temp, v)
				} else {
					temp = append(temp, r.FormValue("src"))
				}
			}
			app.Css = temp
			apps = NetupdateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "4" {
			apps := getApps()
			app := NetgetApp(apps, r.FormValue("pkg"))

			app.Groups = append(app.Groups, r.FormValue("name"))
			os.MkdirAll(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/tmpl/"+r.FormValue("name"), 0777)
			apps = NetupdateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "5" {
			//app := NetgetApp(me.Apps, r.FormValue("pkg"))
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			//update var
			os.Create(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/tmpl/" + r.FormValue("bundle") + "/" + r.FormValue("name") + ".tmpl")
			gos.AddS("template", core.Template{Name: r.FormValue("name"), Bundle: r.FormValue("bundle"), TemplateFile: r.FormValue("bundle") + "/" + r.FormValue("name")})

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
		} else if r.FormValue("type") == "6" {

			if r.FormValue("fmode") == "touch" {

				os.Create(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/web" + r.FormValue("prefix") + "/" + r.FormValue("path"))

			} else if r.FormValue("fmode") == "dir" {
				fmt.Println(os.MkdirAll(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web"+r.FormValue("prefix")+"/"+r.FormValue("path"), 0777))
			} else if r.FormValue("fmode") == "upload" {
				ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web"+r.FormValue("prefix")+"/"+r.FormValue("path"), core.Decode64(nil, []byte(r.FormValue("basesix"))), 0777)
			}

		} else if r.FormValue("type") == "60" {

			if r.FormValue("fmode") == "touch" {
				addstr := ""
				if !strings.Contains(r.FormValue("path"), ".go") {
					addstr = ".go"
				}
				_, err := os.Create(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + r.FormValue("prefix") + r.FormValue("path") + addstr)

				if err != nil {
					npath := strings.Split(r.FormValue("path"), "/")
					os.MkdirAll(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("prefix")+strings.Join(npath[:len(npath)-1], "/"), 0777)
					os.Create(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + r.FormValue("prefix") + r.FormValue("path") + addstr)

				}

			} else if r.FormValue("fmode") == "dir" {
				os.MkdirAll(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("prefix")+r.FormValue("path"), 0777)
			} else if r.FormValue("fmode") == "upload" {
				ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("prefix")+r.FormValue("path"), core.Decode64(nil, []byte(r.FormValue("basesix"))), 0777)
			}

		} else if r.FormValue("type") == "7" {
			err := os.Rename(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web"+r.FormValue("prefix"), os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web/"+r.FormValue("path"))

			if err != nil {
				response = NetbAlert(Alertbs{Type: "danger", Text: "Failed to move resource : " + err.Error()})
			} else {
				response = NetbAlert(Alertbs{Type: "success", Text: "Operation succeeded"})
			}

		} else if r.FormValue("type") == "70" {
			err := os.Rename(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("prefix"), os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/"+r.FormValue("path"))
			if err != nil {
				response = NetbAlert(Alertbs{Type: "danger", Text: "Failed to move resource : " + err.Error()})
			} else {
				response = NetbAlert(Alertbs{Type: "success", Text: "Operation succeeded"})
			}
		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/put" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		me := SoftUser{Email: "Strukture user", Username: "Strukture user"}

		if r.FormValue("type") == "0" {

			//fmt.Println(m)
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("id") + "/gos.gxml")

			gos.Set(r.FormValue("put"), r.FormValue("var"))
			//   fmt.Println(gos)

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("id") + "/gos.gxml")

		} else if r.FormValue("type") == "1" {
			ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/"+r.FormValue("target"), []byte(r.FormValue("data")), 0644)
			response = NetbAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
		} else if r.FormValue("type") == "2" {
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			gos.Update("template", r.FormValue("id"), r.FormValue("struct"))
			// fmt.Println(gos)

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			response = "Template interface saved!"
		} else if r.FormValue("type") == "3" {
			ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web"+r.FormValue("target"), []byte(r.FormValue("data")), 0777)

			response = NetbAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
		} else if r.FormValue("type") == "30" {
			ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("target"), []byte(r.FormValue("data")), 0777)

			response = NetbAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
		} else if r.FormValue("type") == "4" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/structs.dsl"
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			ioutil.WriteFile(filep, []byte(r.FormValue("data")), 0644)
			//marhal and add
			vgos := core.CreateVGos(filep)
			//fmt.Println(vgos,"Gos")
			gos.MStructs(vgos.Structs)
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = NetbAlert(Alertbs{Type: "warning", Text: "Interfaces saved!"})

		} else if r.FormValue("type") == "5" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/objects.dsl"
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			ioutil.WriteFile(filep, []byte(r.FormValue("data")), 0644)
			//marhal and add
			vgos := core.CreateVGos(filep)
			//fmt.Println(vgos,"Gos")
			gos.MObjects(vgos.Objects)
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = NetbAlert(Alertbs{Type: "warning", Text: "Objects saved!"})

		} else if r.FormValue("type") == "6" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/methods.dsl"
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			ioutil.WriteFile(filep, []byte(r.FormValue("data")), 0644)
			//marhal and add
			vgos := core.CreateVGos(filep)
			//fmt.Println(vgos,"Gos")
			gos.MMethod(vgos.Methods)
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = NetbAlert(Alertbs{Type: "warning", Text: "Pipelines saved!"})

		} else if r.FormValue("type") == "7" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")
			//write file

			gos.Add("end", "", r.FormValue("path"))
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

		} else if r.FormValue("type") == "8" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")
			//write file

			gos.Add("timer", "", r.FormValue("name"))
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")
		} else if r.FormValue("type") == "9" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file

			gos.Update("end", r.FormValue("id"), core.Endpoint{Path: r.FormValue("path"), Type: r.FormValue("typ")})
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = "OK"

		} else if r.FormValue("type") == "13r" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file

			gos.UpdateMethod(r.FormValue("target"), r.FormValue("data"))
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = NetbAlert(Alertbs{Type: "warning", Text: "Endpoint code saved!"})

		} else if r.FormValue("type") == "10" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file

			gos.Update("timer", r.FormValue("id"), core.Timer{Name: r.FormValue("name"), Method: r.FormValue("method"), Unit: r.FormValue("unit"), Interval: r.FormValue("interval")})
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			response = "OK"
		} else if r.FormValue("type") == "11" {

			hasher := sha512.New512_256()

			if string(hasher.Sum([]byte(r.FormValue("cpassword")))) == string(me.Password) {
				me.Password = hasher.Sum([]byte(r.FormValue("npassword")))
				response = NetbAlert(Alertbs{Type: "success", Text: "Password updated"})
			} else {
				response = NetbAlert(Alertbs{Type: "danger", Text: "Error incorrect current password"})
			}

		} else if r.FormValue("type") == "12" {
			me.Email = r.FormValue("email")
			response = NetbAlert(Alertbs{Type: "success", Text: "Email updated"})
		} else if r.FormValue("type") == "13" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			gos.Main = r.FormValue("data")
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "14" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			gos.Init_Func = r.FormValue("data")
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "15" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			gos.Session = r.FormValue("data")
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "16" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			gos.Package = r.FormValue("var")
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		} else if r.FormValue("type") == "17" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			//write file
			gos.Shutdown = r.FormValue("data")
			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		}

		callmet = true
	} else if isURL := (r.URL.Path == "/api/build" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")

		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		os.RemoveAll(gp + "/src/" + r.FormValue("pkg") + "/bindata.go")
		os.RemoveAll(gp + "/src/" + r.FormValue("pkg") + "/application.go")
		logBuilt, _ := core.RunCmdSmart("gos --run --buildcheck")
		fmt.Println(logBuilt)
		passed := false

		if !strings.Contains(logBuilt, "Your build failed,") {
			logBuilt, _ = core.RunCmdSmart("go build")
			if logBuilt != "" {

				debuglog := DebugObj{r.FormValue("pkg"), NetRandTen(), "", logBuilt, time.Now().String(), []DebugNode{}}

				logs := strings.Split(logBuilt, "\n")

				for linen, log := range logs {
					if linen != 0 {
						linedt := strings.Split(log, ":")
						//	fmt.Println(len(linedt))
						dnode := DebugNode{}
						//fmt.Println(linedt)

						//src
						if len(linedt) > 2 {
							dnode.Action = "edit:" + linedt[0] + ":" + linedt[1]
							dnode.CTA = "Update " + linedt[0] + " on line " + linedt[1]
							dnode.Line = strings.Join(linedt[2:], " - ")
							debuglog.Bugs = append(debuglog.Bugs, dnode)
						}

					}
				}

				AddtoLogs(debuglog)
				response = NetbAlert(Alertbs{Type: "danger", Text: "Your build failed, checkout the logs to see why!"})

			} else {
				passed = true
				response = NetbAlert(Alertbs{Type: "success", Text: "Your build passed!"})
			}
		} else {
			debuglog := DebugObj{r.FormValue("pkg"), NetRandTen(), "", logBuilt, time.Now().String(), []DebugNode{}}
			fPart := strings.Split(logBuilt, "Full compiler build log :")
			logs := strings.Split(fPart[1], "\n")

			for linen, log := range logs {
				if linen != 0 {
					linedt := strings.Split(log, ":")
					//	fmt.Println(len(linedt))
					dnode := DebugNode{}
					//	fmt.Println(linedt)
					if linedt[0] == "./application.go" {
						if len(linedt) > 2 {
							il, _ := strconv.Atoi(linedt[1])
							actline := FindString("./application.go", il)
							actline = strings.TrimSpace(actline)
							//find line
							inStructs := FindLine("./structs.dsl", actline)
							if inStructs == -1 {

								inMethods := FindLine("./methods.dsl", actline)
								if inMethods == -1 {

									gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

									inMain := FindinString(gos.Main, actline)

									if inMain == -1 {

										inInit := FindinString(gos.
											Init_Func, actline)
										if inInit == -1 {
											//	fmt.Println(actline)
											for _, v := range gos.Endpoints.Endpoints {

												inEndpoint := FindinString(v.Method, actline)
												if inEndpoint != -1 {
													enode := DebugNode{}
													enode.Action = "service:" + v.Path + " - " + v.Type + ":" + strconv.Itoa(inEndpoint) + ":" + v.Id
													enode.Line = strings.Join(linedt[2:], " - ")
													enode.CTA = "Update webservice " + v.Path
													debuglog.Bugs = append(debuglog.Bugs, enode)
												}
											}
										} else {
											linestring := strconv.Itoa(inInit)
											dnode.Action = "init:" + linestring
											dnode.CTA = "Update Init()"
											dnode.Line = strings.Join(linedt[2:], " - ")
											debuglog.Bugs = append(debuglog.Bugs, dnode)
										}

									} else {
										linestring := strconv.Itoa(inMain)
										dnode.Action = "main:" + linestring
										dnode.CTA = "Update main()"
										dnode.Line = strings.Join(linedt[2:], " - ")
										debuglog.Bugs = append(debuglog.Bugs, dnode)
									}

								} else {
									linestring := strconv.Itoa(inMethods)
									dnode.Action = "meth:" + linestring
									dnode.CTA = "Update pipelines."
									dnode.Line = strings.Join(linedt[2:], " - ")
									debuglog.Bugs = append(debuglog.Bugs, dnode)
								}

							} else {
								linestring := strconv.Itoa(inStructs)
								dnode.Action = "structs:" + linestring
								dnode.CTA = "Update interfaces."
								dnode.Line = strings.Join(linedt[2:], " - ")
								debuglog.Bugs = append(debuglog.Bugs, dnode)
							}

						}
					} else {
						//src
						if len(linedt) > 2 {
							dnode.Action = "edit:" + linedt[0] + ":" + linedt[1]
							dnode.CTA = "Update " + linedt[0] + " on line " + linedt[1]
							dnode.Line = strings.Join(linedt[2:], " - ")
							debuglog.Bugs = append(debuglog.Bugs, dnode)
						}
					}

				}
			}

			AddtoLogs(debuglog)
			response = NetbAlert(Alertbs{Type: "danger", Text: "Your build failed, checkout the logs to see why!"})

		}

		//DebugLogs.Insert(dObj)

		apps := getApps()
		sapp := NetgetApp(apps, r.FormValue("pkg"))

		sapp.Passed = passed
		sapp.LatestBuild = time.Now().String()
		apps = NetupdateApp(apps, r.FormValue("pkg"), sapp)
		saveApps(apps)

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/start" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")

		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		apps := getApps()
		sapp := NetgetApp(apps, r.FormValue("pkg"))

		if sapp.Passed {

			if sapp.Pid != "" {
				core.RunCmdB("kill -3 " + sapp.Pid)
				response = NetbAlert(Alertbs{Type: "success", Text: "Build stopped."})
			}

			pkSpl := strings.Split(r.FormValue("pkg"), "/")
			if Windows {
				go core.RunCmd("cmd /C gos --run") //live reload on windows...
			} else {
				shscript := `#!/bin/bash  
									cmd="./` + pkSpl[len(pkSpl)-1] + ` "
									eval "${cmd}" >main.log &disown
									exit 0`

				ioutil.WriteFile("runsc", []byte(shscript), 0777)
				go core.RunCmdSmart("sh runsc &>/dev/null")
			}

			time.Sleep(time.Second * 5)
			raw, _ := ioutil.ReadFile("main.log")
			lines := strings.Split(string(raw), "\n")
			sapp.Pid = lines[0]
			if Windows {
				response = NetbAlert(Alertbs{Type: "success", Text: "Server up"})
			} else {
				response = NetbAlert(Alertbs{Type: "success", Text: "Your server is up at PID : " + sapp.Pid})
			}
		} else {
			response = NetbAlert(Alertbs{Type: "danger", Text: "Your latest build failed."})

		}

		//DebugLogs.Insert(dObj)
		apps = NetupdateApp(apps, r.FormValue("pkg"), sapp)
		saveApps(apps)

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/stop" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")

		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		apps := getApps()
		sapp := NetgetApp(apps, r.FormValue("pkg"))

		if sapp.Pid == "" {
			response = NetbAlert(Alertbs{Type: "danger", Text: "No build running."})
		} else {
			core.RunCmdB("kill -3 " + sapp.Pid)
			response = NetbAlert(Alertbs{Type: "success", Text: "Build stopped."})
		}

		//DebugLogs.Insert(dObj)
		apps = NetupdateApp(apps, r.FormValue("pkg"), sapp)
		saveApps(apps)

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/bin" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		os.Chdir(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg"))
		os.Remove(strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".binary.zip ")
		bPath := strings.Split(r.FormValue("pkg"), "/")
		gp := os.ExpandEnv("$GOPATH")
		//coreTemplate,_ := core.LoadGos(gp + "/src/" +  r.FormValue("pkg") + "/gos.gxml")

		//core.RunCmdB("gos ru " +  r.FormValue("pkg") + " gos.gxml web tmpl")
		os.Remove(gp + "/src/" + r.FormValue("pkg") + "/server_out.go")

		//core.Process(coreTemplate,gp + "/src/" +  r.FormValue("pkg"), "web","tmpl")
		core.RunCmdB("gos --export")
		zipname := strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".binary.zip"
		if Windows {
			bPath[len(bPath)-1] += ".exe"
			zipit(bPath[len(bPath)-1], zipname)
		} else {
			core.RunCmdB("zip -r " + zipname + " " + bPath[len(bPath)-1])
		}
		time.Sleep(500 * time.Millisecond)

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipname))
		http.ServeFile(w, r, zipname)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/export" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		os.Chdir(os.ExpandEnv("$GOPATH") + "/src/")
		os.Remove(strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".zip ")
		pkgpath := r.FormValue("pkg")
		zipname := strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".zip"
		if Windows {
			pkgpath = strings.Replace(pkgpath, "/", "\\", -1)
			zipit(pkgpath, zipname)
		} else {
			core.RunCmdB("zip -r " + zipname + " " + pkgpath + "/")
		}
		time.Sleep(500 * time.Millisecond)

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipname))
		http.ServeFile(w, r, strings.Replace(r.FormValue("pkg"), "/", ".", -1)+".zip")

		callmet = true
	} else if isURL := (r.URL.Path == "/api/complete" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		prefx := r.FormValue("pref")
		ret := []bson.M{}
		//return {name: ea.word, value: ea.insert, score: 0, meta: ea.meta}
		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
		score := 0
		for _, v := range gos.Variables {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "Global variable | " + v.Type})
			}

		}

		for _, v := range gos.RootImports {

			if strings.Contains(v.Src, prefx) {
				score = score + 1
				paths := strings.Split(v.Src, "/")
				ret = append(ret, bson.M{"name": v.Src, "value": paths[len(paths)-1], "score": score, "meta": "package"})
			}

		}

		for _, v := range gos.Header.Structs {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "Interface"})
			}

		}

		for _, v := range gos.Header.Objects {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "{{Interface func group}}"})
			}

		}

		for _, v := range gos.Methods.Methods {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name + " | ", "score": score, "meta": "{{Template pipeline}}"})
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": "Net" + v.Name + "(" + v.Variables + ")", "score": score, "meta": "Method"})
			}

		}

		for _, v := range gos.Templates.Templates {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name + " ", "score": score, "meta": "{{Template reference}}"})
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": "Netb" + v.Name + "(" + v.Struct + "{})", "score": score, "meta": "Template"})
			}

		}

		response = mResponse(ret)

		callmet = true
	} else if isURL := (r.URL.Path == "/api/console" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		if strings.Contains(r.FormValue("command"), "cd") {
			parts := strings.Fields(r.FormValue("command"))

			if len(parts) == 1 {
				os.Chdir("")
			} else {
				os.Chdir(parts[1])
			}
			if dir, err := os.Getwd(); err != nil {
				response = "Changed directory to " + dir
			}

		} else {
			data, err := core.RunCmdSmart(r.FormValue("command"))
			if err != nil {
				response = fmt.Sprintf("Error:: %s", err) + "" + data
			} else {
				response = data
			}
		}
		w.Write([]byte(response))

		response = ""

		callmet = true
	}

	if callmet {
		session.Save(r, w)
		session = nil
		if response != "" {
			//Unmarshal json
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
		}
		return
	}
	session = nil
	return
}

func DebugTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	lastline := 0
	linestring := ""
	defer func() {
		if n := recover(); n != nil {
			log.Println()
			// log.Println(n)
			log.Println("Error on line :", lastline+1, ":"+strings.TrimSpace(linestring))
			//http.Redirect(w,r,"",307)
		}
	}()

	p, err := loadPage(r.URL.Path)
	filename := tmpl + ".tmpl"
	body, err := Asset(filename)
	session, er := store.Get(r, "session-")

	if er != nil {
		session, er = store.New(r, "session-")
	}
	p.Session = session
	p.R = r
	if err != nil {
		log.Print(err)

	} else {

		lines := strings.Split(string(body), "\n")
		// log.Println( lines )
		linebuffer := ""
		waitend := false
		open := 0
		for i, line := range lines {

			processd := false

			if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				linebuffer += line
				waitend = true

				endstr := ""
				processd = true
				if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {

					open++

				}
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if waitend && !processd && !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end")) {
				linebuffer += line

				endstr := ""
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}

			}

			if !waitend && !processd {
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if !processd && (strings.Contains(line, "{{end") || strings.Contains(line, "{{ end")) {
				open--

				if open == 0 {
					waitend = false

				}
			}
		}

	}

}

func DebugTemplatePath(tmpl string, intrf interface{}) {
	lastline := 0
	linestring := ""
	defer func() {
		if n := recover(); n != nil {

			log.Println("Error on line :", lastline+1, ":"+strings.TrimSpace(linestring))
			log.Println(n)
			//http.Redirect(w,r,"",307)
		}
	}()

	filename := tmpl
	body, err := Asset(filename)

	if err != nil {
		log.Print(err)

	} else {

		lines := strings.Split(string(body), "\n")
		// log.Println( lines )
		linebuffer := ""
		waitend := false
		open := 0
		for i, line := range lines {

			processd := false

			if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				linebuffer += line
				waitend = true

				endstr := ""
				if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {

					open++

				}

				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec

				processd = true
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if waitend && !processd && !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {
				linebuffer += line

				endstr := ""
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}

			}

			if !waitend && !processd {
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(TemplateFuncStore)
				t, _ = t.Parse(string([]byte(fmt.Sprintf("%s%s", linebuffer))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if !processd && (strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {
				open--

				if open == 0 {
					waitend = false

				}
			}
		}

	}

}
func Handler(w http.ResponseWriter, r *http.Request) {
	var p *gosweb.Page
	p, err := loadPage(r.URL.Path)
	var session *sessions.Session
	var er error
	if session, er = store.Get(r, "session-"); er != nil {
		session, _ = store.New(r, "session-")
	}

	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusNotFound)

		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			//
			return
		}
		pag.R = r
		pag.Session = session
		if p != nil {
			p.Session = nil
			p.Body = nil
			p.R = nil
			p = nil
		}

		if pag.IsResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) //""
		}
		session = nil

		return
	}

	if !p.IsResource {
		w.Header().Set("Content-Type", "text/html")
		p.Session = session
		p.R = r
		renderTemplate(w, p) //fmt.Sprintf("web%s", r.URL.Path)
		session.Save(r, w)
		// log.Println(w)
	} else {
		w.Header().Set("Cache-Control", "public")
		if strings.Contains(r.URL.Path, ".css") {
			w.Header().Add("Content-Type", "text/css")
		} else if strings.Contains(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "application/javascript")
		} else {
			w.Header().Add("Content-Type", http.DetectContentType(p.Body))
		}

		w.Write(p.Body)
	}

	p.Session = nil
	p.Body = nil
	p.R = nil
	p = nil
	session = nil

	return
}

var WebCache = gosweb.NewCache()

func loadPage(title string) (*gosweb.Page, error) {

	if lPage, ok := WebCache.Get(title); ok {
		return &lPage, nil
	}

	var nPage = gosweb.Page{}
	if roottitle := (title == "/"); roottitle {
		webbase := "web/"
		fname := fmt.Sprintf("%s%s", webbase, "index.html")
		body, err := Asset(fname)
		if err != nil {
			fname = fmt.Sprintf("%s%s", webbase, "index.tmpl")
			body, err = Asset(fname)
			if err != nil {
				return nil, err
			}
			nPage.Body = body
			WebCache.Put(title, nPage)
			body = nil
			return &nPage, nil
		}
		nPage.Body = body
		nPage.IsResource = true
		WebCache.Put(title, nPage)
		body = nil
		return &nPage, nil

	}

	filename := fmt.Sprintf("web%s.tmpl", title)

	if body, err := Asset(filename); err != nil {
		filename = fmt.Sprintf("web%s.html", title)

		if body, err = Asset(filename); err != nil {
			filename = fmt.Sprintf("web%s", title)

			if body, err = Asset(filename); err != nil {
				return nil, err
			} else {
				if strings.Contains(title, ".tmpl") {
					return nil, nil
				}
				nPage.Body = body
				nPage.IsResource = true
				WebCache.Put(title, nPage)
				body = nil
				return &nPage, nil
			}
		} else {
			nPage.Body = body
			nPage.IsResource = true
			WebCache.Put(title, nPage)
			body = nil
			return &nPage, nil
		}
	} else {
		nPage.Body = body
		WebCache.Put(title, nPage)
		body = nil
		return &nPage, nil
	}

}

var Windows bool
var upgrader = websocket.Upgrader{}

func init() {

	gob.Register(&SoftUser{})

}

type FSCs struct {
	Path string
	Form Forms
}

func NetcastFSCs(args ...interface{}) *FSCs {

	s := FSCs{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructFSCs() *FSCs { return &FSCs{} }

type Dex struct {
	Misc string
	Text string
	Link string
}

func NetcastDex(args ...interface{}) *Dex {

	s := Dex{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructDex() *Dex { return &Dex{} }

type SoftUser struct {
	Username         string
	Email            string
	Password         []byte
	Apps             []App
	Docker           string
	TrialEnd         int64
	StripeID, FLogin string
}

func NetcastSoftUser(args ...interface{}) *SoftUser {

	s := SoftUser{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructSoftUser() *SoftUser { return &SoftUser{} }

type NoStruct struct {
}

func NetcastNoStruct(args ...interface{}) *NoStruct {

	s := NoStruct{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructNoStruct() *NoStruct { return &NoStruct{} }

type USettings struct {
	LastPaid string
	Email    string
	StripeID string
}

func NetcastUSettings(args ...interface{}) *USettings {

	s := USettings{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructUSettings() *USettings { return &USettings{} }

type App struct {
	Type             string
	Name             string
	PublicName       string
	Css              []string
	Groups           []string
	Passed, Running  bool
	LatestBuild, Pid string
}

func NetcastApp(args ...interface{}) *App {

	s := App{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructApp() *App { return &App{} }

type TemplateEdits struct {
	SavesTo, PKG, PreviewLink, ID, Mime string
	File                                []byte
	Settings                            rPut
}

func NetcastTemplateEdits(args ...interface{}) *TemplateEdits {

	s := TemplateEdits{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructTemplateEdits() *TemplateEdits { return &TemplateEdits{} }

type WebRootEdits struct {
	SavesTo, Type, PreviewLink, ID, PKG string
	File                                []byte
}

func NetcastWebRootEdits(args ...interface{}) *WebRootEdits {

	s := WebRootEdits{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructWebRootEdits() *WebRootEdits { return &WebRootEdits{} }

type TEditor struct {
	PKG, Type, LType string
	CreateForm       rPut
}

func NetcastTEditor(args ...interface{}) *TEditor {

	s := TEditor{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructTEditor() *TEditor { return &TEditor{} }

type Navbars struct {
	Mode string
	ID   string
}

func NetcastNavbars(args ...interface{}) *Navbars {

	s := Navbars{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructNavbars() *Navbars { return &Navbars{} }

type sModal struct {
	Title   string
	Body    string
	Color   string
	Buttons []sButton
	Form    Forms
}

func NetcastsModal(args ...interface{}) *sModal {

	s := sModal{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsModal() *sModal { return &sModal{} }

type Forms struct {
	Link    string
	Inputs  []Inputs
	Buttons []sButton
	CTA     string
	Class   string
}

func NetcastForms(args ...interface{}) *Forms {

	s := Forms{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructForms() *Forms { return &Forms{} }

type sButton struct {
	Text  string
	Class string
	Link  string
}

func NetcastsButton(args ...interface{}) *sButton {

	s := sButton{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsButton() *sButton { return &sButton{} }

type sTab struct {
	Buttons []sButton
}

func NetcastsTab(args ...interface{}) *sTab {

	s := sTab{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsTab() *sTab { return &sTab{} }

type DForm struct {
	Text, Link string
}

func NetcastDForm(args ...interface{}) *DForm {

	s := DForm{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructDForm() *DForm { return &DForm{} }

type Alertbs struct {
	Type     string
	Text     string
	Redirect string
}

func NetcastAlertbs(args ...interface{}) *Alertbs {

	s := Alertbs{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructAlertbs() *Alertbs { return &Alertbs{} }

type Inputs struct {
	Misc    string
	Text    string
	Name    string
	Type    string
	Options []string
	Value   string
}

func NetcastInputs(args ...interface{}) *Inputs {

	s := Inputs{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructInputs() *Inputs { return &Inputs{} }

type Aput struct {
	Link, Param, Value string
}

func NetcastAput(args ...interface{}) *Aput {

	s := Aput{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructAput() *Aput { return &Aput{} }

type rPut struct {
	Link     string
	DLink    string
	Inputs   []Inputs
	Count    string
	ListLink string
}

func NetcastrPut(args ...interface{}) *rPut {

	s := rPut{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructrPut() *rPut { return &rPut{} }

type sSWAL struct {
	Title, Type, Text string
}

func NetcastsSWAL(args ...interface{}) *sSWAL {

	s := sSWAL{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsSWAL() *sSWAL { return &sSWAL{} }

type sPackageEdit struct {
	Type, Mainf, Shutdown, Initf, Sessionf                  string
	IType, Package, Port, Key, Name, Ffpage, Erpage, Domain Aput
	Css                                                     rPut
	Imports                                                 []rPut
	Variables                                               []rPut
	CssFiles                                                []rPut
	CreateVar                                               rPut
	CreateImport                                            rPut
	TName                                                   string
}

func NetcastsPackageEdit(args ...interface{}) *sPackageEdit {

	s := sPackageEdit{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsPackageEdit() *sPackageEdit { return &sPackageEdit{} }

type DebugObj struct {
	PKG, Id, Username, RawLog, Time string
	Bugs                            []DebugNode
}

func NetcastDebugObj(args ...interface{}) *DebugObj {

	s := DebugObj{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructDebugObj() *DebugObj { return &DebugObj{} }

type DebugNode struct {
	Action, Line, CTA string
}

func NetcastDebugNode(args ...interface{}) *DebugNode {

	s := DebugNode{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructDebugNode() *DebugNode { return &DebugNode{} }

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

func NetcastPkgItem(args ...interface{}) *PkgItem {

	s := PkgItem{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructPkgItem() *PkgItem { return &PkgItem{} }

type sROC struct {
	Name      string
	CompLog   []byte
	Build     bool
	Time, Pid string
}

func NetcastsROC(args ...interface{}) *sROC {

	s := sROC{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructsROC() *sROC { return &sROC{} }

type vHuf struct {
	Type, PKG string
	Edata     []byte
}

func NetcastvHuf(args ...interface{}) *vHuf {

	s := vHuf{}
	mapp := args[0].(db.O)
	if _, ok := mapp["_id"]; ok {
		mapp["Id"] = mapp["_id"]
	}
	data, _ := json.Marshal(&mapp)

	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println(err.Error())
	}

	return &s
}
func NetstructvHuf() *vHuf { return &vHuf{} }

type myDemoObject Dex

//
func NetBindMisc(args ...interface{}) Dex {
	misc := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = misc.(string)
	return Nav

}

//
func NetListPlugins(args ...interface{}) []string {

	return getPlugins()

}

//
func NetBindID(args ...interface{}) Dex {
	id := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = id.(string)
	return Nav

}

//
func NetRandTen(args ...interface{}) string {

	return core.NewLen(10)

}

//
func NetFragmentize(args ...interface{}) (finall string) {
	inn := args[0]

	finall = strings.Replace(inn.(string), ".tmpl", "", -1)
	return

}

//
func NetparseLog(args ...interface{}) string {
	cline := args[0]

	calls := strings.Split(cline.(string), ":")
	actionText := ""
	if calls[0] == "service" {
		actionText = "The line is located in  Web service ( " + calls[1] + ") at line: " + calls[2]
	} else if calls[0] == "init" {
		actionText = "The line is located in your package Init func at line: " + calls[1]
	} else if calls[0] == "main" {
		actionText = "The line is located in your package Main func at line: " + calls[1]
	} else if calls[0] == "structs" {
		actionText = "The line is located in your package Interfaces at line: " + calls[1]
	} else if calls[0] == "meth" {
		actionText = "The line is located in your package template pipelines at line: " + calls[1]
	}
	return actionText

}

//
func NetanyBugs(args ...interface{}) (ajet bool) {
	packge := args[0]

	ajet = false //,err := DebugLogs.Find(bson.M{"pkg":packge.(string), "username":usernam.(string)}).Count()

	bugs := GetLogs(packge.(string))
	sapp := NetgetApp(getApps(), packge.(string))

	if len(bugs) > 0 {
		ajet = true
	}

	if sapp.Pid != "" {
		ajet = true
	}

	return

}

//
func NetPluginJS(args ...interface{}) string {

	plugins := getPlugins()
	jsstring := ""
	for _, v := range plugins {
		data, err := ioutil.ReadFile(os.ExpandEnv("$GOPATH") + "/src/" + v + "/index.js")
		if err != nil {
			fmt.Println("Error loading ", v)
		} else {
			jsstring = jsstring + "\n" + string(data)
		}
	}
	return "<script>" + jsstring + "</script>"

}

//
func NetFindmyBugs(args ...interface{}) (ajet []DebugObj) {
	packge := args[0]

	ajet = GetLogs(packge.(string))
	sapp := NetgetApp(getApps(), packge.(string))

	if sapp.Pid != "" {
		activLog := DebugObj{Time: "Server", Bugs: []DebugNode{}}
		ajet = append([]DebugObj{activLog}, ajet...)
	}
	return

}

//
func NetisExpired(args ...interface{}) bool {
	current := args[0]
	strip := args[1]

	if time.Now().Unix() > current.(int64) {
		if strip.(string) == "" {
			return true
		} else {
			return false
		}
	}

	return false

}

//
func NetgetTemplate(args ...interface{}) core.Template {
	templates := args[0]
	name := args[1]

	s := reflect.ValueOf(templates)
	slice := make([]App, s.Len())
	for i, _ := range slice {
		v := s.Index(i).Interface().(core.Template)

		if v.Name == name.(string) {
			return v
		}
	}

	return core.Template{}

}

//
func NetmConsole(args ...interface{}) string {

	return ""

}

//
func NetmPut(args ...interface{}) string {

	//response = "OK"

	return ""

}

//
func NetupdateApp(args ...interface{}) []App {
	apps := args[0]
	name := args[1]
	app := args[2]

	s := reflect.ValueOf(apps)
	n := make([]App, s.Len())
	slice := make([]App, s.Len())
	for i, _ := range slice {
		v := s.Index(i).Interface().(App)

		if v.Name == name.(string) {
			n = append(n, app.(App))
		} else if v.Name != "" {
			n = append(n, v)
		}
	}
	return n

}

//
func NetgetApp(args ...interface{}) App {
	apps := args[0]
	name := args[1]

	s := reflect.ValueOf(apps)
	slice := make([]App, s.Len())
	for i, _ := range slice {
		v := s.Index(i).Interface().(App)

		if v.Name == name.(string) {
			return v
		}
	}

	return App{}

}

func templateFNCss(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (css) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDCss = "tmpl/css.tmpl"

func NetCss(args ...interface{}) string {

	localid := templateIDCss
	var d *Dex
	defer templateFNCss(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Css")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bCss(d Dex) string {
	return NetbCss(d)
}

//
func NetbCss(d Dex) string {
	localid := templateIDCss
	defer templateFNCss(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Css")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcCss(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cCss(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcCss(args[0])
	} else {
		d = NetcCss()
	}
	return
}

func templateFNJS(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (js) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDJS = "tmpl/js.tmpl"

func NetJS(args ...interface{}) string {

	localid := templateIDJS
	var d *Dex
	defer templateFNJS(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("JS")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bJS(d Dex) string {
	return NetbJS(d)
}

//
func NetbJS(d Dex) string {
	localid := templateIDJS
	defer templateFNJS(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("JS")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcJS(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cJS(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcJS(args[0])
	} else {
		d = NetcJS()
	}
	return
}

func templateFNFA(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/fa) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDFA = "tmpl/ui/fa.tmpl"

func NetFA(args ...interface{}) string {

	localid := templateIDFA
	var d *Dex
	defer templateFNFA(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("FA")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bFA(d Dex) string {
	return NetbFA(d)
}

//
func NetbFA(d Dex) string {
	localid := templateIDFA
	defer templateFNFA(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("FA")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcFA(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cFA(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcFA(args[0])
	} else {
		d = NetcFA()
	}
	return
}

func templateFNPluginList(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/pluginlist) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPluginList = "tmpl/ui/pluginlist.tmpl"

func NetPluginList(args ...interface{}) string {

	localid := templateIDPluginList
	var d *gosweb.NoStruct
	defer templateFNPluginList(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &gosweb.NoStruct{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PluginList")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bPluginList(d gosweb.NoStruct) string {
	return NetbPluginList(d)
}

//
func NetbPluginList(d gosweb.NoStruct) string {
	localid := templateIDPluginList
	defer templateFNPluginList(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PluginList")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = gosweb.NoStruct{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcPluginList(args ...interface{}) (d gosweb.NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = gosweb.NoStruct{}
	}
	return
}

func cPluginList(args ...interface{}) (d gosweb.NoStruct) {
	if len(args) > 0 {
		d = NetcPluginList(args[0])
	} else {
		d = NetcPluginList()
	}
	return
}

func templateFNLogin(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/login) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDLogin = "tmpl/ui/login.tmpl"

func NetLogin(args ...interface{}) string {

	localid := templateIDLogin
	var d *Dex
	defer templateFNLogin(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Login")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bLogin(d Dex) string {
	return NetbLogin(d)
}

//
func NetbLogin(d Dex) string {
	localid := templateIDLogin
	defer templateFNLogin(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Login")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcLogin(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cLogin(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcLogin(args[0])
	} else {
		d = NetcLogin()
	}
	return
}

func templateFNModal(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/modal) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDModal = "tmpl/ui/modal.tmpl"

func NetModal(args ...interface{}) string {

	localid := templateIDModal
	var d *sModal
	defer templateFNModal(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sModal{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Modal")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bModal(d sModal) string {
	return NetbModal(d)
}

//
func NetbModal(d sModal) string {
	localid := templateIDModal
	defer templateFNModal(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Modal")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sModal{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcModal(args ...interface{}) (d sModal) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sModal{}
	}
	return
}

func cModal(args ...interface{}) (d sModal) {
	if len(args) > 0 {
		d = NetcModal(args[0])
	} else {
		d = NetcModal()
	}
	return
}

func templateFNxButton(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/sbutton) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDxButton = "tmpl/ui/sbutton.tmpl"

func NetxButton(args ...interface{}) string {

	localid := templateIDxButton
	var d *sButton
	defer templateFNxButton(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sButton{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("xButton")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bxButton(d sButton) string {
	return NetbxButton(d)
}

//
func NetbxButton(d sButton) string {
	localid := templateIDxButton
	defer templateFNxButton(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("xButton")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sButton{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcxButton(args ...interface{}) (d sButton) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sButton{}
	}
	return
}

func cxButton(args ...interface{}) (d sButton) {
	if len(args) > 0 {
		d = NetcxButton(args[0])
	} else {
		d = NetcxButton()
	}
	return
}

func templateFNjButton(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/jbutton) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDjButton = "tmpl/ui/user/forms/jbutton.tmpl"

func NetjButton(args ...interface{}) string {

	localid := templateIDjButton
	var d *sButton
	defer templateFNjButton(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sButton{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("jButton")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bjButton(d sButton) string {
	return NetbjButton(d)
}

//
func NetbjButton(d sButton) string {
	localid := templateIDjButton
	defer templateFNjButton(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("jButton")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sButton{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcjButton(args ...interface{}) (d sButton) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sButton{}
	}
	return
}

func cjButton(args ...interface{}) (d sButton) {
	if len(args) > 0 {
		d = NetcjButton(args[0])
	} else {
		d = NetcjButton()
	}
	return
}

func templateFNPUT(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/aput) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPUT = "tmpl/ui/user/forms/aput.tmpl"

func NetPUT(args ...interface{}) string {

	localid := templateIDPUT
	var d *Aput
	defer templateFNPUT(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Aput{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PUT")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bPUT(d Aput) string {
	return NetbPUT(d)
}

//
func NetbPUT(d Aput) string {
	localid := templateIDPUT
	defer templateFNPUT(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PUT")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Aput{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcPUT(args ...interface{}) (d Aput) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Aput{}
	}
	return
}

func cPUT(args ...interface{}) (d Aput) {
	if len(args) > 0 {
		d = NetcPUT(args[0])
	} else {
		d = NetcPUT()
	}
	return
}

func templateFNGroup(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/tab) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDGroup = "tmpl/ui/user/forms/tab.tmpl"

func NetGroup(args ...interface{}) string {

	localid := templateIDGroup
	var d *sTab
	defer templateFNGroup(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sTab{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Group")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bGroup(d sTab) string {
	return NetbGroup(d)
}

//
func NetbGroup(d sTab) string {
	localid := templateIDGroup
	defer templateFNGroup(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Group")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sTab{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcGroup(args ...interface{}) (d sTab) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sTab{}
	}
	return
}

func cGroup(args ...interface{}) (d sTab) {
	if len(args) > 0 {
		d = NetcGroup(args[0])
	} else {
		d = NetcGroup()
	}
	return
}

func templateFNRegister(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/register) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDRegister = "tmpl/ui/register.tmpl"

func NetRegister(args ...interface{}) string {

	localid := templateIDRegister
	var d *Dex
	defer templateFNRegister(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Register")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bRegister(d Dex) string {
	return NetbRegister(d)
}

//
func NetbRegister(d Dex) string {
	localid := templateIDRegister
	defer templateFNRegister(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Register")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcRegister(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cRegister(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcRegister(args[0])
	} else {
		d = NetcRegister()
	}
	return
}

func templateFNAlert(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/alert) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDAlert = "tmpl/ui/alert.tmpl"

func NetAlert(args ...interface{}) string {

	localid := templateIDAlert
	var d *Alertbs
	defer templateFNAlert(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Alertbs{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Alert")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bAlert(d Alertbs) string {
	return NetbAlert(d)
}

//
func NetbAlert(d Alertbs) string {
	localid := templateIDAlert
	defer templateFNAlert(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Alert")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Alertbs{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcAlert(args ...interface{}) (d Alertbs) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Alertbs{}
	}
	return
}

func cAlert(args ...interface{}) (d Alertbs) {
	if len(args) > 0 {
		d = NetcAlert(args[0])
	} else {
		d = NetcAlert()
	}
	return
}

func templateFNStructEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/structs) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDStructEditor = "tmpl/editor/structs.tmpl"

func NetStructEditor(args ...interface{}) string {

	localid := templateIDStructEditor
	var d *vHuf
	defer templateFNStructEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &vHuf{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("StructEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bStructEditor(d vHuf) string {
	return NetbStructEditor(d)
}

//
func NetbStructEditor(d vHuf) string {
	localid := templateIDStructEditor
	defer templateFNStructEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("StructEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = vHuf{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcStructEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = vHuf{}
	}
	return
}

func cStructEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		d = NetcStructEditor(args[0])
	} else {
		d = NetcStructEditor()
	}
	return
}

func templateFNMethodEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/methods) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDMethodEditor = "tmpl/editor/methods.tmpl"

func NetMethodEditor(args ...interface{}) string {

	localid := templateIDMethodEditor
	var d *vHuf
	defer templateFNMethodEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &vHuf{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("MethodEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bMethodEditor(d vHuf) string {
	return NetbMethodEditor(d)
}

//
func NetbMethodEditor(d vHuf) string {
	localid := templateIDMethodEditor
	defer templateFNMethodEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("MethodEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = vHuf{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcMethodEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = vHuf{}
	}
	return
}

func cMethodEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		d = NetcMethodEditor(args[0])
	} else {
		d = NetcMethodEditor()
	}
	return
}

func templateFNObjectEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/objects) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDObjectEditor = "tmpl/editor/objects.tmpl"

func NetObjectEditor(args ...interface{}) string {

	localid := templateIDObjectEditor
	var d *vHuf
	defer templateFNObjectEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &vHuf{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ObjectEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bObjectEditor(d vHuf) string {
	return NetbObjectEditor(d)
}

//
func NetbObjectEditor(d vHuf) string {
	localid := templateIDObjectEditor
	defer templateFNObjectEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ObjectEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = vHuf{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcObjectEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = vHuf{}
	}
	return
}

func cObjectEditor(args ...interface{}) (d vHuf) {
	if len(args) > 0 {
		d = NetcObjectEditor(args[0])
	} else {
		d = NetcObjectEditor()
	}
	return
}

func templateFNEndpointEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/endpoints) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDEndpointEditor = "tmpl/editor/endpoints.tmpl"

func NetEndpointEditor(args ...interface{}) string {

	localid := templateIDEndpointEditor
	var d *TEditor
	defer templateFNEndpointEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &TEditor{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bEndpointEditor(d TEditor) string {
	return NetbEndpointEditor(d)
}

//
func NetbEndpointEditor(d TEditor) string {
	localid := templateIDEndpointEditor
	defer templateFNEndpointEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = TEditor{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcEndpointEditor(args ...interface{}) (d TEditor) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = TEditor{}
	}
	return
}

func cEndpointEditor(args ...interface{}) (d TEditor) {
	if len(args) > 0 {
		d = NetcEndpointEditor(args[0])
	} else {
		d = NetcEndpointEditor()
	}
	return
}

func templateFNTimerEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/timers) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDTimerEditor = "tmpl/editor/timers.tmpl"

func NetTimerEditor(args ...interface{}) string {

	localid := templateIDTimerEditor
	var d *TEditor
	defer templateFNTimerEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &TEditor{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TimerEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bTimerEditor(d TEditor) string {
	return NetbTimerEditor(d)
}

//
func NetbTimerEditor(d TEditor) string {
	localid := templateIDTimerEditor
	defer templateFNTimerEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TimerEditor")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = TEditor{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcTimerEditor(args ...interface{}) (d TEditor) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = TEditor{}
	}
	return
}

func cTimerEditor(args ...interface{}) (d TEditor) {
	if len(args) > 0 {
		d = NetcTimerEditor(args[0])
	} else {
		d = NetcTimerEditor()
	}
	return
}

func templateFNFSC(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/fsc) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDFSC = "tmpl/ui/fsc.tmpl"

func NetFSC(args ...interface{}) string {

	localid := templateIDFSC
	var d *FSCs
	defer templateFNFSC(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &FSCs{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("FSC")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bFSC(d FSCs) string {
	return NetbFSC(d)
}

//
func NetbFSC(d FSCs) string {
	localid := templateIDFSC
	defer templateFNFSC(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("FSC")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = FSCs{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcFSC(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = FSCs{}
	}
	return
}

func cFSC(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		d = NetcFSC(args[0])
	} else {
		d = NetcFSC()
	}
	return
}

func templateFNMV(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/mv) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDMV = "tmpl/ui/mv.tmpl"

func NetMV(args ...interface{}) string {

	localid := templateIDMV
	var d *FSCs
	defer templateFNMV(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &FSCs{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("MV")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bMV(d FSCs) string {
	return NetbMV(d)
}

//
func NetbMV(d FSCs) string {
	localid := templateIDMV
	defer templateFNMV(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("MV")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = FSCs{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcMV(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = FSCs{}
	}
	return
}

func cMV(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		d = NetcMV(args[0])
	} else {
		d = NetcMV()
	}
	return
}

func templateFNRM(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/rm) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDRM = "tmpl/ui/user/rm.tmpl"

func NetRM(args ...interface{}) string {

	localid := templateIDRM
	var d *FSCs
	defer templateFNRM(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &FSCs{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("RM")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bRM(d FSCs) string {
	return NetbRM(d)
}

//
func NetbRM(d FSCs) string {
	localid := templateIDRM
	defer templateFNRM(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("RM")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = FSCs{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcRM(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = FSCs{}
	}
	return
}

func cRM(args ...interface{}) (d FSCs) {
	if len(args) > 0 {
		d = NetcRM(args[0])
	} else {
		d = NetcRM()
	}
	return
}

func templateFNWebRootEdit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/webrootedit) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDWebRootEdit = "tmpl/ui/user/panel/webrootedit.tmpl"

func NetWebRootEdit(args ...interface{}) string {

	localid := templateIDWebRootEdit
	var d *WebRootEdits
	defer templateFNWebRootEdit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &WebRootEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bWebRootEdit(d WebRootEdits) string {
	return NetbWebRootEdit(d)
}

//
func NetbWebRootEdit(d WebRootEdits) string {
	localid := templateIDWebRootEdit
	defer templateFNWebRootEdit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = WebRootEdits{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcWebRootEdit(args ...interface{}) (d WebRootEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = WebRootEdits{}
	}
	return
}

func cWebRootEdit(args ...interface{}) (d WebRootEdits) {
	if len(args) > 0 {
		d = NetcWebRootEdit(args[0])
	} else {
		d = NetcWebRootEdit()
	}
	return
}

func templateFNWebRootEdittwo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/webtwo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDWebRootEdittwo = "tmpl/ui/user/panel/webtwo.tmpl"

func NetWebRootEdittwo(args ...interface{}) string {

	localid := templateIDWebRootEdittwo
	var d *WebRootEdits
	defer templateFNWebRootEdittwo(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &WebRootEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdittwo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bWebRootEdittwo(d WebRootEdits) string {
	return NetbWebRootEdittwo(d)
}

//
func NetbWebRootEdittwo(d WebRootEdits) string {
	localid := templateIDWebRootEdittwo
	defer templateFNWebRootEdittwo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdittwo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = WebRootEdits{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcWebRootEdittwo(args ...interface{}) (d WebRootEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = WebRootEdits{}
	}
	return
}

func cWebRootEdittwo(args ...interface{}) (d WebRootEdits) {
	if len(args) > 0 {
		d = NetcWebRootEdittwo(args[0])
	} else {
		d = NetcWebRootEdittwo()
	}
	return
}

func templateFNuSettings(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/settings) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDuSettings = "tmpl/editor/settings.tmpl"

func NetuSettings(args ...interface{}) string {

	localid := templateIDuSettings
	var d *USettings
	defer templateFNuSettings(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &USettings{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("uSettings")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func buSettings(d USettings) string {
	return NetbuSettings(d)
}

//
func NetbuSettings(d USettings) string {
	localid := templateIDuSettings
	defer templateFNuSettings(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("uSettings")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = USettings{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcuSettings(args ...interface{}) (d USettings) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = USettings{}
	}
	return
}

func cuSettings(args ...interface{}) (d USettings) {
	if len(args) > 0 {
		d = NetcuSettings(args[0])
	} else {
		d = NetcuSettings()
	}
	return
}

func templateFNForm(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/form) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDForm = "tmpl/ui/user/forms/form.tmpl"

func NetForm(args ...interface{}) string {

	localid := templateIDForm
	var d *Forms
	defer templateFNForm(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Forms{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Form")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bForm(d Forms) string {
	return NetbForm(d)
}

//
func NetbForm(d Forms) string {
	localid := templateIDForm
	defer templateFNForm(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Form")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Forms{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcForm(args ...interface{}) (d Forms) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Forms{}
	}
	return
}

func cForm(args ...interface{}) (d Forms) {
	if len(args) > 0 {
		d = NetcForm(args[0])
	} else {
		d = NetcForm()
	}
	return
}

func templateFNSWAL(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/swal) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDSWAL = "tmpl/ui/user/forms/swal.tmpl"

func NetSWAL(args ...interface{}) string {

	localid := templateIDSWAL
	var d *sSWAL
	defer templateFNSWAL(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sSWAL{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("SWAL")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bSWAL(d sSWAL) string {
	return NetbSWAL(d)
}

//
func NetbSWAL(d sSWAL) string {
	localid := templateIDSWAL
	defer templateFNSWAL(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("SWAL")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sSWAL{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcSWAL(args ...interface{}) (d sSWAL) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sSWAL{}
	}
	return
}

func cSWAL(args ...interface{}) (d sSWAL) {
	if len(args) > 0 {
		d = NetcSWAL(args[0])
	} else {
		d = NetcSWAL()
	}
	return
}

func templateFNROC(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/roc) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDROC = "tmpl/ui/user/panel/roc.tmpl"

func NetROC(args ...interface{}) string {

	localid := templateIDROC
	var d *sROC
	defer templateFNROC(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sROC{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ROC")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bROC(d sROC) string {
	return NetbROC(d)
}

//
func NetbROC(d sROC) string {
	localid := templateIDROC
	defer templateFNROC(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ROC")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sROC{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcROC(args ...interface{}) (d sROC) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sROC{}
	}
	return
}

func cROC(args ...interface{}) (d sROC) {
	if len(args) > 0 {
		d = NetcROC(args[0])
	} else {
		d = NetcROC()
	}
	return
}

func templateFNRPUT(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/forms/rput) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDRPUT = "tmpl/ui/user/forms/rput.tmpl"

func NetRPUT(args ...interface{}) string {

	localid := templateIDRPUT
	var d *rPut
	defer templateFNRPUT(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &rPut{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("RPUT")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bRPUT(d rPut) string {
	return NetbRPUT(d)
}

//
func NetbRPUT(d rPut) string {
	localid := templateIDRPUT
	defer templateFNRPUT(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("RPUT")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = rPut{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcRPUT(args ...interface{}) (d rPut) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = rPut{}
	}
	return
}

func cRPUT(args ...interface{}) (d rPut) {
	if len(args) > 0 {
		d = NetcRPUT(args[0])
	} else {
		d = NetcRPUT()
	}
	return
}

func templateFNPackageEdit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/package) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPackageEdit = "tmpl/ui/user/panel/package.tmpl"

func NetPackageEdit(args ...interface{}) string {

	localid := templateIDPackageEdit
	var d *sPackageEdit
	defer templateFNPackageEdit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &sPackageEdit{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PackageEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bPackageEdit(d sPackageEdit) string {
	return NetbPackageEdit(d)
}

//
func NetbPackageEdit(d sPackageEdit) string {
	localid := templateIDPackageEdit
	defer templateFNPackageEdit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PackageEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = sPackageEdit{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcPackageEdit(args ...interface{}) (d sPackageEdit) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = sPackageEdit{}
	}
	return
}

func cPackageEdit(args ...interface{}) (d sPackageEdit) {
	if len(args) > 0 {
		d = NetcPackageEdit(args[0])
	} else {
		d = NetcPackageEdit()
	}
	return
}

func templateFNDelete(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/delete) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDDelete = "tmpl/ui/user/panel/delete.tmpl"

func NetDelete(args ...interface{}) string {

	localid := templateIDDelete
	var d *DForm
	defer templateFNDelete(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &DForm{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Delete")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bDelete(d DForm) string {
	return NetbDelete(d)
}

//
func NetbDelete(d DForm) string {
	localid := templateIDDelete
	defer templateFNDelete(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Delete")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = DForm{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcDelete(args ...interface{}) (d DForm) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = DForm{}
	}
	return
}

func cDelete(args ...interface{}) (d DForm) {
	if len(args) > 0 {
		d = NetcDelete(args[0])
	} else {
		d = NetcDelete()
	}
	return
}

func templateFNWelcome(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/welcome) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDWelcome = "tmpl/ui/welcome.tmpl"

func NetWelcome(args ...interface{}) string {

	localid := templateIDWelcome
	var d *Dex
	defer templateFNWelcome(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Welcome")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bWelcome(d Dex) string {
	return NetbWelcome(d)
}

//
func NetbWelcome(d Dex) string {
	localid := templateIDWelcome
	defer templateFNWelcome(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Welcome")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcWelcome(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cWelcome(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcWelcome(args[0])
	} else {
		d = NetcWelcome()
	}
	return
}

func templateFNStripe(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/stripe) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDStripe = "tmpl/ui/stripe.tmpl"

func NetStripe(args ...interface{}) string {

	localid := templateIDStripe
	var d *Dex
	defer templateFNStripe(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Stripe")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bStripe(d Dex) string {
	return NetbStripe(d)
}

//
func NetbStripe(d Dex) string {
	localid := templateIDStripe
	defer templateFNStripe(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Stripe")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcStripe(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cStripe(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcStripe(args[0])
	} else {
		d = NetcStripe()
	}
	return
}

func templateFNDebugger(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/debugger) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDDebugger = "tmpl/ui/debugger.tmpl"

func NetDebugger(args ...interface{}) string {

	localid := templateIDDebugger
	var d *DebugObj
	defer templateFNDebugger(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &DebugObj{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Debugger")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bDebugger(d DebugObj) string {
	return NetbDebugger(d)
}

//
func NetbDebugger(d DebugObj) string {
	localid := templateIDDebugger
	defer templateFNDebugger(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Debugger")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = DebugObj{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcDebugger(args ...interface{}) (d DebugObj) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = DebugObj{}
	}
	return
}

func cDebugger(args ...interface{}) (d DebugObj) {
	if len(args) > 0 {
		d = NetcDebugger(args[0])
	} else {
		d = NetcDebugger()
	}
	return
}

func templateFNTemplateEdit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/templateEditor) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDTemplateEdit = "tmpl/ui/user/panel/templateEditor.tmpl"

func NetTemplateEdit(args ...interface{}) string {

	localid := templateIDTemplateEdit
	var d *TemplateEdits
	defer templateFNTemplateEdit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &TemplateEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bTemplateEdit(d TemplateEdits) string {
	return NetbTemplateEdit(d)
}

//
func NetbTemplateEdit(d TemplateEdits) string {
	localid := templateIDTemplateEdit
	defer templateFNTemplateEdit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEdit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = TemplateEdits{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcTemplateEdit(args ...interface{}) (d TemplateEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = TemplateEdits{}
	}
	return
}

func cTemplateEdit(args ...interface{}) (d TemplateEdits) {
	if len(args) > 0 {
		d = NetcTemplateEdit(args[0])
	} else {
		d = NetcTemplateEdit()
	}
	return
}

func templateFNTemplateEditTwo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/tpetwo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDTemplateEditTwo = "tmpl/ui/user/panel/tpetwo.tmpl"

func NetTemplateEditTwo(args ...interface{}) string {

	localid := templateIDTemplateEditTwo
	var d *TemplateEdits
	defer templateFNTemplateEditTwo(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &TemplateEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEditTwo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bTemplateEditTwo(d TemplateEdits) string {
	return NetbTemplateEditTwo(d)
}

//
func NetbTemplateEditTwo(d TemplateEdits) string {
	localid := templateIDTemplateEditTwo
	defer templateFNTemplateEditTwo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEditTwo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = TemplateEdits{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcTemplateEditTwo(args ...interface{}) (d TemplateEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = TemplateEdits{}
	}
	return
}

func cTemplateEditTwo(args ...interface{}) (d TemplateEdits) {
	if len(args) > 0 {
		d = NetcTemplateEditTwo(args[0])
	} else {
		d = NetcTemplateEditTwo()
	}
	return
}

func templateFNInput(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/input) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDInput = "tmpl/ui/input.tmpl"

func NetInput(args ...interface{}) string {

	localid := templateIDInput
	var d *Inputs
	defer templateFNInput(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Inputs{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Input")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bInput(d Inputs) string {
	return NetbInput(d)
}

//
func NetbInput(d Inputs) string {
	localid := templateIDInput
	defer templateFNInput(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Input")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Inputs{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcInput(args ...interface{}) (d Inputs) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Inputs{}
	}
	return
}

func cInput(args ...interface{}) (d Inputs) {
	if len(args) > 0 {
		d = NetcInput(args[0])
	} else {
		d = NetcInput()
	}
	return
}

func templateFNDebuggerNode(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/debugnode) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDDebuggerNode = "tmpl/ui/debugnode.tmpl"

func NetDebuggerNode(args ...interface{}) string {

	localid := templateIDDebuggerNode
	var d *DebugObj
	defer templateFNDebuggerNode(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &DebugObj{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("DebuggerNode")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bDebuggerNode(d DebugObj) string {
	return NetbDebuggerNode(d)
}

//
func NetbDebuggerNode(d DebugObj) string {
	localid := templateIDDebuggerNode
	defer templateFNDebuggerNode(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("DebuggerNode")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = DebugObj{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcDebuggerNode(args ...interface{}) (d DebugObj) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = DebugObj{}
	}
	return
}

func cDebuggerNode(args ...interface{}) (d DebugObj) {
	if len(args) > 0 {
		d = NetcDebuggerNode(args[0])
	} else {
		d = NetcDebuggerNode()
	}
	return
}

func templateFNButton(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/button) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDButton = "tmpl/ui/button.tmpl"

func NetButton(args ...interface{}) string {

	localid := templateIDButton
	var d *Dex
	defer templateFNButton(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Button")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bButton(d Dex) string {
	return NetbButton(d)
}

//
func NetbButton(d Dex) string {
	localid := templateIDButton
	defer templateFNButton(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Button")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcButton(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cButton(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcButton(args[0])
	} else {
		d = NetcButton()
	}
	return
}

func templateFNSubmit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/submit) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDSubmit = "tmpl/ui/submit.tmpl"

func NetSubmit(args ...interface{}) string {

	localid := templateIDSubmit
	var d *Dex
	defer templateFNSubmit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Submit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bSubmit(d Dex) string {
	return NetbSubmit(d)
}

//
func NetbSubmit(d Dex) string {
	localid := templateIDSubmit
	defer templateFNSubmit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Submit")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcSubmit(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cSubmit(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcSubmit(args[0])
	} else {
		d = NetcSubmit()
	}
	return
}

func templateFNLogo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (logo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDLogo = "tmpl/logo.tmpl"

func NetLogo(args ...interface{}) string {

	localid := templateIDLogo
	var d *Dex
	defer templateFNLogo(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Logo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bLogo(d Dex) string {
	return NetbLogo(d)
}

//
func NetbLogo(d Dex) string {
	localid := templateIDLogo
	defer templateFNLogo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Logo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcLogo(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cLogo(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcLogo(args[0])
	} else {
		d = NetcLogo()
	}
	return
}

func templateFNNavbar(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/navbar) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDNavbar = "tmpl/ui/navbar.tmpl"

func NetNavbar(args ...interface{}) string {

	localid := templateIDNavbar
	var d *Dex
	defer templateFNNavbar(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Navbar")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bNavbar(d Dex) string {
	return NetbNavbar(d)
}

//
func NetbNavbar(d Dex) string {
	localid := templateIDNavbar
	defer templateFNNavbar(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Navbar")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcNavbar(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cNavbar(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcNavbar(args[0])
	} else {
		d = NetcNavbar()
	}
	return
}

func templateFNNavCustom(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/navbars) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDNavCustom = "tmpl/ui/navbars.tmpl"

func NetNavCustom(args ...interface{}) string {

	localid := templateIDNavCustom
	var d *Navbars
	defer templateFNNavCustom(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Navbars{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavCustom")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bNavCustom(d Navbars) string {
	return NetbNavCustom(d)
}

//
func NetbNavCustom(d Navbars) string {
	localid := templateIDNavCustom
	defer templateFNNavCustom(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavCustom")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Navbars{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcNavCustom(args ...interface{}) (d Navbars) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Navbars{}
	}
	return
}

func cNavCustom(args ...interface{}) (d Navbars) {
	if len(args) > 0 {
		d = NetcNavCustom(args[0])
	} else {
		d = NetcNavCustom()
	}
	return
}

func templateFNNavMain(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/navmain) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDNavMain = "tmpl/ui/navmain.tmpl"

func NetNavMain(args ...interface{}) string {

	localid := templateIDNavMain
	var d *Dex
	defer templateFNNavMain(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavMain")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bNavMain(d Dex) string {
	return NetbNavMain(d)
}

//
func NetbNavMain(d Dex) string {
	localid := templateIDNavMain
	defer templateFNNavMain(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavMain")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcNavMain(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cNavMain(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcNavMain(args[0])
	} else {
		d = NetcNavMain()
	}
	return
}

func templateFNNavPKG(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/navpkg) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDNavPKG = "tmpl/ui/navpkg.tmpl"

func NetNavPKG(args ...interface{}) string {

	localid := templateIDNavPKG
	var d *Dex
	defer templateFNNavPKG(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavPKG")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bNavPKG(d Dex) string {
	return NetbNavPKG(d)
}

//
func NetbNavPKG(d Dex) string {
	localid := templateIDNavPKG
	defer templateFNNavPKG(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavPKG")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcNavPKG(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cNavPKG(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcNavPKG(args[0])
	} else {
		d = NetcNavPKG()
	}
	return
}

func templateFNCrashedPage(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/crashedpage) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDCrashedPage = "tmpl/ui/crashedpage.tmpl"

func NetCrashedPage(args ...interface{}) string {

	localid := templateIDCrashedPage
	var d *Dex
	defer templateFNCrashedPage(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("CrashedPage")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bCrashedPage(d Dex) string {
	return NetbCrashedPage(d)
}

//
func NetbCrashedPage(d Dex) string {
	localid := templateIDCrashedPage
	defer templateFNCrashedPage(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("CrashedPage")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcCrashedPage(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cCrashedPage(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcCrashedPage(args[0])
	} else {
		d = NetcCrashedPage()
	}
	return
}

func templateFNEndpointTesting(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/endpointtester) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDEndpointTesting = "tmpl/ui/endpointtester.tmpl"

func NetEndpointTesting(args ...interface{}) string {

	localid := templateIDEndpointTesting
	var d *Dex
	defer templateFNEndpointTesting(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointTesting")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bEndpointTesting(d Dex) string {
	return NetbEndpointTesting(d)
}

//
func NetbEndpointTesting(d Dex) string {
	localid := templateIDEndpointTesting
	defer templateFNEndpointTesting(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointTesting")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcEndpointTesting(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cEndpointTesting(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcEndpointTesting(args[0])
	} else {
		d = NetcEndpointTesting()
	}
	return
}

func templateFNKanBan(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/kanban) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDKanBan = "tmpl/ui/kanban.tmpl"

func NetKanBan(args ...interface{}) string {

	localid := templateIDKanBan
	var d *Dex
	defer templateFNKanBan(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("KanBan")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bKanBan(d Dex) string {
	return NetbKanBan(d)
}

//
func NetbKanBan(d Dex) string {
	localid := templateIDKanBan
	defer templateFNKanBan(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("KanBan")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcKanBan(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cKanBan(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcKanBan(args[0])
	} else {
		d = NetcKanBan()
	}
	return
}

func templateFNNavPromo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/navpromo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDNavPromo = "tmpl/ui/navpromo.tmpl"

func NetNavPromo(args ...interface{}) string {

	localid := templateIDNavPromo
	var d *Dex
	defer templateFNNavPromo(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavPromo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bNavPromo(d Dex) string {
	return NetbNavPromo(d)
}

//
func NetbNavPromo(d Dex) string {
	localid := templateIDNavPromo
	defer templateFNNavPromo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("NavPromo")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcNavPromo(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = Dex{}
	}
	return
}

func cNavPromo(args ...interface{}) (d Dex) {
	if len(args) > 0 {
		d = NetcNavPromo(args[0])
	} else {
		d = NetcNavPromo()
	}
	return
}

func dummy_timer() {
	dg := time.Second * 5
	log.Println(dg)
}
func main() {
	fmt.Fprintf(os.Stdout, "%v\n", os.Getpid())

	dfd = os.ExpandEnv("$GOPATH")
	Windows = strings.Contains(runtime.GOOS, "windows")
	if dfd == "" {
		fmt.Println("Using temporary $GOPATH")
		if Windows {
			os.Chdir(os.ExpandEnv("$USERPROFILE"))
		} else {
			os.Chdir(os.ExpandEnv("$HOME"))
		}

		err := os.MkdirAll("workspace/", 0777)
		if err != nil {
			fmt.Println(err.Error())

		} else {
			//download go
			os.MkdirAll("workspace/src", 0777)
			os.MkdirAll("workspace/bin", 0777)
			cwd, _ := os.Getwd()
			cwd = cwd + "/workspace"
			os.Setenv("GOPATH", cwd)
			pathbin := os.ExpandEnv("$PATH")
			if Windows {
				os.Setenv("PATH", pathbin+":"+strings.Replace(cwd+"/bin", "/", "\\", -1))
			} else {
				os.Setenv("PATH", pathbin+":"+cwd+"/bin")
			}
			_, goinvs := core.RunCmdSmart("go help build")
			if goinvs != nil {
				fmt.Println("If you do not have GO, remember to download and install GO to complete installation. Find Go here : https://golang.org/dl/ . ***Installing GO requires sudo permission.")

			}
			dfd = cwd

		}
	}

	dir := os.ExpandEnv("$GOPATH") + "/src/github.com/cheikhshift/gos"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Downloading GoS")
		_, err := core.RunCmdSmart("go get github.com/cheikhshift/gos")
		if err != nil {
			color.Red("Please install GO : https://golang.org/dl/ ")
		} else {
			core.RunCmdSmart("gos deps")
		}
	}

	apps := getApps()
	newapps := []App{}
	for _, app := range apps {
		if app.Name != "" {
			app.Pid = ""
			newapps = append(newapps, app)
		}
	}
	saveApps(newapps)

	log.Println("Strukture up on port 8884")
	if len(os.Args) == 1 && !Windows {
		if isMac := strings.Contains(runtime.GOOS, "arwin"); isMac {
			core.RunCmd("open http://localhost:8884/index")
		} else {
			core.RunCmd("xdg-open http://localhost:8884/index")
		}
	} else if len(os.Args) == 1 && Windows {
		core.RunCmd("cmd /C start http://localhost:8884/index")
	}
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
		Domain:   "",
	}

	port := ":8884"
	if envport := os.ExpandEnv("$PORT"); envport != "" {
		port = fmt.Sprintf(":%s", envport)
	}
	log.Printf("Listenning on Port %v\n", port)

	//+++extendgxmlmain+++

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)
	http.Handle("/dist/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "web"}))
	http.HandleFunc("/", MakeHandler(Handler))

	h := &http.Server{Addr: port}

	go func() {
		errgos := h.ListenAndServe()
		if errgos != nil {
			log.Fatal(errgos)
		}
	}()

	<-stop

	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	h.Shutdown(ctx)

	log.Println("Server gracefully stopped")

}

//+++extendgxmlroot+++
