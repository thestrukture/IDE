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

func net_sessionGet(key string, s *sessions.Session) string {
	return s.Values[key].(string)
}

func UrlAtZ(url, base string) (isURL bool) {
	isURL = strings.Index(url, base) == 0
	return
}

func net_sessionDelete(s *sessions.Session) string {
	//keys := make([]string, len(s.Values))

	//i := 0
	for k := range s.Values {
		// keys[i] = k.(string)
		net_sessionRemove(k.(string), s)
		//i++
	}

	return ""
}

func net_sessionRemove(key string, s *sessions.Session) string {
	delete(s.Values, key)
	return ""
}
func net_sessionKey(key string, s *sessions.Session) bool {
	if _, ok := s.Values[key]; ok {
		//do something here
		return true
	}

	return false
}

func net_add(x, v float64) float64 {
	return v + x
}

func net_subs(x, v float64) float64 {
	return v - x
}

func net_multiply(x, v float64) float64 {
	return v * x
}

func net_divided(x, v float64) float64 {
	return v / x
}

func net_sessionGetInt(key string, s *sessions.Session) interface{} {
	return s.Values[key]
}

func net_sessionSet(key string, value string, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}
func net_sessionSetInt(key string, value interface{}, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}

func dbDummy() {
	smap := db.O{}
	smap["key"] = "set"
	log.Println(smap)
}

func net_importcss(s string) string {
	return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
}

func net_importjs(s string) string {
	return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
}

