/*

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"i", "value": 200, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"f", "value": 3.14, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"s", "value": "100", "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"o", "value": {"id": 9, "name": "x"}}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"o2", "value": {"id": "config", "value": {"n": 1, "b": false, "s": "ok"}}}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"a", "value": ["ok", "true"]}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"b", "value": true, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"n", "value": null, "tag": "simple"}' | jq

curl -XGET http://localhost:8080/config/i  | jq
curl -XGET http://localhost:8080/config/f  | jq
curl -XGET http://localhost:8080/config/s  | jq
curl -XGET http://localhost:8080/config/o  | jq
curl -XGET http://localhost:8080/config/o2 | jq
curl -XGET http://localhost:8080/config/a  | jq
curl -XGET http://localhost:8080/config/b  | jq
curl -XGET http://localhost:8080/config/n  | jq

curl -XGET http://localhost:8080/config/tag/simple | jq

// TODO:
// 1) curl -XGET http://localhost:8080/config/?q=tag1:name;tag2:*;tag3:footer,header;tagPrefix*:copyright
// 2) cache response payload

*/

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

type ConfigVO struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
}

type Config struct {
	Id    string `datastore:"-" goon:"id"`
	Key   string
	Value []byte
	Tag   string
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getByKey(w, r)
	case "POST":
		post(w, r)
	default:
		w.Write([]byte(`{"error": "ERROR unsupported method"}`))
	}
}

func configTagHandler(w http.ResponseWriter, r *http.Request) {
	getByTag(w, r)
}

func getByTag(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	p := strings.LastIndex(url, "/") + 1
	tag := url[p:]

	ctx := appengine.NewContext(r)
	q := datastore.
		NewQuery("Config").
		Filter("Tag =", tag)
	configs := make([]Config, 0)
	_, err := q.GetAll(ctx, &configs)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed GetAll from datastore: %s"}`, err.Error())))
		return
	}

	data := make(map[string]ConfigVO)
	for _, c := range configs {
		var value interface{}
		json.Unmarshal(c.Value, &value)
		vo := ConfigVO{Key: c.Key, Value: value, Tag: c.Tag}
		data[vo.Key] = vo
	}
	log.Infof(appengine.NewContext(r), "\n 🔴 %+v", data)

	payload, _ := json.Marshal(data)
	w.Write(payload)
}

func getByKey(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	p := strings.LastIndex(url, "/") + 1
	key := url[p:]

	c := Config{Id: "id-" + key}
	goon.NewGoon(r).Get(&c)
	log.Infof(appengine.NewContext(r), "\n 🔴 %+v", c)

	var value interface{}
	json.Unmarshal(c.Value, &value)
	vo := ConfigVO{Key: c.Key, Value: value, Tag: c.Tag}
	log.Infof(appengine.NewContext(r), "\n ☢️ %+v", vo)

	payload, _ := json.Marshal(vo)
	w.Write(payload)
}

func post(w http.ResponseWriter, r *http.Request) {
	create(w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
	vo := &ConfigVO{}
	err := json.NewDecoder(r.Body).Decode(&vo)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed decode request body: %s"}`, err.Error())))
		return
	}
	defer r.Body.Close()
	log.Infof(appengine.NewContext(r), "\n 🔴 %+v", vo)

	value, _ := json.Marshal(vo.Value)
	dc := Config{Id: "id-" + vo.Key, Key: vo.Key, Value: value, Tag: vo.Tag}
	log.Infof(appengine.NewContext(r), "\n ☢️ %+v", dc)

	key, err := goon.NewGoon(r).Put(&dc)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}
