package main

import (
	"bufio"
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cheikhshift/db"
	"github.com/cheikhshift/gos/core"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/fatih/color"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

var store = sessions.NewCookieStore([]byte("something-secretive-is-what-a-gorrilla-needs"))

type NoStruct struct {
	/* emptystruct */
}

func NetsessionGet(key string, s *sessions.Session) string {
	return s.Values[key].(string)
}

func UrlAtZ(url, base string) (isURL bool) {
	isURL = strings.Index(url, base) == 0
	return
}

func NetsessionDelete(s *sessions.Session) string {
	//keys := make([]string, len(s.Values))

	//i := 0
	for k := range s.Values {
		// keys[i] = k.(string)
		NetsessionRemove(k.(string), s)
		//i++
	}

	return ""
}

func NetsessionRemove(key string, s *sessions.Session) string {
	delete(s.Values, key)
	return ""
}
func NetsessionKey(key string, s *sessions.Session) bool {
	if _, ok := s.Values[key]; ok {
		//do something here
		return true
	}

	return false
}

func Netadd(x, v float64) float64 {
	return v + x
}

func Netsubs(x, v float64) float64 {
	return v - x
}

func Netmultiply(x, v float64) float64 {
	return v * x
}

func Netdivided(x, v float64) float64 {
	return v / x
}

func NetsessionGetInt(key string, s *sessions.Session) interface{} {
	return s.Values[key]
}

func NetsessionSet(key string, value string, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}
func NetsessionSetInt(key string, value interface{}, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}

func dbDummy() {
	smap := db.O{}
	smap["key"] = "set"
	log.Println(smap)
}

func Netimportcss(s string) string {
	return fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\" /> ", s)
}

func Netimportjs(s string) string {
	return fmt.Sprintf("<script type=\"text/javascript\" src=\"%s\" ></script> ", s)
}

func formval(s string, r *http.Request) string {
	return r.FormValue(s)
}

func renderTemplate(w http.ResponseWriter, p *Page) bool {
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

			if pag.isResource {
				w.Write(pag.Body)
			} else {
				pag.R = p.R
				pag.Session = p.Session
				renderTemplate(w, pag) //"

			}
		}
	}()

	t := template.New("PageWrapper")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(p.Body))
	outp := new(bytes.Buffer)
	err := t.Execute(outp, p)
	if err != nil {
		log.Println(err.Error())
		DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			return false
		}
		pag.R = p.R
		pag.Session = p.Session
		p = nil
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) // ""

		}
		return false
	}

	p.Session.Save(p.R, w)

	fmt.Fprintf(w, html.UnescapeString(outp.String()))

	return true

}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string, *sessions.Session)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var session *sessions.Session
		var er error
		if session, er = store.Get(r, "session-"); er != nil {
			session, _ = store.New(r, "session-")
		}
		if attmpt := apiAttempt(w, r, session); !attmpt {
			fn(w, r, "", session)
		} else {
			context.Clear(r)
		}

	}
}

func mResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}
func apiAttempt(w http.ResponseWriter, r *http.Request, session *sessions.Session) (callmet bool) {
	var response string
	response = ""

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

		}
		callmet = true

	}

	if isURL := (r.URL.Path == "/api/pkg-bugs" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		bugs := GetLogs(r.FormValue("pkg"))
		sapp := NetgetApp(getApps(), r.FormValue("pkg"))
		if len(bugs) == 0 || sapp.Passed {
			response = "{}"
		} else {
			response = mResponse(bugs[0])
		}

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/empty" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		ClearLogs(r.FormValue("pkg"))
		response = NetbAlert(Alertbs{Type: "success", Text: "Your build logs are cleared."})

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/tester/" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")
		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		logfull, _ := core.RunCmdSmart("gos " + r.FormValue("mode") + " " + r.FormValue("c"))
		response = html.EscapeString(logfull)

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/create" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/delete" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/rename" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/new" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/act" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/put" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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

		}

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/build" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/start" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/stop" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/bin" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/export" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/complete" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

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
	}

	if isURL := (r.URL.Path == "/api/console" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

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
		if response != "" {
			//Unmarshal json
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
		}
		return
	}
	return
}
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate(body))
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate(body))
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate(body))
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
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
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer))))
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
func Handler(w http.ResponseWriter, r *http.Request, contxt string, session *sessions.Session) {
	var p *Page
	p, err := loadPage(r.URL.Path)

	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusNotFound)

		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			//context.Clear(r)
			return
		}
		pag.R = r
		pag.Session = session
		p = nil
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) //""
		}
		context.Clear(r)
		return
	}

	if !p.isResource {
		w.Header().Set("Content-Type", "text/html")
		p.Session = session
		p.R = r
		renderTemplate(w, p) //fmt.Sprintf("web%s", r.URL.Path)

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

	p.R = nil
	p.Session = nil
	p = nil
	context.Clear(r)
	return
}

