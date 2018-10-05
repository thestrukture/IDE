package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var lock *sync.RWMutex = new(sync.RWMutex)
var connections = []*websocket.Conn{}

var dfd string

var dockerLarge = `FROM %s
ENV WEBAPP /go/src/server
RUN mkdir -p ${WEBAPP}
COPY . ${WEBAPP}
ENV PORT=%s 
WORKDIR ${WEBAPP}
RUN go install
RUN rm -rf ${WEBAPP} 
EXPOSE %s
CMD server
`

func generateComposeFile(r *http.Request) {

	mongo := r.FormValue("mongo")
	mPort := r.FormValue("mPort")

	redis := r.FormValue("redis")
	rPort := r.FormValue("rPort")

	postgr := r.FormValue("postgres")
	dbname := r.FormValue("dbname")
	username := r.FormValue("username")
	pass := r.FormValue("pass")
	pPort := r.FormValue("pPort")

	fport := r.FormValue("fport")
	image := r.FormValue("image")
	port := r.FormValue("port")

	pkg := r.FormValue("pkg")
	path := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg, "compose-file.yml")
	var fileCompleted string

	if mongo == "false" && redis == "false" && postgr == "false" {

		composefile := strings.Replace(composeBase, "links:", "", -1)

		fileCompleted = fmt.Sprintf(composefile, image, port, fport, "", "")

	} else {
		links := "%s"

		services := ""

		if mongo == "true" {
			links = fmt.Sprintf(links, "            - mongodb\n%s")
			services += fmt.Sprintf(mongotemp, mPort)
		}

		if postgr == "true" {
			links = fmt.Sprintf(links, "            - postgres\n%s")

			services += fmt.Sprintf(postgrestemp, pPort, username, pass, dbname)
		}

		if redis == "true" {
			links = fmt.Sprintf(links, "            - redis\n%s")
			services += fmt.Sprintf(redistemp, rPort)
		}

		links = fmt.Sprintf(links, "")

		fileCompleted = fmt.Sprintf(composeBase, image, port, fport, links, services)

	}

	err := ioutil.WriteFile(path, []byte(fileCompleted), 0700)
	if err != nil {
		log.Println(err)
	}

}

var composeBase string = `version: '2'
services:

    # Application container


    go:
        image: %s
        ports:
            - "%s:%s"
        links:
%s


    %s`

var mongotemp string = `mongodb:
        image: mvertes/alpine-mongo:3.2.3
        restart: unless-stopped
        ports:
            - "27017:%s"


            `

var postgrestemp string = `postgres:
        image: onjin/alpine-postgres:9.5
        restart: unless-stopped
        ports:
            - "5432:%s"
        environment:
            LC_ALL: C.UTF-8
            POSTGRES_USER: %s
            POSTGRES_PASSWORD: %s
            POSTGRES_DB: %s

            `

var redistemp string = `redis:
        image: sickp/alpine-redis:3.2.2
        restart: unless-stopped
        ports:
            - "6379:%s"

            `

var dockerSmall = `FROM %s as builder
ENV WEBAPP /go/src/server
RUN mkdir -p ${WEBAPP}
COPY . ${WEBAPP}
ENV PORT=%s 
WORKDIR ${WEBAPP}
RUN go install

# start from scratch
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]

EXPOSE %s
CMD server
`

var addjsstr = ` <script type="text/javascript">

					$(".marker .row",".endp-view").each(function(e,i){
							var attr = $(this).attr('proc-set');

							
							if (typeof attr !== typeof undefined && attr !== false) {
							  // Element has this attribute
								return;
							} else {
								   $(this).attr('proc-set', "can edit")
							}
							var tabid = $(this).parents(".tabview").attr("id")
							$( ".col-xs-12:nth-child(3)",this).css("text-align","left")
							var mrker = this
					 		$(".col-xs-12:nth-child(3)",this).prepend($("<button class='btn edt-code btn-success' path='" + $(this).attr("path") + "' tab-id='" + tabid + "'>Edit <span class='hidden-md-down'>code</span></button>").click(function(){
					 		 
					 			$.ajax({url: $(this).attr("path").replace("put?type=9", "get?type=13r"),error:function(err){
					 			
					 				$(".tabview.active .code-bin").html(err.responseText)
					 				}
							 	})
							 })
							)

					 	

					 })
				 	
				 	</script>`

type reader struct {
	Conn *websocket.Conn
}

func (r *reader) OnData(b []byte) bool {
	r.Conn.WriteMessage(1, b)
	return false
}

func (r *reader) OnError(b []byte) bool {
	if r.Conn != nil {
		r.Conn.WriteMessage(1, b)
	}
	return false
}

func (r *reader) OnTimeout() {

}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

// http://blog.ralch.com/tutorial/golang-working-with-zip/
func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func processLog(pkg, buildLog string) DebugObj {

	dObj := DebugObj{}
	logSlice := strings.Split(buildLog, "Full compiler build log :")
	fmt.Println(logSlice)

	return dObj
}

