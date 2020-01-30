package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url1 = "https://api.github.com/users/cn007b"
	url2 = "https://realtimelog.herokuapp.com/sddjklskj"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func Get1() (*http.Response, error) {
	r, err := http.Get(url1)

	return r, err
}

func Get2() (User, error) {
	r, _ := http.Get(url1)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	u := User{}
	json.Unmarshal(body, &u)

	return u, nil
}

func Get3() (User, error) {
	req, _ := http.NewRequest(http.MethodGet, url1, nil)
	c := http.Client{}
	r, _ := c.Do(req)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	u := User{}
	json.Unmarshal(body, &u)

	return u, nil
}

func Post1() (*http.Response, error) {
	j, _ := json.Marshal(map[string]string{"code": "200", "status": "ok"})
	r, err := http.Post(url2, "application/json", bytes.NewBuffer(j))

	return r, err
}

func Post2() (*http.Response, error) {
	j, _ := json.Marshal(map[string]string{"code": "200", "status": "ok, 2nd time"})
	req, _ := http.NewRequest(http.MethodPost, url2, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	c := http.Client{}
	r, err := c.Do(req)

	return r, err
}
