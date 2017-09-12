package core

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	//"go/types"
)

const (
	StdLen  = 16
	UUIDLen = 20
)

var primitives = []string{"Bool",
	"Int",
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"Uint",
	"Uint8",
	"Uint16",
	"Uint32",
	"Uint64",
	"Uintptr",
	"Float32",
	"Float64",
	"Complex64",
	"Complex128",
	"String",
	"UnsafePointer",
	"UntypedBool",
	"UntypedInt",
	"UntypedRune",
	"UntypedFloat",
	"UntypedComplex",
	"UntypedString",
	"UntypedNil",
	"Byte",
	"Rune"}

var GOHOME = os.ExpandEnv("$GOPATH") + "/src/"
var available_methods []string
var int_methods []string
var api_methods []string
var int_mappings []string
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var StdNums = []byte("OYZ0123456789")
var DAMP = "&&"
var AMP = "&"

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//Local DEBUG tools of Gos

func IsInImports(imports []*ast.ImportSpec, unfoundvar string) bool {

	for _, v := range imports {
		ab := strings.Split(strings.Replace(v.Path.Value, "\"", "", 2), "/")
		if ab[len(ab)-1] == unfoundvar {
			return true
		}

	}

	return false
}

func isBuiltin(pkgj string) bool {

	for _, v := range primitives {
		if v == pkgj || pkgj == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func IsInSlice(qry string, slic []string) bool {
	for _, j := range slic {
		if j == qry {
			return true
		}
	}
	return false
}

func (d *gos) DeleteEnd(id string) {

		temp := []Endpoint{}
		for _, v := range d.Endpoints.Endpoints {
			if v.Id != id  {
				temp = append(temp, v)
			}
		}
		d.Endpoints.Endpoints = temp

}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CheckFile(source string) (file, line, reason string) {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.

	o, err := parser.ParseFile(fset, source, nil, parser.SpuriousErrors)
	if err != nil {
		eb := strings.Split(err.Error(), ":")
		if len(eb) > 3 {
			file = eb[0]
			line = eb[1]
			reason = eb[3]
		}
		return

	}

	if len(o.Unresolved) > 0 {
		file = "UR"
		line = "UR"
		nset := []string{}

		for _, v := range o.Unresolved {
			if !IsInImports(o.Imports, v.Name) && !IsInSlice(v.Name, nset) && !isBuiltin(v.Name) {
				nset = append(nset, v.Name)
			}
		}

		reason = strings.Join(nset, ",")

		/* if len(nset) == 0 {

			file = "COMP"
			line = "COMP"

			log,_ := RunCmdSmart("sh gobuild.sh error.go")

			probs := strings.Split(log,"\n")

			for k,m := range probs {
				if k > 0 {
					errors = append(errors, m)
				}
			}


		} */

	}

	return
}

func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, er := os.Stat(source)
		if er != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

//updates
// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// ensure dest dir does not already exist

	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		err = os.MkdirAll(dest, fi.Mode())
		if err != nil {
			return err
		}
	}

	// create dest dir

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = os.MkdirAll(dfp, fi.Mode())
			if err != nil {
				log.Println(err)
			}
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

// A struct for returning custom error messages
type CustomError struct {
	What string
}

// Returns the error message defined in What as a string
func (e *CustomError) Error() string {
	return e.What
}

func NewLen(length int) string {
	return NewLenChars(length, StdChars)
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func NewID(length int) string {
	return NewLenChars(length, StdNums)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func Process(template *gos, r string, web string, tmpl string) (local_string string) {
	// r = GOHOME + GoS Project
	arch := gosArch{}

	mathFuncs := ` 

			 func net_add(x,v float64) float64 {
					return v + x				   
			 }

			 func net_subs(x,v float64) float64 {
				   return v - x
			 }

			 func net_multiply(x,v float64) float64 {
				   return v * x
			 }

			 func net_divided(x,v float64) float64 {
				   return v/x
			 }

	`

	if template.Type == "locale" {
		local_string = `package main 
import (`

		// if template.Type == "webapp" {

		net_imports := []string{"net/http", "time", "github.com/gorilla/sessions", "bytes", "encoding/json", "fmt", "html", "html/template", "strings", "net/http/httptest", "reflect", "os", "unsafe"}
		/*
			Methods before so that we can create to correct delegate method for each object
		*/

		for _, imp := range template.Methods.Methods {
			if !contains(available_methods, imp.Name) {
				available_methods = append(available_methods, imp.Name)
			}
		}
		apiraw := ``
		for _, imp := range template.Endpoints.Endpoints {
			/*	if !contains(api_methods, imp.Method) {
					api_methods = append(api_methods, imp.Method)
				}
			*/
			apiraw += ` 
				if  r.URL.Path == "` + imp.Path + `" && r.Method == strings.ToUpper("` + imp.Type + `") { 
					` + strings.Replace(imp.Method, `&#38;`, `&`, -1) + `
					callmet = true
				}
				`

		}
		timeline := ``
		for _, imp := range template.Timers.Timers {
			/*if !contains(api_methods, imp.Method) {
				api_methods = append(api_methods,imp.Method)
			} */
			//meth := template.findMethod(imp.Method)
			timeline += `
			` + imp.Name + ` := time.NewTicker(time.` + imp.Unit + ` * ` + imp.Interval + `)
					    go func() {
					        for _ = range ` + imp.Name + `.C {
					           ` + strings.Replace(imp.Method, `&#38;`, `&`, -1) + `
					        }
					    }()
    `
		}

		//fmt.Printf("APi Methods %v\n",api_methods)
		netMa := `template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt`
		for _, imp := range available_methods {
			if !contains(api_methods, imp) && template.findMethod(imp).Keeplocal != "true" {
				netMa += `,"` + imp + `" : net_` + imp
			}
		}
		int_lok := []string{}

		for _, imp := range template.Header.Objects {
			//struct return and function

			if !contains(int_lok, imp.Name) {
				int_lok = append(int_lok, imp.Name)
				netMa += `,"` + imp.Name + `" : net_` + imp.Name
			}
		}

		for _, imp := range template.Templates.Templates {

			netMa += `,"` + imp.Name + `" : net_` + imp.Name
			netMa += `,"b` + imp.Name + `" : net_b` + imp.Name
			netMa += `,"c` + imp.Name + `" : net_c` + imp.Name
		}
		netMa += `}`

		for _, imp := range template.RootImports {
			//fmt.Println(imp)
			if !strings.Contains(imp.Src, ".gxml") {
				dir := os.ExpandEnv("$GOPATH") + "/src/" + imp.Src
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					color.Red("Package not found")
					fmt.Println("∑ Downloading Package " + imp.Src)

					RunCmdSmart("go get " + imp.Src)
				}
				if !contains(net_imports, imp.Src) {
					net_imports = append(net_imports, imp.Src)
				}
			} else {
				pathsplit := strings.Split(imp.Src, "/")
				gosName := pathsplit[len(pathsplit)-1]
				pathsplit = pathsplit[:len(pathsplit)-1]
				dir := os.ExpandEnv("$GOPATH") + "/src/" + strings.Join(pathsplit, "/")
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					color.Red("Package not found")
					fmt.Println("∑ Downloading Package " + strings.Join(pathsplit, "/"))
					RunCmdSmart("go get " + strings.Join(pathsplit, "/"))
				}
				//split and replace last section
				fmt.Println("∑ Processing XML Yåå ", pathsplit)
				xmlPackageDir := os.ExpandEnv("$GOPATH") + "/src/" + strings.Join(pathsplit, "/") + "/"
				xml_iter, _ := LoadGos(xmlPackageDir + gosName)
				if xml_iter.Package != "" {
					//copy gole with given path -
					fmt.Println("Installing Resources into project!")
					//delete prior to copy
					//	RemoveContents(r + "/" + web + "/" + xml_iter.Package)
					//	RemoveContents(r + "/" + tmpl + "/" + xml_iter.Package)

					//CopyDir(xmlPackageDir + xml_iter.Web, r + "/" + web + "/" + xml_iter.Package)
					//CopyDir(xmlPackageDir + xml_iter.Tmpl, r + "/" + tmpl + "/" + xml_iter.Package)
				} else {
					fmt.Println("∑ Error, Couldn't import your files no package name specified")
				}
			}
		}

		//	fmt.Println(template.Methods.Methods[0].Name)

		for _, imp := range net_imports {
			local_string += `
			"` + imp + `"`
		}
		local_string += `
		)
				var store = sessions.NewCookieStore([]byte("` + template.Key + `"))
				type NoStruct struct {
					/* emptystruct */
				}
				func net_sessionGet(key string,s *sessions.Session) string {
					return s.Values[key].(string)
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

				func net_sessionRemove(key string,s *sessions.Session) string {
					delete(s.Values, key)
					return ""
				}
				func net_sessionKey(key string,s *sessions.Session) bool {					
				 if _, ok := s.Values[key]; ok {
					    //do something here
				 		return true
					}

					return false
				}

				` + mathFuncs + `

				func net_sessionGetInt(key string,s *sessions.Session) interface{} {
					return s.Values[key]
				}

				func net_sessionSet(key string, value string,s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}
				func net_sessionSetInt(key string, value interface{},s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}

				func net_importcss(s string) string {
					return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
				}

				func net_importjs(s string) string {
					return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
				}



				func formval(s string, r*http.Request) string {
					return r.FormValue(s)
				}
			
				func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page) {
				     filename :=  tmpl  + ".tmpl"
				    body, err := Asset(filename)
				    session, er := store.Get(r, "session-")

				 	if er != nil {
				           session,er = store.New(r,"session-")
				    }
				    p.Session = session
				    p.R = r
				    if err != nil {
				       fmt.Print(err)
				    } else {
				    t := template.New("PageWrapper")
				    t = t.Funcs(` + netMa + `)
				    t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				    outp := new(bytes.Buffer)
				    error := t.Execute(outp, p)
				    if error != nil {
				    	  fmt.Fprintf(w, error.Error() )
				    return
				    } 

				    p.Session.Save(r, w)

				    fmt.Fprintf(w, html.UnescapeString(outp.String()) )
				    }
				}

				func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
				  return func(w http.ResponseWriter, r *http.Request) {
				  	if !apiAttempt(w,r) {
				      fn(w, r, "")
				  	}
				  }
				} 

				func mHandler(w http.ResponseWriter, r *http.Request) {
				  	
				  	if !apiAttempt(w,r) {
				      handler(w, r, "")
				  	}
				  
				} 
				func mResponse(v interface{}) string {
					data,_ := json.Marshal(&v)
					return string(data)
				}
				func apiAttempt(w http.ResponseWriter, r *http.Request) bool {
					session, er := store.Get(r, "session-")
					response := ""
					if er != nil {
						session,_ = store.New(r, "session-")
					}
					callmet := false

					` + apiraw + `

					if callmet {
						session.Save(r,w)
						if response != "" {
							//Unmarshal json
							w.Header().Set("Access-Control-Allow-Origin", "*")
							w.Header().Set("Content-Type",  "application/json")
							w.Write([]byte(response))
						}
						return true
					}
					return false
				}

				func handler(w http.ResponseWriter, r *http.Request, context string) {
				  // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
				  p,err := loadPage(r.URL.Path , context,r,w)
				  if err != nil {
				  		fmt.Println(err)
				        http.Error(w, err.Error(), http.StatusInternalServerError)
				        return
				  }

				   w.Header().Set("Cache-Control",  "public")
				  if !p.isResource {
				  		w.Header().Set("Content-Type",  "text/html")
				  		
				        renderTemplate(w, r,  "` + web + `" + r.URL.Path, p)
				  } else {
				  	  	if strings.Contains(r.URL.Path, ".css") {
				  	  		w.Header().Add("Content-Type",  "text/css")
				  	  	} else if strings.Contains(r.URL.Path, ".js") {
				  	  		w.Header().Add("Content-Type",  "application/javascript")
				  	  	} else {
				  	  	w.Header().Add("Content-Type",  http.DetectContentType(p.Body))
				  	  	}

				      w.Write(p.Body)
				  }
				}

				func loadPage(title string, servlet string,r *http.Request,w http.ResponseWriter) (*Page,error) {
				    filename :=  "` + web + `" + title + ".tmpl" 
				    body, err := Asset(filename)
				    if err != nil {
				      filename = "` + web + `" + title + ".html"
				     
				      body, err = Asset(filename)
				      if err != nil {
				         filename = "` + web + `" + title
				         body, err = Asset(filename)
				         if err != nil {
				            return nil, err
				         } else {
				          if strings.Contains(title, ".tmpl") || title == "/" {
				            return &Page{Title: title, Body: body,isResource: false,request: nil}, nil
				          } else {
				          	 return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				          }
				         
				         }
				      } else {
				         return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				      }
				    } 
				    //load custom struts
				    return &Page{Title: title, Body: body,isResource:false,request:r}, nil
				}
				func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}
				func equalz(args ...interface{}) bool {
		    	    if args[0] == args[1] {
		        	return true;
				    }
				    return false;
				 }
				 func nequalz(args ...interface{}) bool {
				    if args[0] != args[1] {
				        return true;
				    }
				    return false;
				 }

				 func netlt(x,v float64) bool {
				    if x < v {
				        return true;
				    }
				    return false;
				 }
				 func netgt(x,v float64) bool {
				    if x > v {
				        return true;
				    }
				    return false;
				 }
				 func netlte(x,v float64) bool {
				    if x <= v {
				        return true;
				    }
				    return false;
				 }
				 func netgte(x,v float64) bool {
				    if x >= v {
				        return true;
				    }
				    return false;
				 }
				 type Page struct {
					    Title string
					    Body  []byte
					    request *http.Request
					    isResource bool
					    R *http.Request
					    Session *sessions.Session
					}`
		for _, imp := range template.Variables {
			local_string += `
						var ` + imp.Name + ` ` + imp.Type
		}
		if template.Init_Func != "" {
			local_string += `
			func init(){
				` + template.Init_Func + `
			}`

		}

		//Lets Do structs
		for _, imp := range template.Header.Structs {
			if !contains(arch.objects, imp.Name) {
				fmt.Println("Processing Struct : " + imp.Name)
				arch.objects = append(arch.objects, imp.Name)
				local_string += `
			type ` + imp.Name + ` struct {`
				local_string += imp.Attributes
				local_string += `
			}`
			}
		}

		for _, imp := range template.Header.Objects {
			local_string += `
			type ` + imp.Name + ` ` + imp.Templ
		}

		//Create an object map
		for _, imp := range template.Header.Objects {
			//struct return and function
			fmt.Println("∑ Processing object :" + imp.Name)
			if !contains(available_methods, imp.Name) {
				//addcontructor
				available_methods = append(available_methods, imp.Name)
				int_methods = append(int_methods, imp.Name)
				local_string += `
				func  net_` + imp.Name + `(args ...interface{}) (d ` + imp.Templ + `){
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					return
					} else {
						d = ` + imp.Templ + `{} 
						return
					}
				}`

			}

			delegateMethods := strings.Split(imp.Methods, "\n")

			for _, im := range delegateMethods {

				if stripSpaces(im) != "" {
					fmt.Println(imp.Name + "->" + im)
					function_map := strings.Split(im, ")")

					if !contains(int_mappings, function_map[0]+imp.Templ) {
						int_mappings = append(int_mappings, function_map[0]+imp.Templ)
						funcsp := strings.Split(function_map[0], "(")
						meth := template.findMethod(stripSpaces(funcsp[0]))

						//process limits and keep local deritives
						if meth.Autoface == "" || meth.Autoface == "true" {

							/*

							 */
							procc_funcs := true
							fmt.Println()

							if meth.Limit != "" {
								if !contains(strings.Split(meth.Limit, ","), imp.Name) {
									procc_funcs = false
								}
							}

							if contains(api_methods, meth.Name) {
								procc_funcs = false
							}

							objectName := meth.Object
							if objectName == "" {
								objectName = "object"
							}
							if procc_funcs {
								if !contains(int_methods, stripSpaces(funcsp[0])) && meth.Name != "000" {
									int_methods = append(int_methods, stripSpaces(funcsp[0]))
								}
								local_string += `
					  	func  net_` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1]+`, `+objectName+` `+imp.Templ, ",") + `) ` + stripSpaces(function_map[1])
								if stripSpaces(function_map[1]) == "" {
									local_string += ` string`
								}

								local_string += ` {
									` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

								if stripSpaces(function_map[1]) == "" {
									local_string += ` 
								return ""
							`
								}
								local_string += ` 
						}`

								if meth.Keeplocal == "false" || meth.Keeplocal == "" {
									local_string += `
						func (` + objectName + ` ` + imp.Templ + `) ` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1], ",") + `) ` + stripSpaces(function_map[1])

									local_string += ` {
							` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

									local_string += `
						}`
								}
							}
						}

					}
				}
			}

			//create Unused methods methods
			fmt.Println(int_methods)
			for _, imp := range available_methods {
				if !contains(int_methods, imp) && !contains(api_methods, imp) {
					fmt.Println("Processing : " + imp)
					meth := template.findMethod(imp)
					addedit := false
					if meth.Returntype == "" {
						meth.Returntype = "string"
						addedit = true
					}
					local_string += `
						func net_` + meth.Name + `(args ...interface{}) ` + meth.Returntype + ` {
							`
					for k, nam := range strings.Split(meth.Variables, ",") {
						if nam != "" {
							local_string += nam + ` := ` + `args[` + strconv.Itoa(k) + `]
								`
						}
					}
					local_string += strings.Replace(meth.Method, `&#38;`, `&`, -1)
					if addedit {
						local_string += `
						 return ""
						 `
					}
					local_string += `
						}`
				}
			}
			for _, imp := range template.Templates.Templates {
				local_string += `
				func  net_` + imp.Name + `(args ...interface{}) string {
					var d ` + imp.Struct + `
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return ""
					}
					} else {
						d = ` + imp.Struct + `{}
					}

					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
			
				    error := t.Execute(output, &d)
				    if error != nil {
				    fmt.Print(error)
				    } 
					return html.UnescapeString(output.String())
				}`
				local_string += `
				func  net_b` + imp.Name + `(d ` + imp.Struct + `) string {
					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
			
				    error := t.Execute(output, &d)
				    if error != nil {
				    fmt.Print(error)
				    } 
					return html.UnescapeString(output.String())
				}`
				local_string += `
				func  net_c` + imp.Name + `(args ...interface{}) (d ` + imp.Struct + `) {
					if len(args) > 0 {
					var jsonBlob = []byte(args[0].(string))
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return 
					}
					} else {
						d = ` + imp.Struct + `{}
					}
    				return
				}`
			}

			//Methods have been added

			local_string += `
			func dummy_timer(){
				dg := time.Second *5
				fmt.Println(dg)
			}`

			local_string += `

			func main() {
				` + template.Main

			local_string += `
					 ` + timeline + `
					path := os.Args[1]
				var params []byte 

				
				if len(os.Args) > 3 {
				params = []byte(os.Args[3]) 
				}
				//params


				req, err := http.NewRequest(os.Args[2], "http://example.com" + path,bytes.NewReader(params) )
				if err != nil {
					fmt.Printf("error",err)
				}

				w := httptest.NewRecorder()
				mHandler(w, req)

				fmt.Printf("%s", w.Body.String())
			}`

			fmt.Println("Saving file to " + r + "/" + template.Output)
			d1 := []byte(local_string)
			_ = ioutil.WriteFile(r+"/"+template.Output, d1, 0644)

		}

	} else if template.Type == "webapp" {
		local_string = `package main 
import (`

		// if template.Type == "webapp" {
		if !template.Prod {
			template.ErrorPage = ""
			template.NPage = ""
		}

		var TraceOpt string
		if template.Debug == "on" {
			TraceOpt = `TraceTwo(2)`
		}

		net_imports := []string{"net/http", "time", "github.com/gorilla/sessions", "github.com/gorilla/context", "errors", "github.com/cheikhshift/db", "github.com/elazarl/go-bindata-assetfs", "bytes", "encoding/json", "fmt", "html", "html/template", "github.com/fatih/color", "strings", "reflect", "unsafe", "os", "bufio", "log", "io/ioutil", "runtime/trace"}
		/*
			Methods before so that we can create to correct delegate method for each object
		*/

		for _, imp := range template.Methods.Methods {
			if !contains(available_methods, imp.Name) {
				available_methods = append(available_methods, imp.Name)
			}
		}
		apiraw := ``
		for _, imp := range template.Endpoints.Endpoints {
			est := ``
			if !template.Prod {
				est = `	
					lastLine := ""
					defer func() {
					       if n := recover(); n != nil {
					          fmt.Println("Web request (` + imp.Path +`) failed at line :",GetLine("` + template.Name + `", lastLine),"Of file:` + template.Name + ` :"` + `, strings.TrimSpace(lastLine))
					          fmt.Println("Reason : ",n)
					          ` + TraceOpt + `
					          http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
						}()`
				setv := strings.Split(imp.Method, "\n")
				for _, line := range setv {
					est += `
						lastLine = ` + "`" + line + "`" + `
						` + line
				}

			} else {
				est = strings.Replace(imp.Method, `&#38;`, `&`, -1)
			}
			if imp.Type == "f" {

				apiraw += ` 
				if   strings.Contains(r.URL.Path, "` + imp.Path + `")  { 
					` + est + `
					context.Clear(r)
					` + TraceOpt + `
				}
				`
			}
		}
		for _, imp := range template.Endpoints.Endpoints {
			est := ``
			if !template.Prod {
				est = `	
					lastLine := ""
					defer func() {
					       if n := recover(); n != nil {
					          fmt.Println("Web request (` + imp.Path + `) failed at line :",GetLine("` + template.Name + `", lastLine),"Of file:` + template.Name + ` :"` + `, strings.TrimSpace(lastLine))
					          fmt.Println("Reason : ",n)
					          http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
						}()`
				setv := strings.Split(imp.Method, "\n")
				for _, line := range setv {
					est += `
						lastLine = ` + "`" + line + "`" + `
						` + line
				}

			} else {
				est = strings.Replace(imp.Method, `&#38;`, `&`, -1)
			}
			if imp.Type == "star" {

				apiraw += ` 
				if   strings.Contains(r.URL.Path, "` + imp.Path + `")  { 
					` + est + `
					
					 ` + TraceOpt + `
					callmet = true
				}
				`
			} else if imp.Type != "f" {

				apiraw += ` 
				if  r.URL.Path == "` + imp.Path + `" && r.Method == strings.ToUpper("` + imp.Type + `") { 
					` + est + `
					context.Clear(r)
					 ` + TraceOpt + `
					callmet = true
				}
				`
			}

		}
		timeline := ``
		for _, imp := range template.Timers.Timers {

			timeline += `
			` + imp.Name + ` := time.NewTicker(time.` + imp.Unit + ` * ` + imp.Interval + `)
					    go func() {
					        for _ = range ` + imp.Name + `.C {
					           ` + strings.Replace(imp.Method, `&#38;`, `&`, -1) + `
					        }
					    }()
    `
		}

		//fmt.Printf("APi Methods %v\n",api_methods)
		netMa := `template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt`
		for _, imp := range available_methods {
			if !contains(api_methods, imp) && template.findMethod(imp).Keeplocal != "true" {
				netMa += `,"` + imp + `" : net_` + imp
			}
		}
		int_lok := []string{}

		/*	for _,imp := range template.RootImports {
					//fmt.Println(imp)
				if strings.Contains(imp.Src,".gxml") {

					pathsplit := strings.Split(imp.Src,"/")
					gosName := pathsplit[len(pathsplit) - 1]
					pathsplit = pathsplit[:len(pathsplit)-1]
					if _, err := os.Stat(TrimSuffix(os.ExpandEnv("$GOPATH"), "/" ) + "/src/"  + strings.Join(pathsplit,"/")); os.IsNotExist(err){
							color.Red("Package not found")
							fmt.Println("∑ Downloading Package " + strings.Join(pathsplit,"/"))
							RunCmdSmart("go get " + strings.Join(pathsplit,"/"))
					}
					//split and replace last section
					fmt.Println("∑ Processing XML Yåå ", pathsplit)
					xmlPackageDir := TrimSuffix(os.ExpandEnv("$GOPATH"), "/" ) + "/src/" + strings.Join(pathsplit,"/") + "/"
						//copy gole with given path -
						fmt.Println("Installing Resources into project!")
						//delete prior to copy
					//	RemoveContents(r + "/" + web + "/" + xml_iter.Package)
					//	RemoveContents(r + "/" + tmpl + "/" + xml_iter.Package)
					//	CopyDir(xmlPackageDir + xml_iter.Web, r + "/" + web + "/" + xml_iter.Package)
					//	CopyDir(xmlPackageDir + xml_iter.Tmpl, r + "/" + tmpl + "/" + xml_iter.Package)
					//	template.MergeWith( xmlPackageDir + gosName)
					//	fmt.Println(template)

				}
			}
		*/
		for _, imp := range template.RootImports {
			if !strings.Contains(imp.Src, ".gxml") {
				//fmt.Println(TrimSuffix(os.ExpandEnv("$GOPATH"), "/" ) + "/src/" + imp.Src )
				if _, err := os.Stat(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + imp.Src); os.IsNotExist(err) {
					color.Red("Package not found")
					fmt.Println("∑ Downloading Package " + imp.Src)
					RunCmdSmart("go get " + imp.Src)
				}
				if !contains(net_imports, imp.Src) {
					net_imports = append(net_imports, imp.Src)
				}
			}
		}

		for _, imp := range template.Header.Objects {
			//struct return and function

			if !contains(int_lok, imp.Name) {
				int_lok = append(int_lok, imp.Name)
				netMa += `,"` + imp.Name + `" : net_` + imp.Name
			}
		}

		for _, imp := range template.Templates.Templates {

			netMa += `,"` + imp.Name + `" : net_` + imp.Name
			netMa += `,"b` + imp.Name + `" : net_b` + imp.Name
			netMa += `,"c` + imp.Name + `" : net_c` + imp.Name
		}

		//	fmt.Println(template.Methods.Methods[0].Name)

		for _, imp := range net_imports {
			local_string += `
			"` + imp + `"`
		}
		var structs_string string
		//Lets Do structs
		structs_string = ``
		for _, imp := range template.Header.Structs {
			if !contains(arch.objects, imp.Name) {
				fmt.Println("Processing Struct : " + imp.Name)
				arch.objects = append(arch.objects, imp.Name)
				structs_string += `
			type ` + imp.Name + ` struct {`
				structs_string += imp.Attributes
				structs_string += `
			}

			func  net_cast` + imp.Name + `(args ...interface{}) *` + imp.Name + `  {
				
				s := ` + imp.Name + `{}
				mapp := args[0].(db.O)
				if _, ok := mapp["_id"]; ok {
					mapp["Id"] = mapp["_id"]
				}
				data,_ := json.Marshal(&mapp)
				
				err := json.Unmarshal(data, &s) 
				if err != nil {
					fmt.Println(err.Error())
				}
				
				return &s
			}
			func net_struct` + imp.Name + `() *` + imp.Name + `{ return &` + imp.Name + `{} }`
				netMa += `,"` + imp.Name + `" : net_struct` + imp.Name
				netMa += `,"is` + imp.Name + `" : net_cast` + imp.Name

			}
		}

		netMa += `}`

		local_string += `
		)
				var store = sessions.NewCookieStore([]byte("` + template.Key + `"))

				type NoStruct struct {
					/* emptystruct */
				}

				func net_sessionGet(key string,s *sessions.Session) string {
					return s.Values[key].(string)
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

				func net_sessionRemove(key string,s *sessions.Session) string {
					delete(s.Values, key)
					return ""
				}
				func net_sessionKey(key string,s *sessions.Session) bool {					
				 if _, ok := s.Values[key]; ok {
					    //do something here
				 		return true
					}

					return false
				}

				` + mathFuncs + `

				func net_sessionGetInt(key string,s *sessions.Session) interface{} {
					return s.Values[key]
				}

				func net_sessionSet(key string, value string,s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}
				func net_sessionSetInt(key string, value interface{},s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}

				func dbDummy() {
					smap := db.O{}
					smap["key"] = "set"
					fmt.Println(smap)
				}

				func TraceTwo( sec int64) {
  
				  /*
				  	if durationExceedsWriteTimeout(r, sec) {
				  		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				  		w.Header().Set("X-Go-Pprof", "1")
				  		w.WriteHeader(http.StatusBadRequest)
				  		fmt.Fprintln(w, "profile duration exceeds server's WriteTimeout")
				  		return
				  	} */
				  
				  	// Set Content Type assuming trace.Start will work,
				  	// because if it does it starts writing.
				  	 var w bytes.Buffer
				  	if err := trace.Start(&w); err != nil {
				  		// trace.Start failed, so no writes yet.
				  		// Can change header back to text content and send error code.
				  		fmt.Println("Stack trace failed.")
				  		return
				  	}
				  	
				  	trace.Stop()
				  	ioutil.WriteFile("__heap", w.Bytes(), 0777)
				  	//fmt.Println(w.String())
				  }
				func Trace( nam string) {
  
				  /*
				  	if durationExceedsWriteTimeout(r, sec) {
				  		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				  		w.Header().Set("X-Go-Pprof", "1")
				  		w.WriteHeader(http.StatusBadRequest)
				  		fmt.Fprintln(w, "profile duration exceeds server's WriteTimeout")
				  		return
				  	} */
				  
				  	// Set Content Type assuming trace.Start will work,
				  	// because if it does it starts writing.
				  	 var w bytes.Buffer
				  	if err := trace.Start(&w); err != nil {
				  		// trace.Start failed, so no writes yet.
				  		// Can change header back to text content and send error code.
				  		fmt.Println("Stack trace failed.")
				  		return
				  	}
				  	
				  	trace.Stop()
				  	ioutil.WriteFile(nam, w.Bytes(), 0777)
				  	//fmt.Println(w.String())
				  }

				func net_importcss(s string) string {
					return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
				}

				func net_importjs(s string) string {
					return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
				}



				func formval(s string, r*http.Request) string {
					return r.FormValue(s)
				}
			
				func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page)  bool {
				     defer func() {
					        if n := recover(); n != nil {
					           	 color.Red("Error loading template in path : ` + web + `" + r.URL.Path + ".tmpl reason :" )
					           	 fmt.Println(n)
					           	 DebugTemplate( w,r ,"` + web + `" + r.URL.Path)
					           	 http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
					    }()

				    filename :=  tmpl  + ".tmpl"
				    body, err := Asset(filename)
				    session, er := store.Get(r, "session-")

				 	if er != nil {
				           session,er = store.New(r,"session-")
				    }
				    p.Session = session
				    p.R = r
				    if err != nil {
				      // fmt.Print(err)
				    	return false
				    } else {
				    t := template.New("PageWrapper")
				    t = t.Funcs(` + netMa + `)
				    t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				    outp := new(bytes.Buffer)
				    error := t.Execute(outp, p)
				    if error != nil {
				    fmt.Println(error.Error())
				    	 DebugTemplate( w,r ,"` + web + `" + r.URL.Path)
				    	 http.Redirect(w,r,"` + template.ErrorPage + `",301)
				    return false
				    }  else {
				    p.Session.Save(r, w)

				    fmt.Fprintf(w, html.UnescapeString(outp.String()) )
				    ` + TraceOpt + `
				    return true
					}
				    }
				}

				func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
				  return func(w http.ResponseWriter, r *http.Request) {
				  	if !apiAttempt(w,r) {
				      fn(w, r, "")
				  	}
				  	
				  }
				} 

				func mHandler(w http.ResponseWriter, r *http.Request) {
				  	
				  	if !apiAttempt(w,r) {
				      handler(w, r, "")
				  	}
				  
				} 
				func mResponse(v interface{}) string {
					data,_ := json.Marshal(&v)
					return string(data)
				}
				func apiAttempt(w http.ResponseWriter, r *http.Request) bool {
					session, er := store.Get(r, "session-")
					response := ""
					if er != nil {
						session,_ = store.New(r, "session-")
					}
					callmet := false

					` + apiraw + `

					if callmet {
						session.Save(r,w)
						if response != "" {
							//Unmarshal json
							w.Header().Set("Access-Control-Allow-Origin", "*")
							w.Header().Set("Content-Type",  "application/json")
							w.Write([]byte(response))
						}
						return true
					}
					return false
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
				func DebugTemplate(w http.ResponseWriter,r *http.Request,tmpl string){
					lastline := 0
					linestring := ""
					defer func() {
					       if n := recover(); n != nil {
					           	fmt.Println()
					           	// fmt.Println(n)
					           			fmt.Println("Error on line :", lastline ,":" + strings.TrimSpace(linestring)) 
					           	 //http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
					    }()	

					p,err := loadPage(r.URL.Path , "",r,w)
					filename :=  tmpl  + ".tmpl"
				    body, err := Asset(filename)
				    session, er := store.Get(r, "session-")

				 	if er != nil {
				           session,er = store.New(r,"session-")
				    }
				    p.Session = session
				    p.R = r
				    if err != nil {
				       	fmt.Print(err)
				    	
				    } else {
				    
				  
				   
				    lines := strings.Split(string(body), "\n")
				   // fmt.Println( lines )
				    linebuffer := ""
				    waitend := false
				    open := 0
				    for i, line := range lines {
				    	
				    	

				   

				    	if waitend {
				    		linebuffer += line

				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(` + netMa + `)
					    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(linebuffer + endstr, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
					    	lastline = i
					    	linestring =  line
					    	error := t.Execute(outp, p)
						    if error != nil {
						   		fmt.Println("Error on line :", i + 1,line,error.Error())   
						    } 

				    	}

				    if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				    		linebuffer += line
				    		waitend = true
				    		open++;
				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(` + netMa + `)
					    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(linebuffer + endstr, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
					    	lastline = i
					    	linestring =  line
					    	error := t.Execute(outp, p)
						    if error != nil {
						   		fmt.Println("Error on line :", i + 1,line,error.Error())   
						    } 
				    	}

				    	if !waitend {
				    	outp := new(bytes.Buffer)  
				    	t := template.New("PageWrapper")
				    	t = t.Funcs(` + netMa + `)
				    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(line, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				    	lastline = i
				    	linestring = line
				    	error := t.Execute(outp, p)
					    if error != nil {
					   		fmt.Println("Error on line :", i + 1,line,error.Error())   
					    }  
						}

						if  strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") {
							open--

							if open == 0 {
							waitend = false
				    		
							}
				    	}
				    }
				    
					
				    }

				}

			func DebugTemplatePath(tmpl string, intrf interface{}){
					lastline := 0
					linestring := ""
					defer func() {
					       if n := recover(); n != nil {
					         
					           			fmt.Println("Error on line :", lastline + 1,":" + strings.TrimSpace(linestring)) 
					           			fmt.Println(n)
					           	 //http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
					    }()	

				
					filename :=  tmpl  
				    body, err := Asset(filename)
				   
				    if err != nil {
				       	fmt.Print(err)
				    	
				    } else {
				    
				  
				   
				    lines := strings.Split(string(body), "\n")
				   // fmt.Println( lines )
				    linebuffer := ""
				    waitend := false
				    open := 0
				    for i, line := range lines {
				    	
				    	

				   

				    	if waitend {
				    		linebuffer += line

				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(` + netMa + `)
					    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(linebuffer + endstr, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
					    	lastline = i
					    	linestring =  line
					    	error := t.Execute(outp, intrf)
						    if error != nil {
						   		fmt.Println("Error on line :", i + 1,line,error.Error())   
						    } 

				    	}

				    	if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				    		linebuffer += line
				    		waitend = true
				    		open++;
				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(` + netMa + `)
					    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(linebuffer + endstr, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
					    	lastline = i
					    	linestring =  line
					    	error := t.Execute(outp, intrf)
						    if error != nil {
						   		fmt.Println("Error on line :", i + 1,line,error.Error())   
						    } 
				    	}

				    	if !waitend {
				    	outp := new(bytes.Buffer)  
				    	t := template.New("PageWrapper")
				    	t = t.Funcs(` + netMa + `)
				    	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(line, "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				    	lastline = i
				    	linestring = line
				    	error := t.Execute(outp, intrf)
					    if error != nil {
					   		fmt.Println("Error on line :", i + 1,line,error.Error())   
					    }  
						}

						if  strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}"){
							open--

							if open == 0 {
							waitend = false
				    		
							}
				    	}
				    }
				    
					
				    }

				}
			func handler(w http.ResponseWriter, r *http.Request, contxt string) {
				  // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
				  p,err := loadPage(r.URL.Path , contxt,r,w)
				  if err != nil {
				  	fmt.Println(err)
				  	` + TraceOpt + `
				        http.Redirect(w,r,"` + template.NPage + `",307)
				        context.Clear(r)
				        return
				  }

				   w.Header().Set("Cache-Control",  "public")
				  if !p.isResource {
				  		w.Header().Set("Content-Type",  "text/html")
				  		    defer func() {
					       if n := recover(); n != nil {
					           	 color.Red("Error loading template in path : ` + web + `" + r.URL.Path + ".tmpl reason :" )
					           	 fmt.Println(n)
					           	 DebugTemplate( w,r ,"` + web + `" + r.URL.Path)
					           	 http.Redirect(w,r,"` + template.ErrorPage + `",307)
					           	 context.Clear(r)
					           	 ` + TraceOpt + `
					        }
					    }()	
				      	renderTemplate(w, r,  "` + web + `" + r.URL.Path, p)
				     
				     // fmt.Println(w)
				  } else {
				  		if strings.Contains(r.URL.Path, ".css") {
				  	  		w.Header().Add("Content-Type",  "text/css")
				  	  	} else if strings.Contains(r.URL.Path, ".js") {
				  	  		w.Header().Add("Content-Type",  "application/javascript")
				  	  	} else {
				  	  	w.Header().Add("Content-Type",  http.DetectContentType(p.Body))
				  	  	}
				  	 
				  	 
				      w.Write(p.Body)
				  }

				  context.Clear(r)
				  
				}

				func loadPage(title string, servlet string,r *http.Request,w http.ResponseWriter) (*Page,error) {
				    filename :=  "` + web + `" + title + ".tmpl"
				    if title == "/" {
				      http.Redirect(w,r,"/index",302)
				    }
				    body, err := Asset(filename)
				    if err != nil {
				      filename = "` + web + `" + title + ".html"
				      if title == "/" {
				    	filename = "` + web + `/index.html"
				    	}
				      body, err = Asset(filename)
				      if err != nil {
				         filename = "` + web + `" + title
				         body, err = Asset(filename)
				         if err != nil {
				            return nil, err
				         } else {
				          if strings.Contains(title, ".tmpl") || title == "/" {
				              return nil,nil
				          }
				          return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				         }
				      } else {
				         return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				      }
				    } 
				    //load custom struts
				    ` + TraceOpt + `
				    return &Page{Title: title, Body: body,isResource:false,request:r}, nil
				}
				func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}
				func equalz(args ...interface{}) bool {
		    	    if args[0] == args[1] {
		        	return true;
				    }
				    return false;
				 }
				 func nequalz(args ...interface{}) bool {
				    if args[0] != args[1] {
				        return true;
				    }
				    return false;
				 }

				 func netlt(x,v float64) bool {
				    if x < v {
				        return true;
				    }
				    return false;
				 }
				 func netgt(x,v float64) bool {
				    if x > v {
				        return true;
				    }
				    return false;
				 }
				 func netlte(x,v float64) bool {
				    if x <= v {
				        return true;
				    }
				    return false;
				 }

				 func GetLine(fname string , match string )  int {
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
									if strings.Contains(scanner.Text(), match ) {
												    		
												    		return intx
												    	}

								}


					return -1
				}
				 func netgte(x,v float64) bool {
				    if x >= v {
				        return true;
				    }
				    return false;
				 }
				 type Page struct {
					    Title string
					    Body  []byte
					    request *http.Request
					    isResource bool
					    R *http.Request
					    Session *sessions.Session
					}`
		for _, imp := range template.Variables {
			local_string += `
						var ` + imp.Name + ` ` + imp.Type
		}
		if template.Init_Func != "" {
			local_string += `
			func init(){
				` + template.Init_Func + `
			}`

		}

		local_string += structs_string

		for _, imp := range template.Header.Objects {
			local_string += `
			type ` + imp.Name + ` ` + imp.Templ
		}

		//Create an object map
		for _, imp := range template.Header.Objects {
			//struct return and function
			fmt.Println("∑ Processing object :" + imp.Name)
			if !contains(available_methods, imp.Name) {
				//addcontructor
				if imp.Templ == "" {
					imp.Templ = "NoStruct"
				}

				available_methods = append(available_methods, imp.Name)
				int_methods = append(int_methods, imp.Name)
				local_string += `
				func  net_` + imp.Name + `(args ...interface{}) (d ` + imp.Templ + `){
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					return
					} else {
						d = ` + imp.Templ + `{} 
						return
					}
				}`

			}

			delegateMethods := strings.Split(imp.Methods, "\n")

			for _, im := range delegateMethods {

				if stripSpaces(im) != "" {
					fmt.Println(imp.Name + "->" + im)
					function_map := strings.Split(im, ")")

					if !contains(int_mappings, function_map[0]+imp.Templ) {
						int_mappings = append(int_mappings, function_map[0]+imp.Templ)
						funcsp := strings.Split(function_map[0], "(")
						meth := template.findMethod(stripSpaces(funcsp[0]))

						//process limits and keep local deritives
						if meth.Autoface == "" || meth.Autoface == "true" {

							/*

							 */
							procc_funcs := true
							fmt.Println()

							if meth.Limit != "" {
								if !contains(strings.Split(meth.Limit, ","), imp.Name) {
									procc_funcs = false
								}
							}

							if contains(api_methods, meth.Name) {
								procc_funcs = false
							}

							objectName := meth.Object
							if objectName == "" {
								objectName = "object"
							}
							if procc_funcs {
								if !contains(int_methods, stripSpaces(funcsp[0])) && meth.Name != "000" {
									int_methods = append(int_methods, stripSpaces(funcsp[0]))
								}
								local_string += `
					  	func  net_` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1]+`, `+objectName+` `+imp.Templ, ",") + `) ` + stripSpaces(function_map[1])
								if stripSpaces(function_map[1]) == "" {
									local_string += ` string`
								}

								local_string += ` {
									` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

								if stripSpaces(function_map[1]) == "" {
									local_string += ` 
								return ""
							`
								}
								local_string += ` 
						}`

								if meth.Keeplocal == "false" || meth.Keeplocal == "" {
									local_string += `
						func (` + objectName + ` ` + imp.Templ + `) ` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1], ",") + `) ` + stripSpaces(function_map[1])

									local_string += ` {
							` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

									local_string += `
						}`
								}
							}
						}

					}
				}
			}

			//create Unused methods methods
		//	fmt.Println(int_methods)
		
		

			


		}

			for _, imp := range available_methods {
				if !contains(int_methods, imp) && !contains(api_methods, imp) {
					fmt.Println("Processing : " + imp)
					meth := template.findMethod(imp)
					addedit := false
					if meth.Returntype == "" {
						meth.Returntype = "string"
						addedit = true
					}
					local_string += `
						func net_` + meth.Name + `(args ...interface{}) ` + meth.Returntype + ` {
							`
					for k, nam := range strings.Split(meth.Variables, ",") {
						if nam != "" {
							local_string += nam + ` := ` + `args[` + strconv.Itoa(k) + `]
								`
						}
					}

					est := ``
					if !template.Prod {
						est = `	
							lastLine := ""
							defer func() {
							       if n := recover(); n != nil {
							          fmt.Println("Pipeline failed at line :",GetLine("` + template.Name + `", lastLine),"Of file:` + template.Name + ` :"` + `, strings.TrimSpace(lastLine))
							          fmt.Println("Reason : ",n)
							         
							        }
								}()`
						setv := strings.Split(meth.Method, "\n")
						for _, line := range setv {
							if len(line) > 4 {
								est += `
								lastLine = ` + "`" + line + "`" + `
								` + line
							}
						}

					} else {
						est = strings.Replace(meth.Method, `&#38;`, `&`, -1)
					}
					local_string += est
					if addedit {
						local_string += `
						 return ""
						 `
					}
					local_string += `
						}`
				}
			}

			for _, imp := range template.Templates.Templates {
				if imp.Struct == "" {
					imp.Struct = "NoStruct"
				}
				local_string += `


				func  net_` + imp.Name + `(args ...interface{}) string {
					var d ` + imp.Struct + `
					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
						defer func() {
					       if n := recover(); n != nil {
					           	   color.Red("Error loading template in path (` + imp.Name + `) : " + filename )
					           	// fmt.Println(n)
					           		DebugTemplatePath(filename, &d)	
					           	 //http.Redirect(w,r,"` + template.ErrorPage + `",307)
					        }
					    }()	
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return ""
					}
					} else {
						d = ` + imp.Struct + `{}
					}

					
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
					
					
				    error := t.Execute(output, &d)
				    if error != nil {
				   color.Red("Error processing template " + filename)
				   DebugTemplatePath(filename, &d)	
				    } 
					return html.UnescapeString(output.String())
					
				}`
				local_string += `
					func  b` + imp.Name + `(d ` + imp.Struct + `) string {
						return net_b` + imp.Name + `(d)
					}

				func  net_b` + imp.Name + `(d ` + imp.Struct + `) string {
					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
					
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				 defer func() {
					        if n := recover(); n != nil {
					           	color.Red("Error loading template in path (` + imp.Name + `) : " + filename )
					           	DebugTemplatePath(filename, &d)	
					        }
					    }()
				    error := t.Execute(output, &d)
				    if error != nil {
				    fmt.Print(error)
				    } 
					return html.UnescapeString(output.String())
				}`
				local_string += `
				func  net_c` + imp.Name + `(args ...interface{}) (d ` + imp.Struct + `) {
					if len(args) > 0 {
					var jsonBlob = []byte(args[0].(string))
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return 
					}
					} else {
						d = ` + imp.Struct + `{}
					}
    				return
				}

				func  c` + imp.Name + `(args ...interface{}) (d ` + imp.Struct + `) {
					if len(args) > 0 {
					var jsonBlob = []byte(args[0].(string))
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return 
					}
					} else {
						d = ` + imp.Struct + `{}
					}
    				return
				}`
			}

			//Methods have been added

			local_string += `
			func dummy_timer(){
				dg := time.Second *5
				fmt.Println(dg)
			}`

			local_string += `

			func main() {
				fmt.Printf("%d\n", os.Getpid())
				` + template.Main

			local_string += ` 
					 ` + timeline + `
					 fmt.Printf("Listenning on Port %v\n", "` + template.Port + `")
					 http.HandleFunc( "/",  makeHandler(handler))
					 store.Options = &sessions.Options{
						    Path:     "/",
						    MaxAge:   86400 * 7,
						    HttpOnly: true,
						    Secure : true,
						    Domain : "` + template.Domain + `",
						}
					 http.Handle("/dist/",  http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "` + web + `"}))
					errgos := http.ListenAndServe(":` + template.Port + `", nil)
					if errgos != nil {
						log.Fatal(errgos)
					} 

					}`

			fmt.Println("Saving file to " + r + "/" + template.Output)
			d1 := []byte(local_string)
			_ = ioutil.WriteFile(r+"/"+template.Output, d1, 0644)

	} else if template.Type == "bind" {
		local_string = `package ` + template.Package + ` 
				
	import (`

		net_imports := []string{"time", "os", "bytes", "encoding/json", "fmt", "html", "html/template", "io/ioutil", "strings", "reflect", "unsafe", "crypto/aes", "crypto/cipher", "crypto/rand", "io", "encoding/base64", "errors"}

		/*
			Methods before so that we can create to correct delegate method for each object
		*/

		for _, imp := range template.Methods.Methods {
			if !contains(available_methods, imp.Name) {
				available_methods = append(available_methods, imp.Name)
			}
		}
		apiraw := ``
		for _, imp := range template.Endpoints.Endpoints {
			if !contains(api_methods, imp.Method) {
				api_methods = append(api_methods, imp.Method)
			}
			meth := template.findMethod(imp.Method)
			apiraw += ` 
				if  path == "` + imp.Path + `" && method == strings.ToUpper("` + imp.Type + `") { 
					` + strings.Replace(meth.Method, `&#38;`, `&`, -1) + `
					callmet = true
				}
				`

		}

		fmt.Printf("APi Methods %v\n", api_methods)
		netMa := `template.FuncMap{"GetLocation": net_supportGetLocation,"Run": net_supportRunjs,"PlaySound" : net_supportSoundPlay,"StopSound" : net_supportSoundStop,"SetVolume" : net_supportSoundSetVolume,"GetVolume" : net_supportSoundGetVolume, "isPlaying" : net_supportSoundisPlaying,"trackMotion": net_supportMotionStart,"stopMotion" : net_supportMotionStop,"ShowLoad" : net_supportShowload, "HideLoad" : net_supportHideLoad, "Device": net_supportDevice,"TakePicture" : net_supportTakePicture, "Notify" : net_supportNotify,"AbsolutePath" : net_supportFileAbsPath,"Download" : net_supportFileDownload,"Download_lg" : net_supportFileDownloadLarge,"Base64" : net_supportBase64,"DeleteRes" : net_supportDeleteFile , "Height":net_layerHeight,"Width": net_layerWidth,"push":net_pushView,"dismiss":net_dismissView,"dismissAt": net_dismissViewatInt,"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sDelete" : deleteSession,"sRemove" : net_RemoveSessionKey,"sExist": net_SessionKeyExists,"sSet" : net_SetSessionKey,"sSetField": net_SetSessionField,"sGet" : net_GetSession,"sGetString" : net_GetSessionString, "sGetN" : net_GetSessionFloat,"Get" : paramGet,"eq": equalz, "neq" : nequalz, "lte" : netlt`
		for _, imp := range available_methods {
			if !contains(api_methods, imp) && template.findMethod(imp).Keeplocal != "true" {
				netMa += `,"` + imp + `" : net_` + imp
			}
		}
		int_lok := []string{}

		for _, imp := range template.Header.Objects {
			//struct return and function

			if !contains(int_lok, imp.Name) {
				int_lok = append(int_lok, imp.Name)
				netMa += `,"` + imp.Name + `" : net_` + imp.Name
			}
		}

		for _, imp := range template.Header.Structs {
			netMa += `,"is` + imp.Name + `":net_is` + imp.Name
		}

		for _, imp := range template.Templates.Templates {

			netMa += `,"` + imp.Name + `" : net_` + imp.Name
			netMa += `,"b` + imp.Name + `" : b` + imp.Name
			netMa += `,"c` + imp.Name + `" : c` + imp.Name
		}
		netMa += `}`

		for _, imp := range template.RootImports {
			//fmt.Println(imp)
			if !strings.Contains(imp.Src, ".gxml") {
				if imp.Download == "true" {
					fmt.Println("∑ Downloading Package " + imp.Src)
					RunCmdSmart("go get " + imp.Src)
				}
				if !contains(net_imports, imp.Src) {
					net_imports = append(net_imports, imp.Src)
				}
			}
		}

		//fmt.Println(template.Methods.Methods[0].Name)

		for _, imp := range net_imports {
			local_string += `
			"` + imp + `"`
		}
		local_string += `
		)

                type Flow interface {
         			PushView(url string)
         			DismissView()
         			DismissViewatInt(index int)
         			Width() float64
         			Height() float64
         			Device() int
         			ShowLoad()
         			HideLoad()
         			RunJS(line string)

         			Play(path string)
        			Stop()
        			SetVolume(power int)
        			GetVolume() int
        			IsPlaying() bool
        			PlayFromWebRoot(path string)

        			RequestLocation()
        			TrackMotion()
        			StopMotion()

        			CreatePictureNamed(name string)
        			OpenAppLink(url string)

    
         			Notify(title string,message string)

         			AbsolutePath(file string) string
         			Download(url string, target string) bool
         			DownloadLarge(url string, target string)
         			Base64String(target string) string
         			GetBytes(target string) []byte
         			GetBytesFromUrl(target string) []byte
         			DeleteDirectory(path string) bool
         			DeleteFile(path string) bool
         			
         		}

         		


         		func net_supportGetLocation(flow Flow) string {
         			flow.RequestLocation()
         			return ""
         		}

         		func net_supportRunjs(jss string,flow Flow) string {
         			flow.RunJS(jss)
         			return ""
         		}

         		// sound funcs 

         		func net_supportSoundPlay(file string,flow Flow) string {
         			flow.PlayFromWebRoot(file)
         			return ""
         		}

         		func net_supportSoundStop(flow Flow) string {
         			flow.Stop()
         			return ""
         		}

         		func net_supportSoundSetVolume(level int, flow Flow) string {
         			flow.SetVolume(level)
         			return ""
         		}

         		func net_supportSoundGetVolume(flow Flow) int {
         			return flow.GetVolume()
         		}

         		func net_supportSoundisPlaying(flow Flow) bool {
         			return flow.IsPlaying()
         		}

         		// end sound funcs 

         		func net_supportMotionStart(flow Flow) string {
         			flow.TrackMotion()
         			return ""
         		}

         		func net_supportMotionStop(flow Flow) string {
         			flow.StopMotion()
         			return ""
         		}

         		func net_supportDevice(flow Flow) int {
         			return flow.Device()
         		}

         		func net_supportShowload(flow Flow) string {
         			flow.ShowLoad()
         			return ""
         		}

         		func net_supportHideLoad(flow Flow) string {
         			flow.HideLoad()
         			return ""
         		}

         		func net_supportTakePicture(pic string,flow Flow) string {
         			flow.CreatePictureNamed(pic)
         			return ""
         		}

         		func net_supportNotify(title string,message string,flow Flow) string {
         			flow.Notify(title,message)
         			return ""
         		}

         		// start file manager 
     

         		func net_supportFileAbsPath(path string, file Flow) string {
         			return file.AbsolutePath(path)
         		}

         		func net_supportFileDownload(url string,target string, file Flow) bool {
         			return file.Download(url,target);
         		}

         		func net_supportFileDownloadLarge(url string, target string, file Flow) string {
         			file.DownloadLarge(url, target)
         			return ""
         		}

         		func net_supportBase64(path string,file Flow) string {
         			return file.Base64String(path)
         		}

         		func net_supportGetBytes(target string, file Flow) []byte {
         			return file.GetBytes(target)
         		}

         		func net_supportGetBytesFromUrl(target string, file Flow) []byte {
         			return file.GetBytesFromUrl(target)
         		}

         		func net_supportDeleteFolder(path string,file Flow) bool {
         			return file.DeleteDirectory(path)
         		}

         		func net_supportDeleteFile(path string,file Flow) bool {
         			return file.DeleteFile(path)
         		}


         		// End file manager 
         		func net_pushView(url string,flow Flow) string {
         			flow.PushView(url)
         			return ""
         		}

         		func net_dismissView(flow Flow) string {
         			flow.DismissView()
         			return ""
         		}

         		func net_layerWidth(flow Flow) float64 {
         			return flow.Width()
         		}
         		func net_layerHeight(flow Flow) float64 {
         			return flow.Height()
         		}

         		func net_dismissViewatInt(ind int,flow Flow) string {
         			flow.DismissViewatInt(ind)
         			return ""
         		}
				
				var key = []byte("` + template.Key + `")

				func net_importcss(s string) string {
					return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
				}

				func net_importjs(s string) string {
					return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
				}

			
					 type page struct {
					    Title string
					    Body  []byte
					 	Parameters map[string]interface{}
					 	Session session
					 	Layer Flow
					    isResource bool
					}

				type session struct {
					Values map[string]interface{}
					//custom props
					` + template.Session + `
				
				}

				func paramGet(ke string,f map[string]interface{}) string {
					if _, ok := f[ke]; ok {
					return f[ke].(string)
					} else {
						return ""
					}
				}
			
				func dummy_timer(){
					dg := time.Second *5
					fmt.Println(dg)
				}

				func LoadUrl(path string,bod []byte,method string,flow Flow)[]byte { 
								
				body := new(bytes.Buffer)
				body.Write(bod)
				var f interface{}
				if bod != nil {
				_ = json.Unmarshal(bod, &f)
				}
				data,proceed := apiAttempt(path,method,bod,flow)				
				if proceed {
					return data
				} else {

								 p,err := loadPage(path)
								  if err != nil {
								  	fmt.Println(err)
								        return []byte("Error ")
								  }

								  if !p.isResource {
								      p.Parameters = f.(map[string]interface{}) 
								      p.Session = openSession()
								      p.Layer = flow
								      return   []byte(html.UnescapeString(string(renderTemplate("web" + path, p))))
								  } else {
								       return p.Body
								  }

					return bod
				}
								 
				}

				func net_SetSessionField(key string, arg interface{}) string {
					s := openSessionMap()
					s[key] = arg
					keepSessionMap(s)
					return ""
				}
				func net_SetSessionKey(key string, arg interface{}) string {
					s := openSession()
					s.Values[key] = arg
					keepSession(s)
					return ""
				}

				func net_SessionKeyExists(key string) bool {
					s := openSession()
					 if _, ok := s.Values[key]; ok {
					    //do something here
				 		return true
					}

					return false
				}

				` + mathFuncs + `

				func net_GetSession(key string) interface{} {
					s := openSession() 
					return s.Values[key]
				}
				func net_GetSessionString(key string) string {
					s := openSession() 
					if _, ok := s.Values[key]; ok {
					return s.Values[key].(string)
					} else {
						return ""
					}
				}
				func net_GetSessionFloat(key string) float64 {
					s := openSession() 
					if _, ok := s.Values[key]; ok {
					return s.Values[key].(float64)
					} else {
						return 0
					}
				}

				func net_RemoveSessionField(key string) string {
					s := openSessionMap()
					delete(s,key)
					//save here
					keepSessionMap(s)
					return ""
				}

				func net_RemoveSessionKey(key string) string {
					s := openSession()
					delete(s.Values,key)
					//save here
					keepSession(s)
					return ""
				}

				func deleteSession() string {
					os.Remove(os.TempDir() + "/session")
					return ""
				}

				func encrypt(text []byte) ([]byte, error) {
				    block, err := aes.NewCipher(key)
				    if err != nil {
				        return nil, err
				    }
				    b := base64.StdEncoding.EncodeToString(text)
				    ciphertext := make([]byte, aes.BlockSize+len(b))
				    iv := ciphertext[:aes.BlockSize]
				    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
				        return nil, err
				    }
				    cfb := cipher.NewCFBEncrypter(block, iv)
				    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
				    return ciphertext, nil
				}

				func decrypt(text []byte) ([]byte, error) {
				    block, err := aes.NewCipher(key)
				    if err != nil {
				        return nil, err
				    }
				    if len(text) < aes.BlockSize {
				        return nil, errors.New("ciphertext too short")
				    }
				    iv := text[:aes.BlockSize]
				    text = text[aes.BlockSize:]
				    cfb := cipher.NewCFBDecrypter(block, iv)
				    cfb.XORKeyStream(text, text)
				    data, err := base64.StdEncoding.DecodeString(string(text))
				    if err != nil {
				        return nil, err
				    }
				    return data, nil
				}


				func openSession() session {
				  body, err := ioutil.ReadFile(os.TempDir() + "/session")
    				if err != nil {
    						s := session{Values:make(map[string]interface{})}
    						return s
    				}
    				var d session
    				data,_ := decrypt(body)
    				err = json.Unmarshal(data, &d)
					if err != nil {
						fmt.Println("error:", err)
						return session{}
					}
					return d
				}

				func openSessionMap() map[string]interface{} {
				  body, err := ioutil.ReadFile(os.TempDir() + "/session")
    				if err != nil {
    						s := make(map[string]interface{})
    						return s
    				}
    				var d interface{}
    				data,_ := decrypt(body)
    				err = json.Unmarshal(data, &d)
					if err != nil {
						fmt.Println("error:", err)
						return make(map[string]interface{})
					}
					return d.(map[string]interface{})
				}

			

				
				func keepSession(s session){
				
					data,er := encrypt([]byte(mResponse(s)))
					if er != nil {
						fmt.Println(er)
						return
					}
					err := ioutil.WriteFile(os.TempDir() + "/session", data,0644)
					if err != nil {
						fmt.Println(err)
					}
				}

					func keepSessionMap(s interface{}){
					fmt.Println(mResponse(s))
					data,er := encrypt([]byte(mResponse(s)))
					if er != nil {
						fmt.Println(er)
						return
					}
					err := ioutil.WriteFile(os.TempDir() + "/session", data,0644)
					if err != nil {
						fmt.Println(err)
					}
				}


				func renderTemplate(tmpl string, f*page) []byte {
				   filename :=  tmpl  + ".tmpl"
				   body, err := Asset(filename)
				   outp := new(bytes.Buffer)
				    if err != nil {
				       fmt.Print(err)
				    } else {
				    t := template.New("PageWrapper")
				    t = t.Funcs(` + netMa + `)
				      t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
				   
				    error := t.Execute(outp, f)
				    if error != nil {
				    fmt.Print(error)
				    return nil
				    } 

				   return outp.Bytes()

				    
				    }
				    return outp.Bytes()
				}

				func loadPage(title string) (*page,error) {
				    filename :=  "web" + title + ".tmpl"
				    if title == "/" {
				    	filename = "web/index.tmpl"
				    }
				    body, err := Asset(filename)
				    if err != nil {
				      filename = "web" + title + ".html"
				      if title == "/" {
				    	filename = "web/index.html"
				    	}
				      body, err = Asset(filename)
				      if err != nil {
				         filename = "web" + title
				         body, err = Asset(filename)
				         if err != nil {
				            return nil, err
				         } else {
				          if strings.Contains(title, ".tmpl") || title == "/" {
				              return nil,nil
				          }
				          return &page{Title: title, Body: body,isResource: true}, nil
				         }
				      } else {
				         return &page{Title: title, Body: body,isResource: true}, nil
				      }
				    } 
				    //load custom struts
				    return &page{Title: title, Body: body,isResource:false}, nil
				}
				func apiAttempt(path string, method string,bod []byte,layer Flow) ([]byte,bool) {
				//	session, er := store.Get(r, "session-")
					response := ""
					session := openSession()
					callmet := false
					var f interface{}
					if bod != nil {
					_ = json.Unmarshal(bod, &f)
					}

					` + apiraw + `
				

					if callmet {
						keepSession(session)
						
						if response != "" {
							
							return []byte(response),true
						}
					
					}
					return []byte(""),false
				} 


			
				func mResponse(v interface{}) string {
					data,_ := json.Marshal(&v)
					return string(data)
				}
				
			

			
				func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}
				func equalz(args ...interface{}) bool {
		    	    if args[0] == args[1] {
		        	return true;
				    }
				    return false;
				 }
				 func nequalz(args ...interface{}) bool {
				    if args[0] != args[1] {
				        return true;
				    }
				    return false;
				 }

				 func netlt(x,v float64) bool {
				    if x < v {
				        return true;
				    }
				    return false;
				 }
				 func netgt(x,v float64) bool {
				    if x > v {
				        return true;
				    }
				    return false;
				 }
				 func netlte(x,v float64) bool {
				    if x <= v {
				        return true;
				    }
				    return false;
				 }
				 func netgte(x,v float64) bool {
				    if x >= v {
				        return true;
				    }
				    return false;
				 }
				`
		for _, imp := range template.Variables {
			local_string += `
						var ` + imp.Name + ` ` + imp.Type
		}
		if template.Init_Func != "" {
			local_string += `
			func init(){
				` + template.Init_Func + `
			}`

		}

		//Lets Do structs
		for _, imp := range template.Header.Structs {
			if !contains(arch.objects, imp.Name) {
				fmt.Println("Processing Struct : " + imp.Name)
				arch.objects = append(arch.objects, imp.Name)
				local_string += `
			type ` + imp.Name + ` struct {`
				local_string += imp.Attributes
				local_string += `
			}`

				local_string += `
			func net_is` + imp.Name + ` (arg interface{}) ` + imp.Name + ` {`
				local_string += `
				return arg.(` + imp.Name + `)
			}`
			}
		}

		for _, imp := range template.Header.Objects {
			local_string += `
			type ` + imp.Name + ` ` + imp.Templ
		}

		//Create an object map
		for _, imp := range template.Header.Objects {
			//struct return and function
			fmt.Println("∑ Processing object :" + imp.Name)
			if !contains(available_methods, imp.Name) {
				//addcontructor
				available_methods = append(available_methods, imp.Name)
				int_methods = append(int_methods, imp.Name)
				local_string += `
				func  net_` + imp.Name + `(args ...interface{}) (d ` + imp.Templ + `){
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					return
					} else {
						d = ` + imp.Templ + `{} 
						return
					}
				}`

			}

			delegateMethods := strings.Split(imp.Methods, "\n")

			for _, im := range delegateMethods {

				if stripSpaces(im) != "" {
					fmt.Println(imp.Name + "->" + im)
					function_map := strings.Split(im, ")")

					if !contains(int_mappings, function_map[0]+imp.Templ) {
						int_mappings = append(int_mappings, function_map[0]+imp.Templ)
						funcsp := strings.Split(function_map[0], "(")
						meth := template.findMethod(stripSpaces(funcsp[0]))

						//process limits and keep local deritives
						if meth.Autoface == "" || meth.Autoface == "true" {

							/*

							 */
							procc_funcs := true
							fmt.Println()

							if meth.Limit != "" {
								if !contains(strings.Split(meth.Limit, ","), imp.Name) {
									procc_funcs = false
								}
							}

							if contains(api_methods, meth.Name) {
								procc_funcs = false
							}

							objectName := meth.Object
							if objectName == "" {
								objectName = "object"
							}
							if procc_funcs {
								if !contains(int_methods, stripSpaces(funcsp[0])) && meth.Name != "000" {
									int_methods = append(int_methods, stripSpaces(funcsp[0]))
								}
								local_string += `
					  	func  net_` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1]+`, `+objectName+` `+imp.Templ, ",") + `) ` + stripSpaces(function_map[1])
								if stripSpaces(function_map[1]) == "" {
									local_string += ` string`
								}

								local_string += ` {
									` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

								if stripSpaces(function_map[1]) == "" {
									local_string += ` 
								return ""
							`
								}
								local_string += ` 
						}`

								if meth.Keeplocal == "false" || meth.Keeplocal == "" {
									local_string += `
						func (` + objectName + ` ` + imp.Templ + `) ` + stripSpaces(funcsp[0]) + `(` + strings.Trim(funcsp[1], ",") + `) ` + stripSpaces(function_map[1])

									local_string += ` {
							` + strings.Replace(meth.Method, `&#38;`, `&`, -1)

									local_string += `
						}`
								}
							}
						}

					}
				}
			}

			//create Unused methods methods
			fmt.Println(int_methods)
			for _, imp := range available_methods {
				if !contains(int_methods, imp) && !contains(api_methods, imp) {
					fmt.Println("Processing : " + imp)
					meth := template.findMethod(imp)
					addedit := false
					if meth.Returntype == "" {
						meth.Returntype = "string"
						addedit = true
					}
					local_string += `
						func net_` + meth.Name + `(args ...interface{}) ` + meth.Returntype + ` {
							`
					for k, nam := range strings.Split(meth.Variables, ",") {
						if nam != "" {
							local_string += nam + ` := ` + `args[` + strconv.Itoa(k) + `]
								`
						}
					}
					local_string += strings.Replace(meth.Method, `&#38;`, `&`, -1)
					if addedit {
						local_string += `
						 return ""
						 `
					}
					local_string += `
						}`
				}
			}
			for _, imp := range template.Templates.Templates {
				local_string += `
				func  net_` + imp.Name + `(args ...interface{}) string {
					var d ` + imp.Struct + `
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return ""
					}
					} else {
						d = ` + imp.Struct + `{}
					}

					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
			
				    error := t.Execute(output, &d)
				    if error != nil {
				    fmt.Print(error)
				    } 
					return html.UnescapeString(output.String())
				}`
				local_string += `
				func net_b` + imp.Name + `(d ` + imp.Struct + `) string {
					return  b` + imp.Name + `(d)
				}
				func  b` + imp.Name + `(d ` + imp.Struct + `) string {
					filename :=  "` + tmpl + `/` + imp.TemplateFile + `.tmpl"
    				body, er := Asset(filename)
    				if er != nil {
    					return ""
    				}
    				 output := new(bytes.Buffer) 
					t := template.New("` + imp.Name + `")
    				t = t.Funcs(` + netMa + `)
				  	t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"` + "`" + `", ` + "`" + `\"` + "`" + ` ,-1) )
			
				    error := t.Execute(output, &d)
				    if error != nil {
				    fmt.Print(error)
				    } 
					return html.UnescapeString(output.String())
				}`
				local_string += `
				func  c` + imp.Name + `(args ...interface{}) (d ` + imp.Struct + `) {
					if len(args) > 0 {
					var jsonBlob = []byte(args[0].(string))
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return 
					}
					} else {
						d = ` + imp.Struct + `{}
					}
    				return
				}
				func  net_c` + imp.Name + `(args ...interface{}) (d ` + imp.Struct + `) {
					if len(args) > 0 {
					var jsonBlob = []byte(args[0].(string))
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return 
					}
					} else {
						d = ` + imp.Struct + `{}
					}
    				return
				}`
			}

			//Methods have been added

			fmt.Println("Saving file to " + r + "/" + template.Output)
			d1 := []byte(local_string)
			_ = ioutil.WriteFile(r+"/"+template.Output, d1, 0644)

		}
	}

	return
}

