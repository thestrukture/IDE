package templates

import (
	"fmt"
	"html/template"
	"strings"

	gosweb "github.com/cheikhshift/gos/web"

	"github.com/thestrukture/IDE/api/assets"
	methods "github.com/thestrukture/IDE/api/methods"
	"github.com/thestrukture/IDE/types"
)

var TemplateFuncStore template.FuncMap

var templateCache = gosweb.NewTemplateCache()
var WebCache = gosweb.NewCache()

func StoreNetfn() int {
	// List of pipelines linked to each template.
	TemplateFuncStore = template.FuncMap{"a": gosweb.Netadd, "s": gosweb.Netsubs, "m": gosweb.Netmultiply, "d": gosweb.Netdivided, "js": gosweb.Netimportjs, "css": gosweb.Netimportcss, "sd": gosweb.NetsessionDelete, "sr": gosweb.NetsessionRemove, "sc": gosweb.NetsessionKey, "ss": gosweb.NetsessionSet, "sso": gosweb.NetsessionSetInt, "sgo": gosweb.NetsessionGetInt, "sg": gosweb.NetsessionGet, "form": gosweb.Formval, "eq": gosweb.Equalz, "neq": gosweb.Nequalz, "lte": gosweb.Netlt, "BindMisc": methods.BindMisc, "ListPlugins": methods.ListPlugins, "BindID": methods.BindID, "RandTen": methods.RandTen, "Fragmentize": methods.Fragmentize, "ParseLog": methods.ParseLog, "AnyBugs": methods.AnyBugs, "PluginJS": methods.PluginJS, "FindmyBugs": methods.FindmyBugs, "IsExpired": methods.IsExpired, "GetTemplate": methods.GetTemplate, "MConsole": methods.MConsole, "MPut": methods.MPut, "UpdateApp": methods.UpdateApp, "GetApp": methods.GetApp, "Css": NetCss, "bCss": NetbCss, "cCss": NetcCss, "JS": NetJS, "bJS": NetbJS, "cJS": NetcJS, "FA": NetFA, "bFA": NetbFA, "cFA": NetcFA, "PluginList": NetPluginList, "bPluginList": NetbPluginList, "cPluginList": NetcPluginList, "Login": NetLogin, "bLogin": NetbLogin, "cLogin": NetcLogin, "Modal": NetModal, "bModal": NetbModal, "cModal": NetcModal, "xButton": NetxButton, "bxButton": NetbxButton, "cxButton": NetcxButton, "jButton": NetjButton, "bjButton": NetbjButton, "cjButton": NetcjButton, "PUT": NetPUT, "bPUT": NetbPUT, "cPUT": NetcPUT, "Group": NetGroup, "bGroup": NetbGroup, "cGroup": NetcGroup, "Register": NetRegister, "bRegister": NetbRegister, "cRegister": NetcRegister, "Alert": NetAlert, "bAlert": NetbAlert, "cAlert": NetcAlert, "StructEditor": NetStructEditor, "bStructEditor": NetbStructEditor, "cStructEditor": NetcStructEditor, "MethodEditor": NetMethodEditor, "bMethodEditor": NetbMethodEditor, "cMethodEditor": NetcMethodEditor, "ObjectEditor": NetObjectEditor, "bObjectEditor": NetbObjectEditor, "cObjectEditor": NetcObjectEditor, "EndpointEditor": NetEndpointEditor, "bEndpointEditor": NetbEndpointEditor, "cEndpointEditor": NetcEndpointEditor, "TimerEditor": NetTimerEditor, "bTimerEditor": NetbTimerEditor, "cTimerEditor": NetcTimerEditor, "FSC": NetFSC, "bFSC": NetbFSC, "cFSC": NetcFSC, "MV": NetMV, "bMV": NetbMV, "cMV": NetcMV, "RM": NetRM, "bRM": NetbRM, "cRM": NetcRM, "WebRootEdit": NetWebRootEdit, "bWebRootEdit": NetbWebRootEdit, "cWebRootEdit": NetcWebRootEdit, "WebRootEdittwo": NetWebRootEdittwo, "bWebRootEdittwo": NetbWebRootEdittwo, "cWebRootEdittwo": NetcWebRootEdittwo, "uSettings": NetuSettings, "buSettings": NetbuSettings, "cuSettings": NetcuSettings, "Form": NetForm, "bForm": NetbForm, "cForm": NetcForm, "SWAL": NetSWAL, "bSWAL": NetbSWAL, "cSWAL": NetcSWAL, "ROC": NetROC, "bROC": NetbROC, "cROC": NetcROC, "RPUT": NetRPUT, "bRPUT": NetbRPUT, "cRPUT": NetcRPUT, "PackageEdit": NetPackageEdit, "bPackageEdit": NetbPackageEdit, "cPackageEdit": NetcPackageEdit, "Delete": NetDelete, "bDelete": NetbDelete, "cDelete": NetcDelete, "Welcome": NetWelcome, "bWelcome": NetbWelcome, "cWelcome": NetcWelcome, "Stripe": NetStripe, "bStripe": NetbStripe, "cStripe": NetcStripe, "Debugger": NetDebugger, "bDebugger": NetbDebugger, "cDebugger": NetcDebugger, "TemplateEdit": NetTemplateEdit, "bTemplateEdit": NetbTemplateEdit, "cTemplateEdit": NetcTemplateEdit, "TemplateEditTwo": NetTemplateEditTwo, "bTemplateEditTwo": NetbTemplateEditTwo, "cTemplateEditTwo": NetcTemplateEditTwo, "Input": NetInput, "bInput": NetbInput, "cInput": NetcInput, "DebuggerNode": NetDebuggerNode, "bDebuggerNode": NetbDebuggerNode, "cDebuggerNode": NetcDebuggerNode, "Button": NetButton, "bButton": NetbButton, "cButton": NetcButton, "Submit": NetSubmit, "bSubmit": NetbSubmit, "cSubmit": NetcSubmit, "Logo": NetLogo, "bLogo": NetbLogo, "cLogo": NetcLogo, "Navbar": NetNavbar, "bNavbar": NetbNavbar, "cNavbar": NetcNavbar, "NavCustom": NetNavCustom, "bNavCustom": NetbNavCustom, "cNavCustom": NetcNavCustom, "NavMain": NetNavMain, "bNavMain": NetbNavMain, "cNavMain": NetcNavMain, "NavPKG": NetNavPKG, "bNavPKG": NetbNavPKG, "cNavPKG": NetcNavPKG, "CrashedPage": NetCrashedPage, "bCrashedPage": NetbCrashedPage, "cCrashedPage": NetcCrashedPage, "EndpointTesting": NetEndpointTesting, "bEndpointTesting": NetbEndpointTesting, "cEndpointTesting": NetcEndpointTesting, "KanBan": NetKanBan, "bKanBan": NetbKanBan, "cKanBan": NetcKanBan, "Docker": NetDocker, "bDocker": NetbDocker, "cDocker": NetcDocker, "NavPromo": NetNavPromo, "bNavPromo": NetbNavPromo, "cNavPromo": NetcNavPromo, "FSCs": types.NetstructFSCs, "isFSCs": types.NetcastFSCs, "Dex": types.NetstructDex, "isDex": types.NetcastDex, "SoftUser": types.NetstructSoftUser, "isSoftUser": types.NetcastSoftUser, "USettings": types.NetstructUSettings, "isUSettings": types.NetcastUSettings, "App": types.NetstructApp, "isApp": types.NetcastApp, "TemplateEdits": types.NetstructTemplateEdits, "isTemplateEdits": types.NetcastTemplateEdits, "WebRootEdits": types.NetstructWebRootEdits, "isWebRootEdits": types.NetcastWebRootEdits, "TEditor": types.NetstructTEditor, "isTEditor": types.NetcastTEditor, "Navbars": types.NetstructNavbars, "isNavbars": types.NetcastNavbars, "SModal": types.NetstructSModal, "isSModal": types.NetcastSModal, "Forms": types.NetstructForms, "isForms": types.NetcastForms, "SButton": types.NetstructSButton, "isSButton": types.NetcastSButton, "STab": types.NetstructSTab, "isSTab": types.NetcastSTab, "DForm": types.NetstructDForm, "isDForm": types.NetcastDForm, "Alertbs": types.NetstructAlertbs, "isAlertbs": types.NetcastAlertbs, "Inputs": types.NetstructInputs, "isInputs": types.NetcastInputs, "Aput": types.NetstructAput, "isAput": types.NetcastAput, "RPut": types.NetstructRPut, "isRPut": types.NetcastRPut, "SSWAL": types.NetstructSSWAL, "isSSWAL": types.NetcastSSWAL, "SPackageEdit": types.NetstructSPackageEdit, "isSPackageEdit": types.NetcastSPackageEdit, "DebugObj": types.NetstructDebugObj, "isDebugObj": types.NetcastDebugObj, "DebugNode": types.NetstructDebugNode, "isDebugNode": types.NetcastDebugNode, "PkgItem": types.NetstructPkgItem, "isPkgItem": types.NetcastPkgItem, "SROC": types.NetstructSROC, "isSROC": types.NetcastSROC, "VHuf": types.NetstructVHuf, "isVHuf": types.NetcastVHuf}
	return 0
}

var FuncStored = StoreNetfn()

func LoadPage(title string) (*gosweb.Page, error) {

	if lPage, ok := WebCache.Get(title); ok {
		return &lPage, nil
	}

	var nPage = gosweb.Page{}
	if roottitle := (title == "/"); roottitle {
		webbase := "web/"
		fname := fmt.Sprintf("%s%s", webbase, "index.html")
		body, err := assets.Asset(fname)
		if err != nil {
			fname = fmt.Sprintf("%s%s", webbase, "index.tmpl")
			body, err = assets.Asset(fname)
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

	if body, err := assets.Asset(filename); err != nil {
		filename = fmt.Sprintf("web%s.html", title)

		if body, err = assets.Asset(filename); err != nil {
			filename = fmt.Sprintf("web%s", title)

			if body, err = assets.Asset(filename); err != nil {
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
