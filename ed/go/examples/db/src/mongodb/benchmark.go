package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	Id int
	Sha1 string
	Path string
}

func (d *Document) Print() {
	fmt.Printf("%d, %s, %s \n", d.Id, d.Sha1, d.Path)
}

func main() {
	now := time.Now()
	nanos := now.UnixNano()

	action()

	now2 := time.Now()
	nanos2 := now2.UnixNano()

	fmt.Printf("Took: %d microseconds", (nanos2 - nanos) / 1000)
}

func action() {
	connStr := "dbu:dbp@xmongo:27017/test"
	session, _ := mgo.Dial(connStr)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("device_file")

	query := []bson.M{
		bson.M{"$match": bson.M{"count_devices_for_file_sandbox": bson.M{"$gt": 0}}},
		bson.M{"$unwind": "$sandbox"},
		bson.M{"$project": bson.M{"_id": 0, "id": 1, "sha1": 1, "path": "$sandbox.path"}},
		bson.M{"$sort": bson.M{"id": -1, "path": 1}},
		bson.M{"$skip": 100},
		bson.M{"$limit": 10}}
	var docs []Document
	c.Pipe(query).All(&docs)

	for _, d := range docs {
		d.Print()
	}
}