func RunFile(root string, file string) {
	fmt.Println("∑ Running " + root + "/" + file)
	exe_cmd("go run " + root + "/" + file)
}

func RunCmd(cmd string) {
	exe_cmd(cmd)
}

func RunCmdString(cmd string) string {
	parts := strings.Fields(cmd)
	fmt.Println(cmd)
	var out *exec.Cmd
	if len(parts) == 5 {
		fmt.Println("Match")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4])
	} else if len(parts) == 9 {
		fmt.Println("Match GPGl")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6], parts[7], parts[8])
	} else if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2])
	} else if len(parts) == 1 {
		out = exec.Command(parts[0])
	} else {
		out = exec.Command(parts[0], parts[1])
	}

	var ou bytes.Buffer
	out.Stdout = &ou
	err := out.Run()
	if err != nil {
		fmt.Println("Error")
	}
	return ou.String()
}

func RunCmdByte(cmd string) []byte {
	parts := strings.Fields(cmd)
	fmt.Println(cmd)
	var out *exec.Cmd
	if len(parts) == 5 {
		fmt.Println("Match")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4])
	} else if len(parts) == 9 {
		fmt.Println("Match GPGl")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6], parts[7], parts[8])
	} else if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2])
	} else if len(parts) == 1 {
		out = exec.Command(parts[0])
	} else {
		out = exec.Command(parts[0], parts[1])
	}

	var ou bytes.Buffer
	out.Stdout = &ou
	err := out.Run()
	if err != nil {
		fmt.Println("Error")
	}
	return ou.Bytes()
}

