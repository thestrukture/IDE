package main

import (
	//	"bufio"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"sort"
	"strings"
)

var dfd string

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

/*
func GetLine(fname string, match string) int {
	intx := 0
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Could not find a source file")
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
} */

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
		fmt.Println(err.Error())
		// os.Exit(1)
	}

	var c []App
	json.Unmarshal(raw, &c)
	return c
}

func getPlugins() []string {

	raw, err := ioutil.ReadFile(dfd + "/plugins.json")
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
	}

	var c []string
	json.Unmarshal(raw, &c)
	return c
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
