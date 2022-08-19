package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type RequestData1 struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

type RequestData2 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// open http://localhost:8080/file
func main() {
	http.HandleFunc("/get", get1)
	http.HandleFunc("/post", post2)
	http.HandleFunc("/file", file)
	http.HandleFunc("/uploadFile", uploadFile)
	http.ListenAndServe(":8080", nil)
}

// @see curl 'http://localhost:8080/get?f=foo&b=bar'
func get1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Printf("Query: %+v", query)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}

func post1(w http.ResponseWriter, r *http.Request) {
	// body1, err := ioutil.ReadAll(r.Body)
	body1, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rd1 := RequestData1{}
	err = json.Unmarshal(body1, &rd1)
	if err != nil {
		panic(err)
	}

	// body2, err := ioutil.ReadAll(r.Body)
	body2, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rd2 := RequestData2{}
	err = json.Unmarshal(body2, &rd2)
	if err != nil {
		panic(err) // 2018/07/19 17:56:01 http: panic serving [::1]:54581: unexpected end of JSON input
	}

	fmt.Printf("rd1: %+v \nrd2: %+v", rd1, rd2)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}

// @see curl -X POST 'http://localhost:8080/post' -H 'Content-Type: application/json' -d '{"code":"200", "status": "OK", "message": "200 OK"}'
func post2(w http.ResponseWriter, r *http.Request) {
	// body, err := ioutil.ReadAll(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	rd1 := RequestData1{}
	if err = json.Unmarshal(body, &rd1); err != nil {
		panic(err)
	}

	rd2 := RequestData2{}
	err = json.Unmarshal(body, &rd2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("rd1: %+v \nrd2: %+v", rd1, rd2)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}

func file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postFile(w, r)
		return
	}
	getFile(w, r)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<body>
		<form action="file" method="post" enctype="multipart/form-data">
			<input type="text" id="msg" name="msg">
			<input type="file" id="file" name="file">
			<input type="submit" name="submit" value="Upload">
		</form>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func postFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	fmt.Fprintf(w, "File name %s\n", name[0])

	var Buf bytes.Buffer
	_, err3 := io.Copy(&Buf, file)
	if err3 != nil {
		panic(err3)
	}

	contents := Buf.String()
	fmt.Fprintf(w, "%v", contents)

	Buf.Reset()
}

// @example: go run ed/go/examples/http/http.server.go
// @example: curl localhost:8080/uploadFile
func uploadFile(w http.ResponseWriter, r *http.Request) {
	// echo 'It works!)' > /tmp/test.txt
	filename := "/tmp/test.txt"
	// image:
	// filename := "/tmp/i.png"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		panic(err)
	}

	io.Copy(part, file)

	writer.WriteField("msg", "foo")
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080/file", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Fprintf(w, "response Status: %v", res.Status)
	// b, _ := ioutil.ReadAll(res.Body)
	b, _ := io.ReadAll(res.Body)
	fmt.Fprintf(w, "response Body: %s", b)
	// for images:
	// fmt.Fprintf(w, "%s", b)
}