func RunCmdSmartB(cmd string) ([]byte, error) {
	parts := strings.Fields(cmd)
	fmt.Println(parts[0], parts[1:])
	var out *exec.Cmd

	if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else {
		out = exec.Command(parts[0], parts[1:]...)
	}
	var ou, our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	err := out.Run()
	fmt.Println(our.String())
	if err != nil {
		return our.Bytes(), err
	}
	return ou.Bytes(), nil
}

func RunCmdSmart(cmd string) (string, error) {
	parts := strings.Fields(cmd)
		//fmt.Println(parts)
	var out *exec.Cmd

	if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else {
		out = exec.Command(parts[0], parts[1:]...)
	}

	var ou, our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	fmt.Println(BytesToString(our.Bytes()))
	err := out.Run()
	if err != nil {
		//	fmt.Println("%v", err.Error())
		return our.String(), err
	}
	return ou.String(), nil
}

func RunCmdSmarttwo(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	//	fmt.Println(parts[0],parts[1:])
	var out *exec.Cmd

	out = exec.Command(parts[0], parts[1:]...)

	var ou, our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	fmt.Println(BytesToString(our.Bytes()))
	err := out.Run()
	if err != nil {
		//	fmt.Println("%v", err.Error())
		return our.String(), err
	}
	return ou.String(), nil
}