func loadPage(title string) (*Page, error) {

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
			return &Page{Body: body, isResource: false}, nil
		}

		return &Page{Body: body, isResource: true}, nil

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
				return &Page{Body: body, isResource: true}, nil
			}
		} else {
			return &Page{Body: body, isResource: true}, nil
		}
	} else {
		return &Page{Body: body, isResource: false}, nil
	}

}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
func equalz(args ...interface{}) bool {
	if args[0] == args[1] {
		return true
	}
	return false
}
func nequalz(args ...interface{}) bool {
	if args[0] != args[1] {
		return true
	}
	return false
}

func netlt(x, v float64) bool {
	if x < v {
		return true
	}
	return false
}
func netgt(x, v float64) bool {
	if x > v {
		return true
	}
	return false
}
func netlte(x, v float64) bool {
	if x <= v {
		return true
	}
	return false
}

func GetLine(fname string, match string) int {
	intx := 0
	file, err := os.Open(fname)
	if err != nil {
		color.Red("Could not find a source file")
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intx = intx + 1
		if strings.Contains(scanner.Text(), match) {

			return intx
		}

	}

	return -1
}
func netgte(x, v float64) bool {
	if x >= v {
		return true
	}
	return false
}

type Page struct {
	Title      string
	Body       []byte
	request    *http.Request
	isResource bool
	R          *http.Request
	Session    *sessions.Session
}

func ReadyTemplate(body []byte) string {
	return strings.Replace(strings.Replace(strings.Replace(string(body), "/{", "\"{", -1), "}/", "}\"", -1), "`", "\"", -1)
}

var Windows bool

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
	Type, Mainf, Initf, Sessionf                            string
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

func NetBindMisc(args ...interface{}) Dex {
	misc := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = misc.(string)
	return Nav

}
func NetListPlugins(args ...interface{}) []string {

	return getPlugins()

}
func NetBindID(args ...interface{}) Dex {
	id := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = id.(string)
	return Nav

}
func NetRandTen(args ...interface{}) string {

	return core.NewLen(10)

}
func NetFragmentize(args ...interface{}) (finall string) {
	inn := args[0]

	finall = strings.Replace(inn.(string), ".tmpl", "", -1)
	return

}
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
func NetmConsole(args ...interface{}) string {

	return ""

}
func NetmPut(args ...interface{}) string {

	//response = "OK"

	return ""

}
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

func NetCss(args ...interface{}) string {

	var d Dex
	filename := "tmpl/css.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Css) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Css")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bCss(d Dex) string {
	return NetbCss(d)
}

func NetbCss(d Dex) string {

	filename := "tmpl/css.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Css")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Css) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BCss(intstr interface{}) string {
	return NetCss(intstr)
}

func NetJS(args ...interface{}) string {

	var d Dex
	filename := "tmpl/js.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (JS) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("JS")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bJS(d Dex) string {
	return NetbJS(d)
}

func NetbJS(d Dex) string {

	filename := "tmpl/js.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("JS")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (JS) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BJS(intstr interface{}) string {
	return NetJS(intstr)
}

func NetFA(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/fa.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (FA) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FA")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bFA(d Dex) string {
	return NetbFA(d)
}

func NetbFA(d Dex) string {

	filename := "tmpl/ui/fa.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FA")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (FA) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BFA(intstr interface{}) string {
	return NetFA(intstr)
}

