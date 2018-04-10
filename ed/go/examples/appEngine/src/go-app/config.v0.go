package go_app

import (
	"encoding/json"
	"fmt"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"
)

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

func configHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getByKey(w, r)
	default:
		w.Write([]byte(`{"error": "ERROR unsupported method"}`))
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	conf1(w, r)
	conf2(w, r)
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
