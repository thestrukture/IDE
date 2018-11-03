package types

import (
	"encoding/json"
	"log"

	"github.com/cheikhshift/db"
)

// Assert first argument as struct
// FSCs
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

// Assert first argument as struct
// Dex
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

// Assert first argument as struct
// SoftUser
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

// Assert first argument as struct
// USettings
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

// Assert first argument as struct
// App
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

// Assert first argument as struct
// TemplateEdits
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

// Assert first argument as struct
// WebRootEdits
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

// Assert first argument as struct
// TEditor
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

// Assert first argument as struct
// Navbars
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

// Assert first argument as struct
// SModal
func NetcastSModal(args ...interface{}) *SModal {

	s := SModal{}
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
func NetstructSModal() *SModal { return &SModal{} }

// Assert first argument as struct
// Forms
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

// Assert first argument as struct
// SButton
func NetcastSButton(args ...interface{}) *SButton {

	s := SButton{}
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
func NetstructSButton() *SButton { return &SButton{} }

// Assert first argument as struct
// STab
func NetcastSTab(args ...interface{}) *STab {

	s := STab{}
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
func NetstructSTab() *STab { return &STab{} }

// Assert first argument as struct
// DForm
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

// Assert first argument as struct
// Alertbs
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

// Assert first argument as struct
// Inputs
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

// Assert first argument as struct
// Aput
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

// Assert first argument as struct
// RPut
func NetcastRPut(args ...interface{}) *RPut {

	s := RPut{}
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
func NetstructRPut() *RPut { return &RPut{} }

// Assert first argument as struct
// SSWAL
func NetcastSSWAL(args ...interface{}) *SSWAL {

	s := SSWAL{}
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
func NetstructSSWAL() *SSWAL { return &SSWAL{} }

// Assert first argument as struct
// SPackageEdit
func NetcastSPackageEdit(args ...interface{}) *SPackageEdit {

	s := SPackageEdit{}
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
func NetstructSPackageEdit() *SPackageEdit { return &SPackageEdit{} }

// Assert first argument as struct
// DebugObj
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

// Assert first argument as struct
// DebugNode
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

// Assert first argument as struct
// PkgItem
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

// Assert first argument as struct
// SROC
func NetcastSROC(args ...interface{}) *SROC {

	s := SROC{}
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
func NetstructSROC() *SROC { return &SROC{} }

// Assert first argument as struct
// VHuf
func NetcastVHuf(args ...interface{}) *VHuf {

	s := VHuf{}
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
func NetstructVHuf() *VHuf { return &VHuf{} }