func NetPluginList(args ...interface{}) string {

	var d NoStruct
	filename := "tmpl/ui/pluginlist.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PluginList) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = NoStruct{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PluginList")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPluginList(d NoStruct) string {
	return NetbPluginList(d)
}

func NetbPluginList(d NoStruct) string {

	filename := "tmpl/ui/pluginlist.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PluginList")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PluginList) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
}
func NetcPluginList(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func cPluginList(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func BPluginList(intstr interface{}) string {
	return NetPluginList(intstr)
}

func NetLogin(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/login.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Login) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Login")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bLogin(d Dex) string {
	return NetbLogin(d)
}

func NetbLogin(d Dex) string {

	filename := "tmpl/ui/login.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Login")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Login) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BLogin(intstr interface{}) string {
	return NetLogin(intstr)
}

func NetModal(args ...interface{}) string {

	var d sModal
	filename := "tmpl/ui/modal.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Modal) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sModal{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Modal")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bModal(d sModal) string {
	return NetbModal(d)
}

func NetbModal(d sModal) string {

	filename := "tmpl/ui/modal.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Modal")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Modal) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BModal(intstr interface{}) string {
	return NetModal(intstr)
}

func NetxButton(args ...interface{}) string {

	var d sButton
	filename := "tmpl/ui/sbutton.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (xButton) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sButton{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("xButton")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bxButton(d sButton) string {
	return NetbxButton(d)
}

func NetbxButton(d sButton) string {

	filename := "tmpl/ui/sbutton.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("xButton")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (xButton) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BxButton(intstr interface{}) string {
	return NetxButton(intstr)
}

func NetjButton(args ...interface{}) string {

	var d sButton
	filename := "tmpl/ui/user/forms/jbutton.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (jButton) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sButton{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("jButton")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bjButton(d sButton) string {
	return NetbjButton(d)
}

func NetbjButton(d sButton) string {

	filename := "tmpl/ui/user/forms/jbutton.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("jButton")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (jButton) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BjButton(intstr interface{}) string {
	return NetjButton(intstr)
}

func NetPUT(args ...interface{}) string {

	var d Aput
	filename := "tmpl/ui/user/forms/aput.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PUT) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Aput{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PUT")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPUT(d Aput) string {
	return NetbPUT(d)
}

func NetbPUT(d Aput) string {

	filename := "tmpl/ui/user/forms/aput.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PUT")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PUT) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BPUT(intstr interface{}) string {
	return NetPUT(intstr)
}

func NetGroup(args ...interface{}) string {

	var d sTab
	filename := "tmpl/ui/user/forms/tab.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Group) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sTab{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Group")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bGroup(d sTab) string {
	return NetbGroup(d)
}

func NetbGroup(d sTab) string {

	filename := "tmpl/ui/user/forms/tab.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Group")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Group) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BGroup(intstr interface{}) string {
	return NetGroup(intstr)
}

func NetRegister(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/register.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Register) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Register")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRegister(d Dex) string {
	return NetbRegister(d)
}

func NetbRegister(d Dex) string {

	filename := "tmpl/ui/register.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Register")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Register) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BRegister(intstr interface{}) string {
	return NetRegister(intstr)
}

func NetAlert(args ...interface{}) string {

	var d Alertbs
	filename := "tmpl/ui/alert.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Alert) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Alertbs{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Alert")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bAlert(d Alertbs) string {
	return NetbAlert(d)
}

func NetbAlert(d Alertbs) string {

	filename := "tmpl/ui/alert.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Alert")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Alert) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BAlert(intstr interface{}) string {
	return NetAlert(intstr)
}

