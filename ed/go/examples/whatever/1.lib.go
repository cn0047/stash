package lib

import (
	"strconv"
)

str := strconv.FormatInt(int64, 10)                     // int64 -> str
str := strconv.FormatInt(int64(int32), 10)              // int32 -> str
int64, _ := strconv.ParseInt(str, 10, 32); int32(int64) // str   -> int32

if !regexp.MustCompile(`^[\d]+$`).MatchString(ds) {}


// ReverseInt32 returns 321 for 123 etc.
func ReverseInt32(n int32) int32 {
	s := strconv.FormatInt(int64(n), 10)

	rs := ""
	for i := 0; i < len(s); i++ {
		rs = string(s[i]) + rs
	}

	r, _ := strconv.ParseInt(rs, 10, 32)

	return int32(r)
}

// SwapInIntSlice swaps i element with j element in array a.
func SwapInIntSlice(a []int, i int, j int) []int {
	v1 := a[i]
	v2 := a[j]

	a = append(a[:i], append([]int{v2}, a[i+1:]...)...)
	a = append(a[:j], append([]int{v1}, a[j+1:]...)...)

	return a
}

// SwapInStr swaps i char with j char in string s.
func SwapInStr(s string, i int, j int) string {
	c1 := s[i]
	c2 := s[j]

	s = s[:i] + string(c2) + s[i+1:]
	s = s[:j] + string(c1) + s[j+1:]

	return s
}

// LeftPadding adds char x to left part of string s up to final length n.
func LeftPadding(s string, x string, n int) string {
	if len(s) >= n {
		return s
	}

	m := n - len(s)
	res := ""
	for i := 0; i < m; i++ {
		res += x
	}

	res += s

	return res
}

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

	resBody, err := ioutil.ReadAll(res.Body)
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
