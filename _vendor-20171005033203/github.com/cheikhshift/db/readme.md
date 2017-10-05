# Go-Express
![enter image description here](https://camo.githubusercontent.com/ed1230a48b6946283ee0d76185a726b49ba58254/68747470733a2f2f7472617669732d63692e6f72672f746f6f6c732f676f6465702e737667)

Simple Mongo client for Go. The package aims to match Express seamless integration with NodeJs and MongoDb.

[Index](#index)

## Requirements

 - MongoDb (access to a server)
 - [Mgo](https://godoc.org/labix.org/v2/mgo#Query) 

	
#### Dependencies
(GoDeps supported as well.)
Manually:

		go get github.com/asaskevich/govalidator
		go get gopkg.in/mgo.v2

#### install
	
	go get github.com/cheikhshift/db


## Import

	import (
		"log" //optional
		"github.com/cheikhshift/db"
	)

## Connect

The snippet below will connect to your `database` @ your `localhost`

	dbs,err := db.Connect("localhost","database") 
	if err != nil {
		log.Fatal(err)
	}

Format of server URI : 

	[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]

## Add new model

Define a new struct as usual (fields Created and Id are system generated) :

	type MyObject struct {
		Id bson.ObjectId `bson:"_id,omitempty"`
		TestField string `valid:"unique"`
		FieldTwo string `valid:"email,required"`
		Created time.Time //timestamp local format
	}

The following tag will ensure no duplicates happen while saving. 

	TestField string `valid:"unique"`

The following tag `valid` will ensure your field data is an email.

	 FieldTwo string `valid:"email,required"`

## Save a model
	
	obj := NewObject{FieldTwo: "value" }
	dbs.New(&obj)	
	err = dbs.Save(&obj)
	if err != nil {
		log.Fatal(err)
	}

## Retrieve model

The following query will return one object. Within the Find function you may use all of the usual Mongo DB query functions. Remember to add strings even around `$` prefixed parameters. The following example will retrieve one result. Any struct field is converted to lower case on database save, when performing a query type struct field names without capital letters. 

	query := MyObject{}
	dbs.Q(query).Find(db.O{"fieldtwo":"value"}).One(&query)

	fmt.Println(query)

## Delete model
This function will fail if your struct's `Id` field is blank.

	obj := NewObject{FieldTwo: "value" }
	err := dbs.Add(&obj)	
	if err != nil {
		log.Fatal(err)
	}
	
	
	dbs.Remove(&obj)


#Index
	
	package github.com/cheikhshift/db

## Structs

### DB
	type DB struct {
		MoDb *mgo.Database
	}
	
### O
	type O bson.M
	

## Functions

### Connect
Create a new database connection.

	func Connect(url string, db string) (DB,error)

-url : MongoDb server URI to connect to.
- db: database to use.


## Close

Close database connection.

	func (d DB) Close()

### Validate
Verify if model is valid. This [package](https://github.com/asaskevich/govalidator) is used for validation, Read more about supported types [here](https://github.com/asaskevich/govalidator).

	func (d DB) PreVerify(item interface{}) error

- item : Model (struct) to verify.

### Query
More information about mgo.Query [here](https://godoc.org/labix.org/v2/mgo#Query). Once your query is retrieved you may count ,limit, skip as you please.

	func (d DB) Query(item interface{}, query interface{}) *mgo.Query

- item : Model to set query to.
- query : `db.O`  map of parameters to query.


### Add
Add the specified models to your database.

	func (d DB) Add(items ...interface{}) error {

- items : models to save.

### Save
Update the model.

	func (d DB) Save(item interface{}) error


- item : model to update.


### Remove
Remove model from database. If the model has no value for field Id this function will fail.

	func (d DB) Remove(items ...interface{}) error 

- items : items to remove.

### RemoveAll
Remove items via `db.O` query.


	func (d DB) RemoveAll(item interface{},query  interface{}) (*mgo.ChangeInfo,error)


- item : model to set query to.
- query :  map of parameters to query.

### UpdateAll ...Save `All` 
Update models in your database via query.

	func (d DB) UpdateAll(item interface{},query interface{}, update interface{}) (*mgo.ChangeInfo,error) 

- item : model to set query to.
- query : map of parameters to query.
- update : map of updates to make. Remember to use MongoDb selections , ie : `$ne`.