func NetStructEditor(args ...interface{}) string {

	var d vHuf
	filename := "tmpl/editor/structs.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (StructEditor) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = vHuf{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("StructEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bStructEditor(d vHuf) string {
	return NetbStructEditor(d)
}

func NetbStructEditor(d vHuf) string {

	filename := "tmpl/editor/structs.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("StructEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (StructEditor) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BStructEditor(intstr interface{}) string {
	return NetStructEditor(intstr)
}

func NetMethodEditor(args ...interface{}) string {

	var d vHuf
	filename := "tmpl/editor/methods.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (MethodEditor) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = vHuf{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MethodEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bMethodEditor(d vHuf) string {
	return NetbMethodEditor(d)
}

func NetbMethodEditor(d vHuf) string {

	filename := "tmpl/editor/methods.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MethodEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (MethodEditor) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BMethodEditor(intstr interface{}) string {
	return NetMethodEditor(intstr)
}

func NetObjectEditor(args ...interface{}) string {

	var d vHuf
	filename := "tmpl/editor/objects.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ObjectEditor) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = vHuf{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ObjectEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bObjectEditor(d vHuf) string {
	return NetbObjectEditor(d)
}

func NetbObjectEditor(d vHuf) string {

	filename := "tmpl/editor/objects.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ObjectEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ObjectEditor) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BObjectEditor(intstr interface{}) string {
	return NetObjectEditor(intstr)
}

func NetEndpointEditor(args ...interface{}) string {

	var d TEditor
	filename := "tmpl/editor/endpoints.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (EndpointEditor) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = TEditor{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bEndpointEditor(d TEditor) string {
	return NetbEndpointEditor(d)
}

func NetbEndpointEditor(d TEditor) string {

	filename := "tmpl/editor/endpoints.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (EndpointEditor) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BEndpointEditor(intstr interface{}) string {
	return NetEndpointEditor(intstr)
}

func NetTimerEditor(args ...interface{}) string {

	var d TEditor
	filename := "tmpl/editor/timers.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TimerEditor) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = TEditor{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TimerEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTimerEditor(d TEditor) string {
	return NetbTimerEditor(d)
}

func NetbTimerEditor(d TEditor) string {

	filename := "tmpl/editor/timers.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TimerEditor")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TimerEditor) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BTimerEditor(intstr interface{}) string {
	return NetTimerEditor(intstr)
}

func NetFSC(args ...interface{}) string {

	var d FSCs
	filename := "tmpl/ui/fsc.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (FSC) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = FSCs{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FSC")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bFSC(d FSCs) string {
	return NetbFSC(d)
}

func NetbFSC(d FSCs) string {

	filename := "tmpl/ui/fsc.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FSC")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (FSC) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BFSC(intstr interface{}) string {
	return NetFSC(intstr)
}

func NetMV(args ...interface{}) string {

	var d FSCs
	filename := "tmpl/ui/mv.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (MV) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = FSCs{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MV")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bMV(d FSCs) string {
	return NetbMV(d)
}

func NetbMV(d FSCs) string {

	filename := "tmpl/ui/mv.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MV")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (MV) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BMV(intstr interface{}) string {
	return NetMV(intstr)
}

func NetRM(args ...interface{}) string {

	var d FSCs
	filename := "tmpl/ui/user/rm.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (RM) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = FSCs{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RM")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRM(d FSCs) string {
	return NetbRM(d)
}

func NetbRM(d FSCs) string {

	filename := "tmpl/ui/user/rm.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RM")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (RM) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BRM(intstr interface{}) string {
	return NetRM(intstr)
}

func NetWebRootEdit(args ...interface{}) string {

	var d WebRootEdits
	filename := "tmpl/ui/user/panel/webrootedit.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (WebRootEdit) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = WebRootEdits{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWebRootEdit(d WebRootEdits) string {
	return NetbWebRootEdit(d)
}

func NetbWebRootEdit(d WebRootEdits) string {

	filename := "tmpl/ui/user/panel/webrootedit.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (WebRootEdit) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BWebRootEdit(intstr interface{}) string {
	return NetWebRootEdit(intstr)
}

func NetWebRootEdittwo(args ...interface{}) string {

	var d WebRootEdits
	filename := "tmpl/ui/user/panel/webtwo.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (WebRootEdittwo) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = WebRootEdits{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdittwo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWebRootEdittwo(d WebRootEdits) string {
	return NetbWebRootEdittwo(d)
}

func NetbWebRootEdittwo(d WebRootEdits) string {

	filename := "tmpl/ui/user/panel/webtwo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdittwo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (WebRootEdittwo) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BWebRootEdittwo(intstr interface{}) string {
	return NetWebRootEdittwo(intstr)
}

func NetuSettings(args ...interface{}) string {

	var d USettings
	filename := "tmpl/editor/settings.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (uSettings) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = USettings{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("uSettings")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func buSettings(d USettings) string {
	return NetbuSettings(d)
}

func NetbuSettings(d USettings) string {

	filename := "tmpl/editor/settings.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("uSettings")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (uSettings) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BuSettings(intstr interface{}) string {
	return NetuSettings(intstr)
}

func NetForm(args ...interface{}) string {

	var d Forms
	filename := "tmpl/ui/user/forms/form.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Form) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Forms{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Form")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bForm(d Forms) string {
	return NetbForm(d)
}

func NetbForm(d Forms) string {

	filename := "tmpl/ui/user/forms/form.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Form")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Form) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BForm(intstr interface{}) string {
	return NetForm(intstr)
}

