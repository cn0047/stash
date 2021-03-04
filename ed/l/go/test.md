Test
-

## Test

````sh
# Local directory mode (caching is disabled)
test -v

# List mode (caches successful package test)
go test math
gotest ./...
go test .

# If run tests with non-test flags - the result is not cached.
# The idiomatic way to disable test caching explicitly is to use -count=1.
go test . -count=1

testdata # directory which is ignored by go (magic)

# to get more test info
go test -gcflags
go build -gcflags '-m'
go build -gcflags '-m -m' ./main.go

go test -short
go test -v
go test -cover
go test -race
go test -cpu
go test -json
go test -timeout 2s
go test -list
go test -parallel

goapp test -v transactional/users/service -run TestGetServiceAccountsForAdmin

# coverage
go test -coverpkg mylib, fmt mylib
# example
GOPATH=$PWD/ed/go/examples/install
cd $GOPATH/lib ;\
  go test -v ;\
  go test -cover -coverprofile=coverage.out ;\
  go tool cover -html=coverage.out -o=coverage.html ;\
  open coverage.html

# ✅
go test -count=1 -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o=coverage.html && rm coverage.out
open coverage.html
````

````golang
func TestXxx()      // test
func BenchmarkXxx() // benchmark
func ExampleXxx()   // example (godoc)
````

````golang
import (
  "testing"
)
func TestX(ts *testing.T) { // test suite
  ts.Run("testCase", func(t *testing.T) {
  })
}

t.Log("GivenTheNeedToTestTheSendJSONEndpoint.") {
  //code
}

t.Run("TestCase", func(t *testing.T) {})
t.Errorf("Got: %v, want: %v", a, e) // ✅ & CONTINUE
t.Error("Msg")                      // mark test as failed  & CONTINUE further test execution
t.Skip("Msg")                       // mark test as skipped & STOP further test execution
t.SkipNow()                         // mark test as skipped & STOP further test execution
t.Skipped()                         // is test skipped
t.Fail()                            // mark test as failed  & CONTINUE further test execution
t.Failed()                          // is test failed
t.FailNow()                         // mark test as failed  & STOP further test execution
t.Fatal()                           // mark test as failed  & STOP further test execution
t.Fatalf("Msg")                     // mark test as failed  & STOP further test execution

b.RunParallel(func(pb *testing.PB) {})
if testing.Short() { // go test -short }
````

Helpers:

````golang
{
  httpReq, httpBody, err1 := NewUploadRequest(url, field, filePath)
  err2 := httpBody.WriteField("video_name", "v.mov")
  err4 := httpBody.Close()
}

func NewUploadRequest(
  url string, fieldName string, pathToFile string,
) (httpReq *http.Request, httpBodyWriter *multipart.Writer, err error) {
  f, err1 := os.Open(pathToFile)
  if err1 != nil {
    return nil, nil, fmt.Errorf(
      "failed to open file: %v, error: %v", pathToFile, err1,
    )
  }

  fStat, err3 := f.Stat()
  if err3 != nil {
    return nil, nil, fmt.Errorf("failed to get file stat, error: %v", err3)
  }

  httpBody := &bytes.Buffer{}

  httpBodyWriter = multipart.NewWriter(httpBody)
  part, err4 := httpBodyWriter.CreateFormFile(fieldName, filepath.Base(f.Name()))
  if err4 != nil {
    return nil, nil, fmt.Errorf("failed to create FormFile, error: %v", err4)
  }

  _, err5 := io.Copy(part, f)
  if err5 != nil {
    return nil, nil, fmt.Errorf(
      "failed to copy file into FormFile, error: %v", err5,
    )
  }

  err2 := f.Close()
  if err2 != nil {
    return nil, nil, fmt.Errorf("failed to close file, error: %v", err2)
  }

  httpReq, err6 := http.NewRequest(http.MethodPost, url, httpBody)
  if err6 != nil {
    return nil, nil, fmt.Errorf(
      "failed to create http request, error: %v", err6,
    )
  }

  contentLength := strconv.FormatInt(fStat.Size(), 10)
  httpReq.Header.Set(echo.HeaderContentLength, contentLength)
  httpReq.Header.Set(echo.HeaderContentType, httpBodyWriter.FormDataContentType())

  return httpReq, httpBodyWriter, nil
}
````

## Bench

````golang
b.ResetTimer()     // to reset time and start benchmark (to skip preparations steps)
b.SetParallelism() // sets the number of goroutines used by RunParallel
b.ReportAllocs()   // enables malloc statistics
````

Bench output:

````sh
goos: linux
goarch: amd64
BenchmarkFibComplete-4     1000000        2185 ns/op      3911 B/op        0 allocs/op
PASS
ok    _/gh/ed/go/examples/bench 2.446s
````

Means that the loop ran 100000000 times at a speed of 2185 ns (nanosecond) per loop.

## Asserts

````golang
func assert(t *testing.T, actual interface{}, expected interface{}) {
func assert(t *testing.T, expected interface{}, actual interface{}) {
  if actual != expected {
    t.Errorf("Got: %#v, want: %#v", actual, expected)
  }
}

func assertNil(t *testing.T, v interface{}) {
  if v != nil {
    t.Errorf("Got: %#v, want: nil", v)
  }
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
  if actual != expected {
    t.Errorf("Got: %#v, want: %#v", actual, expected)
  }
}

func assertContains(t *testing.T, haystack []byte, needle string) {
  s := string(haystack)
  if !strings.Contains(s, needle) {
    t.Errorf("Got: %v which doesn't contain: %v", s, needle)
  }
}

func assertContains(t *testing.T, str string, substr string) {
  if !strings.Contains(str, substr) {
    t.Errorf("Got: %v which doesn't contain: %v", str, substr)
  }
}
````
