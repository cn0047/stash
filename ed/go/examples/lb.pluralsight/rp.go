package main

import (
	"io"
	"net/http"
	"net/url"
)

var (
	appservers   = []string{}
	currentIndex = 0
	client       = http.Client{Transport: &transport}
)

func processRequests() {
	for {
		select {
		case request := <-requestCh:
			println("request")
			if len(appservers) == 0 {
				request.w.WriteHeader(http.StatusInternalServerError)
				request.w.Write([]byte("No app servers found"))
				request.doneCh <- struct{}{}
				continue
			}
			currentIndex++
			if currentIndex == len(appservers) {
				currentIndex = 0
			}
			host := appservers[currentIndex]
			go processRequest(host, request)
		}
	}
}

func processRequest(host string, request *webRequest) {
	hostURL, _ := url.Parse(request.r.URL.String())
	hostURL.Scheme = "https"
	hostURL.Host = host
	println(host)
	println(hostURL.String())
	req, _ := http.NewRequest(request.r.Method, hostURL.String(), request.r.Body)
	for k, v := range request.r.Header {
		values := ""
		for _, headerValue := range v {
			values += headerValue + " "
		}
		req.Header.Add(k, values)
	}

	resp, err := client.Do(req)

	if err != nil {
		request.w.WriteHeader(http.StatusInternalServerError)
		request.doneCh <- struct{}{}
		return
	}

	for k, v := range resp.Header {
		values := ""
		for _, headerValue := range v {
			values += headerValue + " "
		}
		request.w.Header().Add(k, values)
	}
	io.Copy(request.w, resp.Body)

	request.doneCh <- struct{}{}
}
