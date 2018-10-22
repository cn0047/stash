package go_app

import (
  "cloud.google.com/go/storage"
  "fmt"
  "golang.org/x/net/context"
  "google.golang.org/appengine" // Required external App Engine library
  "google.golang.org/appengine/file"
  "google.golang.org/appengine/urlfetch"
  "io"
  "io/ioutil"
  "net/http"
  "os"
)

func init() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/x", xHandler)

  appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  msg := "ago - ok"
  w.Write([]byte(msg))
}

func xHandler(w http.ResponseWriter, r *http.Request) {
  //downloadFile(w, r)
  downloadToGCStorage(w, r)
  //writeToGCStorage(w, r)
  readFromGCStorage(w, r)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)

  filePath := "downloaded.data"

  out, err := os.Create(filePath)
  // open /tmp/x.html: operation not permitted
  // open x.html: no such file or directory
  if err != nil {
    panic("[1] " + err.Error())
  }
  defer out.Close()

  resp := getResp(ctx)
  defer resp.Body.Close()

  _, err = io.Copy(out, resp.Body)
  // permission denied
  if err != nil {
    panic("[3] " + err.Error())
    return
  }
}

func getResp(ctx context.Context) *http.Response {
  url := "https://ago-dot-itisgnp.appspot.com"
  client := urlfetch.Client(ctx)
  resp, err := client.Get(url)
  if err != nil {
    panic("[getResp] " + err.Error())
  }

  return resp
}

func downloadToGCStorage(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  fileName := "test/downloaded.data"

  resp := getResp(ctx)
  defer resp.Body.Close()

  b, c := getBucket(ctx)
  defer c.Close()

  wc := b.Object(fileName).NewWriter(ctx)
  wc.ContentType = "text/plain"

  _, err := io.Copy(wc, resp.Body)
  if err != nil {
    panic("[1] " + err.Error())
    return
  }

  if err := wc.Close(); err != nil {
    panic(fmt.Errorf("[2] createFile: unable to close file %q: %v", fileName, err))
  }
}

func getBucket(ctx context.Context) (*storage.BucketHandle, *storage.Client) {
  bucketName, err := file.DefaultBucketName(ctx)
  if err != nil {
    panic("[getBucket 1] " + err.Error())
  }

  client, err := storage.NewClient(ctx)
  if err != nil {
    panic("[getBucket 2] " + err.Error())
  }

  bucket := client.Bucket(bucketName)

  return bucket, client
}

func writeToGCStorage(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  fileName := "test/downloaded.data"

  b, c := getBucket(ctx)
  defer c.Close()

  wc := b.Object(fileName).NewWriter(ctx)
  wc.ContentType = "text/plain"
  wc.Metadata = map[string]string{
    "x-goog-meta-foo": "foo",
    "x-goog-meta-bar": "bar",
  }

  if _, err := wc.Write([]byte("it works!\n")); err != nil {
    panic(fmt.Errorf("[3] createFile: unable to write data to file %q: %v", fileName, err))
  }
  if err := wc.Close(); err != nil {
    panic(fmt.Errorf("[5] createFile: unable to close file %q: %v", fileName, err))
  }
}

func readFromGCStorage(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  fileName := "test/downloaded.data"

  b, c := getBucket(ctx)
  defer c.Close()

  rc, err := b.Object(fileName).NewReader(ctx)
  if err != nil {
    panic("[3] " + err.Error())
  }
  defer rc.Close()

  data, err := ioutil.ReadAll(rc)
  if err != nil {
    panic("[4] " + err.Error())
  }

  fmt.Fprintf(w, "<br>%s", data)
}
