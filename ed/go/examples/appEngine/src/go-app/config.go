package go_app

import (
	"encoding/json"
	"fmt"
	"github.com/mjibson/goon"
	"net/http"
)

type CreateConfigRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type IntConfig struct {
	Id    string `datastore:"-" goon:"id"`
	Key   string
	Value int
}

type StrConfig struct {
	Id    string `datastore:"-" goon:"id"`
	Key   string
	Value string
}

type Config struct {
	Id    string `datastore:"-" goon:"id"`
	Key   string
	Value interface{}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get()
	case "POST":
		post(w, r)
	default:
		w.Write([]byte(`{"error": "ERROR unsupported method"}`))
	}
}

func get() {

}

func post(w http.ResponseWriter, r *http.Request) {
	//conf1(w, r)
	//conf2(w, r)
	conf0(w, r)
}

func conf0(w http.ResponseWriter, r *http.Request) {
	dc := &CreateConfigRequest{}
	err := json.NewDecoder(r.Body).Decode(&dc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed decode request body: %s"}`, err.Error())))
		return
	}
	defer r.Body.Close()

	//log.Infof(appengine.NewContext(r), "\n 🔴 %+v", dc)
	//switch t := dc.Value.(type) {
	//default:
	//	log.Infof(appengine.NewContext(r), "\n ✴️ %T", t)
	//}

	//dc.Id = "id-" + dc.Key
	////dc.Id = "1"
	//
	//key, err := goon.NewGoon(r).Put(dc)
	//if err != nil {
	//	w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
	//	return
	//}
	//w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}

func conf1(w http.ResponseWriter, r *http.Request) {
	dc := &IntConfig{Id: "IntConfig1", Key: "IntConfig1", Value: 200}

	key, err := goon.NewGoon(r).Put(dc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}

func conf2(w http.ResponseWriter, r *http.Request) {
	dc := &StrConfig{Id: "StrConfig1", Key: "StrConfig1", Value: "100"}

	key, err := goon.NewGoon(r).Put(dc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}