func RunCmdSmartZ(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	//	fmt.Println(parts[0],parts[1:])
	var out *exec.Cmd

	out = exec.Command(parts[0], parts[1])

	var ou, our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	fmt.Println(BytesToString(our.Bytes()))
	err := out.Run()
	if err != nil {
		//	fmt.Println("%v", err.Error())
		return ou.String() + our.String(), err
	}
	return ou.String(), nil
}

func RunCmdSmartP(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	//	fmt.Println(parts[0],parts[1:])
	var out *exec.Cmd

	out = exec.Command(parts[0], parts[1], parts[2], "-benchmem")

	var ou, our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	fmt.Println(BytesToString(our.Bytes()))
	err := out.Run()
	if err != nil {
		//	fmt.Println("%v", err.Error())
		return ou.String() + our.String(), err
	}
	return ou.String(), nil
}

func RunCmdSmartCmb(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	fmt.Println(parts[0], parts[1:])
	var out *exec.Cmd

	out = exec.Command(parts[0], parts[1:]...)

	ou, err := out.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(ou), nil
}

func BytesToString(b []byte) string {
	return string(b)
}

func RunCmdB(cmd string) {
	parts := strings.Fields(cmd)
	fmt.Println(cmd)
	var out *exec.Cmd
	if len(parts) == 5 {
		fmt.Println("Match")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4])
	} else if len(parts) == 9 {
		fmt.Println("Match GPGl")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6], parts[7], parts[8])
	} else if len(parts) == 8 {
		fmt.Println("Match decrypt GPGl")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6], parts[7], "pass")

	} else if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2])
	} else if len(parts) == 1 {
		out = exec.Command(parts[0])
	} else {
		out = exec.Command(parts[0], parts[1], "2>&1")
	}

	var ou bytes.Buffer
	out.Stdout = &ou
	err := out.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ou.String())
}