func formval(s string, r *http.Request) string {
	return r.FormValue(s)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page, session *sessions.Session) bool {
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path : web%s.tmpl reason : %s", r.URL.Path, n))

			DebugTemplate(w, r, fmt.Sprintf("web%s", r.URL.Path))
			w.WriteHeader(http.StatusInternalServerError)

			pag, err := loadPage("")
			if err != nil {
				log.Println(err.Error())
				return
			}
			if pag.isResource {
				w.Write(pag.Body)
			} else {
				renderTemplate(w, r, "web", pag, session)

			}
		}
	}()

	p.Session = session
	p.R = r

	t := template.New("PageWrapper")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(p.Body))
	outp := new(bytes.Buffer)
	err := t.Execute(outp, p)
	if err != nil {
		log.Println(err.Error())
		DebugTemplate(w, r, fmt.Sprintf("web%s", r.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		pag, err := loadPage("")
		if err != nil {
			log.Println(err.Error())
			return false
		}
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, r, "web", pag, session)

		}
		return false
	}

	p.Session.Save(r, w)

	fmt.Fprintf(w, html.UnescapeString(outp.String()))

	return true

}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string, *sessions.Session)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var session *sessions.Session
		var er error
		if session, er = store.Get(r, "session-"); er != nil {
			session, _ = store.New(r, "session-")
		}
		if attmpt := apiAttempt(w, r, session); !attmpt {
			fn(w, r, "", session)
		}

		session = nil
		context.Clear(r)
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
			sapp := net_getApp(getApps(), r.FormValue("id"))
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

			response = net_bPackageEdit(editor)

		} else if r.FormValue("type") == "2" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.Variables {
				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "is", Type: "text", Text: "Variable type", Value: v.Type})
				varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Variable name", Value: v.Name})
				response = response + net_bRPUT(rPut{DLink: "/api/delete?type=0&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Count: "4", Link: "/api/act?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Inputs: varf})
			}

		} else if r.FormValue("type") == "3" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.RootImports {
				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Package path", Value: v.Src})

				response = response + net_bRPUT(rPut{DLink: "/api/delete?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Count: "6", Link: "/api/act?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Inputs: varf})
			}

		} else if r.FormValue("type") == "4" {
			sapp := net_getApp(getApps(), r.FormValue("pkg"))

			for _, v := range sapp.Css {

				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "src", Type: "text", Text: "Path to css lib", Value: v})

				response = response + net_bRPUT(rPut{DLink: "/api/delete?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v, Count: "6", Link: "/api/act?type=3&pkg=" + r.FormValue("pkg") + "&id=" + v, Inputs: varf})
			}

		} else if r.FormValue("type") == "5" {
			id := strings.Split(r.FormValue("id"), "@pkg:")
			data, _ := ioutil.ReadFile(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/tmpl/" + id[1] + ".tmpl")

			data = []byte(html.EscapeString(string(data)))
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

			template := net_getTemplate(gos.Templates.Templates, id[1])

			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Value: template.Struct, Name: "struct", Text: "Interface to use with template"})

			response = net_bTemplateEdit(TemplateEdits{SavesTo: "tmpl/" + id[1] + ".tmpl", ID: net_RandTen(), PKG: r.FormValue("space"), Mime: "html", File: data, Settings: rPut{Link: "/api/put?type=2&id=" + id[1] + "&pkg=" + r.FormValue("space"), Inputs: varf, Count: "6"}})
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
			response = net_bWebRootEdit(WebRootEdits{SavesTo: id[1], Type: ftype, File: data, ID: net_RandTen(), PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "60" {
			id := strings.Split(r.FormValue("id"), "@pkg:")
			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + id[1]

			data, _ := ioutil.ReadFile(filep)
			data = []byte(html.EscapeString(string(data)))
			response = net_bWebRootEdittwo(WebRootEdits{SavesTo: id[1], Type: "golang", File: data, ID: net_RandTen(), PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "7" {
			sapp := net_getApp(getApps(), r.FormValue("space"))
			response = net_bROC(sROC{Name: r.FormValue("space"), Build: sapp.Passed, Time: sapp.LatestBuild, Pid: sapp.Pid})
		} else if r.FormValue("type") == "8" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/structs.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = net_bStructEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "9" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/objects.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = net_bObjectEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "10" {

			filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/methods.dsl"

			b, e := ioutil.ReadFile(filep)
			if e != nil {
				b = []byte("<gos&gt; \n \n </gos&gt; ")
			} else {
				b = []byte(html.EscapeString(string(b[:len(b)])))
			}
			response = net_bMethodEditor(vHuf{Edata: b, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "11" {

			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "path", Type: "text", Text: "Endpoint path"})
			kput := rPut{ListLink: "/api/get?type=13&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=7&space=" + r.FormValue("space")}
			response = net_bEndpointEditor(TEditor{CreateForm: kput, PKG: r.FormValue("space")})

		} else if r.FormValue("type") == "12" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Timer name"})
			kput := rPut{ListLink: "/api/get?type=14&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=8&space=" + r.FormValue("space")}
			response = net_bTimerEditor(TEditor{CreateForm: kput, PKG: r.FormValue("space")})
		} else if r.FormValue("type") == "13" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

			for _, v := range gos.Endpoints.Endpoints {

				varf := []Inputs{}
				varf = append(varf, Inputs{Name: "path", Type: "text", Text: "Endpoint path", Value: v.Path})
				//varf = append(varf, Inputs{Name:"method", Type:"text",Text:"Endpoint method",Value:v.Method})
				varf = append(varf, Inputs{Name: "typ", Type: "text", Text: "Request type : GET,POST,PUT,DELETE,f,star...", Value: v.Type})

				response = response + net_bRPUT(rPut{DLink: "/api/delete?type=7&pkg=" + r.FormValue("space") + "&path=" + v.Id, Link: "/api/put?type=9&id=" + v.Id + "&pkg=" + r.FormValue("space"), Count: "12", Inputs: varf}) + addjsstr

			}

		} else if r.FormValue("type") == "13r" {

			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			for _, v := range gos.Endpoints.Endpoints {

				if v.Id == r.FormValue("id") {
					id := net_RandTen()
					response = net_bTemplateEditTwo(TemplateEdits{SavesTo: "gosforceasapi/" + r.FormValue("id") + "++()/", ID: id, PKG: r.FormValue("pkg"), Mime: "golang", File: []byte(v.Method)})
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
				response = response + net_bRPUT(rPut{DLink: "/api/delete?type=8&pkg=" + r.FormValue("space") + "&name=" + v.Name, Link: "/api/put?type=10&id=" + v.Name + "&pkg=" + r.FormValue("space"), Count: "2", Inputs: varf})

			}

		} else if r.FormValue("type") == "15" {

			tempx := net_buSettings(USettings{StripeID: me.StripeID, LastPaid: "Date", Email: me.Email})
			response = net_bModal(sModal{Title: "Account settings", Body: tempx, Color: "orange"})
		} else if r.FormValue("type") == "16" {

			response = net_bDebugger(DebugObj{PKG: r.FormValue("space"), Username: ""})
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

			response = net_bDebuggerNode(tDebugNode)

		} else if r.FormValue("type") == "18" {

			response = net_bEndpointTesting(Dex{Misc: r.FormValue("space")})

		}
		callmet = true

	}

	if isURL := (r.URL.Path == "/api/pkg-bugs" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		bugs := GetLogs(r.FormValue("pkg"))
		sapp := net_getApp(getApps(), r.FormValue("pkg"))
		if len(bugs) == 0 || sapp.Passed {
			response = "{}"
		} else {
			response = mResponse(bugs[0])
		}

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/empty" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		ClearLogs(r.FormValue("pkg"))
		response = net_bAlert(Alertbs{Type: "success", Text: "Your build logs are cleared."})

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
			app := net_getApp(getApps(), r.FormValue("pkg"))
			app.Css = append(app.Css, r.FormValue("src"))
			apps = net_updateApp(getApps(), r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "3" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Bundle name"})

			response = net_bForm(Forms{Link: "/api/act?type=4&pkg=" + r.FormValue("pkg"), CTA: "Create Bundle", Class: "warning", Inputs: varf})

		} else if r.FormValue("type") == "4" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Name: "name", Type: "text", Text: "Template name"})

			response = net_bForm(Forms{Link: "/api/act?type=5&pkg=" + r.FormValue("pkg") + "&bundle=" + r.FormValue("bundle"), CTA: "Create Template file", Class: "warning", Inputs: varf})

		} else if r.FormValue("type") == "5" {
			//prefix pkg
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Text: "Path"})
			varf = append(varf, Inputs{Type: "hidden", Name: "basesix"})
			varf = append(varf, Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

			response = net_bFSC(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=6&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
		} else if r.FormValue("type") == "50" {
			//prefix pkg
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Text: "Path"})
			varf = append(varf, Inputs{Type: "hidden", Name: "basesix"})
			varf = append(varf, Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

			response = net_bFSC(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=60&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
		} else if r.FormValue("type") == "6" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

			response = net_bMV(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=7&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
		} else if r.FormValue("type") == "60" {
			varf := []Inputs{}
			varf = append(varf, Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

			response = net_bMV(FSCs{Path: r.FormValue("path"), Form: Forms{Link: "/api/act?type=70&pkg=" + r.FormValue("pkg") + "&folder=" + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
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
			app := net_getApp(apps, r.FormValue("pkg"))
			temp := []string{}
			for _, v := range app.Css {
				if v != r.FormValue("id") {
					temp = append(temp, v)
				}
			}
			app.Css = temp
			apps = net_updateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "3" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = net_bDelete(DForm{Text: "Are you sure you want to delete the package " + r.FormValue("pkg"), Link: "type=3&pkg=" + r.FormValue("pkg")})

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

				response = net_bAlert(Alertbs{Type: "success", Text: "Success package " + r.FormValue("pkg") + " was removed. Please reload page to close all linked resources.", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "4" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = net_bDelete(DForm{Text: "Are you sure you want to delete the bundle " + r.FormValue("bundle") + " and all of its sub templates", Link: "type=4&bundle=" + r.FormValue("bundle") + "&pkg=" + r.FormValue("pkg")})
			} else {
				//delete bundle
				apps := getApps()
				sapp := net_getApp(apps, r.FormValue("pkg"))

				replac := []string{}

				for _, v := range sapp.Groups {

					if r.FormValue("bundle") != v {
						replac = append(replac, v)
					}

				}

				sapp.Groups = replac
				apps = net_updateApp(apps, sapp.Name, sapp)
				saveApps(apps)
				gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
				os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/tmpl/" + r.FormValue("bundle"))
				gos.Delete("bundle", r.FormValue("bundle"))

				gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

				response = net_bAlert(Alertbs{Type: "success", Text: "Success bundle was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "5" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = net_bDelete(DForm{Text: "Are you sure you want to delete the template " + r.FormValue("tmpl"), Link: "type=5&tmpl=" + r.FormValue("tmpl") + "&pkg=" + r.FormValue("pkg")})
			} else {
				//delete

				os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/tmpl/" + r.FormValue("tmpl") + ".tmpl")

				gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
				parsedStr := strings.Split(r.FormValue("tmpl"), "/")
				gos.Delete("template", parsedStr[len(parsedStr)-1])

				gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

				response = net_bAlert(Alertbs{Type: "success", Text: "Success template " + r.FormValue("tmpl") + " was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "6" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = net_bDelete(DForm{Text: "Are you sure you want to delete the web resource at " + r.FormValue("path"), Link: "type=6&conf=do&path=" + r.FormValue("path") + "&pkg=" + r.FormValue("pkg")})

			} else {
				//delete
				if r.FormValue("isDir") == "Yes" {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/web" + r.FormValue("path"))
				} else {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/web" + r.FormValue("path"))
				}
				response = net_bAlert(Alertbs{Type: "success", Text: "Success resource at " + r.FormValue("path") + " was removed!", Redirect: "javascript:updateTree()"})
			}

		} else if r.FormValue("type") == "60" {
			//pkg
			if r.FormValue("conf") != "do" {
				response = net_bDelete(DForm{Text: "Are you sure you want to delete the resource at " + r.FormValue("path"), Link: "type=60&conf=do&path=" + r.FormValue("path") + "&pkg=" + r.FormValue("pkg")})

			} else {
				//delete
				if r.FormValue("isDir") == "Yes" {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/" + r.FormValue("path"))
				} else {
					os.RemoveAll(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/" + r.FormValue("path"))
				}
				response = net_bAlert(Alertbs{Type: "success", Text: "Success resource at " + r.FormValue("path") + " removed!", Redirect: "javascript:updateTree()"})
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
			response = net_bModal(sModal{Body: "", Title: "New Package", Color: "#ededed", Form: Forms{Link: "/api/act", CTA: "Create Package", Class: "warning btn-block", Buttons: []sButton{}, Inputs: inputs}})
		} else if r.FormValue("type") == "100" {
			inputs := []Inputs{}
			inputs = append(inputs, Inputs{Type: "text", Name: "name", Misc: "required", Text: "Plugin install path"})
			inputs = append(inputs, Inputs{Type: "hidden", Name: "type", Value: "100"})
			response = net_bModal(sModal{Body: "", Title: "PLUGINS", Color: "#ededed", Form: Forms{Link: "/api/act", CTA: "ADD", Class: "warning btn-block", Buttons: []sButton{}, Inputs: inputs}})
		} else if r.FormValue("type") == "101" {

			response = bPluginList(NoStruct{})
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
			response = net_bAlert(Alertbs{Type: "warning", Text: "Success package " + r.FormValue("name") + " was created!", Redirect: "javascript:updateTree()"})
		} else if r.FormValue("type") == "100" {
			plugins := getPlugins()
			plugins = append(plugins, r.FormValue("name"))

			//Users.Update(bson.M{"uid": me.UID}, me)

			_, err := core.RunCmdSmart("go get " + r.FormValue("name"))
			if err != nil {
				response = net_bAlert(Alertbs{Type: "warning", Text: "Error, could not find plugin.", Redirect: "#"})
			} else {
				savePlugins(plugins)
				response = net_bAlert(Alertbs{Type: "success", Text: "Success plugin " + r.FormValue("name") + " installed! Reload the page to activate plugin.", Redirect: "javascript:GetPlugins()"})
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
			app := net_getApp(apps, r.FormValue("pkg"))
			temp := []string{}
			for _, v := range app.Css {
				if v != r.FormValue("id") {
					temp = append(temp, v)
				} else {
					temp = append(temp, r.FormValue("src"))
				}
			}
			app.Css = temp
			apps = net_updateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "4" {
			apps := getApps()
			app := net_getApp(apps, r.FormValue("pkg"))

			app.Groups = append(app.Groups, r.FormValue("name"))
			os.MkdirAll(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/tmpl/"+r.FormValue("name"), 0777)
			apps = net_updateApp(apps, r.FormValue("pkg"), app)
			saveApps(apps)
			//Users.Update(bson.M{"uid":me.UID}, me)
		} else if r.FormValue("type") == "5" {
			//app := net_getApp(me.Apps, r.FormValue("pkg"))
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
				response = net_bAlert(Alertbs{Type: "danger", Text: "Failed to move resource : " + err.Error()})
			} else {
				response = net_bAlert(Alertbs{Type: "success", Text: "Operation succeeded"})
			}

		} else if r.FormValue("type") == "70" {
			err := os.Rename(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("prefix"), os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/"+r.FormValue("path"))
			if err != nil {
				response = net_bAlert(Alertbs{Type: "danger", Text: "Failed to move resource : " + err.Error()})
			} else {
				response = net_bAlert(Alertbs{Type: "success", Text: "Operation succeeded"})
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
			response = net_bAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
		} else if r.FormValue("type") == "2" {
			gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

			gos.Update("template", r.FormValue("id"), r.FormValue("struct"))
			// fmt.Println(gos)

			gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
			response = "Template interface saved!"
		} else if r.FormValue("type") == "3" {
			ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+"/web"+r.FormValue("target"), []byte(r.FormValue("data")), 0777)

			response = net_bAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
		} else if r.FormValue("type") == "30" {
			ioutil.WriteFile(os.ExpandEnv("$GOPATH")+"/src/"+r.FormValue("pkg")+r.FormValue("target"), []byte(r.FormValue("data")), 0777)

			response = net_bAlert(Alertbs{Type: "warning", Text: r.FormValue("target") + " saved!"})
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

			response = net_bAlert(Alertbs{Type: "warning", Text: "Interfaces saved!"})

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

			response = net_bAlert(Alertbs{Type: "warning", Text: "Objects saved!"})

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

			response = net_bAlert(Alertbs{Type: "warning", Text: "Pipelines saved!"})

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

			response = net_bAlert(Alertbs{Type: "warning", Text: "Endpoint code saved!"})

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
				response = net_bAlert(Alertbs{Type: "success", Text: "Password updated"})
			} else {
				response = net_bAlert(Alertbs{Type: "danger", Text: "Error incorrect current password"})
			}

		} else if r.FormValue("type") == "12" {
			me.Email = r.FormValue("email")
			response = net_bAlert(Alertbs{Type: "success", Text: "Email updated"})
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
		//	fmt.Println(logBuilt)
		passed := false

		if !strings.Contains(logBuilt, "Your build failed,") {
			logBuilt, _ = core.RunCmdSmart("go build")
			if logBuilt != "" {

				debuglog := DebugObj{r.FormValue("pkg"), net_RandTen(), "", logBuilt, time.Now().String(), []DebugNode{}}

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
				response = net_bAlert(Alertbs{Type: "danger", Text: "Your build failed, checkout the logs to see why!"})

			} else {
				passed = true
				response = net_bAlert(Alertbs{Type: "success", Text: "Your build passed!"})
			}
		} else {
			debuglog := DebugObj{r.FormValue("pkg"), net_RandTen(), "", logBuilt, time.Now().String(), []DebugNode{}}
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
			response = net_bAlert(Alertbs{Type: "danger", Text: "Your build failed, checkout the logs to see why!"})

		}

		//DebugLogs.Insert(dObj)

		apps := getApps()
		sapp := net_getApp(apps, r.FormValue("pkg"))

		sapp.Passed = passed
		sapp.LatestBuild = time.Now().String()
		apps = net_updateApp(apps, r.FormValue("pkg"), sapp)
		saveApps(apps)

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/start" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")

		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		apps := getApps()
		sapp := net_getApp(apps, r.FormValue("pkg"))

		if sapp.Passed {

			if sapp.Pid != "" {
				core.RunCmdB("kill -3 " + sapp.Pid)
				response = net_bAlert(Alertbs{Type: "success", Text: "Build stopped."})
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
				response = net_bAlert(Alertbs{Type: "success", Text: "Server up"})
			} else {
				response = net_bAlert(Alertbs{Type: "success", Text: "Your server is up at PID : " + sapp.Pid})
			}
		} else {
			response = net_bAlert(Alertbs{Type: "danger", Text: "Your latest build failed."})

		}

		//DebugLogs.Insert(dObj)
		apps = net_updateApp(apps, r.FormValue("pkg"), sapp)
		saveApps(apps)

		//Users.Update(bson.M{"uid":me.UID}, me)

		callmet = true
	}

	if isURL := (r.URL.Path == "/api/stop" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		gp := os.ExpandEnv("$GOPATH")

		os.Chdir(gp + "/src/" + r.FormValue("pkg"))
		apps := getApps()
		sapp := net_getApp(apps, r.FormValue("pkg"))

		if sapp.Pid == "" {
			response = net_bAlert(Alertbs{Type: "danger", Text: "No build running."})
		} else {
			core.RunCmdB("kill -3 " + sapp.Pid)
			response = net_bAlert(Alertbs{Type: "success", Text: "Build stopped."})
		}

		//DebugLogs.Insert(dObj)
		apps = net_updateApp(apps, r.FormValue("pkg"), sapp)
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
				ret = append(ret, bson.M{"name": v.Name, "value": "net_" + v.Name + "(" + v.Variables + ")", "score": score, "meta": "Method"})
			}

		}

		for _, v := range gos.Templates.Templates {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name + " ", "score": score, "meta": "{{Template reference}}"})
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": "net_b" + v.Name + "(" + v.Struct + "{})", "score": score, "meta": "Template"})
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
				t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func handler(w http.ResponseWriter, r *http.Request, contxt string, session *sessions.Session) {
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
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, r, "web", pag, session)
		}
		return
	}

	if !p.isResource {
		w.Header().Set("Content-Type", "text/html")

		renderTemplate(w, r, fmt.Sprintf("web%s", r.URL.Path), p, session)

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

	//context.Clear(r)

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

func net_castFSCs(args ...interface{}) *FSCs {

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
func net_structFSCs() *FSCs { return &FSCs{} }

type Dex struct {
	Misc string
	Text string
	Link string
}

func net_castDex(args ...interface{}) *Dex {

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
func net_structDex() *Dex { return &Dex{} }

type SoftUser struct {
	Username         string
	Email            string
	Password         []byte
	Apps             []App
	Docker           string
	TrialEnd         int64
	StripeID, FLogin string
}

func net_castSoftUser(args ...interface{}) *SoftUser {

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
func net_structSoftUser() *SoftUser { return &SoftUser{} }

type USettings struct {
	LastPaid string
	Email    string
	StripeID string
}

func net_castUSettings(args ...interface{}) *USettings {

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
func net_structUSettings() *USettings { return &USettings{} }

type App struct {
	Type             string
	Name             string
	PublicName       string
	Css              []string
	Groups           []string
	Passed, Running  bool
	LatestBuild, Pid string
}

func net_castApp(args ...interface{}) *App {

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
func net_structApp() *App { return &App{} }

type TemplateEdits struct {
	SavesTo, PKG, PreviewLink, ID, Mime string
	File                                []byte
	Settings                            rPut
}

func net_castTemplateEdits(args ...interface{}) *TemplateEdits {

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
func net_structTemplateEdits() *TemplateEdits { return &TemplateEdits{} }

type WebRootEdits struct {
	SavesTo, Type, PreviewLink, ID, PKG string
	File                                []byte
}

func net_castWebRootEdits(args ...interface{}) *WebRootEdits {

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
func net_structWebRootEdits() *WebRootEdits { return &WebRootEdits{} }

type TEditor struct {
	PKG, Type, LType string
	CreateForm       rPut
}

func net_castTEditor(args ...interface{}) *TEditor {

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
func net_structTEditor() *TEditor { return &TEditor{} }

type Navbars struct {
	Mode string
	ID   string
}

func net_castNavbars(args ...interface{}) *Navbars {

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
func net_structNavbars() *Navbars { return &Navbars{} }

type sModal struct {
	Title   string
	Body    string
	Color   string
	Buttons []sButton
	Form    Forms
}

func net_castsModal(args ...interface{}) *sModal {

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
func net_structsModal() *sModal { return &sModal{} }

type Forms struct {
	Link    string
	Inputs  []Inputs
	Buttons []sButton
	CTA     string
	Class   string
}

func net_castForms(args ...interface{}) *Forms {

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
func net_structForms() *Forms { return &Forms{} }

type sButton struct {
	Text  string
	Class string
	Link  string
}

func net_castsButton(args ...interface{}) *sButton {

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
func net_structsButton() *sButton { return &sButton{} }

type sTab struct {
	Buttons []sButton
}

func net_castsTab(args ...interface{}) *sTab {

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
func net_structsTab() *sTab { return &sTab{} }

type DForm struct {
	Text, Link string
}

func net_castDForm(args ...interface{}) *DForm {

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
func net_structDForm() *DForm { return &DForm{} }

type Alertbs struct {
	Type     string
	Text     string
	Redirect string
}

func net_castAlertbs(args ...interface{}) *Alertbs {

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
func net_structAlertbs() *Alertbs { return &Alertbs{} }

type Inputs struct {
	Misc    string
	Text    string
	Name    string
	Type    string
	Options []string
	Value   string
}

func net_castInputs(args ...interface{}) *Inputs {

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
func net_structInputs() *Inputs { return &Inputs{} }

type Aput struct {
	Link, Param, Value string
}

func net_castAput(args ...interface{}) *Aput {

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
func net_structAput() *Aput { return &Aput{} }

type rPut struct {
	Link     string
	DLink    string
	Inputs   []Inputs
	Count    string
	ListLink string
}

func net_castrPut(args ...interface{}) *rPut {

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
func net_structrPut() *rPut { return &rPut{} }

type sSWAL struct {
	Title, Type, Text string
}

func net_castsSWAL(args ...interface{}) *sSWAL {

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
func net_structsSWAL() *sSWAL { return &sSWAL{} }

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

func net_castsPackageEdit(args ...interface{}) *sPackageEdit {

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
func net_structsPackageEdit() *sPackageEdit { return &sPackageEdit{} }

type DebugObj struct {
	PKG, Id, Username, RawLog, Time string
	Bugs                            []DebugNode
}

func net_castDebugObj(args ...interface{}) *DebugObj {

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
func net_structDebugObj() *DebugObj { return &DebugObj{} }

type DebugNode struct {
	Action, Line, CTA string
}

func net_castDebugNode(args ...interface{}) *DebugNode {

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
func net_structDebugNode() *DebugNode { return &DebugNode{} }

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

func net_castPkgItem(args ...interface{}) *PkgItem {

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
func net_structPkgItem() *PkgItem { return &PkgItem{} }

type sROC struct {
	Name      string
	CompLog   []byte
	Build     bool
	Time, Pid string
}

func net_castsROC(args ...interface{}) *sROC {

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
func net_structsROC() *sROC { return &sROC{} }

type vHuf struct {
	Type, PKG string
	Edata     []byte
}

func net_castvHuf(args ...interface{}) *vHuf {

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
func net_structvHuf() *vHuf { return &vHuf{} }

type myDemoObject Dex

func net_BindMisc(args ...interface{}) Dex {
	misc := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = misc.(string)
	return Nav

}
func net_ListPlugins(args ...interface{}) []string {

	return getPlugins()

}
func net_BindID(args ...interface{}) Dex {
	id := args[0]
	nav := args[1]

	Nav := nav.(Dex)
	Nav.Misc = id.(string)
	return Nav

}
func net_RandTen(args ...interface{}) string {

	return core.NewLen(10)

}
func net_Fragmentize(args ...interface{}) (finall string) {
	inn := args[0]

	finall = strings.Replace(inn.(string), ".tmpl", "", -1)
	return

}
func net_parseLog(args ...interface{}) string {
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
func net_anyBugs(args ...interface{}) (ajet bool) {
	packge := args[0]

	ajet = false //,err := DebugLogs.Find(bson.M{"pkg":packge.(string), "username":usernam.(string)}).Count()

	bugs := GetLogs(packge.(string))
	sapp := net_getApp(getApps(), packge.(string))

	if len(bugs) > 0 {
		ajet = true
	}

	if sapp.Pid != "" {
		ajet = true
	}

	return

}
func net_PluginJS(args ...interface{}) string {

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
func net_FindmyBugs(args ...interface{}) (ajet []DebugObj) {
	packge := args[0]

	ajet = GetLogs(packge.(string))
	sapp := net_getApp(getApps(), packge.(string))

	if sapp.Pid != "" {
		activLog := DebugObj{Time: "Server", Bugs: []DebugNode{}}
		ajet = append([]DebugObj{activLog}, ajet...)
	}
	return

}
func net_isExpired(args ...interface{}) bool {
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
func net_getTemplate(args ...interface{}) core.Template {
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
func net_mConsole(args ...interface{}) string {

	return ""

}
func net_mPut(args ...interface{}) string {

	//response = "OK"

	return ""

}
func net_updateApp(args ...interface{}) []App {
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
func net_getApp(args ...interface{}) App {
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

func net_Css(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bCss(d Dex) string {
	return net_bCss(d)
}

func net_bCss(d Dex) string {
	filename := "tmpl/css.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Css")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cCss(args ...interface{}) (d Dex) {
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

func net_JS(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bJS(d Dex) string {
	return net_bJS(d)
}

func net_bJS(d Dex) string {
	filename := "tmpl/js.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("JS")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cJS(args ...interface{}) (d Dex) {
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

func net_FA(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bFA(d Dex) string {
	return net_bFA(d)
}

func net_bFA(d Dex) string {
	filename := "tmpl/ui/fa.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FA")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cFA(args ...interface{}) (d Dex) {
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

func net_PluginList(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPluginList(d NoStruct) string {
	return net_bPluginList(d)
}

func net_bPluginList(d NoStruct) string {
	filename := "tmpl/ui/pluginlist.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PluginList")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cPluginList(args ...interface{}) (d NoStruct) {
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

func net_Login(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bLogin(d Dex) string {
	return net_bLogin(d)
}

func net_bLogin(d Dex) string {
	filename := "tmpl/ui/login.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Login")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cLogin(args ...interface{}) (d Dex) {
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

func net_Modal(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bModal(d sModal) string {
	return net_bModal(d)
}

func net_bModal(d sModal) string {
	filename := "tmpl/ui/modal.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Modal")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cModal(args ...interface{}) (d sModal) {
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

func net_xButton(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bxButton(d sButton) string {
	return net_bxButton(d)
}

func net_bxButton(d sButton) string {
	filename := "tmpl/ui/sbutton.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("xButton")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cxButton(args ...interface{}) (d sButton) {
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

func net_jButton(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bjButton(d sButton) string {
	return net_bjButton(d)
}

func net_bjButton(d sButton) string {
	filename := "tmpl/ui/user/forms/jbutton.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("jButton")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cjButton(args ...interface{}) (d sButton) {
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

func net_PUT(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPUT(d Aput) string {
	return net_bPUT(d)
}

func net_bPUT(d Aput) string {
	filename := "tmpl/ui/user/forms/aput.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PUT")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cPUT(args ...interface{}) (d Aput) {
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

func net_Group(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bGroup(d sTab) string {
	return net_bGroup(d)
}

func net_bGroup(d sTab) string {
	filename := "tmpl/ui/user/forms/tab.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Group")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cGroup(args ...interface{}) (d sTab) {
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

func net_Register(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRegister(d Dex) string {
	return net_bRegister(d)
}

func net_bRegister(d Dex) string {
	filename := "tmpl/ui/register.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Register")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cRegister(args ...interface{}) (d Dex) {
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

func net_Alert(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bAlert(d Alertbs) string {
	return net_bAlert(d)
}

func net_bAlert(d Alertbs) string {
	filename := "tmpl/ui/alert.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Alert")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cAlert(args ...interface{}) (d Alertbs) {
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

func net_StructEditor(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bStructEditor(d vHuf) string {
	return net_bStructEditor(d)
}

func net_bStructEditor(d vHuf) string {
	filename := "tmpl/editor/structs.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("StructEditor")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cStructEditor(args ...interface{}) (d vHuf) {
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

func net_MethodEditor(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bMethodEditor(d vHuf) string {
	return net_bMethodEditor(d)
}

func net_bMethodEditor(d vHuf) string {
	filename := "tmpl/editor/methods.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MethodEditor")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cMethodEditor(args ...interface{}) (d vHuf) {
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

func net_ObjectEditor(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bObjectEditor(d vHuf) string {
	return net_bObjectEditor(d)
}

func net_bObjectEditor(d vHuf) string {
	filename := "tmpl/editor/objects.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ObjectEditor")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cObjectEditor(args ...interface{}) (d vHuf) {
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

func net_EndpointEditor(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bEndpointEditor(d TEditor) string {
	return net_bEndpointEditor(d)
}

func net_bEndpointEditor(d TEditor) string {
	filename := "tmpl/editor/endpoints.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointEditor")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cEndpointEditor(args ...interface{}) (d TEditor) {
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

func net_TimerEditor(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTimerEditor(d TEditor) string {
	return net_bTimerEditor(d)
}

func net_bTimerEditor(d TEditor) string {
	filename := "tmpl/editor/timers.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TimerEditor")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cTimerEditor(args ...interface{}) (d TEditor) {
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

func net_FSC(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bFSC(d FSCs) string {
	return net_bFSC(d)
}

func net_bFSC(d FSCs) string {
	filename := "tmpl/ui/fsc.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("FSC")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cFSC(args ...interface{}) (d FSCs) {
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

func net_MV(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bMV(d FSCs) string {
	return net_bMV(d)
}

func net_bMV(d FSCs) string {
	filename := "tmpl/ui/mv.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("MV")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cMV(args ...interface{}) (d FSCs) {
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

func net_RM(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRM(d FSCs) string {
	return net_bRM(d)
}

func net_bRM(d FSCs) string {
	filename := "tmpl/ui/user/rm.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RM")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cRM(args ...interface{}) (d FSCs) {
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

func net_WebRootEdit(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWebRootEdit(d WebRootEdits) string {
	return net_bWebRootEdit(d)
}

func net_bWebRootEdit(d WebRootEdits) string {
	filename := "tmpl/ui/user/panel/webrootedit.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdit")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cWebRootEdit(args ...interface{}) (d WebRootEdits) {
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

func net_WebRootEdittwo(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWebRootEdittwo(d WebRootEdits) string {
	return net_bWebRootEdittwo(d)
}

func net_bWebRootEdittwo(d WebRootEdits) string {
	filename := "tmpl/ui/user/panel/webtwo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("WebRootEdittwo")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cWebRootEdittwo(args ...interface{}) (d WebRootEdits) {
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

func net_uSettings(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func buSettings(d USettings) string {
	return net_buSettings(d)
}

func net_buSettings(d USettings) string {
	filename := "tmpl/editor/settings.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("uSettings")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cuSettings(args ...interface{}) (d USettings) {
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

func net_Form(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bForm(d Forms) string {
	return net_bForm(d)
}

func net_bForm(d Forms) string {
	filename := "tmpl/ui/user/forms/form.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Form")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cForm(args ...interface{}) (d Forms) {
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

func net_SWAL(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bSWAL(d sSWAL) string {
	return net_bSWAL(d)
}

func net_bSWAL(d sSWAL) string {
	filename := "tmpl/ui/user/forms/swal.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("SWAL")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cSWAL(args ...interface{}) (d sSWAL) {
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

func net_ROC(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bROC(d sROC) string {
	return net_bROC(d)
}

func net_bROC(d sROC) string {
	filename := "tmpl/ui/user/panel/roc.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ROC")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cROC(args ...interface{}) (d sROC) {
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

func net_RPUT(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bRPUT(d rPut) string {
	return net_bRPUT(d)
}

func net_bRPUT(d rPut) string {
	filename := "tmpl/ui/user/forms/rput.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("RPUT")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cRPUT(args ...interface{}) (d rPut) {
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

func net_PackageEdit(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bPackageEdit(d sPackageEdit) string {
	return net_bPackageEdit(d)
}

func net_bPackageEdit(d sPackageEdit) string {
	filename := "tmpl/ui/user/panel/package.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("PackageEdit")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cPackageEdit(args ...interface{}) (d sPackageEdit) {
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

func net_Delete(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDelete(d DForm) string {
	return net_bDelete(d)
}

func net_bDelete(d DForm) string {
	filename := "tmpl/ui/user/panel/delete.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Delete")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cDelete(args ...interface{}) (d DForm) {
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

func net_Welcome(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bWelcome(d Dex) string {
	return net_bWelcome(d)
}

func net_bWelcome(d Dex) string {
	filename := "tmpl/ui/welcome.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Welcome")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cWelcome(args ...interface{}) (d Dex) {
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

func net_Stripe(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bStripe(d Dex) string {
	return net_bStripe(d)
}

func net_bStripe(d Dex) string {
	filename := "tmpl/ui/stripe.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Stripe")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cStripe(args ...interface{}) (d Dex) {
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

func net_Debugger(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDebugger(d DebugObj) string {
	return net_bDebugger(d)
}

func net_bDebugger(d DebugObj) string {
	filename := "tmpl/ui/debugger.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Debugger")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cDebugger(args ...interface{}) (d DebugObj) {
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

func net_TemplateEdit(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTemplateEdit(d TemplateEdits) string {
	return net_bTemplateEdit(d)
}

func net_bTemplateEdit(d TemplateEdits) string {
	filename := "tmpl/ui/user/panel/templateEditor.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEdit")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cTemplateEdit(args ...interface{}) (d TemplateEdits) {
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

func net_TemplateEditTwo(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bTemplateEditTwo(d TemplateEdits) string {
	return net_bTemplateEditTwo(d)
}

func net_bTemplateEditTwo(d TemplateEdits) string {
	filename := "tmpl/ui/user/panel/tpetwo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("TemplateEditTwo")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cTemplateEditTwo(args ...interface{}) (d TemplateEdits) {
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

func net_Input(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bInput(d Inputs) string {
	return net_bInput(d)
}

func net_bInput(d Inputs) string {
	filename := "tmpl/ui/input.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Input")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cInput(args ...interface{}) (d Inputs) {
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

func net_DebuggerNode(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bDebuggerNode(d DebugObj) string {
	return net_bDebuggerNode(d)
}

func net_bDebuggerNode(d DebugObj) string {
	filename := "tmpl/ui/debugnode.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("DebuggerNode")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cDebuggerNode(args ...interface{}) (d DebugObj) {
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

func net_Button(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bButton(d Dex) string {
	return net_bButton(d)
}

func net_bButton(d Dex) string {
	filename := "tmpl/ui/button.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Button")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cButton(args ...interface{}) (d Dex) {
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

func net_Submit(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bSubmit(d Dex) string {
	return net_bSubmit(d)
}

func net_bSubmit(d Dex) string {
	filename := "tmpl/ui/submit.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Submit")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cSubmit(args ...interface{}) (d Dex) {
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

func net_Logo(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bLogo(d Dex) string {
	return net_bLogo(d)
}

func net_bLogo(d Dex) string {
	filename := "tmpl/logo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Logo")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cLogo(args ...interface{}) (d Dex) {
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

func net_Navbar(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavbar(d Dex) string {
	return net_bNavbar(d)
}

func net_bNavbar(d Dex) string {
	filename := "tmpl/ui/navbar.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("Navbar")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cNavbar(args ...interface{}) (d Dex) {
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

func net_NavCustom(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavCustom(d Navbars) string {
	return net_bNavCustom(d)
}

func net_bNavCustom(d Navbars) string {
	filename := "tmpl/ui/navbars.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavCustom")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cNavCustom(args ...interface{}) (d Navbars) {
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

func net_NavMain(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavMain(d Dex) string {
	return net_bNavMain(d)
}

func net_bNavMain(d Dex) string {
	filename := "tmpl/ui/navmain.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavMain")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cNavMain(args ...interface{}) (d Dex) {
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

func net_NavPKG(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavPKG(d Dex) string {
	return net_bNavPKG(d)
}

func net_bNavPKG(d Dex) string {
	filename := "tmpl/ui/navpkg.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPKG")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cNavPKG(args ...interface{}) (d Dex) {
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

func net_CrashedPage(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bCrashedPage(d Dex) string {
	return net_bCrashedPage(d)
}

func net_bCrashedPage(d Dex) string {
	filename := "tmpl/ui/crashedpage.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("CrashedPage")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cCrashedPage(args ...interface{}) (d Dex) {
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

func net_EndpointTesting(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bEndpointTesting(d Dex) string {
	return net_bEndpointTesting(d)
}

func net_bEndpointTesting(d Dex) string {
	filename := "tmpl/ui/endpointtester.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("EndpointTesting")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cEndpointTesting(args ...interface{}) (d Dex) {
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

func net_NavPromo(args ...interface{}) string {
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
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bNavPromo(d Dex) string {
	return net_bNavPromo(d)
}

func net_bNavPromo(d Dex) string {
	filename := "tmpl/ui/navpromo.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("NavPromo")
	t = t.Funcs(template.FuncMap{"a": net_add, "s": net_subs, "m": net_multiply, "d": net_divided, "js": net_importjs, "css": net_importcss, "sd": net_sessionDelete, "sr": net_sessionRemove, "sc": net_sessionKey, "ss": net_sessionSet, "sso": net_sessionSetInt, "sgo": net_sessionGetInt, "sg": net_sessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "BindMisc": net_BindMisc, "ListPlugins": net_ListPlugins, "BindID": net_BindID, "RandTen": net_RandTen, "Fragmentize": net_Fragmentize, "parseLog": net_parseLog, "anyBugs": net_anyBugs, "PluginJS": net_PluginJS, "FindmyBugs": net_FindmyBugs, "isExpired": net_isExpired, "getTemplate": net_getTemplate, "mConsole": net_mConsole, "mPut": net_mPut, "updateApp": net_updateApp, "getApp": net_getApp, "Css": net_Css, "bCss": net_bCss, "cCss": net_cCss, "JS": net_JS, "bJS": net_bJS, "cJS": net_cJS, "FA": net_FA, "bFA": net_bFA, "cFA": net_cFA, "PluginList": net_PluginList, "bPluginList": net_bPluginList, "cPluginList": net_cPluginList, "Login": net_Login, "bLogin": net_bLogin, "cLogin": net_cLogin, "Modal": net_Modal, "bModal": net_bModal, "cModal": net_cModal, "xButton": net_xButton, "bxButton": net_bxButton, "cxButton": net_cxButton, "jButton": net_jButton, "bjButton": net_bjButton, "cjButton": net_cjButton, "PUT": net_PUT, "bPUT": net_bPUT, "cPUT": net_cPUT, "Group": net_Group, "bGroup": net_bGroup, "cGroup": net_cGroup, "Register": net_Register, "bRegister": net_bRegister, "cRegister": net_cRegister, "Alert": net_Alert, "bAlert": net_bAlert, "cAlert": net_cAlert, "StructEditor": net_StructEditor, "bStructEditor": net_bStructEditor, "cStructEditor": net_cStructEditor, "MethodEditor": net_MethodEditor, "bMethodEditor": net_bMethodEditor, "cMethodEditor": net_cMethodEditor, "ObjectEditor": net_ObjectEditor, "bObjectEditor": net_bObjectEditor, "cObjectEditor": net_cObjectEditor, "EndpointEditor": net_EndpointEditor, "bEndpointEditor": net_bEndpointEditor, "cEndpointEditor": net_cEndpointEditor, "TimerEditor": net_TimerEditor, "bTimerEditor": net_bTimerEditor, "cTimerEditor": net_cTimerEditor, "FSC": net_FSC, "bFSC": net_bFSC, "cFSC": net_cFSC, "MV": net_MV, "bMV": net_bMV, "cMV": net_cMV, "RM": net_RM, "bRM": net_bRM, "cRM": net_cRM, "WebRootEdit": net_WebRootEdit, "bWebRootEdit": net_bWebRootEdit, "cWebRootEdit": net_cWebRootEdit, "WebRootEdittwo": net_WebRootEdittwo, "bWebRootEdittwo": net_bWebRootEdittwo, "cWebRootEdittwo": net_cWebRootEdittwo, "uSettings": net_uSettings, "buSettings": net_buSettings, "cuSettings": net_cuSettings, "Form": net_Form, "bForm": net_bForm, "cForm": net_cForm, "SWAL": net_SWAL, "bSWAL": net_bSWAL, "cSWAL": net_cSWAL, "ROC": net_ROC, "bROC": net_bROC, "cROC": net_cROC, "RPUT": net_RPUT, "bRPUT": net_bRPUT, "cRPUT": net_cRPUT, "PackageEdit": net_PackageEdit, "bPackageEdit": net_bPackageEdit, "cPackageEdit": net_cPackageEdit, "Delete": net_Delete, "bDelete": net_bDelete, "cDelete": net_cDelete, "Welcome": net_Welcome, "bWelcome": net_bWelcome, "cWelcome": net_cWelcome, "Stripe": net_Stripe, "bStripe": net_bStripe, "cStripe": net_cStripe, "Debugger": net_Debugger, "bDebugger": net_bDebugger, "cDebugger": net_cDebugger, "TemplateEdit": net_TemplateEdit, "bTemplateEdit": net_bTemplateEdit, "cTemplateEdit": net_cTemplateEdit, "TemplateEditTwo": net_TemplateEditTwo, "bTemplateEditTwo": net_bTemplateEditTwo, "cTemplateEditTwo": net_cTemplateEditTwo, "Input": net_Input, "bInput": net_bInput, "cInput": net_cInput, "DebuggerNode": net_DebuggerNode, "bDebuggerNode": net_bDebuggerNode, "cDebuggerNode": net_cDebuggerNode, "Button": net_Button, "bButton": net_bButton, "cButton": net_cButton, "Submit": net_Submit, "bSubmit": net_bSubmit, "cSubmit": net_cSubmit, "Logo": net_Logo, "bLogo": net_bLogo, "cLogo": net_cLogo, "Navbar": net_Navbar, "bNavbar": net_bNavbar, "cNavbar": net_cNavbar, "NavCustom": net_NavCustom, "bNavCustom": net_bNavCustom, "cNavCustom": net_cNavCustom, "NavMain": net_NavMain, "bNavMain": net_bNavMain, "cNavMain": net_cNavMain, "NavPKG": net_NavPKG, "bNavPKG": net_bNavPKG, "cNavPKG": net_cNavPKG, "CrashedPage": net_CrashedPage, "bCrashedPage": net_bCrashedPage, "cCrashedPage": net_cCrashedPage, "EndpointTesting": net_EndpointTesting, "bEndpointTesting": net_bEndpointTesting, "cEndpointTesting": net_cEndpointTesting, "NavPromo": net_NavPromo, "bNavPromo": net_bNavPromo, "cNavPromo": net_cNavPromo, "FSCs": net_structFSCs, "isFSCs": net_castFSCs, "Dex": net_structDex, "isDex": net_castDex, "SoftUser": net_structSoftUser, "isSoftUser": net_castSoftUser, "USettings": net_structUSettings, "isUSettings": net_castUSettings, "App": net_structApp, "isApp": net_castApp, "TemplateEdits": net_structTemplateEdits, "isTemplateEdits": net_castTemplateEdits, "WebRootEdits": net_structWebRootEdits, "isWebRootEdits": net_castWebRootEdits, "TEditor": net_structTEditor, "isTEditor": net_castTEditor, "Navbars": net_structNavbars, "isNavbars": net_castNavbars, "sModal": net_structsModal, "issModal": net_castsModal, "Forms": net_structForms, "isForms": net_castForms, "sButton": net_structsButton, "issButton": net_castsButton, "sTab": net_structsTab, "issTab": net_castsTab, "DForm": net_structDForm, "isDForm": net_castDForm, "Alertbs": net_structAlertbs, "isAlertbs": net_castAlertbs, "Inputs": net_structInputs, "isInputs": net_castInputs, "Aput": net_structAput, "isAput": net_castAput, "rPut": net_structrPut, "isrPut": net_castrPut, "sSWAL": net_structsSWAL, "issSWAL": net_castsSWAL, "sPackageEdit": net_structsPackageEdit, "issPackageEdit": net_castsPackageEdit, "DebugObj": net_structDebugObj, "isDebugObj": net_castDebugObj, "DebugNode": net_structDebugNode, "isDebugNode": net_castDebugNode, "PkgItem": net_structPkgItem, "isPkgItem": net_castPkgItem, "sROC": net_structsROC, "issROC": net_castsROC, "vHuf": net_structvHuf, "isvHuf": net_castvHuf})
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
func net_cNavPromo(args ...interface{}) (d Dex) {
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
	http.HandleFunc("/", makeHandler(handler))

	http.Handle("/dist/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "web"}))

	errgos := http.ListenAndServe(port, nil)
	if errgos != nil {
		log.Fatal(errgos)
	}

}