func NetSWAL(args ...interface{}) string {

	var d sSWAL
	filename := "tmpl/ui/user/forms/swal.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (SWAL) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sSWAL{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("SWAL")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bSWAL(d sSWAL) string {
	return NetbSWAL(d)
}

func NetbSWAL(d sSWAL) string {

	filename := "tmpl/ui/user/forms/swal.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("SWAL")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (SWAL) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BSWAL(intstr interface{}) string {
	return NetSWAL(intstr)
}

func NetROC(args ...interface{}) string {

	var d sROC
	filename := "tmpl/ui/user/panel/roc.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ROC) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sROC{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ROC")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bROC(d sROC) string {
	return NetbROC(d)
}

func NetbROC(d sROC) string {

	filename := "tmpl/ui/user/panel/roc.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ROC")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ROC) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BROC(intstr interface{}) string {
	return NetROC(intstr)
}

func NetRPUT(args ...interface{}) string {

	var d rPut
	filename := "tmpl/ui/user/forms/rput.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (RPUT) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = rPut{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RPUT")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRPUT(d rPut) string {
	return NetbRPUT(d)
}

func NetbRPUT(d rPut) string {

	filename := "tmpl/ui/user/forms/rput.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RPUT")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (RPUT) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BRPUT(intstr interface{}) string {
	return NetRPUT(intstr)
}

func NetPackageEdit(args ...interface{}) string {

	var d sPackageEdit
	filename := "tmpl/ui/user/panel/package.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PackageEdit) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = sPackageEdit{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PackageEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPackageEdit(d sPackageEdit) string {
	return NetbPackageEdit(d)
}

func NetbPackageEdit(d sPackageEdit) string {

	filename := "tmpl/ui/user/panel/package.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PackageEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (PackageEdit) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BPackageEdit(intstr interface{}) string {
	return NetPackageEdit(intstr)
}

func NetDelete(args ...interface{}) string {

	var d DForm
	filename := "tmpl/ui/user/panel/delete.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Delete) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = DForm{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Delete")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDelete(d DForm) string {
	return NetbDelete(d)
}

func NetbDelete(d DForm) string {

	filename := "tmpl/ui/user/panel/delete.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Delete")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Delete) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BDelete(intstr interface{}) string {
	return NetDelete(intstr)
}

func NetWelcome(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/welcome.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Welcome) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Welcome")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWelcome(d Dex) string {
	return NetbWelcome(d)
}

func NetbWelcome(d Dex) string {

	filename := "tmpl/ui/welcome.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Welcome")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Welcome) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BWelcome(intstr interface{}) string {
	return NetWelcome(intstr)
}

func NetStripe(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/stripe.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Stripe) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Stripe")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bStripe(d Dex) string {
	return NetbStripe(d)
}

func NetbStripe(d Dex) string {

	filename := "tmpl/ui/stripe.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Stripe")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Stripe) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BStripe(intstr interface{}) string {
	return NetStripe(intstr)
}

func NetDebugger(args ...interface{}) string {

	var d DebugObj
	filename := "tmpl/ui/debugger.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Debugger) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = DebugObj{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Debugger")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDebugger(d DebugObj) string {
	return NetbDebugger(d)
}

func NetbDebugger(d DebugObj) string {

	filename := "tmpl/ui/debugger.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Debugger")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Debugger) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BDebugger(intstr interface{}) string {
	return NetDebugger(intstr)
}