func RunCmdA(cm string) error {

	parts := strings.Fields(cm)
	fmt.Println(parts)
	var cmd *exec.Cmd
	fmt.Println("Match decrypt GPGl")
	cmd = exec.Command(parts[0], parts[1], parts[2], "--no-tty", parts[3], parts[4], parts[5], parts[6])
	inpipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Pipe : ", err)
	}
	io.WriteString(inpipe, "")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println("Result: " + out.String())
	return nil
}

func exe_cmd(cmd string) {
		defer func() {
		if n := recover(); n != nil {
			
			fmt.Println(n)
		}
	}()
	parts := strings.Fields(cmd)
	fmt.Println(cmd)
	var out *exec.Cmd
	if len(parts) == 5 {
		fmt.Println("Match")
		out = exec.Command(parts[0], parts[1], parts[2], parts[3], parts[4])
	} else if len(parts) == 4 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2])
	} else if len(parts) == 1 {
		out = exec.Command(parts[0])
	} else {
		out = exec.Command(parts[0], parts[1])
	}
	stdoutStderr, err := out.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func Exe_Stall(cmd string, chn chan bool) {
	//fmt.Println(cmd)
	parts := strings.Fields(cmd)
	var out *exec.Cmd
	if len(parts) > 3 {
		out = exec.Command(parts[0], parts[1], parts[2], parts[3])
	} else if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2], "2>&1")
	} else if len(parts) == 1 {
		out = exec.Command(parts[0], "2>&1")
	} else {
		out = exec.Command(parts[0], parts[1], "2>&1")
	}
	stdout, err := out.StdoutPipe()

	if err != nil {
		fmt.Println("error occurred")
		fmt.Printf("%s", err)
	}
	out.Start()
	r := bufio.NewReader(stdout)

	t := false

	go func() {
		for !t {
			line, _, _ := r.ReadLine()
			if string(line) != "" {
				log.Println(string(line))
			}

		}
	}()
	tch, m := <-chn
	if m && tch {
		t = tch
	}
	log.Println("Killing proc.")
	if err := out.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
	}
}

