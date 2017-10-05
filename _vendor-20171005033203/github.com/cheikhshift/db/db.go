package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
	//	"fmt"
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
)

type DB struct {
	MoDb *mgo.Database
	S    *mgo.Session
}

type O bson.M

type MyObject struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	TestField string        `valid:"unique"`
	FieldTwo  string        `valid:"email,unique,required"`
	Created   time.Time     //timestamp local format
}

/*
govalidator.TagMap["unique"] = govalidator.Validator(func(item interface{}) bool {
	return true
}) */

func (d DB) Close() {
	d.S.Close()
}

func Connect(url string, db string) (DB, error) {

	session, err := mgo.Dial(url)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		return DB{}, err
	}
	connection := DB{session.DB(db), session}
	govalidator.TagMap["unique"] = govalidator.Validator(func(str string) bool {
		return true
	})

	return connection, nil
}

//update, find, upsert,count, add, find one
//Id bson.ObjectId `bson:"_id,omitempty"`
func (d DB) Count(item interface{}) (int, error) {
	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	return collection.Find(nil).Count()
}

func (d DB) C(item interface{}, ret interface{}) error {
	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	iter := collection.Find(nil)
	err := iter.All(&ret)
	if err != nil {
		return err
	}
	return nil
}

func mResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}

func RResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}

func ToBson(data string) bson.M {

	res := bson.M{}
	json.Unmarshal([]byte(data), &res)

	return res
}

func (d DB) PreVerify(item interface{}) error {
	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)

	_, err := govalidator.ValidateStruct(item)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(item).Elem()
	//t := reflect.TypeOf(item)
	bso := ToBson(mResponse(item))
	//fmt.Println(bso)
	for i := 0; i < v.NumField(); i++ {

		field := v.Type().Field(i)
		//field := t.Field(i)

		if strings.Contains(string(field.Tag), "unique") {
			chec := bson.M{}
			chec[strings.ToLower(field.Name)] = bso[strings.ToLower(field.Name)]
			if chec[strings.ToLower(field.Name)] == nil {
				chec[strings.ToLower(field.Name)] = bso[field.Name]
			}
			if _, ok := bso["Id"]; ok {
				if bso["Id"].(string) != "" {
					chec["_id"] = bson.M{"$ne": bson.ObjectIdHex(bso["Id"].(string))}
				}
			}

			count, err := collection.Find(chec).Count()

			if err != nil {
				return err
			}

			if count == 0 {
				return nil
			} else {
				return errors.New("Value set for " + field.Name + " already exists within your collection")
			}
		}

	}

	return nil
}

func (d DB) Q(item interface{}) *mgo.Collection {
	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	return collection
}

func (d DB) Query(item interface{}, query interface{}) *mgo.Query {
	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	return collection.Find(query)
}

func (d DB) Remove(items ...interface{}) error {

	for _, item := range items {
		object := strings.Split(reflect.TypeOf(item).String(), ".")
		Collection := object[len(object)-1]
		collection := d.MoDb.C(Collection)
		bso := ToBson(mResponse(item))

		err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(bso["Id"].(string))})
		if err != nil {
			return err
		}

	}

	return nil
}

func (d DB) RemoveAll(item interface{}, query interface{}) (*mgo.ChangeInfo, error) {

	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	ci, err := collection.RemoveAll(query)
	if err != nil {
		return nil, err
	}

	return ci, nil
}

func (d DB) UpdateAll(item interface{}, query interface{}, update interface{}) (*mgo.ChangeInfo, error) {

	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	collection := d.MoDb.C(Collection)
	ci, err := collection.UpdateAll(query, update)
	if err != nil {
		return nil, err
	}

	return ci, nil
}

func (d DB) Upsert(item interface{}) error {

	err := d.PreVerify(item)
	if err != nil {
		return err
	}

	object := strings.Split(reflect.TypeOf(item).String(), ".")
	Collection := object[len(object)-1]
	bso := ToBson(mResponse(item))
	stype := reflect.ValueOf(item).Elem()
	upd := stype.FieldByName("Updated")
	if upd.IsValid() {
		upd.Set(reflect.ValueOf(time.Now()))
	}

	_, err = d.MoDb.C(Collection).Upsert(bson.M{"_id": bson.ObjectIdHex(bso["Id"].(string))}, item)

	if err != nil {
		//perform verification too
		return err
	}

	return nil
}

func (d DB) Update(items interface{}) error {
	return d.Upsert(items)
}

func (d DB) Save(items interface{}) error {
	return d.Upsert(items)
}

func (d DB) New(item interface{}) interface{} {
	// object := strings.Split(reflect.TypeOf(item).String(),".")
	//Collection := object[len(object) - 1]

	stype := reflect.ValueOf(item).Elem()
	field := stype.FieldByName("Id")
	if field.IsValid() {
		field.Set(reflect.ValueOf(bson.NewObjectId()))
	}

	field = stype.FieldByName("Created")
	if field.IsValid() {
		field.Set(reflect.ValueOf(time.Now()))
	}
	//err := d.MoDb.C(Collection).Insert(item)

	return item
}

func (d DB) Add(items ...interface{}) error {

	for _, item := range items {

		err := d.PreVerify(item)
		if err != nil {
			return err
		}

		object := strings.Split(reflect.TypeOf(item).String(), ".")
		Collection := object[len(object)-1]
		//fmt.Println(reflect.TypeOf(item).String())
		stype := reflect.ValueOf(item).Elem()
		field := stype.FieldByName("Id")
		if field.IsValid() {
			field.Set(reflect.ValueOf(bson.NewObjectId()))
		}

		field = stype.FieldByName("Created")
		if field.IsValid() {
			field.Set(reflect.ValueOf(time.Now()))
		}

		field = stype.FieldByName("Updated")
		if field.IsValid() {
			field.Set(reflect.ValueOf(time.Now()))
		}

		err = d.MoDb.C(Collection).Insert(item)
		if err != nil {
			//perform verification too
			return err
		}

	}

	return nil
}
