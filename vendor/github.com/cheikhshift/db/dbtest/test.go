package main

import (
"github.com/cheikhshift/db"
"fmt"
"log"

)


type MyObject db.MyObject


func main() {
	dbs,err := db.Connect("localhost","database") 
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Save object")
	obj := MyObject{TestField:"second", FieldTwo: "Ivalud@update.com"}
	dbs.New(&obj)
	
	err = dbs.Save(&obj)
	if err != nil {
		fmt.Println(err)
	}
	query := MyObject{}
	dbs.Q(query).Find(db.O{"fieldtwo": "Ivalud@update.com"}).One(&query)

	fmt.Println(query)
	dbs.Close()
	/* 
		Directly add new items withoout calling dbs.New
	
	dbs.Add(&MyObject{})
	dbs.Add(&MyObject{TestField: "Value"})
	//dbs.Add(&MyObject{})

	size, err := dbs.Count(&MyObject{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Size of `MyObject` collection is " , size)

	//dbs.Add(MyObject{TestField: "Object"})
	*/
}