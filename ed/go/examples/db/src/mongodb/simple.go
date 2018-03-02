package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Test struct {
	Code int
}

func main() {
	session, err := mgo.Dial("dbu:dbp@xmongo:27017/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("test")
	err = c.Insert(&Test{202}, &Test{204})
	if err != nil {
		log.Fatal(err)
	}

	result := Test{}
	err = c.Find(bson.M{"code": 204}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", result)
}