func Walk() {
	err := filepath.Walk(os.ExpandEnv("$GOPATH")+"/src/", visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

func getApps() []App {

	raw, err := ioutil.ReadFile(dfd + "/apps.json")
	if err != nil {
		return []App{}
		// os.Exit(1)
	}

	var c []App
	json.Unmarshal(raw, &c)
	return c
}

func getPlugins() []string {
	var c []string
	raw, err := ioutil.ReadFile(dfd + "/plugins.json")
	if err != nil {
		return []string{}
	}

	json.Unmarshal(raw, &c)
	return c
}

func saveKanBan(pkg, raw string) {
	gpath := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg, "kanban.json")
	fmt.Println("Writing ", gpath)
	ioutil.WriteFile(gpath, []byte(raw), 0700)
}

func getKanBan(pkg string) (ret map[string]interface{}) {
	gpath := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg, "kanban.json")
	fmt.Println("Reading ", gpath)
	raw, err := ioutil.ReadFile(gpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(raw, &ret)

	return

}

func AddConnection(c *websocket.Conn) {
	lock.Lock()
	connections = append(connections, c)
	lock.Unlock()
}

func Broadcast(m []byte) {
	lock.Lock()

	for index, c := range connections {
		if c != nil {
			err := c.WriteMessage(1, m)
			if err != nil {
				log.Println("write:", err)
				connections[index] = nil
			}
		}
	}

	lock.Unlock()
}

func pushGit(pkg string) {
	gpath := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg)

	fmt.Println("pushing ", gpath)

	os.Chdir(gpath)

	glog, err := core.RunCmdSmart(fmt.Sprintf("git push"))

	if err != nil {
		log.Println(err)
	}

	fmt.Print(glog)

}

func commitGit(pkg, message string) bool {
	gpath := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg)
	cmd := fmt.Sprintf("git commit -m \"%s\" -a ", message)

	fmt.Println("committing ", gpath)

	os.Chdir(gpath)

	fmt.Println("Running : ", cmd)
	glog, err := core.RunCmdSmart(cmd)

	if err != nil {
		log.Print(err)
		return true
	}

	log.Print(glog)

	return false

}

func FindinString(data string, match string) int {

	lines := strings.Split(data, "\n")
	//fmt.Println(lines)

	for i, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.Contains(trim, match) {
			return (i + 1)
		}
	}

	return -1
}

func FindString(fname string, match int) string {
	file, _ := os.Open(fname)
	scanner := bufio.NewScanner(file)
	inm := 0
	for scanner.Scan() {
		inm++
		//fmt.Println("%+V", inm)
		lin := scanner.Text()
		if inm == match {

			return lin
		}

	}

	return ""
}

func FindLine(fname string, match string) int {
	intx := 0
	file, err := os.Open(fname)
	if err != nil {
		//color.Red("Could not find a source file")
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intx = intx + 1
		trim := strings.TrimSpace(scanner.Text())
		if strings.Contains(trim, match) {

			return intx
		}

	}

	return -1
}

func reverse(numbers []DebugObj) []DebugObj {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func GetLogs(pkg string) []DebugObj {
	var c []DebugObj
	raw, err := ioutil.ReadFile(dfd + "/logs.json")
	if err != nil {
		//fmt.Println(err.Error())
		// os.Exit(1)
	} else {
		json.Unmarshal(raw, &c)
	}

	c = reverse(c)
	fSet := []DebugObj{}
	for _, dobj := range c {
		if dobj.PKG == pkg {
			fSet = append(fSet, dobj)
		}
	}

	return fSet
}

func ClearLogs(pkg string) {
	var c []DebugObj
	raw, err := ioutil.ReadFile(dfd + "/logs.json")
	if err != nil {
		//fmt.Println(err.Error())
		// os.Exit(1)
	} else {
		json.Unmarshal(raw, &c)
	}

	newlogs := []DebugObj{}
	for _, log := range c {
		if log.PKG != pkg {
			newlogs = append(newlogs, log)
		}
	}
	c = newlogs

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())

	}

	ioutil.WriteFile(dfd+"/logs.json", bytes, 0777)
}

func AddtoLogs(log DebugObj) {
	var c []DebugObj
	raw, err := ioutil.ReadFile(dfd + "/logs.json")
	if err != nil {
		//fmt.Println(err.Error())
		// os.Exit(1)
	} else {
		json.Unmarshal(raw, &c)
	}

	c = append(c, log)
	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())

	}

	ioutil.WriteFile(dfd+"/logs.json", bytes, 0777)

}

func saveApps(app interface{}) {
	bytes, err := json.Marshal(app)
	if err != nil {
		fmt.Println(err.Error())

	}
	ioutil.WriteFile(dfd+"/apps.json", bytes, 0777)

}

func savePlugins(plugins []string) {

	filtss := []string{}
	for _, v := range plugins {
		if v != "" {
			filtss = append(filtss, v)
		}
	}
	bytes, err := json.Marshal(filtss)
	if err != nil {
		fmt.Println(err.Error())

	}

	ioutil.WriteFile(dfd+"/plugins.json", bytes, 0777)

}