func Exe_Stalll(cmd string) {
	fmt.Println(cmd)
	parts := strings.Fields(cmd)
	var out *exec.Cmd
	if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2], "2>&1")
	} else if len(parts) == 1 {
		out = exec.Command(parts[0], "2>&1")
	} else {
		out = exec.Command(parts[0], parts[1], "2>&1")
	}
	stdout, err := out.StdoutPipe()

	if err != nil {
		fmt.Println("error occurred")
		fmt.Printf("%s", err)
	}
	out.Start()
	r := bufio.NewReader(stdout)

	t := false
	for !t {
		line, _, _ := r.ReadLine()
		if string(line) != "" {
			fmt.Println(string(line))
		}

	}

}

func Exe_BG(cmd string) {
	fmt.Println(cmd)
	parts := strings.Fields(cmd)
	var out *exec.Cmd
	if len(parts) > 2 {
		out = exec.Command(parts[0], parts[1], parts[2])
	} else if len(parts) == 1 {
		out = exec.Command(parts[0])
	} else {
		out = exec.Command(parts[0], parts[1])
	}
	stdout, err := out.StdoutPipe()
	if err != nil {
		fmt.Println("error occurred")
		fmt.Printf("%s", err)
	}
	out.Start()
	r := bufio.NewReader(stdout)
	t := false
	for !t {
		line, _, _ := r.ReadLine()
		if string(line) != "" {
			fmt.Println(string(line))
		}
	}
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

func (d *gos) findStruct(name string) Struct {
	for _, imp := range d.Header.Structs {
		if imp.Name == name {
			return imp
		}
	}
	return Struct{Name: "000"}
}

func (d *gos) findMethod(name string) Method {
	for _, imp := range d.Methods.Methods {
		if imp.Name == name {
			return imp
		}
	}
	return Method{Name: "000"}
}

func (d *gos) PSaveGos(path string) {
	b, _ := xml.Marshal(d)
	ioutil.WriteFile(path, b, 0644)
}

func (d *gos) Delete(typ, id string) {
	if typ == "var" {
		temp := []GlobalVariables{}
		for _, v := range d.Variables {
			if v.Name != id {
				temp = append(temp, v)
			}
		}
		d.Variables = temp
	} else if typ == "import" {
		temp := []Import{}
		for _, v := range d.RootImports {
			if v.Src != id {
				temp = append(temp, v)
			}
		}
		d.RootImports = temp
	} else if typ == "timer" {

		temp := []Timer{}
		for _, v := range d.Timers.Timers {
			if v.Name != id {
				temp = append(temp, v)
			}
		}
		d.Timers.Timers = temp

	} else if typ == "end" {
		temp := []Endpoint{}
		for _, v := range d.Endpoints.Endpoints {
			if v.Path != id {
				temp = append(temp, v)
			}
		}
		d.Endpoints.Endpoints = temp
	} else if typ == "template" {

		temp := []Template{}
		for _, v := range d.Templates.Templates {
			if v.Name != id {
				temp = append(temp, v)
			}
		}
		d.Templates.Templates = temp

	} else if typ == "bundle" {

		temp := []Template{}
		for _, v := range d.Templates.Templates {
			if v.Bundle != id {
				temp = append(temp, v)
			}
		}
		d.Templates.Templates = temp

	}

}

func CreateVGos(path string) *VGos {
	v := &VGos{}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	d := xml.NewDecoder(bytes.NewReader(body))
	d.Entity = map[string]string{
		"&": "&",
	}
	err = d.Decode(&v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	return v
}

func (d *gos) MStructs(ne []Struct) {
	d.Header.Structs = ne
}

func (d *gos) MObjects(ne []Object) {
	d.Header.Objects = ne
}

func (d *gos) MMethod(ne []Method) {
	d.Methods.Methods = ne
}

func PLoadGos(path string) (*gos, *Error) {
	fmt.Println("∑ loading " + path)
	v := &gos{}
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &Error{code: 404, reason: "file not found! @ " + path}
	}

	//obj := Error{}
	//fmt.Println(obj);
	d := xml.NewDecoder(bytes.NewReader(body))
	d.Strict = false
	err = d.Decode(&v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, nil
	}

	for _, imp := range v.RootImports {
		//fmt.Println(imp.Src)
		if strings.Contains(imp.Src, ".gxml") {
			v.MergeWithV(os.ExpandEnv("$GOPATH") + "/" + strings.Trim(imp.Src, "/"))
			//copy files
		}
	}

	return v, nil
}

func (d *gos) Add(sec, typ, name string) {
	if sec == "var" {
		d.Variables = append(d.Variables, GlobalVariables{Name: name, Type: typ})
	} else if sec == "import" {
				d.RootImports = append(d.RootImports, Import{Src: name})
	} else if sec == "end" {
		d.Endpoints.Endpoints = append(d.Endpoints.Endpoints, Endpoint{Path: name,Id: NewID(15)})
	} else if sec == "timer" {
		d.Timers.Timers = append(d.Timers.Timers, Timer{Name: name})
	}
}

func (d *gos) AddS(sec string, typ interface{}) {
	if sec == "template" {
		d.Templates.Templates = append(d.Templates.Templates, typ.(Template))
	} else if sec == "import" {
		//d.RootImports = append(d.RootImports,Import{Src: name})
	}
}

func Decode64(decBuf, enc []byte) []byte {
	e64 := base64.StdEncoding
	maxDecLen := e64.DecodedLen(len(enc))
	if decBuf == nil || len(decBuf) < maxDecLen {
		decBuf = make([]byte, maxDecLen)
	}
	n, err := e64.Decode(decBuf, enc)
	_ = err
	return decBuf[0:n]
}

func (d *gos) UpdateMethod(id string, data string) {
	
		temp := []Endpoint{}
		for _, v := range d.Endpoints.Endpoints {
			if id == v.Id {
				v.Method = data
				temp = append(temp, v)
			} else {
				temp = append(temp, v)
			}
		}
		d.Endpoints.Endpoints = temp
	
}

func (d *gos) Update(sec, id string, update interface{}) {
	if sec == "var" {
		temp := []GlobalVariables{}
		for _, v := range d.Variables {
			if id == v.Name {
				temp = append(temp, update.(GlobalVariables))
			} else {
				temp = append(temp, v)
			}
		}
		d.Variables = temp
	} else if sec == "import" {
		//d.RootImports = append(d.RootImports,Import{Src: name})
		temp := []Import{}
		for _, v := range d.RootImports {
			if id == v.Src {
				temp = append(temp, update.(Import))
			} else {
				temp = append(temp, v)
			}
		}
		d.RootImports = temp
	} else if sec == "template" {
		temp := []Template{}
		for _, v := range d.Templates.Templates {
			if id == v.Name {
				v.Struct = update.(string)
				temp = append(temp, v)
			} else {
				temp = append(temp, v)
			}
		}
		d.Templates.Templates = temp
	} else if sec == "timer" {
		temp := []Timer{}
		for _, v := range d.Timers.Timers {
			if id == v.Name {

				temp = append(temp, update.(Timer))
			} else {
				temp = append(temp, v)
			}
		}
		d.Timers.Timers = temp
	} else if sec == "end" {
		temp := []Endpoint{}
		for _, v := range d.Endpoints.Endpoints {
			if id == v.Id {
				upd := update.(Endpoint)
				upd.Method = v.Method
				upd.Id = id
				temp = append(temp, upd)
			} else {
				temp = append(temp, v)
			}
		}
		d.Endpoints.Endpoints = temp
	}
}

func (d *gos) Set(attr, value string) {
	if attr == "app" {
		d.Type = value
	} else if attr == "port" {
		d.Port = value
	} else if attr == "key" {
		d.Key = value
	} else if attr == "erpage" {
		d.ErrorPage = value
	} else if attr == "fpage" {
		d.NPage = value
	} else if attr == "domain" {
		d.Domain = value
	}
}

func VLoadGos(path string) (gos, *Error) {
	fmt.Println("∑ loading " + path)
	v := gos{}
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return v, &Error{code: 404, reason: "file not found! @ " + path}
	}

	//obj := Error{}
	//fmt.Println(obj);
	d := xml.NewDecoder(bytes.NewReader(body))
	d.Strict = false
	err = d.Decode(&v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return v, nil
	}

	if v.FolderRoot == "" {
		ab := strings.Split(path, "/")
		v.FolderRoot = strings.Join(ab[:len(ab)-1], "/") + "/"
	}
	//process mergs
	for _, imp := range v.RootImports {
		//fmt.Println(imp.Src)
		if strings.Contains(imp.Src, ".gxml") {
			srcP := strings.Split(imp.Src, "/")
			dir := strings.Join(srcP[:len(srcP)-1], "/")
			if _, err := os.Stat(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + dir); os.IsNotExist(err) {
				// path/to/whatever does not exist
				//fmt.Println("")
				RunCmdSmart("go get " + dir)
			}
		} else {
			dir := TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(imp.Src, "/")
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				// path/to/whatever does not exist
				//fmt.Println("")
				RunCmdSmart("go get " + imp.Src)
			}
		}

		if strings.Contains(imp.Src, ".gxml") {
			v.MergeWith(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(imp.Src, "/"))
			//copy files
		}
		//
	}

	return v, nil
}