func NetTemplateEdit(args ...interface{}) string {

	var d TemplateEdits
	filename := "tmpl/ui/user/panel/templateEditor.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TemplateEdit) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = TemplateEdits{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTemplateEdit(d TemplateEdits) string {
	return NetbTemplateEdit(d)
}

func NetbTemplateEdit(d TemplateEdits) string {

	filename := "tmpl/ui/user/panel/templateEditor.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEdit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TemplateEdit) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BTemplateEdit(intstr interface{}) string {
	return NetTemplateEdit(intstr)
}

func NetTemplateEditTwo(args ...interface{}) string {

	var d TemplateEdits
	filename := "tmpl/ui/user/panel/tpetwo.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TemplateEditTwo) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = TemplateEdits{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEditTwo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTemplateEditTwo(d TemplateEdits) string {
	return NetbTemplateEditTwo(d)
}

func NetbTemplateEditTwo(d TemplateEdits) string {

	filename := "tmpl/ui/user/panel/tpetwo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEditTwo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (TemplateEditTwo) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BTemplateEditTwo(intstr interface{}) string {
	return NetTemplateEditTwo(intstr)
}

func NetInput(args ...interface{}) string {

	var d Inputs
	filename := "tmpl/ui/input.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Input) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Inputs{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Input")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bInput(d Inputs) string {
	return NetbInput(d)
}

func NetbInput(d Inputs) string {

	filename := "tmpl/ui/input.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Input")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Input) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BInput(intstr interface{}) string {
	return NetInput(intstr)
}

func NetDebuggerNode(args ...interface{}) string {

	var d DebugObj
	filename := "tmpl/ui/debugnode.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (DebuggerNode) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = DebugObj{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("DebuggerNode")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDebuggerNode(d DebugObj) string {
	return NetbDebuggerNode(d)
}

func NetbDebuggerNode(d DebugObj) string {

	filename := "tmpl/ui/debugnode.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("DebuggerNode")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (DebuggerNode) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BDebuggerNode(intstr interface{}) string {
	return NetDebuggerNode(intstr)
}

func NetButton(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/button.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Button) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Button")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bButton(d Dex) string {
	return NetbButton(d)
}

func NetbButton(d Dex) string {

	filename := "tmpl/ui/button.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Button")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Button) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BButton(intstr interface{}) string {
	return NetButton(intstr)
}

func NetSubmit(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/submit.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Submit) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Submit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bSubmit(d Dex) string {
	return NetbSubmit(d)
}

func NetbSubmit(d Dex) string {

	filename := "tmpl/ui/submit.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Submit")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Submit) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BSubmit(intstr interface{}) string {
	return NetSubmit(intstr)
}

func NetLogo(args ...interface{}) string {

	var d Dex
	filename := "tmpl/logo.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Logo) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Logo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bLogo(d Dex) string {
	return NetbLogo(d)
}

func NetbLogo(d Dex) string {

	filename := "tmpl/logo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Logo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Logo) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BLogo(intstr interface{}) string {
	return NetLogo(intstr)
}

func NetNavbar(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/navbar.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Navbar) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Navbar")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavbar(d Dex) string {
	return NetbNavbar(d)
}

func NetbNavbar(d Dex) string {

	filename := "tmpl/ui/navbar.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Navbar")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (Navbar) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BNavbar(intstr interface{}) string {
	return NetNavbar(intstr)
}

func NetNavCustom(args ...interface{}) string {

	var d Navbars
	filename := "tmpl/ui/navbars.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavCustom) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Navbars{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavCustom")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavCustom(d Navbars) string {
	return NetbNavCustom(d)
}

func NetbNavCustom(d Navbars) string {

	filename := "tmpl/ui/navbars.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavCustom")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavCustom) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BNavCustom(intstr interface{}) string {
	return NetNavCustom(intstr)
}

func NetNavMain(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/navmain.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavMain) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavMain")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavMain(d Dex) string {
	return NetbNavMain(d)
}

func NetbNavMain(d Dex) string {

	filename := "tmpl/ui/navmain.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavMain")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavMain) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BNavMain(intstr interface{}) string {
	return NetNavMain(intstr)
}

func NetNavPKG(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/navpkg.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavPKG) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPKG")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavPKG(d Dex) string {
	return NetbNavPKG(d)
}

