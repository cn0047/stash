package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func transmitRequest(w http.ResponseWriter, r *http.Request, targetURI string) {
	req, err := http.NewRequest(r.Method, targetURI, r.Body)
	if err != nil {
		err(w, err, "failed to create new request")
		return
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", r.Header.Get("authorization"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err(w, err, "failed to create new client")
		return
	}

	// resBody, err := ioutil.ReadAll(res.Body)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		err(w, err, "failed to read response body")
		return
	}
	fmt.Printf("proxy response: %s", resBody)

	if _, err := io.Copy(w, bytes.NewReader(resBody)); err != nil {
		err(w, err, "failed to copy response")
		return
	}
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
}

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	target, err := url.Parse(h.config.ExternalService.BaseURL)
	if err != nil {
		h.SendInternalError(w, err, "failed to parse url")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Add("Access-Control-Allow-Origin", h.config.CORS.AllowOrigin)
		return nil
	}

	r.URL.Scheme = target.Scheme
	r.URL.Host = target.Host
	r.URL.Path = "/x"
	r.Host = target.Host

	proxy.ServeHTTP(w, r)
}

func rtl(data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
	}
	_, err2 := http.Post("https://realtimelog.herokuapp.com:443/rkc8q6llprn", "application/json", bytes.NewBuffer(j))
	if err2 != nil {
	}
}

func GetPi() float32 {
	return 3.14
}