func LoadGos(path string) (*gos, *Error) {
	fmt.Println("∑ loading " + path)
	v := gos{}
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &Error{code: 404, reason: "file not found! @ " + path}
	}

	//obj := Error{}
	//fmt.Println(obj);
	d := xml.NewDecoder(bytes.NewReader(body))
	d.Strict = false
	err = d.Decode(&v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, nil
	}

	if v.FolderRoot == "" {
		ab := strings.Split(path, "/")
		v.FolderRoot = strings.Join(ab[:len(ab)-1], "/") + "/"
	}
	//process mergs
	for _, imp := range v.RootImports {
		//fmt.Println(imp.Src)
		if strings.Contains(imp.Src, ".gxml") {
			srcP := strings.Split(imp.Src, "/")
			dir := strings.Join(srcP[:len(srcP)-1], "/")
			if _, err := os.Stat(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + dir); os.IsNotExist(err) {
				// path/to/whatever does not exist
				//fmt.Println("")
				RunCmdSmart("go get " + dir)
			}
		} else {
			dir := TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(imp.Src, "/")
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				// path/to/whatever does not exist
				//fmt.Println("")
				RunCmdSmart("go get " + imp.Src)
			}
		}

		if strings.Contains(imp.Src, ".gxml") {
			v.MergeWith(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(imp.Src, "/"))
			//copy files
		}
		//
	}

	return &v, nil
}

