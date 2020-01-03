package di

import (
	"gopkg.in/mgo.v2"
)

var (
	containerInstance container
)

type container struct {
	initializedItems map[string]bool

	mongoDB *mgo.Database
	mongoDBSession *mgo.Session
}

func Init() {
	containerInstance.initializedItems = make(map[string]bool)
}

func Deinit() {
	go containerInstance.mongoDBSession.Close()
}

func GetMongoDB() *mgo.Database {
	if !containerInstance.initializedItems["MongoDB"] {
		InitMongoDB()
	}

	return containerInstance.mongoDB
}

func InitMongoDB() {
	session, err := mgo.Dial(MongoDBConnectionString)
	if err != nil {
		panic("RUNTIME-ERROR-DB-1: " + err.Error())
	}
	session.SetMode(mgo.Monotonic, true)

	containerInstance.mongoDB = session.DB("test")
	containerInstance.mongoDBSession = session
	containerInstance.initializedItems["MongoDB"] = true
}
