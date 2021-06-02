package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Test struct {
	Code int
}

func main() {
	//u := "dbu:dbp@xmongo:27017/test"
	u := "dbu:dbp@localhost:27017/test"
	session, err := mgo.Dial(u)
	if err != nil {
		panic(fmt.Errorf("failed to dial mongo, err: %w", err))
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("test")

	insert(c)
	findByCode(c, 204)
	findByCode(c, 205)
}

func insert(c *mgo.Collection) {
	err := c.Insert(&Test{Code: 202}, &Test{Code: 204})
	if err != nil {
		panic(fmt.Errorf("failed to perform insert, err: %w", err))
	}
}

func find(c *mgo.Collection, code int) {
	result := Test{}
	err := c.Find(bson.M{"code": code}).One(&result)
	if err != nil {
		panic(fmt.Errorf("failed to perform find, err: %w", err))
	}
	fmt.Printf("%+v\n", result)
}

func findByCode(c *mgo.Collection, code int) {
	find(c, code)
}