func NetbNavPKG(d Dex) string {

	filename := "tmpl/ui/navpkg.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPKG")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavPKG) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BNavPKG(intstr interface{}) string {
	return NetNavPKG(intstr)
}

func NetCrashedPage(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/crashedpage.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (CrashedPage) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("CrashedPage")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bCrashedPage(d Dex) string {
	return NetbCrashedPage(d)
}

func NetbCrashedPage(d Dex) string {

	filename := "tmpl/ui/crashedpage.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("CrashedPage")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (CrashedPage) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BCrashedPage(intstr interface{}) string {
	return NetCrashedPage(intstr)
}

func NetEndpointTesting(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/endpointtester.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (EndpointTesting) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointTesting")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bEndpointTesting(d Dex) string {
	return NetbEndpointTesting(d)
}

func NetbEndpointTesting(d Dex) string {

	filename := "tmpl/ui/endpointtester.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointTesting")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (EndpointTesting) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BEndpointTesting(intstr interface{}) string {
	return NetEndpointTesting(intstr)
}

func NetNavPromo(args ...interface{}) string {

	var d Dex
	filename := "tmpl/ui/navpromo.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavPromo) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return ""
		}
	} else {
		d = Dex{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPromo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavPromo(d Dex) string {
	return NetbNavPromo(d)
}

func NetbNavPromo(d Dex) string {

	filename := "tmpl/ui/navpromo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPromo")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": NetBindMisc, "ListPlugins": NetListPlugins, "BindID": NetBindID, "RandTen": NetRandTen, "Fragmentize": NetFragmentize, "parseLog": NetparseLog, "anyBugs": NetanyBugs, "PluginJS": NetPluginJS, "FindmyBugs": NetFindmyBugs, "isExpired": NetisExpired, "getTemplate": NetgetTemplate, "mConsole": NetmConsole, "mPut": NetmPut, "updateApp": NetupdateApp, "getApp": NetgetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": NetstructFSCs, "isFSCs": NetcastFSCs, "Dex": NetstructDex, "isDex": NetcastDex, "SoftUser": NetstructSoftUser, "isSoftUser": NetcastSoftUser, "USettings": NetstructUSettings, "isUSettings": NetcastUSettings, "App": NetstructApp, "isApp": NetcastApp, "TemplateEdits": NetstructTemplateEdits, "isTemplateEdits": NetcastTemplateEdits, "WebRootEdits": NetstructWebRootEdits, "isWebRootEdits": NetcastWebRootEdits, "TEditor": NetstructTEditor, "isTEditor": NetcastTEditor, "Navbars": NetstructNavbars, "isNavbars": NetcastNavbars, "sModal": NetstructsModal, "issModal": NetcastsModal, "Forms": NetstructForms, "isForms": NetcastForms, "sButton": NetstructsButton, "issButton": NetcastsButton, "sTab": NetstructsTab, "issTab": NetcastsTab, "DForm": NetstructDForm, "isDForm": NetcastDForm, "Alertbs": NetstructAlertbs, "isAlertbs": NetcastAlertbs, "Inputs": NetstructInputs, "isInputs": NetcastInputs, "Aput": NetstructAput, "isAput": NetcastAput, "rPut": NetstructrPut, "isrPut": NetcastrPut, "sSWAL": NetstructsSWAL, "issSWAL": NetcastsSWAL, "sPackageEdit": NetstructsPackageEdit, "issPackageEdit": NetcastsPackageEdit, "DebugObj": NetstructDebugObj, "isDebugObj": NetcastDebugObj, "DebugNode": NetstructDebugNode, "isDebugNode": NetcastDebugNode, "PkgItem": NetstructPkgItem, "isPkgItem": NetcastPkgItem, "sROC": NetstructsROC, "issROC": NetcastsROC, "vHuf": NetstructvHuf, "isvHuf": NetcastvHuf})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (NavPromo) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
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

func BNavPromo(intstr interface{}) string {
	return NetNavPromo(intstr)
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
	http.HandleFunc("/", MakeHandler(Handler))

	http.Handle("/dist/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "web"}))

	errgos := http.ListenAndServe(port, nil)
	if errgos != nil {
		log.Fatal(errgos)
	}

}
