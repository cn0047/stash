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

	insert(c)
	findByCode(c, 204)
	findByCode(c, 205)
}

func insert(c *mgo.Collection) {
	err := c.Insert(&Test{202}, &Test{204})
	if err != nil {
		log.Fatal(err)
	}
}

func find(c *mgo.Collection, code int) {
	result := Test{}
	err := c.Find(bson.M{"code": code}).One(&result)
	if err != nil {
		log.Fatal("NOT FOUND.")
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func findByCode(c *mgo.Collection, code int) {
	find(c, code)
}