func (d *gos) MergeWithV(target string) {
	fmt.Println("∑ Merging " + target)
	imp, err := LoadGos(target)
	if err != nil {
		fmt.Println(err)
	} else {

		for _, im := range imp.RootImports {
			if strings.Contains(im.Src, ".gxml") {
				imp.MergeWithV(os.ExpandEnv("$GOPATH") + "/" + strings.Trim(im.Src, "/"))
				//copy files
			} else {
				d.RootImports = append(d.RootImports, im)
			}
		}

		if imp.FolderRoot == "" {
			ab := strings.Split(target, "/")
			imp.FolderRoot = strings.Join(ab[:len(ab)-1], "/") + "/"
		}
		//d.RootImports = append(imp.RootImports,d.RootImports...)
		d.Header.Structs = append(imp.Header.Structs, d.Header.Structs...)
		d.Header.Objects = append(imp.Header.Objects, d.Header.Objects...)
		d.Methods.Methods = append(imp.Methods.Methods, d.Methods.Methods...)
		d.Timers.Timers = append(imp.Timers.Timers, d.Timers.Timers...)
		//Specialize method for templates
		//d.Variables = append(imp.Variables, d.Variables...)
		if imp.Package != "" && imp.Type == "package" {
			fmt.Println("Parsing Prefixes for " + imp.Package)
			for _, im := range imp.Templates.Templates {
				im.TemplateFile = imp.Package + "/" + im.TemplateFile
				d.Templates.Templates = append(d.Templates.Templates, im)
			}
		} else {
			d.Templates.Templates = append(imp.Templates.Templates, d.Templates.Templates...)
		}
		//copy files
		d.Endpoints.Endpoints = append(imp.Endpoints.Endpoints, d.Endpoints.Endpoints...)
	}

	d.Init_Func = d.Init_Func + ` 
	` + imp.Init_Func
}

func (d *gos) MergeWith(target string) {
	fmt.Println("∑ Merging " + target)
	imp, err := LoadGos(target)
	if err != nil {
		fmt.Println(err)
	} else {

		for _, im := range imp.RootImports {

			if strings.Contains(im.Src, ".gxml") {
				srcP := strings.Split(im.Src, "/")
				dir := strings.Join(srcP[:len(srcP)-1], "/")
				if _, err := os.Stat(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + dir); os.IsNotExist(err) {
					// path/to/whatever does not exist
					//fmt.Println("")
					RunCmdSmart("go get " + dir)
				}
			} else {
				dir := TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(im.Src, "/")
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					// path/to/whatever does not exist
					//fmt.Println("")
					RunCmdSmart("go get " + im.Src)
				}
			}
			if strings.Contains(im.Src, ".gxml") {
				imp.MergeWith(TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + strings.Trim(im.Src, "/"))
				//copy files
			} else {
				d.RootImports = append(d.RootImports, im)
			}
		}

		//d.RootImports = append(imp.RootImports,d.RootImports...)
		d.Header.Structs = append(imp.Header.Structs, d.Header.Structs...)
		d.Header.Objects = append(imp.Header.Objects, d.Header.Objects...)
		d.Methods.Methods = append(imp.Methods.Methods, d.Methods.Methods...)
		d.Timers.Timers = append(imp.Timers.Timers, d.Timers.Timers...)
		//Specialize method for templates
		d.Variables = append(imp.Variables, d.Variables...)

		if imp.Package != "" && imp.Type == "package" {
			fmt.Println("Parsing Prefixes for " + imp.Package)
			for _, im := range imp.Templates.Templates {
				im.TemplateFile = imp.Package + "/" + im.TemplateFile
				d.Templates.Templates = append(d.Templates.Templates, im)
			}
		} else {
			d.Templates.Templates = append(imp.Templates.Templates, d.Templates.Templates...)
		}

		if imp.Tmpl == "" {
			imp.Tmpl = "tmpl"
		}

		if imp.Web == "" {
			imp.Web = "web"
		}

		if d.Tmpl == "" {
			d.Tmpl = "tmpl"
		}

		if d.Web == "" {
			d.Web = "web"
		}
		d.Init_Func = d.Init_Func + ` 
	` + imp.Init_Func
		d.Main = d.Main + ` 
	` + imp.Main
		os.MkdirAll(d.Tmpl+"/"+imp.Package, 0777)
		os.MkdirAll(d.Web+"/"+imp.Package, 0777)
		CopyDir(imp.FolderRoot+imp.Tmpl, d.Tmpl+"/"+imp.Package)
		CopyDir(imp.FolderRoot+imp.Web, d.Web+"/"+imp.Package)

		//copy files
		d.Endpoints.Endpoints = append(imp.Endpoints.Endpoints, d.Endpoints.Endpoints...)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DoubleInput(p1 string, p2 string) (r1 string, r2 string) {
	fmt.Println(p1)
	fmt.Scanln(&r1)
	fmt.Println(p2)
	fmt.Scanln(&r2)
	return
}

func AskForConfirmation() bool {
	var response string
	fmt.Println("Please type yes or no and then press enter:")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return AskForConfirmation()
	}
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func getPy() string {
	return ``
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
